package static

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/rose839/Rose/internal/rose/base"
)

// Function实现了FunctionType接口
var _ FunctionType = (*Function)(nil)

// Function用于表示函数定义
type Function struct {
	*BaseSymbol             // 函数属于符号
	*BaseScope              // 函数会创建新的作用域
	Parameterss []*Variable // 函数的参数
	ReturnType  base.Type   // 函数的返回类型
	ClosureVars []*Variable // 闭包变量, 即闭包所引用的外部变量
}

func NewFunction(name string, scope base.Scope, ctx antlr.ParserRuleContext) *Function {
	return &Function{
		BaseSymbol: NewBaseSymbol(name, scope, ctx),
		BaseScope:  NewBaseScope(),
	}
}

func (f *Function) GetReturnType() base.Type {
	return f.ReturnType
}

func (f *Function) GetParamTypes() []base.Type {
	var types []base.Type
	for _, p := range f.Parameterss {
		types = append(types, p.t)
	}
	return types
}

func (f *Function) MatchParamTypes(types []base.Type) bool {
	if len(f.Parameterss) != len(types) {
		return false
	}
	for i, p := range f.Parameterss {
		if !p.t.IsType(types[i]) {
			return false
		}
	}
	return true
}

func (f *Function) IsType(t base.Type) bool {
	return false
}
