package structs

var cube [opEnumCount][typeEnumCount][typeEnumCount]int

func InitCube() {
    cube[ASG][Int_t][Int_t]         = Int_t
    cube[ASG][Int_t][Float_t]       = Int_t
    cube[ASG][Int_t][Bool_t]        = Error_t
    cube[ASG][Int_t][Tensor_t]      = Error_t

    cube[ASG][Float_t][Int_t]       = Float_t
    cube[ASG][Float_t][Float_t]     = Float_t
    cube[ASG][Float_t][Bool_t]      = Error_t
    cube[ASG][Float_t][Tensor_t]    = Error_t

    cube[ASG][Bool_t][Int_t]        = Error_t
    cube[ASG][Bool_t][Float_t]      = Error_t
    cube[ASG][Bool_t][Bool_t]       = Bool_t
    cube[ASG][Bool_t][Tensor_t]     = Error_t

    cube[ASG][Tensor_t][Int_t]      = Error_t
    cube[ASG][Tensor_t][Float_t]    = Error_t
    cube[ASG][Tensor_t][Bool_t]     = Error_t
    cube[ASG][Tensor_t][Tensor_t]   = Tensor_t



    cube[ADD][Int_t][Int_t]         = Int_t
    cube[ADD][Int_t][Float_t]       = Float_t
    cube[ADD][Int_t][Bool_t]        = Error_t
    cube[ADD][Int_t][Tensor_t]      = Error_t

    cube[ADD][Float_t][Int_t]       = Float_t
    cube[ADD][Float_t][Float_t]     = Float_t
    cube[ADD][Float_t][Bool_t]      = Error_t
    cube[ADD][Float_t][Tensor_t]    = Error_t

    cube[ADD][Bool_t][Int_t]        = Error_t
    cube[ADD][Bool_t][Float_t]      = Error_t
    cube[ADD][Bool_t][Bool_t]       = Error_t
    cube[ADD][Bool_t][Tensor_t]     = Error_t

    cube[ADD][Tensor_t][Int_t]      = Error_t
    cube[ADD][Tensor_t][Float_t]    = Error_t
    cube[ADD][Tensor_t][Bool_t]     = Error_t
    cube[ADD][Tensor_t][Tensor_t]   = Tensor_t



    cube[SUB][Int_t][Int_t]         = Int_t
    cube[SUB][Int_t][Float_t]       = Float_t
    cube[SUB][Int_t][Bool_t]        = Error_t
    cube[SUB][Int_t][Tensor_t]      = Error_t

    cube[SUB][Float_t][Int_t]       = Float_t
    cube[SUB][Float_t][Float_t]     = Float_t
    cube[SUB][Float_t][Bool_t]      = Error_t
    cube[SUB][Float_t][Tensor_t]    = Error_t

    cube[SUB][Bool_t][Int_t]        = Error_t
    cube[SUB][Bool_t][Float_t]      = Error_t
    cube[SUB][Bool_t][Bool_t]       = Error_t
    cube[SUB][Bool_t][Tensor_t]     = Error_t

    cube[SUB][Tensor_t][Int_t]      = Error_t
    cube[SUB][Tensor_t][Float_t]    = Error_t
    cube[SUB][Tensor_t][Bool_t]     = Error_t
    cube[SUB][Tensor_t][Tensor_t]   = Tensor_t



    cube[MUL][Int_t][Int_t]         = Int_t
    cube[MUL][Int_t][Float_t]       = Float_t
    cube[MUL][Int_t][Bool_t]        = Error_t
    cube[MUL][Int_t][Tensor_t]      = Error_t

    cube[MUL][Float_t][Int_t]       = Float_t
    cube[MUL][Float_t][Float_t]     = Float_t
    cube[MUL][Float_t][Bool_t]      = Error_t
    cube[MUL][Float_t][Tensor_t]    = Error_t

    cube[MUL][Bool_t][Int_t]        = Error_t
    cube[MUL][Bool_t][Float_t]      = Error_t
    cube[MUL][Bool_t][Bool_t]       = Error_t
    cube[MUL][Bool_t][Tensor_t]     = Error_t

    cube[MUL][Tensor_t][Int_t]      = Error_t
    cube[MUL][Tensor_t][Float_t]    = Error_t
    cube[MUL][Tensor_t][Bool_t]     = Error_t
    cube[MUL][Tensor_t][Tensor_t]   = Tensor_t



    cube[DIV][Int_t][Int_t]         = Int_t
    cube[DIV][Int_t][Float_t]       = Float_t
    cube[DIV][Int_t][Bool_t]        = Error_t
    cube[DIV][Int_t][Tensor_t]      = Error_t

    cube[DIV][Float_t][Int_t]       = Float_t
    cube[DIV][Float_t][Float_t]     = Float_t
    cube[DIV][Float_t][Bool_t]      = Error_t
    cube[DIV][Float_t][Tensor_t]    = Error_t

    cube[DIV][Bool_t][Int_t]        = Error_t
    cube[DIV][Bool_t][Float_t]      = Error_t
    cube[DIV][Bool_t][Bool_t]       = Error_t
    cube[DIV][Bool_t][Tensor_t]     = Error_t

    cube[DIV][Tensor_t][Int_t]      = Error_t
    cube[DIV][Tensor_t][Float_t]    = Error_t
    cube[DIV][Tensor_t][Bool_t]     = Error_t
    cube[DIV][Tensor_t][Tensor_t]   = Tensor_t



    cube[AND][Int_t][Int_t]         = Error_t
    cube[AND][Int_t][Float_t]       = Error_t
    cube[AND][Int_t][Bool_t]        = Error_t
    cube[AND][Int_t][Tensor_t]      = Error_t

    cube[AND][Float_t][Int_t]       = Error_t
    cube[AND][Float_t][Float_t]     = Error_t
    cube[AND][Float_t][Bool_t]      = Error_t
    cube[AND][Float_t][Tensor_t]    = Error_t

    cube[AND][Bool_t][Int_t]        = Error_t
    cube[AND][Bool_t][Float_t]      = Error_t
    cube[AND][Bool_t][Bool_t]       = Bool_t
    cube[AND][Bool_t][Tensor_t]     = Error_t

    cube[AND][Tensor_t][Int_t]      = Error_t
    cube[AND][Tensor_t][Float_t]    = Error_t
    cube[AND][Tensor_t][Bool_t]     = Error_t
    cube[AND][Tensor_t][Tensor_t]   = Error_t



    cube[OR][Int_t][Int_t]          = Error_t
    cube[OR][Int_t][Float_t]        = Error_t
    cube[OR][Int_t][Bool_t]         = Error_t
    cube[OR][Int_t][Tensor_t]       = Error_t

    cube[OR][Float_t][Int_t]        = Error_t
    cube[OR][Float_t][Float_t]      = Error_t
    cube[OR][Float_t][Bool_t]       = Error_t
    cube[OR][Float_t][Tensor_t]     = Error_t

    cube[OR][Bool_t][Int_t]         = Error_t
    cube[OR][Bool_t][Float_t]       = Error_t
    cube[OR][Bool_t][Bool_t]        = Bool_t
    cube[OR][Bool_t][Tensor_t]      = Error_t

    cube[OR][Tensor_t][Int_t]       = Error_t
    cube[OR][Tensor_t][Float_t]     = Error_t
    cube[OR][Tensor_t][Bool_t]      = Error_t
    cube[OR][Tensor_t][Tensor_t]    = Error_t



    cube[NOT][Int_t][Int_t]         = Error_t
    cube[NOT][Int_t][Float_t]       = Error_t
    cube[NOT][Int_t][Bool_t]        = Error_t
    cube[NOT][Int_t][Tensor_t]      = Error_t

    cube[NOT][Float_t][Int_t]       = Error_t
    cube[NOT][Float_t][Float_t]     = Error_t
    cube[NOT][Float_t][Bool_t]      = Error_t
    cube[NOT][Float_t][Tensor_t]    = Error_t

    cube[NOT][Bool_t][Int_t]        = Error_t
    cube[NOT][Bool_t][Float_t]      = Error_t
    cube[NOT][Bool_t][Bool_t]       = Bool_t
    cube[NOT][Bool_t][Tensor_t]     = Error_t

    cube[NOT][Tensor_t][Int_t]      = Error_t
    cube[NOT][Tensor_t][Float_t]    = Error_t
    cube[NOT][Tensor_t][Bool_t]     = Error_t
    cube[NOT][Tensor_t][Tensor_t]   = Error_t



    cube[EQ][Int_t][Int_t]          = Bool_t
    cube[EQ][Int_t][Float_t]        = Bool_t
    cube[EQ][Int_t][Bool_t]         = Error_t
    cube[EQ][Int_t][Tensor_t]       = Error_t

    cube[EQ][Float_t][Int_t]        = Bool_t
    cube[EQ][Float_t][Float_t]      = Bool_t
    cube[EQ][Float_t][Bool_t]       = Error_t
    cube[EQ][Float_t][Tensor_t]     = Error_t

    cube[EQ][Bool_t][Int_t]         = Error_t
    cube[EQ][Bool_t][Float_t]       = Error_t
    cube[EQ][Bool_t][Bool_t]        = Bool_t
    cube[EQ][Bool_t][Tensor_t]      = Error_t

    cube[EQ][Tensor_t][Int_t]       = Error_t
    cube[EQ][Tensor_t][Float_t]     = Error_t
    cube[EQ][Tensor_t][Bool_t]      = Error_t
    cube[EQ][Tensor_t][Tensor_t]    = Bool_t



    cube[NEQ][Int_t][Int_t]         = Bool_t
    cube[NEQ][Int_t][Float_t]       = Bool_t
    cube[NEQ][Int_t][Bool_t]        = Error_t
    cube[NEQ][Int_t][Tensor_t]      = Error_t

    cube[NEQ][Float_t][Int_t]       = Bool_t
    cube[NEQ][Float_t][Float_t]     = Bool_t
    cube[NEQ][Float_t][Bool_t]      = Error_t
    cube[NEQ][Float_t][Tensor_t]    = Error_t

    cube[NEQ][Bool_t][Int_t]        = Error_t
    cube[NEQ][Bool_t][Float_t]      = Error_t
    cube[NEQ][Bool_t][Bool_t]       = Bool_t
    cube[NEQ][Bool_t][Tensor_t]     = Error_t

    cube[NEQ][Tensor_t][Int_t]      = Error_t
    cube[NEQ][Tensor_t][Float_t]    = Error_t
    cube[NEQ][Tensor_t][Bool_t]     = Error_t
    cube[NEQ][Tensor_t][Tensor_t]   = Bool_t



    cube[LES][Int_t][Int_t]         = Bool_t
    cube[LES][Int_t][Float_t]       = Bool_t
    cube[LES][Int_t][Bool_t]        = Error_t
    cube[LES][Int_t][Tensor_t]      = Error_t

    cube[LES][Float_t][Int_t]       = Bool_t
    cube[LES][Float_t][Float_t]     = Bool_t
    cube[LES][Float_t][Bool_t]      = Error_t
    cube[LES][Float_t][Tensor_t]    = Error_t

    cube[LES][Bool_t][Int_t]        = Error_t
    cube[LES][Bool_t][Float_t]      = Error_t
    cube[LES][Bool_t][Bool_t]       = Bool_t
    cube[LES][Bool_t][Tensor_t]     = Error_t

    cube[LES][Tensor_t][Int_t]      = Error_t
    cube[LES][Tensor_t][Float_t]    = Error_t
    cube[LES][Tensor_t][Bool_t]     = Error_t
    cube[LES][Tensor_t][Tensor_t]   = Bool_t



    cube[GRE][Int_t][Int_t]         = Bool_t
    cube[GRE][Int_t][Float_t]       = Bool_t
    cube[GRE][Int_t][Bool_t]        = Error_t
    cube[GRE][Int_t][Tensor_t]      = Error_t

    cube[GRE][Float_t][Int_t]       = Bool_t
    cube[GRE][Float_t][Float_t]     = Bool_t
    cube[GRE][Float_t][Bool_t]      = Error_t
    cube[GRE][Float_t][Tensor_t]    = Error_t

    cube[GRE][Bool_t][Int_t]        = Error_t
    cube[GRE][Bool_t][Float_t]      = Error_t
    cube[GRE][Bool_t][Bool_t]       = Bool_t
    cube[GRE][Bool_t][Tensor_t]     = Error_t

    cube[GRE][Tensor_t][Int_t]      = Error_t
    cube[GRE][Tensor_t][Float_t]    = Error_t
    cube[GRE][Tensor_t][Bool_t]     = Error_t
    cube[GRE][Tensor_t][Tensor_t]   = Bool_t



    cube[LEQ][Int_t][Int_t]         = Bool_t
    cube[LEQ][Int_t][Float_t]       = Bool_t
    cube[LEQ][Int_t][Bool_t]        = Error_t
    cube[LEQ][Int_t][Tensor_t]      = Error_t

    cube[LEQ][Float_t][Int_t]       = Bool_t
    cube[LEQ][Float_t][Float_t]     = Bool_t
    cube[LEQ][Float_t][Bool_t]      = Error_t
    cube[LEQ][Float_t][Tensor_t]    = Error_t

    cube[LEQ][Bool_t][Int_t]        = Error_t
    cube[LEQ][Bool_t][Float_t]      = Error_t
    cube[LEQ][Bool_t][Bool_t]       = Bool_t
    cube[LEQ][Bool_t][Tensor_t]     = Error_t

    cube[LEQ][Tensor_t][Int_t]      = Error_t
    cube[LEQ][Tensor_t][Float_t]    = Error_t
    cube[LEQ][Tensor_t][Bool_t]     = Error_t
    cube[LEQ][Tensor_t][Tensor_t]   = Bool_t



    cube[GEQ][Int_t][Int_t]         = Bool_t
    cube[GEQ][Int_t][Float_t]       = Bool_t
    cube[GEQ][Int_t][Bool_t]        = Error_t
    cube[GEQ][Int_t][Tensor_t]      = Error_t

    cube[GEQ][Float_t][Int_t]       = Bool_t
    cube[GEQ][Float_t][Float_t]     = Bool_t
    cube[GEQ][Float_t][Bool_t]      = Error_t
    cube[GEQ][Float_t][Tensor_t]    = Error_t

    cube[GEQ][Bool_t][Int_t]        = Error_t
    cube[GEQ][Bool_t][Float_t]      = Error_t
    cube[GEQ][Bool_t][Bool_t]       = Bool_t
    cube[GEQ][Bool_t][Tensor_t]     = Error_t

    cube[GEQ][Tensor_t][Int_t]      = Error_t
    cube[GEQ][Tensor_t][Float_t]    = Error_t
    cube[GEQ][Tensor_t][Bool_t]     = Error_t
    cube[GEQ][Tensor_t][Tensor_t]   = Bool_t
}
