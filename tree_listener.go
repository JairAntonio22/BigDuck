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

    filename    string

    valid       bool
    debug       bool

    // variables to keep semantic analysis state
    symtable    structs.SymTable

    symqueue    structs.Queue
    typequeue   structs.Queue
    paramqueue  structs.Queue
    dimqueue    structs.Queue

    in_decl     bool
    in_args     bool
    in_stmt     bool

    scope       int
    argc        int

    ret_type    int
    curr_proc   string

    // variables to generarte intermidiate representation code
    ir_code     []structs.Tac
    data_seg    []structs.Tac

    opstack     structs.Stack
    argstack    structs.Stack
    typestack   structs.Stack

    jmpstack    structs.Stack
    skipqueue   structs.Queue
    breakqueue  structs.Queue
    eraqueue    structs.Queue

    pc          int
    tmpc        int
    paramc      int
    loopstyle   int
    startpoint  int
    startproc   string

    curr_line   int
    curr_col    int
    curr_pcall  string
    curr_tensor structs.Symbol
    curr_dim    int

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
    l.ir_code = append(l.ir_code, structs.Tac{Op: structs.ERA})
    l.pc++
    l.ir_code = append(l.ir_code, structs.Tac{Op: structs.GOPROC})
    l.pc++
 }

func (l *BigDuckListener) ExitProgram(c *parser.ProgramContext) {
    if l.valid {
        _, sym, _ := l.symtable.Lookup(l.startproc)

        l.ir_code[0].Args[0] = strconv.Itoa(sym.Ic)
        l.ir_code[0].Args[1] = strconv.Itoa(sym.Fc)
        l.ir_code[0].Args[2] = strconv.Itoa(sym.Bc)
        l.ir_code[0].Address[0] = sym.Ic
        l.ir_code[0].Address[1] = sym.Fc
        l.ir_code[0].Address[2] = sym.Bc

        l.ir_code[1].Args[2] = l.startproc
        l.ir_code[1].Address[2] = l.startpoint

        l.data_seg = l.memmap.GetDataSegment()
        l.GenerateObjFile()
    }

    if l.debug {
        for index, code := range l.data_seg {
            fmt.Printf("%3d ", index);
            code.Print()
        }

        fmt.Println()

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
            var baddress int
            stype := structs.TypeFromString[c.GetText()]

            if len(var_dim) > 0 {
                size := 1

                for _, dim := range var_dim {
                    size *= dim
                }

                baddress = l.memmap.GetAddress(l.scope, name, stype)

                for i := 1; i < size; i++ {
                    l.memmap.GetAddress(l.scope, name + strconv.Itoa(i), stype)
                }
            }

            l.symtable.Insert(
                l.scope,
                name,
                structs.Symbol {
                    Stype: stype,
                    Dim: var_dim,
                    Baddress: baddress})

            if l.in_args {
                l.typequeue.Push(stype)
                l.paramqueue.Push(name)
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
        } else if n <= 0 {
            l.valid = false
            fmt.Printf(
                "line %d:%d tensor dimension must be greater than 0\n",
                c.GetStart().GetLine(), c.GetStart().GetColumn())
        } else {
            l.dimqueue.Push(n)
        }
    }
}

