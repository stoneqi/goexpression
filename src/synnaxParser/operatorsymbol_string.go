// Code generated by "stringer -type=OperatorSymbol"; DO NOT EDIT.

package parserSecond

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OXXX-0]
	_ = x[LITERAL-1]
	_ = x[PLUS-2]
	_ = x[MINUS-3]
	_ = x[EXCLAMATION-4]
	_ = x[EN_NOT-5]
	_ = x[CARET-6]
	_ = x[STAR-7]
	_ = x[AMPERSAND-8]
	_ = x[DIV-9]
	_ = x[MOD-10]
	_ = x[LSHIFT-11]
	_ = x[RSHIFT-12]
	_ = x[BIT_CLEAR-13]
	_ = x[OR-14]
	_ = x[EQUALS-15]
	_ = x[NOT_EQUALS-16]
	_ = x[LESS-17]
	_ = x[LESS_OR_EQUALS-18]
	_ = x[GREATER-19]
	_ = x[GREATER_OR_EQUALS-20]
	_ = x[LOGICAL_AND-21]
	_ = x[LOGICAL_OR-22]
	_ = x[LOGICAL_IN-23]
	_ = x[OCALLFUNC-24]
	_ = x[BOOL-25]
	_ = x[FLOATE64-26]
	_ = x[STRING-27]
	_ = x[SLICE-28]
	_ = x[INDEX-29]
	_ = x[SLICEPARAMS-30]
	_ = x[FUNCCALL-31]
	_ = x[EXPRESSIONLIST-32]
	_ = x[EXPRESSIONCOLON-33]
	_ = x[NIL-34]
}

const _OperatorSymbol_name = "OXXXLITERALPLUSMINUSEXCLAMATIONEN_NOTCARETSTARAMPERSANDDIVMODLSHIFTRSHIFTBIT_CLEAROREQUALSNOT_EQUALSLESSLESS_OR_EQUALSGREATERGREATER_OR_EQUALSLOGICAL_ANDLOGICAL_ORLOGICAL_INOCALLFUNCBOOLFLOATE64STRINGSLICEINDEXSLICEPARAMSFUNCCALLEXPRESSIONLISTEXPRESSIONCOLONNIL"

var _OperatorSymbol_index = [...]uint16{0, 4, 11, 15, 20, 31, 37, 42, 46, 55, 58, 61, 67, 73, 82, 84, 90, 100, 104, 118, 125, 142, 153, 163, 173, 182, 186, 194, 200, 205, 210, 221, 229, 243, 258, 261}

func (i OperatorSymbol) String() string {
	if i >= OperatorSymbol(len(_OperatorSymbol_index)-1) {
		return "OperatorSymbol(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OperatorSymbol_name[_OperatorSymbol_index[i]:_OperatorSymbol_index[i+1]]
}
