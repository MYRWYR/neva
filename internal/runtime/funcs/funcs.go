package funcs

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nevalang/neva/internal/runtime"
)

func read(ctx context.Context, io runtime.FuncIO) (func(), error) {
	sig, err := io.In.Port("sig")
	if err != nil {
		return nil, err
	}
	vout, err := io.Out.Port("v")
	if err != nil {
		return nil, err
	}
	return func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			select {
			case <-ctx.Done():
				return
			case <-sig:
				select {
				case <-ctx.Done():
					return
				default:
					bb, _, err := reader.ReadLine()
					if err != nil {
						panic(err)
					}
					select {
					case <-ctx.Done():
						return
					case vout <- runtime.NewStrMsg(string(bb)):
					}
				}
			}
		}
	}, nil
}

func print(ctx context.Context, io runtime.FuncIO) (func(), error) { //nolint:predeclared
	vin, err := io.In.Port("v")
	if err != nil {
		return nil, err
	}
	vout, err := io.Out.Port("v")
	if err != nil {
		return nil, err
	}
	return func() {
		for {
			select {
			case <-ctx.Done():
				return
			case v := <-vin:
				select {
				case <-ctx.Done():
					return
				default:
					fmt.Println(v.String())
					select {
					case <-ctx.Done():
						return
					case vout <- v:
					}
				}
			}
		}
	}, nil
}

func lock(ctx context.Context, io runtime.FuncIO) (func(), error) {
	vin, err := io.In.Port("v")
	if err != nil {
		return nil, err
	}
	sig, err := io.In.Port("sig")
	if err != nil {
		return nil, err
	}
	vout, err := io.Out.Port("v")
	if err != nil {
		return nil, err
	}
	return func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-sig:
				select {
				case <-ctx.Done():
					return
				case v := <-vin:
					select {
					case <-ctx.Done():
						return
					case vout <- v:
					}
				}
			}
		}
	}, nil
}

func constant(ctx context.Context, io runtime.FuncIO) (func(), error) {
	msg := ctx.Value("msg")
	if msg == nil {
		return nil, errors.New("ctx msg not found")
	}

	v, ok := msg.(runtime.Msg)
	if !ok {
		return nil, errors.New("ctx value is not runtime message")
	}

	vout, err := io.Out.Port("v")
	if err != nil {
		return nil, err
	}

	return func() {
		for {
			select {
			case <-ctx.Done():
				return
			case vout <- v:
			}
		}
	}, nil
}

func void(ctx context.Context, io runtime.FuncIO) (func(), error) {
	vin, err := io.In.Port("v")
	if err != nil {
		return nil, err
	}

	return func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-vin:
			}
		}
	}, nil
}

func add(ctx context.Context, io runtime.FuncIO) (func(), error) {
	a, err := io.In.Port("a")
	if err != nil {
		return nil, err
	}
	b, err := io.In.Port("b")
	if err != nil {
		return nil, err
	}
	vout, err := io.Out.Port("v")
	if err != nil {
		return nil, err
	}

	handler := func(a, b runtime.Msg) runtime.Msg {
		return runtime.NewIntMsg(a.Int() + b.Int())
	}

	return func() {
		for {
			select {
			case <-ctx.Done():
				return
			case v1 := <-a:
				select {
				case <-ctx.Done():
					return
				case v2 := <-b:
					select {
					case <-ctx.Done():
						return
					case vout <- handler(v1, v2):
					}
				}
			}
		}
	}, nil
}

func parseInt(ctx context.Context, io runtime.FuncIO) (func(), error) {
	vin, err := io.In.Port("v")
	if err != nil {
		return nil, err
	}

	vout, err := io.Out.Port("v")
	if err != nil {
		return nil, err
	}

	errout, err := io.Out.Port("err")
	if err != nil {
		return nil, err
	}

	parseFunc := func(str string) (runtime.Msg, error) {
		v, err := strconv.Atoi(str)
		if err != nil {
			return nil, errors.New(strings.TrimPrefix(err.Error(), "strconv.Atoi: "))
		}
		return runtime.NewIntMsg(int64(v)), nil
	}

	return func() {
		for {
			select {
			case <-ctx.Done():
				return
			case str := <-vin:
				v, err := parseFunc(str.Str())
				if err != nil {
					select {
					case <-ctx.Done():
						return
					case errout <- runtime.NewStrMsg(err.Error()):
					}
					continue
				}
				select {
				case <-ctx.Done():
					return
				case vout <- v:
				}
			}
		}
	}, nil
}

func Registry() map[string]runtime.Func {
	return map[string]runtime.Func{
		"Read":     read,
		"Print":    print,
		"Lock":     lock,
		"Const":    constant,
		"Add":      add,
		"ParseInt": parseInt,
		"Void":     void,
	}
}
