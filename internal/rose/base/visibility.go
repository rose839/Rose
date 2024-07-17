package base

// Visibility用于表示符号的可见性
type Visibility int

const (
	Private   Visibility = iota // 私有
	Protected                   // 保护
	Public                      // 公有
)
