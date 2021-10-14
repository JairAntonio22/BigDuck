
ANTLR=/usr/local/lib/antlr-4.9-complete.jar

all: BigDuck

test: BigDuck
	clear
	./BigDuck test_expr.duck

BigDuck: parser structs *.go
	go build -o BigDuck *.go

parser: BigDuck.g4
	java -jar $(ANTLR) BigDuck.g4 -Dlanguage=Go -o parser

clean:
	rm -fr parser BigDuck
