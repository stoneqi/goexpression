// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // goexpression
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

type goexpressionParser struct {
	*antlr.BaseParser
}

var goexpressionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func goexpressionParserInit() {
	staticData := &goexpressionParserStaticData
	staticData.literalNames = []string{
		"", "", "", "'and'", "'in'", "'not'", "'or'", "'('", "')'", "'{'", "'}'",
		"'['", "']'", "'='", "','", "';'", "'?'", "':'", "'.'", "'nil'", "'||'",
		"'&&'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'|'", "'/'",
		"'%'", "'<<'", "'>>'", "'&^'", "'!'", "'+'", "'-'", "'^'", "'*'", "'&'",
	}
	staticData.symbolicNames = []string{
		"", "EN_TRUE", "EN_FALSE", "EN_AND", "EN_IN", "EN_NOT", "EN_OR", "L_PAREN",
		"R_PAREN", "L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET", "ASSIGN",
		"COMMA", "SEMI", "QUESTION", "COLON", "DOT", "NIL_LIT", "LOGICAL_OR",
		"LOGICAL_AND", "EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS", "GREATER",
		"GREATER_OR_EQUALS", "OR", "DIV", "MOD", "LSHIFT", "RSHIFT", "BIT_CLEAR",
		"EXCLAMATION", "PLUS", "MINUS", "CARET", "STAR", "AMPERSAND", "DECIMAL_LIT",
		"BINARY_LIT", "OCTAL_LIT", "HEX_LIT", "FLOAT_LIT", "DECIMAL_FLOAT_LIT",
		"HEX_FLOAT_LIT", "IDENTIFIER", "RAW_STRING_LIT", "INTERPRETED_STRING_LIT",
		"SINGLE_STRING_LIT", "WS", "COMMENT", "TERMINATOR", "LINE_COMMENT",
		"WS_NLSEMI", "COMMENT_NLSEMI", "LINE_COMMENT_NLSEMI", "EOS",
	}
	staticData.ruleNames = []string{
		"expressionStmt", "expression", "primaryExpr", "operand", "operandName",
		"slice_", "expressionColon", "index", "basicLit", "nil_lit", "en_bool",
		"float_lit", "integer", "expressionList", "arguments", "eos", "identifier",
		"string_",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 155, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 43,
		8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 72, 8, 1, 10, 1, 12, 1, 75, 9, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 86, 8, 2, 5, 2, 88,
		8, 2, 10, 2, 12, 2, 91, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 3, 3, 101, 8, 3, 1, 3, 3, 3, 104, 8, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 3, 8, 125, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 5, 13, 138, 8, 13, 10, 13, 12, 13, 141,
		9, 13, 1, 14, 1, 14, 3, 14, 145, 8, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 17, 0, 2, 2, 4, 18, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 0, 8, 2, 0, 5, 5, 34, 37, 2,
		0, 29, 33, 38, 39, 2, 0, 28, 28, 35, 37, 1, 0, 22, 27, 1, 0, 1, 2, 1, 0,
		40, 43, 2, 1, 15, 15, 58, 58, 1, 0, 49, 50, 160, 0, 36, 1, 0, 0, 0, 2,
		42, 1, 0, 0, 0, 4, 76, 1, 0, 0, 0, 6, 103, 1, 0, 0, 0, 8, 105, 1, 0, 0,
		0, 10, 107, 1, 0, 0, 0, 12, 111, 1, 0, 0, 0, 14, 115, 1, 0, 0, 0, 16, 124,
		1, 0, 0, 0, 18, 126, 1, 0, 0, 0, 20, 128, 1, 0, 0, 0, 22, 130, 1, 0, 0,
		0, 24, 132, 1, 0, 0, 0, 26, 134, 1, 0, 0, 0, 28, 142, 1, 0, 0, 0, 30, 148,
		1, 0, 0, 0, 32, 150, 1, 0, 0, 0, 34, 152, 1, 0, 0, 0, 36, 37, 3, 2, 1,
		0, 37, 1, 1, 0, 0, 0, 38, 39, 6, 1, -1, 0, 39, 43, 3, 4, 2, 0, 40, 41,
		7, 0, 0, 0, 41, 43, 3, 2, 1, 10, 42, 38, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0,
		43, 73, 1, 0, 0, 0, 44, 45, 10, 9, 0, 0, 45, 46, 7, 1, 0, 0, 46, 72, 3,
		2, 1, 10, 47, 48, 10, 8, 0, 0, 48, 49, 7, 2, 0, 0, 49, 72, 3, 2, 1, 9,
		50, 51, 10, 7, 0, 0, 51, 52, 7, 3, 0, 0, 52, 72, 3, 2, 1, 8, 53, 54, 10,
		6, 0, 0, 54, 55, 5, 21, 0, 0, 55, 72, 3, 2, 1, 7, 56, 57, 10, 5, 0, 0,
		57, 58, 5, 20, 0, 0, 58, 72, 3, 2, 1, 6, 59, 60, 10, 4, 0, 0, 60, 61, 5,
		3, 0, 0, 61, 72, 3, 2, 1, 5, 62, 63, 10, 3, 0, 0, 63, 64, 5, 6, 0, 0, 64,
		72, 3, 2, 1, 4, 65, 66, 10, 2, 0, 0, 66, 67, 5, 4, 0, 0, 67, 72, 3, 2,
		1, 3, 68, 69, 10, 1, 0, 0, 69, 70, 5, 16, 0, 0, 70, 72, 3, 12, 6, 0, 71,
		44, 1, 0, 0, 0, 71, 47, 1, 0, 0, 0, 71, 50, 1, 0, 0, 0, 71, 53, 1, 0, 0,
		0, 71, 56, 1, 0, 0, 0, 71, 59, 1, 0, 0, 0, 71, 62, 1, 0, 0, 0, 71, 65,
		1, 0, 0, 0, 71, 68, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0,
		73, 74, 1, 0, 0, 0, 74, 3, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 77, 6, 2,
		-1, 0, 77, 78, 3, 6, 3, 0, 78, 89, 1, 0, 0, 0, 79, 85, 10, 1, 0, 0, 80,
		81, 5, 18, 0, 0, 81, 86, 5, 47, 0, 0, 82, 86, 3, 14, 7, 0, 83, 86, 3, 10,
		5, 0, 84, 86, 3, 28, 14, 0, 85, 80, 1, 0, 0, 0, 85, 82, 1, 0, 0, 0, 85,
		83, 1, 0, 0, 0, 85, 84, 1, 0, 0, 0, 86, 88, 1, 0, 0, 0, 87, 79, 1, 0, 0,
		0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 5, 1,
		0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 104, 3, 16, 8, 0, 93, 104, 3, 8, 4, 0,
		94, 95, 5, 7, 0, 0, 95, 96, 3, 2, 1, 0, 96, 97, 5, 8, 0, 0, 97, 104, 1,
		0, 0, 0, 98, 100, 5, 9, 0, 0, 99, 101, 3, 26, 13, 0, 100, 99, 1, 0, 0,
		0, 100, 101, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 104, 5, 10, 0, 0, 103,
		92, 1, 0, 0, 0, 103, 93, 1, 0, 0, 0, 103, 94, 1, 0, 0, 0, 103, 98, 1, 0,
		0, 0, 104, 7, 1, 0, 0, 0, 105, 106, 3, 32, 16, 0, 106, 9, 1, 0, 0, 0, 107,
		108, 5, 11, 0, 0, 108, 109, 3, 12, 6, 0, 109, 110, 5, 12, 0, 0, 110, 11,
		1, 0, 0, 0, 111, 112, 3, 2, 1, 0, 112, 113, 5, 17, 0, 0, 113, 114, 3, 2,
		1, 0, 114, 13, 1, 0, 0, 0, 115, 116, 5, 11, 0, 0, 116, 117, 3, 2, 1, 0,
		117, 118, 5, 12, 0, 0, 118, 15, 1, 0, 0, 0, 119, 125, 3, 18, 9, 0, 120,
		125, 3, 20, 10, 0, 121, 125, 3, 24, 12, 0, 122, 125, 3, 34, 17, 0, 123,
		125, 3, 22, 11, 0, 124, 119, 1, 0, 0, 0, 124, 120, 1, 0, 0, 0, 124, 121,
		1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 123, 1, 0, 0, 0, 125, 17, 1, 0,
		0, 0, 126, 127, 5, 19, 0, 0, 127, 19, 1, 0, 0, 0, 128, 129, 7, 4, 0, 0,
		129, 21, 1, 0, 0, 0, 130, 131, 5, 44, 0, 0, 131, 23, 1, 0, 0, 0, 132, 133,
		7, 5, 0, 0, 133, 25, 1, 0, 0, 0, 134, 139, 3, 2, 1, 0, 135, 136, 5, 14,
		0, 0, 136, 138, 3, 2, 1, 0, 137, 135, 1, 0, 0, 0, 138, 141, 1, 0, 0, 0,
		139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 27, 1, 0, 0, 0, 141, 139,
		1, 0, 0, 0, 142, 144, 5, 7, 0, 0, 143, 145, 3, 26, 13, 0, 144, 143, 1,
		0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147, 5, 8, 0,
		0, 147, 29, 1, 0, 0, 0, 148, 149, 7, 6, 0, 0, 149, 31, 1, 0, 0, 0, 150,
		151, 5, 47, 0, 0, 151, 33, 1, 0, 0, 0, 152, 153, 7, 7, 0, 0, 153, 35, 1,
		0, 0, 0, 10, 42, 71, 73, 85, 89, 100, 103, 124, 139, 144,
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

// goexpressionParserInit initializes any static state used to implement goexpressionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewgoexpressionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoexpressionParserInit() {
	staticData := &goexpressionParserStaticData
	staticData.once.Do(goexpressionParserInit)
}

// NewgoexpressionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewgoexpressionParser(input antlr.TokenStream) *goexpressionParser {
	GoexpressionParserInit()
	this := new(goexpressionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &goexpressionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// goexpressionParser tokens.
const (
	goexpressionParserEOF                    = antlr.TokenEOF
	goexpressionParserEN_TRUE                = 1
	goexpressionParserEN_FALSE               = 2
	goexpressionParserEN_AND                 = 3
	goexpressionParserEN_IN                  = 4
	goexpressionParserEN_NOT                 = 5
	goexpressionParserEN_OR                  = 6
	goexpressionParserL_PAREN                = 7
	goexpressionParserR_PAREN                = 8
	goexpressionParserL_CURLY                = 9
	goexpressionParserR_CURLY                = 10
	goexpressionParserL_BRACKET              = 11
	goexpressionParserR_BRACKET              = 12
	goexpressionParserASSIGN                 = 13
	goexpressionParserCOMMA                  = 14
	goexpressionParserSEMI                   = 15
	goexpressionParserQUESTION               = 16
	goexpressionParserCOLON                  = 17
	goexpressionParserDOT                    = 18
	goexpressionParserNIL_LIT                = 19
	goexpressionParserLOGICAL_OR             = 20
	goexpressionParserLOGICAL_AND            = 21
	goexpressionParserEQUALS                 = 22
	goexpressionParserNOT_EQUALS             = 23
	goexpressionParserLESS                   = 24
	goexpressionParserLESS_OR_EQUALS         = 25
	goexpressionParserGREATER                = 26
	goexpressionParserGREATER_OR_EQUALS      = 27
	goexpressionParserOR                     = 28
	goexpressionParserDIV                    = 29
	goexpressionParserMOD                    = 30
	goexpressionParserLSHIFT                 = 31
	goexpressionParserRSHIFT                 = 32
	goexpressionParserBIT_CLEAR              = 33
	goexpressionParserEXCLAMATION            = 34
	goexpressionParserPLUS                   = 35
	goexpressionParserMINUS                  = 36
	goexpressionParserCARET                  = 37
	goexpressionParserSTAR                   = 38
	goexpressionParserAMPERSAND              = 39
	goexpressionParserDECIMAL_LIT            = 40
	goexpressionParserBINARY_LIT             = 41
	goexpressionParserOCTAL_LIT              = 42
	goexpressionParserHEX_LIT                = 43
	goexpressionParserFLOAT_LIT              = 44
	goexpressionParserDECIMAL_FLOAT_LIT      = 45
	goexpressionParserHEX_FLOAT_LIT          = 46
	goexpressionParserIDENTIFIER             = 47
	goexpressionParserRAW_STRING_LIT         = 48
	goexpressionParserINTERPRETED_STRING_LIT = 49
	goexpressionParserSINGLE_STRING_LIT      = 50
	goexpressionParserWS                     = 51
	goexpressionParserCOMMENT                = 52
	goexpressionParserTERMINATOR             = 53
	goexpressionParserLINE_COMMENT           = 54
	goexpressionParserWS_NLSEMI              = 55
	goexpressionParserCOMMENT_NLSEMI         = 56
	goexpressionParserLINE_COMMENT_NLSEMI    = 57
	goexpressionParserEOS                    = 58
)

// goexpressionParser rules.
const (
	goexpressionParserRULE_expressionStmt  = 0
	goexpressionParserRULE_expression      = 1
	goexpressionParserRULE_primaryExpr     = 2
	goexpressionParserRULE_operand         = 3
	goexpressionParserRULE_operandName     = 4
	goexpressionParserRULE_slice_          = 5
	goexpressionParserRULE_expressionColon = 6
	goexpressionParserRULE_index           = 7
	goexpressionParserRULE_basicLit        = 8
	goexpressionParserRULE_nil_lit         = 9
	goexpressionParserRULE_en_bool         = 10
	goexpressionParserRULE_float_lit       = 11
	goexpressionParserRULE_integer         = 12
	goexpressionParserRULE_expressionList  = 13
	goexpressionParserRULE_arguments       = 14
	goexpressionParserRULE_eos             = 15
	goexpressionParserRULE_identifier      = 16
	goexpressionParserRULE_string_         = 17
)

// IExpressionStmtContext is an interface to support dynamic dispatch.
type IExpressionStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionStmtContext differentiates from other interfaces.
	IsExpressionStmtContext()
}

type ExpressionStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStmtContext() *ExpressionStmtContext {
	var p = new(ExpressionStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goexpressionParserRULE_expressionStmt
	return p
}

func (*ExpressionStmtContext) IsExpressionStmtContext() {}

func NewExpressionStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStmtContext {
	var p = new(ExpressionStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goexpressionParserRULE_expressionStmt

	return p
}

func (s *ExpressionStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStmtContext) Expression() IExpressionContext {
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

func (s *ExpressionStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.EnterExpressionStmt(s)
	}
}

func (s *ExpressionStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.ExitExpressionStmt(s)
	}
}

func (s *ExpressionStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goexpressionVisitor:
		return t.VisitExpressionStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goexpressionParser) ExpressionStmt() (localctx IExpressionStmtContext) {
	this := p
	_ = this

	localctx = NewExpressionStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, goexpressionParserRULE_expressionStmt)

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
		p.SetState(36)
		p.expression(0)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetUnary_op returns the unary_op token.
	GetUnary_op() antlr.Token

	// GetMul_op returns the mul_op token.
	GetMul_op() antlr.Token

	// GetAdd_op returns the add_op token.
	GetAdd_op() antlr.Token

	// GetRel_op returns the rel_op token.
	GetRel_op() antlr.Token

	// SetUnary_op sets the unary_op token.
	SetUnary_op(antlr.Token)

	// SetMul_op sets the mul_op token.
	SetMul_op(antlr.Token)

	// SetAdd_op sets the add_op token.
	SetAdd_op(antlr.Token)

	// SetRel_op sets the rel_op token.
	SetRel_op(antlr.Token)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	unary_op antlr.Token
	mul_op   antlr.Token
	add_op   antlr.Token
	rel_op   antlr.Token
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goexpressionParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goexpressionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetUnary_op() antlr.Token { return s.unary_op }

func (s *ExpressionContext) GetMul_op() antlr.Token { return s.mul_op }

func (s *ExpressionContext) GetAdd_op() antlr.Token { return s.add_op }

func (s *ExpressionContext) GetRel_op() antlr.Token { return s.rel_op }

func (s *ExpressionContext) SetUnary_op(v antlr.Token) { s.unary_op = v }

func (s *ExpressionContext) SetMul_op(v antlr.Token) { s.mul_op = v }

func (s *ExpressionContext) SetAdd_op(v antlr.Token) { s.add_op = v }

func (s *ExpressionContext) SetRel_op(v antlr.Token) { s.rel_op = v }

func (s *ExpressionContext) PrimaryExpr() IPrimaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
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

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(goexpressionParserPLUS, 0)
}

func (s *ExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(goexpressionParserMINUS, 0)
}

func (s *ExpressionContext) EXCLAMATION() antlr.TerminalNode {
	return s.GetToken(goexpressionParserEXCLAMATION, 0)
}

func (s *ExpressionContext) EN_NOT() antlr.TerminalNode {
	return s.GetToken(goexpressionParserEN_NOT, 0)
}

func (s *ExpressionContext) CARET() antlr.TerminalNode {
	return s.GetToken(goexpressionParserCARET, 0)
}

func (s *ExpressionContext) STAR() antlr.TerminalNode {
	return s.GetToken(goexpressionParserSTAR, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(goexpressionParserDIV, 0)
}

func (s *ExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(goexpressionParserMOD, 0)
}

func (s *ExpressionContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(goexpressionParserLSHIFT, 0)
}

func (s *ExpressionContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(goexpressionParserRSHIFT, 0)
}

func (s *ExpressionContext) AMPERSAND() antlr.TerminalNode {
	return s.GetToken(goexpressionParserAMPERSAND, 0)
}

func (s *ExpressionContext) BIT_CLEAR() antlr.TerminalNode {
	return s.GetToken(goexpressionParserBIT_CLEAR, 0)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(goexpressionParserOR, 0)
}

func (s *ExpressionContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(goexpressionParserEQUALS, 0)
}

func (s *ExpressionContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(goexpressionParserNOT_EQUALS, 0)
}

func (s *ExpressionContext) LESS() antlr.TerminalNode {
	return s.GetToken(goexpressionParserLESS, 0)
}

func (s *ExpressionContext) LESS_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(goexpressionParserLESS_OR_EQUALS, 0)
}

func (s *ExpressionContext) GREATER() antlr.TerminalNode {
	return s.GetToken(goexpressionParserGREATER, 0)
}

func (s *ExpressionContext) GREATER_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(goexpressionParserGREATER_OR_EQUALS, 0)
}

