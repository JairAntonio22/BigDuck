package structs

var Cube [opEnumCount][typeEnumCount][typeEnumCount]int

func InitCube() {
    Cube[ASG][Int_t][Int_t]         = Int_t
    Cube[ASG][Int_t][Float_t]       = Int_t
    Cube[ASG][Int_t][Bool_t]        = Error_t
    Cube[ASG][Int_t][Tensor_t]      = Error_t

    Cube[ASG][Float_t][Int_t]       = Float_t
    Cube[ASG][Float_t][Float_t]     = Float_t
    Cube[ASG][Float_t][Bool_t]      = Error_t
    Cube[ASG][Float_t][Tensor_t]    = Error_t

    Cube[ASG][Bool_t][Int_t]        = Error_t
    Cube[ASG][Bool_t][Float_t]      = Error_t
    Cube[ASG][Bool_t][Bool_t]       = Bool_t
    Cube[ASG][Bool_t][Tensor_t]     = Error_t

    Cube[ASG][Tensor_t][Int_t]      = Error_t
    Cube[ASG][Tensor_t][Float_t]    = Error_t
    Cube[ASG][Tensor_t][Bool_t]     = Error_t
    Cube[ASG][Tensor_t][Tensor_t]   = Tensor_t



    Cube[ADD][Int_t][Int_t]         = Int_t
    Cube[ADD][Int_t][Float_t]       = Float_t
    Cube[ADD][Int_t][Bool_t]        = Error_t
    Cube[ADD][Int_t][Tensor_t]      = Error_t

    Cube[ADD][Float_t][Int_t]       = Float_t
    Cube[ADD][Float_t][Float_t]     = Float_t
    Cube[ADD][Float_t][Bool_t]      = Error_t
    Cube[ADD][Float_t][Tensor_t]    = Error_t

    Cube[ADD][Bool_t][Int_t]        = Error_t
    Cube[ADD][Bool_t][Float_t]      = Error_t
    Cube[ADD][Bool_t][Bool_t]       = Error_t
    Cube[ADD][Bool_t][Tensor_t]     = Error_t

    Cube[ADD][Tensor_t][Int_t]      = Error_t
    Cube[ADD][Tensor_t][Float_t]    = Error_t
    Cube[ADD][Tensor_t][Bool_t]     = Error_t
    Cube[ADD][Tensor_t][Tensor_t]   = Tensor_t



    Cube[SUB][Int_t][Int_t]         = Int_t
    Cube[SUB][Int_t][Float_t]       = Float_t
    Cube[SUB][Int_t][Bool_t]        = Error_t
    Cube[SUB][Int_t][Tensor_t]      = Error_t

    Cube[SUB][Float_t][Int_t]       = Float_t
    Cube[SUB][Float_t][Float_t]     = Float_t
    Cube[SUB][Float_t][Bool_t]      = Error_t
    Cube[SUB][Float_t][Tensor_t]    = Error_t

    Cube[SUB][Bool_t][Int_t]        = Error_t
    Cube[SUB][Bool_t][Float_t]      = Error_t
    Cube[SUB][Bool_t][Bool_t]       = Error_t
    Cube[SUB][Bool_t][Tensor_t]     = Error_t

    Cube[SUB][Tensor_t][Int_t]      = Error_t
    Cube[SUB][Tensor_t][Float_t]    = Error_t
    Cube[SUB][Tensor_t][Bool_t]     = Error_t
    Cube[SUB][Tensor_t][Tensor_t]   = Tensor_t



    Cube[MUL][Int_t][Int_t]         = Int_t
    Cube[MUL][Int_t][Float_t]       = Float_t
    Cube[MUL][Int_t][Bool_t]        = Error_t
    Cube[MUL][Int_t][Tensor_t]      = Error_t

    Cube[MUL][Float_t][Int_t]       = Float_t
    Cube[MUL][Float_t][Float_t]     = Float_t
    Cube[MUL][Float_t][Bool_t]      = Error_t
    Cube[MUL][Float_t][Tensor_t]    = Error_t

    Cube[MUL][Bool_t][Int_t]        = Error_t
    Cube[MUL][Bool_t][Float_t]      = Error_t
    Cube[MUL][Bool_t][Bool_t]       = Error_t
    Cube[MUL][Bool_t][Tensor_t]     = Error_t

    Cube[MUL][Tensor_t][Int_t]      = Error_t
    Cube[MUL][Tensor_t][Float_t]    = Error_t
    Cube[MUL][Tensor_t][Bool_t]     = Error_t
    Cube[MUL][Tensor_t][Tensor_t]   = Tensor_t



    Cube[DIV][Int_t][Int_t]         = Int_t
    Cube[DIV][Int_t][Float_t]       = Float_t
    Cube[DIV][Int_t][Bool_t]        = Error_t
    Cube[DIV][Int_t][Tensor_t]      = Error_t

    Cube[DIV][Float_t][Int_t]       = Float_t
    Cube[DIV][Float_t][Float_t]     = Float_t
    Cube[DIV][Float_t][Bool_t]      = Error_t
    Cube[DIV][Float_t][Tensor_t]    = Error_t

    Cube[DIV][Bool_t][Int_t]        = Error_t
    Cube[DIV][Bool_t][Float_t]      = Error_t
    Cube[DIV][Bool_t][Bool_t]       = Error_t
    Cube[DIV][Bool_t][Tensor_t]     = Error_t

    Cube[DIV][Tensor_t][Int_t]      = Error_t
    Cube[DIV][Tensor_t][Float_t]    = Error_t
    Cube[DIV][Tensor_t][Bool_t]     = Error_t
    Cube[DIV][Tensor_t][Tensor_t]   = Tensor_t



    Cube[AND][Int_t][Int_t]         = Error_t
    Cube[AND][Int_t][Float_t]       = Error_t
    Cube[AND][Int_t][Bool_t]        = Error_t
    Cube[AND][Int_t][Tensor_t]      = Error_t

    Cube[AND][Float_t][Int_t]       = Error_t
    Cube[AND][Float_t][Float_t]     = Error_t
    Cube[AND][Float_t][Bool_t]      = Error_t
    Cube[AND][Float_t][Tensor_t]    = Error_t

    Cube[AND][Bool_t][Int_t]        = Error_t
    Cube[AND][Bool_t][Float_t]      = Error_t
    Cube[AND][Bool_t][Bool_t]       = Bool_t
    Cube[AND][Bool_t][Tensor_t]     = Error_t

    Cube[AND][Tensor_t][Int_t]      = Error_t
    Cube[AND][Tensor_t][Float_t]    = Error_t
    Cube[AND][Tensor_t][Bool_t]     = Error_t
    Cube[AND][Tensor_t][Tensor_t]   = Error_t



    Cube[OR][Int_t][Int_t]          = Error_t
    Cube[OR][Int_t][Float_t]        = Error_t
    Cube[OR][Int_t][Bool_t]         = Error_t
    Cube[OR][Int_t][Tensor_t]       = Error_t

    Cube[OR][Float_t][Int_t]        = Error_t
    Cube[OR][Float_t][Float_t]      = Error_t
    Cube[OR][Float_t][Bool_t]       = Error_t
    Cube[OR][Float_t][Tensor_t]     = Error_t

    Cube[OR][Bool_t][Int_t]         = Error_t
    Cube[OR][Bool_t][Float_t]       = Error_t
    Cube[OR][Bool_t][Bool_t]        = Bool_t
    Cube[OR][Bool_t][Tensor_t]      = Error_t

    Cube[OR][Tensor_t][Int_t]       = Error_t
    Cube[OR][Tensor_t][Float_t]     = Error_t
    Cube[OR][Tensor_t][Bool_t]      = Error_t
    Cube[OR][Tensor_t][Tensor_t]    = Error_t



    Cube[NOT][Int_t][Int_t]         = Error_t
    Cube[NOT][Int_t][Float_t]       = Error_t
    Cube[NOT][Int_t][Bool_t]        = Error_t
    Cube[NOT][Int_t][Tensor_t]      = Error_t

    Cube[NOT][Float_t][Int_t]       = Error_t
    Cube[NOT][Float_t][Float_t]     = Error_t
    Cube[NOT][Float_t][Bool_t]      = Error_t
    Cube[NOT][Float_t][Tensor_t]    = Error_t

    Cube[NOT][Bool_t][Int_t]        = Error_t
    Cube[NOT][Bool_t][Float_t]      = Error_t
    Cube[NOT][Bool_t][Bool_t]       = Bool_t
    Cube[NOT][Bool_t][Tensor_t]     = Error_t

    Cube[NOT][Tensor_t][Int_t]      = Error_t
    Cube[NOT][Tensor_t][Float_t]    = Error_t
    Cube[NOT][Tensor_t][Bool_t]     = Error_t
    Cube[NOT][Tensor_t][Tensor_t]   = Error_t



    Cube[EQ][Int_t][Int_t]          = Bool_t
    Cube[EQ][Int_t][Float_t]        = Bool_t
    Cube[EQ][Int_t][Bool_t]         = Error_t
    Cube[EQ][Int_t][Tensor_t]       = Error_t

    Cube[EQ][Float_t][Int_t]        = Bool_t
    Cube[EQ][Float_t][Float_t]      = Bool_t
    Cube[EQ][Float_t][Bool_t]       = Error_t
    Cube[EQ][Float_t][Tensor_t]     = Error_t

    Cube[EQ][Bool_t][Int_t]         = Error_t
    Cube[EQ][Bool_t][Float_t]       = Error_t
    Cube[EQ][Bool_t][Bool_t]        = Bool_t
    Cube[EQ][Bool_t][Tensor_t]      = Error_t

    Cube[EQ][Tensor_t][Int_t]       = Error_t
    Cube[EQ][Tensor_t][Float_t]     = Error_t
    Cube[EQ][Tensor_t][Bool_t]      = Error_t
    Cube[EQ][Tensor_t][Tensor_t]    = Bool_t



    Cube[NEQ][Int_t][Int_t]         = Bool_t
    Cube[NEQ][Int_t][Float_t]       = Bool_t
    Cube[NEQ][Int_t][Bool_t]        = Error_t
    Cube[NEQ][Int_t][Tensor_t]      = Error_t

    Cube[NEQ][Float_t][Int_t]       = Bool_t
    Cube[NEQ][Float_t][Float_t]     = Bool_t
    Cube[NEQ][Float_t][Bool_t]      = Error_t
    Cube[NEQ][Float_t][Tensor_t]    = Error_t

    Cube[NEQ][Bool_t][Int_t]        = Error_t
    Cube[NEQ][Bool_t][Float_t]      = Error_t
    Cube[NEQ][Bool_t][Bool_t]       = Bool_t
    Cube[NEQ][Bool_t][Tensor_t]     = Error_t

    Cube[NEQ][Tensor_t][Int_t]      = Error_t
    Cube[NEQ][Tensor_t][Float_t]    = Error_t
    Cube[NEQ][Tensor_t][Bool_t]     = Error_t
    Cube[NEQ][Tensor_t][Tensor_t]   = Bool_t



    Cube[LES][Int_t][Int_t]         = Bool_t
    Cube[LES][Int_t][Float_t]       = Bool_t
    Cube[LES][Int_t][Bool_t]        = Error_t
    Cube[LES][Int_t][Tensor_t]      = Error_t

    Cube[LES][Float_t][Int_t]       = Bool_t
    Cube[LES][Float_t][Float_t]     = Bool_t
    Cube[LES][Float_t][Bool_t]      = Error_t
    Cube[LES][Float_t][Tensor_t]    = Error_t

    Cube[LES][Bool_t][Int_t]        = Error_t
    Cube[LES][Bool_t][Float_t]      = Error_t
    Cube[LES][Bool_t][Bool_t]       = Bool_t
    Cube[LES][Bool_t][Tensor_t]     = Error_t

    Cube[LES][Tensor_t][Int_t]      = Error_t
    Cube[LES][Tensor_t][Float_t]    = Error_t
    Cube[LES][Tensor_t][Bool_t]     = Error_t
    Cube[LES][Tensor_t][Tensor_t]   = Bool_t



    Cube[GRE][Int_t][Int_t]         = Bool_t
    Cube[GRE][Int_t][Float_t]       = Bool_t
    Cube[GRE][Int_t][Bool_t]        = Error_t
    Cube[GRE][Int_t][Tensor_t]      = Error_t

    Cube[GRE][Float_t][Int_t]       = Bool_t
    Cube[GRE][Float_t][Float_t]     = Bool_t
    Cube[GRE][Float_t][Bool_t]      = Error_t
    Cube[GRE][Float_t][Tensor_t]    = Error_t

    Cube[GRE][Bool_t][Int_t]        = Error_t
    Cube[GRE][Bool_t][Float_t]      = Error_t
    Cube[GRE][Bool_t][Bool_t]       = Bool_t
    Cube[GRE][Bool_t][Tensor_t]     = Error_t

    Cube[GRE][Tensor_t][Int_t]      = Error_t
    Cube[GRE][Tensor_t][Float_t]    = Error_t
    Cube[GRE][Tensor_t][Bool_t]     = Error_t
    Cube[GRE][Tensor_t][Tensor_t]   = Bool_t



    Cube[LEQ][Int_t][Int_t]         = Bool_t
    Cube[LEQ][Int_t][Float_t]       = Bool_t
    Cube[LEQ][Int_t][Bool_t]        = Error_t
    Cube[LEQ][Int_t][Tensor_t]      = Error_t

    Cube[LEQ][Float_t][Int_t]       = Bool_t
    Cube[LEQ][Float_t][Float_t]     = Bool_t
    Cube[LEQ][Float_t][Bool_t]      = Error_t
    Cube[LEQ][Float_t][Tensor_t]    = Error_t

    Cube[LEQ][Bool_t][Int_t]        = Error_t
    Cube[LEQ][Bool_t][Float_t]      = Error_t
    Cube[LEQ][Bool_t][Bool_t]       = Bool_t
    Cube[LEQ][Bool_t][Tensor_t]     = Error_t

    Cube[LEQ][Tensor_t][Int_t]      = Error_t
    Cube[LEQ][Tensor_t][Float_t]    = Error_t
    Cube[LEQ][Tensor_t][Bool_t]     = Error_t
    Cube[LEQ][Tensor_t][Tensor_t]   = Bool_t



    Cube[GEQ][Int_t][Int_t]         = Bool_t
    Cube[GEQ][Int_t][Float_t]       = Bool_t
    Cube[GEQ][Int_t][Bool_t]        = Error_t
    Cube[GEQ][Int_t][Tensor_t]      = Error_t

    Cube[GEQ][Float_t][Int_t]       = Bool_t
    Cube[GEQ][Float_t][Float_t]     = Bool_t
    Cube[GEQ][Float_t][Bool_t]      = Error_t
    Cube[GEQ][Float_t][Tensor_t]    = Error_t

    Cube[GEQ][Bool_t][Int_t]        = Error_t
    Cube[GEQ][Bool_t][Float_t]      = Error_t
    Cube[GEQ][Bool_t][Bool_t]       = Bool_t
    Cube[GEQ][Bool_t][Tensor_t]     = Error_t

    Cube[GEQ][Tensor_t][Int_t]      = Error_t
    Cube[GEQ][Tensor_t][Float_t]    = Error_t
    Cube[GEQ][Tensor_t][Bool_t]     = Error_t
    Cube[GEQ][Tensor_t][Tensor_t]   = Bool_t
}
