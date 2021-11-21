

ANTLR=/usr/local/lib/antlr-4.9-complete.jar

TARGET=duck

all: $(TARGET)

install: $(TARGET)
	cp duck /usr/local/bin

uninstall:
	rm -f /usr/local/bin/$(TARGET)

test: $(TARGET)
	rm -fr examples/*.quack
	clear
	./$(TARGET) examples/test.duck
	./$(TARGET) run examples/test.quack

benchmark: $(TARGET)
	rm -fr examples/*.quack
	clear
	time -p ./$(TARGET) examples/test_expr.duck
	time -p ./$(TARGET) run examples/test_expr.quack

$(TARGET): parser structs *.go
	go build -o $(TARGET) *.go

parser: BigDuck.g4
	java -jar $(ANTLR) BigDuck.g4 -Dlanguage=Go -o parser

clean:
	rm -fr parser ducc examples/*.quack /usr/local/bin/$(TARGET)
