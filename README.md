# the BUND

There are many programming languages out there, and I do dare to introduce a new one. The BUND programming language carries a legacy from a number of other languages that came before the BUND.

Bund is an interpreted high-level general-purpose programming language. It is shared a features from different “ancestors” while not being a "next generation" of any of them. Out of many BUND ancestors, I can specifically mention FORTH, Lisp and Haskell (as very distant relative). Like a FORTH, BUND is stack-based, but it is different from FORTH in respect that it is not a procedural language. It is functional. It is also different in how BUND stack machine works. Like LISP and Haskell (remember, that one is very distant relative, but relative nevertheless), BUND is functional. It is supports lambda calculus with named and anonymous lambdas. Lambda functions in BUND are functionally strict, they are taking only one parameter and always return a value. And we will discuss them in a separate tutorial. In addition to a functions, BUND offers separate concept of “operators”, which essentially functions, but taking two parameters and returning a value at all times. Like in Haskell, BUND offers a “variable immutability” and unlike in Haskell and LISP, BUND do not have a variables. Unlike LISP, BUND virtual machine operates with Stack not with List and this BUND’s feature differs it from Haskell as well.

## What the BUND is looks like ?

![Hello World](Documentation/HelloWorld.png)
