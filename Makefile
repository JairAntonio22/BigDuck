
ANTLR=/usr/local/lib/antlr-4.9-complete.jar

DCC_FILES=compiler.go tree_listener.go tac_gen.go
DRUN_FILES=vmachine.go

TARGET_CC=ducc
TARGET_VM=drun

all: $(TARGET_CC)

test: $(TARGET_CC)
	clear
	./$(TARGET_CC) examples/test_expr.duck

$(TARGET_VM): $(DRUN_FILES)
	go build -o $(TARGET_VM) $(DRUN_FILES)

$(TARGET_CC): parser structs $(DCC_FILES)
	go build -o $(TARGET_CC) $(DCC_FILES)

parser: BigDuck.g4
	java -jar $(ANTLR) BigDuck.g4 -Dlanguage=Go -o parser

clean:
	rm -fr parser ducc
