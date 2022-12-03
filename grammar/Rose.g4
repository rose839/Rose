// 该文件定义rose语言的语法文法
grammar Rose;

// 导入词法文法文件
import Lexer;

// 类定义文法
classDeclaration
    : CLASS IDENTIFIER
      (EXTENDS typeType)?
      (IMPLEMENTS typeList)?
      classBody
    ;

// 类体文法
classBody
    : '{' classBodyDeclaration* '}'
    ;

// 接口体定义
// interfaceBody
//     : '{' interfaceBodyDeclaration* '}'
//     ;

// 类成员文法
classBodyDeclaration
    : ';'
    //| STATIC? block
    | memberDeclaration
    ;

// 成员文法
memberDeclaration
    : functionDeclaration
//    | genericFunctionDeclaration
    | fieldDeclaration
    // | constructorDeclaration
    // | genericConstructorDeclaration
//     | interfaceDeclaration
    // | annotationTypeDeclaration
     | classDeclaration
    // | enumDeclaration
    ;

// 函数和方法文法
functionDeclaration
    : typeTypeOrVoid? IDENTIFIER formalParameters ('[' ']')*
      (THROWS qualifiedNameList)?
      functionBody
    ;

// 函数和方法体文法
functionBody
    : block
    | ';'
    ;

// 函数和方法返回值文法
typeTypeOrVoid
    : typeType
    | VOID
    ;

// '.'分割的限定名文法
qualifiedNameList
    : qualifiedName (',' qualifiedName)*
    ;

// 形参文法
formalParameters
    : '(' formalParameterList? ')'
    ;

// 形参列表文法
formalParameterList
    : formalParameter (',' formalParameter)* (',' lastFormalParameter)?
    | lastFormalParameter
    ;

// 单个形参文法
formalParameter
    : variableModifier* typeType variableDeclaratorId
    ;

// 可变参文法
lastFormalParameter
    : variableModifier* typeType '...' variableDeclaratorId
    ;

// 变量修饰符文法
variableModifier
    : FINAL
    //| annotation
    ;

// 单个'.'分割的限定名文法
qualifiedName
    : IDENTIFIER ('.' IDENTIFIER)*
    ;

// 属性定义文法
fieldDeclaration
    //: typeType variableDeclarators ';'
    : variableDeclarators ';'
    ;

// 构造函数定义文法
constructorDeclaration
    : IDENTIFIER formalParameters (THROWS qualifiedNameList)? constructorBody=block
    ;

// 变量定义文法
variableDeclarators
    : typeType variableDeclarator (',' variableDeclarator)*
    ;

// 单个变量定义文法
variableDeclarator
    : variableDeclaratorId ('=' variableInitializer)?
    ;

// 变量文法
variableDeclaratorId
    : IDENTIFIER ('[' ']')*
    ;

// 变量初始化的值部分的文法
variableInitializer
    : arrayInitializer
    | expression
    ;

// 变量初始化的值为数组的值部分的文法
arrayInitializer
    : '{' (variableInitializer (',' variableInitializer)* (',')? )? '}'
    ;

// 类或者接口名文法
classOrInterfaceType
    : IDENTIFIER ('.' IDENTIFIER)*
    //: IDENTIFIER
    ;

// 类型参数文法
typeArgument
    : typeType
    | '?' ((EXTENDS | SUPER) typeType)?
    ;

// 字面量文法
literal
    : integerLiteral
    | floatLiteral
    | CHAR_LITERAL
    | STRING_LITERAL
    | BOOL_LITERAL
    | NULL_LITERAL
    ;

// 整数字面量文法
integerLiteral
    : DECIMAL_LITERAL
    | HEX_LITERAL
    | OCT_LITERAL
    | BINARY_LITERAL
    ;

// 浮点数字面量文法
floatLiteral
    : FLOAT_LITERAL
    | HEX_FLOAT_LITERAL
    ;

// 代表整个程序的文法
prog
    : blockStatements
    ;

// 代码块文法
block
    : '{' blockStatements '}'
    ;

// 语句文法
blockStatements
    : blockStatement*
    ;

// 单个语句文法
blockStatement
    : variableDeclarators ';'
    | statement
   // | localTypeDeclaration
    | functionDeclaration
    | classDeclaration
    ;

