package static

import "github.com/rose839/Rose/internal/rose/base"

var _ base.Scope = (*BaseScope)(nil)

// BaseScope表示作用域基础结构
type BaseScope struct {
	Symbols []base.Symbol // 作用域内的符号集合
}

func NewBaseScope() *BaseScope {
	return &BaseScope{
		Symbols: make([]base.Symbol, 0),
	}
}

// AddSymbol向作用域中添加符号
func (s *BaseScope) AddSymbol(symbol base.Symbol) {
	s.Symbols = append(s.Symbols, symbol)
	symbol.SetScope(s)
}

// GetName获取作用域名称
func (s *BaseScope) GetName() string {
	panic("implement me")
}

// GetVariable从作用域中获取变量
func (s *BaseScope) GetVariable(name string) base.Symbol {
	return nil
}

// GetFunction从作用域中获取函数
func (s *BaseScope) GetFunction(name string, paramTypes []base.Type) base.Symbol {
	return nil
}

// GetClass从作用域中获取类
func (s *BaseScope) GetClass(name string) base.Symbol {
	return nil
}
