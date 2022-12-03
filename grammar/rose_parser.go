// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package grammar // Rose
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type RoseParser struct {
	*antlr.BaseParser
}

var roseParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func roseParserInit() {
	staticData := &roseParserStaticData
	staticData.literalNames = []string{
		"", "'abstract'", "'assert'", "'boolean'", "'break'", "'byte'", "'case'",
		"'catch'", "'char'", "'class'", "'const'", "'continue'", "'default'",
		"'do'", "'double'", "'else'", "'enum'", "'extends'", "'final'", "'finally'",
		"'float'", "'for'", "'if'", "'goto'", "'implements'", "'import'", "'instanceof'",
		"'int'", "'interface'", "'long'", "'native'", "'new'", "'package'",
		"'private'", "'protected'", "'public'", "'return'", "'short'", "'static'",
		"'strictfp'", "'super'", "'switch'", "'synchronized'", "'this'", "'throw'",
		"'throws'", "'transient'", "'try'", "'void'", "'volatile'", "'while'",
		"'function'", "'string'", "", "", "", "", "", "", "", "", "", "'null'",
		"'('", "')'", "'{'", "'}'", "'['", "']'", "';'", "','", "'.'", "'='",
		"'>'", "'<'", "'!'", "'~'", "'?'", "':'", "'=='", "'<='", "'>='", "'!='",
		"'&&'", "'||'", "'++'", "'--'", "'+'", "'-'", "'*'", "'/'", "'&'", "'|'",
		"'^'", "'%'", "'+='", "'-='", "'*='", "'/='", "'&='", "'|='", "'^='",
		"'%='", "'<<='", "'>>='", "'>>>='", "'->'", "'::'", "'@'", "'...'",
	}
	staticData.symbolicNames = []string{
		"", "ABSTRACT", "ASSERT", "BOOLEAN", "BREAK", "BYTE", "CASE", "CATCH",
		"CHAR", "CLASS", "CONST", "CONTINUE", "DEFAULT", "DO", "DOUBLE", "ELSE",
		"ENUM", "EXTENDS", "FINAL", "FINALLY", "FLOAT", "FOR", "IF", "GOTO",
		"IMPLEMENTS", "IMPORT", "INSTANCEOF", "INT", "INTERFACE", "LONG", "NATIVE",
		"NEW", "PACKAGE", "PRIVATE", "PROTECTED", "PUBLIC", "RETURN", "SHORT",
		"STATIC", "STRICTFP", "SUPER", "SWITCH", "SYNCHRONIZED", "THIS", "THROW",
		"THROWS", "TRANSIENT", "TRY", "VOID", "VOLATILE", "WHILE", "FUNCTION",
		"STRING", "DECIMAL_LITERAL", "HEX_LITERAL", "OCT_LITERAL", "BINARY_LITERAL",
		"FLOAT_LITERAL", "HEX_FLOAT_LITERAL", "BOOL_LITERAL", "CHAR_LITERAL",
		"STRING_LITERAL", "NULL_LITERAL", "LPAREN", "RPAREN", "LBRACE", "RBRACE",
		"LBRACK", "RBRACK", "SEMI", "COMMA", "DOT", "ASSIGN", "GT", "LT", "BANG",
		"TILDE", "QUESTION", "COLON", "EQUAL", "LE", "GE", "NOTEQUAL", "AND",
		"OR", "INC", "DEC", "ADD", "SUB", "MUL", "DIV", "BITAND", "BITOR", "CARET",
		"MOD", "ADD_ASSIGN", "SUB_ASSIGN", "MUL_ASSIGN", "DIV_ASSIGN", "AND_ASSIGN",
		"OR_ASSIGN", "XOR_ASSIGN", "MOD_ASSIGN", "LSHIFT_ASSIGN", "RSHIFT_ASSIGN",
		"URSHIFT_ASSIGN", "ARROW", "COLONCOLON", "AT", "ELLIPSIS", "WS", "COMMENT",
		"LINE_COMMENT", "IDENTIFIER",
	}
	staticData.ruleNames = []string{
		"classDeclaration", "classBody", "classBodyDeclaration", "memberDeclaration",
		"functionDeclaration", "functionBody", "typeTypeOrVoid", "qualifiedNameList",
		"formalParameters", "formalParameterList", "formalParameter", "lastFormalParameter",
		"variableModifier", "qualifiedName", "fieldDeclaration", "constructorDeclaration",
		"variableDeclarators", "variableDeclarator", "variableDeclaratorId",
		"variableInitializer", "arrayInitializer", "classOrInterfaceType", "typeArgument",
		"literal", "integerLiteral", "floatLiteral", "prog", "block", "blockStatements",
		"blockStatement", "statement", "switchBlockStatementGroup", "switchLabel",
		"forControl", "forInit", "enhancedForControl", "parExpression", "expressionList",
		"functionCall", "expression", "primary", "typeList", "typeType", "functionType",
		"primitiveType", "creator", "superSuffix", "arguments",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 113, 587, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 101, 8, 0, 1, 0, 1, 0, 3, 0, 105,
		8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 111, 8, 1, 10, 1, 12, 1, 114, 9, 1,
		1, 1, 1, 1, 1, 2, 1, 2, 3, 2, 120, 8, 2, 1, 3, 1, 3, 1, 3, 3, 3, 125, 8,
		3, 1, 4, 3, 4, 128, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 134, 8, 4, 10,
		4, 12, 4, 137, 9, 4, 1, 4, 1, 4, 3, 4, 141, 8, 4, 1, 4, 1, 4, 1, 5, 1,
		5, 3, 5, 147, 8, 5, 1, 6, 1, 6, 3, 6, 151, 8, 6, 1, 7, 1, 7, 1, 7, 5, 7,
		156, 8, 7, 10, 7, 12, 7, 159, 9, 7, 1, 8, 1, 8, 3, 8, 163, 8, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 170, 8, 9, 10, 9, 12, 9, 173, 9, 9, 1, 9,
		1, 9, 3, 9, 177, 8, 9, 1, 9, 3, 9, 180, 8, 9, 1, 10, 5, 10, 183, 8, 10,
		10, 10, 12, 10, 186, 9, 10, 1, 10, 1, 10, 1, 10, 1, 11, 5, 11, 192, 8,
		11, 10, 11, 12, 11, 195, 9, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 13, 5, 13, 206, 8, 13, 10, 13, 12, 13, 209, 9, 13, 1,
		14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 218, 8, 15, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 226, 8, 16, 10, 16, 12, 16, 229,
		9, 16, 1, 17, 1, 17, 1, 17, 3, 17, 234, 8, 17, 1, 18, 1, 18, 1, 18, 5,
		18, 239, 8, 18, 10, 18, 12, 18, 242, 9, 18, 1, 19, 1, 19, 3, 19, 246, 8,
		19, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 252, 8, 20, 10, 20, 12, 20, 255,
		9, 20, 1, 20, 3, 20, 258, 8, 20, 3, 20, 260, 8, 20, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 21, 5, 21, 267, 8, 21, 10, 21, 12, 21, 270, 9, 21, 1, 22, 1,
		22, 1, 22, 1, 22, 3, 22, 276, 8, 22, 3, 22, 278, 8, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 3, 23, 286, 8, 23, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 5, 28, 299, 8, 28, 10,
		28, 12, 28, 302, 9, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29,
		310, 8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 318, 8, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30,
		340, 8, 30, 10, 30, 12, 30, 343, 9, 30, 1, 30, 5, 30, 346, 8, 30, 10, 30,
		12, 30, 349, 9, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 355, 8, 30, 1, 30,
		1, 30, 1, 30, 3, 30, 360, 8, 30, 1, 30, 1, 30, 1, 30, 3, 30, 365, 8, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 375, 8,
		30, 1, 31, 4, 31, 378, 8, 31, 11, 31, 12, 31, 379, 1, 31, 4, 31, 383, 8,
		31, 11, 31, 12, 31, 384, 1, 32, 1, 32, 1, 32, 3, 32, 390, 8, 32, 1, 32,
		1, 32, 1, 32, 3, 32, 395, 8, 32, 1, 33, 1, 33, 3, 33, 399, 8, 33, 1, 33,
		1, 33, 3, 33, 403, 8, 33, 1, 33, 1, 33, 3, 33, 407, 8, 33, 3, 33, 409,
		8, 33, 1, 34, 1, 34, 3, 34, 413, 8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 5, 37, 427, 8, 37,
		10, 37, 12, 37, 430, 9, 37, 1, 38, 1, 38, 1, 38, 3, 38, 435, 8, 38, 1,
		38, 1, 38, 1, 38, 1, 38, 3, 38, 441, 8, 38, 1, 38, 1, 38, 1, 38, 1, 38,
		3, 38, 447, 8, 38, 1, 38, 3, 38, 450, 8, 38, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 39, 1, 39, 1, 39, 3, 39, 459, 8, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39,
		475, 8, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39, 513, 8, 39, 1, 39,
		1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 5, 39, 525,
		8, 39, 10, 39, 12, 39, 528, 9, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 3, 40, 538, 8, 40, 1, 41, 1, 41, 1, 41, 5, 41, 543, 8,
		41, 10, 41, 12, 41, 546, 9, 41, 1, 42, 1, 42, 1, 42, 3, 42, 551, 8, 42,
		1, 42, 1, 42, 5, 42, 555, 8, 42, 10, 42, 12, 42, 558, 9, 42, 1, 43, 1,
		43, 1, 43, 1, 43, 3, 43, 564, 8, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45,
		1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 3, 46, 577, 8, 46, 3, 46, 579,
		8, 46, 1, 47, 1, 47, 3, 47, 583, 8, 47, 1, 47, 1, 47, 1, 47, 0, 1, 78,
		48, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
		72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 0, 12, 2, 0, 17, 17, 40,
		40, 1, 0, 53, 56, 1, 0, 57, 58, 1, 0, 85, 88, 1, 0, 75, 76, 2, 0, 89, 90,
		94, 94, 1, 0, 87, 88, 2, 0, 73, 74, 80, 81, 2, 0, 79, 79, 82, 82, 2, 0,
		72, 72, 95, 105, 1, 0, 85, 86, 9, 0, 3, 3, 5, 5, 8, 8, 14, 14, 20, 20,
		27, 27, 29, 29, 37, 37, 52, 52, 644, 0, 96, 1, 0, 0, 0, 2, 108, 1, 0, 0,
		0, 4, 119, 1, 0, 0, 0, 6, 124, 1, 0, 0, 0, 8, 127, 1, 0, 0, 0, 10, 146,
		1, 0, 0, 0, 12, 150, 1, 0, 0, 0, 14, 152, 1, 0, 0, 0, 16, 160, 1, 0, 0,
		0, 18, 179, 1, 0, 0, 0, 20, 184, 1, 0, 0, 0, 22, 193, 1, 0, 0, 0, 24, 200,
		1, 0, 0, 0, 26, 202, 1, 0, 0, 0, 28, 210, 1, 0, 0, 0, 30, 213, 1, 0, 0,
		0, 32, 221, 1, 0, 0, 0, 34, 230, 1, 0, 0, 0, 36, 235, 1, 0, 0, 0, 38, 245,
		1, 0, 0, 0, 40, 247, 1, 0, 0, 0, 42, 263, 1, 0, 0, 0, 44, 277, 1, 0, 0,
		0, 46, 285, 1, 0, 0, 0, 48, 287, 1, 0, 0, 0, 50, 289, 1, 0, 0, 0, 52, 291,
		1, 0, 0, 0, 54, 293, 1, 0, 0, 0, 56, 300, 1, 0, 0, 0, 58, 309, 1, 0, 0,
		0, 60, 374, 1, 0, 0, 0, 62, 377, 1, 0, 0, 0, 64, 394, 1, 0, 0, 0, 66, 408,
		1, 0, 0, 0, 68, 412, 1, 0, 0, 0, 70, 414, 1, 0, 0, 0, 72, 419, 1, 0, 0,
		0, 74, 423, 1, 0, 0, 0, 76, 449, 1, 0, 0, 0, 78, 458, 1, 0, 0, 0, 80, 537,
		1, 0, 0, 0, 82, 539, 1, 0, 0, 0, 84, 550, 1, 0, 0, 0, 86, 559, 1, 0, 0,
		0, 88, 567, 1, 0, 0, 0, 90, 569, 1, 0, 0, 0, 92, 578, 1, 0, 0, 0, 94, 580,
		1, 0, 0, 0, 96, 97, 5, 9, 0, 0, 97, 100, 5, 113, 0, 0, 98, 99, 5, 17, 0,
		0, 99, 101, 3, 84, 42, 0, 100, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101,
		104, 1, 0, 0, 0, 102, 103, 5, 24, 0, 0, 103, 105, 3, 82, 41, 0, 104, 102,
		1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 3, 2,
		1, 0, 107, 1, 1, 0, 0, 0, 108, 112, 5, 65, 0, 0, 109, 111, 3, 4, 2, 0,
		110, 109, 1, 0, 0, 0, 111, 114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112,
		113, 1, 0, 0, 0, 113, 115, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 116,
		5, 66, 0, 0, 116, 3, 1, 0, 0, 0, 117, 120, 5, 69, 0, 0, 118, 120, 3, 6,
		3, 0, 119, 117, 1, 0, 0, 0, 119, 118, 1, 0, 0, 0, 120, 5, 1, 0, 0, 0, 121,
		125, 3, 8, 4, 0, 122, 125, 3, 28, 14, 0, 123, 125, 3, 0, 0, 0, 124, 121,
		1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 123, 1, 0, 0, 0, 125, 7, 1, 0, 0,
		0, 126, 128, 3, 12, 6, 0, 127, 126, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128,
		129, 1, 0, 0, 0, 129, 130, 5, 113, 0, 0, 130, 135, 3, 16, 8, 0, 131, 132,
		5, 67, 0, 0, 132, 134, 5, 68, 0, 0, 133, 131, 1, 0, 0, 0, 134, 137, 1,
		0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 140, 1, 0, 0,
		0, 137, 135, 1, 0, 0, 0, 138, 139, 5, 45, 0, 0, 139, 141, 3, 14, 7, 0,
		140, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142,
		143, 3, 10, 5, 0, 143, 9, 1, 0, 0, 0, 144, 147, 3, 54, 27, 0, 145, 147,
		5, 69, 0, 0, 146, 144, 1, 0, 0, 0, 146, 145, 1, 0, 0, 0, 147, 11, 1, 0,
		0, 0, 148, 151, 3, 84, 42, 0, 149, 151, 5, 48, 0, 0, 150, 148, 1, 0, 0,
		0, 150, 149, 1, 0, 0, 0, 151, 13, 1, 0, 0, 0, 152, 157, 3, 26, 13, 0, 153,
		154, 5, 70, 0, 0, 154, 156, 3, 26, 13, 0, 155, 153, 1, 0, 0, 0, 156, 159,
		1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 15, 1, 0,
		0, 0, 159, 157, 1, 0, 0, 0, 160, 162, 5, 63, 0, 0, 161, 163, 3, 18, 9,
		0, 162, 161, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164,
		165, 5, 64, 0, 0, 165, 17, 1, 0, 0, 0, 166, 171, 3, 20, 10, 0, 167, 168,
		5, 70, 0, 0, 168, 170, 3, 20, 10, 0, 169, 167, 1, 0, 0, 0, 170, 173, 1,
		0, 0, 0, 171, 169, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 176, 1, 0, 0,
		0, 173, 171, 1, 0, 0, 0, 174, 175, 5, 70, 0, 0, 175, 177, 3, 22, 11, 0,
		176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 180, 1, 0, 0, 0, 178,
		180, 3, 22, 11, 0, 179, 166, 1, 0, 0, 0, 179, 178, 1, 0, 0, 0, 180, 19,
		1, 0, 0, 0, 181, 183, 3, 24, 12, 0, 182, 181, 1, 0, 0, 0, 183, 186, 1,
		0, 0, 0, 184, 182, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 187, 1, 0, 0,
		0, 186, 184, 1, 0, 0, 0, 187, 188, 3, 84, 42, 0, 188, 189, 3, 36, 18, 0,
		189, 21, 1, 0, 0, 0, 190, 192, 3, 24, 12, 0, 191, 190, 1, 0, 0, 0, 192,
		195, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 196,
		1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 196, 197, 3, 84, 42, 0, 197, 198, 5,
		109, 0, 0, 198, 199, 3, 36, 18, 0, 199, 23, 1, 0, 0, 0, 200, 201, 5, 18,
		0, 0, 201, 25, 1, 0, 0, 0, 202, 207, 5, 113, 0, 0, 203, 204, 5, 71, 0,
		0, 204, 206, 5, 113, 0, 0, 205, 203, 1, 0, 0, 0, 206, 209, 1, 0, 0, 0,
		207, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 27, 1, 0, 0, 0, 209, 207,
		1, 0, 0, 0, 210, 211, 3, 32, 16, 0, 211, 212, 5, 69, 0, 0, 212, 29, 1,
		0, 0, 0, 213, 214, 5, 113, 0, 0, 214, 217, 3, 16, 8, 0, 215, 216, 5, 45,
		0, 0, 216, 218, 3, 14, 7, 0, 217, 215, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0,
		218, 219, 1, 0, 0, 0, 219, 220, 3, 54, 27, 0, 220, 31, 1, 0, 0, 0, 221,
		222, 3, 84, 42, 0, 222, 227, 3, 34, 17, 0, 223, 224, 5, 70, 0, 0, 224,
		226, 3, 34, 17, 0, 225, 223, 1, 0, 0, 0, 226, 229, 1, 0, 0, 0, 227, 225,
		1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 33, 1, 0, 0, 0, 229, 227, 1, 0,
		0, 0, 230, 233, 3, 36, 18, 0, 231, 232, 5, 72, 0, 0, 232, 234, 3, 38, 19,
		0, 233, 231, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 35, 1, 0, 0, 0, 235,
		240, 5, 113, 0, 0, 236, 237, 5, 67, 0, 0, 237, 239, 5, 68, 0, 0, 238, 236,
		1, 0, 0, 0, 239, 242, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 240, 241, 1, 0,
		0, 0, 241, 37, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 243, 246, 3, 40, 20, 0,
		244, 246, 3, 78, 39, 0, 245, 243, 1, 0, 0, 0, 245, 244, 1, 0, 0, 0, 246,
		39, 1, 0, 0, 0, 247, 259, 5, 65, 0, 0, 248, 253, 3, 38, 19, 0, 249, 250,
		5, 70, 0, 0, 250, 252, 3, 38, 19, 0, 251, 249, 1, 0, 0, 0, 252, 255, 1,
		0, 0, 0, 253, 251, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 257, 1, 0, 0,
		0, 255, 253, 1, 0, 0, 0, 256, 258, 5, 70, 0, 0, 257, 256, 1, 0, 0, 0, 257,
		258, 1, 0, 0, 0, 258, 260, 1, 0, 0, 0, 259, 248, 1, 0, 0, 0, 259, 260,
		1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 262, 5, 66, 0, 0, 262, 41, 1, 0,
		0, 0, 263, 268, 5, 113, 0, 0, 264, 265, 5, 71, 0, 0, 265, 267, 5, 113,
		0, 0, 266, 264, 1, 0, 0, 0, 267, 270, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0,
		268, 269, 1, 0, 0, 0, 269, 43, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 271, 278,
		3, 84, 42, 0, 272, 275, 5, 77, 0, 0, 273, 274, 7, 0, 0, 0, 274, 276, 3,
		84, 42, 0, 275, 273, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 278, 1, 0,
		0, 0, 277, 271, 1, 0, 0, 0, 277, 272, 1, 0, 0, 0, 278, 45, 1, 0, 0, 0,
		279, 286, 3, 48, 24, 0, 280, 286, 3, 50, 25, 0, 281, 286, 5, 60, 0, 0,
		282, 286, 5, 61, 0, 0, 283, 286, 5, 59, 0, 0, 284, 286, 5, 62, 0, 0, 285,
		279, 1, 0, 0, 0, 285, 280, 1, 0, 0, 0, 285, 281, 1, 0, 0, 0, 285, 282,
		1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 285, 284, 1, 0, 0, 0, 286, 47, 1, 0,
		0, 0, 287, 288, 7, 1, 0, 0, 288, 49, 1, 0, 0, 0, 289, 290, 7, 2, 0, 0,
		290, 51, 1, 0, 0, 0, 291, 292, 3, 56, 28, 0, 292, 53, 1, 0, 0, 0, 293,
		294, 5, 65, 0, 0, 294, 295, 3, 56, 28, 0, 295, 296, 5, 66, 0, 0, 296, 55,
		1, 0, 0, 0, 297, 299, 3, 58, 29, 0, 298, 297, 1, 0, 0, 0, 299, 302, 1,
		0, 0, 0, 300, 298, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 57, 1, 0, 0,
		0, 302, 300, 1, 0, 0, 0, 303, 304, 3, 32, 16, 0, 304, 305, 5, 69, 0, 0,
		305, 310, 1, 0, 0, 0, 306, 310, 3, 60, 30, 0, 307, 310, 3, 8, 4, 0, 308,
		310, 3, 0, 0, 0, 309, 303, 1, 0, 0, 0, 309, 306, 1, 0, 0, 0, 309, 307,
		1, 0, 0, 0, 309, 308, 1, 0, 0, 0, 310, 59, 1, 0, 0, 0, 311, 375, 3, 54,
		27, 0, 312, 313, 5, 22, 0, 0, 313, 314, 3, 72, 36, 0, 314, 317, 3, 60,
		30, 0, 315, 316, 5, 15, 0, 0, 316, 318, 3, 60, 30, 0, 317, 315, 1, 0, 0,
		0, 317, 318, 1, 0, 0, 0, 318, 375, 1, 0, 0, 0, 319, 320, 5, 21, 0, 0, 320,
		321, 5, 63, 0, 0, 321, 322, 3, 66, 33, 0, 322, 323, 5, 64, 0, 0, 323, 324,
		3, 60, 30, 0, 324, 375, 1, 0, 0, 0, 325, 326, 5, 50, 0, 0, 326, 327, 3,
		72, 36, 0, 327, 328, 3, 60, 30, 0, 328, 375, 1, 0, 0, 0, 329, 330, 5, 13,
		0, 0, 330, 331, 3, 60, 30, 0, 331, 332, 5, 50, 0, 0, 332, 333, 3, 72, 36,
		0, 333, 334, 5, 69, 0, 0, 334, 375, 1, 0, 0, 0, 335, 336, 5, 41, 0, 0,
		336, 337, 3, 72, 36, 0, 337, 341, 5, 65, 0, 0, 338, 340, 3, 62, 31, 0,
		339, 338, 1, 0, 0, 0, 340, 343, 1, 0, 0, 0, 341, 339, 1, 0, 0, 0, 341,
		342, 1, 0, 0, 0, 342, 347, 1, 0, 0, 0, 343, 341, 1, 0, 0, 0, 344, 346,
		3, 64, 32, 0, 345, 344, 1, 0, 0, 0, 346, 349, 1, 0, 0, 0, 347, 345, 1,
		0, 0, 0, 347, 348, 1, 0, 0, 0, 348, 350, 1, 0, 0, 0, 349, 347, 1, 0, 0,
		0, 350, 351, 5, 66, 0, 0, 351, 375, 1, 0, 0, 0, 352, 354, 5, 36, 0, 0,
		353, 355, 3, 78, 39, 0, 354, 353, 1, 0, 0, 0, 354, 355, 1, 0, 0, 0, 355,
		356, 1, 0, 0, 0, 356, 375, 5, 69, 0, 0, 357, 359, 5, 4, 0, 0, 358, 360,
		5, 113, 0, 0, 359, 358, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0, 360, 361, 1,
		0, 0, 0, 361, 375, 5, 69, 0, 0, 362, 364, 5, 11, 0, 0, 363, 365, 5, 113,
		0, 0, 364, 363, 1, 0, 0, 0, 364, 365, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0,
		366, 375, 5, 69, 0, 0, 367, 375, 5, 69, 0, 0, 368, 369, 3, 78, 39, 0, 369,
		370, 5, 69, 0, 0, 370, 375, 1, 0, 0, 0, 371, 372, 5, 113, 0, 0, 372, 373,
		5, 78, 0, 0, 373, 375, 3, 60, 30, 0, 374, 311, 1, 0, 0, 0, 374, 312, 1,
		0, 0, 0, 374, 319, 1, 0, 0, 0, 374, 325, 1, 0, 0, 0, 374, 329, 1, 0, 0,
		0, 374, 335, 1, 0, 0, 0, 374, 352, 1, 0, 0, 0, 374, 357, 1, 0, 0, 0, 374,
		362, 1, 0, 0, 0, 374, 367, 1, 0, 0, 0, 374, 368, 1, 0, 0, 0, 374, 371,
		1, 0, 0, 0, 375, 61, 1, 0, 0, 0, 376, 378, 3, 64, 32, 0, 377, 376, 1, 0,
		0, 0, 378, 379, 1, 0, 0, 0, 379, 377, 1, 0, 0, 0, 379, 380, 1, 0, 0, 0,
		380, 382, 1, 0, 0, 0, 381, 383, 3, 58, 29, 0, 382, 381, 1, 0, 0, 0, 383,
		384, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 63, 1,
		0, 0, 0, 386, 389, 5, 6, 0, 0, 387, 390, 3, 78, 39, 0, 388, 390, 5, 113,
		0, 0, 389, 387, 1, 0, 0, 0, 389, 388, 1, 0, 0, 0, 390, 391, 1, 0, 0, 0,
		391, 395, 5, 78, 0, 0, 392, 393, 5, 12, 0, 0, 393, 395, 5, 78, 0, 0, 394,
		386, 1, 0, 0, 0, 394, 392, 1, 0, 0, 0, 395, 65, 1, 0, 0, 0, 396, 409, 3,
		70, 35, 0, 397, 399, 3, 68, 34, 0, 398, 397, 1, 0, 0, 0, 398, 399, 1, 0,
		0, 0, 399, 400, 1, 0, 0, 0, 400, 402, 5, 69, 0, 0, 401, 403, 3, 78, 39,
		0, 402, 401, 1, 0, 0, 0, 402, 403, 1, 0, 0, 0, 403, 404, 1, 0, 0, 0, 404,
		406, 5, 69, 0, 0, 405, 407, 3, 74, 37, 0, 406, 405, 1, 0, 0, 0, 406, 407,
		1, 0, 0, 0, 407, 409, 1, 0, 0, 0, 408, 396, 1, 0, 0, 0, 408, 398, 1, 0,
		0, 0, 409, 67, 1, 0, 0, 0, 410, 413, 3, 32, 16, 0, 411, 413, 3, 74, 37,
		0, 412, 410, 1, 0, 0, 0, 412, 411, 1, 0, 0, 0, 413, 69, 1, 0, 0, 0, 414,
		415, 3, 84, 42, 0, 415, 416, 3, 36, 18, 0, 416, 417, 5, 78, 0, 0, 417,
		418, 3, 78, 39, 0, 418, 71, 1, 0, 0, 0, 419, 420, 5, 63, 0, 0, 420, 421,
		3, 78, 39, 0, 421, 422, 5, 64, 0, 0, 422, 73, 1, 0, 0, 0, 423, 428, 3,
		78, 39, 0, 424, 425, 5, 70, 0, 0, 425, 427, 3, 78, 39, 0, 426, 424, 1,
		0, 0, 0, 427, 430, 1, 0, 0, 0, 428, 426, 1, 0, 0, 0, 428, 429, 1, 0, 0,
		0, 429, 75, 1, 0, 0, 0, 430, 428, 1, 0, 0, 0, 431, 432, 5, 113, 0, 0, 432,
		434, 5, 63, 0, 0, 433, 435, 3, 74, 37, 0, 434, 433, 1, 0, 0, 0, 434, 435,
		1, 0, 0, 0, 435, 436, 1, 0, 0, 0, 436, 450, 5, 64, 0, 0, 437, 438, 5, 43,
		0, 0, 438, 440, 5, 63, 0, 0, 439, 441, 3, 74, 37, 0, 440, 439, 1, 0, 0,
		0, 440, 441, 1, 0, 0, 0, 441, 442, 1, 0, 0, 0, 442, 450, 5, 64, 0, 0, 443,
		444, 5, 40, 0, 0, 444, 446, 5, 63, 0, 0, 445, 447, 3, 74, 37, 0, 446, 445,
		1, 0, 0, 0, 446, 447, 1, 0, 0, 0, 447, 448, 1, 0, 0, 0, 448, 450, 5, 64,
		0, 0, 449, 431, 1, 0, 0, 0, 449, 437, 1, 0, 0, 0, 449, 443, 1, 0, 0, 0,
		450, 77, 1, 0, 0, 0, 451, 452, 6, 39, -1, 0, 452, 459, 3, 80, 40, 0, 453,
		459, 3, 76, 38, 0, 454, 455, 7, 3, 0, 0, 455, 459, 3, 78, 39, 15, 456,
		457, 7, 4, 0, 0, 457, 459, 3, 78, 39, 14, 458, 451, 1, 0, 0, 0, 458, 453,
		1, 0, 0, 0, 458, 454, 1, 0, 0, 0, 458, 456, 1, 0, 0, 0, 459, 526, 1, 0,
		0, 0, 460, 461, 10, 13, 0, 0, 461, 462, 7, 5, 0, 0, 462, 525, 3, 78, 39,
		14, 463, 464, 10, 12, 0, 0, 464, 465, 7, 6, 0, 0, 465, 525, 3, 78, 39,
		13, 466, 474, 10, 11, 0, 0, 467, 468, 5, 74, 0, 0, 468, 475, 5, 74, 0,
		0, 469, 470, 5, 73, 0, 0, 470, 471, 5, 73, 0, 0, 471, 475, 5, 73, 0, 0,
		472, 473, 5, 73, 0, 0, 473, 475, 5, 73, 0, 0, 474, 467, 1, 0, 0, 0, 474,
		469, 1, 0, 0, 0, 474, 472, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0, 476, 525,
		3, 78, 39, 12, 477, 478, 10, 10, 0, 0, 478, 479, 7, 7, 0, 0, 479, 525,
		3, 78, 39, 11, 480, 481, 10, 8, 0, 0, 481, 482, 7, 8, 0, 0, 482, 525, 3,
		78, 39, 9, 483, 484, 10, 7, 0, 0, 484, 485, 5, 91, 0, 0, 485, 525, 3, 78,
		39, 8, 486, 487, 10, 6, 0, 0, 487, 488, 5, 93, 0, 0, 488, 525, 3, 78, 39,
		7, 489, 490, 10, 5, 0, 0, 490, 491, 5, 92, 0, 0, 491, 525, 3, 78, 39, 6,
		492, 493, 10, 4, 0, 0, 493, 494, 5, 83, 0, 0, 494, 525, 3, 78, 39, 5, 495,
		496, 10, 3, 0, 0, 496, 497, 5, 84, 0, 0, 497, 525, 3, 78, 39, 4, 498, 499,
		10, 2, 0, 0, 499, 500, 5, 77, 0, 0, 500, 501, 3, 78, 39, 0, 501, 502, 5,
		78, 0, 0, 502, 503, 3, 78, 39, 3, 503, 525, 1, 0, 0, 0, 504, 505, 10, 1,
		0, 0, 505, 506, 7, 9, 0, 0, 506, 525, 3, 78, 39, 1, 507, 508, 10, 19, 0,
		0, 508, 512, 5, 71, 0, 0, 509, 513, 5, 113, 0, 0, 510, 513, 3, 76, 38,
		0, 511, 513, 5, 43, 0, 0, 512, 509, 1, 0, 0, 0, 512, 510, 1, 0, 0, 0, 512,
		511, 1, 0, 0, 0, 513, 525, 1, 0, 0, 0, 514, 515, 10, 18, 0, 0, 515, 516,
		5, 67, 0, 0, 516, 517, 3, 78, 39, 0, 517, 518, 5, 68, 0, 0, 518, 525, 1,
		0, 0, 0, 519, 520, 10, 16, 0, 0, 520, 525, 7, 10, 0, 0, 521, 522, 10, 9,
		0, 0, 522, 523, 5, 26, 0, 0, 523, 525, 3, 84, 42, 0, 524, 460, 1, 0, 0,
		0, 524, 463, 1, 0, 0, 0, 524, 466, 1, 0, 0, 0, 524, 477, 1, 0, 0, 0, 524,
		480, 1, 0, 0, 0, 524, 483, 1, 0, 0, 0, 524, 486, 1, 0, 0, 0, 524, 489,
		1, 0, 0, 0, 524, 492, 1, 0, 0, 0, 524, 495, 1, 0, 0, 0, 524, 498, 1, 0,
		0, 0, 524, 504, 1, 0, 0, 0, 524, 507, 1, 0, 0, 0, 524, 514, 1, 0, 0, 0,
		524, 519, 1, 0, 0, 0, 524, 521, 1, 0, 0, 0, 525, 528, 1, 0, 0, 0, 526,
		524, 1, 0, 0, 0, 526, 527, 1, 0, 0, 0, 527, 79, 1, 0, 0, 0, 528, 526, 1,
		0, 0, 0, 529, 530, 5, 63, 0, 0, 530, 531, 3, 78, 39, 0, 531, 532, 5, 64,
		0, 0, 532, 538, 1, 0, 0, 0, 533, 538, 5, 43, 0, 0, 534, 538, 5, 40, 0,
		0, 535, 538, 3, 46, 23, 0, 536, 538, 5, 113, 0, 0, 537, 529, 1, 0, 0, 0,
		537, 533, 1, 0, 0, 0, 537, 534, 1, 0, 0, 0, 537, 535, 1, 0, 0, 0, 537,
		536, 1, 0, 0, 0, 538, 81, 1, 0, 0, 0, 539, 544, 3, 84, 42, 0, 540, 541,
		5, 70, 0, 0, 541, 543, 3, 84, 42, 0, 542, 540, 1, 0, 0, 0, 543, 546, 1,
		0, 0, 0, 544, 542, 1, 0, 0, 0, 544, 545, 1, 0, 0, 0, 545, 83, 1, 0, 0,
		0, 546, 544, 1, 0, 0, 0, 547, 551, 3, 42, 21, 0, 548, 551, 3, 86, 43, 0,
		549, 551, 3, 88, 44, 0, 550, 547, 1, 0, 0, 0, 550, 548, 1, 0, 0, 0, 550,
		549, 1, 0, 0, 0, 551, 556, 1, 0, 0, 0, 552, 553, 5, 67, 0, 0, 553, 555,
		5, 68, 0, 0, 554, 552, 1, 0, 0, 0, 555, 558, 1, 0, 0, 0, 556, 554, 1, 0,
		0, 0, 556, 557, 1, 0, 0, 0, 557, 85, 1, 0, 0, 0, 558, 556, 1, 0, 0, 0,
		559, 560, 5, 51, 0, 0, 560, 561, 3, 12, 6, 0, 561, 563, 5, 63, 0, 0, 562,
		564, 3, 82, 41, 0, 563, 562, 1, 0, 0, 0, 563, 564, 1, 0, 0, 0, 564, 565,
		1, 0, 0, 0, 565, 566, 5, 64, 0, 0, 566, 87, 1, 0, 0, 0, 567, 568, 7, 11,
		0, 0, 568, 89, 1, 0, 0, 0, 569, 570, 5, 113, 0, 0, 570, 571, 3, 94, 47,
		0, 571, 91, 1, 0, 0, 0, 572, 579, 3, 94, 47, 0, 573, 574, 5, 71, 0, 0,
		574, 576, 5, 113, 0, 0, 575, 577, 3, 94, 47, 0, 576, 575, 1, 0, 0, 0, 576,
		577, 1, 0, 0, 0, 577, 579, 1, 0, 0, 0, 578, 572, 1, 0, 0, 0, 578, 573,
		1, 0, 0, 0, 579, 93, 1, 0, 0, 0, 580, 582, 5, 63, 0, 0, 581, 583, 3, 74,
		37, 0, 582, 581, 1, 0, 0, 0, 582, 583, 1, 0, 0, 0, 583, 584, 1, 0, 0, 0,
		584, 585, 5, 64, 0, 0, 585, 95, 1, 0, 0, 0, 66, 100, 104, 112, 119, 124,
		127, 135, 140, 146, 150, 157, 162, 171, 176, 179, 184, 193, 207, 217, 227,
		233, 240, 245, 253, 257, 259, 268, 275, 277, 285, 300, 309, 317, 341, 347,
		354, 359, 364, 374, 379, 384, 389, 394, 398, 402, 406, 408, 412, 428, 434,
		440, 446, 449, 458, 474, 512, 524, 526, 537, 544, 550, 556, 563, 576, 578,
		582,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// RoseParserInit initializes any static state used to implement RoseParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRoseParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RoseParserInit() {
	staticData := &roseParserStaticData
	staticData.once.Do(roseParserInit)
}

// NewRoseParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRoseParser(input antlr.TokenStream) *RoseParser {
	RoseParserInit()
	this := new(RoseParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &roseParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// RoseParser tokens.
const (
	RoseParserEOF               = antlr.TokenEOF
	RoseParserABSTRACT          = 1
	RoseParserASSERT            = 2
	RoseParserBOOLEAN           = 3
	RoseParserBREAK             = 4
	RoseParserBYTE              = 5
	RoseParserCASE              = 6
	RoseParserCATCH             = 7
	RoseParserCHAR              = 8
	RoseParserCLASS             = 9
	RoseParserCONST             = 10
	RoseParserCONTINUE          = 11
	RoseParserDEFAULT           = 12
	RoseParserDO                = 13
	RoseParserDOUBLE            = 14
	RoseParserELSE              = 15
	RoseParserENUM              = 16
	RoseParserEXTENDS           = 17
	RoseParserFINAL             = 18
	RoseParserFINALLY           = 19
	RoseParserFLOAT             = 20
	RoseParserFOR               = 21
	RoseParserIF                = 22
	RoseParserGOTO              = 23
	RoseParserIMPLEMENTS        = 24
	RoseParserIMPORT            = 25
	RoseParserINSTANCEOF        = 26
	RoseParserINT               = 27
	RoseParserINTERFACE         = 28
	RoseParserLONG              = 29
	RoseParserNATIVE            = 30
	RoseParserNEW               = 31
	RoseParserPACKAGE           = 32
	RoseParserPRIVATE           = 33
	RoseParserPROTECTED         = 34
	RoseParserPUBLIC            = 35
	RoseParserRETURN            = 36
	RoseParserSHORT             = 37
	RoseParserSTATIC            = 38
	RoseParserSTRICTFP          = 39
	RoseParserSUPER             = 40
	RoseParserSWITCH            = 41
	RoseParserSYNCHRONIZED      = 42
	RoseParserTHIS              = 43
	RoseParserTHROW             = 44
	RoseParserTHROWS            = 45
	RoseParserTRANSIENT         = 46
	RoseParserTRY               = 47
	RoseParserVOID              = 48
	RoseParserVOLATILE          = 49
	RoseParserWHILE             = 50
	RoseParserFUNCTION          = 51
	RoseParserSTRING            = 52
	RoseParserDECIMAL_LITERAL   = 53
	RoseParserHEX_LITERAL       = 54
	RoseParserOCT_LITERAL       = 55
	RoseParserBINARY_LITERAL    = 56
	RoseParserFLOAT_LITERAL     = 57
	RoseParserHEX_FLOAT_LITERAL = 58
	RoseParserBOOL_LITERAL      = 59
	RoseParserCHAR_LITERAL      = 60
	RoseParserSTRING_LITERAL    = 61
	RoseParserNULL_LITERAL      = 62
	RoseParserLPAREN            = 63
	RoseParserRPAREN            = 64
	RoseParserLBRACE            = 65
	RoseParserRBRACE            = 66
	RoseParserLBRACK            = 67
	RoseParserRBRACK            = 68
	RoseParserSEMI              = 69
	RoseParserCOMMA             = 70
	RoseParserDOT               = 71
	RoseParserASSIGN            = 72
	RoseParserGT                = 73
	RoseParserLT                = 74
	RoseParserBANG              = 75
	RoseParserTILDE             = 76
	RoseParserQUESTION          = 77
	RoseParserCOLON             = 78
	RoseParserEQUAL             = 79
	RoseParserLE                = 80
	RoseParserGE                = 81
	RoseParserNOTEQUAL          = 82
	RoseParserAND               = 83
	RoseParserOR                = 84
	RoseParserINC               = 85
	RoseParserDEC               = 86
	RoseParserADD               = 87
	RoseParserSUB               = 88
	RoseParserMUL               = 89
	RoseParserDIV               = 90
	RoseParserBITAND            = 91
	RoseParserBITOR             = 92
	RoseParserCARET             = 93
	RoseParserMOD               = 94
	RoseParserADD_ASSIGN        = 95
	RoseParserSUB_ASSIGN        = 96
	RoseParserMUL_ASSIGN        = 97
	RoseParserDIV_ASSIGN        = 98
	RoseParserAND_ASSIGN        = 99
	RoseParserOR_ASSIGN         = 100
	RoseParserXOR_ASSIGN        = 101
	RoseParserMOD_ASSIGN        = 102
	RoseParserLSHIFT_ASSIGN     = 103
	RoseParserRSHIFT_ASSIGN     = 104
	RoseParserURSHIFT_ASSIGN    = 105
	RoseParserARROW             = 106
	RoseParserCOLONCOLON        = 107
	RoseParserAT                = 108
	RoseParserELLIPSIS          = 109
	RoseParserWS                = 110
	RoseParserCOMMENT           = 111
	RoseParserLINE_COMMENT      = 112
	RoseParserIDENTIFIER        = 113
)

// RoseParser rules.
const (
	RoseParserRULE_classDeclaration          = 0
	RoseParserRULE_classBody                 = 1
	RoseParserRULE_classBodyDeclaration      = 2
	RoseParserRULE_memberDeclaration         = 3
	RoseParserRULE_functionDeclaration       = 4
	RoseParserRULE_functionBody              = 5
	RoseParserRULE_typeTypeOrVoid            = 6
	RoseParserRULE_qualifiedNameList         = 7
	RoseParserRULE_formalParameters          = 8
	RoseParserRULE_formalParameterList       = 9
	RoseParserRULE_formalParameter           = 10
	RoseParserRULE_lastFormalParameter       = 11
	RoseParserRULE_variableModifier          = 12
	RoseParserRULE_qualifiedName             = 13
	RoseParserRULE_fieldDeclaration          = 14
	RoseParserRULE_constructorDeclaration    = 15
	RoseParserRULE_variableDeclarators       = 16
	RoseParserRULE_variableDeclarator        = 17
	RoseParserRULE_variableDeclaratorId      = 18
	RoseParserRULE_variableInitializer       = 19
	RoseParserRULE_arrayInitializer          = 20
	RoseParserRULE_classOrInterfaceType      = 21
	RoseParserRULE_typeArgument              = 22
	RoseParserRULE_literal                   = 23
	RoseParserRULE_integerLiteral            = 24
	RoseParserRULE_floatLiteral              = 25
	RoseParserRULE_prog                      = 26
	RoseParserRULE_block                     = 27
	RoseParserRULE_blockStatements           = 28
	RoseParserRULE_blockStatement            = 29
	RoseParserRULE_statement                 = 30
	RoseParserRULE_switchBlockStatementGroup = 31
	RoseParserRULE_switchLabel               = 32
	RoseParserRULE_forControl                = 33
	RoseParserRULE_forInit                   = 34
	RoseParserRULE_enhancedForControl        = 35
	RoseParserRULE_parExpression             = 36
	RoseParserRULE_expressionList            = 37
	RoseParserRULE_functionCall              = 38
	RoseParserRULE_expression                = 39
	RoseParserRULE_primary                   = 40
	RoseParserRULE_typeList                  = 41
	RoseParserRULE_typeType                  = 42
	RoseParserRULE_functionType              = 43
	RoseParserRULE_primitiveType             = 44
	RoseParserRULE_creator                   = 45
	RoseParserRULE_superSuffix               = 46
	RoseParserRULE_arguments                 = 47
)

// IClassDeclarationContext is an interface to support dynamic dispatch.
type IClassDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassDeclarationContext differentiates from other interfaces.
	IsClassDeclarationContext()
}

type ClassDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassDeclarationContext() *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_classDeclaration
	return p
}

