grammar BigDuck;

// Tokens

// Reserved Keywords
PROC    : 'proc';
RETURN  : 'return';
IF      : 'if';
ELSE    : 'else';
LOOP    : 'loop';
BREAK   : 'break';
SKIP    : 'skip';
AND     : 'and';
OR      : 'or';
NOT     : 'not';
VAR     : 'var';
INT     : 'int';
FLOAT   : 'float';
BOOL    : 'bool';
TRUE    : 'true';
FALSE   : 'false';

// Literals
fragment DIGIT  : [0-9];
fragment DIGITS : DIGIT+;
fragment LETTER : [A-Za-z]+;
fragment SIGN   : [+-]?;

CTE_INT     : SIGN DIGITS;
CTE_FLOAT   : DIGITS ('.' DIGITS)? ([Ee] SIGN? DIGITS)?;
CTE_STRING  : '"' ~('"')* '"';
ID          : LETTER (LETTER | DIGIT)*;
WS          : [ \n\t\r] -> skip;

// Syntax
program     : vars_decl procs_decl

vars_decl   : var_decl vars_decl
            | ;

var_decl    : VAR ID nextVar type nextTypes ';' nextVarDecl;
nextVar     : ',' ID nextVar
            | ;
nextTypes   : ',' ID nextVar type nextTypes
            | ;
nextVarDecl : var_decl nextVarDecl
            | ;

type        : scalar
            | tensor;

scalar      : INT
            | FLOAT
            | BOOL;

tensor      : dim scalar;
dim         : '[' num_expr ']';
