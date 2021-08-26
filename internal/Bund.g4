grammar Bund;

expressions
  : root_term*
;

root_term
   : ( ns
     | block
   )
;

ns
  : '[' name=NAME ':' (body+=term)* ';;'
  ;

block
  : '(' (body+=term)* ')'
  ;

term
  : ( ns
    | block
    | call_term
    | operator_term
    | ref_call_term
    | ref_operator_term
    | boolean_term
    | integer_term
    | float_term
    | string_term
    | complex_term
    | glob_term
    | datablock_term
    | matchblock_term
    | logicblock_term
    | mode_term
    | separate_term
    | function_term
    | lambda_term
    | operation_term
    | file_term
    | unixcmd_term
    | json_term
  );

data
  : ( boolean_term
    | integer_term
    | float_term
    | string_term
    | complex_term
    | call_term
    | operator_term
    | separate_term
    | glob_term
    | lambda_term
    | ref_call_term
  );

call_term
  : VALUE=(SYSF|NAME) (':(' FUNCTOR=(SYS|SYSF|NAME) ')')?
  ;

operator_term
  : VALUE=CMD (':(' FUNCTOR=(SYS|SYSF|NAME) ')')?
  ;

ref_call_term:     '`' VALUE=(SYSF|NAME) (':(' FUNCTOR=(SYS|SYSF|NAME) ')')? ;
ref_operator_term: '`' VALUE=CMD (':(' FUNCTOR=(SYS|SYSF|NAME) ')')? ;

boolean_term: VALUE=(TRUE|FALSE)(':(' FUNCTOR=(SYSF|NAME) ')')? ;
integer_term: VALUE=INTEGER(':(' FUNCTOR=(SYSF|NAME) ')')? ;
float_term:   VALUE=(FLOAT_NUMBER|'+Inf'|'NaN'|'-Inf'|'Inf')(':(' FUNCTOR=(SYSF|NAME) ')')? ;
string_term:  (PRE=NAME '@')? VALUE=STRING(':(' FUNCTOR=(SYSF|NAME) ')')? ;
complex_term: VALUE=COMPLEX_NUMBER(':(' FUNCTOR=(SYSF|NAME) ')')? ;
glob_term:    VALUE=GLOB(':(' FUNCTOR=(SYSF|NAME) ')')? ;
file_term:    VALUE=URI(':(' FUNCTOR=(SYSF|NAME) ')')? ;
unixcmd_term: VALUE=UNIXCMD(':(' FUNCTOR=(SYSF|NAME) ')')? ;
json_term:    VALUE=JSON(':(' FUNCTOR=(SYSF|NAME) ')')? ;


mode_term:    VALUE=(TOBEGIN|TOEND) ;

separate_term: VALUE=SEPARATE ;

datablock_term
  : '(*' (body+=data)* ')'(':(' FUNCTOR=(SYS|NAME) ')')?
  ;

matchblock_term
  : '(?' (body+=data)+ ')'
  ;

logicblock_term
  : HDR=('(true'|'(false') (body+=term)* ')'
  ;

function_term
  : '[' name=NAME ']' (body+=term)* '.'
  ;

lambda_term
  : '[]' (body+=term)* '.'
  ;


operation_term
  : '[[' name=CMD ']]' (body+=term)* '.'
  ;


INTEGER
  :  (SIGN)? DECIMAL_INTEGER
  ;

DECIMAL_INTEGER
  : NON_ZERO_DIGIT DIGIT*
  | '0'+
  ;

FLOAT_NUMBER
  : EXPONENT_OR_POINT_FLOAT
  ;

STRING
  : SHORT_STRING
  | LONG_STRING
  ;

COMPLEX_NUMBER
  : '(' FLOAT_NUMBER SIGN FLOAT_NUMBER 'i)'
  ;


TRUE
  : 'true'
  | 'True'
  | 'TRUE'
  | 'T'
  | 'yes'
  | 'Yes'
  | 'YES'
  ;

