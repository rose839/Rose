package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Command是子命令数据结构
type Command struct {
	usage    string         // 子命令使用信息
	desc     string         // 子命令描述
	options  CliOptions     // 子命令的命令行选项
	commands []*Command     // 子命令的子命令集合
	runFunc  RunCommandFunc // 子命令的执行函数
}

// CommandOption用于Command实例的选项模式
type CommandOption func(*Command)

// RunCommandFunc用于定义Command实例的启动函数
type RunCommandFunc func(args []string) error

// WithCommandOptions设置Command实例的命令行选项
func WithCommandOptions(opt CliOptions) CommandOption {
	return func(c *Command) {
		c.options = opt
	}
}

// WithCommandRunFunc设置Command实例的命令行选项的执行函数
func WithCommandRunFunc(run RunCommandFunc) CommandOption {
	return func(c *Command) {
		c.runFunc = run
	}
}

// NewCommand创建一个新的Command实例
func NewCommand(usage string, desc string, opts ...CommandOption) *Command {
	c := &Command{
		usage: usage,
		desc:  desc,
	}

	for _, o := range opts {
		o(c)
	}

	return c
}

// AddCommand添加子命令到Command实例
func (c *Command) AddCommand(cmd *Command) {
	c.commands = append(c.commands, cmd)
}

// AddCommands添加多个子命令到Command实例
func (c *Command) AddCommands(cmds ...*Command) {
	c.commands = append(c.commands, cmds...)
}

// cobraCommand将Command实例转换为cobra命令
func (c *Command) cobraCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   c.usage,
		Short: c.desc,
	}
	cmd.SetOutput(os.Stdout)
	cmd.Flags().SortFlags = true

	// 向Command实例添加子命令
	if len(c.commands) > 0 {
		for _, command := range c.commands {
			cmd.AddCommand(command.cobraCommand())
		}
	}

	// 设置子命令执行函数
	if c.runFunc != nil {
		cmd.Run = c.runCommand
	}

	// 设置子命令的flags
	if c.options != nil {
		for _, f := range c.options.Flags().FlagSets {
			cmd.Flags().AddFlagSet(f)
		}
	}

	// 添加help flag到该子命令
	addHelpCommandFlag(c.usage, cmd.Flags())

	return cmd
}

// runCommand为cobra命令的执行函数
func (c *Command) runCommand(cmd *cobra.Command, args []string) {
	if c.runFunc != nil {
		if err := c.runFunc(args); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}
}