func (s *ExpressionContext) LOGICAL_AND() antlr.TerminalNode {
	return s.GetToken(goexpressionParserLOGICAL_AND, 0)
}

func (s *ExpressionContext) LOGICAL_OR() antlr.TerminalNode {
	return s.GetToken(goexpressionParserLOGICAL_OR, 0)
}

func (s *ExpressionContext) EN_AND() antlr.TerminalNode {
	return s.GetToken(goexpressionParserEN_AND, 0)
}

func (s *ExpressionContext) EN_OR() antlr.TerminalNode {
	return s.GetToken(goexpressionParserEN_OR, 0)
}

func (s *ExpressionContext) EN_IN() antlr.TerminalNode {
	return s.GetToken(goexpressionParserEN_IN, 0)
}

func (s *ExpressionContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(goexpressionParserQUESTION, 0)
}

func (s *ExpressionContext) ExpressionColon() IExpressionColonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionColonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionColonContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goexpressionVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goexpressionParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *goexpressionParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, goexpressionParserRULE_expression, _p)
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
	p.SetState(42)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case goexpressionParserEN_TRUE, goexpressionParserEN_FALSE, goexpressionParserL_PAREN, goexpressionParserL_CURLY, goexpressionParserNIL_LIT, goexpressionParserDECIMAL_LIT, goexpressionParserBINARY_LIT, goexpressionParserOCTAL_LIT, goexpressionParserHEX_LIT, goexpressionParserFLOAT_LIT, goexpressionParserIDENTIFIER, goexpressionParserINTERPRETED_STRING_LIT, goexpressionParserSINGLE_STRING_LIT:
		{
			p.SetState(39)
			p.primaryExpr(0)
		}

	case goexpressionParserEN_NOT, goexpressionParserEXCLAMATION, goexpressionParserPLUS, goexpressionParserMINUS, goexpressionParserCARET:
		{
			p.SetState(40)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressionContext).unary_op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&257698037792) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressionContext).unary_op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(41)
			p.expression(10)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(71)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, goexpressionParserRULE_expression)
				p.SetState(44)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(45)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).mul_op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&841276719104) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).mul_op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(46)
					p.expression(10)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, goexpressionParserRULE_expression)
				p.SetState(47)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(48)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).add_op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&240786604032) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).add_op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(49)
					p.expression(9)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, goexpressionParserRULE_expression)
				p.SetState(50)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(51)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).rel_op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&264241152) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).rel_op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(52)
					p.expression(8)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, goexpressionParserRULE_expression)
				p.SetState(53)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(54)
					p.Match(goexpressionParserLOGICAL_AND)
				}
				{
					p.SetState(55)
					p.expression(7)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, goexpressionParserRULE_expression)
				p.SetState(56)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(57)
					p.Match(goexpressionParserLOGICAL_OR)
				}
				{
					p.SetState(58)
					p.expression(6)
				}

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, goexpressionParserRULE_expression)
				p.SetState(59)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(60)
					p.Match(goexpressionParserEN_AND)
				}
				{
					p.SetState(61)
					p.expression(5)
				}

			case 7:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, goexpressionParserRULE_expression)
				p.SetState(62)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(63)
					p.Match(goexpressionParserEN_OR)
				}
				{
					p.SetState(64)
					p.expression(4)
				}

			case 8:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, goexpressionParserRULE_expression)
				p.SetState(65)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(66)
					p.Match(goexpressionParserEN_IN)
				}
				{
					p.SetState(67)
					p.expression(3)
				}

			case 9:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, goexpressionParserRULE_expression)
				p.SetState(68)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(69)
					p.Match(goexpressionParserQUESTION)
				}
				{
					p.SetState(70)
					p.ExpressionColon()
				}

			}

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goexpressionParserRULE_primaryExpr
	return p
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goexpressionParserRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *PrimaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *PrimaryExprContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *PrimaryExprContext) Slice_() ISlice_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISlice_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISlice_Context)
}