func (*ClassDeclarationContext) IsClassDeclarationContext() {}

func NewClassDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_classDeclaration

	return p
}

func (s *ClassDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDeclarationContext) CLASS() antlr.TerminalNode {
	return s.GetToken(RoseParserCLASS, 0)
}

func (s *ClassDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RoseParserIDENTIFIER, 0)
}

func (s *ClassDeclarationContext) ClassBody() IClassBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassBodyContext)
}

func (s *ClassDeclarationContext) EXTENDS() antlr.TerminalNode {
	return s.GetToken(RoseParserEXTENDS, 0)
}

func (s *ClassDeclarationContext) TypeType() ITypeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *ClassDeclarationContext) IMPLEMENTS() antlr.TerminalNode {
	return s.GetToken(RoseParserIMPLEMENTS, 0)
}

func (s *ClassDeclarationContext) TypeList() ITypeListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeListContext)
}

func (s *ClassDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterClassDeclaration(s)
	}
}

func (s *ClassDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitClassDeclaration(s)
	}
}

func (s *ClassDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitClassDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) ClassDeclaration() (localctx IClassDeclarationContext) {
	this := p
	_ = this

	localctx = NewClassDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RoseParserRULE_classDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(RoseParserCLASS)
	}
	{
		p.SetState(97)
		p.Match(RoseParserIDENTIFIER)
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RoseParserEXTENDS {
		{
			p.SetState(98)
			p.Match(RoseParserEXTENDS)
		}
		{
			p.SetState(99)
			p.TypeType()
		}

	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RoseParserIMPLEMENTS {
		{
			p.SetState(102)
			p.Match(RoseParserIMPLEMENTS)
		}
		{
			p.SetState(103)
			p.TypeList()
		}

	}
	{
		p.SetState(106)
		p.ClassBody()
	}

	return localctx
}

// IClassBodyContext is an interface to support dynamic dispatch.
type IClassBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassBodyContext differentiates from other interfaces.
	IsClassBodyContext()
}

type ClassBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassBodyContext() *ClassBodyContext {
	var p = new(ClassBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_classBody
	return p
}

func (*ClassBodyContext) IsClassBodyContext() {}

func NewClassBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassBodyContext {
	var p = new(ClassBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_classBody

	return p
}

func (s *ClassBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassBodyContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(RoseParserLBRACE, 0)
}

func (s *ClassBodyContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(RoseParserRBRACE, 0)
}

func (s *ClassBodyContext) AllClassBodyDeclaration() []IClassBodyDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IClassBodyDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IClassBodyDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IClassBodyDeclarationContext); ok {
			tst[i] = t.(IClassBodyDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ClassBodyContext) ClassBodyDeclaration(i int) IClassBodyDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassBodyDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassBodyDeclarationContext)
}

func (s *ClassBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterClassBody(s)
	}
}

func (s *ClassBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitClassBody(s)
	}
}

func (s *ClassBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitClassBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) ClassBody() (localctx IClassBodyContext) {
	this := p
	_ = this

	localctx = NewClassBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RoseParserRULE_classBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(RoseParserLBRACE)
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7037012528874280) != 0 || _la == RoseParserSEMI || _la == RoseParserIDENTIFIER {
		{
			p.SetState(109)
			p.ClassBodyDeclaration()
		}

		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(115)
		p.Match(RoseParserRBRACE)
	}

	return localctx
}

// IClassBodyDeclarationContext is an interface to support dynamic dispatch.
type IClassBodyDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassBodyDeclarationContext differentiates from other interfaces.
	IsClassBodyDeclarationContext()
}

type ClassBodyDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassBodyDeclarationContext() *ClassBodyDeclarationContext {
	var p = new(ClassBodyDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_classBodyDeclaration
	return p
}

func (*ClassBodyDeclarationContext) IsClassBodyDeclarationContext() {}

func NewClassBodyDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassBodyDeclarationContext {
	var p = new(ClassBodyDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_classBodyDeclaration

	return p
}

func (s *ClassBodyDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassBodyDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RoseParserSEMI, 0)
}

func (s *ClassBodyDeclarationContext) MemberDeclaration() IMemberDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemberDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMemberDeclarationContext)
}

func (s *ClassBodyDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassBodyDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassBodyDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterClassBodyDeclaration(s)
	}
}

func (s *ClassBodyDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitClassBodyDeclaration(s)
	}
}

func (s *ClassBodyDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitClassBodyDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) ClassBodyDeclaration() (localctx IClassBodyDeclarationContext) {
	this := p
	_ = this

	localctx = NewClassBodyDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RoseParserRULE_classBodyDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RoseParserSEMI:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.Match(RoseParserSEMI)
		}

	case RoseParserBOOLEAN, RoseParserBYTE, RoseParserCHAR, RoseParserCLASS, RoseParserDOUBLE, RoseParserFLOAT, RoseParserINT, RoseParserLONG, RoseParserSHORT, RoseParserVOID, RoseParserFUNCTION, RoseParserSTRING, RoseParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.MemberDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMemberDeclarationContext is an interface to support dynamic dispatch.
type IMemberDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberDeclarationContext differentiates from other interfaces.
	IsMemberDeclarationContext()
}

type MemberDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberDeclarationContext() *MemberDeclarationContext {
	var p = new(MemberDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_memberDeclaration
	return p
}

func (*MemberDeclarationContext) IsMemberDeclarationContext() {}

func NewMemberDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberDeclarationContext {
	var p = new(MemberDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_memberDeclaration

	return p
}

func (s *MemberDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberDeclarationContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *MemberDeclarationContext) FieldDeclaration() IFieldDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldDeclarationContext)
}

func (s *MemberDeclarationContext) ClassDeclaration() IClassDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassDeclarationContext)
}

func (s *MemberDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterMemberDeclaration(s)
	}
}

func (s *MemberDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitMemberDeclaration(s)
	}
}

func (s *MemberDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitMemberDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) MemberDeclaration() (localctx IMemberDeclarationContext) {
	this := p
	_ = this

	localctx = NewMemberDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RoseParserRULE_memberDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.FunctionDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.FieldDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(123)
			p.ClassDeclaration()
		}

	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RoseParserIDENTIFIER, 0)
}

func (s *FunctionDeclarationContext) FormalParameters() IFormalParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormalParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormalParametersContext)
}

func (s *FunctionDeclarationContext) FunctionBody() IFunctionBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *FunctionDeclarationContext) TypeTypeOrVoid() ITypeTypeOrVoidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeTypeOrVoidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeTypeOrVoidContext)
}

func (s *FunctionDeclarationContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(RoseParserLBRACK)
}

func (s *FunctionDeclarationContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(RoseParserLBRACK, i)
}

func (s *FunctionDeclarationContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(RoseParserRBRACK)
}

func (s *FunctionDeclarationContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(RoseParserRBRACK, i)
}

func (s *FunctionDeclarationContext) THROWS() antlr.TerminalNode {
	return s.GetToken(RoseParserTHROWS, 0)
}

func (s *FunctionDeclarationContext) QualifiedNameList() IQualifiedNameListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifiedNameListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameListContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	this := p
	_ = this

	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RoseParserRULE_functionDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(127)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(126)
			p.TypeTypeOrVoid()
		}

	}
	{
		p.SetState(129)
		p.Match(RoseParserIDENTIFIER)
	}
	{
		p.SetState(130)
		p.FormalParameters()
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RoseParserLBRACK {
		{
			p.SetState(131)
			p.Match(RoseParserLBRACK)
		}
		{
			p.SetState(132)
			p.Match(RoseParserRBRACK)
		}

		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RoseParserTHROWS {
		{
			p.SetState(138)
			p.Match(RoseParserTHROWS)
		}
		{
			p.SetState(139)
			p.QualifiedNameList()
		}

	}
	{
		p.SetState(142)
		p.FunctionBody()
	}

	return localctx
}

// IFunctionBodyContext is an interface to support dynamic dispatch.
type IFunctionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionBodyContext differentiates from other interfaces.
	IsFunctionBodyContext()
}

type FunctionBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionBodyContext() *FunctionBodyContext {
	var p = new(FunctionBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_functionBody
	return p
}

func (*FunctionBodyContext) IsFunctionBodyContext() {}

func NewFunctionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionBodyContext {
	var p = new(FunctionBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_functionBody

	return p
}

func (s *FunctionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionBodyContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionBodyContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RoseParserSEMI, 0)
}

func (s *FunctionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterFunctionBody(s)
	}
}

func (s *FunctionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitFunctionBody(s)
	}
}

func (s *FunctionBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitFunctionBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) FunctionBody() (localctx IFunctionBodyContext) {
	this := p
	_ = this

	localctx = NewFunctionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RoseParserRULE_functionBody)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(146)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RoseParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.Block()
		}

	case RoseParserSEMI:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
			p.Match(RoseParserSEMI)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeTypeOrVoidContext is an interface to support dynamic dispatch.
type ITypeTypeOrVoidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeTypeOrVoidContext differentiates from other interfaces.
	IsTypeTypeOrVoidContext()
}

type TypeTypeOrVoidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeTypeOrVoidContext() *TypeTypeOrVoidContext {
	var p = new(TypeTypeOrVoidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_typeTypeOrVoid
	return p
}

func (*TypeTypeOrVoidContext) IsTypeTypeOrVoidContext() {}

func NewTypeTypeOrVoidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeTypeOrVoidContext {
	var p = new(TypeTypeOrVoidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_typeTypeOrVoid

	return p
}

func (s *TypeTypeOrVoidContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeTypeOrVoidContext) TypeType() ITypeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *TypeTypeOrVoidContext) VOID() antlr.TerminalNode {
	return s.GetToken(RoseParserVOID, 0)
}

func (s *TypeTypeOrVoidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeTypeOrVoidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeTypeOrVoidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterTypeTypeOrVoid(s)
	}
}

func (s *TypeTypeOrVoidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitTypeTypeOrVoid(s)
	}
}

func (s *TypeTypeOrVoidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitTypeTypeOrVoid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) TypeTypeOrVoid() (localctx ITypeTypeOrVoidContext) {
	this := p
	_ = this

	localctx = NewTypeTypeOrVoidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RoseParserRULE_typeTypeOrVoid)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(150)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RoseParserBOOLEAN, RoseParserBYTE, RoseParserCHAR, RoseParserDOUBLE, RoseParserFLOAT, RoseParserINT, RoseParserLONG, RoseParserSHORT, RoseParserFUNCTION, RoseParserSTRING, RoseParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.TypeType()
		}

	case RoseParserVOID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(149)
			p.Match(RoseParserVOID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IQualifiedNameListContext is an interface to support dynamic dispatch.
type IQualifiedNameListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedNameListContext differentiates from other interfaces.
	IsQualifiedNameListContext()
}

type QualifiedNameListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedNameListContext() *QualifiedNameListContext {
	var p = new(QualifiedNameListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_qualifiedNameList
	return p
}

func (*QualifiedNameListContext) IsQualifiedNameListContext() {}

func NewQualifiedNameListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedNameListContext {
	var p = new(QualifiedNameListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_qualifiedNameList

	return p
}

func (s *QualifiedNameListContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedNameListContext) AllQualifiedName() []IQualifiedNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQualifiedNameContext); ok {
			len++
		}
	}

	tst := make([]IQualifiedNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQualifiedNameContext); ok {
			tst[i] = t.(IQualifiedNameContext)
			i++
		}
	}

	return tst
}

func (s *QualifiedNameListContext) QualifiedName(i int) IQualifiedNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifiedNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameContext)
}

func (s *QualifiedNameListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RoseParserCOMMA)
}

func (s *QualifiedNameListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RoseParserCOMMA, i)
}

func (s *QualifiedNameListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedNameListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedNameListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterQualifiedNameList(s)
	}
}

func (s *QualifiedNameListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitQualifiedNameList(s)
	}
}

func (s *QualifiedNameListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitQualifiedNameList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) QualifiedNameList() (localctx IQualifiedNameListContext) {
	this := p
	_ = this

	localctx = NewQualifiedNameListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RoseParserRULE_qualifiedNameList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.QualifiedName()
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RoseParserCOMMA {
		{
			p.SetState(153)
			p.Match(RoseParserCOMMA)
		}
		{
			p.SetState(154)
			p.QualifiedName()
		}

		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFormalParametersContext is an interface to support dynamic dispatch.
type IFormalParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParametersContext differentiates from other interfaces.
	IsFormalParametersContext()
}

type FormalParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParametersContext() *FormalParametersContext {
	var p = new(FormalParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_formalParameters
	return p
}

func (*FormalParametersContext) IsFormalParametersContext() {}

func NewFormalParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParametersContext {
	var p = new(FormalParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_formalParameters

	return p
}

func (s *FormalParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParametersContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RoseParserLPAREN, 0)
}

func (s *FormalParametersContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RoseParserRPAREN, 0)
}

func (s *FormalParametersContext) FormalParameterList() IFormalParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormalParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormalParameterListContext)
}

func (s *FormalParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterFormalParameters(s)
	}
}

func (s *FormalParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitFormalParameters(s)
	}
}

func (s *FormalParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitFormalParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) FormalParameters() (localctx IFormalParametersContext) {
	this := p
	_ = this

	localctx = NewFormalParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RoseParserRULE_formalParameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(RoseParserLPAREN)
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6755537552425256) != 0 || _la == RoseParserIDENTIFIER {
		{
			p.SetState(161)
			p.FormalParameterList()
		}

	}
	{
		p.SetState(164)
		p.Match(RoseParserRPAREN)
	}

	return localctx
}

// IFormalParameterListContext is an interface to support dynamic dispatch.
type IFormalParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParameterListContext differentiates from other interfaces.
	IsFormalParameterListContext()
}

type FormalParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParameterListContext() *FormalParameterListContext {
	var p = new(FormalParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_formalParameterList
	return p
}

func (*FormalParameterListContext) IsFormalParameterListContext() {}

func NewFormalParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParameterListContext {
	var p = new(FormalParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_formalParameterList

	return p
}

func (s *FormalParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParameterListContext) AllFormalParameter() []IFormalParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFormalParameterContext); ok {
			len++
		}
	}

	tst := make([]IFormalParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFormalParameterContext); ok {
			tst[i] = t.(IFormalParameterContext)
			i++
		}
	}

	return tst
}

func (s *FormalParameterListContext) FormalParameter(i int) IFormalParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormalParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormalParameterContext)
}

func (s *FormalParameterListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RoseParserCOMMA)
}

func (s *FormalParameterListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RoseParserCOMMA, i)
}

func (s *FormalParameterListContext) LastFormalParameter() ILastFormalParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILastFormalParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILastFormalParameterContext)
}

func (s *FormalParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterFormalParameterList(s)
	}
}

func (s *FormalParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitFormalParameterList(s)
	}
}

func (s *FormalParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitFormalParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) FormalParameterList() (localctx IFormalParameterListContext) {
	this := p
	_ = this

	localctx = NewFormalParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RoseParserRULE_formalParameterList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			p.FormalParameter()
		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(167)
					p.Match(RoseParserCOMMA)
				}
				{
					p.SetState(168)
					p.FormalParameter()
				}

			}
			p.SetState(173)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RoseParserCOMMA {
			{
				p.SetState(174)
				p.Match(RoseParserCOMMA)
			}
			{
				p.SetState(175)
				p.LastFormalParameter()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(178)
			p.LastFormalParameter()
		}

	}

	return localctx
}

// IFormalParameterContext is an interface to support dynamic dispatch.
type IFormalParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParameterContext differentiates from other interfaces.
	IsFormalParameterContext()
}

type FormalParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParameterContext() *FormalParameterContext {
	var p = new(FormalParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_formalParameter
	return p
}

func (*FormalParameterContext) IsFormalParameterContext() {}

func NewFormalParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParameterContext {
	var p = new(FormalParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_formalParameter

	return p
}

func (s *FormalParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParameterContext) TypeType() ITypeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *FormalParameterContext) VariableDeclaratorId() IVariableDeclaratorIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclaratorIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorIdContext)
}

func (s *FormalParameterContext) AllVariableModifier() []IVariableModifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableModifierContext); ok {
			len++
		}
	}

	tst := make([]IVariableModifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableModifierContext); ok {
			tst[i] = t.(IVariableModifierContext)
			i++
		}
	}

	return tst
}

func (s *FormalParameterContext) VariableModifier(i int) IVariableModifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableModifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableModifierContext)
}

func (s *FormalParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterFormalParameter(s)
	}
}

func (s *FormalParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitFormalParameter(s)
	}
}

func (s *FormalParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitFormalParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) FormalParameter() (localctx IFormalParameterContext) {
	this := p
	_ = this

	localctx = NewFormalParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RoseParserRULE_formalParameter)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RoseParserFINAL {
		{
			p.SetState(181)
			p.VariableModifier()
		}

		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(187)
		p.TypeType()
	}
	{
		p.SetState(188)
		p.VariableDeclaratorId()
	}

	return localctx
}

// ILastFormalParameterContext is an interface to support dynamic dispatch.
type ILastFormalParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLastFormalParameterContext differentiates from other interfaces.
	IsLastFormalParameterContext()
}

type LastFormalParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLastFormalParameterContext() *LastFormalParameterContext {
	var p = new(LastFormalParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_lastFormalParameter
	return p
}

func (*LastFormalParameterContext) IsLastFormalParameterContext() {}

func NewLastFormalParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LastFormalParameterContext {
	var p = new(LastFormalParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_lastFormalParameter

	return p
}

func (s *LastFormalParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *LastFormalParameterContext) TypeType() ITypeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *LastFormalParameterContext) ELLIPSIS() antlr.TerminalNode {
	return s.GetToken(RoseParserELLIPSIS, 0)
}

func (s *LastFormalParameterContext) VariableDeclaratorId() IVariableDeclaratorIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclaratorIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorIdContext)
}

func (s *LastFormalParameterContext) AllVariableModifier() []IVariableModifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableModifierContext); ok {
			len++
		}
	}

	tst := make([]IVariableModifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableModifierContext); ok {
			tst[i] = t.(IVariableModifierContext)
			i++
		}
	}

	return tst
}

func (s *LastFormalParameterContext) VariableModifier(i int) IVariableModifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableModifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableModifierContext)
}

func (s *LastFormalParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LastFormalParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LastFormalParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterLastFormalParameter(s)
	}
}

func (s *LastFormalParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitLastFormalParameter(s)
	}
}

func (s *LastFormalParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitLastFormalParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) LastFormalParameter() (localctx ILastFormalParameterContext) {
	this := p
	_ = this

	localctx = NewLastFormalParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RoseParserRULE_lastFormalParameter)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RoseParserFINAL {
		{
			p.SetState(190)
			p.VariableModifier()
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(196)
		p.TypeType()
	}
	{
		p.SetState(197)
		p.Match(RoseParserELLIPSIS)
	}
	{
		p.SetState(198)
		p.VariableDeclaratorId()
	}

	return localctx
}

// IVariableModifierContext is an interface to support dynamic dispatch.
type IVariableModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableModifierContext differentiates from other interfaces.
	IsVariableModifierContext()
}

type VariableModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableModifierContext() *VariableModifierContext {
	var p = new(VariableModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_variableModifier
	return p
}

func (*VariableModifierContext) IsVariableModifierContext() {}

func NewVariableModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableModifierContext {
	var p = new(VariableModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_variableModifier

	return p
}

func (s *VariableModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableModifierContext) FINAL() antlr.TerminalNode {
	return s.GetToken(RoseParserFINAL, 0)
}

func (s *VariableModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterVariableModifier(s)
	}
}

func (s *VariableModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitVariableModifier(s)
	}
}

func (s *VariableModifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitVariableModifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) VariableModifier() (localctx IVariableModifierContext) {
	this := p
	_ = this

	localctx = NewVariableModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RoseParserRULE_variableModifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(RoseParserFINAL)
	}

	return localctx
}

// IQualifiedNameContext is an interface to support dynamic dispatch.
type IQualifiedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedNameContext differentiates from other interfaces.
	IsQualifiedNameContext()
}

type QualifiedNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedNameContext() *QualifiedNameContext {
	var p = new(QualifiedNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_qualifiedName
	return p
}

func (*QualifiedNameContext) IsQualifiedNameContext() {}

func NewQualifiedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedNameContext {
	var p = new(QualifiedNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_qualifiedName

	return p
}

func (s *QualifiedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedNameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(RoseParserIDENTIFIER)
}

func (s *QualifiedNameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(RoseParserIDENTIFIER, i)
}

func (s *QualifiedNameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(RoseParserDOT)
}

func (s *QualifiedNameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(RoseParserDOT, i)
}

func (s *QualifiedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterQualifiedName(s)
	}
}

func (s *QualifiedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitQualifiedName(s)
	}
}

