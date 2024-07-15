package rose

import (
	"github.com/rose839/Rose/pkg/app"
)

const commandDesc = `Rose is a scripting language designed to provide a simple and easy-to-use programming experience.`

// 创建rose app对象
func NewApp(basename string) *app.App {
	opts := NewOptions()
	application := app.NewApp(
		"Rose Language",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithNoConfig(),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *Options) app.RunFunc {
	return func(basename string) error {
		config, err := CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}

		print(config.Code)
		return nil
	}
}
