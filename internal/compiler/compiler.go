package compiler

import (
	"errors"
	"fmt"

	"github.com/emil14/neva/internal/compiler/program"
	runtime "github.com/emil14/neva/internal/runtime/program"
)

var (
	ErrParsing    = errors.New("failed to parse module")
	ErrValidation = errors.New("module is invalid")
	ErrInternal   = errors.New("internal error")
)

type (
	SRCParser interface {
		Module([]byte) (program.Module, error)
		Program(program.Program) ([]byte, error)
	}

	Translator interface {
		Translate(program.Program) (runtime.Program, error)
	}

	Validator interface {
		Validate(program.Module) error // todo validate program
	}

	Coder interface {
		Code(runtime.Program) ([]byte, error)
	}

	PkgDescriptor struct {
		Root    string
		Modules map[string][]byte
	}
)

type Compiler struct {
	srcParser  SRCParser
	validator  Validator
	translator Translator
	coder      Coder
	operators  map[string]program.Operator
}

func (c Compiler) BuildProgram(pkgd PkgDescriptor) (runtime.Program, program.Program, error) {
	scope := c.defaultScope(len(pkgd.Modules))
	for k, v := range pkgd.Modules {
		mod, err := c.compileModule(v)
		if err != nil {
			return runtime.Program{}, program.Program{}, err
		}
		scope[k] = mod
	}

	if err := c.resolveDeps(scope); err != nil {
		return runtime.Program{}, program.Program{}, err
	}

	prog := program.Program{
		Root:  pkgd.Root,
		Scope: scope,
	}

	rprog, err := c.translator.Translate(prog)
	if err != nil {
		return runtime.Program{}, program.Program{}, err
	}

	return rprog, prog, nil
}

func (c Compiler) resolveDeps(scope map[string]program.Component) error {
	for componentName, component := range scope {
		if _, ok := component.(program.Operator); ok {
			continue
		}

		mod, ok := component.(program.Module)
		if !ok {
			return fmt.Errorf("unknown component type")
		}

		for depName, wantIO := range mod.Deps {
			dep := scope[depName]
			if dep == nil {
				return fmt.Errorf("dep %s not found for %s", depName, componentName)
			}

			if err := wantIO.Compare(dep.Interface()); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c Compiler) compileModule(mod []byte) (program.Module, error) {
	parsed, err := c.srcParser.Module(mod)
	if err != nil {
		return program.Module{}, fmt.Errorf("%w: %v", ErrParsing, err)
	}

	if err := c.validator.Validate(parsed); err != nil {
		return program.Module{}, fmt.Errorf("%w: %v", ErrValidation, err)
	}

	return parsed, nil
}

func (c Compiler) compileProgram(
	modules map[string][]byte,
	root string,
	scope map[string]program.Component,
) (program.Program, error) {
	var mod program.Module

	if _, ok := scope[root]; !ok {
		parsed, err := c.srcParser.Module(modules[root])
		if err != nil {
			return program.Program{}, fmt.Errorf("%w: %v", ErrParsing, err)
		}

		if err := c.validator.Validate(parsed); err != nil {
			return program.Program{}, fmt.Errorf("%w: %v", ErrValidation, err)
		}

		mod = parsed
	}

	for name, dep := range mod.Deps {
		prog, err := c.compileProgram(modules, name, scope)
		if err != nil {
			return program.Program{}, err
		}

		subroot, ok := prog.Scope[prog.Root]
		if !ok {
			return program.Program{}, fmt.Errorf("%w", ErrInternal)
		}

		if err := subroot.Interface().Compare(dep); err != nil {
			return program.Program{}, err
		}
	}

	scope[root] = mod

	return program.Program{
		Root:  root,
		Scope: scope,
	}, nil
}

func (c Compiler) defaultScope(padding int) map[string]program.Component {
	m := make(map[string]program.Component, len(c.operators)+padding)
	for i := range c.operators {
		m[c.operators[i].Name] = c.operators[i]
	}
	return m
}

func New(p SRCParser, v Validator, t Translator, c Coder, ops map[string]program.Operator) (Compiler, error) {
	if p == nil || v == nil || t == nil || c == nil || ops == nil {
		return Compiler{}, fmt.Errorf("%w: failed to build compiler", ErrInternal)
	}

	return Compiler{
		srcParser:  p,
		validator:  v,
		translator: t,
		coder:      c,
		operators:  ops,
	}, nil
}

func MustNew(p SRCParser, v Validator, t Translator, c Coder, ops map[string]program.Operator) Compiler {
	cmp, err := New(p, v, t, c, ops)
	if err != nil {
		panic(err)
	}
	return cmp
}
