package static

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/rose839/Rose/internal/rose/base"
)

// Function实现了FunctionType接口
var _ FunctionType = (*Function)(nil)

// Function用于表示函数定义(函数定义是函数类型，但不能用来声明变量类型)
type Function struct {
	*BaseSymbol             // 函数属于符号
	*BaseScope              // 函数会创建新的作用域
	Parameters  []*Variable // 函数的参数
	ReturnType  base.Type   // 函数的返回类型
	ClosureVars []*Variable // 自由变量, 即闭包所引用的外部变量
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
	for _, p := range f.Parameters {
		types = append(types, p.t)
	}
	return types
}

func (f *Function) MatchParamTypes(types []base.Type) bool {
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

func (f *Function) IsType(t base.Type) bool {
	target, ok := t.(FunctionType)
	if !ok {
		return false
	}

	// 检查参数类型是否匹配
	if !f.MatchParamTypes(target.GetParamTypes()) {
		return false
	}

	// 检查返回类型是否匹配
	if !f.ReturnType.IsType(target.GetReturnType()) {
		return false
	}

	return true
}

// IsMethod判断函数是否为方法
func (f *Function) IsMethod() bool {
	scope := f.GetScope()
	if _, ok := scope.(Class); ok {
		return true
	}

	return false
}

// IsConstructor判断函数是否为构造函数
func (f *Function) IsConstructor() bool {
	if f.IsMethod() && f.GetName() == f.GetScope().GetName() {
		return true
	}

	return false
}
