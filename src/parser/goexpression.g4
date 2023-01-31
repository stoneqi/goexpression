grammar goexpression;

expressionStmt: expression;
expression:
	primaryExpr
	| unary_op = (
		PLUS
		| MINUS
		| EXCLAMATION
        | EN_NOT
		| CARET
//		| STAR
//		| AMPERSAND
	) expression
	| expression mul_op = (
		STAR
		| DIV
		| MOD
		| LSHIFT
		| RSHIFT
		| AMPERSAND
		| BIT_CLEAR
	) expression
	| expression add_op = (PLUS | MINUS | OR | CARET) expression
	| expression rel_op = (
		EQUALS
		| NOT_EQUALS
		| LESS
		| LESS_OR_EQUALS
		| GREATER
		| GREATER_OR_EQUALS
	) expression
	| expression LOGICAL_AND expression
	| expression LOGICAL_OR expression
    | expression EN_AND expression
    | expression EN_OR expression
    | expression EN_IN expression
	| expression QUESTION expressionColon;

primaryExpr:
	operand
	| primaryExpr (
        (DOT IDENTIFIER)
		| index
		| slice_
		| arguments
	);



operand: basicLit | operandName | L_PAREN expression R_PAREN |  L_CURLY expressionList R_CURLY;

operandName: identifier;
slice_:
	L_BRACKET (
		expressionColon
	) R_BRACKET; 

expressionColon: expression COLON expression;	

index: L_BRACKET expression R_BRACKET;

// literal: basicLit;

basicLit:
	nil_lit
    | en_bool
	| integer
	| string_
	| float_lit;

nil_lit: NIL_LIT;
en_bool: EN_FALSE | EN_TRUE;
float_lit:FLOAT_LIT;
integer:
	DECIMAL_LIT
	| BINARY_LIT
	| OCTAL_LIT
	| HEX_LIT;

expressionList: expression (COMMA expression)*;
arguments:
	L_PAREN (
		expressionList
	)? R_PAREN;

eos:
	SEMI
	| EOF
	| EOS;

identifier: IDENTIFIER;
string_: RAW_STRING_LIT | INTERPRETED_STRING_LIT | SINGLE_STRING_LIT;

// bool
EN_TRUE:'True' 
    | 'true' 
    | 'TRUE' ;
EN_FALSE: 'False' 
    | 'false' 
    | 'FALSE';

//logic_en
EN_AND : 'and';
EN_IN : 'in';
EN_NOT : 'not';
EN_OR : 'or';


// Punctuation

L_PAREN                : '(';
R_PAREN                : ')';
L_CURLY                : '{';
R_CURLY                : '}';
L_BRACKET              : '[';
R_BRACKET              : ']';
ASSIGN                 : '=';
COMMA                  : ',';
SEMI                   : ';';
QUESTION               :'?';
COLON                  : ':';
DOT                    : '.';
NIL_LIT                : 'nil';

// Logical

LOGICAL_OR             : '||';
LOGICAL_AND            : '&&';

// Relation operators

EQUALS                 : '==';
NOT_EQUALS             : '!=';
LESS                   : '<';
LESS_OR_EQUALS         : '<=';
GREATER                : '>';
GREATER_OR_EQUALS      : '>=';

// Arithmetic operators

OR                     : '|';
DIV                    : '/';
MOD                    : '%';
LSHIFT                 : '<<';
RSHIFT                 : '>>';
BIT_CLEAR              : '&^';

// Unary operators

EXCLAMATION            : '!';

// Mixed operators

PLUS                   : '+';
MINUS                  : '-';
CARET                  : '^';
STAR                   : '*';
AMPERSAND              : '&';

// Number literals

DECIMAL_LIT            : ('0' | [1-9] ('_'? [0-9])*);
BINARY_LIT             : '0' [bB] ('_'? BIN_DIGIT)+;
OCTAL_LIT              : '0' [oO] ('_'? OCTAL_DIGIT)+;
HEX_LIT                : '0' [xX]  ('_'? HEX_DIGIT)+;


FLOAT_LIT : (DECIMAL_FLOAT_LIT | HEX_FLOAT_LIT);

DECIMAL_FLOAT_LIT      : DECIMALS ('.' DECIMALS? EXPONENT? | EXPONENT)
                       | '.' DECIMALS EXPONENT?
                       ;

HEX_FLOAT_LIT          : '0' [xX] HEX_MANTISSA HEX_EXPONENT
                       ;

fragment HEX_MANTISSA  : ('_'? HEX_DIGIT)+ ('.' ( '_'? HEX_DIGIT )*)?
                       | '.' HEX_DIGIT ('_'? HEX_DIGIT)*;

fragment HEX_EXPONENT  : [pP] [+-]? DECIMALS;

//name literals
IDENTIFIER
    : [a-zA-Z_]([a-zA-Z_0-9])*
    ;

// String literals

RAW_STRING_LIT         : '`' ~'`'*                      '`';
INTERPRETED_STRING_LIT : '"' (~["\\] | ESCAPED_VALUE)*  '"';
SINGLE_STRING_LIT : '\'' (~["\\] | ESCAPED_VALUE)*  '\'';

// Hidden tokens
WS                     : [ \t]+             -> channel(HIDDEN);
COMMENT                : '/*' .*? '*/'      -> channel(HIDDEN);
TERMINATOR             : [\r\n]+            -> channel(HIDDEN);
LINE_COMMENT           : '//' ~[\r\n]*      -> channel(HIDDEN);

// Fragments
fragment ESCAPED_VALUE
    : '\\' ('u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
           | 'U' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
           | [abfnrtv\\'"]
           | OCTAL_DIGIT OCTAL_DIGIT OCTAL_DIGIT
           | 'x' HEX_DIGIT HEX_DIGIT)
    ;
fragment DECIMALS
    : [0-9] ('_'? [0-9])*
    ;
fragment OCTAL_DIGIT
    : [0-7]
    ;
fragment HEX_DIGIT
    : [0-9a-fA-F]
    ;
fragment BIN_DIGIT
    : [01]
    ;
fragment EXPONENT
    : [eE] [+-]? DECIMALS
    ;
// Treat whitespace as normal
WS_NLSEMI                     : [ \t]+             -> channel(HIDDEN);
// Ignore any comments that only span one line
COMMENT_NLSEMI                : '/*' ~[\r\n]*? '*/'      -> channel(HIDDEN);
LINE_COMMENT_NLSEMI : '//' ~[\r\n]*      -> channel(HIDDEN);
// Emit an EOS token for any newlines, semicolon, multiline comments or the EOF and 
//return to normal lexing
EOS:              ([\r\n]+ | ';' | '/*' .*? '*/' | EOF)            -> mode(DEFAULT_MODE);
// Did not find an EOS, so go back to normal lexing