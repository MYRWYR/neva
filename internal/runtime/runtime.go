package runtime

import (
	"errors"
	"fmt"

	"github.com/emil14/respect/internal/runtime/program"
)

var ErrPortNotFound = errors.New("port not found")

type (
	Runtime struct { // TODO: hide implementation
		cnctr Connector
	}

	Connector interface {
		ConnectSubnet([]Connection)
		ConnectOperator(string, IO) error
	}
)

func (r Runtime) Run(p program.Program) (IO, error) {
	return r.spawnNode("root", p.Scope, p.RootNodeMeta)
}

func (r Runtime) spawnNode(
	nodeName string,
	scope map[string]program.Component,
	nodeMeta program.WorkerNodeMeta,
) (IO, error) {
	component, ok := scope[nodeMeta.ComponentName]
	if !ok {
		return IO{}, fmt.Errorf("component not found: %s", nodeMeta.ComponentName)
	}

	io := r.nodeIO(nodeMeta)

	if component.Operator != "" {
		if err := r.cnctr.ConnectOperator(component.Operator, io); err != nil {
			return IO{}, fmt.Errorf("connect operator: %w", err)
		}
		return r.patchIO(nodeMeta, io, nodeName), nil
	}

	subnetIO := map[string]IO{
		"in":  {Out: io.In}, // for subnet 'in' node is sender
		"out": {In: io.Out}, // and 'out' is receiver
	}

	if l := len(component.Const); l > 0 {
		out := make(Ports, l)
		for name, cnst := range component.Const {
			addr := program.PortAddr{Node: "const", Port: name}
			out[addr] = r.constOutport(cnst)
		}
		subnetIO["const"] = IO{Out: out}
	}

	for workerNodeName, workerNodeMeta := range component.WorkerNodesMeta {
		nodeIO, err := r.spawnNode(workerNodeName, scope, workerNodeMeta) // <- recursion
		if err != nil {
			return IO{}, err
		}
		subnetIO[workerNodeName] = nodeIO
	}

	cc, err := r.connections(subnetIO, component.Net)
	if err != nil {
		return IO{}, err
	}

	r.cnctr.ConnectSubnet(cc)

	return r.patchIO(nodeMeta, io, nodeName), nil
}

func (r Runtime) constOutport(cnst program.Const) chan Msg {
	var msg Msg

	switch cnst.Type {
	case program.IntType:
		msg = NewIntMsg(cnst.IntValue)
	}

	ch := make(chan Msg)
	go func() {
		for {
			ch <- msg
		}
	}()

	return ch
}

// patchIO replaces "in" and "out" node names with worker name from parent network
func (r Runtime) patchIO(meta program.WorkerNodeMeta, io IO, nodeName string) IO {
	io2 := IO{
		In:  map[program.PortAddr]chan Msg{},
		Out: map[program.PortAddr]chan Msg{},
	}
	for addr, ch := range io.In {
		addr.Node = nodeName
		io2.In[addr] = ch
	}
	for addr, ch := range io.Out {
		addr.Node = nodeName
		io2.Out[addr] = ch
	}
	return io2
}

// connections initializes channels for network.
func (r Runtime) connections(nodesIO map[string]IO, net []program.Connection) ([]Connection, error) {
	cc := make([]Connection, len(net))

	for i, c := range net {
		fromNodeIO, ok := nodesIO[c.From.Node]
		if !ok {
			return nil, fmt.Errorf("not found IO for node %s", c.From.Node)
		}

		sender, ok := fromNodeIO.Out[c.From]
		if !ok {
			return nil, fmt.Errorf("outport %s not found", c.From)
		}

		receivers := make([]Port, len(c.To))
		for j, toAddr := range c.To {
			toNodeIO, ok := nodesIO[toAddr.Node]
			if !ok {
				return nil, fmt.Errorf("io for receiver node not found: %s", toAddr.Node)
			}

			receiver, ok := toNodeIO.In[toAddr]
			if !ok {
				return nil, fmt.Errorf("inport not found %s", toAddr)
			}

			receivers[j] = Port{Ch: receiver, Addr: toAddr}
		}

		cc[i] = Connection{
			From: Port{Ch: sender, Addr: c.From},
			To:   receivers,
		}
	}

	return cc, nil
}

// nodeIO creates channels for node.
func (r Runtime) nodeIO(nodeMeta program.WorkerNodeMeta) IO {
	in := make(map[program.PortAddr]chan Msg)

	for port, slots := range nodeMeta.In {
		addr := program.PortAddr{Port: port, Node: "in"}

		if slots == 0 {
			in[addr] = make(chan Msg)
			continue
		}

		for idx := uint8(0); idx < slots; idx++ {
			addr.Idx = idx
			in[addr] = make(chan Msg)
		}
	}

	outports := make(map[program.PortAddr]chan Msg)

	for port, slots := range nodeMeta.Out {
		addr := program.PortAddr{Port: port, Node: "out"}

		if slots == 0 {
			outports[addr] = make(chan Msg)
			continue
		}

		for idx := uint8(0); idx < slots; idx++ {
			addr.Idx = idx
			outports[addr] = make(chan Msg)
		}
	}

	return IO{in, outports}
}

// Operator returns error if IO doesn't fit.
type Operator func(IO) error

// Connection represents sender-receiver pair.
type Connection struct {
	From Port
	To   []Port
}

// Port maps network address with the real channel.
type Port struct {
	Ch   chan Msg
	Addr program.PortAddr
}

// IO represents node's input and output ports.
type IO struct {
	In, Out Ports
}

// Ports maps network addresses to real channels.
type Ports map[program.PortAddr]chan Msg

// PortGroup returns all port-chanells associated with the given array port name.
func (ports Ports) PortGroup(arrPort string) ([]chan Msg, error) {
	cc := []chan Msg{}

	for addr, ch := range ports {
		if addr.Port == arrPort {
			cc = append(cc, ch)
		}
	}

	if len(cc) == 0 {
		return nil, fmt.Errorf("ErrArrPortNotFound: %s", arrPort)
	}

	return cc, nil
}

func (ports Ports) Port(name string) (chan Msg, error) {
	for addr, ch := range ports {
		if addr.Port != name {
			continue
		}
		if addr.Idx > 0 {
			return nil, fmt.Errorf("unexpected arr port %v", addr)
		}
		return ch, nil
	}

	return nil, fmt.Errorf("%w: want '%s', got: %v", ErrPortNotFound, name, ports)
}

func New(connector Connector) Runtime {
	return Runtime{
		cnctr: connector,
	}
}

type AbsPortAddr struct {
	port Port
	path []string
}
