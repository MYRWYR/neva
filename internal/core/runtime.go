package core

import (
	"fmt"
)

type Runtime struct {
	env   map[string]Module // TODO move out
	cache map[string]bool
}

const (
	tmpBuf     = 0
	tmpArrSize = 10
)

func (r Runtime) Run(name string) (NodeIO, error) {
	mod, ok := r.env[name]
	if !ok {
		return NodeIO{}, errModNotFound(name)
	}

	modInterface := mod.Interface()

	if nmod, ok := mod.(nativeModule); ok {
		io := r.nodeIO(modInterface.In, modInterface.Out)
		go nmod.impl(io)
		return io, nil
	}

	cmod, ok := mod.(customModule)
	if !ok {
		return NodeIO{}, errUnknownModType(name, mod)
	}

	if !r.cache[name] {
		if err := r.resolveDeps(cmod.deps); err != nil {
			return NodeIO{}, err
		}
		r.cache[name] = true
	}

	nodesIO := make(map[string]NodeIO, 2+len(cmod.workers))

	nodesIO["in"] = r.nodeIO(
		nil,
		OutportsInterface(modInterface.In),
	)
	nodesIO["out"] = r.nodeIO(
		InportsInterface(modInterface.Out),
		nil,
	)

	for w, dep := range cmod.workers {
		io, err := r.Run(dep)
		if err != nil {
			return NodeIO{}, err
		}

		nodesIO[w] = io
	}

	net, err := r.connections(nodesIO, cmod.net)
	if err != nil {
		return NodeIO{}, err
	}

	r.connectAll(net)

	return NodeIO{
		in:  nodeInports(nodesIO["in"].out),
		out: nodeOutports(nodesIO["out"].in),
	}, nil
}

func (r Runtime) connections(io map[string]NodeIO, net []StreamDef) ([]connection, error) {
	rels := make([]connection, len(net))

	for i, rel := range net {
		sender := r.chanByPoint(rel.Sender, io[rel.Sender.NodeName()])

		receivers := make([]chan Msg, len(rel.Recievers))
		for i, receiver := range rel.Recievers {
			receivers[i] = r.chanByPoint(receiver, io[receiver.NodeName()])
		}

		rels[i] = connection{
			Sender:    sender,
			Receivers: receivers,
		}
	}

	return rels, nil
}

func (r Runtime) chanByPoint(point PortPoint, nodeIO NodeIO) chan Msg {
	var result chan Msg

	arrpoint, ok := point.(ArrPortPoint)
	if ok {
		arrport, err := nodeIO.ArrOutport(arrpoint.Port)
		if err != nil {
			panic(err)
		}

		if uint8(len(arrport)) < arrpoint.Index {
			panic("arrport to small")
		}

		result = arrport[arrpoint.Index]
	} else {
		normPoint, ok := point.(NormPortPoint)
		if !ok {
			panic(fmt.Sprintf("%T", point))
		}

		normPort, err := nodeIO.NormOut(normPoint.Port)
		if err != nil {
			panic(err)
		}

		result = normPort
	}

	return result
}

func (r Runtime) resolveDeps(deps Interfaces) error {
	for dep := range deps {
		mod, ok := r.env[dep]
		if !ok {
			return errModNotFound(dep)
		}

		i := mod.Interface()
		err := i.Compare(deps[dep])
		if err != nil {
			return fmt.Errorf("unresolved dependency '%s': %w", dep, err)
		}
	}

	return nil
}

func (r Runtime) nodeIO(in InportsInterface, out OutportsInterface) NodeIO {
	inports := r.Ports(PortsInterface(in))
	outports := r.Ports(PortsInterface(out))

	return NodeIO{
		nodeInports(inports),
		nodeOutports(outports),
	}
}

func (r Runtime) Ports(ports PortsInterface) nodePorts {
	result := make(nodePorts, len(ports))

	for port, typ := range ports {
		if typ.Arr {
			cc := make([]chan Msg, tmpArrSize)
			for i := range cc {
				cc[i] = make(chan Msg)
			}
			result[port] = cc
			continue
		}

		result[port] = make(chan Msg)
	}

	return result
}

func (r Runtime) connectAll(cc []connection) {
	for i := range cc {
		go r.connect(cc[i])
	}
}

func (m Runtime) connect(c connection) {
	for msg := range c.Sender {
		for _, r := range c.Receivers {
			select {
			case r <- msg:
				continue
			default:
				go func() { r <- msg }()
			}
		}
	}
}

type Port interface{}

func New(env map[string]Module) Runtime {
	return Runtime{
		env:   env,
		cache: map[string]bool{},
	}
}
