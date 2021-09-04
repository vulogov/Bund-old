all: Bund.g4
	rm -rf parser
	antlr -Dlanguage=Go -o parser -atn Bund.g4
	antlr -Dlanguage=Go -o parser  Bund.g4
	( cd ./parser ; ../../script/generateBNF.sh )