func (s *QualifiedNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitQualifiedName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) QualifiedName() (localctx IQualifiedNameContext) {
	this := p
	_ = this

	localctx = NewQualifiedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RoseParserRULE_qualifiedName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Match(RoseParserIDENTIFIER)
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RoseParserDOT {
		{
			p.SetState(203)
			p.Match(RoseParserDOT)
		}
		{
			p.SetState(204)
			p.Match(RoseParserIDENTIFIER)
		}

		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldDeclarationContext is an interface to support dynamic dispatch.
type IFieldDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldDeclarationContext differentiates from other interfaces.
	IsFieldDeclarationContext()
}

type FieldDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDeclarationContext() *FieldDeclarationContext {
	var p = new(FieldDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_fieldDeclaration
	return p
}

func (*FieldDeclarationContext) IsFieldDeclarationContext() {}

func NewFieldDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDeclarationContext {
	var p = new(FieldDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_fieldDeclaration

	return p
}

func (s *FieldDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDeclarationContext) VariableDeclarators() IVariableDeclaratorsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclaratorsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorsContext)
}

func (s *FieldDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RoseParserSEMI, 0)
}

func (s *FieldDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterFieldDeclaration(s)
	}
}

func (s *FieldDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitFieldDeclaration(s)
	}
}

func (s *FieldDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitFieldDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) FieldDeclaration() (localctx IFieldDeclarationContext) {
	this := p
	_ = this

	localctx = NewFieldDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RoseParserRULE_fieldDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.VariableDeclarators()
	}
	{
		p.SetState(211)
		p.Match(RoseParserSEMI)
	}

	return localctx
}

// IConstructorDeclarationContext is an interface to support dynamic dispatch.
type IConstructorDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetConstructorBody returns the constructorBody rule contexts.
	GetConstructorBody() IBlockContext

	// SetConstructorBody sets the constructorBody rule contexts.
	SetConstructorBody(IBlockContext)

	// IsConstructorDeclarationContext differentiates from other interfaces.
	IsConstructorDeclarationContext()
}

type ConstructorDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	constructorBody IBlockContext
}

func NewEmptyConstructorDeclarationContext() *ConstructorDeclarationContext {
	var p = new(ConstructorDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_constructorDeclaration
	return p
}

func (*ConstructorDeclarationContext) IsConstructorDeclarationContext() {}

func NewConstructorDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstructorDeclarationContext {
	var p = new(ConstructorDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_constructorDeclaration

	return p
}

func (s *ConstructorDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstructorDeclarationContext) GetConstructorBody() IBlockContext { return s.constructorBody }

func (s *ConstructorDeclarationContext) SetConstructorBody(v IBlockContext) { s.constructorBody = v }

func (s *ConstructorDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RoseParserIDENTIFIER, 0)
}

func (s *ConstructorDeclarationContext) FormalParameters() IFormalParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormalParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormalParametersContext)
}

func (s *ConstructorDeclarationContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ConstructorDeclarationContext) THROWS() antlr.TerminalNode {
	return s.GetToken(RoseParserTHROWS, 0)
}

func (s *ConstructorDeclarationContext) QualifiedNameList() IQualifiedNameListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifiedNameListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameListContext)
}

func (s *ConstructorDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstructorDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstructorDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterConstructorDeclaration(s)
	}
}

func (s *ConstructorDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitConstructorDeclaration(s)
	}
}

func (s *ConstructorDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitConstructorDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) ConstructorDeclaration() (localctx IConstructorDeclarationContext) {
	this := p
	_ = this

	localctx = NewConstructorDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RoseParserRULE_constructorDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(RoseParserIDENTIFIER)
	}
	{
		p.SetState(214)
		p.FormalParameters()
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RoseParserTHROWS {
		{
			p.SetState(215)
			p.Match(RoseParserTHROWS)
		}
		{
			p.SetState(216)
			p.QualifiedNameList()
		}

	}
	{
		p.SetState(219)

		var _x = p.Block()

		localctx.(*ConstructorDeclarationContext).constructorBody = _x
	}

	return localctx
}

// IVariableDeclaratorsContext is an interface to support dynamic dispatch.
type IVariableDeclaratorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorsContext differentiates from other interfaces.
	IsVariableDeclaratorsContext()
}

type VariableDeclaratorsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorsContext() *VariableDeclaratorsContext {
	var p = new(VariableDeclaratorsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_variableDeclarators
	return p
}

func (*VariableDeclaratorsContext) IsVariableDeclaratorsContext() {}

func NewVariableDeclaratorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorsContext {
	var p = new(VariableDeclaratorsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_variableDeclarators

	return p
}

func (s *VariableDeclaratorsContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorsContext) TypeType() ITypeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *VariableDeclaratorsContext) AllVariableDeclarator() []IVariableDeclaratorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableDeclaratorContext); ok {
			len++
		}
	}

	tst := make([]IVariableDeclaratorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableDeclaratorContext); ok {
			tst[i] = t.(IVariableDeclaratorContext)
			i++
		}
	}

	return tst
}

func (s *VariableDeclaratorsContext) VariableDeclarator(i int) IVariableDeclaratorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclaratorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorContext)
}

func (s *VariableDeclaratorsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RoseParserCOMMA)
}

func (s *VariableDeclaratorsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RoseParserCOMMA, i)
}

func (s *VariableDeclaratorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterVariableDeclarators(s)
	}
}

func (s *VariableDeclaratorsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitVariableDeclarators(s)
	}
}

func (s *VariableDeclaratorsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitVariableDeclarators(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) VariableDeclarators() (localctx IVariableDeclaratorsContext) {
	this := p
	_ = this

	localctx = NewVariableDeclaratorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, RoseParserRULE_variableDeclarators)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.TypeType()
	}
	{
		p.SetState(222)
		p.VariableDeclarator()
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RoseParserCOMMA {
		{
			p.SetState(223)
			p.Match(RoseParserCOMMA)
		}
		{
			p.SetState(224)
			p.VariableDeclarator()
		}

		p.SetState(229)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVariableDeclaratorContext is an interface to support dynamic dispatch.
type IVariableDeclaratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorContext differentiates from other interfaces.
	IsVariableDeclaratorContext()
}

type VariableDeclaratorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorContext() *VariableDeclaratorContext {
	var p = new(VariableDeclaratorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_variableDeclarator
	return p
}

func (*VariableDeclaratorContext) IsVariableDeclaratorContext() {}

func NewVariableDeclaratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorContext {
	var p = new(VariableDeclaratorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_variableDeclarator

	return p
}

func (s *VariableDeclaratorContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorContext) VariableDeclaratorId() IVariableDeclaratorIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclaratorIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorIdContext)
}

func (s *VariableDeclaratorContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(RoseParserASSIGN, 0)
}

func (s *VariableDeclaratorContext) VariableInitializer() IVariableInitializerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableInitializerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableInitializerContext)
}

func (s *VariableDeclaratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterVariableDeclarator(s)
	}
}

func (s *VariableDeclaratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitVariableDeclarator(s)
	}
}

func (s *VariableDeclaratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitVariableDeclarator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) VariableDeclarator() (localctx IVariableDeclaratorContext) {
	this := p
	_ = this

	localctx = NewVariableDeclaratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RoseParserRULE_variableDeclarator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.VariableDeclaratorId()
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RoseParserASSIGN {
		{
			p.SetState(231)
			p.Match(RoseParserASSIGN)
		}
		{
			p.SetState(232)
			p.VariableInitializer()
		}

	}

	return localctx
}

// IVariableDeclaratorIdContext is an interface to support dynamic dispatch.
type IVariableDeclaratorIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorIdContext differentiates from other interfaces.
	IsVariableDeclaratorIdContext()
}

type VariableDeclaratorIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorIdContext() *VariableDeclaratorIdContext {
	var p = new(VariableDeclaratorIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_variableDeclaratorId
	return p
}

func (*VariableDeclaratorIdContext) IsVariableDeclaratorIdContext() {}

func NewVariableDeclaratorIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorIdContext {
	var p = new(VariableDeclaratorIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_variableDeclaratorId

	return p
}

func (s *VariableDeclaratorIdContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorIdContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RoseParserIDENTIFIER, 0)
}

func (s *VariableDeclaratorIdContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(RoseParserLBRACK)
}

func (s *VariableDeclaratorIdContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(RoseParserLBRACK, i)
}

func (s *VariableDeclaratorIdContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(RoseParserRBRACK)
}

func (s *VariableDeclaratorIdContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(RoseParserRBRACK, i)
}

func (s *VariableDeclaratorIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterVariableDeclaratorId(s)
	}
}

func (s *VariableDeclaratorIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitVariableDeclaratorId(s)
	}
}

func (s *VariableDeclaratorIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitVariableDeclaratorId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) VariableDeclaratorId() (localctx IVariableDeclaratorIdContext) {
	this := p
	_ = this

	localctx = NewVariableDeclaratorIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, RoseParserRULE_variableDeclaratorId)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(RoseParserIDENTIFIER)
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RoseParserLBRACK {
		{
			p.SetState(236)
			p.Match(RoseParserLBRACK)
		}
		{
			p.SetState(237)
			p.Match(RoseParserRBRACK)
		}

		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVariableInitializerContext is an interface to support dynamic dispatch.
type IVariableInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableInitializerContext differentiates from other interfaces.
	IsVariableInitializerContext()
}

type VariableInitializerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableInitializerContext() *VariableInitializerContext {
	var p = new(VariableInitializerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_variableInitializer
	return p
}

func (*VariableInitializerContext) IsVariableInitializerContext() {}

func NewVariableInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableInitializerContext {
	var p = new(VariableInitializerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_variableInitializer

	return p
}

func (s *VariableInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableInitializerContext) ArrayInitializer() IArrayInitializerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayInitializerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayInitializerContext)
}

func (s *VariableInitializerContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableInitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterVariableInitializer(s)
	}
}

func (s *VariableInitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitVariableInitializer(s)
	}
}

func (s *VariableInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitVariableInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) VariableInitializer() (localctx IVariableInitializerContext) {
	this := p
	_ = this

	localctx = NewVariableInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, RoseParserRULE_variableInitializer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(245)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RoseParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(243)
			p.ArrayInitializer()
		}

	case RoseParserSUPER, RoseParserTHIS, RoseParserDECIMAL_LITERAL, RoseParserHEX_LITERAL, RoseParserOCT_LITERAL, RoseParserBINARY_LITERAL, RoseParserFLOAT_LITERAL, RoseParserHEX_FLOAT_LITERAL, RoseParserBOOL_LITERAL, RoseParserCHAR_LITERAL, RoseParserSTRING_LITERAL, RoseParserNULL_LITERAL, RoseParserLPAREN, RoseParserBANG, RoseParserTILDE, RoseParserINC, RoseParserDEC, RoseParserADD, RoseParserSUB, RoseParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(244)
			p.expression(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArrayInitializerContext is an interface to support dynamic dispatch.
type IArrayInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayInitializerContext differentiates from other interfaces.
	IsArrayInitializerContext()
}

type ArrayInitializerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayInitializerContext() *ArrayInitializerContext {
	var p = new(ArrayInitializerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_arrayInitializer
	return p
}

func (*ArrayInitializerContext) IsArrayInitializerContext() {}

func NewArrayInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayInitializerContext {
	var p = new(ArrayInitializerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_arrayInitializer

	return p
}

func (s *ArrayInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayInitializerContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(RoseParserLBRACE, 0)
}

func (s *ArrayInitializerContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(RoseParserRBRACE, 0)
}

func (s *ArrayInitializerContext) AllVariableInitializer() []IVariableInitializerContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableInitializerContext); ok {
			len++
		}
	}

	tst := make([]IVariableInitializerContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableInitializerContext); ok {
			tst[i] = t.(IVariableInitializerContext)
			i++
		}
	}

	return tst
}

func (s *ArrayInitializerContext) VariableInitializer(i int) IVariableInitializerContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableInitializerContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableInitializerContext)
}

func (s *ArrayInitializerContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RoseParserCOMMA)
}

func (s *ArrayInitializerContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RoseParserCOMMA, i)
}

func (s *ArrayInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayInitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterArrayInitializer(s)
	}
}

func (s *ArrayInitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitArrayInitializer(s)
	}
}

func (s *ArrayInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitArrayInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) ArrayInitializer() (localctx IArrayInitializerContext) {
	this := p
	_ = this

	localctx = NewArrayInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, RoseParserRULE_arrayInitializer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(RoseParserLBRACE)
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8997303650091008) != 0 || (int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&281474992442369) != 0 {
		{
			p.SetState(248)
			p.VariableInitializer()
		}
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(249)
					p.Match(RoseParserCOMMA)
				}
				{
					p.SetState(250)
					p.VariableInitializer()
				}

			}
			p.SetState(255)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
		}
		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RoseParserCOMMA {
			{
				p.SetState(256)
				p.Match(RoseParserCOMMA)
			}

		}

	}
	{
		p.SetState(261)
		p.Match(RoseParserRBRACE)
	}

	return localctx
}

// IClassOrInterfaceTypeContext is an interface to support dynamic dispatch.
type IClassOrInterfaceTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassOrInterfaceTypeContext differentiates from other interfaces.
	IsClassOrInterfaceTypeContext()
}

type ClassOrInterfaceTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassOrInterfaceTypeContext() *ClassOrInterfaceTypeContext {
	var p = new(ClassOrInterfaceTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_classOrInterfaceType
	return p
}

func (*ClassOrInterfaceTypeContext) IsClassOrInterfaceTypeContext() {}

func NewClassOrInterfaceTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassOrInterfaceTypeContext {
	var p = new(ClassOrInterfaceTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_classOrInterfaceType

	return p
}

func (s *ClassOrInterfaceTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassOrInterfaceTypeContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(RoseParserIDENTIFIER)
}

func (s *ClassOrInterfaceTypeContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(RoseParserIDENTIFIER, i)
}

func (s *ClassOrInterfaceTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(RoseParserDOT)
}

func (s *ClassOrInterfaceTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(RoseParserDOT, i)
}

func (s *ClassOrInterfaceTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassOrInterfaceTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassOrInterfaceTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterClassOrInterfaceType(s)
	}
}

func (s *ClassOrInterfaceTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitClassOrInterfaceType(s)
	}
}

func (s *ClassOrInterfaceTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitClassOrInterfaceType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) ClassOrInterfaceType() (localctx IClassOrInterfaceTypeContext) {
	this := p
	_ = this

	localctx = NewClassOrInterfaceTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, RoseParserRULE_classOrInterfaceType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.Match(RoseParserIDENTIFIER)
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(264)
				p.Match(RoseParserDOT)
			}
			{
				p.SetState(265)
				p.Match(RoseParserIDENTIFIER)
			}

		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
	}

	return localctx
}

// ITypeArgumentContext is an interface to support dynamic dispatch.
type ITypeArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeArgumentContext differentiates from other interfaces.
	IsTypeArgumentContext()
}

type TypeArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeArgumentContext() *TypeArgumentContext {
	var p = new(TypeArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_typeArgument
	return p
}

func (*TypeArgumentContext) IsTypeArgumentContext() {}

func NewTypeArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeArgumentContext {
	var p = new(TypeArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_typeArgument

	return p
}

func (s *TypeArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeArgumentContext) TypeType() ITypeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *TypeArgumentContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(RoseParserQUESTION, 0)
}

func (s *TypeArgumentContext) EXTENDS() antlr.TerminalNode {
	return s.GetToken(RoseParserEXTENDS, 0)
}

func (s *TypeArgumentContext) SUPER() antlr.TerminalNode {
	return s.GetToken(RoseParserSUPER, 0)
}

func (s *TypeArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterTypeArgument(s)
	}
}

func (s *TypeArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitTypeArgument(s)
	}
}

func (s *TypeArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitTypeArgument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) TypeArgument() (localctx ITypeArgumentContext) {
	this := p
	_ = this

	localctx = NewTypeArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, RoseParserRULE_typeArgument)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(277)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RoseParserBOOLEAN, RoseParserBYTE, RoseParserCHAR, RoseParserDOUBLE, RoseParserFLOAT, RoseParserINT, RoseParserLONG, RoseParserSHORT, RoseParserFUNCTION, RoseParserSTRING, RoseParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)
			p.TypeType()
		}

	case RoseParserQUESTION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(272)
			p.Match(RoseParserQUESTION)
		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RoseParserEXTENDS || _la == RoseParserSUPER {
			{
				p.SetState(273)
				_la = p.GetTokenStream().LA(1)

				if !(_la == RoseParserEXTENDS || _la == RoseParserSUPER) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(274)
				p.TypeType()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) IntegerLiteral() IIntegerLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *LiteralContext) FloatLiteral() IFloatLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatLiteralContext)
}

