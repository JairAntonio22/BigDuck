
ANTLR=/usr/local/lib/antlr-4.9-complete.jar

all: parser
	go build -o BigDuck main.go 

parser: BigDuck.g4
	java -jar $(ANTLR) -Dlanguage=Go -o parser BigDuck.g4

clean:
	rm -r parser BigDuck
