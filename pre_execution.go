package goexpression

/*
Recurses through all operators in the entire tree, eliding operators where both sides are literals.
*/
func elideLiterals(root *evaluationNode) *evaluationNode {

	if root.LeftOperator != nil {
		root.LeftOperator = elideLiterals(root.LeftOperator)
	}

	for i := 0; i < len(root.RightOperator); i++ {
		root.RightOperator[i] = elideLiterals(root.RightOperator[i])
	}

	return elideStage(root)
}

/*
Elides a specific singleExpr, if possible.
Returns the unmodified [root] singleExpr if it cannot or should not be elided.
Otherwise, returns a new singleExpr representing the condensed value from the elided stages.
*/
func elideStage(root *evaluationNode) *evaluationNode {

	var leftValue, rightValue, result interface{}
	var err error

	if root.LeftOperator == nil && root.RightOperator == nil {
		return root
	}
	// don't elide some operators
	switch root.Symbol {
	//case SEPARATE:
	//	fallthrough
	//case INDEX:
	//	return root
	}

	if root.LeftOperator != nil {
		if root.LeftOperator.Symbol != BOOL &&
			root.LeftOperator.Symbol != FLOATE64 &&
			root.LeftOperator.Symbol != NIL &&
			root.LeftOperator.Symbol != STRING &&
			root.LeftOperator.Symbol != LITERAL {
			return root
		}
		leftValue, err = root.LeftOperator.Operator(nil, nil, nil)
		if err != nil {
			return root
		}
	}

	if root.RightOperator != nil {
		var rightValueTemp []interface{}
		for _, node := range root.RightOperator {
			if err == nil && node != nil {
				if node.Symbol != BOOL && node.Symbol != FLOATE64 && node.Symbol != NIL && node.Symbol != STRING &&
					node.Symbol != LITERAL {
					return root
				}
				operatorRet, err := node.Operator(nil, nil, nil)
				if err != nil {
					return root
				} else {
					rightValueTemp = append(rightValueTemp, operatorRet)
				}
			} else {
				return root
			}
		}
		rightValue = rightValueTemp
		if len(rightValueTemp) == 1 {
			rightValue = rightValueTemp[0]
		}
	}
	// typcheck, since the grammar checker is a bit loose with which operator symbols go together.
	err = typeCheck(root.LeftTypeCheck, leftValue, root.Symbol, root.TypeErrorFormat)
	if err != nil {
		return root
	}

	err = typeCheck(root.RightTypeCheck, rightValue, root.Symbol, root.TypeErrorFormat)
	if err != nil {
		return root
	}

	if root.TypeCheck != nil && !root.TypeCheck(leftValue, rightValue) {
		return root
	}

	// pre-calculate, and return a new singleExpr representing the result.
	result, err = root.Operator(leftValue, rightValue, nil)
	if err != nil {
		return root
	}

	return &evaluationNode{
		Symbol:    LITERAL,
		Operator:  makeLiteralOperator(result),
		RawString: root.RawString,
	}
}
