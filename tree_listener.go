package main

import (
    "fmt"
    "strconv"
    "./structs"
    "./parser"
    "github.com/antlr/antlr4/runtime/Go/antlr"
)

// BigDuckListener structure

type BigDuckListener struct {
    *parser.BaseBigDuckListener

    valid       bool
    debug       bool

    // variables to keep semantic analysis state
    symtable    structs.SymTable

    symqueue    structs.Queue
    typequeue   structs.Queue
    dimqueue    structs.Queue

    in_decl     bool
    in_args     bool

    scope       int
    argc        int

    ret_type    int
    curr_proc   string

    // variables to generarte intermidiate representation code
    ir_code     []structs.Tac

    opstack     structs.Stack
    argstack    structs.Stack
    typestack   structs.Stack

    jmpstack    structs.Stack
    skipqueue   structs.Queue
    breakqueue  structs.Queue

    pc          int
    tmpc        int
    paramc      int
    loopstyle   int
    startpoint  int

    curr_line   int
    curr_col    int
    curr_pcall  string

    memmap      structs.MemMap
}

func (l *BigDuckListener) VisitErrorNode(node antlr.ErrorNode) {
    l.valid = false
}

// program     
func (l *BigDuckListener) EnterProgram(c *parser.ProgramContext) {
    l.symtable = structs.NewSymTable()
    l.memmap = structs.NewMemMap()
    l.ret_type = structs.Void_t
    l.scope = structs.Global
    l.valid = true
    structs.InitCube()
    l.GenerateJmpTAC(structs.JMP)
 }

func (l *BigDuckListener) ExitProgram(c *parser.ProgramContext) {
    l.FillJmpTAC(0, l.startpoint)

    if l.debug {
        for index, code := range l.ir_code {
            fmt.Printf("%3d ", index);
            code.Print()
        }

        fmt.Println()
    }
}

// vars_decl   

// var_decl    
func (l *BigDuckListener) EnterVar_decl(c *parser.Var_declContext) {
    if !l.valid {
        return
    }

    l.in_decl = true
    l.symqueue.Push(c.ID().GetText())
}

func (l *BigDuckListener) ExitVar_decl(c *parser.Var_declContext) {
    if !l.valid {
        return
    }

    l.in_decl = false
}

// nextVar     
func (l *BigDuckListener) EnterNextVar(c *parser.NextVarContext) {
    if !l.valid {
        return
    }

    if c.ID() != nil {
        l.symqueue.Push(c.ID().GetText())
    }
}

// nextVarDecl 

// var_type    

// scalar      
func (l *BigDuckListener) EnterScalar(c *parser.ScalarContext) {
    if !l.valid {
        return
    }

    var var_dim []int

    if !l.dimqueue.Empty() {
        for !l.dimqueue.Empty() {
            item, _ := l.dimqueue.Pop()

            if n, ok := item.(int); ok {
                var_dim = append(var_dim, n)
            }
        }
    }

    for !l.symqueue.Empty() {
        item, _ := l.symqueue.Pop()
        name, _ := item.(string)
        _, _, exists := l.symtable.Lookup(name)

        if exists {
            l.valid = false
            fmt.Printf(
                "line %d:%d duplicate symbol %s\n",
                c.GetStart().GetLine(), c.GetStart().GetColumn(), name)
        } else {
            stype := structs.TypeFromString[c.GetText()]

            l.symtable.Insert(
                l.scope,
                name,
                structs.Symbol {
                    Stype: stype,
                    Dim: var_dim})

            if l.in_args {
                l.typequeue.Push(stype)
            }
        }
    }
}

// tensor      

// dim         
func (l *BigDuckListener) EnterDim(c *parser.DimContext) {
    if !l.valid {
        return
    }

    if l.in_decl {
        n, err := strconv.Atoi(c.Num_expr().GetText())

        if err != nil {
            l.valid = false
            fmt.Printf(
                "line %d:%d tensor dimension must be constant\n",
                c.GetStart().GetLine(), c.GetStart().GetColumn())
        } else {
            l.dimqueue.Push(n)
        }
    }
}

// nextDim     

// procs_decl  

