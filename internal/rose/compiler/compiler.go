package compiler

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/rose839/Rose/grammar"
)

// Compiler用于编译rose脚本
type Compiler struct {
	Lexer  *grammar.RoseLexer  // 词法分析器
	Parser *grammar.RoseParser // 语法分析器
}

// NewCompiler创建一个新的Compiler对象
func NewCompiler(Code string) *Compiler {
	return &Compiler{}
}

// Compile用于编译rose脚本
func (c *Compiler) Compile(Code string) error {
	// 词法分析
	inputStream := antlr.NewInputStream(Code)
	c.Lexer = grammar.NewRoseLexer(inputStream)

	// 语法分析
	c.Parser = grammar.NewRoseParser(antlr.NewCommonTokenStream(c.Lexer, antlr.LexerDefaultTokenChannel))

	// walker := antlr.NewParseTreeWalker()

	// 开始语法分析
	tree := c.Parser.Prog()

	// 输出语法树
	print(tree.GetText())

	return nil
}