func (s *LiteralContext) CHAR_LITERAL() antlr.TerminalNode {
	return s.GetToken(RoseParserCHAR_LITERAL, 0)
}

func (s *LiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(RoseParserSTRING_LITERAL, 0)
}

func (s *LiteralContext) BOOL_LITERAL() antlr.TerminalNode {
	return s.GetToken(RoseParserBOOL_LITERAL, 0)
}

func (s *LiteralContext) NULL_LITERAL() antlr.TerminalNode {
	return s.GetToken(RoseParserNULL_LITERAL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, RoseParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(285)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RoseParserDECIMAL_LITERAL, RoseParserHEX_LITERAL, RoseParserOCT_LITERAL, RoseParserBINARY_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(279)
			p.IntegerLiteral()
		}

	case RoseParserFLOAT_LITERAL, RoseParserHEX_FLOAT_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(280)
			p.FloatLiteral()
		}

	case RoseParserCHAR_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(281)
			p.Match(RoseParserCHAR_LITERAL)
		}

	case RoseParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(282)
			p.Match(RoseParserSTRING_LITERAL)
		}

	case RoseParserBOOL_LITERAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(283)
			p.Match(RoseParserBOOL_LITERAL)
		}

	case RoseParserNULL_LITERAL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(284)
			p.Match(RoseParserNULL_LITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntegerLiteralContext is an interface to support dynamic dispatch.
type IIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerLiteralContext differentiates from other interfaces.
	IsIntegerLiteralContext()
}

type IntegerLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerLiteralContext() *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_integerLiteral
	return p
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(RoseParserDECIMAL_LITERAL, 0)
}

func (s *IntegerLiteralContext) HEX_LITERAL() antlr.TerminalNode {
	return s.GetToken(RoseParserHEX_LITERAL, 0)
}

func (s *IntegerLiteralContext) OCT_LITERAL() antlr.TerminalNode {
	return s.GetToken(RoseParserOCT_LITERAL, 0)
}

func (s *IntegerLiteralContext) BINARY_LITERAL() antlr.TerminalNode {
	return s.GetToken(RoseParserBINARY_LITERAL, 0)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitIntegerLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	this := p
	_ = this

	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, RoseParserRULE_integerLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135107988821114880) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFloatLiteralContext is an interface to support dynamic dispatch.
type IFloatLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatLiteralContext differentiates from other interfaces.
	IsFloatLiteralContext()
}

type FloatLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLiteralContext() *FloatLiteralContext {
	var p = new(FloatLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_floatLiteral
	return p
}

func (*FloatLiteralContext) IsFloatLiteralContext() {}

func NewFloatLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_floatLiteral

	return p
}

func (s *FloatLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLiteralContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(RoseParserFLOAT_LITERAL, 0)
}

func (s *FloatLiteralContext) HEX_FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(RoseParserHEX_FLOAT_LITERAL, 0)
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) FloatLiteral() (localctx IFloatLiteralContext) {
	this := p
	_ = this

	localctx = NewFloatLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, RoseParserRULE_floatLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RoseParserFLOAT_LITERAL || _la == RoseParserHEX_FLOAT_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) BlockStatements() IBlockStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementsContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, RoseParserRULE_prog)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.BlockStatements()
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(RoseParserLBRACE, 0)
}

func (s *BlockContext) BlockStatements() IBlockStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementsContext)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(RoseParserRBRACE, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, RoseParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(RoseParserLBRACE)
	}
	{
		p.SetState(294)
		p.BlockStatements()
	}
	{
		p.SetState(295)
		p.Match(RoseParserRBRACE)
	}

	return localctx
}

// IBlockStatementsContext is an interface to support dynamic dispatch.
type IBlockStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockStatementsContext differentiates from other interfaces.
	IsBlockStatementsContext()
}

type BlockStatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementsContext() *BlockStatementsContext {
	var p = new(BlockStatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_blockStatements
	return p
}

func (*BlockStatementsContext) IsBlockStatementsContext() {}

func NewBlockStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementsContext {
	var p = new(BlockStatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_blockStatements

	return p
}

func (s *BlockStatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementsContext) AllBlockStatement() []IBlockStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockStatementContext); ok {
			len++
		}
	}

	tst := make([]IBlockStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockStatementContext); ok {
			tst[i] = t.(IBlockStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockStatementsContext) BlockStatement(i int) IBlockStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *BlockStatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterBlockStatements(s)
	}
}

func (s *BlockStatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitBlockStatements(s)
	}
}

func (s *BlockStatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitBlockStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) BlockStatements() (localctx IBlockStatementsContext) {
	this := p
	_ = this

	localctx = NewBlockStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, RoseParserRULE_blockStatements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-832123465340104) != 0 || (int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&281474992442385) != 0 {
		{
			p.SetState(297)
			p.BlockStatement()
		}

		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBlockStatementContext is an interface to support dynamic dispatch.
type IBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockStatementContext differentiates from other interfaces.
	IsBlockStatementContext()
}

type BlockStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementContext() *BlockStatementContext {
	var p = new(BlockStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_blockStatement
	return p
}

func (*BlockStatementContext) IsBlockStatementContext() {}

func NewBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementContext {
	var p = new(BlockStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_blockStatement

	return p
}

func (s *BlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementContext) VariableDeclarators() IVariableDeclaratorsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclaratorsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorsContext)
}

func (s *BlockStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RoseParserSEMI, 0)
}

func (s *BlockStatementContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockStatementContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *BlockStatementContext) ClassDeclaration() IClassDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassDeclarationContext)
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterBlockStatement(s)
	}
}

func (s *BlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitBlockStatement(s)
	}
}

func (s *BlockStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitBlockStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) BlockStatement() (localctx IBlockStatementContext) {
	this := p
	_ = this

	localctx = NewBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, RoseParserRULE_blockStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(303)
			p.VariableDeclarators()
		}
		{
			p.SetState(304)
			p.Match(RoseParserSEMI)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(306)
			p.Statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(307)
			p.FunctionDeclaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(308)
			p.ClassDeclaration()
		}

	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdentifierLabel returns the identifierLabel token.
	GetIdentifierLabel() antlr.Token

	// SetIdentifierLabel sets the identifierLabel token.
	SetIdentifierLabel(antlr.Token)

	// GetBlockLabel returns the blockLabel rule contexts.
	GetBlockLabel() IBlockContext

	// GetStatementExpression returns the statementExpression rule contexts.
	GetStatementExpression() IExpressionContext

	// SetBlockLabel sets the blockLabel rule contexts.
	SetBlockLabel(IBlockContext)

	// SetStatementExpression sets the statementExpression rule contexts.
	SetStatementExpression(IExpressionContext)

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser              antlr.Parser
	blockLabel          IBlockContext
	statementExpression IExpressionContext
	identifierLabel     antlr.Token
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) GetIdentifierLabel() antlr.Token { return s.identifierLabel }

func (s *StatementContext) SetIdentifierLabel(v antlr.Token) { s.identifierLabel = v }

func (s *StatementContext) GetBlockLabel() IBlockContext { return s.blockLabel }

func (s *StatementContext) GetStatementExpression() IExpressionContext { return s.statementExpression }

func (s *StatementContext) SetBlockLabel(v IBlockContext) { s.blockLabel = v }

func (s *StatementContext) SetStatementExpression(v IExpressionContext) { s.statementExpression = v }

func (s *StatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) IF() antlr.TerminalNode {
	return s.GetToken(RoseParserIF, 0)
}

func (s *StatementContext) ParExpression() IParExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParExpressionContext)
}

func (s *StatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(RoseParserELSE, 0)
}

func (s *StatementContext) FOR() antlr.TerminalNode {
	return s.GetToken(RoseParserFOR, 0)
}

func (s *StatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RoseParserLPAREN, 0)
}

func (s *StatementContext) ForControl() IForControlContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForControlContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForControlContext)
}

func (s *StatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RoseParserRPAREN, 0)
}

func (s *StatementContext) WHILE() antlr.TerminalNode {
	return s.GetToken(RoseParserWHILE, 0)
}

func (s *StatementContext) DO() antlr.TerminalNode {
	return s.GetToken(RoseParserDO, 0)
}

func (s *StatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(RoseParserSEMI, 0)
}

func (s *StatementContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(RoseParserSWITCH, 0)
}

func (s *StatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(RoseParserLBRACE, 0)
}

func (s *StatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(RoseParserRBRACE, 0)
}

func (s *StatementContext) AllSwitchBlockStatementGroup() []ISwitchBlockStatementGroupContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitchBlockStatementGroupContext); ok {
			len++
		}
	}

	tst := make([]ISwitchBlockStatementGroupContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitchBlockStatementGroupContext); ok {
			tst[i] = t.(ISwitchBlockStatementGroupContext)
			i++
		}
	}

	return tst
}

func (s *StatementContext) SwitchBlockStatementGroup(i int) ISwitchBlockStatementGroupContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchBlockStatementGroupContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchBlockStatementGroupContext)
}

func (s *StatementContext) AllSwitchLabel() []ISwitchLabelContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitchLabelContext); ok {
			len++
		}
	}

	tst := make([]ISwitchLabelContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitchLabelContext); ok {
			tst[i] = t.(ISwitchLabelContext)
			i++
		}
	}

	return tst
}

func (s *StatementContext) SwitchLabel(i int) ISwitchLabelContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchLabelContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchLabelContext)
}

func (s *StatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(RoseParserRETURN, 0)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) BREAK() antlr.TerminalNode {
	return s.GetToken(RoseParserBREAK, 0)
}

func (s *StatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RoseParserIDENTIFIER, 0)
}

func (s *StatementContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(RoseParserCONTINUE, 0)
}

func (s *StatementContext) COLON() antlr.TerminalNode {
	return s.GetToken(RoseParserCOLON, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, RoseParserRULE_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(311)

			var _x = p.Block()

			localctx.(*StatementContext).blockLabel = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(312)
			p.Match(RoseParserIF)
		}
		{
			p.SetState(313)
			p.ParExpression()
		}
		{
			p.SetState(314)
			p.Statement()
		}
		p.SetState(317)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(315)
				p.Match(RoseParserELSE)
			}
			{
				p.SetState(316)
				p.Statement()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(319)
			p.Match(RoseParserFOR)
		}
		{
			p.SetState(320)
			p.Match(RoseParserLPAREN)
		}
		{
			p.SetState(321)
			p.ForControl()
		}
		{
			p.SetState(322)
			p.Match(RoseParserRPAREN)
		}
		{
			p.SetState(323)
			p.Statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(325)
			p.Match(RoseParserWHILE)
		}
		{
			p.SetState(326)
			p.ParExpression()
		}
		{
			p.SetState(327)
			p.Statement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(329)
			p.Match(RoseParserDO)
		}
		{
			p.SetState(330)
			p.Statement()
		}
		{
			p.SetState(331)
			p.Match(RoseParserWHILE)
		}
		{
			p.SetState(332)
			p.ParExpression()
		}
		{
			p.SetState(333)
			p.Match(RoseParserSEMI)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(335)
			p.Match(RoseParserSWITCH)
		}
		{
			p.SetState(336)
			p.ParExpression()
		}
		{
			p.SetState(337)
			p.Match(RoseParserLBRACE)
		}
		p.SetState(341)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(338)
					p.SwitchBlockStatementGroup()
				}

			}
			p.SetState(343)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
		}
		p.SetState(347)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RoseParserCASE || _la == RoseParserDEFAULT {
			{
				p.SetState(344)
				p.SwitchLabel()
			}

			p.SetState(349)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(350)
			p.Match(RoseParserRBRACE)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(352)
			p.Match(RoseParserRETURN)
		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8997303650091008) != 0 || (int64((_la-75)) & ^0x3f) == 0 && ((int64(1)<<(_la-75))&274877922307) != 0 {
			{
				p.SetState(353)
				p.expression(0)
			}

		}
		{
			p.SetState(356)
			p.Match(RoseParserSEMI)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(357)
			p.Match(RoseParserBREAK)
		}
		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RoseParserIDENTIFIER {
			{
				p.SetState(358)
				p.Match(RoseParserIDENTIFIER)
			}

		}
		{
			p.SetState(361)
			p.Match(RoseParserSEMI)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(362)
			p.Match(RoseParserCONTINUE)
		}
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RoseParserIDENTIFIER {
			{
				p.SetState(363)
				p.Match(RoseParserIDENTIFIER)
			}

		}
		{
			p.SetState(366)
			p.Match(RoseParserSEMI)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(367)
			p.Match(RoseParserSEMI)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(368)

			var _x = p.expression(0)

			localctx.(*StatementContext).statementExpression = _x
		}
		{
			p.SetState(369)
			p.Match(RoseParserSEMI)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(371)

			var _m = p.Match(RoseParserIDENTIFIER)

			localctx.(*StatementContext).identifierLabel = _m
		}
		{
			p.SetState(372)
			p.Match(RoseParserCOLON)
		}
		{
			p.SetState(373)
			p.Statement()
		}

	}

	return localctx
}

// ISwitchBlockStatementGroupContext is an interface to support dynamic dispatch.
type ISwitchBlockStatementGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchBlockStatementGroupContext differentiates from other interfaces.
	IsSwitchBlockStatementGroupContext()
}

type SwitchBlockStatementGroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchBlockStatementGroupContext() *SwitchBlockStatementGroupContext {
	var p = new(SwitchBlockStatementGroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_switchBlockStatementGroup
	return p
}

func (*SwitchBlockStatementGroupContext) IsSwitchBlockStatementGroupContext() {}

func NewSwitchBlockStatementGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchBlockStatementGroupContext {
	var p = new(SwitchBlockStatementGroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_switchBlockStatementGroup

	return p
}

func (s *SwitchBlockStatementGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchBlockStatementGroupContext) AllSwitchLabel() []ISwitchLabelContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitchLabelContext); ok {
			len++
		}
	}

	tst := make([]ISwitchLabelContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitchLabelContext); ok {
			tst[i] = t.(ISwitchLabelContext)
			i++
		}
	}

	return tst
}

func (s *SwitchBlockStatementGroupContext) SwitchLabel(i int) ISwitchLabelContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchLabelContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchLabelContext)
}

func (s *SwitchBlockStatementGroupContext) AllBlockStatement() []IBlockStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockStatementContext); ok {
			len++
		}
	}

	tst := make([]IBlockStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockStatementContext); ok {
			tst[i] = t.(IBlockStatementContext)
			i++
		}
	}

	return tst
}

func (s *SwitchBlockStatementGroupContext) BlockStatement(i int) IBlockStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *SwitchBlockStatementGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchBlockStatementGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchBlockStatementGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterSwitchBlockStatementGroup(s)
	}
}

func (s *SwitchBlockStatementGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitSwitchBlockStatementGroup(s)
	}
}

func (s *SwitchBlockStatementGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitSwitchBlockStatementGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) SwitchBlockStatementGroup() (localctx ISwitchBlockStatementGroupContext) {
	this := p
	_ = this

	localctx = NewSwitchBlockStatementGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, RoseParserRULE_switchBlockStatementGroup)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(377)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == RoseParserCASE || _la == RoseParserDEFAULT {
		{
			p.SetState(376)
			p.SwitchLabel()
		}

		p.SetState(379)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-832123465340104) != 0 || (int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&281474992442385) != 0 {
		{
			p.SetState(381)
			p.BlockStatement()
		}

		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISwitchLabelContext is an interface to support dynamic dispatch.
type ISwitchLabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEnumConstantName returns the enumConstantName token.
	GetEnumConstantName() antlr.Token

	// SetEnumConstantName sets the enumConstantName token.
	SetEnumConstantName(antlr.Token)

	// GetConstantExpression returns the constantExpression rule contexts.
	GetConstantExpression() IExpressionContext

	// SetConstantExpression sets the constantExpression rule contexts.
	SetConstantExpression(IExpressionContext)

	// IsSwitchLabelContext differentiates from other interfaces.
	IsSwitchLabelContext()
}

type SwitchLabelContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	constantExpression IExpressionContext
	enumConstantName   antlr.Token
}

