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
		"HEX_FLOAT_LIT", "IDENTIFIER", "INTERPRETED_STRING_LIT", "WS", "COMMENT",
		"TERMINATOR", "LINE_COMMENT", "WS_NLSEMI", "COMMENT_NLSEMI", "LINE_COMMENT_NLSEMI",
		"EOS",
	}
	staticData.ruleNames = []string{
		"EN_TRUE", "EN_FALSE", "EN_AND", "EN_IN", "EN_NOT", "EN_OR", "L_PAREN",
		"R_PAREN", "L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET", "ASSIGN",
		"COMMA", "SEMI", "QUESTION", "COLON", "DOT", "NIL_LIT", "LOGICAL_OR",
		"LOGICAL_AND", "EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS", "GREATER",
		"GREATER_OR_EQUALS", "OR", "DIV", "MOD", "LSHIFT", "RSHIFT", "BIT_CLEAR",
		"EXCLAMATION", "PLUS", "MINUS", "CARET", "STAR", "AMPERSAND", "DECIMAL_LIT",
		"BINARY_LIT", "OCTAL_LIT", "HEX_LIT", "FLOAT_LIT", "DECIMAL_FLOAT_LIT",
		"HEX_FLOAT_LIT", "HEX_MANTISSA", "HEX_EXPONENT", "IDENTIFIER", "INTERPRETED_STRING_LIT",
		"WS", "COMMENT", "TERMINATOR", "LINE_COMMENT", "ESCAPED_VALUE", "DECIMALS",
		"OCTAL_DIGIT", "HEX_DIGIT", "BIN_DIGIT", "EXPONENT", "WS_NLSEMI", "COMMENT_NLSEMI",
		"LINE_COMMENT_NLSEMI", "EOS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 56, 520, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		62, 2, 63, 7, 63, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 3, 0, 142, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 159, 8, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27,
		1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1,
		32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36,
		1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 3, 39, 255, 8, 39, 1,
		39, 5, 39, 258, 8, 39, 10, 39, 12, 39, 261, 9, 39, 3, 39, 263, 8, 39, 1,
		40, 1, 40, 1, 40, 3, 40, 268, 8, 40, 1, 40, 4, 40, 271, 8, 40, 11, 40,
		12, 40, 272, 1, 41, 1, 41, 1, 41, 3, 41, 278, 8, 41, 1, 41, 4, 41, 281,
		8, 41, 11, 41, 12, 41, 282, 1, 42, 1, 42, 1, 42, 3, 42, 288, 8, 42, 1,
		42, 4, 42, 291, 8, 42, 11, 42, 12, 42, 292, 1, 43, 1, 43, 3, 43, 297, 8,
		43, 1, 44, 1, 44, 1, 44, 3, 44, 302, 8, 44, 1, 44, 3, 44, 305, 8, 44, 1,
		44, 3, 44, 308, 8, 44, 1, 44, 1, 44, 1, 44, 3, 44, 313, 8, 44, 3, 44, 315,
		8, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 3, 46, 323, 8, 46, 1,
		46, 4, 46, 326, 8, 46, 11, 46, 12, 46, 327, 1, 46, 1, 46, 3, 46, 332, 8,
		46, 1, 46, 5, 46, 335, 8, 46, 10, 46, 12, 46, 338, 9, 46, 3, 46, 340, 8,
		46, 1, 46, 1, 46, 1, 46, 3, 46, 345, 8, 46, 1, 46, 5, 46, 348, 8, 46, 10,
		46, 12, 46, 351, 9, 46, 3, 46, 353, 8, 46, 1, 47, 1, 47, 3, 47, 357, 8,
		47, 1, 47, 1, 47, 1, 48, 1, 48, 5, 48, 363, 8, 48, 10, 48, 12, 48, 366,
		9, 48, 1, 49, 1, 49, 1, 49, 5, 49, 371, 8, 49, 10, 49, 12, 49, 374, 9,
		49, 1, 49, 1, 49, 1, 50, 4, 50, 379, 8, 50, 11, 50, 12, 50, 380, 1, 50,
		1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 5, 51, 389, 8, 51, 10, 51, 12, 51, 392,
		9, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 4, 52, 400, 8, 52, 11,
		52, 12, 52, 401, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 5, 53, 410,
		8, 53, 10, 53, 12, 53, 413, 9, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1,
		54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54,
		1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1,
		54, 1, 54, 3, 54, 443, 8, 54, 1, 55, 1, 55, 3, 55, 447, 8, 55, 1, 55, 5,
		55, 450, 8, 55, 10, 55, 12, 55, 453, 9, 55, 1, 56, 1, 56, 1, 57, 1, 57,
		1, 58, 1, 58, 1, 59, 1, 59, 3, 59, 463, 8, 59, 1, 59, 1, 59, 1, 60, 4,
		60, 468, 8, 60, 11, 60, 12, 60, 469, 1, 60, 1, 60, 1, 61, 1, 61, 1, 61,
		1, 61, 5, 61, 478, 8, 61, 10, 61, 12, 61, 481, 9, 61, 1, 61, 1, 61, 1,
		61, 1, 61, 1, 61, 1, 62, 1, 62, 1, 62, 1, 62, 5, 62, 492, 8, 62, 10, 62,
		12, 62, 495, 9, 62, 1, 62, 1, 62, 1, 63, 4, 63, 500, 8, 63, 11, 63, 12,
		63, 501, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 5, 63, 509, 8, 63, 10, 63,
		12, 63, 512, 9, 63, 1, 63, 1, 63, 1, 63, 3, 63, 517, 8, 63, 1, 63, 1, 63,
		3, 390, 479, 510, 0, 64, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35,
		71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44,
		89, 45, 91, 46, 93, 0, 95, 0, 97, 47, 99, 48, 101, 49, 103, 50, 105, 51,
		107, 52, 109, 0, 111, 0, 113, 0, 115, 0, 117, 0, 119, 0, 121, 53, 123,
		54, 125, 55, 127, 56, 1, 0, 17, 1, 0, 49, 57, 1, 0, 48, 57, 2, 0, 66, 66,
		98, 98, 2, 0, 79, 79, 111, 111, 2, 0, 88, 88, 120, 120, 2, 0, 80, 80, 112,
		112, 2, 0, 43, 43, 45, 45, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57,
		65, 90, 95, 95, 97, 122, 2, 0, 34, 34, 92, 92, 2, 0, 9, 9, 32, 32, 2, 0,
		10, 10, 13, 13, 9, 0, 34, 34, 39, 39, 92, 92, 97, 98, 102, 102, 110, 110,
		114, 114, 116, 116, 118, 118, 1, 0, 48, 55, 3, 0, 48, 57, 65, 70, 97, 102,
		1, 0, 48, 49, 2, 0, 69, 69, 101, 101, 561, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
		0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0,
		0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0,
		0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1,
		0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73,
		1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0,
		81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0,
		0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0,
		0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107,
		1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0,
		0, 127, 1, 0, 0, 0, 1, 141, 1, 0, 0, 0, 3, 158, 1, 0, 0, 0, 5, 160, 1,
		0, 0, 0, 7, 164, 1, 0, 0, 0, 9, 167, 1, 0, 0, 0, 11, 171, 1, 0, 0, 0, 13,
		174, 1, 0, 0, 0, 15, 176, 1, 0, 0, 0, 17, 178, 1, 0, 0, 0, 19, 180, 1,
		0, 0, 0, 21, 182, 1, 0, 0, 0, 23, 184, 1, 0, 0, 0, 25, 186, 1, 0, 0, 0,
		27, 188, 1, 0, 0, 0, 29, 190, 1, 0, 0, 0, 31, 192, 1, 0, 0, 0, 33, 194,
		1, 0, 0, 0, 35, 196, 1, 0, 0, 0, 37, 198, 1, 0, 0, 0, 39, 202, 1, 0, 0,
		0, 41, 205, 1, 0, 0, 0, 43, 208, 1, 0, 0, 0, 45, 211, 1, 0, 0, 0, 47, 214,
		1, 0, 0, 0, 49, 216, 1, 0, 0, 0, 51, 219, 1, 0, 0, 0, 53, 221, 1, 0, 0,
		0, 55, 224, 1, 0, 0, 0, 57, 226, 1, 0, 0, 0, 59, 228, 1, 0, 0, 0, 61, 230,
		1, 0, 0, 0, 63, 233, 1, 0, 0, 0, 65, 236, 1, 0, 0, 0, 67, 239, 1, 0, 0,
		0, 69, 241, 1, 0, 0, 0, 71, 243, 1, 0, 0, 0, 73, 245, 1, 0, 0, 0, 75, 247,
		1, 0, 0, 0, 77, 249, 1, 0, 0, 0, 79, 262, 1, 0, 0, 0, 81, 264, 1, 0, 0,
		0, 83, 274, 1, 0, 0, 0, 85, 284, 1, 0, 0, 0, 87, 296, 1, 0, 0, 0, 89, 314,
		1, 0, 0, 0, 91, 316, 1, 0, 0, 0, 93, 352, 1, 0, 0, 0, 95, 354, 1, 0, 0,
		0, 97, 360, 1, 0, 0, 0, 99, 367, 1, 0, 0, 0, 101, 378, 1, 0, 0, 0, 103,
		384, 1, 0, 0, 0, 105, 399, 1, 0, 0, 0, 107, 405, 1, 0, 0, 0, 109, 416,
		1, 0, 0, 0, 111, 444, 1, 0, 0, 0, 113, 454, 1, 0, 0, 0, 115, 456, 1, 0,
		0, 0, 117, 458, 1, 0, 0, 0, 119, 460, 1, 0, 0, 0, 121, 467, 1, 0, 0, 0,
		123, 473, 1, 0, 0, 0, 125, 487, 1, 0, 0, 0, 127, 516, 1, 0, 0, 0, 129,
		130, 5, 84, 0, 0, 130, 131, 5, 114, 0, 0, 131, 132, 5, 117, 0, 0, 132,
		142, 5, 101, 0, 0, 133, 134, 5, 116, 0, 0, 134, 135, 5, 114, 0, 0, 135,
		136, 5, 117, 0, 0, 136, 142, 5, 101, 0, 0, 137, 138, 5, 84, 0, 0, 138,
		139, 5, 82, 0, 0, 139, 140, 5, 85, 0, 0, 140, 142, 5, 69, 0, 0, 141, 129,
		1, 0, 0, 0, 141, 133, 1, 0, 0, 0, 141, 137, 1, 0, 0, 0, 142, 2, 1, 0, 0,
		0, 143, 144, 5, 70, 0, 0, 144, 145, 5, 97, 0, 0, 145, 146, 5, 108, 0, 0,
		146, 147, 5, 115, 0, 0, 147, 159, 5, 101, 0, 0, 148, 149, 5, 102, 0, 0,
		149, 150, 5, 97, 0, 0, 150, 151, 5, 108, 0, 0, 151, 152, 5, 115, 0, 0,
		152, 159, 5, 101, 0, 0, 153, 154, 5, 70, 0, 0, 154, 155, 5, 65, 0, 0, 155,
		156, 5, 76, 0, 0, 156, 157, 5, 83, 0, 0, 157, 159, 5, 69, 0, 0, 158, 143,
		1, 0, 0, 0, 158, 148, 1, 0, 0, 0, 158, 153, 1, 0, 0, 0, 159, 4, 1, 0, 0,
		0, 160, 161, 5, 97, 0, 0, 161, 162, 5, 110, 0, 0, 162, 163, 5, 100, 0,
		0, 163, 6, 1, 0, 0, 0, 164, 165, 5, 105, 0, 0, 165, 166, 5, 110, 0, 0,
		166, 8, 1, 0, 0, 0, 167, 168, 5, 110, 0, 0, 168, 169, 5, 111, 0, 0, 169,
		170, 5, 116, 0, 0, 170, 10, 1, 0, 0, 0, 171, 172, 5, 111, 0, 0, 172, 173,
		5, 114, 0, 0, 173, 12, 1, 0, 0, 0, 174, 175, 5, 40, 0, 0, 175, 14, 1, 0,
		0, 0, 176, 177, 5, 41, 0, 0, 177, 16, 1, 0, 0, 0, 178, 179, 5, 123, 0,
		0, 179, 18, 1, 0, 0, 0, 180, 181, 5, 125, 0, 0, 181, 20, 1, 0, 0, 0, 182,
		183, 5, 91, 0, 0, 183, 22, 1, 0, 0, 0, 184, 185, 5, 93, 0, 0, 185, 24,
		1, 0, 0, 0, 186, 187, 5, 61, 0, 0, 187, 26, 1, 0, 0, 0, 188, 189, 5, 44,
		0, 0, 189, 28, 1, 0, 0, 0, 190, 191, 5, 59, 0, 0, 191, 30, 1, 0, 0, 0,
		192, 193, 5, 63, 0, 0, 193, 32, 1, 0, 0, 0, 194, 195, 5, 58, 0, 0, 195,
		34, 1, 0, 0, 0, 196, 197, 5, 46, 0, 0, 197, 36, 1, 0, 0, 0, 198, 199, 5,
		110, 0, 0, 199, 200, 5, 105, 0, 0, 200, 201, 5, 108, 0, 0, 201, 38, 1,
		0, 0, 0, 202, 203, 5, 124, 0, 0, 203, 204, 5, 124, 0, 0, 204, 40, 1, 0,
		0, 0, 205, 206, 5, 38, 0, 0, 206, 207, 5, 38, 0, 0, 207, 42, 1, 0, 0, 0,
		208, 209, 5, 61, 0, 0, 209, 210, 5, 61, 0, 0, 210, 44, 1, 0, 0, 0, 211,
		212, 5, 33, 0, 0, 212, 213, 5, 61, 0, 0, 213, 46, 1, 0, 0, 0, 214, 215,
		5, 60, 0, 0, 215, 48, 1, 0, 0, 0, 216, 217, 5, 60, 0, 0, 217, 218, 5, 61,
		0, 0, 218, 50, 1, 0, 0, 0, 219, 220, 5, 62, 0, 0, 220, 52, 1, 0, 0, 0,
		221, 222, 5, 62, 0, 0, 222, 223, 5, 61, 0, 0, 223, 54, 1, 0, 0, 0, 224,
		225, 5, 124, 0, 0, 225, 56, 1, 0, 0, 0, 226, 227, 5, 47, 0, 0, 227, 58,
		1, 0, 0, 0, 228, 229, 5, 37, 0, 0, 229, 60, 1, 0, 0, 0, 230, 231, 5, 60,
		0, 0, 231, 232, 5, 60, 0, 0, 232, 62, 1, 0, 0, 0, 233, 234, 5, 62, 0, 0,
		234, 235, 5, 62, 0, 0, 235, 64, 1, 0, 0, 0, 236, 237, 5, 38, 0, 0, 237,
		238, 5, 94, 0, 0, 238, 66, 1, 0, 0, 0, 239, 240, 5, 33, 0, 0, 240, 68,
		1, 0, 0, 0, 241, 242, 5, 43, 0, 0, 242, 70, 1, 0, 0, 0, 243, 244, 5, 45,
		0, 0, 244, 72, 1, 0, 0, 0, 245, 246, 5, 94, 0, 0, 246, 74, 1, 0, 0, 0,
		247, 248, 5, 42, 0, 0, 248, 76, 1, 0, 0, 0, 249, 250, 5, 38, 0, 0, 250,
		78, 1, 0, 0, 0, 251, 263, 5, 48, 0, 0, 252, 259, 7, 0, 0, 0, 253, 255,
		5, 95, 0, 0, 254, 253, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 256, 1, 0,
		0, 0, 256, 258, 7, 1, 0, 0, 257, 254, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0,
		259, 257, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 263, 1, 0, 0, 0, 261,
		259, 1, 0, 0, 0, 262, 251, 1, 0, 0, 0, 262, 252, 1, 0, 0, 0, 263, 80, 1,
		0, 0, 0, 264, 265, 5, 48, 0, 0, 265, 270, 7, 2, 0, 0, 266, 268, 5, 95,
		0, 0, 267, 266, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0,
		269, 271, 3, 117, 58, 0, 270, 267, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272,
		270, 1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273, 82, 1, 0, 0, 0, 274, 275, 5,
		48, 0, 0, 275, 280, 7, 3, 0, 0, 276, 278, 5, 95, 0, 0, 277, 276, 1, 0,
		0, 0, 277, 278, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 281, 3, 113, 56,
		0, 280, 277, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 280, 1, 0, 0, 0, 282,
		283, 1, 0, 0, 0, 283, 84, 1, 0, 0, 0, 284, 285, 5, 48, 0, 0, 285, 290,
		7, 4, 0, 0, 286, 288, 5, 95, 0, 0, 287, 286, 1, 0, 0, 0, 287, 288, 1, 0,
		0, 0, 288, 289, 1, 0, 0, 0, 289, 291, 3, 115, 57, 0, 290, 287, 1, 0, 0,
		0, 291, 292, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293,
		86, 1, 0, 0, 0, 294, 297, 3, 89, 44, 0, 295, 297, 3, 91, 45, 0, 296, 294,
		1, 0, 0, 0, 296, 295, 1, 0, 0, 0, 297, 88, 1, 0, 0, 0, 298, 307, 3, 111,
		55, 0, 299, 301, 5, 46, 0, 0, 300, 302, 3, 111, 55, 0, 301, 300, 1, 0,
		0, 0, 301, 302, 1, 0, 0, 0, 302, 304, 1, 0, 0, 0, 303, 305, 3, 119, 59,
		0, 304, 303, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 308, 1, 0, 0, 0, 306,
		308, 3, 119, 59, 0, 307, 299, 1, 0, 0, 0, 307, 306, 1, 0, 0, 0, 308, 315,
		1, 0, 0, 0, 309, 310, 5, 46, 0, 0, 310, 312, 3, 111, 55, 0, 311, 313, 3,
		119, 59, 0, 312, 311, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 315, 1, 0,
		0, 0, 314, 298, 1, 0, 0, 0, 314, 309, 1, 0, 0, 0, 315, 90, 1, 0, 0, 0,
		316, 317, 5, 48, 0, 0, 317, 318, 7, 4, 0, 0, 318, 319, 3, 93, 46, 0, 319,
		320, 3, 95, 47, 0, 320, 92, 1, 0, 0, 0, 321, 323, 5, 95, 0, 0, 322, 321,
		1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 326, 3, 115,
		57, 0, 325, 322, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 325, 1, 0, 0, 0,
		327, 328, 1, 0, 0, 0, 328, 339, 1, 0, 0, 0, 329, 336, 5, 46, 0, 0, 330,
		332, 5, 95, 0, 0, 331, 330, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 333,
		1, 0, 0, 0, 333, 335, 3, 115, 57, 0, 334, 331, 1, 0, 0, 0, 335, 338, 1,
		0, 0, 0, 336, 334, 1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 340, 1, 0, 0,
		0, 338, 336, 1, 0, 0, 0, 339, 329, 1, 0, 0, 0, 339, 340, 1, 0, 0, 0, 340,
		353, 1, 0, 0, 0, 341, 342, 5, 46, 0, 0, 342, 349, 3, 115, 57, 0, 343, 345,
		5, 95, 0, 0, 344, 343, 1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345, 346, 1, 0,
		0, 0, 346, 348, 3, 115, 57, 0, 347, 344, 1, 0, 0, 0, 348, 351, 1, 0, 0,
		0, 349, 347, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 353, 1, 0, 0, 0, 351,
		349, 1, 0, 0, 0, 352, 325, 1, 0, 0, 0, 352, 341, 1, 0, 0, 0, 353, 94, 1,
		0, 0, 0, 354, 356, 7, 5, 0, 0, 355, 357, 7, 6, 0, 0, 356, 355, 1, 0, 0,
		0, 356, 357, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358, 359, 3, 111, 55, 0,
		359, 96, 1, 0, 0, 0, 360, 364, 7, 7, 0, 0, 361, 363, 7, 8, 0, 0, 362, 361,
		1, 0, 0, 0, 363, 366, 1, 0, 0, 0, 364, 362, 1, 0, 0, 0, 364, 365, 1, 0,
		0, 0, 365, 98, 1, 0, 0, 0, 366, 364, 1, 0, 0, 0, 367, 372, 5, 34, 0, 0,
		368, 371, 8, 9, 0, 0, 369, 371, 3, 109, 54, 0, 370, 368, 1, 0, 0, 0, 370,
		369, 1, 0, 0, 0, 371, 374, 1, 0, 0, 0, 372, 370, 1, 0, 0, 0, 372, 373,
		1, 0, 0, 0, 373, 375, 1, 0, 0, 0, 374, 372, 1, 0, 0, 0, 375, 376, 5, 34,
		0, 0, 376, 100, 1, 0, 0, 0, 377, 379, 7, 10, 0, 0, 378, 377, 1, 0, 0, 0,
		379, 380, 1, 0, 0, 0, 380, 378, 1, 0, 0, 0, 380, 381, 1, 0, 0, 0, 381,
		382, 1, 0, 0, 0, 382, 383, 6, 50, 0, 0, 383, 102, 1, 0, 0, 0, 384, 385,
		5, 47, 0, 0, 385, 386, 5, 42, 0, 0, 386, 390, 1, 0, 0, 0, 387, 389, 9,
		0, 0, 0, 388, 387, 1, 0, 0, 0, 389, 392, 1, 0, 0, 0, 390, 391, 1, 0, 0,
		0, 390, 388, 1, 0, 0, 0, 391, 393, 1, 0, 0, 0, 392, 390, 1, 0, 0, 0, 393,
		394, 5, 42, 0, 0, 394, 395, 5, 47, 0, 0, 395, 396, 1, 0, 0, 0, 396, 397,
		6, 51, 0, 0, 397, 104, 1, 0, 0, 0, 398, 400, 7, 11, 0, 0, 399, 398, 1,
		0, 0, 0, 400, 401, 1, 0, 0, 0, 401, 399, 1, 0, 0, 0, 401, 402, 1, 0, 0,
		0, 402, 403, 1, 0, 0, 0, 403, 404, 6, 52, 0, 0, 404, 106, 1, 0, 0, 0, 405,
		406, 5, 47, 0, 0, 406, 407, 5, 47, 0, 0, 407, 411, 1, 0, 0, 0, 408, 410,
		8, 11, 0, 0, 409, 408, 1, 0, 0, 0, 410, 413, 1, 0, 0, 0, 411, 409, 1, 0,
		0, 0, 411, 412, 1, 0, 0, 0, 412, 414, 1, 0, 0, 0, 413, 411, 1, 0, 0, 0,
		414, 415, 6, 53, 0, 0, 415, 108, 1, 0, 0, 0, 416, 442, 5, 92, 0, 0, 417,
		418, 5, 117, 0, 0, 418, 419, 3, 115, 57, 0, 419, 420, 3, 115, 57, 0, 420,
		421, 3, 115, 57, 0, 421, 422, 3, 115, 57, 0, 422, 443, 1, 0, 0, 0, 423,
		424, 5, 85, 0, 0, 424, 425, 3, 115, 57, 0, 425, 426, 3, 115, 57, 0, 426,
		427, 3, 115, 57, 0, 427, 428, 3, 115, 57, 0, 428, 429, 3, 115, 57, 0, 429,
		430, 3, 115, 57, 0, 430, 431, 3, 115, 57, 0, 431, 432, 3, 115, 57, 0, 432,
		443, 1, 0, 0, 0, 433, 443, 7, 12, 0, 0, 434, 435, 3, 113, 56, 0, 435, 436,
		3, 113, 56, 0, 436, 437, 3, 113, 56, 0, 437, 443, 1, 0, 0, 0, 438, 439,
		5, 120, 0, 0, 439, 440, 3, 115, 57, 0, 440, 441, 3, 115, 57, 0, 441, 443,
		1, 0, 0, 0, 442, 417, 1, 0, 0, 0, 442, 423, 1, 0, 0, 0, 442, 433, 1, 0,
		0, 0, 442, 434, 1, 0, 0, 0, 442, 438, 1, 0, 0, 0, 443, 110, 1, 0, 0, 0,
		444, 451, 7, 1, 0, 0, 445, 447, 5, 95, 0, 0, 446, 445, 1, 0, 0, 0, 446,
		447, 1, 0, 0, 0, 447, 448, 1, 0, 0, 0, 448, 450, 7, 1, 0, 0, 449, 446,
		1, 0, 0, 0, 450, 453, 1, 0, 0, 0, 451, 449, 1, 0, 0, 0, 451, 452, 1, 0,
		0, 0, 452, 112, 1, 0, 0, 0, 453, 451, 1, 0, 0, 0, 454, 455, 7, 13, 0, 0,
		455, 114, 1, 0, 0, 0, 456, 457, 7, 14, 0, 0, 457, 116, 1, 0, 0, 0, 458,
		459, 7, 15, 0, 0, 459, 118, 1, 0, 0, 0, 460, 462, 7, 16, 0, 0, 461, 463,
		7, 6, 0, 0, 462, 461, 1, 0, 0, 0, 462, 463, 1, 0, 0, 0, 463, 464, 1, 0,
		0, 0, 464, 465, 3, 111, 55, 0, 465, 120, 1, 0, 0, 0, 466, 468, 7, 10, 0,
		0, 467, 466, 1, 0, 0, 0, 468, 469, 1, 0, 0, 0, 469, 467, 1, 0, 0, 0, 469,
		470, 1, 0, 0, 0, 470, 471, 1, 0, 0, 0, 471, 472, 6, 60, 0, 0, 472, 122,
		1, 0, 0, 0, 473, 474, 5, 47, 0, 0, 474, 475, 5, 42, 0, 0, 475, 479, 1,
		0, 0, 0, 476, 478, 8, 11, 0, 0, 477, 476, 1, 0, 0, 0, 478, 481, 1, 0, 0,
		0, 479, 480, 1, 0, 0, 0, 479, 477, 1, 0, 0, 0, 480, 482, 1, 0, 0, 0, 481,
		479, 1, 0, 0, 0, 482, 483, 5, 42, 0, 0, 483, 484, 5, 47, 0, 0, 484, 485,
		1, 0, 0, 0, 485, 486, 6, 61, 0, 0, 486, 124, 1, 0, 0, 0, 487, 488, 5, 47,
		0, 0, 488, 489, 5, 47, 0, 0, 489, 493, 1, 0, 0, 0, 490, 492, 8, 11, 0,
		0, 491, 490, 1, 0, 0, 0, 492, 495, 1, 0, 0, 0, 493, 491, 1, 0, 0, 0, 493,
		494, 1, 0, 0, 0, 494, 496, 1, 0, 0, 0, 495, 493, 1, 0, 0, 0, 496, 497,
		6, 62, 0, 0, 497, 126, 1, 0, 0, 0, 498, 500, 7, 11, 0, 0, 499, 498, 1,
		0, 0, 0, 500, 501, 1, 0, 0, 0, 501, 499, 1, 0, 0, 0, 501, 502, 1, 0, 0,
		0, 502, 517, 1, 0, 0, 0, 503, 517, 5, 59, 0, 0, 504, 505, 5, 47, 0, 0,
		505, 506, 5, 42, 0, 0, 506, 510, 1, 0, 0, 0, 507, 509, 9, 0, 0, 0, 508,
		507, 1, 0, 0, 0, 509, 512, 1, 0, 0, 0, 510, 511, 1, 0, 0, 0, 510, 508,
		1, 0, 0, 0, 511, 513, 1, 0, 0, 0, 512, 510, 1, 0, 0, 0, 513, 514, 5, 42,
		0, 0, 514, 517, 5, 47, 0, 0, 515, 517, 5, 0, 0, 1, 516, 499, 1, 0, 0, 0,
		516, 503, 1, 0, 0, 0, 516, 504, 1, 0, 0, 0, 516, 515, 1, 0, 0, 0, 517,
		518, 1, 0, 0, 0, 518, 519, 6, 63, 1, 0, 519, 128, 1, 0, 0, 0, 44, 0, 141,
		158, 254, 259, 262, 267, 272, 277, 282, 287, 292, 296, 301, 304, 307, 312,
		314, 322, 327, 331, 336, 339, 344, 349, 352, 356, 364, 370, 372, 380, 390,
		401, 411, 442, 446, 451, 462, 469, 479, 493, 501, 510, 516, 2, 0, 1, 0,
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
	goexpressionLexerQUESTION               = 16
	goexpressionLexerCOLON                  = 17
	goexpressionLexerDOT                    = 18
	goexpressionLexerNIL_LIT                = 19
	goexpressionLexerLOGICAL_OR             = 20
	goexpressionLexerLOGICAL_AND            = 21
	goexpressionLexerEQUALS                 = 22
	goexpressionLexerNOT_EQUALS             = 23
	goexpressionLexerLESS                   = 24
	goexpressionLexerLESS_OR_EQUALS         = 25
	goexpressionLexerGREATER                = 26
	goexpressionLexerGREATER_OR_EQUALS      = 27
	goexpressionLexerOR                     = 28
	goexpressionLexerDIV                    = 29
	goexpressionLexerMOD                    = 30
	goexpressionLexerLSHIFT                 = 31
	goexpressionLexerRSHIFT                 = 32
	goexpressionLexerBIT_CLEAR              = 33
	goexpressionLexerEXCLAMATION            = 34
	goexpressionLexerPLUS                   = 35
	goexpressionLexerMINUS                  = 36
	goexpressionLexerCARET                  = 37
	goexpressionLexerSTAR                   = 38
	goexpressionLexerAMPERSAND              = 39
	goexpressionLexerDECIMAL_LIT            = 40
	goexpressionLexerBINARY_LIT             = 41
	goexpressionLexerOCTAL_LIT              = 42
	goexpressionLexerHEX_LIT                = 43
	goexpressionLexerFLOAT_LIT              = 44
	goexpressionLexerDECIMAL_FLOAT_LIT      = 45
	goexpressionLexerHEX_FLOAT_LIT          = 46
	goexpressionLexerIDENTIFIER             = 47
	goexpressionLexerINTERPRETED_STRING_LIT = 48
	goexpressionLexerWS                     = 49
	goexpressionLexerCOMMENT                = 50
	goexpressionLexerTERMINATOR             = 51
	goexpressionLexerLINE_COMMENT           = 52
	goexpressionLexerWS_NLSEMI              = 53
	goexpressionLexerCOMMENT_NLSEMI         = 54
	goexpressionLexerLINE_COMMENT_NLSEMI    = 55
	goexpressionLexerEOS                    = 56
)
