package static

import "github.com/rose839/Rose/internal/rose/base"

// Class实现了Type接口
var _ base.Type = (*Class)(nil)

// Class表示Rose语言中的类定义
type Class struct {
	*BaseSymbol // 函数属于符号
	*BaseScope  // 函数会创建新的作用域
}
