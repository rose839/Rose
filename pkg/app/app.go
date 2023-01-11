// Package app用于表示应用的命令行框架
package app

import (
	"fmt"
	"os"
	"thunder_fire/pkg/term"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// App是命令行应用的主数据结构，使用NewApp方法创建
type App struct {
	basename    string               // 应用名
	name        string               // 应用简短描述
	description string               // 应用描述
	options     CliOptions           // 命令行选项
	runFunc     RunFunc              // 应用的启动函数
	silence     bool                 // 是否设置为silence模式，该模式下应用不会输出应用启动信息、配置信息和版本信息
	noVersion   bool                 // 是否设置为不提供-v、--version选项，即不输出版本信息
	noConfig    bool                 // 是否设置为不提供-c、--config选项，即不支持配置文件
	commands    []*Command           // 应用定义的子命令，没有子命令时为nil
	args        cobra.PositionalArgs // 非选项参数的验证函数
	cmd         *cobra.Command       // cobra中的根命令
}

// Options用于App的选项模式
type Option func(*App)

// RunFunc用于定义应用的启动函数
type RunFunc func(basename string) error

// WithOptions设置app的命令行选项
func WithOptions(opt CliOptions) Option {
	return func(a *App) {
		a.options = opt
	}
}

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

// WithNoVersion设置app noVersion
func WithNoVersion() Option {
	return func(a *App) {
		a.noVersion = true
	}
}

// WithNoConfig设置app noConfig
func WithNoConfig() Option {
	return func(a *App) {
		a.noConfig = true
	}
}

// WithValidArgs设置非选项参数的自定义验证函数
func WithValidArgs(args cobra.PositionalArgs) Option {
	return func(app *App) {
		app.args = args
	}
}

// WithDefaultValidArgs设置非选项参数的默认验证函数
func WithDefaultValidArgs() Option {
	return func(app *App) {
		app.args = func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		}
	}
}

// NewApp用于创建一个app实例，基于给定的app简短描述、二进制名和其它配置选项
func NewApp(name string, basename string, opts ...Option) *App {
	app := &App{
		name:     name,
		basename: basename,
	}

	for _, o := range opts {
		o(app)
	}

	app.buildCommand()

	return app
}

// buildCommand用于构建命令行框架
func (app *App) buildCommand() {
	// 创建命令行的cobra根命令
	cmd := cobra.Command{
		Use:           app.basename,
		Short:         app.name,
		Long:          app.description,
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          app.args,
	}
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Flags().SortFlags = true
	InitFlags(cmd.Flags())

	// 应用设置了子命令，将子命令添加到根命令上
	if len(app.commands) > 0 {
		for _, command := range app.commands {
			cmd.AddCommand(command.cobraCommand())
		}

		// 存在子命令时则额外添加help子命令
		cmd.SetHelpCommand(helpCommand(app.basename))
	}

	if app.runFunc != nil {
		cmd.RunE = app.runCommand
	}

	// 添加flags到根命令
	var namedFlagSets NamedFlagSets
	if app.options != nil {
		namedFlagSets = app.options.Flags()
		fs := cmd.Flags()
		for _, f := range namedFlagSets.FlagSets {
			fs.AddFlagSet(f)
		}
	}

	// 添加-v、--verison选项
	if !app.noVersion {

	}

	// 添加-c、--config选项
	if !app.noConfig {
		addConfigFlag(app.basename, namedFlagSets.FlagSet("global"))
	}

	// 添加help flag到global命令flag set中
	namedFlagSets.FlagSet("global").BoolP("help", "h", false, fmt.Sprintf("help for %s", cmd.Name()))

	// 添加global flag set到根命令
	cmd.Flags().AddFlagSet(namedFlagSets.FlagSet("global"))

	addCmdTemplate(&cmd, namedFlagSets)
	app.cmd = &cmd
}

// Run用于启动应用
func (app *App) Run() {
	if err := app.cmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

// Command返回app实例的cobra命令
func (app *App) Command() *cobra.Command {
	return app.cmd
}

func (app *App) runCommand(cmd *cobra.Command, args []string) error {
	// 应用指定了version flag，输出版本信息，然后退出程序
	if !app.noVersion {

	}

	// 应用指定了配置文件
	if !app.noConfig {
		// 命令行配置优先级高于配置文件，在此处绑定命令行配置到viper
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}

		// 将完整配置同步到应用的配置结构中
		if err := viper.Unmarshal(app.options); err != nil {
			return err
		}
	}

	// 执行应用指定的配置操作方法
	if app.options != nil {
		if err := app.applyOptionRules(); err != nil {
			return err
		}
	}
	// 运行应用指定的执行函数
	if app.runFunc != nil {
		return app.runFunc(app.basename)
	}

	return nil
}

// applyOptionRules用于执行应用指定的配置操作方法，比如配置补全、验证等
func (app *App) applyOptionRules() error {
	// 配置补全
	if completeableOptions, ok := app.options.(CompleteableOptions); ok {
		if err := completeableOptions.Complete(); err != nil {
			return err
		}
	}

	// 配置验证
	if err := app.options.Validate(); err != nil {
		return err
	}

	// 配置输出
	if _, ok := app.options.(PrintableOptions); ok && !app.silence {

	}

	return nil
}

// AddCommand添加子命令到app实例
func (a *App) AddCommand(cmd *Command) {
	a.commands = append(a.commands, cmd)
}

// AddCommands添加多个子命令到app实例
func (a *App) AddCommands(cmds ...*Command) {
	a.commands = append(a.commands, cmds...)
}

// addCmdTemplate用于添加命令的usage和help函数
func addCmdTemplate(cmd *cobra.Command, namedFlagSets NamedFlagSets) {
	usageFmt := "Usage:\n  %s\n"
	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
		PrintSections(cmd.OutOrStderr(), namedFlagSets, cols)

		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
		PrintSections(cmd.OutOrStdout(), namedFlagSets, cols)
	})
}