// rbracket
func (l *BigDuckListener) EnterRbracket(c *parser.RbracketContext) {
    if !l.valid {
        return
    }

    if l.in_decl {
        return
    }

    item := l.argstack.Top()
    result, _ := item.(string)
    item = l.typestack.Top()
    rtype, _ := item.(int)


    if rtype != structs.Int_t {
        l.valid = false
        fmt.Printf(
            "line %d:%d index value must be of type int\n",
            c.GetStart().GetLine(), c.GetStart().GetColumn())
        return
    }

    scope := structs.Local
    _, _, exists := l.symtable.Lookup(result)

    if !exists || (len(result) >= 1 && result[0] != 't') {
        scope = structs.Global
    }

    dim := l.curr_tensor.Dim[l.curr_dim]

    address1:= l.memmap.GetAddress(scope, result, rtype)
    address2:= l.memmap.GetAddress(
        structs.Global, strconv.Itoa(dim), structs.Int_t)

    l.ir_code = append(
        l.ir_code, structs.Tac{
            Op: structs.ASSERT,
            Args: [3]string{result, "0", strconv.Itoa(dim)},
            Address: [3]int{address1, 0, address2}})
    l.pc++

    if l.curr_dim + 1 < len(l.curr_tensor.Dim) {
        m := 1

        for i := l.curr_dim + 1; i < len(l.curr_tensor.Dim); i++ {
            m *= l.curr_tensor.Dim[i]
        }

        l.argstack.Push(strconv.Itoa(m))
        l.typestack.Push(structs.Int_t)

        l.PushOp(structs.OpFromString["*"])
        l.GenerateOpTAC()


    } else {
        l.argstack.Push(strconv.Itoa(l.curr_tensor.Baddress))
        l.typestack.Push(structs.Int_t)

        l.PushOp(structs.OpFromString["+"])
        l.GenerateOpTAC()
    }

    l.curr_dim++

    if l.curr_dim > len(l.curr_tensor.Dim) {
        l.valid = false
        fmt.Printf(
            "line %d:%d tensor access does not match with dimensions\n",
            c.GetStart().GetLine(),
            c.GetStart().GetColumn())
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

    if l.ret_type != structs.Void_t {
        l.symtable.Insert(
            structs.Global,
            "_" + l.curr_proc,
            structs.Symbol{Stype: l.ret_type})
    }

    _, sym, _ := l.symtable.Lookup(l.curr_proc)

    sym.Ic = l.memmap.Typecount[structs.Local][structs.Int_t]
    sym.Fc = l.memmap.Typecount[structs.Local][structs.Float_t]
    sym.Bc = l.memmap.Typecount[structs.Local][structs.Bool_t]

    l.symtable.Update(l.curr_proc, sym)

    for !l.eraqueue.Empty() {
        item, _ := l.eraqueue.Pop()
        pc, _ := item.(int)
        l.ir_code[pc].Args[0] = strconv.Itoa(sym.Ic)
        l.ir_code[pc].Args[1] = strconv.Itoa(sym.Fc)
        l.ir_code[pc].Args[2] = strconv.Itoa(sym.Bc)

        l.ir_code[pc].Address[0] = sym.Ic
        l.ir_code[pc].Address[1] = sym.Fc
        l.ir_code[pc].Address[2] = sym.Bc
    }

    if l.debug {
        l.symtable.Print()
    }

    l.symtable.ClearLocalScope()
    l.memmap.ClearLocalScope()
    l.scope = structs.Global
    l.ret_type = structs.Void_t
    l.argc = 0
    l.tmpc = 0
}

// proc_info   
func (l *BigDuckListener) ExitProc_info(c *parser.Proc_infoContext) {
    if !l.valid {
        return
    }

    var typeArgs []int
    var paddress []int

    for i := 0; !l.typequeue.Empty(); i++ {
        item, _ := l.typequeue.Pop()
        stype, _ := item.(int)
        typeArgs = append(typeArgs, stype)

        item, _ = l.paramqueue.Pop()
        name, _ := item.(string)
        paddress = append(
            paddress, l.memmap.GetAddress(structs.Local, name, stype))
    }

    _, sym, _ := l.symtable.Lookup(l.curr_proc)

    sym.Stype = structs.Proc_t
    sym.Argc = l.argc
    sym.TypeArgs = typeArgs
    sym.RetType = l.ret_type
    sym.Paddress = paddress

    l.symtable.Update(l.curr_proc, sym)
}

// sign        
func (l *BigDuckListener) EnterSign(c *parser.SignContext) {
    if !l.valid {
        return
    }

    l.startpoint = l.pc
    l.startproc = c.ID().GetText()
    _, _, exists := l.symtable.Lookup(c.ID().GetText())

    if exists {
        l.valid = false
        fmt.Printf(
            "line %d:%d duplicate symbol %s\n",
            c.GetStart().GetLine(), c.GetStart().GetColumn(), c.ID().GetText())
    } else {
        l.symtable.Insert(l.scope, c.ID().GetText(), structs.Symbol{
            Startpoint: l.startpoint})
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
        if c.TRUE() != nil {
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
        if c.CTE_INT() != nil {
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

// variable
func (l *BigDuckListener) EnterVariable(c *parser.VariableContext) {
    _, sym, exists := l.symtable.Lookup(c.ID().GetText())

    if !exists {
        l.valid = false
        fmt.Printf(
            "line %d:%d variable %s was not declared\n",
            c.GetStart().GetLine(),
            c.GetStart().GetColumn(),
            c.ID().GetText())
        return
    }

    l.argstack.Push(c.ID().GetText())
    l.typestack.Push(sym.Stype)

    if c.Dim() != nil {
        l.argstack.Pop()
        l.typestack.Pop()

        if len(sym.Dim) == 0 {
            l.valid = false
            fmt.Printf(
                "line %d:%d variable %s is scalar\n",
                c.GetStart().GetLine(),
                c.GetStart().GetColumn(),
                c.ID().GetText())
        } else {
            l.curr_dim = 0
            l.curr_tensor = sym
            l.PushOp(structs.LPAREN)
        }
    }
}

func (l *BigDuckListener) ExitVariable(c *parser.VariableContext) {
    if c.Dim() != nil {
        l.PushOp(structs.RPAREN)

        if l.curr_dim != len(l.curr_tensor.Dim) {
            l.valid = false
            fmt.Printf(
                "line %d:%d tensor access does not match with dimensions\n",
                c.GetStart().GetLine(),
                c.GetStart().GetColumn())
        }
    }
}

// t_access
func (l *BigDuckListener) EnterT_access(c *parser.T_accessContext) {
}

// t_end
func (l *BigDuckListener) EnterT_end(c *parser.T_endContext) {
}

// proc_call   
func (l *BigDuckListener) EnterProc_call(c *parser.Proc_callContext) {
    if !l.valid {
        return
    }

    l.paramc = 0
    l.curr_pcall = c.ID().GetText()

    _, sym, exists := l.symtable.Lookup(l.curr_pcall)

    if !exists {
        l.valid = false
        fmt.Printf(
            "line %d:%d procedure %s not declared\n",
            c.GetStart().GetLine(), c.GetStart().GetColumn(), c.ID().GetText())
        return
    }

    if l.curr_proc == l.curr_pcall {
        l.eraqueue.Push(l.pc)
    }

    l.ir_code = append(
        l.ir_code,
        structs.Tac{
            Op: structs.ERA,
            Args: [3]string{
                strconv.Itoa(sym.Ic),
                strconv.Itoa(sym.Fc),
                strconv.Itoa(sym.Bc)},
            Address: [3]int{sym.Ic, sym.Fc, sym.Bc}})
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

    } else if l.in_stmt {
        l.ir_code = append(
            l.ir_code,
            structs.Tac{
                Op: structs.GOPROC,
                Args: [3]string{"", "", c.ID().GetText()},
                Address: [3]int{0, 0, sym.Startpoint}})
        l.pc++

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

// void_proc        
func (l *BigDuckListener) EnterVoid_proc(c *parser.Void_procContext) {
    l.in_stmt = true;
}

func (l *BigDuckListener) ExitVoid_proc(c *parser.Void_procContext) {
    l.in_stmt = false;
}

// assignment  
func (l *BigDuckListener) EnterAssignment(c *parser.AssignmentContext) {
    if !l.valid {
        return
    }

    l.PushOp(structs.OpFromString["<-"])
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

// print_r   
func (l *BigDuckListener) ExitPrint_r(c * parser.Print_rContext) {
    if !l.valid {
        return
    }

    l.ir_code[l.pc - 1].Op = structs.PRINTLN
}

// pparam    
// pparamTerm
func (l *BigDuckListener) ExitPparamTerm(c * parser.PparamTermContext) {
    if !l.valid {
        return
    }

    if c.CTE_STRING() != nil {
        address := l.memmap.GetAddress(
            structs.Global, c.CTE_STRING().GetText(), structs.String_t)

        l.ir_code = append(
            l.ir_code, structs.Tac{
                Op: structs.PRINT,
                Args: [3]string{"", "", c.CTE_STRING().GetText()},
                Address: [3]int{0, 0, address}})
        l.pc++

    } else {
        l.GeneratePrintTAC()
    }
}

// pnextParam
