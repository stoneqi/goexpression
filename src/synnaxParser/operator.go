package parserSecond

type OperatorSymbol uint8

// Node ops.
const (
	OXXX OperatorSymbol = iota
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
	FUNCPARAMS
	NIL
)

func (this OperatorSymbol) String() string {

	switch this {
	case OR:
		return "||"
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	}
	return ""
}