// 普通语句文法
statement
    : blockLabel=block
    // | ASSERT expression (':' expression)? ';'
    | IF parExpression statement (ELSE statement)?
    | FOR '(' forControl ')' statement
    | WHILE parExpression statement
    | DO statement WHILE parExpression ';'
    //| TRY block (catchClause+ finallyBlock? | finallyBlock)
    //| TRY resourceSpecification block catchClause* finallyBlock?
    | SWITCH parExpression '{' switchBlockStatementGroup* switchLabel* '}'
    //| SYNCHRONIZED parExpression block
    | RETURN expression? ';'
    //| THROW expression ';'
    | BREAK IDENTIFIER? ';'
    | CONTINUE IDENTIFIER? ';'
    | SEMI
    | statementExpression=expression ';'
    | identifierLabel=IDENTIFIER ':' statement
    ;

// switch语句的case部分的文法
switchBlockStatementGroup
    : switchLabel+ blockStatement+
    ;

// switch语句的case 标签部分的文法
switchLabel
    : CASE (constantExpression=expression | enumConstantName=IDENTIFIER) ':'
    | DEFAULT ':'
    ;

// for循环整个条件部分的文法
forControl
    : enhancedForControl
    | forInit? ';' expression? ';' forUpdate=expressionList?
    ;

// for循环整个条件初始化部分的文法
forInit
    : variableDeclarators
    | expressionList
    ;

// for(xxx : xxx)简化形式的for循环条件部分的文法
enhancedForControl
    : typeType variableDeclaratorId ':' expression
    ;

// 部分条件语句，比如if、while等条件部分的文法
parExpression
    : '(' expression ')'
    ;

// 多个表达式的文法
expressionList
    : expression (',' expression)*
    ;

// 函数调用文法
functionCall
    : IDENTIFIER '(' expressionList? ')'
    | THIS '(' expressionList? ')'
    | SUPER '(' expressionList? ')'
    ;

// 表达式文法
expression
    : primary
    | expression bop='.'
      ( IDENTIFIER
      | functionCall
      | THIS
    //   | NEW nonWildcardTypeArguments? innerCreator
    //   | SUPER superSuffix
    //   | explicitGenericInvocation
      )
    | expression '[' expression ']'
    | functionCall
    // | NEW creator   //不用new关键字，而是用类名相同的函数直接生成对象，类似于python的文法
    // | '(' typeType ')' expression
    | expression postfix=('++' | '--')
    | prefix=('+'|'-'|'++'|'--') expression
    | prefix=('~'|'!') expression
    | expression bop=('*'|'/'|'%') expression  
    | expression bop=('+'|'-') expression 
    | expression ('<' '<' | '>' '>' '>' | '>' '>') expression
    | expression bop=('<=' | '>=' | '>' | '<') expression
    | expression bop=INSTANCEOF typeType
    | expression bop=('==' | '!=') expression
    | expression bop='&' expression
    | expression bop='^' expression
    | expression bop='|' expression
    | expression bop='&&' expression
    | expression bop='||' expression
    | expression bop='?' expression ':' expression
    | <assoc=right> expression
      bop=('=' | '+=' | '-=' | '*=' | '/=' | '&=' | '|=' | '^=' | '>>=' | '>>>=' | '<<=' | '%=')
      expression
    // | lambdaExpression

    // functionReference
    // | expression '::' typeArguments? IDENTIFIER
    // | typeType '::' (typeArguments? IDENTIFIER | NEW)
    // | classType '::' typeArguments? NEW
    ;

// 基础表达式文法
primary
    : '(' expression ')'
    | THIS
    | SUPER
    | literal
    | IDENTIFIER
    // | typeTypeOrVoid '.' CLASS
    ;

// 类型列表文法
typeList
    : typeType (',' typeType)*
    ;

// 单个类型文法
typeType
    : (classOrInterfaceType| functionType | primitiveType) ('[' ']')*
    ;

// 函数类型文法
functionType
    : FUNCTION typeTypeOrVoid '(' typeList? ')'
    ;

// 基础类型文法
primitiveType
    : BOOLEAN
    | CHAR
    | BYTE
    | SHORT
    | INT
    | LONG
    | FLOAT
    | DOUBLE
    | STRING
    ;

creator
    : IDENTIFIER arguments
    ;

superSuffix
    : arguments
    | '.' IDENTIFIER arguments?
    ;

// 实参文法
arguments
    : '(' expressionList? ')'
    ;
