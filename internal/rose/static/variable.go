package static

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/rose839/Rose/internal/rose/base"
)

// Variable用于表示变量
type Variable struct {
	*BaseSymbol
	t base.Type // 变量的类型
}

// NewVariable创建一个新的变量
func NewVariable(name string, scope base.Scope, ctx antlr.ParserRuleContext) *Variable {
	return &Variable{
		BaseSymbol: NewBaseSymbol(name, scope, ctx),
		t:          nil,
	}
}