// proc_decl   
func (l *BigDuckListener) ExitProc_decl(c *parser.Proc_declContext) {
    if !l.valid {
        return
    }

    l.ir_code = append(l.ir_code, structs.Tac{Op: structs.ENDPROC})
    l.pc++

    var typeArgs []int

    for !l.typequeue.Empty() {
        item, _ := l.typequeue.Pop()
        stype, _ := item.(int)
        typeArgs = append(typeArgs, stype)
    }

    _, sym, _ := l.symtable.Lookup(l.curr_proc)

    for _, argtype := range typeArgs {
        switch argtype {
        case structs.Int_t:
            sym.Ic += 1

        case structs.Float_t:
            sym.Fc += 1

        case structs.Bool_t:
            sym.Bc += 1
        }
    }

    sym.Stype = structs.Proc_t
    sym.Argc = l.argc
    sym.TypeArgs = typeArgs
    sym.RetType = l.ret_type

    l.symtable.Update(l.curr_proc, sym)

    if l.ret_type != structs.Void_t {
        l.symtable.Insert(
            structs.Global,
            "_" + l.curr_proc,
            structs.Symbol{Stype: l.ret_type})
    }

    l.symtable.ClearLocalScope()
    l.memmap.ClearLocalScope()
    l.scope = structs.Global
    l.ret_type = structs.Void_t
    l.argc = 0
    l.tmpc = 0
}

// sign        
func (l *BigDuckListener) EnterSign(c *parser.SignContext) {
    if !l.valid {
        return
    }

    l.startpoint = l.pc
    _, _, exists := l.symtable.Lookup(c.ID().GetText())

    if exists {
        l.valid = false
        fmt.Printf(
            "line %d:%d duplicate symbol %s\n",
            c.GetStart().GetLine(), c.GetStart().GetColumn(), c.ID().GetText())
    } else {
        l.symtable.Insert(l.scope, c.ID().GetText(), structs.Symbol{})
        l.curr_proc = c.ID().GetText()
        l.scope = structs.Local
    }
}

// args        
func (l *BigDuckListener) EnterArgs(c *parser.ArgsContext) {
    if !l.valid {
        return
    }

    if c.ID() != nil {
        l.symqueue.Push(c.ID().GetText())
        l.in_decl = true
        l.in_args = true
        l.argc++
    }
}

func (l *BigDuckListener) ExitArgs(c *parser.ArgsContext) {
    if !l.valid {
        return
    }

    l.in_decl = false
    l.in_args = false
}

// nextTypes   
func (l *BigDuckListener) EnterNextTypes(c *parser.NextTypesContext) {
    if !l.valid {
        return
    }

    if c.ID() != nil {
        l.symqueue.Push(c.ID().GetText())
        l.argc++
    }
}

// nextArg     
func (l *BigDuckListener) EnterNextArg(c *parser.NextArgContext) {
    if !l.valid {
        return
    }

    if c.ID() != nil {
        l.symqueue.Push(c.ID().GetText())
        l.argc++
    }
}

// ret_type    
func (l *BigDuckListener) EnterRet_type(c *parser.Ret_typeContext) {
    if !l.valid {
        return
    }

    l.ret_type = structs.TypeFromString[c.Scalar().GetText()]
}

// bool_expr   

// nextBool    
func (l *BigDuckListener) EnterNextBool(c *parser.NextBoolContext) {
    if !l.valid {
        return
    }

    if c.GetText() != "" {
        l.PushOp(structs.OpFromString["or"])
    }
}

// and_expr    
func (l *BigDuckListener) ExitAnd_expr(c *parser.And_exprContext) {
    if !l.valid {
        return
    }

    if l.TopOp() == structs.OR {
        l.curr_line = c.GetStart().GetLine()
        l.curr_col = c.GetStart().GetColumn()
        l.GenerateOpTAC()
    }
}

// nextAnd     
func (l *BigDuckListener) EnterNextAnd(c *parser.NextAndContext) {
    if !l.valid {
        return
    }

    if c.GetText() != "" {
        l.PushOp(structs.OpFromString["and"])
    }
}

// not_expr    
func (l *BigDuckListener) EnterNot_expr(c *parser.Not_exprContext) {
    if !l.valid {
        return
    }

    if c.NOT() != nil {
        l.PushOp(structs.OpFromString["not"])
    }
}

func (l *BigDuckListener) ExitNot_expr(c *parser.Not_exprContext) {
    if !l.valid {
        return
    }

    if c.NOT() != nil {
        l.curr_line = c.GetStart().GetLine()
        l.curr_col = c.GetStart().GetColumn()
        l.GenerateOpTAC()
    }
}

// bool_term   
func (l *BigDuckListener) EnterBool_term(c *parser.Bool_termContext) {
    if !l.valid {
        return
    }

    if c.Bool_expr() != nil {
        l.PushOp(structs.LPAREN)
    }
}

