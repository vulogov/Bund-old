# Features

- Add !"" unixcmd datatype. Content of the string will be considered as UNIX command which will be executed and STDOUT will be placed as string object to the stack
- Add j"" json datatype. Content of the string will be considered as JSON which will be parsed and stored as the value of the object of type json, placed on the stack.
- str/Strip and str/Lines functions
- fs/File/* functions
- Functors now a "first class citizens".
```
/* Converting string to json with functor json */
( '{"answer": 42}':(json) )
```
- JSON now is a datatype and have a functor for converting the strings
- JSON querying and generation functions as json/*
- new functor "all" - executing operations over all content of  the stack
- new loop "loop" - running loop until externally interrupted
- new functions in args* tree. They will query what you pass through --args
- new loop function 'times' which call function n times
- new loop function 'over' which pushes data from (* ) to stack and call function
- new function 'type' returns name of the type of data element on the stack
- new function 'seq' - generating sequence of the values in DBLOCK
- new function 'rotateFront'/'rotate', 'rotateBack' - rotating the stack front to back or back to front
- new datatype and opcode http
- string datatype supports pre-function feature
- data retrieval from remote http/https endpoints
- http* functions for retrieval from remote http/https endpoints

# Bugfixes

- Fixing typos in MAT matrix files
- Fix bug in sleep function causing overflow if sleep specified in integers
- Fix the fact that \` references was declared, but not implemented