func NewEmptySwitchLabelContext() *SwitchLabelContext {
	var p = new(SwitchLabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_switchLabel
	return p
}

func (*SwitchLabelContext) IsSwitchLabelContext() {}

func NewSwitchLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchLabelContext {
	var p = new(SwitchLabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_switchLabel

	return p
}

func (s *SwitchLabelContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchLabelContext) GetEnumConstantName() antlr.Token { return s.enumConstantName }

func (s *SwitchLabelContext) SetEnumConstantName(v antlr.Token) { s.enumConstantName = v }

func (s *SwitchLabelContext) GetConstantExpression() IExpressionContext { return s.constantExpression }

func (s *SwitchLabelContext) SetConstantExpression(v IExpressionContext) { s.constantExpression = v }

func (s *SwitchLabelContext) CASE() antlr.TerminalNode {
	return s.GetToken(RoseParserCASE, 0)
}

func (s *SwitchLabelContext) COLON() antlr.TerminalNode {
	return s.GetToken(RoseParserCOLON, 0)
}

func (s *SwitchLabelContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SwitchLabelContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RoseParserIDENTIFIER, 0)
}

func (s *SwitchLabelContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(RoseParserDEFAULT, 0)
}

func (s *SwitchLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchLabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchLabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterSwitchLabel(s)
	}
}

func (s *SwitchLabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitSwitchLabel(s)
	}
}

func (s *SwitchLabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitSwitchLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) SwitchLabel() (localctx ISwitchLabelContext) {
	this := p
	_ = this

	localctx = NewSwitchLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, RoseParserRULE_switchLabel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(394)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RoseParserCASE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(386)
			p.Match(RoseParserCASE)
		}
		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(387)

				var _x = p.expression(0)

				localctx.(*SwitchLabelContext).constantExpression = _x
			}

		case 2:
			{
				p.SetState(388)

				var _m = p.Match(RoseParserIDENTIFIER)

				localctx.(*SwitchLabelContext).enumConstantName = _m
			}

		}
		{
			p.SetState(391)
			p.Match(RoseParserCOLON)
		}

	case RoseParserDEFAULT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(392)
			p.Match(RoseParserDEFAULT)
		}
		{
			p.SetState(393)
			p.Match(RoseParserCOLON)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IForControlContext is an interface to support dynamic dispatch.
type IForControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetForUpdate returns the forUpdate rule contexts.
	GetForUpdate() IExpressionListContext

	// SetForUpdate sets the forUpdate rule contexts.
	SetForUpdate(IExpressionListContext)

	// IsForControlContext differentiates from other interfaces.
	IsForControlContext()
}

type ForControlContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	forUpdate IExpressionListContext
}

func NewEmptyForControlContext() *ForControlContext {
	var p = new(ForControlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_forControl
	return p
}

func (*ForControlContext) IsForControlContext() {}

func NewForControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForControlContext {
	var p = new(ForControlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_forControl

	return p
}

func (s *ForControlContext) GetParser() antlr.Parser { return s.parser }

func (s *ForControlContext) GetForUpdate() IExpressionListContext { return s.forUpdate }

func (s *ForControlContext) SetForUpdate(v IExpressionListContext) { s.forUpdate = v }

func (s *ForControlContext) EnhancedForControl() IEnhancedForControlContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnhancedForControlContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnhancedForControlContext)
}

func (s *ForControlContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(RoseParserSEMI)
}

func (s *ForControlContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(RoseParserSEMI, i)
}

func (s *ForControlContext) ForInit() IForInitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForInitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForInitContext)
}

func (s *ForControlContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForControlContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ForControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForControlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterForControl(s)
	}
}

func (s *ForControlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitForControl(s)
	}
}

func (s *ForControlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitForControl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) ForControl() (localctx IForControlContext) {
	this := p
	_ = this

	localctx = NewForControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, RoseParserRULE_forControl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(396)
			p.EnhancedForControl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2241766097927896) != 0 || (int64((_la-75)) & ^0x3f) == 0 && ((int64(1)<<(_la-75))&274877922307) != 0 {
			{
				p.SetState(397)
				p.ForInit()
			}

		}
		{
			p.SetState(400)
			p.Match(RoseParserSEMI)
		}
		p.SetState(402)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8997303650091008) != 0 || (int64((_la-75)) & ^0x3f) == 0 && ((int64(1)<<(_la-75))&274877922307) != 0 {
			{
				p.SetState(401)
				p.expression(0)
			}

		}
		{
			p.SetState(404)
			p.Match(RoseParserSEMI)
		}
		p.SetState(406)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8997303650091008) != 0 || (int64((_la-75)) & ^0x3f) == 0 && ((int64(1)<<(_la-75))&274877922307) != 0 {
			{
				p.SetState(405)

				var _x = p.ExpressionList()

				localctx.(*ForControlContext).forUpdate = _x
			}

		}

	}

	return localctx
}

// IForInitContext is an interface to support dynamic dispatch.
type IForInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForInitContext differentiates from other interfaces.
	IsForInitContext()
}

type ForInitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForInitContext() *ForInitContext {
	var p = new(ForInitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_forInit
	return p
}

func (*ForInitContext) IsForInitContext() {}

func NewForInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForInitContext {
	var p = new(ForInitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_forInit

	return p
}

func (s *ForInitContext) GetParser() antlr.Parser { return s.parser }

func (s *ForInitContext) VariableDeclarators() IVariableDeclaratorsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclaratorsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorsContext)
}

func (s *ForInitContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ForInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterForInit(s)
	}
}

func (s *ForInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitForInit(s)
	}
}

func (s *ForInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitForInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) ForInit() (localctx IForInitContext) {
	this := p
	_ = this

	localctx = NewForInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, RoseParserRULE_forInit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(410)
			p.VariableDeclarators()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(411)
			p.ExpressionList()
		}

	}

	return localctx
}

// IEnhancedForControlContext is an interface to support dynamic dispatch.
type IEnhancedForControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnhancedForControlContext differentiates from other interfaces.
	IsEnhancedForControlContext()
}

type EnhancedForControlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnhancedForControlContext() *EnhancedForControlContext {
	var p = new(EnhancedForControlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_enhancedForControl
	return p
}

func (*EnhancedForControlContext) IsEnhancedForControlContext() {}

func NewEnhancedForControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnhancedForControlContext {
	var p = new(EnhancedForControlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_enhancedForControl

	return p
}

func (s *EnhancedForControlContext) GetParser() antlr.Parser { return s.parser }

func (s *EnhancedForControlContext) TypeType() ITypeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *EnhancedForControlContext) VariableDeclaratorId() IVariableDeclaratorIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclaratorIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorIdContext)
}

func (s *EnhancedForControlContext) COLON() antlr.TerminalNode {
	return s.GetToken(RoseParserCOLON, 0)
}

func (s *EnhancedForControlContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EnhancedForControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnhancedForControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnhancedForControlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterEnhancedForControl(s)
	}
}

func (s *EnhancedForControlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitEnhancedForControl(s)
	}
}

func (s *EnhancedForControlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitEnhancedForControl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) EnhancedForControl() (localctx IEnhancedForControlContext) {
	this := p
	_ = this

	localctx = NewEnhancedForControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, RoseParserRULE_enhancedForControl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(414)
		p.TypeType()
	}
	{
		p.SetState(415)
		p.VariableDeclaratorId()
	}
	{
		p.SetState(416)
		p.Match(RoseParserCOLON)
	}
	{
		p.SetState(417)
		p.expression(0)
	}

	return localctx
}

// IParExpressionContext is an interface to support dynamic dispatch.
type IParExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParExpressionContext differentiates from other interfaces.
	IsParExpressionContext()
}

type ParExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParExpressionContext() *ParExpressionContext {
	var p = new(ParExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_parExpression
	return p
}

func (*ParExpressionContext) IsParExpressionContext() {}

func NewParExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParExpressionContext {
	var p = new(ParExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_parExpression

	return p
}

func (s *ParExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ParExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RoseParserLPAREN, 0)
}

func (s *ParExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RoseParserRPAREN, 0)
}

func (s *ParExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterParExpression(s)
	}
}

func (s *ParExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitParExpression(s)
	}
}

func (s *ParExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitParExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) ParExpression() (localctx IParExpressionContext) {
	this := p
	_ = this

	localctx = NewParExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, RoseParserRULE_parExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(419)
		p.Match(RoseParserLPAREN)
	}
	{
		p.SetState(420)
		p.expression(0)
	}
	{
		p.SetState(421)
		p.Match(RoseParserRPAREN)
	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RoseParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RoseParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) ExpressionList() (localctx IExpressionListContext) {
	this := p
	_ = this

	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, RoseParserRULE_expressionList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(423)
		p.expression(0)
	}
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RoseParserCOMMA {
		{
			p.SetState(424)
			p.Match(RoseParserCOMMA)
		}
		{
			p.SetState(425)
			p.expression(0)
		}

		p.SetState(430)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RoseParserIDENTIFIER, 0)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RoseParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RoseParserRPAREN, 0)
}

func (s *FunctionCallContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *FunctionCallContext) THIS() antlr.TerminalNode {
	return s.GetToken(RoseParserTHIS, 0)
}

func (s *FunctionCallContext) SUPER() antlr.TerminalNode {
	return s.GetToken(RoseParserSUPER, 0)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) FunctionCall() (localctx IFunctionCallContext) {
	this := p
	_ = this

	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, RoseParserRULE_functionCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(449)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RoseParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(431)
			p.Match(RoseParserIDENTIFIER)
		}
		{
			p.SetState(432)
			p.Match(RoseParserLPAREN)
		}
		p.SetState(434)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8997303650091008) != 0 || (int64((_la-75)) & ^0x3f) == 0 && ((int64(1)<<(_la-75))&274877922307) != 0 {
			{
				p.SetState(433)
				p.ExpressionList()
			}

		}
		{
			p.SetState(436)
			p.Match(RoseParserRPAREN)
		}

	case RoseParserTHIS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(437)
			p.Match(RoseParserTHIS)
		}
		{
			p.SetState(438)
			p.Match(RoseParserLPAREN)
		}
		p.SetState(440)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8997303650091008) != 0 || (int64((_la-75)) & ^0x3f) == 0 && ((int64(1)<<(_la-75))&274877922307) != 0 {
			{
				p.SetState(439)
				p.ExpressionList()
			}

		}
		{
			p.SetState(442)
			p.Match(RoseParserRPAREN)
		}

	case RoseParserSUPER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(443)
			p.Match(RoseParserSUPER)
		}
		{
			p.SetState(444)
			p.Match(RoseParserLPAREN)
		}
		p.SetState(446)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8997303650091008) != 0 || (int64((_la-75)) & ^0x3f) == 0 && ((int64(1)<<(_la-75))&274877922307) != 0 {
			{
				p.SetState(445)
				p.ExpressionList()
			}

		}
		{
			p.SetState(448)
			p.Match(RoseParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPrefix returns the prefix token.
	GetPrefix() antlr.Token

	// GetBop returns the bop token.
	GetBop() antlr.Token

	// GetPostfix returns the postfix token.
	GetPostfix() antlr.Token

	// SetPrefix sets the prefix token.
	SetPrefix(antlr.Token)

	// SetBop sets the bop token.
	SetBop(antlr.Token)

	// SetPostfix sets the postfix token.
	SetPostfix(antlr.Token)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	prefix  antlr.Token
	bop     antlr.Token
	postfix antlr.Token
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetPrefix() antlr.Token { return s.prefix }

func (s *ExpressionContext) GetBop() antlr.Token { return s.bop }

func (s *ExpressionContext) GetPostfix() antlr.Token { return s.postfix }

func (s *ExpressionContext) SetPrefix(v antlr.Token) { s.prefix = v }

func (s *ExpressionContext) SetBop(v antlr.Token) { s.bop = v }

func (s *ExpressionContext) SetPostfix(v antlr.Token) { s.postfix = v }

func (s *ExpressionContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *ExpressionContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) ADD() antlr.TerminalNode {
	return s.GetToken(RoseParserADD, 0)
}

func (s *ExpressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(RoseParserSUB, 0)
}

func (s *ExpressionContext) INC() antlr.TerminalNode {
	return s.GetToken(RoseParserINC, 0)
}

func (s *ExpressionContext) DEC() antlr.TerminalNode {
	return s.GetToken(RoseParserDEC, 0)
}

func (s *ExpressionContext) TILDE() antlr.TerminalNode {
	return s.GetToken(RoseParserTILDE, 0)
}

func (s *ExpressionContext) BANG() antlr.TerminalNode {
	return s.GetToken(RoseParserBANG, 0)
}

func (s *ExpressionContext) MUL() antlr.TerminalNode {
	return s.GetToken(RoseParserMUL, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(RoseParserDIV, 0)
}

func (s *ExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(RoseParserMOD, 0)
}

func (s *ExpressionContext) AllLT() []antlr.TerminalNode {
	return s.GetTokens(RoseParserLT)
}

func (s *ExpressionContext) LT(i int) antlr.TerminalNode {
	return s.GetToken(RoseParserLT, i)
}

func (s *ExpressionContext) AllGT() []antlr.TerminalNode {
	return s.GetTokens(RoseParserGT)
}

func (s *ExpressionContext) GT(i int) antlr.TerminalNode {
	return s.GetToken(RoseParserGT, i)
}

func (s *ExpressionContext) LE() antlr.TerminalNode {
	return s.GetToken(RoseParserLE, 0)
}

func (s *ExpressionContext) GE() antlr.TerminalNode {
	return s.GetToken(RoseParserGE, 0)
}

func (s *ExpressionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(RoseParserEQUAL, 0)
}

func (s *ExpressionContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(RoseParserNOTEQUAL, 0)
}

func (s *ExpressionContext) BITAND() antlr.TerminalNode {
	return s.GetToken(RoseParserBITAND, 0)
}

func (s *ExpressionContext) CARET() antlr.TerminalNode {
	return s.GetToken(RoseParserCARET, 0)
}

func (s *ExpressionContext) BITOR() antlr.TerminalNode {
	return s.GetToken(RoseParserBITOR, 0)
}

func (s *ExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(RoseParserAND, 0)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(RoseParserOR, 0)
}

func (s *ExpressionContext) COLON() antlr.TerminalNode {
	return s.GetToken(RoseParserCOLON, 0)
}

func (s *ExpressionContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(RoseParserQUESTION, 0)
}

func (s *ExpressionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(RoseParserASSIGN, 0)
}

func (s *ExpressionContext) ADD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(RoseParserADD_ASSIGN, 0)
}

func (s *ExpressionContext) SUB_ASSIGN() antlr.TerminalNode {
	return s.GetToken(RoseParserSUB_ASSIGN, 0)
}

func (s *ExpressionContext) MUL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(RoseParserMUL_ASSIGN, 0)
}

func (s *ExpressionContext) DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(RoseParserDIV_ASSIGN, 0)
}

func (s *ExpressionContext) AND_ASSIGN() antlr.TerminalNode {
	return s.GetToken(RoseParserAND_ASSIGN, 0)
}

func (s *ExpressionContext) OR_ASSIGN() antlr.TerminalNode {
	return s.GetToken(RoseParserOR_ASSIGN, 0)
}

func (s *ExpressionContext) XOR_ASSIGN() antlr.TerminalNode {
	return s.GetToken(RoseParserXOR_ASSIGN, 0)
}

func (s *ExpressionContext) RSHIFT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(RoseParserRSHIFT_ASSIGN, 0)
}

func (s *ExpressionContext) URSHIFT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(RoseParserURSHIFT_ASSIGN, 0)
}

func (s *ExpressionContext) LSHIFT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(RoseParserLSHIFT_ASSIGN, 0)
}

func (s *ExpressionContext) MOD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(RoseParserMOD_ASSIGN, 0)
}

func (s *ExpressionContext) DOT() antlr.TerminalNode {
	return s.GetToken(RoseParserDOT, 0)
}

func (s *ExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RoseParserIDENTIFIER, 0)
}

func (s *ExpressionContext) THIS() antlr.TerminalNode {
	return s.GetToken(RoseParserTHIS, 0)
}

func (s *ExpressionContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(RoseParserLBRACK, 0)
}

func (s *ExpressionContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(RoseParserRBRACK, 0)
}