func (l *BigDuckListener) ExitBool_term(c *parser.Bool_termContext) {
    if !l.valid {
        return
    }

    if c.Bool_expr() == nil {
        if c.ID() != nil {
            _, sym, exists := l.symtable.Lookup(c.ID().GetText())

            if !exists {
                l.valid = false
                fmt.Printf(
                    "line %d:%d variable %s was not declared\n",
                    c.GetStart().GetLine(),
                    c.GetStart().GetColumn(),
                    c.ID().GetText())
            }

            l.argstack.Push(c.ID().GetText())
            l.typestack.Push(sym.Stype)

        } else if c.TRUE() != nil {
            l.argstack.Push("#t")
            l.typestack.Push(structs.Bool_t)

        } else if c.FALSE() != nil {
            l.argstack.Push("#f")
            l.typestack.Push(structs.Bool_t)
        }

    } else {
        l.PushOp(structs.RPAREN)
    }

    top := l.TopOp()

    if top == structs.AND {
        l.curr_line = c.GetStart().GetLine()
        l.curr_col = c.GetStart().GetColumn()
        l.GenerateOpTAC()
    }
}

// rel_expr    
func (l *BigDuckListener) ExitRel_expr(c *parser.Rel_exprContext) {
    if !l.valid {
        return
    }

    l.curr_line = c.GetStart().GetLine()
    l.curr_col = c.GetStart().GetColumn()
    l.GenerateOpTAC()
}

// opRel       
func (l *BigDuckListener) EnterOpRel(c *parser.OpRelContext) {
    if !l.valid {
        return
    }

    l.PushOp(structs.OpFromString[c.RelOp().GetText()])
}

// relOp       
// num_expr    
// nextSum     
func (l *BigDuckListener) EnterNextSum(c *parser.NextSumContext) {
    if !l.valid {
        return
    }

    if c.GetText() != "" {
        if c.GetText()[0] == '+' {
            l.PushOp(structs.OpFromString["+"])

        } else if c.GetText()[0] == '-' {
            l.PushOp(structs.OpFromString["-"])
        }
    }
}

// prod_expr   
func (l *BigDuckListener) ExitProd_expr(c *parser.Prod_exprContext) {
    if !l.valid {
        return
    }

    if l.TopOp() == structs.ADD || l.TopOp() == structs.SUB {
        l.curr_line = c.GetStart().GetLine()
        l.curr_col = c.GetStart().GetColumn()
        l.GenerateOpTAC()
    }
}

// nextProd    
func (l *BigDuckListener) EnterNextProd(c *parser.NextProdContext) {
    if !l.valid {
        return
    }

    if c.GetText() != "" {
        if c.GetText()[0] == '*' {
            l.PushOp(structs.OpFromString["*"])

        } else if c.GetText()[0] == '/' {
            l.PushOp(structs.OpFromString["/"])
        }
    }
}

// factor      
func (l *BigDuckListener) EnterFactor(c *parser.FactorContext) {
    if !l.valid {
        return
    }

    if c.Num_expr() != nil {
        l.PushOp(structs.LPAREN)
    }
}

func (l *BigDuckListener) ExitFactor(c *parser.FactorContext) {
    if !l.valid {
        return
    }

    if c.Num_expr() == nil {
        if c.ID() != nil {
            _, sym, exists := l.symtable.Lookup(c.ID().GetText())

            if !exists {
                l.valid = false
                fmt.Printf(
                    "line %d:%d variable %s was not declared\n",
                    c.GetStart().GetLine(),
                    c.GetStart().GetColumn(),
                    c.ID().GetText())
            }

            l.argstack.Push(c.ID().GetText())
            l.typestack.Push(sym.Stype)

        } else if c.CTE_INT() != nil {
            l.argstack.Push(c.CTE_INT().GetText())
            l.typestack.Push(structs.Int_t)

        } else if c.CTE_FLOAT() != nil {
            l.argstack.Push(c.CTE_FLOAT().GetText())
            l.typestack.Push(structs.Float_t)
        }

    } else {
        l.PushOp(structs.RPAREN)
    }

    top := l.TopOp()

    if top == structs.MUL || top == structs.DIV {
        l.curr_line = c.GetStart().GetLine()
        l.curr_col = c.GetStart().GetColumn()
        l.GenerateOpTAC()
    }
}

