package static

import "github.com/rose839/Rose/internal/rose/base"

// FunctionType表示Rose语言中的函数类型
type FunctionType interface {
	base.Type
	GetReturnType() base.Type         // 获取函数的返回类型
	GetParamTypes() []base.Type       // 获取函数的参数类型
	MatchParamTypes([]base.Type) bool // 判断参数类型是否匹配
}
