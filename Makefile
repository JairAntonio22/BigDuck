
ANTLR=/usr/local/lib/antlr-4.9-complete.jar

DCC_FILES=\
	  compiler.go\
	  tree_listener.go\
	  tac_gen.go

DRUN_FILES=vmachine.go

all: dcc

test: dcc
	clear
	./dcc examples/test_expr.duck

drun: $(DRUN_FILES)
	go build -o drun $(DRUN_FILES)

dcc: parser structs $(DCC_FILES)
	go build -o dcc $(DCC_FILES)

parser: BigDuck.g4
	java -jar $(ANTLR) BigDuck.g4 -Dlanguage=Go -o parser

clean:
	rm -fr parser dcc