// proc_call   
func (l *BigDuckListener) EnterProc_call(c *parser.Proc_callContext) {
    if !l.valid {
        return
    }

    l.paramc = 0
    l.curr_pcall = c.ID().GetText()

    _, _, exists := l.symtable.Lookup(c.ID().GetText())

    if !exists {
        l.valid = false
        fmt.Printf(
            "line %d:%d procedure %s not declared\n",
            c.GetStart().GetLine(), c.GetStart().GetColumn(), c.ID().GetText())
        return
    }

    l.ir_code = append(
        l.ir_code,
        structs.Tac{
            Op: structs.ERA,
            Args: [3]string{"","", c.ID().GetText()}})
    l.pc++
}

func (l *BigDuckListener) ExitProc_call(c *parser.Proc_callContext) {
    if !l.valid {
        return
    }

    _, sym, _ := l.symtable.Lookup(c.ID().GetText())

    if sym.RetType != structs.Void_t {
        l.curr_line = c.GetStart().GetLine()
        l.curr_col = c.GetStart().GetColumn()
        l.GenerateProcCallRetTAC(c.ID().GetText())
    } else {
        l.valid = false
        fmt.Printf(
            "line %d:%d void procedure used in expression\n",
            l.curr_line,
            l.curr_col)
        return
    }

    if l.paramc != sym.Argc {
        l.valid = false
        fmt.Printf(
            "line %d:%d procedure %s expected %d arguments, given %d\n",
            c.GetStart().GetLine(),
            c.GetStart().GetColumn(),
            c.ID().GetText(),
            sym.Argc,
            l.paramc)
    }
}

// param 

// paramTerm
func (l *BigDuckListener) EnterParamTerm(c *parser.ParamTermContext) {
    if !l.valid {
        return
    }

    l.PushOp(structs.LPAREN)
}

func (l *BigDuckListener) ExitParamTerm(c *parser.ParamTermContext) {
    if !l.valid {
        return
    }

    l.GenerateParamTAC()
    l.PushOp(structs.RPAREN)
}

// nextParam 

// block       

// stmts       

// stmt        

// assignment  
func (l *BigDuckListener) EnterAssignment(c *parser.AssignmentContext) {
    if !l.valid {
        return
    }

    l.PushOp(structs.OpFromString["<-"])

    _, sym, exists := l.symtable.Lookup(c.ID().GetText())

    if !exists {
        l.valid = false
        fmt.Printf(
            "line %d:%d variable %s was not declared\n",
            c.GetStart().GetLine(),
            c.GetStart().GetColumn(),
            c.ID().GetText())
    }

    l.argstack.Push(c.ID().GetText())
    l.typestack.Push(sym.Stype)
}

func (l *BigDuckListener) ExitAssignment(c *parser.AssignmentContext) {
    if !l.valid {
        return
    }

    l.curr_line = c.GetStart().GetLine()
    l.curr_col = c.GetStart().GetColumn()
    l.GenerateOpTAC()
}

// condition   

// bodyCond    
func (l *BigDuckListener) EnterBodyCond(c *parser.BodyCondContext) {
    if !l.valid {
        return
    }

    l.jmpstack.Push(l.pc)
    l.curr_line = c.GetStart().GetLine()
    l.curr_col = c.GetStart().GetColumn()
    l.GenerateJmpTAC(structs.JMF)
}

// endIfBlock  
func (l *BigDuckListener) EnterEndIfBlock(c *parser.EndIfBlockContext) {
    if !l.valid {
        return
    }

    item, _ := l.jmpstack.Pop()
    index, _ := item.(int)

    if c.Alter() == nil {
        l.FillJmpTAC(index, l.pc)
    } else {
        l.FillJmpTAC(index, l.pc + 1)
    }
}

// alter       
func (l *BigDuckListener) EnterAlter(c *parser.AlterContext) {
    if !l.valid {
        return
    }

    l.jmpstack.Push(l.pc)
    l.GenerateJmpTAC(structs.JMP)
}

func (l *BigDuckListener) ExitAlter(c *parser.AlterContext) {
    if !l.valid {
        return
    }

    item, _ := l.jmpstack.Pop()
    index, _ := item.(int)
    l.FillJmpTAC(index, l.pc)
}