func (s *PrimaryExprContext) Arguments() IArgumentsContext {
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

func (s *PrimaryExprContext) DOT() antlr.TerminalNode {
	return s.GetToken(goexpressionParserDOT, 0)
}

func (s *PrimaryExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(goexpressionParserIDENTIFIER, 0)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goexpressionVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goexpressionParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	return p.primaryExpr(0)
}

func (p *goexpressionParser) primaryExpr(_p int) (localctx IPrimaryExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, goexpressionParserRULE_primaryExpr, _p)

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
	{
		p.SetState(77)
		p.Operand()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, goexpressionParserRULE_primaryExpr)
			p.SetState(79)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(80)
					p.Match(goexpressionParserDOT)
				}
				{
					p.SetState(81)
					p.Match(goexpressionParserIDENTIFIER)
				}

			case 2:
				{
					p.SetState(82)
					p.Index()
				}

			case 3:
				{
					p.SetState(83)
					p.Slice_()
				}

			case 4:
				{
					p.SetState(84)
					p.Arguments()
				}

			}

		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goexpressionParserRULE_operand
	return p
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goexpressionParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) BasicLit() IBasicLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasicLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBasicLitContext)
}

func (s *OperandContext) OperandName() IOperandNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandNameContext)
}

func (s *OperandContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(goexpressionParserL_PAREN, 0)
}

