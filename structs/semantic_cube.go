package structs

var Cube [opEnumCount][typeEnumCount][typeEnumCount]int

func InitCube() {
    Cube[ASG][Int_t][Int_t]         = Int_t
    Cube[ASG][Int_t][Float_t]       = Int_t
    Cube[ASG][Int_t][Bool_t]        = Error_t

    Cube[ASG][Float_t][Int_t]       = Float_t
    Cube[ASG][Float_t][Float_t]     = Float_t
    Cube[ASG][Float_t][Bool_t]      = Error_t

    Cube[ASG][Bool_t][Int_t]        = Error_t
    Cube[ASG][Bool_t][Float_t]      = Error_t
    Cube[ASG][Bool_t][Bool_t]       = Bool_t



    Cube[ADD][Int_t][Int_t]         = Int_t
    Cube[ADD][Int_t][Float_t]       = Float_t
    Cube[ADD][Int_t][Bool_t]        = Error_t

    Cube[ADD][Float_t][Int_t]       = Float_t
    Cube[ADD][Float_t][Float_t]     = Float_t
    Cube[ADD][Float_t][Bool_t]      = Error_t

    Cube[ADD][Bool_t][Int_t]        = Error_t
    Cube[ADD][Bool_t][Float_t]      = Error_t
    Cube[ADD][Bool_t][Bool_t]       = Error_t



    Cube[SUB][Int_t][Int_t]         = Int_t
    Cube[SUB][Int_t][Float_t]       = Float_t
    Cube[SUB][Int_t][Bool_t]        = Error_t

    Cube[SUB][Float_t][Int_t]       = Float_t
    Cube[SUB][Float_t][Float_t]     = Float_t
    Cube[SUB][Float_t][Bool_t]      = Error_t

    Cube[SUB][Bool_t][Int_t]        = Error_t
    Cube[SUB][Bool_t][Float_t]      = Error_t
    Cube[SUB][Bool_t][Bool_t]       = Error_t



    Cube[MUL][Int_t][Int_t]         = Int_t
    Cube[MUL][Int_t][Float_t]       = Float_t
    Cube[MUL][Int_t][Bool_t]        = Error_t

    Cube[MUL][Float_t][Int_t]       = Float_t
    Cube[MUL][Float_t][Float_t]     = Float_t
    Cube[MUL][Float_t][Bool_t]      = Error_t

    Cube[MUL][Bool_t][Int_t]        = Error_t
    Cube[MUL][Bool_t][Float_t]      = Error_t
    Cube[MUL][Bool_t][Bool_t]       = Error_t



    Cube[DIV][Int_t][Int_t]         = Int_t
    Cube[DIV][Int_t][Float_t]       = Float_t
    Cube[DIV][Int_t][Bool_t]        = Error_t

    Cube[DIV][Float_t][Int_t]       = Float_t
    Cube[DIV][Float_t][Float_t]     = Float_t
    Cube[DIV][Float_t][Bool_t]      = Error_t

    Cube[DIV][Bool_t][Int_t]        = Error_t
    Cube[DIV][Bool_t][Float_t]      = Error_t
    Cube[DIV][Bool_t][Bool_t]       = Error_t



    Cube[AND][Int_t][Int_t]         = Error_t
    Cube[AND][Int_t][Float_t]       = Error_t
    Cube[AND][Int_t][Bool_t]        = Error_t

    Cube[AND][Float_t][Int_t]       = Error_t
    Cube[AND][Float_t][Float_t]     = Error_t
    Cube[AND][Float_t][Bool_t]      = Error_t

    Cube[AND][Bool_t][Int_t]        = Error_t
    Cube[AND][Bool_t][Float_t]      = Error_t
    Cube[AND][Bool_t][Bool_t]       = Bool_t



    Cube[OR][Int_t][Int_t]          = Error_t
    Cube[OR][Int_t][Float_t]        = Error_t
    Cube[OR][Int_t][Bool_t]         = Error_t

    Cube[OR][Float_t][Int_t]        = Error_t
    Cube[OR][Float_t][Float_t]      = Error_t
    Cube[OR][Float_t][Bool_t]       = Error_t

    Cube[OR][Bool_t][Int_t]         = Error_t
    Cube[OR][Bool_t][Float_t]       = Error_t
    Cube[OR][Bool_t][Bool_t]        = Bool_t



    Cube[NOT][Int_t][Int_t]         = Error_t
    Cube[NOT][Int_t][Float_t]       = Error_t
    Cube[NOT][Int_t][Bool_t]        = Error_t

    Cube[NOT][Float_t][Int_t]       = Error_t
    Cube[NOT][Float_t][Float_t]     = Error_t
    Cube[NOT][Float_t][Bool_t]      = Error_t

    Cube[NOT][Bool_t][Int_t]        = Error_t
    Cube[NOT][Bool_t][Float_t]      = Error_t
    Cube[NOT][Bool_t][Bool_t]       = Bool_t



    Cube[EQ][Int_t][Int_t]          = Bool_t
    Cube[EQ][Int_t][Float_t]        = Bool_t
    Cube[EQ][Int_t][Bool_t]         = Error_t

    Cube[EQ][Float_t][Int_t]        = Bool_t
    Cube[EQ][Float_t][Float_t]      = Bool_t
    Cube[EQ][Float_t][Bool_t]       = Error_t

    Cube[EQ][Bool_t][Int_t]         = Error_t
    Cube[EQ][Bool_t][Float_t]       = Error_t
    Cube[EQ][Bool_t][Bool_t]        = Bool_t



    Cube[NEQ][Int_t][Int_t]         = Bool_t
    Cube[NEQ][Int_t][Float_t]       = Bool_t
    Cube[NEQ][Int_t][Bool_t]        = Error_t

    Cube[NEQ][Float_t][Int_t]       = Bool_t
    Cube[NEQ][Float_t][Float_t]     = Bool_t
    Cube[NEQ][Float_t][Bool_t]      = Error_t

    Cube[NEQ][Bool_t][Int_t]        = Error_t
    Cube[NEQ][Bool_t][Float_t]      = Error_t
    Cube[NEQ][Bool_t][Bool_t]       = Bool_t



    Cube[LES][Int_t][Int_t]         = Bool_t
    Cube[LES][Int_t][Float_t]       = Bool_t
    Cube[LES][Int_t][Bool_t]        = Error_t

    Cube[LES][Float_t][Int_t]       = Bool_t
    Cube[LES][Float_t][Float_t]     = Bool_t
    Cube[LES][Float_t][Bool_t]      = Error_t

    Cube[LES][Bool_t][Int_t]        = Error_t
    Cube[LES][Bool_t][Float_t]      = Error_t
    Cube[LES][Bool_t][Bool_t]       = Bool_t



    Cube[GRE][Int_t][Int_t]         = Bool_t
    Cube[GRE][Int_t][Float_t]       = Bool_t
    Cube[GRE][Int_t][Bool_t]        = Error_t

    Cube[GRE][Float_t][Int_t]       = Bool_t
    Cube[GRE][Float_t][Float_t]     = Bool_t
    Cube[GRE][Float_t][Bool_t]      = Error_t

    Cube[GRE][Bool_t][Int_t]        = Error_t
    Cube[GRE][Bool_t][Float_t]      = Error_t
    Cube[GRE][Bool_t][Bool_t]       = Bool_t



    Cube[LEQ][Int_t][Int_t]         = Bool_t
    Cube[LEQ][Int_t][Float_t]       = Bool_t
    Cube[LEQ][Int_t][Bool_t]        = Error_t

    Cube[LEQ][Float_t][Int_t]       = Bool_t
    Cube[LEQ][Float_t][Float_t]     = Bool_t
    Cube[LEQ][Float_t][Bool_t]      = Error_t

    Cube[LEQ][Bool_t][Int_t]        = Error_t
    Cube[LEQ][Bool_t][Float_t]      = Error_t
    Cube[LEQ][Bool_t][Bool_t]       = Bool_t



    Cube[GEQ][Int_t][Int_t]         = Bool_t
    Cube[GEQ][Int_t][Float_t]       = Bool_t
    Cube[GEQ][Int_t][Bool_t]        = Error_t

    Cube[GEQ][Float_t][Int_t]       = Bool_t
    Cube[GEQ][Float_t][Float_t]     = Bool_t
    Cube[GEQ][Float_t][Bool_t]      = Error_t

    Cube[GEQ][Bool_t][Int_t]        = Error_t
    Cube[GEQ][Bool_t][Float_t]      = Error_t
    Cube[GEQ][Bool_t][Bool_t]       = Bool_t



    Cube[SIN][Int_t][Int_t]         = Float_t
    Cube[SIN][Int_t][Float_t]       = Error_t
    Cube[SIN][Int_t][Bool_t]        = Error_t

    Cube[SIN][Float_t][Int_t]       = Error_t
    Cube[SIN][Float_t][Float_t]     = Float_t
    Cube[SIN][Float_t][Bool_t]      = Error_t

    Cube[SIN][Bool_t][Int_t]        = Error_t
    Cube[SIN][Bool_t][Float_t]      = Error_t
    Cube[SIN][Bool_t][Bool_t]       = Error_t



    Cube[ASIN][Int_t][Int_t]         = Float_t
    Cube[ASIN][Int_t][Float_t]       = Error_t
    Cube[ASIN][Int_t][Bool_t]        = Error_t

    Cube[ASIN][Float_t][Int_t]       = Error_t
    Cube[ASIN][Float_t][Float_t]     = Float_t
    Cube[ASIN][Float_t][Bool_t]      = Error_t

    Cube[ASIN][Bool_t][Int_t]        = Error_t
    Cube[ASIN][Bool_t][Float_t]      = Error_t
    Cube[ASIN][Bool_t][Bool_t]       = Error_t



    Cube[COS][Int_t][Int_t]         = Float_t
    Cube[COS][Int_t][Float_t]       = Error_t
    Cube[COS][Int_t][Bool_t]        = Error_t

    Cube[COS][Float_t][Int_t]       = Error_t
    Cube[COS][Float_t][Float_t]     = Float_t
    Cube[COS][Float_t][Bool_t]      = Error_t

    Cube[COS][Bool_t][Int_t]        = Error_t
    Cube[COS][Bool_t][Float_t]      = Error_t
    Cube[COS][Bool_t][Bool_t]       = Error_t



    Cube[ACOS][Int_t][Int_t]         = Float_t
    Cube[ACOS][Int_t][Float_t]       = Error_t
    Cube[ACOS][Int_t][Bool_t]        = Error_t

    Cube[ACOS][Float_t][Int_t]       = Error_t
    Cube[ACOS][Float_t][Float_t]     = Float_t
    Cube[ACOS][Float_t][Bool_t]      = Error_t

    Cube[ACOS][Bool_t][Int_t]        = Error_t
    Cube[ACOS][Bool_t][Float_t]      = Error_t
    Cube[ACOS][Bool_t][Bool_t]       = Error_t



    Cube[TAN][Int_t][Int_t]         = Float_t
    Cube[TAN][Int_t][Float_t]       = Error_t
    Cube[TAN][Int_t][Bool_t]        = Error_t

    Cube[TAN][Float_t][Int_t]       = Error_t
    Cube[TAN][Float_t][Float_t]     = Float_t
    Cube[TAN][Float_t][Bool_t]      = Error_t

    Cube[TAN][Bool_t][Int_t]        = Error_t
    Cube[TAN][Bool_t][Float_t]      = Error_t
    Cube[TAN][Bool_t][Bool_t]       = Error_t



    Cube[ATAN][Int_t][Int_t]         = Float_t
    Cube[ATAN][Int_t][Float_t]       = Error_t
    Cube[ATAN][Int_t][Bool_t]        = Error_t

    Cube[ATAN][Float_t][Int_t]       = Error_t
    Cube[ATAN][Float_t][Float_t]     = Float_t
    Cube[ATAN][Float_t][Bool_t]      = Error_t

    Cube[ATAN][Bool_t][Int_t]        = Error_t
    Cube[ATAN][Bool_t][Float_t]      = Error_t
    Cube[ATAN][Bool_t][Bool_t]       = Error_t



    Cube[ATAN2][Int_t][Int_t]        = Float_t
    Cube[ATAN2][Int_t][Float_t]      = Float_t
    Cube[ATAN2][Int_t][Bool_t]       = Error_t

    Cube[ATAN2][Float_t][Int_t]      = Float_t
    Cube[ATAN2][Float_t][Float_t]    = Float_t
    Cube[ATAN2][Float_t][Bool_t]     = Error_t

    Cube[ATAN2][Bool_t][Int_t]       = Error_t
    Cube[ATAN2][Bool_t][Float_t]     = Error_t
    Cube[ATAN2][Bool_t][Bool_t]      = Error_t



    Cube[EXP][Int_t][Int_t]         = Float_t
    Cube[EXP][Int_t][Float_t]       = Float_t
    Cube[EXP][Int_t][Bool_t]        = Error_t

    Cube[EXP][Float_t][Int_t]       = Float_t
    Cube[EXP][Float_t][Float_t]     = Float_t
    Cube[EXP][Float_t][Bool_t]      = Error_t

    Cube[EXP][Bool_t][Int_t]        = Error_t
    Cube[EXP][Bool_t][Float_t]      = Error_t
    Cube[EXP][Bool_t][Bool_t]       = Error_t



    Cube[LN][Int_t][Int_t]          = Float_t
    Cube[LN][Int_t][Float_t]        = Error_t
    Cube[LN][Int_t][Bool_t]         = Error_t

    Cube[LN][Float_t][Int_t]        = Error_t
    Cube[LN][Float_t][Float_t]      = Float_t
    Cube[LN][Float_t][Bool_t]       = Error_t

    Cube[LN][Bool_t][Int_t]         = Error_t
    Cube[LN][Bool_t][Float_t]       = Error_t
    Cube[LN][Bool_t][Bool_t]        = Error_t



    Cube[SQRT][Int_t][Int_t]        = Float_t
    Cube[SQRT][Int_t][Float_t]      = Float_t
    Cube[SQRT][Int_t][Bool_t]       = Error_t

    Cube[SQRT][Float_t][Int_t]      = Float_t
    Cube[SQRT][Float_t][Float_t]    = Float_t
    Cube[SQRT][Float_t][Bool_t]     = Error_t

    Cube[SQRT][Bool_t][Int_t]       = Error_t
    Cube[SQRT][Bool_t][Float_t]     = Error_t
    Cube[SQRT][Bool_t][Bool_t]      = Error_t



    Cube[POW][Int_t][Int_t]         = Float_t
    Cube[POW][Int_t][Float_t]       = Float_t
    Cube[POW][Int_t][Bool_t]        = Error_t

    Cube[POW][Float_t][Int_t]       = Float_t
    Cube[POW][Float_t][Float_t]     = Float_t
    Cube[POW][Float_t][Bool_t]      = Error_t

    Cube[POW][Bool_t][Int_t]        = Error_t
    Cube[POW][Bool_t][Float_t]      = Error_t
    Cube[POW][Bool_t][Bool_t]       = Error_t



    Cube[LOG][Int_t][Int_t]         = Float_t
    Cube[LOG][Int_t][Float_t]       = Float_t
    Cube[LOG][Int_t][Bool_t]        = Error_t

    Cube[LOG][Float_t][Int_t]       = Float_t
    Cube[LOG][Float_t][Float_t]     = Float_t
    Cube[LOG][Float_t][Bool_t]      = Error_t

    Cube[LOG][Bool_t][Int_t]        = Error_t
    Cube[LOG][Bool_t][Float_t]      = Error_t
    Cube[LOG][Bool_t][Bool_t]       = Error_t



    Cube[MOD][Int_t][Int_t]         = Int_t
    Cube[MOD][Int_t][Float_t]       = Float_t
    Cube[MOD][Int_t][Bool_t]        = Error_t

    Cube[MOD][Float_t][Int_t]       = Float_t
    Cube[MOD][Float_t][Float_t]     = Float_t
    Cube[MOD][Float_t][Bool_t]      = Error_t

    Cube[MOD][Bool_t][Int_t]        = Error_t
    Cube[MOD][Bool_t][Float_t]      = Error_t
    Cube[MOD][Bool_t][Bool_t]       = Error_t



    Cube[ABS][Int_t][Int_t]         = Float_t
    Cube[ABS][Int_t][Float_t]       = Error_t
    Cube[ABS][Int_t][Bool_t]        = Error_t

    Cube[ABS][Float_t][Int_t]       = Error_t
    Cube[ABS][Float_t][Float_t]     = Float_t
    Cube[ABS][Float_t][Bool_t]      = Error_t

    Cube[ABS][Bool_t][Int_t]        = Error_t
    Cube[ABS][Bool_t][Float_t]      = Error_t
    Cube[ABS][Bool_t][Bool_t]       = Error_t



    Cube[CEIL][Int_t][Int_t]        = Float_t
    Cube[CEIL][Int_t][Float_t]      = Error_t
    Cube[CEIL][Int_t][Bool_t]       = Error_t

    Cube[CEIL][Float_t][Int_t]      = Error_t
    Cube[CEIL][Float_t][Float_t]    = Float_t
    Cube[CEIL][Float_t][Bool_t]     = Error_t

    Cube[CEIL][Bool_t][Int_t]       = Error_t
    Cube[CEIL][Bool_t][Float_t]     = Error_t
    Cube[CEIL][Bool_t][Bool_t]      = Error_t



    Cube[FLOOR][Int_t][Int_t]       = Float_t
    Cube[FLOOR][Int_t][Float_t]     = Error_t
    Cube[FLOOR][Int_t][Bool_t]      = Error_t

    Cube[FLOOR][Float_t][Int_t]     = Error_t
    Cube[FLOOR][Float_t][Float_t]   = Float_t
    Cube[FLOOR][Float_t][Bool_t]    = Error_t

    Cube[FLOOR][Bool_t][Int_t]      = Error_t
    Cube[FLOOR][Bool_t][Float_t]    = Error_t
    Cube[FLOOR][Bool_t][Bool_t]     = Error_t



    Cube[MEAN][Int_t][Int_t]        = Float_t
    Cube[MEAN][Int_t][Float_t]      = Float_t
    Cube[MEAN][Int_t][Bool_t]       = Error_t

    Cube[MEAN][Float_t][Int_t]      = Float_t
    Cube[MEAN][Float_t][Float_t]    = Float_t
    Cube[MEAN][Float_t][Bool_t]     = Error_t

    Cube[MEAN][Bool_t][Int_t]       = Error_t
    Cube[MEAN][Bool_t][Float_t]     = Error_t
    Cube[MEAN][Bool_t][Bool_t]      = Error_t



    Cube[MEDIAN][Int_t][Int_t]     = Float_t
    Cube[MEDIAN][Int_t][Float_t]   = Float_t
    Cube[MEDIAN][Int_t][Bool_t]    = Error_t

    Cube[MEDIAN][Float_t][Int_t]   = Float_t
    Cube[MEDIAN][Float_t][Float_t] = Float_t
    Cube[MEDIAN][Float_t][Bool_t]  = Error_t

    Cube[MEDIAN][Bool_t][Int_t]    = Error_t
    Cube[MEDIAN][Bool_t][Float_t]  = Error_t
    Cube[MEDIAN][Bool_t][Bool_t]   = Error_t



    Cube[MODE][Int_t][Int_t]       = Float_t
    Cube[MODE][Int_t][Float_t]     = Float_t
    Cube[MODE][Int_t][Bool_t]      = Error_t

    Cube[MODE][Float_t][Int_t]     = Float_t
    Cube[MODE][Float_t][Float_t]   = Float_t
    Cube[MODE][Float_t][Bool_t]    = Error_t

    Cube[MODE][Bool_t][Int_t]      = Error_t
    Cube[MODE][Bool_t][Float_t]    = Error_t
    Cube[MODE][Bool_t][Bool_t]     = Error_t
}