FALSE
  : 'false'
  | 'False'
  | 'FALSE'
  | 'F'
  | 'no'
  | 'No'
  | 'NO'
  ;

SYS
  : ('_'|'@'|'%'|'?')+
  ;

CMD
  : ('↑'|'√'|'≉'|'≈'|'∐'|'∏'|'∇'|'∆'|'∪'|'∩'|'∉'|'∈'|'⊉'|'⊈'|'⊇'|'⊆'|'⊅'|'⊄'|'⊃'|'⊂'|'÷'|'\\'|'+'|'-'|'&'|'='|'<'|'>'|'*'|'×')+
  ;

SYSF
  : ('∧'|'∨'|'#'|','|'$'|'^'|'_'|'!'|'¬'|'∀')+
  ;

SLASH:   '/' ;

NAME
  : ID_START ID_CONTINUE*
  ;

TOBEGIN: ':' ;
TOEND:   ';' ;

GLOB
  : GLOB_PATTERN
  ;

URI
  : URI_PATTERN
  ;

UNIXCMD
  : CMD_PATTERN
  ;

JSON
  : JSON_PATTERN
  ;

SEPARATE
  : '|'
  ;


COMMENT
  : '##'  .*? [\r\n] -> skip
  ;
BLOCK_COMMENT
  :   '/*' .*? '*/' -> skip
  ;
WS
  : [ \r\n\t]+ -> skip
  ;
SHEBANG
  : '#' '!' ~('\n'|'\r')* -> channel(HIDDEN)
  ;

fragment NON_ZERO_DIGIT
  : [1-9]
  ;

fragment DIGIT
  : [0-9]
  ;

fragment SIGN
  : ('+' | '-')
  ;

fragment EXPONENT_OR_POINT_FLOAT
  : (SIGN)? ([0-9]+ | POINT_FLOAT) [eE] [+-]? [0-9]+
  | (SIGN)? POINT_FLOAT
  ;

fragment EXPONENT_OR_POINT_UFLOAT
  : ('U'|'u') ([0-9]+ | POINT_FLOAT) [eE] [+-]? [0-9]+
  | ('U'|'u') POINT_FLOAT
  ;

fragment POINT_FLOAT
  : [0-9]* '.' [0-9]+
  | [0-9]+ '.'
  ;

fragment GLOB_PATTERN
  : 'g\'' ('\\' (RN | .) | ~[\\\r\n'])* '\''
  | 'g"'  ('\\' (RN | .) | ~[\\\r\n"])* '"'
  ;

fragment URI_PATTERN
  : 'f\'' ('\\' (RN | .) | ~[\\\r\n'])* '\''
  | 'f"'  ('\\' (RN | .) | ~[\\\r\n"])* '"'
  ;

fragment CMD_PATTERN
  : '!\'' ('\\' (RN | .) | ~[\\\r\n'])* '\''
  | '!"'  ('\\' (RN | .) | ~[\\\r\n"])* '"'
  ;

fragment JSON_PATTERN
  : 'j\'' ('\\' (RN | .) | ~[\\\r\n'])* '\''
  | 'j"'  ('\\' (RN | .) | ~[\\\r\n"])* '"'
  ;

fragment RN
  : '\r'? '\n'
  ;

fragment SHORT_STRING
  : '\'' ('\\' (RN | .) | ~[\\\r\n'])* '\''
  | '"'  ('\\' (RN | .) | ~[\\\r\n"])* '"'
  ;

fragment LONG_STRING
  : '\'\'\'' LONG_STRING_ITEM*? '\'\'\''
  | '"""' LONG_STRING_ITEM*? '"""'
  ;

fragment LONG_STRING_ITEM
  : ~'\\'
  | '\\' (RN | .)
  ;

fragment ID_START
 : ([A-Z]|[a-z]|SLASH)
 | [a-z]
 ;

fragment ID_CONTINUE
 : ID_START
 | [0-9]
 | SLASH
 ;
