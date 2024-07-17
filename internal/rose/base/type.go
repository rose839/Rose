package base

// Type表示Rose语言中的类型
type Type interface {
	GetName() string    // 获取类型名称
	GetScope() Scope    // 获取类型所在的作用域
	IsType(t Type) bool // 判断是否为某种类型
}
