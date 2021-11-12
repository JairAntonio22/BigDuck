
ANTLR=/usr/local/lib/antlr-4.9-complete.jar

DCC_FILES=main.go compiler.go tree_listener.go tac_gen.go vm.go

TARGET_CC=duck

all: $(TARGET_CC)

test: $(TARGET_CC)
	clear
	./$(TARGET_CC) examples/test_expr.duck
	./$(TARGET_CC) run examples/test_expr.quack

$(TARGET_CC): parser structs $(DCC_FILES)
	go build -o $(TARGET_CC) $(DCC_FILES)

parser: BigDuck.g4
	java -jar $(ANTLR) BigDuck.g4 -Dlanguage=Go -o parser

clean:
	rm -fr parser ducc
