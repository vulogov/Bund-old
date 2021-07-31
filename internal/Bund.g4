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
    );


SLASH:   '/' ;

NAME
  : ID_START ID_CONTINUE*
  ;

COMMENT
  : '##' ~[\r\n]* -> skip
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

fragment ID_START
 : ([A-Z]|[a-z]|SLASH)
 | [a-z]
 ;

fragment ID_CONTINUE
 : ID_START
 | [0-9]
 | SLASH
 ;
