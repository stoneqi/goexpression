// Package main provides ...
package parserSecond

type evaluationOperator func(left interface{}, right interface{}, parameters Parameters) (interface{}, error)
type stageTypeCheck func(value interface{}) bool
type stageCombinedTypeCheck func(left interface{}, right interface{}) bool

type evaluationNode struct {
	symbol OperatorSymbol

	leftOperator, rightOperator *evaluationNode

	// the operation that will be used to evaluate this stage (such as adding [left] to [right] and return the result)
	operator evaluationOperator

	// ensures that both left and right values are appropriate for this stage. Returns an error if they aren't operable.
	leftTypeCheck  stageTypeCheck
	rightTypeCheck stageTypeCheck

	// if specified, will override whatever is used in "leftTypeCheck" and "rightTypeCheck".
	// primarily used for specific operators that don't care which side a given type is on, but still requires one side to be of a given type
	// (like string concat)
	typeCheck stageCombinedTypeCheck

	// regardless of which type check is used, this string format will be used as the error message for type errors
	typeErrorFormat string
}

func (eval *evaluationNode) swapWith(other *evaluationNode) {

	temp := *other
	other.setToNonOperator(*eval)
	eval.setToNonOperator(temp)
}

func (eval *evaluationNode) setToNonOperator(other evaluationNode) {

	eval.symbol = other.symbol
	eval.operator = other.operator
	eval.leftTypeCheck = other.leftTypeCheck
	eval.rightTypeCheck = other.rightTypeCheck
	eval.typeCheck = other.typeCheck
	eval.typeErrorFormat = other.typeErrorFormat
}

func (eval *evaluationNode) isShortCircuitable() bool {

	switch eval.symbol {
	case LOGICAL_AND:
		fallthrough
	case LOGICAL_OR:
		//fallthrough
		//case TERNARY_TRUE:
		//	fallthrough
		//case TERNARY_FALSE:
		//	fallthrough
		//case COALESCE:
		return true
	}

	return false
}
