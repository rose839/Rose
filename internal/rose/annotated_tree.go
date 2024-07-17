package rose

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/rose839/Rose/internal/rose/base"
)

// AnnotatedTree用于表示语义分析后的语法树扩展结构。
// 语义分析的结果都放在该结构中，并与AST的节点建立关联。包括：
//
//	1.类型信息，包括基本类型和用户自定义类型；
//	2.变量和函数调用的消解；
//	3.作用域Scope。
type AnnotatedTree struct {
	AST         *antlr.ParseTree                        // 原始语法树
	Node2Symbol map[antlr.ParserRuleContext]base.Symbol // AST节点对应的符号
	Node2Scope  map[antlr.ParserRuleContext]*base.Scope // AST节点对应的作用域
	Node2Type   map[antlr.ParserRuleContext]base.Type   // AST节点推断出的类型
}
