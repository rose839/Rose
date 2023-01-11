package app

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/pflag"
)

// NamedFlagSets用于存储命名flag set
type NamedFlagSets struct {
	Order    []string                  // 记录flag set被创建的顺序
	FlagSets map[string]*pflag.FlagSet // 记录flag set名和flag set的对应关系
}

// FlagSet根据给定的名称获取对应的flag set，不存在时创建新的flag set
func (nfs *NamedFlagSets) FlagSet(name string) *pflag.FlagSet {
	if nfs.FlagSets == nil {
		nfs.FlagSets = make(map[string]*pflag.FlagSet)
	}

	if _, ok := nfs.FlagSets[name]; !ok {
		nfs.FlagSets[name] = pflag.NewFlagSet(name, pflag.ExitOnError)
		nfs.Order = append(nfs.Order, name)
	}

	return nfs.FlagSets[name]
}

// WordSepNormalizeFunc用于将flag名中的"_"替换为"-"
func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		nname := strings.ReplaceAll(name, "_", "-")

		return pflag.NormalizedName(nname)
	}
	return pflag.NormalizedName(name)
}

// InitFlags用于对flag set做一些初始化操作
func InitFlags(flags *pflag.FlagSet) {
	flags.SetNormalizeFunc(WordSepNormalizeFunc)
	flags.AddGoFlagSet(flag.CommandLine)
}

// PrintSections输出格式化的命名flag set
func PrintSections(w io.Writer, fss NamedFlagSets, cols int) {
	for _, name := range fss.Order {
		fs := fss.FlagSets[name]
		if !fs.HasFlags() {
			continue
		}

		wideFS := pflag.NewFlagSet("", pflag.ExitOnError)
		wideFS.AddFlagSet(fs)

		var zzz string
		if cols > 24 {
			zzz = strings.Repeat("z", cols-24)
			wideFS.Int(zzz, 0, strings.Repeat("z", cols-24))
		}

		var buf bytes.Buffer
		fmt.Fprintf(&buf, "\n%s flags:\n\n%s", strings.ToUpper(name[:1])+name[1:], wideFS.FlagUsagesWrapped(cols))

		if cols > 24 {
			i := strings.Index(buf.String(), zzz)
			lines := strings.Split(buf.String()[:i], "\n")
			fmt.Fprint(w, strings.Join(lines[:len(lines)-1], "\n"))
			fmt.Fprintln(w)
		} else {
			fmt.Fprint(w, buf.String())
		}
	}
}