func (s *OperandContext) Expression() IExpressionContext {
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

func (s *OperandContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(goexpressionParserR_PAREN, 0)
}

func (s *OperandContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(goexpressionParserL_CURLY, 0)
}

func (s *OperandContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(goexpressionParserR_CURLY, 0)
}

func (s *OperandContext) ExpressionList() IExpressionListContext {
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

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.ExitOperand(s)
	}
}

func (s *OperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goexpressionVisitor:
		return t.VisitOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goexpressionParser) Operand() (localctx IOperandContext) {
	this := p
	_ = this

	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, goexpressionParserRULE_operand)
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

	p.SetState(103)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case goexpressionParserEN_TRUE, goexpressionParserEN_FALSE, goexpressionParserNIL_LIT, goexpressionParserDECIMAL_LIT, goexpressionParserBINARY_LIT, goexpressionParserOCTAL_LIT, goexpressionParserHEX_LIT, goexpressionParserFLOAT_LIT, goexpressionParserINTERPRETED_STRING_LIT, goexpressionParserSINGLE_STRING_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.BasicLit()
		}

	case goexpressionParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.OperandName()
		}

	case goexpressionParserL_PAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(94)
			p.Match(goexpressionParserL_PAREN)
		}
		{
			p.SetState(95)
			p.expression(0)
		}
		{
			p.SetState(96)
			p.Match(goexpressionParserR_PAREN)
		}

	case goexpressionParserL_CURLY:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(98)
			p.Match(goexpressionParserL_CURLY)
		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1863929907643046) != 0 {
			{
				p.SetState(99)
				p.ExpressionList()
			}

		}
		{
			p.SetState(102)
			p.Match(goexpressionParserR_CURLY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperandNameContext is an interface to support dynamic dispatch.
type IOperandNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandNameContext differentiates from other interfaces.
	IsOperandNameContext()
}

type OperandNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandNameContext() *OperandNameContext {
	var p = new(OperandNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goexpressionParserRULE_operandName
	return p
}

func (*OperandNameContext) IsOperandNameContext() {}

func NewOperandNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandNameContext {
	var p = new(OperandNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goexpressionParserRULE_operandName

	return p
}

func (s *OperandNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandNameContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *OperandNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.EnterOperandName(s)
	}
}

func (s *OperandNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.ExitOperandName(s)
	}
}

func (s *OperandNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goexpressionVisitor:
		return t.VisitOperandName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goexpressionParser) OperandName() (localctx IOperandNameContext) {
	this := p
	_ = this

	localctx = NewOperandNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, goexpressionParserRULE_operandName)

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
		p.SetState(105)
		p.Identifier()
	}

	return localctx
}

