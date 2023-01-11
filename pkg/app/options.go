package app

// CliOptions用于抽象命令行选项
type CliOptions interface {
	Flags() (fss NamedFlagSets) // 获取命名flag sets
	Validate() error            // 验证命令行参数是否合法
}

// CompleteableOptions用于抽象可以被补全的命令行选项
type CompleteableOptions interface {
	Complete() error
}

// PrintableOptions用于抽象自定义打印格式的命令行选项
type PrintableOptions interface {
	String() string
}