func (s *ExpressionContext) TypeType() ITypeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *ExpressionContext) INSTANCEOF() antlr.TerminalNode {
	return s.GetToken(RoseParserINSTANCEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *RoseParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 78
	p.EnterRecursionRule(localctx, 78, RoseParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(452)
			p.Primary()
		}

	case 2:
		{
			p.SetState(453)
			p.FunctionCall()
		}

	case 3:
		{
			p.SetState(454)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressionContext).prefix = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-85)) & ^0x3f) == 0 && ((int64(1)<<(_la-85))&15) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressionContext).prefix = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(455)
			p.expression(15)
		}

	case 4:
		{
			p.SetState(456)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressionContext).prefix = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == RoseParserBANG || _la == RoseParserTILDE) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressionContext).prefix = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(457)
			p.expression(14)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(526)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(524)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RoseParserRULE_expression)
				p.SetState(460)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(461)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-89)) & ^0x3f) == 0 && ((int64(1)<<(_la-89))&35) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(462)
					p.expression(14)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RoseParserRULE_expression)
				p.SetState(463)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(464)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == RoseParserADD || _la == RoseParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(465)
					p.expression(13)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RoseParserRULE_expression)
				p.SetState(466)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				p.SetState(474)
				p.GetErrorHandler().Sync(p)
				switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(467)
						p.Match(RoseParserLT)
					}
					{
						p.SetState(468)
						p.Match(RoseParserLT)
					}

				case 2:
					{
						p.SetState(469)
						p.Match(RoseParserGT)
					}
					{
						p.SetState(470)
						p.Match(RoseParserGT)
					}
					{
						p.SetState(471)
						p.Match(RoseParserGT)
					}

				case 3:
					{
						p.SetState(472)
						p.Match(RoseParserGT)
					}
					{
						p.SetState(473)
						p.Match(RoseParserGT)
					}

				}
				{
					p.SetState(476)
					p.expression(12)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RoseParserRULE_expression)
				p.SetState(477)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(478)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-73)) & ^0x3f) == 0 && ((int64(1)<<(_la-73))&387) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(479)
					p.expression(11)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RoseParserRULE_expression)
				p.SetState(480)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(481)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == RoseParserEQUAL || _la == RoseParserNOTEQUAL) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(482)
					p.expression(9)
				}

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RoseParserRULE_expression)
				p.SetState(483)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(484)

					var _m = p.Match(RoseParserBITAND)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(485)
					p.expression(8)
				}

			case 7:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RoseParserRULE_expression)
				p.SetState(486)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(487)

					var _m = p.Match(RoseParserCARET)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(488)
					p.expression(7)
				}

			case 8:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RoseParserRULE_expression)
				p.SetState(489)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(490)

					var _m = p.Match(RoseParserBITOR)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(491)
					p.expression(6)
				}

			case 9:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RoseParserRULE_expression)
				p.SetState(492)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(493)

					var _m = p.Match(RoseParserAND)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(494)
					p.expression(5)
				}

			case 10:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RoseParserRULE_expression)
				p.SetState(495)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(496)

					var _m = p.Match(RoseParserOR)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(497)
					p.expression(4)
				}

			case 11:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RoseParserRULE_expression)
				p.SetState(498)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(499)

					var _m = p.Match(RoseParserQUESTION)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(500)
					p.expression(0)
				}
				{
					p.SetState(501)
					p.Match(RoseParserCOLON)
				}
				{
					p.SetState(502)
					p.expression(3)
				}

			case 12:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RoseParserRULE_expression)
				p.SetState(504)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(505)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-72)) & ^0x3f) == 0 && ((int64(1)<<(_la-72))&17171480577) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(506)
					p.expression(1)
				}

			case 13:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RoseParserRULE_expression)
				p.SetState(507)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(508)

					var _m = p.Match(RoseParserDOT)

					localctx.(*ExpressionContext).bop = _m
				}
				p.SetState(512)
				p.GetErrorHandler().Sync(p)
				switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(509)
						p.Match(RoseParserIDENTIFIER)
					}

				case 2:
					{
						p.SetState(510)
						p.FunctionCall()
					}

				case 3:
					{
						p.SetState(511)
						p.Match(RoseParserTHIS)
					}

				}

			case 14:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RoseParserRULE_expression)
				p.SetState(514)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(515)
					p.Match(RoseParserLBRACK)
				}
				{
					p.SetState(516)
					p.expression(0)
				}
				{
					p.SetState(517)
					p.Match(RoseParserRBRACK)
				}

			case 15:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RoseParserRULE_expression)
				p.SetState(519)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(520)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).postfix = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == RoseParserINC || _la == RoseParserDEC) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).postfix = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			case 16:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RoseParserRULE_expression)
				p.SetState(521)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(522)

					var _m = p.Match(RoseParserINSTANCEOF)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(523)
					p.TypeType()
				}

			}

		}
		p.SetState(528)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RoseParserLPAREN, 0)
}

func (s *PrimaryContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RoseParserRPAREN, 0)
}

func (s *PrimaryContext) THIS() antlr.TerminalNode {
	return s.GetToken(RoseParserTHIS, 0)
}

func (s *PrimaryContext) SUPER() antlr.TerminalNode {
	return s.GetToken(RoseParserSUPER, 0)
}

func (s *PrimaryContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PrimaryContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RoseParserIDENTIFIER, 0)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) Primary() (localctx IPrimaryContext) {
	this := p
	_ = this

	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, RoseParserRULE_primary)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(537)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RoseParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(529)
			p.Match(RoseParserLPAREN)
		}
		{
			p.SetState(530)
			p.expression(0)
		}
		{
			p.SetState(531)
			p.Match(RoseParserRPAREN)
		}

	case RoseParserTHIS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(533)
			p.Match(RoseParserTHIS)
		}

	case RoseParserSUPER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(534)
			p.Match(RoseParserSUPER)
		}

	case RoseParserDECIMAL_LITERAL, RoseParserHEX_LITERAL, RoseParserOCT_LITERAL, RoseParserBINARY_LITERAL, RoseParserFLOAT_LITERAL, RoseParserHEX_FLOAT_LITERAL, RoseParserBOOL_LITERAL, RoseParserCHAR_LITERAL, RoseParserSTRING_LITERAL, RoseParserNULL_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(535)
			p.Literal()
		}

	case RoseParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(536)
			p.Match(RoseParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeListContext is an interface to support dynamic dispatch.
type ITypeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeListContext differentiates from other interfaces.
	IsTypeListContext()
}

type TypeListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeListContext() *TypeListContext {
	var p = new(TypeListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_typeList
	return p
}

func (*TypeListContext) IsTypeListContext() {}

func NewTypeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeListContext {
	var p = new(TypeListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_typeList

	return p
}

func (s *TypeListContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeListContext) AllTypeType() []ITypeTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeTypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeTypeContext); ok {
			tst[i] = t.(ITypeTypeContext)
			i++
		}
	}

	return tst
}

func (s *TypeListContext) TypeType(i int) ITypeTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeTypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *TypeListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(RoseParserCOMMA)
}

func (s *TypeListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(RoseParserCOMMA, i)
}

func (s *TypeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterTypeList(s)
	}
}

func (s *TypeListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitTypeList(s)
	}
}

func (s *TypeListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitTypeList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) TypeList() (localctx ITypeListContext) {
	this := p
	_ = this

	localctx = NewTypeListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, RoseParserRULE_typeList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(539)
		p.TypeType()
	}
	p.SetState(544)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RoseParserCOMMA {
		{
			p.SetState(540)
			p.Match(RoseParserCOMMA)
		}
		{
			p.SetState(541)
			p.TypeType()
		}

		p.SetState(546)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeTypeContext is an interface to support dynamic dispatch.
type ITypeTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeTypeContext differentiates from other interfaces.
	IsTypeTypeContext()
}

type TypeTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeTypeContext() *TypeTypeContext {
	var p = new(TypeTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_typeType
	return p
}

func (*TypeTypeContext) IsTypeTypeContext() {}

func NewTypeTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeTypeContext {
	var p = new(TypeTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_typeType

	return p
}

func (s *TypeTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeTypeContext) ClassOrInterfaceType() IClassOrInterfaceTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassOrInterfaceTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassOrInterfaceTypeContext)
}

func (s *TypeTypeContext) FunctionType() IFunctionTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionTypeContext)
}

func (s *TypeTypeContext) PrimitiveType() IPrimitiveTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitiveTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *TypeTypeContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(RoseParserLBRACK)
}

func (s *TypeTypeContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(RoseParserLBRACK, i)
}

func (s *TypeTypeContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(RoseParserRBRACK)
}

func (s *TypeTypeContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(RoseParserRBRACK, i)
}

func (s *TypeTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterTypeType(s)
	}
}

func (s *TypeTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitTypeType(s)
	}
}

func (s *TypeTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitTypeType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) TypeType() (localctx ITypeTypeContext) {
	this := p
	_ = this

	localctx = NewTypeTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, RoseParserRULE_typeType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(550)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RoseParserIDENTIFIER:
		{
			p.SetState(547)
			p.ClassOrInterfaceType()
		}

	case RoseParserFUNCTION:
		{
			p.SetState(548)
			p.FunctionType()
		}

	case RoseParserBOOLEAN, RoseParserBYTE, RoseParserCHAR, RoseParserDOUBLE, RoseParserFLOAT, RoseParserINT, RoseParserLONG, RoseParserSHORT, RoseParserSTRING:
		{
			p.SetState(549)
			p.PrimitiveType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(556)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(552)
				p.Match(RoseParserLBRACK)
			}
			{
				p.SetState(553)
				p.Match(RoseParserRBRACK)
			}

		}
		p.SetState(558)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext())
	}

	return localctx
}

// IFunctionTypeContext is an interface to support dynamic dispatch.
type IFunctionTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionTypeContext differentiates from other interfaces.
	IsFunctionTypeContext()
}

type FunctionTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionTypeContext() *FunctionTypeContext {
	var p = new(FunctionTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_functionType
	return p
}

func (*FunctionTypeContext) IsFunctionTypeContext() {}

func NewFunctionTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionTypeContext {
	var p = new(FunctionTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_functionType

	return p
}

func (s *FunctionTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionTypeContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(RoseParserFUNCTION, 0)
}

func (s *FunctionTypeContext) TypeTypeOrVoid() ITypeTypeOrVoidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeTypeOrVoidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeTypeOrVoidContext)
}

func (s *FunctionTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RoseParserLPAREN, 0)
}

func (s *FunctionTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RoseParserRPAREN, 0)
}

func (s *FunctionTypeContext) TypeList() ITypeListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeListContext)
}

func (s *FunctionTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterFunctionType(s)
	}
}

func (s *FunctionTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitFunctionType(s)
	}
}

func (s *FunctionTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitFunctionType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) FunctionType() (localctx IFunctionTypeContext) {
	this := p
	_ = this

	localctx = NewFunctionTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, RoseParserRULE_functionType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(559)
		p.Match(RoseParserFUNCTION)
	}
	{
		p.SetState(560)
		p.TypeTypeOrVoid()
	}
	{
		p.SetState(561)
		p.Match(RoseParserLPAREN)
	}
	p.SetState(563)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6755537552163112) != 0 || _la == RoseParserIDENTIFIER {
		{
			p.SetState(562)
			p.TypeList()
		}

	}
	{
		p.SetState(565)
		p.Match(RoseParserRPAREN)
	}

	return localctx
}

// IPrimitiveTypeContext is an interface to support dynamic dispatch.
type IPrimitiveTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitiveTypeContext differentiates from other interfaces.
	IsPrimitiveTypeContext()
}

type PrimitiveTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveTypeContext() *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_primitiveType
	return p
}

func (*PrimitiveTypeContext) IsPrimitiveTypeContext() {}

func NewPrimitiveTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_primitiveType

	return p
}

func (s *PrimitiveTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveTypeContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(RoseParserBOOLEAN, 0)
}

func (s *PrimitiveTypeContext) CHAR() antlr.TerminalNode {
	return s.GetToken(RoseParserCHAR, 0)
}

func (s *PrimitiveTypeContext) BYTE() antlr.TerminalNode {
	return s.GetToken(RoseParserBYTE, 0)
}

func (s *PrimitiveTypeContext) SHORT() antlr.TerminalNode {
	return s.GetToken(RoseParserSHORT, 0)
}

func (s *PrimitiveTypeContext) INT() antlr.TerminalNode {
	return s.GetToken(RoseParserINT, 0)
}

func (s *PrimitiveTypeContext) LONG() antlr.TerminalNode {
	return s.GetToken(RoseParserLONG, 0)
}

func (s *PrimitiveTypeContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(RoseParserFLOAT, 0)
}

func (s *PrimitiveTypeContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(RoseParserDOUBLE, 0)
}

func (s *PrimitiveTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(RoseParserSTRING, 0)
}

func (s *PrimitiveTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterPrimitiveType(s)
	}
}

func (s *PrimitiveTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitPrimitiveType(s)
	}
}

func (s *PrimitiveTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitPrimitiveType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) PrimitiveType() (localctx IPrimitiveTypeContext) {
	this := p
	_ = this

	localctx = NewPrimitiveTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, RoseParserRULE_primitiveType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(567)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4503737738477864) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICreatorContext is an interface to support dynamic dispatch.
type ICreatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreatorContext differentiates from other interfaces.
	IsCreatorContext()
}

type CreatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreatorContext() *CreatorContext {
	var p = new(CreatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_creator
	return p
}

func (*CreatorContext) IsCreatorContext() {}

func NewCreatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreatorContext {
	var p = new(CreatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_creator

	return p
}

func (s *CreatorContext) GetParser() antlr.Parser { return s.parser }

func (s *CreatorContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RoseParserIDENTIFIER, 0)
}

func (s *CreatorContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *CreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterCreator(s)
	}
}

func (s *CreatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitCreator(s)
	}
}

func (s *CreatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitCreator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) Creator() (localctx ICreatorContext) {
	this := p
	_ = this

	localctx = NewCreatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, RoseParserRULE_creator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(569)
		p.Match(RoseParserIDENTIFIER)
	}
	{
		p.SetState(570)
		p.Arguments()
	}

	return localctx
}

// ISuperSuffixContext is an interface to support dynamic dispatch.
type ISuperSuffixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSuperSuffixContext differentiates from other interfaces.
	IsSuperSuffixContext()
}

type SuperSuffixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySuperSuffixContext() *SuperSuffixContext {
	var p = new(SuperSuffixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_superSuffix
	return p
}

func (*SuperSuffixContext) IsSuperSuffixContext() {}

func NewSuperSuffixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SuperSuffixContext {
	var p = new(SuperSuffixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_superSuffix

	return p
}

func (s *SuperSuffixContext) GetParser() antlr.Parser { return s.parser }

func (s *SuperSuffixContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *SuperSuffixContext) DOT() antlr.TerminalNode {
	return s.GetToken(RoseParserDOT, 0)
}

func (s *SuperSuffixContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RoseParserIDENTIFIER, 0)
}

func (s *SuperSuffixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SuperSuffixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SuperSuffixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterSuperSuffix(s)
	}
}

func (s *SuperSuffixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitSuperSuffix(s)
	}
}

func (s *SuperSuffixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitSuperSuffix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) SuperSuffix() (localctx ISuperSuffixContext) {
	this := p
	_ = this

	localctx = NewSuperSuffixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, RoseParserRULE_superSuffix)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(578)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RoseParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(572)
			p.Arguments()
		}

	case RoseParserDOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(573)
			p.Match(RoseParserDOT)
		}
		{
			p.SetState(574)
			p.Match(RoseParserIDENTIFIER)
		}
		p.SetState(576)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RoseParserLPAREN {
			{
				p.SetState(575)
				p.Arguments()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoseParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoseParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(RoseParserLPAREN, 0)
}

func (s *ArgumentsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(RoseParserRPAREN, 0)
}

func (s *ArgumentsContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoseListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RoseVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RoseParser) Arguments() (localctx IArgumentsContext) {
	this := p
	_ = this

	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, RoseParserRULE_arguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(580)
		p.Match(RoseParserLPAREN)
	}
	p.SetState(582)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8997303650091008) != 0 || (int64((_la-75)) & ^0x3f) == 0 && ((int64(1)<<(_la-75))&274877922307) != 0 {
		{
			p.SetState(581)
			p.ExpressionList()
		}

	}
	{
		p.SetState(584)
		p.Match(RoseParserRPAREN)
	}

	return localctx
}

func (p *RoseParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 39:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RoseParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
