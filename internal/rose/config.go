package rose

import (
	"errors"
	"os"
)

type Config struct {
	Code string // 脚本代码
}

// CreateConfigFromOptions用于从命令行选项创建Config对象
func CreateConfigFromOptions(opts *Options) (*Config, error) {
	if opts.Code != "" {
		return &Config{Code: opts.Code}, nil
	}

	if opts.File != "" {
		// 读取文件内容
		content, err := os.ReadFile(opts.File)
		if err != nil {
			return nil, err
		}

		return &Config{Code: string(content)}, nil
	}

	return nil, errors.New("either -f or -c must be provided")
}