// ISlice_Context is an interface to support dynamic dispatch.
type ISlice_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSlice_Context differentiates from other interfaces.
	IsSlice_Context()
}

type Slice_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySlice_Context() *Slice_Context {
	var p = new(Slice_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goexpressionParserRULE_slice_
	return p
}

func (*Slice_Context) IsSlice_Context() {}

func NewSlice_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Slice_Context {
	var p = new(Slice_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goexpressionParserRULE_slice_

	return p
}

func (s *Slice_Context) GetParser() antlr.Parser { return s.parser }

func (s *Slice_Context) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(goexpressionParserL_BRACKET, 0)
}

func (s *Slice_Context) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(goexpressionParserR_BRACKET, 0)
}

func (s *Slice_Context) ExpressionColon() IExpressionColonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionColonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionColonContext)
}

func (s *Slice_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Slice_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Slice_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.EnterSlice_(s)
	}
}

func (s *Slice_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.ExitSlice_(s)
	}
}

func (s *Slice_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goexpressionVisitor:
		return t.VisitSlice_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goexpressionParser) Slice_() (localctx ISlice_Context) {
	this := p
	_ = this

	localctx = NewSlice_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, goexpressionParserRULE_slice_)

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
		p.SetState(107)
		p.Match(goexpressionParserL_BRACKET)
	}

	{
		p.SetState(108)
		p.ExpressionColon()
	}

	{
		p.SetState(109)
		p.Match(goexpressionParserR_BRACKET)
	}

	return localctx
}

