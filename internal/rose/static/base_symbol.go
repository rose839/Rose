package static

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/rose839/Rose/internal/rose/base"
)

var _ base.Symbol = (*BaseSymbol)(nil)

// Symbol用于表示符号(包括变量、函数、类等)
type BaseSymbol struct {
	Name       string                  // 符号名称
	Scope      base.Scope              // 符号所在的作用域
	Visibility base.Visibility         // 符号的可见性
	Ctx        antlr.ParserRuleContext // 符号关联的AST节点
}

func NewBaseSymbol(name string, scope base.Scope, ctx antlr.ParserRuleContext) *BaseSymbol {
	return &BaseSymbol{
		Name:  name,
		Scope: scope,
		Ctx:   ctx,
	}
}

func (s *BaseSymbol) GetName() string {
	return s.Name
}

func (s *BaseSymbol) GetScope() base.Scope {
	return s.Scope
}

func (s *BaseSymbol) SetScope(scope base.Scope) {
	s.Scope = scope
}
