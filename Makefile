
ANTLR=/usr/local/lib/antlr-4.9-complete.jar

all: parser
	go build -o BigDuck main.go

test: parser
	./BigDuck test.duck

parser: BigDuck.g4
	java -jar $(ANTLR) BigDuck.g4 -Dlanguage=Go -o parser

clean:
	rm -r parser BigDuck
