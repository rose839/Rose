package base

// Scope用于表示作用域
type Scope interface {
	AddSymbol(symbol Symbol)                           // 向作用域中添加符号
	GetName() string                                   // 获取作用域名称
	GetVariable(name string) Symbol                    // 从作用域中获取变量
	GetFunction(name string, paramTypes []Type) Symbol // 从作用域中获取函数
	GetClass(name string) Symbol                       // 从作用域中获取类
}
