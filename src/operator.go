package src

//go:generate stringer -type=OperatorSymbol
type OperatorSymbol uint8

// Node ops.
const (
	OXXX OperatorSymbol = iota
	VALUE
	LITERAL
	PLUS
	MINUS
	EXCLAMATION
	EN_NOT
	CARET
	STAR
	AMPERSAND
	DIV
	MOD
	QUESTION
	LSHIFT
	RSHIFT
	BIT_CLEAR
	OR
	EQUALS
	NOT_EQUALS
	LESS
	LESS_OR_EQUALS
	GREATER
	GREATER_OR_EQUALS
	LOGICAL_AND
	LOGICAL_OR
	LOGICAL_IN
	OCALLFUNC
	BOOL
	FLOATE64
	STRING
	SLICE
	INDEX
	SLICEPARAMS
	FUNCCALL
	EXPRESSIONLIST
	EXPRESSIONCOLON
	NIL
)