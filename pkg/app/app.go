// Package app用于表示应用的命令行框架
package app

import "github.com/spf13/cobra"

// App是命令行应用的主数据结构，使用NewApp方法创建
type App struct {
	basename    string     // 应用名
	name        string     // 应用简短描述
	description string     // 应用描述
	options     CliOptions // 命令行选项
	runFunc     RunFunc    // 应用的启动函数
	silence     bool       // 是否设置为silence模式，该模式下应用不会输出应用启动信息、配置信息和版本信息
	noVersion   bool       // 是否设置为不提供-v、--version选项，即不输出版本信息
	noConfig    bool       //
	commands    []*Command
	args        cobra.PositionalArgs // 非选项参数的验证函数
	cmd         *cobra.Command
}

// Options用于App的选项模式
type Option func(*App)

// RunFunc用于定义应用的启动函数
type RunFunc func(basename string) error

// WithRunFunc用于设置app的应用启动函数
func WithRunFunc(runFunc RunFunc) Option {
	return func(app *App) {
		app.runFunc = runFunc
	}
}

// WithDescription用于设置app的应用描述
func WithDescription(desc string) Option {
	return func(a *App) {
		a.description = desc
	}
}

// WithSilence设置app为silence模式
// 到控制台
func WithSilence() Option {
	return func(app *App) {
		app.silence = true
	}
}

// WithNoVersion设置应用不提供-v、--version选项，即不输出版本信息
func WithNoVersion() Option {
	return func(a *App) {
		a.noVersion = true
	}
}

// WithNoConfig设置应用不提供-c、--config选项，即不支持配置文件
func WithNoConfig() Option {
	return func(a *App) {
		a.noConfig = true
	}
}
