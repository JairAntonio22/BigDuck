grammar BigDuck;

// Tokens

// Reserved Keywords
PROC    : 'proc';
RETURN  : 'return';
IF      : 'if';
ELSE    : 'else';
LOOP    : 'loop';
BREAK   : 'break';
SKIP_W  : 'skip';
AND     : 'and';
OR      : 'or';
NOT     : 'not';
VAR     : 'var';
INT     : 'int';
FLOAT   : 'float';
BOOL    : 'bool';
TRUE    : 'true';
FALSE   : 'false';

// Built-in procedures
PRINT_W : 'print';

// Literals
fragment DIGIT  : [0-9];
fragment DIGITS : DIGIT+;
fragment LETTER : [A-Za-z]+;
fragment SIGN   : '-';

CTE_INT     : SIGN? DIGITS;
CTE_FLOAT   : SIGN? DIGITS ('.' DIGITS)? ([Ee] SIGN? DIGITS)?;
CTE_STRING  : '"' ~('"')* '"';
ID          : LETTER (LETTER | DIGIT | '_')*;
WS          : [ \n\t\r] -> skip;
COMMENT     : '#|' .*? '|#' -> skip;

// Syntax
program     : vars_decl procs_decl;

vars_decl   : var_decl vars_decl
            | ;

var_decl    : VAR ID nextVar var_type ';' nextVarDecl;
nextVar     : ',' ID nextVar
            | ;
nextVarDecl : var_decl nextVarDecl
            | ;

var_type    : scalar | tensor;

scalar      : INT | FLOAT | BOOL;

tensor      : dim scalar;
dim         : '[' num_expr ']' rbracket /*nextDim*/;
rbracket    : ;
nextDim     : dim
            | ;

procs_decl  : proc_decl (procs_decl | );

proc_decl   : sign (ret_type | ) proc_info (var_decl | ) block;
proc_info   : ;

sign        : PROC ID args;

args        : '(' (ID nextArg scalar nextTypes | ) ')';
nextTypes   : ';' ID nextArg scalar nextTypes
            | ;
nextArg     : ',' ID nextArg
            | ;

ret_type    : '->' scalar;


bool_expr   : and_expr nextBool;
nextBool    : OR bool_expr
            | ;

and_expr    : not_expr nextAnd;
nextAnd     : AND and_expr
            | ;

not_expr    : (NOT | ) bool_term;
bool_term   : '(' bool_expr ')'
            | rel_expr
            | TRUE
            | FALSE
            | variable
            | proc_call;

rel_expr    : num_expr opRel;
opRel       : relOp num_expr;
relOp       : '='
            | '/='
            | '<'
            | '>'
            | '>='
            | '<=';

num_expr    : prod_expr nextSum;
nextSum     : ('+' | '-') num_expr
            | ;

prod_expr   : factor nextProd;
nextProd    : ('*' | '/') prod_expr
            | ;

factor      : '(' num_expr ')'
            | variable
            | CTE_INT
            | CTE_FLOAT
            | proc_call;

variable    : ID (t_access dim t_end| );
t_access    : ;
t_end       : ;

proc_call   : ID '(' (param | ) ')';
param       : paramTerm nextParam;
paramTerm   : bool_expr
            | num_expr;
nextParam   : ',' param
            | ;

block       : '{' stmts '}';
stmts       : stmt stmts
            | ;
stmt        : assignment ';'
            | condition
            | loop_stmt
            | ctrl_flow ';'
            | ret_stmt ';'
            | print_r ';'
            | void_proc ';';

void_proc   : proc_call;

assignment  : variable '<-' (num_expr | bool_expr);

condition   : IF bool_expr bodyCond;
bodyCond    : block endIfBlock;
endIfBlock  : (alter | );

alter       : ELSE (condition | block);

loop_stmt   : LOOP (forStyle | whileStyle | infLoop) block;

forStyle    : (assignment | ) ';' forCond ';' ctrl_var;
forCond     : bool_expr;
ctrl_var    : assignment;

whileStyle  : bool_expr;

infLoop     : ;

ctrl_flow   : (BREAK | SKIP_W);

ret_stmt    : RETURN (num_expr | bool_expr | proc_call | );

// Built-in procedures

print_r     : PRINT_W '(' pparam ')';
pparam      : pparamTerm pnextParam;
pparamTerm  : bool_expr
            | num_expr
            | CTE_STRING;
pnextParam  : ',' pparam
            | ;
