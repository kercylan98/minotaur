package parser

type Symbol = byte

const (
	SymbolLeftBracket  Symbol = '['  // 左方括号
	SymbolRightBracket Symbol = ']'  // 右方括号
	SymbolLeftBrace    Symbol = '{'  // 左大括号
	SymbolRightBrace   Symbol = '}'  // 右大括号
	SymbolColon        Symbol = ':'  // 冒号
	SymbolComma        Symbol = ','  // 逗号
	SymbolEqual        Symbol = '='  // 等号
	SymbolQuote        Symbol = '\'' // 单引号
	SymbolDoubleQuote  Symbol = '"'  // 双引号
	SymbolBackslash    Symbol = '\\' // 反斜杠
	SymbolSlash        Symbol = '/'  // 斜杠
	SymbolAsterisk     Symbol = '*'  // 星号
	SymbolDot          Symbol = '.'  // 点
	SymbolHash         Symbol = '#'  // 井
	SymbolAt           Symbol = '@'  // @
	SymbolDollar       Symbol = '$'  // $
	SymbolPercent      Symbol = '%'  // 百分号
	SymbolCaret        Symbol = '^'  // 尖
	SymbolAmpersand    Symbol = '&'  // 与号
	SymbolUnderscore   Symbol = '_'  // 下划线
	SymbolTilde        Symbol = '~'  // 波浪号
	SymbolExclamation  Symbol = '!'  // 感叹号
	SymbolQuestionMark Symbol = '?'  // 问号
	SymbolLeftParen    Symbol = '('  // 左括号
	SymbolRightParen   Symbol = ')'  // 右括号
	SymbolLessThan     Symbol = '<'  // 小于号
	SymbolGreaterThan  Symbol = '>'  // 大于号
	SymbolPipe         Symbol = '|'  // 管道
)
