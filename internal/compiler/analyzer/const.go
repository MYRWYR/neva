package analyzer

import (
	"errors"
	"fmt"

	"github.com/nevalang/neva/internal/compiler"
	src "github.com/nevalang/neva/internal/compiler/sourcecode"
)

var (
	ErrConstSeveralValues = errors.New("Constant cannot have several values at once")
)

func (a Analyzer) analyzeConst(
	constant src.Const,
	scope src.Scope,
) (src.Const, *compiler.Error) {
	if constant.Value.Message == nil && constant.Value.Ref == nil {
		return src.Const{}, &compiler.Error{
			Message: "Constant must either have value or reference to another constant",
			Meta:    &constant.Meta,
		}
	}

	if constant.Value.Message == nil { // is ref
		entity, _, err := scope.Entity(*constant.Value.Ref)
		if err != nil {
			return src.Const{}, &compiler.Error{
				Message: err.Error(),
				Meta:    entity.Meta(),
			}
		}

		if entity.Kind != src.ConstEntity {
			return src.Const{}, &compiler.Error{
				Message: fmt.Sprintf("Constant refers to an entity that is not constant: %v", entity.Kind),
				Meta:    entity.Meta(),
			}
		}

		return a.analyzeConst(entity.Const, scope)
	}

	resolvedType, err := a.analyzeTypeExpr(constant.TypeExpr, scope)
	if err != nil {
		return src.Const{}, compiler.Error{
			Message: "Cannot resolve constant type",
			Meta:    &constant.Meta,
		}.Wrap(err)
	}

	if resolvedType.Lit != nil && resolvedType.Lit.Union != nil {
		return src.Const{}, &compiler.Error{
			Message: "Constant cannot have type union",
			Meta:    &constant.Meta,
		}
	}

	var typeExprStrRepr string
	if inst := resolvedType.Inst; inst != nil {
		typeExprStrRepr = inst.Ref.String()
	} else if lit := resolvedType.Lit; lit != nil {
		if lit.Enum != nil {
			typeExprStrRepr = "enum"
		} else if lit.Struct != nil {
			typeExprStrRepr = "struct"
		}
	}

	switch typeExprStrRepr {
	case "bool":
		if constant.Value.Message.Bool == nil {
			return src.Const{}, &compiler.Error{
				Message: fmt.Sprintf("Boolean value is missing in boolean contant: %v", constant),
				Meta:    &constant.Meta,
			}
		}
		if constant.Value.Message.Int != nil ||
			constant.Value.Message.Float != nil ||
			constant.Value.Message.Str != nil ||
			constant.Value.Message.List != nil ||
			constant.Value.Message.DictOrStruct != nil ||
			constant.Value.Message.Union != nil {
			return src.Const{}, &compiler.Error{
				Message: fmt.Sprintf(
					"Constant cannot have several values at once: %v",
					constant,
				),
				Meta: &constant.Meta,
			}
		}
	case "int":
		if constant.Value.Message.Int == nil {
			return src.Const{}, &compiler.Error{
				Message: fmt.Sprintf("Integer value is missing in integer contant: %v", constant),
				Meta:    &constant.Meta,
			}
		}
		if constant.Value.Message.Bool != nil ||
			constant.Value.Message.Float != nil ||
			constant.Value.Message.Str != nil ||
			constant.Value.Message.List != nil ||
			constant.Value.Message.DictOrStruct != nil ||
			constant.Value.Message.Union != nil {
			return src.Const{}, &compiler.Error{
				Message: fmt.Sprintf(
					"Constant cannot have several values at once: %v",
					constant.Value.Message,
				),
				Meta: &constant.Meta,
			}
		}
	case "float":
		// Float is special case. Constant can have float type expression but integer literal.
		// We must pass this case. Desugarer will turn integer literal into float.
		if constant.Value.Message.Float == nil && constant.Value.Message.Int == nil {
			return src.Const{}, &compiler.Error{
				Message: fmt.Sprintf("Float or integer value is missing in float contant: %v", constant),
				Meta:    &constant.Meta,
			}
		}
		if constant.Value.Message.Float != nil && constant.Value.Message.Int != nil {
			return src.Const{}, &compiler.Error{
				Message: fmt.Sprintf(
					"Constant cannot have several values at once: %v",
					constant.Value.Message,
				),
				Meta: &constant.Meta,
			}
		}
		if constant.Value.Message.Bool != nil ||
			constant.Value.Message.Str != nil ||
			constant.Value.Message.List != nil ||
			constant.Value.Message.DictOrStruct != nil ||
			constant.Value.Message.Union != nil {
			return src.Const{}, &compiler.Error{
				Message: fmt.Sprintf(
					"Constant cannot have several values at once: %v",
					constant.Value.Message,
				),
				Meta: &constant.Meta,
			}
		}
	case "string":
		if constant.Value.Message.Str == nil {
			return src.Const{}, &compiler.Error{
				Message: fmt.Sprintf("String value is missing in string contant: %v", constant),
				Meta:    &constant.Meta,
			}
		}
		if constant.Value.Message.Bool != nil ||
			constant.Value.Message.Int != nil ||
			constant.Value.Message.Float != nil ||
			constant.Value.Message.List != nil ||
			constant.Value.Message.DictOrStruct != nil ||
			constant.Value.Message.Union != nil {
			return src.Const{}, &compiler.Error{
				Message: fmt.Sprintf(
					"Constant cannot have several values at once: %v",
					constant.Value.Message,
				),
				Meta: &constant.Meta,
			}
		}
	case "list":
		if constant.Value.Message.List == nil {
			return src.Const{}, &compiler.Error{
				Message: fmt.Sprintf("List value is missing in list contant: %v", constant),
				Meta:    &constant.Meta,
			}
		}
		if constant.Value.Message.Bool != nil ||
			constant.Value.Message.Int != nil ||
			constant.Value.Message.Float != nil ||
			constant.Value.Message.DictOrStruct != nil ||
			constant.Value.Message.Union != nil {
			return src.Const{}, &compiler.Error{
				Message: fmt.Sprintf(
					"Constant cannot have several values at once: %v",
					constant.Value.Message,
				),
				Meta: &constant.Meta,
			}
		}
	case "map", "struct":
		if constant.Value.Message.DictOrStruct == nil {
			return src.Const{}, &compiler.Error{
				Message: fmt.Sprintf("Map or struct value is missing in map or struct contant: %v", constant),
				Meta:    &constant.Meta,
			}
		}
		if constant.Value.Message.Bool != nil ||
			constant.Value.Message.Int != nil ||
			constant.Value.Message.Float != nil ||
			constant.Value.Message.List != nil ||
			constant.Value.Message.Union != nil {
			return src.Const{}, &compiler.Error{
				Message: fmt.Sprintf(
					"Constant cannot have several values at once: %v",
					constant.Value.Message,
				),
				Meta: &constant.Meta,
			}
		}
	case "enum":
		if constant.Value.Message.Union == nil {
			return src.Const{}, &compiler.Error{
				Message: fmt.Sprintf("Enum value is missing in enum contant: %v", constant),
				Meta:    &constant.Meta,
			}
		}
		if constant.Value.Message.Bool != nil ||
			constant.Value.Message.Int != nil ||
			constant.Value.Message.Float != nil ||
			constant.Value.Message.List != nil ||
			constant.Value.Message.DictOrStruct != nil {
			return src.Const{}, &compiler.Error{
				Message: fmt.Sprintf(
					"Constant cannot have several values at once: %v",
					constant.Value.Message,
				),
				Meta: &constant.Meta,
			}
		}
	}

	return src.Const{
		TypeExpr: resolvedType,
		Value:    constant.Value,
		Meta:     constant.Meta,
	}, nil
}