// loop_stmt   
func (l *BigDuckListener) ExitLoop_stmt(c *parser.Loop_stmtContext) {
    if !l.valid {
        return
    }

    var index, target int
    var item interface{}

    l.GenerateJmpTAC(structs.JMP)

    switch l.loopstyle {
    case structs.InfStyle:
        item, _ = l.jmpstack.Pop()
        target, _ = item.(int)
        l.FillJmpTAC(l.pc - 1, target)

    case structs.WhileStyle:
        item, _ = l.jmpstack.Pop()
        index, _ = item.(int)
        l.FillJmpTAC(index, l.pc)

        item, _ = l.jmpstack.Pop()
        target, _ = item.(int)
        l.FillJmpTAC(l.pc - 1, target)

    case structs.ForStyle:
        item, _ = l.jmpstack.Pop()
        target, _ = item.(int)
        l.FillJmpTAC(l.pc - 1, target)

        item, _ = l.jmpstack.Pop()
        index, _ = item.(int)
        l.FillJmpTAC(index, l.pc)
    }

    for !l.skipqueue.Empty() {
        item, _ = l.skipqueue.Pop()
        index, _ = item.(int)
        l.FillJmpTAC(index, target)
    }

    for !l.breakqueue.Empty() {
        item, _ = l.breakqueue.Pop()
        index, _ = item.(int)
        l.FillJmpTAC(index, l.pc)
    }
}

func (l *BigDuckListener) EnterCtrl_flow(c *parser.Ctrl_flowContext) {
    if !l.valid {
        return
    }

    if c.BREAK() != nil {
        l.breakqueue.Push(l.pc)
        l.GenerateJmpTAC(structs.JMP)
    } else if c.SKIP_W() != nil {
        l.skipqueue.Push(l.pc)
        l.GenerateJmpTAC(structs.JMP)
    }
}

// forStyle    

// forCond     
func (l *BigDuckListener) EnterForCond(c *parser.ForCondContext) {
    if !l.valid {
        return
    }

    l.jmpstack.Push(l.pc)
    l.loopstyle = structs.ForStyle
}

func (l *BigDuckListener) ExitForCond(c *parser.ForCondContext) {
    if !l.valid {
        return
    }

    l.jmpstack.Push(l.pc)
    l.curr_line = c.GetStart().GetLine()
    l.curr_col = c.GetStart().GetColumn()
    l.GenerateJmpTAC(structs.JMT)
    l.jmpstack.Push(l.pc)
    l.GenerateJmpTAC(structs.JMP)
}

// ctrl_var    
func (l *BigDuckListener) EnterCtrl_var(c *parser.Ctrl_varContext) {
    if !l.valid {
        return
    }

    l.jmpstack.Push(l.pc)
}

func (l *BigDuckListener) ExitCtrl_var(c *parser.Ctrl_varContext) {
    if !l.valid {
        return
    }

    l.GenerateJmpTAC(structs.JMP)

    var tmpstack structs.Stack
    var index, target int
    var item interface{}

    for i := 0; i < 3; i ++ {
        item, _ = l.jmpstack.Pop()
        tmpstack.Push(item)
    }

    item, _ = l.jmpstack.Pop()
    target, _ = item.(int)
    l.FillJmpTAC(l.pc - 1, target)

    item, _ = tmpstack.Pop()
    index, _ = item.(int)
    l.FillJmpTAC(index, l.pc)

    for !tmpstack.Empty() {
        item, _ = tmpstack.Pop()
        l.jmpstack.Push(item)
    }
}

// whileStyle  
func (l *BigDuckListener) EnterWhileStyle(c *parser.WhileStyleContext) {
    if !l.valid {
        return
    }

    l.jmpstack.Push(l.pc)
    l.loopstyle = structs.WhileStyle
}

func (l *BigDuckListener) ExitWhileStyle(c *parser.WhileStyleContext) {
    if !l.valid {
        return
    }

    l.jmpstack.Push(l.pc)
    l.curr_line = c.GetStart().GetLine()
    l.curr_col = c.GetStart().GetColumn()
    l.GenerateJmpTAC(structs.JMF)
}

// infLoop     
func (l *BigDuckListener) EnterInfLoop(c *parser.InfLoopContext) {
    if !l.valid {
        return
    }

    l.jmpstack.Push(l.pc)
    l.loopstyle = structs.InfStyle
}

// ctrl_flow   

// ret_stmt    
func (l *BigDuckListener) ExitRet_stmt(c * parser.Ret_stmtContext) {
    if !l.valid {
        return
    }

    l.curr_line = c.GetStart().GetLine()
    l.curr_col = c.GetStart().GetColumn()
    l.GenerateReturnTAC()
}
