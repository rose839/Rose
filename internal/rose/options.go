package rose

import (
	"os"

	"github.com/pkg/errors"
	"github.com/rose839/Rose/pkg/app"
)

// Options需要实现app.CliOptions接口
var _ app.CliOptions = &Options{}

// Options用于定义rose的命令行选项
type Options struct {
	File string `json:"file" mapstructure:"file"` // 脚本文件路径
	Code string `json:"code" mapstructure:"code"` // 脚本代码
}

// NewOptions创建一个新的Options对象
func NewOptions() *Options {
	return &Options{}
}

// Flags返回Options对象的命令行选项
func (o *Options) Flags() (fss app.NamedFlagSets) {
	flagSet := fss.FlagSet("generic")
	flagSet.StringVarP(&o.File, "file", "f", "", "Parse and execute <file>")
	flagSet.StringVarP(&o.Code, "code", "c", "", "Run Rose <code>")
	return fss
}

// Validate用于验证Options对象
func (o *Options) Validate() error {
	if o.File != "" && o.Code != "" {
		return errors.New("only one of -f or -c can be provided")
	}

	if o.File == "" && o.Code == "" {
		return errors.New("either -f or -c must be provided")
	}

	// 判断文件是否存在
	if o.File != "" {
		fileInfo, err := os.Stat(o.File)
		if err != nil {
			return errors.Wrap(err, "Could not open input file: "+o.File)
		}

		if fileInfo.IsDir() {
			return errors.New("Input file is a directory: " + o.File)
		}
	}

	return nil
}
