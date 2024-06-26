parser grammar JavaScriptParser;

options { tokenVocab=JavaScriptLexer; }

num : NUMBER;
id : IDENTIFIER;
str: STRING;
arr: '[' expr (',' expr)* ']';
key : id ':' expr;
obj: '{' key (',' key)* '}';
bool: 'true' | 'false';
this: 'this';

funcdef: 'function' IDENTIFIER '(' paramlist? ')' block;
paramlist: param (',' param)*;
param: IDENTIFIER ('=' expr)?;

funcall: IDENTIFIER '(' exprlist? ')';
exprlist: expr (',' expr)*;
cobj : 'new' IDENTIFIER '(' exprlist? ')';

class: 'class' IDENTIFIER '{' cons? method* '}';
cons: 'constructor' '(' paramlist? ')' block;
method: IDENTIFIER '(' paramlist? ')' block;

program : (global)* ;

global: funcdef
    | stmg
    | class
    ;

expr: '(' expr ')'
    | expr '.' expr
    | expr '[' expr ']'
    | expr ('*' | '/' | '+' | '-') expr
    | expr ('<' | '>' | '==' | '<=' | '>=' | '!=' | '===' | '!==') expr
    | expr ('&&' | '||') expr
    | ('++' | '--') expr
    | expr('++' | '--')
    | ('!' | '-') expr
    | this
    | cobj | funcall | num | id | str | arr | obj
    ;

stmg: stm ';'?;
stm: expr ('=' | '+=' | '-=' | '*=' | '/=') expr
    | 'var' IDENTIFIER ('=' expr)?
    | 'return' expr
    | 'break'
    | 'continue'
    | 'if' '(' expr ')' block ('else' block)?
    | 'while' '(' expr ')' block
    | 'for' '(' stm? ';' stm? ';' stm? ')' block
    | 'switch' '(' expr ')' '{' case*  default? '}'
    | block
    | expr
    ;
    
case : 'case' expr ':' (stm ';'?)*;
default: 'default' ':' (stm ';'?)*;
block: '{' (stm ';'?)* '}';