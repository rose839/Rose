package errors

import (
	"fmt"
	"runtime"
)

// Frame表示堆栈帧内的程序计数器。由于历史原因，如果Frame被解释为uintptr,
// 其值表示程序计数器+1。
type Frame uintptr

// 返回该帧的程序计数器
func (f Frame) pc() uintptr {
	return uintptr(f) - 1
}

// file方法返回Frame对应的文件名
func (f Frame) file() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}

	file, _ := fn.FileLine(f.pc())
	return file
}

// line方法返回Frame对应的行号
func (f Frame) line() int {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return 0
	}
	_, line := fn.FileLine(f.pc())
	return line
}

// name方法返回Frame所在的函数名
func (f Frame) name() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}
	return fn.Name()
}

// 格式化输出该frame:
//
//		%s 文件名
//		%d 行号
//		%n 函数名
//		%v 等同于%s:%d
//	 $+s 函数名\n\t文件相对路径
//	 $+v 等同于%+s:%d
func (f Frame) Format(s fmt.State, verb rune) {

}
