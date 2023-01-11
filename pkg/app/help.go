package app

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	flagHelp          = "help"
	flagHelpShorthand = "H"
)

// helpCommand用于生成help子命令
func helpCommand(name string) *cobra.Command {
	return &cobra.Command{
		Use:   "help [command]",
		Short: "Help about any command.",
		Long: `Help provides help for any command in the application.
Simply type ` + name + ` help [path to command] for full details.`,
		Run: func(c *cobra.Command, args []string) {
			// 找到需要输出help的子命令
			cmd, _, e := c.Root().Find(args)
			if cmd == nil || e != nil {
				// 没有指定子命令或者对于的子命令不存在，则输出跟命令的usage
				c.Printf("Unknown help topic %#q\n", args)
				_ = c.Root().Usage()
			} else {
				// 如果对于子命令没有--help flag，那么在此处添加该flag
				cmd.InitDefaultHelpFlag()

				// 输出该子命令的帮助信息
				_ = cmd.Help()
			}
		},
	}
}

// addHelpFlag在给定的flag set中添加help flag
func addHelpFlag(name string, fs *pflag.FlagSet) {
	fs.BoolP(flagHelp, flagHelpShorthand, false, fmt.Sprintf("Help for %s.", name))
}

// addHelpCommandFlag在给定的flag set中添加help flag
func addHelpCommandFlag(usage string, fs *pflag.FlagSet) {
	fs.BoolP(
		flagHelp,
		flagHelpShorthand,
		false,
		fmt.Sprintf("Help for the %s command.", strings.Split(usage, " ")[0]),
	)
}