// IExpressionColonContext is an interface to support dynamic dispatch.
type IExpressionColonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionColonContext differentiates from other interfaces.
	IsExpressionColonContext()
}

type ExpressionColonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionColonContext() *ExpressionColonContext {
	var p = new(ExpressionColonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goexpressionParserRULE_expressionColon
	return p
}

func (*ExpressionColonContext) IsExpressionColonContext() {}

func NewExpressionColonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionColonContext {
	var p = new(ExpressionColonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goexpressionParserRULE_expressionColon

	return p
}

func (s *ExpressionColonContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionColonContext) AllExpression() []IExpressionContext {
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

func (s *ExpressionColonContext) Expression(i int) IExpressionContext {
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

func (s *ExpressionColonContext) COLON() antlr.TerminalNode {
	return s.GetToken(goexpressionParserCOLON, 0)
}

func (s *ExpressionColonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionColonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionColonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.EnterExpressionColon(s)
	}
}

func (s *ExpressionColonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.ExitExpressionColon(s)
	}
}

func (s *ExpressionColonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goexpressionVisitor:
		return t.VisitExpressionColon(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goexpressionParser) ExpressionColon() (localctx IExpressionColonContext) {
	this := p
	_ = this

	localctx = NewExpressionColonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, goexpressionParserRULE_expressionColon)

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
		p.SetState(111)
		p.expression(0)
	}
	{
		p.SetState(112)
		p.Match(goexpressionParserCOLON)
	}
	{
		p.SetState(113)
		p.expression(0)
	}

	return localctx
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goexpressionParserRULE_index
	return p
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goexpressionParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(goexpressionParserL_BRACKET, 0)
}

func (s *IndexContext) Expression() IExpressionContext {
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

func (s *IndexContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(goexpressionParserR_BRACKET, 0)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goexpressionVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goexpressionParser) Index() (localctx IIndexContext) {
	this := p
	_ = this

	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, goexpressionParserRULE_index)

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
		p.SetState(115)
		p.Match(goexpressionParserL_BRACKET)
	}
	{
		p.SetState(116)
		p.expression(0)
	}
	{
		p.SetState(117)
		p.Match(goexpressionParserR_BRACKET)
	}

	return localctx
}

// IBasicLitContext is an interface to support dynamic dispatch.
type IBasicLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBasicLitContext differentiates from other interfaces.
	IsBasicLitContext()
}

type BasicLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasicLitContext() *BasicLitContext {
	var p = new(BasicLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goexpressionParserRULE_basicLit
	return p
}

func (*BasicLitContext) IsBasicLitContext() {}

func NewBasicLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasicLitContext {
	var p = new(BasicLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goexpressionParserRULE_basicLit

	return p
}

func (s *BasicLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BasicLitContext) Nil_lit() INil_litContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INil_litContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INil_litContext)
}

func (s *BasicLitContext) En_bool() IEn_boolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEn_boolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEn_boolContext)
}

func (s *BasicLitContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *BasicLitContext) String_() IString_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *BasicLitContext) Float_lit() IFloat_litContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloat_litContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloat_litContext)
}

func (s *BasicLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasicLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.EnterBasicLit(s)
	}
}

func (s *BasicLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.ExitBasicLit(s)
	}
}

func (s *BasicLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goexpressionVisitor:
		return t.VisitBasicLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goexpressionParser) BasicLit() (localctx IBasicLitContext) {
	this := p
	_ = this

	localctx = NewBasicLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, goexpressionParserRULE_basicLit)

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

	switch p.GetTokenStream().LA(1) {
	case goexpressionParserNIL_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Nil_lit()
		}

	case goexpressionParserEN_TRUE, goexpressionParserEN_FALSE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.En_bool()
		}

	case goexpressionParserDECIMAL_LIT, goexpressionParserBINARY_LIT, goexpressionParserOCTAL_LIT, goexpressionParserHEX_LIT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(121)
			p.Integer()
		}

	case goexpressionParserINTERPRETED_STRING_LIT, goexpressionParserSINGLE_STRING_LIT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(122)
			p.String_()
		}

	case goexpressionParserFLOAT_LIT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(123)
			p.Float_lit()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INil_litContext is an interface to support dynamic dispatch.
type INil_litContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNil_litContext differentiates from other interfaces.
	IsNil_litContext()
}

type Nil_litContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNil_litContext() *Nil_litContext {
	var p = new(Nil_litContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goexpressionParserRULE_nil_lit
	return p
}

func (*Nil_litContext) IsNil_litContext() {}

func NewNil_litContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Nil_litContext {
	var p = new(Nil_litContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goexpressionParserRULE_nil_lit

	return p
}

func (s *Nil_litContext) GetParser() antlr.Parser { return s.parser }

func (s *Nil_litContext) NIL_LIT() antlr.TerminalNode {
	return s.GetToken(goexpressionParserNIL_LIT, 0)
}

func (s *Nil_litContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Nil_litContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Nil_litContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.EnterNil_lit(s)
	}
}

func (s *Nil_litContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.ExitNil_lit(s)
	}
}

func (s *Nil_litContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goexpressionVisitor:
		return t.VisitNil_lit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goexpressionParser) Nil_lit() (localctx INil_litContext) {
	this := p
	_ = this

	localctx = NewNil_litContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, goexpressionParserRULE_nil_lit)

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
		p.SetState(126)
		p.Match(goexpressionParserNIL_LIT)
	}

	return localctx
}

// IEn_boolContext is an interface to support dynamic dispatch.
type IEn_boolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEn_boolContext differentiates from other interfaces.
	IsEn_boolContext()
}

type En_boolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEn_boolContext() *En_boolContext {
	var p = new(En_boolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goexpressionParserRULE_en_bool
	return p
}

func (*En_boolContext) IsEn_boolContext() {}

func NewEn_boolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *En_boolContext {
	var p = new(En_boolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goexpressionParserRULE_en_bool

	return p
}

func (s *En_boolContext) GetParser() antlr.Parser { return s.parser }

func (s *En_boolContext) EN_FALSE() antlr.TerminalNode {
	return s.GetToken(goexpressionParserEN_FALSE, 0)
}

func (s *En_boolContext) EN_TRUE() antlr.TerminalNode {
	return s.GetToken(goexpressionParserEN_TRUE, 0)
}

func (s *En_boolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *En_boolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *En_boolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.EnterEn_bool(s)
	}
}

func (s *En_boolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.ExitEn_bool(s)
	}
}

func (s *En_boolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goexpressionVisitor:
		return t.VisitEn_bool(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goexpressionParser) En_bool() (localctx IEn_boolContext) {
	this := p
	_ = this

	localctx = NewEn_boolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, goexpressionParserRULE_en_bool)
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
		p.SetState(128)
		_la = p.GetTokenStream().LA(1)

		if !(_la == goexpressionParserEN_TRUE || _la == goexpressionParserEN_FALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFloat_litContext is an interface to support dynamic dispatch.
type IFloat_litContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloat_litContext differentiates from other interfaces.
	IsFloat_litContext()
}

type Float_litContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloat_litContext() *Float_litContext {
	var p = new(Float_litContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goexpressionParserRULE_float_lit
	return p
}

func (*Float_litContext) IsFloat_litContext() {}

func NewFloat_litContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Float_litContext {
	var p = new(Float_litContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goexpressionParserRULE_float_lit

	return p
}

func (s *Float_litContext) GetParser() antlr.Parser { return s.parser }

func (s *Float_litContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(goexpressionParserFLOAT_LIT, 0)
}

func (s *Float_litContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Float_litContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Float_litContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.EnterFloat_lit(s)
	}
}

func (s *Float_litContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.ExitFloat_lit(s)
	}
}

func (s *Float_litContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goexpressionVisitor:
		return t.VisitFloat_lit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goexpressionParser) Float_lit() (localctx IFloat_litContext) {
	this := p
	_ = this

	localctx = NewFloat_litContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, goexpressionParserRULE_float_lit)

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
		p.SetState(130)
		p.Match(goexpressionParserFLOAT_LIT)
	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goexpressionParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goexpressionParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) DECIMAL_LIT() antlr.TerminalNode {
	return s.GetToken(goexpressionParserDECIMAL_LIT, 0)
}

