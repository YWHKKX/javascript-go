lexer grammar JavaScriptLexer;

FUNCTION : 'function';
RETURN : 'return';
VAR : 'var';
CLASS : 'class';
THIS : 'this';
CONS : 'constructor';
NEW : 'new';

NULL: 'null';
UNDEFINED: 'undefined';

IF : 'if';
ELSE : 'else';
WHILE : 'while';
FOR : 'for';
BREAK : 'break';
CONTINUE : 'continue';
CASE : 'case';
SWITCH : 'switch';
DEFAULT : 'default';

TRUE: 'true';
FALSE: 'false';

LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
LBRACKET: '[';
RBRACKET: ']';

NUMBER: [0-9]+('.'[0-9]+)?;
IDENTIFIER: [a-zA-Z] [a-zA-Z0-9]*;
STRING: '"' (~["\r\n] | '\\"')* '"';

COL : ':';
DOT : '.';
COMMA : ',';
SEMICOLON : ';';

ASSIGN : '=';
ADD: '+';
SUB: '-';
MUL: '*';
DIV: '/';
MOD: '%';

INC: '++';
DEC: '--';

ADD_ASSIGN: '+=';
SUB_ASSIGN: '-=';
MUL_ASSIGN: '*=';
DIV_ASSIGN: '/=';
MOD_ASSIGN: '%=';

NOT: '!';
EQ: '==';
NEQ: '!=';
LT: '<';
GT: '>';
LTE: '<=';
GTE: '>=';

AEQ: '===';
NAEQ: '!==';

AND: '&&';
OR: '||';

WS: [ \t\r\n]+ -> skip;
COMMENT: '/*' .*? '*/' -> skip;
LINE_COMMENT: '//' ~[\r\n]* -> skip;
