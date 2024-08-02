package static

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/rose839/Rose/internal/rose/base"
)

// FunctionVarType实现了FunctionType接口
var _ FunctionType = (*FunctionVarType)(nil)

// FunctionVarType表示函数变量类型，用于声明函数类型的变量
type FunctionVarType struct {
	*BaseSymbol             // 函数属于符号
	Parameters  []*Variable // 函数的参数
	ReturnType  base.Type   // 函数的返回类型
}

func NewFunctionVarType(name string, scope base.Scope, ctx antlr.ParserRuleContext) *FunctionVarType {
	return &FunctionVarType{
		BaseSymbol: NewBaseSymbol(name, scope, ctx),
	}
}

func (f *FunctionVarType) GetReturnType() base.Type {
	return f.ReturnType
}

func (f *FunctionVarType) GetParamTypes() []base.Type {
	var types []base.Type

	for _, p := range f.Parameters {
		types = append(types, p.t)
	}

	return types
}

func (f *FunctionVarType) MatchParamTypes(types []base.Type) bool {
	if len(f.Parameters) != len(types) {
		return false
	}

	for i, p := range f.Parameters {
		if !p.t.IsType(types[i]) {
			return false
		}
	}

	return true
}

func (f *FunctionVarType) IsType(t base.Type) bool {
	target, ok := t.(FunctionType)
	if !ok {
		return false
	}

	// 检查参数类型是否匹配
	if !f.MatchParamTypes(target.GetParamTypes()) {
		return false
	}

	// 检查返回类型是否匹配
	return f.ReturnType.IsType(target.GetReturnType())
}