func (s *IntegerContext) BINARY_LIT() antlr.TerminalNode {
	return s.GetToken(goexpressionParserBINARY_LIT, 0)
}

func (s *IntegerContext) OCTAL_LIT() antlr.TerminalNode {
	return s.GetToken(goexpressionParserOCTAL_LIT, 0)
}

func (s *IntegerContext) HEX_LIT() antlr.TerminalNode {
	return s.GetToken(goexpressionParserHEX_LIT, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goexpressionVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goexpressionParser) Integer() (localctx IIntegerContext) {
	this := p
	_ = this

	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, goexpressionParserRULE_integer)
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
		p.SetState(132)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16492674416640) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.RuleIndex = goexpressionParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goexpressionParserRULE_expressionList

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
	return s.GetTokens(goexpressionParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(goexpressionParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goexpressionVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goexpressionParser) ExpressionList() (localctx IExpressionListContext) {
	this := p
	_ = this

	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, goexpressionParserRULE_expressionList)
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
		p.SetState(134)
		p.expression(0)
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == goexpressionParserCOMMA {
		{
			p.SetState(135)
			p.Match(goexpressionParserCOMMA)
		}
		{
			p.SetState(136)
			p.expression(0)
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = goexpressionParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goexpressionParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(goexpressionParserL_PAREN, 0)
}

func (s *ArgumentsContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(goexpressionParserR_PAREN, 0)
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
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goexpressionVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goexpressionParser) Arguments() (localctx IArgumentsContext) {
	this := p
	_ = this

	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, goexpressionParserRULE_arguments)
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
		p.SetState(142)
		p.Match(goexpressionParserL_PAREN)
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1863929907643046) != 0 {
		{
			p.SetState(143)
			p.ExpressionList()
		}

	}
	{
		p.SetState(146)
		p.Match(goexpressionParserR_PAREN)
	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goexpressionParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goexpressionParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) SEMI() antlr.TerminalNode {
	return s.GetToken(goexpressionParserSEMI, 0)
}

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(goexpressionParserEOF, 0)
}

func (s *EosContext) EOS() antlr.TerminalNode {
	return s.GetToken(goexpressionParserEOS, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.ExitEos(s)
	}
}

func (s *EosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goexpressionVisitor:
		return t.VisitEos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goexpressionParser) Eos() (localctx IEosContext) {
	this := p
	_ = this

	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, goexpressionParserRULE_eos)
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
		p.SetState(148)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la - -1)) & ^0x3f) == 0 && ((int64(1)<<(_la - -1))&576460752303489025) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goexpressionParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goexpressionParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(goexpressionParserIDENTIFIER, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goexpressionVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goexpressionParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, goexpressionParserRULE_identifier)

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
		p.SetState(150)
		p.Match(goexpressionParserIDENTIFIER)
	}

	return localctx
}

// IString_Context is an interface to support dynamic dispatch.
type IString_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_Context differentiates from other interfaces.
	IsString_Context()
}

type String_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_Context() *String_Context {
	var p = new(String_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = goexpressionParserRULE_string_
	return p
}

func (*String_Context) IsString_Context() {}

func NewString_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_Context {
	var p = new(String_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = goexpressionParserRULE_string_

	return p
}

func (s *String_Context) GetParser() antlr.Parser { return s.parser }

func (s *String_Context) INTERPRETED_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(goexpressionParserINTERPRETED_STRING_LIT, 0)
}

func (s *String_Context) SINGLE_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(goexpressionParserSINGLE_STRING_LIT, 0)
}

func (s *String_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.EnterString_(s)
	}
}

func (s *String_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(goexpressionListener); ok {
		listenerT.ExitString_(s)
	}
}

func (s *String_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case goexpressionVisitor:
		return t.VisitString_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *goexpressionParser) String_() (localctx IString_Context) {
	this := p
	_ = this

	localctx = NewString_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, goexpressionParserRULE_string_)
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
		_la = p.GetTokenStream().LA(1)

		if !(_la == goexpressionParserINTERPRETED_STRING_LIT || _la == goexpressionParserSINGLE_STRING_LIT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *goexpressionParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 2:
		var t *PrimaryExprContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExprContext)
		}
		return p.PrimaryExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *goexpressionParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *goexpressionParser) PrimaryExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 9:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
