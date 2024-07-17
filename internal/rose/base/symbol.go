package base

// Symbol用于表示符号(包括变量、函数、类等)
type Symbol interface {
	GetName() string      // 获取符号名称
	GetScope() Scope      // 获取符号所在的作用域
	SetScope(scope Scope) // 设置符号所在的作用域
}
