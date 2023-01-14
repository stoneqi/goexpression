// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type goexpressionLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var goexpressionlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func goexpressionlexerLexerInit() {
	staticData := &goexpressionlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "", "'and'", "'in'", "'not'", "'or'", "'('", "')'", "'{'", "'}'",
		"'['", "']'", "'='", "','", "';'", "':'", "'.'", "'nil'", "'||'", "'&&'",
		"'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'|'", "'/'", "'%'", "'<<'",
		"'>>'", "'&^'", "'!'", "'+'", "'-'", "'^'", "'*'", "'&'",
	}
	staticData.symbolicNames = []string{
		"", "EN_TRUE", "EN_FALSE", "EN_AND", "EN_IN", "EN_NOT", "EN_OR", "L_PAREN",
		"R_PAREN", "L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET", "ASSIGN",
		"COMMA", "SEMI", "COLON", "DOT", "NIL_LIT", "LOGICAL_OR", "LOGICAL_AND",
		"EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS",
		"OR", "DIV", "MOD", "LSHIFT", "RSHIFT", "BIT_CLEAR", "EXCLAMATION",
		"PLUS", "MINUS", "CARET", "STAR", "AMPERSAND", "DECIMAL_LIT", "BINARY_LIT",
		"OCTAL_LIT", "HEX_LIT", "FLOAT_LIT", "DECIMAL_FLOAT_LIT", "HEX_FLOAT_LIT",
		"IDENTIFIER", "RAW_STRING_LIT", "INTERPRETED_STRING_LIT", "SINGLE_STRING_LIT",
		"WS", "COMMENT", "TERMINATOR", "LINE_COMMENT", "WS_NLSEMI", "COMMENT_NLSEMI",
		"LINE_COMMENT_NLSEMI", "EOS",
	}
	staticData.ruleNames = []string{
		"EN_TRUE", "EN_FALSE", "EN_AND", "EN_IN", "EN_NOT", "EN_OR", "L_PAREN",
		"R_PAREN", "L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET", "ASSIGN",
		"COMMA", "SEMI", "COLON", "DOT", "NIL_LIT", "LOGICAL_OR", "LOGICAL_AND",
		"EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS",
		"OR", "DIV", "MOD", "LSHIFT", "RSHIFT", "BIT_CLEAR", "EXCLAMATION",
		"PLUS", "MINUS", "CARET", "STAR", "AMPERSAND", "DECIMAL_LIT", "BINARY_LIT",
		"OCTAL_LIT", "HEX_LIT", "FLOAT_LIT", "DECIMAL_FLOAT_LIT", "HEX_FLOAT_LIT",
		"HEX_MANTISSA", "HEX_EXPONENT", "IDENTIFIER", "RAW_STRING_LIT", "INTERPRETED_STRING_LIT",
		"SINGLE_STRING_LIT", "WS", "COMMENT", "TERMINATOR", "LINE_COMMENT",
		"ESCAPED_VALUE", "DECIMALS", "OCTAL_DIGIT", "HEX_DIGIT", "BIN_DIGIT",
		"EXPONENT", "WS_NLSEMI", "COMMENT_NLSEMI", "LINE_COMMENT_NLSEMI", "EOS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 57, 541, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 144, 8, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
		1, 161, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30,
		1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1,
		35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 3, 38, 255, 8, 38,
		1, 38, 5, 38, 258, 8, 38, 10, 38, 12, 38, 261, 9, 38, 3, 38, 263, 8, 38,
		1, 39, 1, 39, 1, 39, 3, 39, 268, 8, 39, 1, 39, 4, 39, 271, 8, 39, 11, 39,
		12, 39, 272, 1, 40, 1, 40, 3, 40, 277, 8, 40, 1, 40, 3, 40, 280, 8, 40,
		1, 40, 4, 40, 283, 8, 40, 11, 40, 12, 40, 284, 1, 41, 1, 41, 1, 41, 3,
		41, 290, 8, 41, 1, 41, 4, 41, 293, 8, 41, 11, 41, 12, 41, 294, 1, 42, 1,
		42, 3, 42, 299, 8, 42, 1, 43, 1, 43, 1, 43, 3, 43, 304, 8, 43, 1, 43, 3,
		43, 307, 8, 43, 1, 43, 3, 43, 310, 8, 43, 1, 43, 1, 43, 1, 43, 3, 43, 315,
		8, 43, 3, 43, 317, 8, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 3,
		45, 325, 8, 45, 1, 45, 4, 45, 328, 8, 45, 11, 45, 12, 45, 329, 1, 45, 1,
		45, 3, 45, 334, 8, 45, 1, 45, 5, 45, 337, 8, 45, 10, 45, 12, 45, 340, 9,
		45, 3, 45, 342, 8, 45, 1, 45, 1, 45, 1, 45, 3, 45, 347, 8, 45, 1, 45, 5,
		45, 350, 8, 45, 10, 45, 12, 45, 353, 9, 45, 3, 45, 355, 8, 45, 1, 46, 1,
		46, 3, 46, 359, 8, 46, 1, 46, 1, 46, 1, 47, 1, 47, 5, 47, 365, 8, 47, 10,
		47, 12, 47, 368, 9, 47, 1, 48, 1, 48, 5, 48, 372, 8, 48, 10, 48, 12, 48,
		375, 9, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 5, 49, 382, 8, 49, 10, 49,
		12, 49, 385, 9, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 5, 50, 392, 8, 50,
		10, 50, 12, 50, 395, 9, 50, 1, 50, 1, 50, 1, 51, 4, 51, 400, 8, 51, 11,
		51, 12, 51, 401, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 52, 5, 52, 410,
		8, 52, 10, 52, 12, 52, 413, 9, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1,
		53, 4, 53, 421, 8, 53, 11, 53, 12, 53, 422, 1, 53, 1, 53, 1, 54, 1, 54,
		1, 54, 1, 54, 5, 54, 431, 8, 54, 10, 54, 12, 54, 434, 9, 54, 1, 54, 1,
		54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55,
		1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1,
		55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 3, 55, 464, 8, 55, 1, 56, 1, 56,
		3, 56, 468, 8, 56, 1, 56, 5, 56, 471, 8, 56, 10, 56, 12, 56, 474, 9, 56,
		1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 3, 60, 484, 8,
		60, 1, 60, 1, 60, 1, 61, 4, 61, 489, 8, 61, 11, 61, 12, 61, 490, 1, 61,
		1, 61, 1, 62, 1, 62, 1, 62, 1, 62, 5, 62, 499, 8, 62, 10, 62, 12, 62, 502,
		9, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 63, 1, 63, 1, 63, 1, 63, 5,
		63, 513, 8, 63, 10, 63, 12, 63, 516, 9, 63, 1, 63, 1, 63, 1, 64, 4, 64,
		521, 8, 64, 11, 64, 12, 64, 522, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 5,
		64, 530, 8, 64, 10, 64, 12, 64, 533, 9, 64, 1, 64, 1, 64, 1, 64, 3, 64,
		538, 8, 64, 1, 64, 1, 64, 3, 411, 500, 531, 0, 65, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45,
		23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63,
		32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81,
		41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 0, 93, 0, 95, 46, 97, 47, 99, 48,
		101, 49, 103, 50, 105, 51, 107, 52, 109, 53, 111, 0, 113, 0, 115, 0, 117,
		0, 119, 0, 121, 0, 123, 54, 125, 55, 127, 56, 129, 57, 1, 0, 18, 1, 0,
		49, 57, 1, 0, 48, 57, 2, 0, 66, 66, 98, 98, 2, 0, 79, 79, 111, 111, 2,
		0, 88, 88, 120, 120, 2, 0, 80, 80, 112, 112, 2, 0, 43, 43, 45, 45, 3, 0,
		65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 96,
		96, 2, 0, 34, 34, 92, 92, 2, 0, 9, 9, 32, 32, 2, 0, 10, 10, 13, 13, 9,
		0, 34, 34, 39, 39, 92, 92, 97, 98, 102, 102, 110, 110, 114, 114, 116, 116,
		118, 118, 1, 0, 48, 55, 3, 0, 48, 57, 65, 70, 97, 102, 1, 0, 48, 49, 2,
		0, 69, 69, 101, 101, 586, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0,
		67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0,
		0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0,
		0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0,
		0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101,
		1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0,
		0, 109, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0, 127, 1,
		0, 0, 0, 0, 129, 1, 0, 0, 0, 1, 143, 1, 0, 0, 0, 3, 160, 1, 0, 0, 0, 5,
		162, 1, 0, 0, 0, 7, 166, 1, 0, 0, 0, 9, 169, 1, 0, 0, 0, 11, 173, 1, 0,
		0, 0, 13, 176, 1, 0, 0, 0, 15, 178, 1, 0, 0, 0, 17, 180, 1, 0, 0, 0, 19,
		182, 1, 0, 0, 0, 21, 184, 1, 0, 0, 0, 23, 186, 1, 0, 0, 0, 25, 188, 1,
		0, 0, 0, 27, 190, 1, 0, 0, 0, 29, 192, 1, 0, 0, 0, 31, 194, 1, 0, 0, 0,
		33, 196, 1, 0, 0, 0, 35, 198, 1, 0, 0, 0, 37, 202, 1, 0, 0, 0, 39, 205,
		1, 0, 0, 0, 41, 208, 1, 0, 0, 0, 43, 211, 1, 0, 0, 0, 45, 214, 1, 0, 0,
		0, 47, 216, 1, 0, 0, 0, 49, 219, 1, 0, 0, 0, 51, 221, 1, 0, 0, 0, 53, 224,
		1, 0, 0, 0, 55, 226, 1, 0, 0, 0, 57, 228, 1, 0, 0, 0, 59, 230, 1, 0, 0,
		0, 61, 233, 1, 0, 0, 0, 63, 236, 1, 0, 0, 0, 65, 239, 1, 0, 0, 0, 67, 241,
		1, 0, 0, 0, 69, 243, 1, 0, 0, 0, 71, 245, 1, 0, 0, 0, 73, 247, 1, 0, 0,
		0, 75, 249, 1, 0, 0, 0, 77, 262, 1, 0, 0, 0, 79, 264, 1, 0, 0, 0, 81, 274,
		1, 0, 0, 0, 83, 286, 1, 0, 0, 0, 85, 298, 1, 0, 0, 0, 87, 316, 1, 0, 0,
		0, 89, 318, 1, 0, 0, 0, 91, 354, 1, 0, 0, 0, 93, 356, 1, 0, 0, 0, 95, 362,
		1, 0, 0, 0, 97, 369, 1, 0, 0, 0, 99, 378, 1, 0, 0, 0, 101, 388, 1, 0, 0,
		0, 103, 399, 1, 0, 0, 0, 105, 405, 1, 0, 0, 0, 107, 420, 1, 0, 0, 0, 109,
		426, 1, 0, 0, 0, 111, 437, 1, 0, 0, 0, 113, 465, 1, 0, 0, 0, 115, 475,
		1, 0, 0, 0, 117, 477, 1, 0, 0, 0, 119, 479, 1, 0, 0, 0, 121, 481, 1, 0,
		0, 0, 123, 488, 1, 0, 0, 0, 125, 494, 1, 0, 0, 0, 127, 508, 1, 0, 0, 0,
		129, 537, 1, 0, 0, 0, 131, 132, 5, 84, 0, 0, 132, 133, 5, 114, 0, 0, 133,
		134, 5, 117, 0, 0, 134, 144, 5, 101, 0, 0, 135, 136, 5, 116, 0, 0, 136,
		137, 5, 114, 0, 0, 137, 138, 5, 117, 0, 0, 138, 144, 5, 101, 0, 0, 139,
		140, 5, 84, 0, 0, 140, 141, 5, 82, 0, 0, 141, 142, 5, 85, 0, 0, 142, 144,
		5, 69, 0, 0, 143, 131, 1, 0, 0, 0, 143, 135, 1, 0, 0, 0, 143, 139, 1, 0,
		0, 0, 144, 2, 1, 0, 0, 0, 145, 146, 5, 70, 0, 0, 146, 147, 5, 97, 0, 0,
		147, 148, 5, 108, 0, 0, 148, 149, 5, 115, 0, 0, 149, 161, 5, 101, 0, 0,
		150, 151, 5, 102, 0, 0, 151, 152, 5, 97, 0, 0, 152, 153, 5, 108, 0, 0,
		153, 154, 5, 115, 0, 0, 154, 161, 5, 101, 0, 0, 155, 156, 5, 70, 0, 0,
		156, 157, 5, 65, 0, 0, 157, 158, 5, 76, 0, 0, 158, 159, 5, 83, 0, 0, 159,
		161, 5, 69, 0, 0, 160, 145, 1, 0, 0, 0, 160, 150, 1, 0, 0, 0, 160, 155,
		1, 0, 0, 0, 161, 4, 1, 0, 0, 0, 162, 163, 5, 97, 0, 0, 163, 164, 5, 110,
		0, 0, 164, 165, 5, 100, 0, 0, 165, 6, 1, 0, 0, 0, 166, 167, 5, 105, 0,
		0, 167, 168, 5, 110, 0, 0, 168, 8, 1, 0, 0, 0, 169, 170, 5, 110, 0, 0,
		170, 171, 5, 111, 0, 0, 171, 172, 5, 116, 0, 0, 172, 10, 1, 0, 0, 0, 173,
		174, 5, 111, 0, 0, 174, 175, 5, 114, 0, 0, 175, 12, 1, 0, 0, 0, 176, 177,
		5, 40, 0, 0, 177, 14, 1, 0, 0, 0, 178, 179, 5, 41, 0, 0, 179, 16, 1, 0,
		0, 0, 180, 181, 5, 123, 0, 0, 181, 18, 1, 0, 0, 0, 182, 183, 5, 125, 0,
		0, 183, 20, 1, 0, 0, 0, 184, 185, 5, 91, 0, 0, 185, 22, 1, 0, 0, 0, 186,
		187, 5, 93, 0, 0, 187, 24, 1, 0, 0, 0, 188, 189, 5, 61, 0, 0, 189, 26,
		1, 0, 0, 0, 190, 191, 5, 44, 0, 0, 191, 28, 1, 0, 0, 0, 192, 193, 5, 59,
		0, 0, 193, 30, 1, 0, 0, 0, 194, 195, 5, 58, 0, 0, 195, 32, 1, 0, 0, 0,
		196, 197, 5, 46, 0, 0, 197, 34, 1, 0, 0, 0, 198, 199, 5, 110, 0, 0, 199,
		200, 5, 105, 0, 0, 200, 201, 5, 108, 0, 0, 201, 36, 1, 0, 0, 0, 202, 203,
		5, 124, 0, 0, 203, 204, 5, 124, 0, 0, 204, 38, 1, 0, 0, 0, 205, 206, 5,
		38, 0, 0, 206, 207, 5, 38, 0, 0, 207, 40, 1, 0, 0, 0, 208, 209, 5, 61,
		0, 0, 209, 210, 5, 61, 0, 0, 210, 42, 1, 0, 0, 0, 211, 212, 5, 33, 0, 0,
		212, 213, 5, 61, 0, 0, 213, 44, 1, 0, 0, 0, 214, 215, 5, 60, 0, 0, 215,
		46, 1, 0, 0, 0, 216, 217, 5, 60, 0, 0, 217, 218, 5, 61, 0, 0, 218, 48,
		1, 0, 0, 0, 219, 220, 5, 62, 0, 0, 220, 50, 1, 0, 0, 0, 221, 222, 5, 62,
		0, 0, 222, 223, 5, 61, 0, 0, 223, 52, 1, 0, 0, 0, 224, 225, 5, 124, 0,
		0, 225, 54, 1, 0, 0, 0, 226, 227, 5, 47, 0, 0, 227, 56, 1, 0, 0, 0, 228,
		229, 5, 37, 0, 0, 229, 58, 1, 0, 0, 0, 230, 231, 5, 60, 0, 0, 231, 232,
		5, 60, 0, 0, 232, 60, 1, 0, 0, 0, 233, 234, 5, 62, 0, 0, 234, 235, 5, 62,
		0, 0, 235, 62, 1, 0, 0, 0, 236, 237, 5, 38, 0, 0, 237, 238, 5, 94, 0, 0,
		238, 64, 1, 0, 0, 0, 239, 240, 5, 33, 0, 0, 240, 66, 1, 0, 0, 0, 241, 242,
		5, 43, 0, 0, 242, 68, 1, 0, 0, 0, 243, 244, 5, 45, 0, 0, 244, 70, 1, 0,
		0, 0, 245, 246, 5, 94, 0, 0, 246, 72, 1, 0, 0, 0, 247, 248, 5, 42, 0, 0,
		248, 74, 1, 0, 0, 0, 249, 250, 5, 38, 0, 0, 250, 76, 1, 0, 0, 0, 251, 263,
		5, 48, 0, 0, 252, 259, 7, 0, 0, 0, 253, 255, 5, 95, 0, 0, 254, 253, 1,
		0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256, 258, 7, 1, 0,
		0, 257, 254, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259,
		260, 1, 0, 0, 0, 260, 263, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 262, 251,
		1, 0, 0, 0, 262, 252, 1, 0, 0, 0, 263, 78, 1, 0, 0, 0, 264, 265, 5, 48,
		0, 0, 265, 270, 7, 2, 0, 0, 266, 268, 5, 95, 0, 0, 267, 266, 1, 0, 0, 0,
		267, 268, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 271, 3, 119, 59, 0, 270,
		267, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 272, 273,
		1, 0, 0, 0, 273, 80, 1, 0, 0, 0, 274, 276, 5, 48, 0, 0, 275, 277, 7, 3,
		0, 0, 276, 275, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 282, 1, 0, 0, 0,
		278, 280, 5, 95, 0, 0, 279, 278, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280,
		281, 1, 0, 0, 0, 281, 283, 3, 115, 57, 0, 282, 279, 1, 0, 0, 0, 283, 284,
		1, 0, 0, 0, 284, 282, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 82, 1, 0,
		0, 0, 286, 287, 5, 48, 0, 0, 287, 292, 7, 4, 0, 0, 288, 290, 5, 95, 0,
		0, 289, 288, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291,
		293, 3, 117, 58, 0, 292, 289, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294, 292,
		1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 84, 1, 0, 0, 0, 296, 299, 3, 87,
		43, 0, 297, 299, 3, 89, 44, 0, 298, 296, 1, 0, 0, 0, 298, 297, 1, 0, 0,
		0, 299, 86, 1, 0, 0, 0, 300, 309, 3, 113, 56, 0, 301, 303, 5, 46, 0, 0,
		302, 304, 3, 113, 56, 0, 303, 302, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304,
		306, 1, 0, 0, 0, 305, 307, 3, 121, 60, 0, 306, 305, 1, 0, 0, 0, 306, 307,
		1, 0, 0, 0, 307, 310, 1, 0, 0, 0, 308, 310, 3, 121, 60, 0, 309, 301, 1,
		0, 0, 0, 309, 308, 1, 0, 0, 0, 310, 317, 1, 0, 0, 0, 311, 312, 5, 46, 0,
		0, 312, 314, 3, 113, 56, 0, 313, 315, 3, 121, 60, 0, 314, 313, 1, 0, 0,
		0, 314, 315, 1, 0, 0, 0, 315, 317, 1, 0, 0, 0, 316, 300, 1, 0, 0, 0, 316,
		311, 1, 0, 0, 0, 317, 88, 1, 0, 0, 0, 318, 319, 5, 48, 0, 0, 319, 320,
		7, 4, 0, 0, 320, 321, 3, 91, 45, 0, 321, 322, 3, 93, 46, 0, 322, 90, 1,
		0, 0, 0, 323, 325, 5, 95, 0, 0, 324, 323, 1, 0, 0, 0, 324, 325, 1, 0, 0,
		0, 325, 326, 1, 0, 0, 0, 326, 328, 3, 117, 58, 0, 327, 324, 1, 0, 0, 0,
		328, 329, 1, 0, 0, 0, 329, 327, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330,
		341, 1, 0, 0, 0, 331, 338, 5, 46, 0, 0, 332, 334, 5, 95, 0, 0, 333, 332,
		1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 335, 337, 3, 117,
		58, 0, 336, 333, 1, 0, 0, 0, 337, 340, 1, 0, 0, 0, 338, 336, 1, 0, 0, 0,
		338, 339, 1, 0, 0, 0, 339, 342, 1, 0, 0, 0, 340, 338, 1, 0, 0, 0, 341,
		331, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342, 355, 1, 0, 0, 0, 343, 344,
		5, 46, 0, 0, 344, 351, 3, 117, 58, 0, 345, 347, 5, 95, 0, 0, 346, 345,
		1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 348, 1, 0, 0, 0, 348, 350, 3, 117,
		58, 0, 349, 346, 1, 0, 0, 0, 350, 353, 1, 0, 0, 0, 351, 349, 1, 0, 0, 0,
		351, 352, 1, 0, 0, 0, 352, 355, 1, 0, 0, 0, 353, 351, 1, 0, 0, 0, 354,
		327, 1, 0, 0, 0, 354, 343, 1, 0, 0, 0, 355, 92, 1, 0, 0, 0, 356, 358, 7,
		5, 0, 0, 357, 359, 7, 6, 0, 0, 358, 357, 1, 0, 0, 0, 358, 359, 1, 0, 0,
		0, 359, 360, 1, 0, 0, 0, 360, 361, 3, 113, 56, 0, 361, 94, 1, 0, 0, 0,
		362, 366, 7, 7, 0, 0, 363, 365, 7, 8, 0, 0, 364, 363, 1, 0, 0, 0, 365,
		368, 1, 0, 0, 0, 366, 364, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367, 96, 1,
		0, 0, 0, 368, 366, 1, 0, 0, 0, 369, 373, 5, 96, 0, 0, 370, 372, 8, 9, 0,
		0, 371, 370, 1, 0, 0, 0, 372, 375, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 373,
		374, 1, 0, 0, 0, 374, 376, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 376, 377,
		5, 96, 0, 0, 377, 98, 1, 0, 0, 0, 378, 383, 5, 34, 0, 0, 379, 382, 8, 10,
		0, 0, 380, 382, 3, 111, 55, 0, 381, 379, 1, 0, 0, 0, 381, 380, 1, 0, 0,
		0, 382, 385, 1, 0, 0, 0, 383, 381, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384,
		386, 1, 0, 0, 0, 385, 383, 1, 0, 0, 0, 386, 387, 5, 34, 0, 0, 387, 100,
		1, 0, 0, 0, 388, 393, 5, 39, 0, 0, 389, 392, 8, 10, 0, 0, 390, 392, 3,
		111, 55, 0, 391, 389, 1, 0, 0, 0, 391, 390, 1, 0, 0, 0, 392, 395, 1, 0,
		0, 0, 393, 391, 1, 0, 0, 0, 393, 394, 1, 0, 0, 0, 394, 396, 1, 0, 0, 0,
		395, 393, 1, 0, 0, 0, 396, 397, 5, 39, 0, 0, 397, 102, 1, 0, 0, 0, 398,
		400, 7, 11, 0, 0, 399, 398, 1, 0, 0, 0, 400, 401, 1, 0, 0, 0, 401, 399,
		1, 0, 0, 0, 401, 402, 1, 0, 0, 0, 402, 403, 1, 0, 0, 0, 403, 404, 6, 51,
		0, 0, 404, 104, 1, 0, 0, 0, 405, 406, 5, 47, 0, 0, 406, 407, 5, 42, 0,
		0, 407, 411, 1, 0, 0, 0, 408, 410, 9, 0, 0, 0, 409, 408, 1, 0, 0, 0, 410,
		413, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0, 411, 409, 1, 0, 0, 0, 412, 414,
		1, 0, 0, 0, 413, 411, 1, 0, 0, 0, 414, 415, 5, 42, 0, 0, 415, 416, 5, 47,
		0, 0, 416, 417, 1, 0, 0, 0, 417, 418, 6, 52, 0, 0, 418, 106, 1, 0, 0, 0,
		419, 421, 7, 12, 0, 0, 420, 419, 1, 0, 0, 0, 421, 422, 1, 0, 0, 0, 422,
		420, 1, 0, 0, 0, 422, 423, 1, 0, 0, 0, 423, 424, 1, 0, 0, 0, 424, 425,
		6, 53, 0, 0, 425, 108, 1, 0, 0, 0, 426, 427, 5, 47, 0, 0, 427, 428, 5,
		47, 0, 0, 428, 432, 1, 0, 0, 0, 429, 431, 8, 12, 0, 0, 430, 429, 1, 0,
		0, 0, 431, 434, 1, 0, 0, 0, 432, 430, 1, 0, 0, 0, 432, 433, 1, 0, 0, 0,
		433, 435, 1, 0, 0, 0, 434, 432, 1, 0, 0, 0, 435, 436, 6, 54, 0, 0, 436,
		110, 1, 0, 0, 0, 437, 463, 5, 92, 0, 0, 438, 439, 5, 117, 0, 0, 439, 440,
		3, 117, 58, 0, 440, 441, 3, 117, 58, 0, 441, 442, 3, 117, 58, 0, 442, 443,
		3, 117, 58, 0, 443, 464, 1, 0, 0, 0, 444, 445, 5, 85, 0, 0, 445, 446, 3,
		117, 58, 0, 446, 447, 3, 117, 58, 0, 447, 448, 3, 117, 58, 0, 448, 449,
		3, 117, 58, 0, 449, 450, 3, 117, 58, 0, 450, 451, 3, 117, 58, 0, 451, 452,
		3, 117, 58, 0, 452, 453, 3, 117, 58, 0, 453, 464, 1, 0, 0, 0, 454, 464,
		7, 13, 0, 0, 455, 456, 3, 115, 57, 0, 456, 457, 3, 115, 57, 0, 457, 458,
		3, 115, 57, 0, 458, 464, 1, 0, 0, 0, 459, 460, 5, 120, 0, 0, 460, 461,
		3, 117, 58, 0, 461, 462, 3, 117, 58, 0, 462, 464, 1, 0, 0, 0, 463, 438,
		1, 0, 0, 0, 463, 444, 1, 0, 0, 0, 463, 454, 1, 0, 0, 0, 463, 455, 1, 0,
		0, 0, 463, 459, 1, 0, 0, 0, 464, 112, 1, 0, 0, 0, 465, 472, 7, 1, 0, 0,
		466, 468, 5, 95, 0, 0, 467, 466, 1, 0, 0, 0, 467, 468, 1, 0, 0, 0, 468,
		469, 1, 0, 0, 0, 469, 471, 7, 1, 0, 0, 470, 467, 1, 0, 0, 0, 471, 474,
		1, 0, 0, 0, 472, 470, 1, 0, 0, 0, 472, 473, 1, 0, 0, 0, 473, 114, 1, 0,
		0, 0, 474, 472, 1, 0, 0, 0, 475, 476, 7, 14, 0, 0, 476, 116, 1, 0, 0, 0,
		477, 478, 7, 15, 0, 0, 478, 118, 1, 0, 0, 0, 479, 480, 7, 16, 0, 0, 480,
		120, 1, 0, 0, 0, 481, 483, 7, 17, 0, 0, 482, 484, 7, 6, 0, 0, 483, 482,
		1, 0, 0, 0, 483, 484, 1, 0, 0, 0, 484, 485, 1, 0, 0, 0, 485, 486, 3, 113,
		56, 0, 486, 122, 1, 0, 0, 0, 487, 489, 7, 11, 0, 0, 488, 487, 1, 0, 0,
		0, 489, 490, 1, 0, 0, 0, 490, 488, 1, 0, 0, 0, 490, 491, 1, 0, 0, 0, 491,
		492, 1, 0, 0, 0, 492, 493, 6, 61, 0, 0, 493, 124, 1, 0, 0, 0, 494, 495,
		5, 47, 0, 0, 495, 496, 5, 42, 0, 0, 496, 500, 1, 0, 0, 0, 497, 499, 8,
		12, 0, 0, 498, 497, 1, 0, 0, 0, 499, 502, 1, 0, 0, 0, 500, 501, 1, 0, 0,
		0, 500, 498, 1, 0, 0, 0, 501, 503, 1, 0, 0, 0, 502, 500, 1, 0, 0, 0, 503,
		504, 5, 42, 0, 0, 504, 505, 5, 47, 0, 0, 505, 506, 1, 0, 0, 0, 506, 507,
		6, 62, 0, 0, 507, 126, 1, 0, 0, 0, 508, 509, 5, 47, 0, 0, 509, 510, 5,
		47, 0, 0, 510, 514, 1, 0, 0, 0, 511, 513, 8, 12, 0, 0, 512, 511, 1, 0,
		0, 0, 513, 516, 1, 0, 0, 0, 514, 512, 1, 0, 0, 0, 514, 515, 1, 0, 0, 0,
		515, 517, 1, 0, 0, 0, 516, 514, 1, 0, 0, 0, 517, 518, 6, 63, 0, 0, 518,
		128, 1, 0, 0, 0, 519, 521, 7, 12, 0, 0, 520, 519, 1, 0, 0, 0, 521, 522,
		1, 0, 0, 0, 522, 520, 1, 0, 0, 0, 522, 523, 1, 0, 0, 0, 523, 538, 1, 0,
		0, 0, 524, 538, 5, 59, 0, 0, 525, 526, 5, 47, 0, 0, 526, 527, 5, 42, 0,
		0, 527, 531, 1, 0, 0, 0, 528, 530, 9, 0, 0, 0, 529, 528, 1, 0, 0, 0, 530,
		533, 1, 0, 0, 0, 531, 532, 1, 0, 0, 0, 531, 529, 1, 0, 0, 0, 532, 534,
		1, 0, 0, 0, 533, 531, 1, 0, 0, 0, 534, 535, 5, 42, 0, 0, 535, 538, 5, 47,
		0, 0, 536, 538, 5, 0, 0, 1, 537, 520, 1, 0, 0, 0, 537, 524, 1, 0, 0, 0,
		537, 525, 1, 0, 0, 0, 537, 536, 1, 0, 0, 0, 538, 539, 1, 0, 0, 0, 539,
		540, 6, 64, 1, 0, 540, 130, 1, 0, 0, 0, 48, 0, 143, 160, 254, 259, 262,
		267, 272, 276, 279, 284, 289, 294, 298, 303, 306, 309, 314, 316, 324, 329,
		333, 338, 341, 346, 351, 354, 358, 366, 373, 381, 383, 391, 393, 401, 411,
		422, 432, 463, 467, 472, 483, 490, 500, 514, 522, 531, 537, 2, 0, 1, 0,
		2, 0, 0,
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

// goexpressionLexerInit initializes any static state used to implement goexpressionLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewgoexpressionLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoexpressionLexerInit() {
	staticData := &goexpressionlexerLexerStaticData
	staticData.once.Do(goexpressionlexerLexerInit)
}

// NewgoexpressionLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewgoexpressionLexer(input antlr.CharStream) *goexpressionLexer {
	GoexpressionLexerInit()
	l := new(goexpressionLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &goexpressionlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "goexpression.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// goexpressionLexer tokens.
const (
	goexpressionLexerEN_TRUE                = 1
	goexpressionLexerEN_FALSE               = 2
	goexpressionLexerEN_AND                 = 3
	goexpressionLexerEN_IN                  = 4
	goexpressionLexerEN_NOT                 = 5
	goexpressionLexerEN_OR                  = 6
	goexpressionLexerL_PAREN                = 7
	goexpressionLexerR_PAREN                = 8
	goexpressionLexerL_CURLY                = 9
	goexpressionLexerR_CURLY                = 10
	goexpressionLexerL_BRACKET              = 11
	goexpressionLexerR_BRACKET              = 12
	goexpressionLexerASSIGN                 = 13
	goexpressionLexerCOMMA                  = 14
	goexpressionLexerSEMI                   = 15
	goexpressionLexerCOLON                  = 16
	goexpressionLexerDOT                    = 17
	goexpressionLexerNIL_LIT                = 18
	goexpressionLexerLOGICAL_OR             = 19
	goexpressionLexerLOGICAL_AND            = 20
	goexpressionLexerEQUALS                 = 21
	goexpressionLexerNOT_EQUALS             = 22
	goexpressionLexerLESS                   = 23
	goexpressionLexerLESS_OR_EQUALS         = 24
	goexpressionLexerGREATER                = 25
	goexpressionLexerGREATER_OR_EQUALS      = 26
	goexpressionLexerOR                     = 27
	goexpressionLexerDIV                    = 28
	goexpressionLexerMOD                    = 29
	goexpressionLexerLSHIFT                 = 30
	goexpressionLexerRSHIFT                 = 31
	goexpressionLexerBIT_CLEAR              = 32
	goexpressionLexerEXCLAMATION            = 33
	goexpressionLexerPLUS                   = 34
	goexpressionLexerMINUS                  = 35
	goexpressionLexerCARET                  = 36
	goexpressionLexerSTAR                   = 37
	goexpressionLexerAMPERSAND              = 38
	goexpressionLexerDECIMAL_LIT            = 39
	goexpressionLexerBINARY_LIT             = 40
	goexpressionLexerOCTAL_LIT              = 41
	goexpressionLexerHEX_LIT                = 42
	goexpressionLexerFLOAT_LIT              = 43
	goexpressionLexerDECIMAL_FLOAT_LIT      = 44
	goexpressionLexerHEX_FLOAT_LIT          = 45
	goexpressionLexerIDENTIFIER             = 46
	goexpressionLexerRAW_STRING_LIT         = 47
	goexpressionLexerINTERPRETED_STRING_LIT = 48
	goexpressionLexerSINGLE_STRING_LIT      = 49
	goexpressionLexerWS                     = 50
	goexpressionLexerCOMMENT                = 51
	goexpressionLexerTERMINATOR             = 52
	goexpressionLexerLINE_COMMENT           = 53
	goexpressionLexerWS_NLSEMI              = 54
	goexpressionLexerCOMMENT_NLSEMI         = 55
	goexpressionLexerLINE_COMMENT_NLSEMI    = 56
	goexpressionLexerEOS                    = 57
)
