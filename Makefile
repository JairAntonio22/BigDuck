
ANTLR=/usr/local/lib/antlr-4.9-complete.jar

TARGET_CC=duck

all: $(TARGET_CC)

test: $(TARGET_CC)
	rm -fr examples/*.quack
	clear
	./$(TARGET_CC) examples/test_expr.duck
	./$(TARGET_CC) run examples/test_expr.quack

benchmark: $(TARGET_CC)
	rm -fr examples/*.quack
	clear
	time -p ./$(TARGET_CC) examples/test_expr.duck
	time -p ./$(TARGET_CC) run examples/test_expr.quack

test_vm: $(TARGET_CC)
	clear
	./$(TARGET_CC) run examples/test_expr.quack

test_comp: $(TARGET_CC)
	clear
	./$(TARGET_CC) examples/test_expr.duck

$(TARGET_CC): parser structs *.go
	go build -o $(TARGET_CC) *.go

parser: BigDuck.g4
	java -jar $(ANTLR) BigDuck.g4 -Dlanguage=Go -o parser

clean:
	rm -fr parser ducc examples/*.quack
