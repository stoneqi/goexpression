// Package main provides ...
package parserSecond

type EvaluationOperator func(left any, right any, parameters Parameters) (any, error)

func (eo *EvaluationOperator) MarshalJSON() ([]byte, error) {
	return []byte("\"EvaluationOperator\""), nil
}

type StageTypeCheck func(value any) bool

func (stc *StageTypeCheck) MarshalJSON() ([]byte, error) {
	return []byte("\"StageTypeCheck\""), nil
}

type StageCombinedTypeCheck func(left any, right any) bool

func (sctc *StageCombinedTypeCheck) MarshalJSON() ([]byte, error) {
	return []byte("\"StageCombinedTypeCheck\""), nil
}

type evaluationNode struct {
	Symbol OperatorSymbol

	LeftOperator  *evaluationNode
	RightOperator []*evaluationNode

	// the operation that will be used to evaluate this stage (such as adding [left] to [right] and return the result)
	Operator EvaluationOperator

	// ensures that both left and right values are appropriate for this stage. Returns an error if they aren't operable.
	LeftTypeCheck  StageTypeCheck
	RightTypeCheck StageTypeCheck

	// if specified, will override whatever is used in "LeftTypeCheck" and "RightTypeCheck".
	// primarily used for specific operators that don't care which side a given type is on, but still requires one side to be of a given type
	// (like string concat)
	TypeCheck StageCombinedTypeCheck

	RawString string

	// regardless of which type check is used, this string format will be used as the error message for type errors
	TypeErrorFormat string
}

func newEvaluationNode() *evaluationNode {
	return &evaluationNode{
		Symbol:          0,
		LeftOperator:    nil,
		RightOperator:   nil,
		Operator:        nil,
		LeftTypeCheck:   nil,
		RightTypeCheck:  nil,
		TypeCheck:       nil,
		TypeErrorFormat: "",
	}
}

func (eval *evaluationNode) swapWith(other *evaluationNode) {

	temp := *other
	other.setToNonOperator(*eval)
	eval.setToNonOperator(temp)
}

func (eval *evaluationNode) setToNonOperator(other evaluationNode) {

	eval.Symbol = other.Symbol
	eval.Operator = other.Operator
	eval.LeftTypeCheck = other.LeftTypeCheck
	eval.RightTypeCheck = other.RightTypeCheck
	eval.TypeCheck = other.TypeCheck
	eval.TypeErrorFormat = other.TypeErrorFormat
}

func (eval *evaluationNode) isShortCircuitable() bool {

	switch eval.Symbol {
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
