// Code generated from BigDuck.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // BigDuck

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BigDuckListener is a complete listener for a parse tree produced by BigDuckParser.
type BigDuckListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterVars_decl is called when entering the vars_decl production.
	EnterVars_decl(c *Vars_declContext)

	// EnterVar_decl is called when entering the var_decl production.
	EnterVar_decl(c *Var_declContext)

	// EnterNextVar is called when entering the nextVar production.
	EnterNextVar(c *NextVarContext)

	// EnterNextVarDecl is called when entering the nextVarDecl production.
	EnterNextVarDecl(c *NextVarDeclContext)

	// EnterVar_type is called when entering the var_type production.
	EnterVar_type(c *Var_typeContext)

	// EnterScalar is called when entering the scalar production.
	EnterScalar(c *ScalarContext)

	// EnterTensor is called when entering the tensor production.
	EnterTensor(c *TensorContext)

	// EnterDim is called when entering the dim production.
	EnterDim(c *DimContext)

	// EnterNextDim is called when entering the nextDim production.
	EnterNextDim(c *NextDimContext)

	// EnterProcs_decl is called when entering the procs_decl production.
	EnterProcs_decl(c *Procs_declContext)

	// EnterProc_decl is called when entering the proc_decl production.
	EnterProc_decl(c *Proc_declContext)

	// EnterSign is called when entering the sign production.
	EnterSign(c *SignContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterNextTypes is called when entering the nextTypes production.
	EnterNextTypes(c *NextTypesContext)

	// EnterNextArg is called when entering the nextArg production.
	EnterNextArg(c *NextArgContext)

	// EnterRet_type is called when entering the ret_type production.
	EnterRet_type(c *Ret_typeContext)

	// EnterBool_expr is called when entering the bool_expr production.
	EnterBool_expr(c *Bool_exprContext)

	// EnterNextBool is called when entering the nextBool production.
	EnterNextBool(c *NextBoolContext)

	// EnterAnd_expr is called when entering the and_expr production.
	EnterAnd_expr(c *And_exprContext)

	// EnterNextAnd is called when entering the nextAnd production.
	EnterNextAnd(c *NextAndContext)

	// EnterNot_expr is called when entering the not_expr production.
	EnterNot_expr(c *Not_exprContext)

	// EnterBool_term is called when entering the bool_term production.
	EnterBool_term(c *Bool_termContext)

	// EnterRel_expr is called when entering the rel_expr production.
	EnterRel_expr(c *Rel_exprContext)

	// EnterOpRel is called when entering the opRel production.
	EnterOpRel(c *OpRelContext)

	// EnterRelOp is called when entering the relOp production.
	EnterRelOp(c *RelOpContext)

	// EnterNum_expr is called when entering the num_expr production.
	EnterNum_expr(c *Num_exprContext)

	// EnterNextSum is called when entering the nextSum production.
	EnterNextSum(c *NextSumContext)

	// EnterProd_expr is called when entering the prod_expr production.
	EnterProd_expr(c *Prod_exprContext)

	// EnterNextProd is called when entering the nextProd production.
	EnterNextProd(c *NextProdContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterProc_call is called when entering the proc_call production.
	EnterProc_call(c *Proc_callContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterNextParam is called when entering the nextParam production.
	EnterNextParam(c *NextParamContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStmts is called when entering the stmts production.
	EnterStmts(c *StmtsContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterBodyCond is called when entering the bodyCond production.
	EnterBodyCond(c *BodyCondContext)

	// EnterAlter is called when entering the alter production.
	EnterAlter(c *AlterContext)

	// EnterLoop_stmt is called when entering the loop_stmt production.
	EnterLoop_stmt(c *Loop_stmtContext)

	// EnterForNotation is called when entering the forNotation production.
	EnterForNotation(c *ForNotationContext)

	// EnterCtrl_flow is called when entering the ctrl_flow production.
	EnterCtrl_flow(c *Ctrl_flowContext)

	// EnterRet_stmt is called when entering the ret_stmt production.
	EnterRet_stmt(c *Ret_stmtContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitVars_decl is called when exiting the vars_decl production.
	ExitVars_decl(c *Vars_declContext)

	// ExitVar_decl is called when exiting the var_decl production.
	ExitVar_decl(c *Var_declContext)

	// ExitNextVar is called when exiting the nextVar production.
	ExitNextVar(c *NextVarContext)

	// ExitNextVarDecl is called when exiting the nextVarDecl production.
	ExitNextVarDecl(c *NextVarDeclContext)

	// ExitVar_type is called when exiting the var_type production.
	ExitVar_type(c *Var_typeContext)

	// ExitScalar is called when exiting the scalar production.
	ExitScalar(c *ScalarContext)

	// ExitTensor is called when exiting the tensor production.
	ExitTensor(c *TensorContext)

	// ExitDim is called when exiting the dim production.
	ExitDim(c *DimContext)

	// ExitNextDim is called when exiting the nextDim production.
	ExitNextDim(c *NextDimContext)

	// ExitProcs_decl is called when exiting the procs_decl production.
	ExitProcs_decl(c *Procs_declContext)

	// ExitProc_decl is called when exiting the proc_decl production.
	ExitProc_decl(c *Proc_declContext)

	// ExitSign is called when exiting the sign production.
	ExitSign(c *SignContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitNextTypes is called when exiting the nextTypes production.
	ExitNextTypes(c *NextTypesContext)

	// ExitNextArg is called when exiting the nextArg production.
	ExitNextArg(c *NextArgContext)

	// ExitRet_type is called when exiting the ret_type production.
	ExitRet_type(c *Ret_typeContext)

	// ExitBool_expr is called when exiting the bool_expr production.
	ExitBool_expr(c *Bool_exprContext)

	// ExitNextBool is called when exiting the nextBool production.
	ExitNextBool(c *NextBoolContext)

	// ExitAnd_expr is called when exiting the and_expr production.
	ExitAnd_expr(c *And_exprContext)

	// ExitNextAnd is called when exiting the nextAnd production.
	ExitNextAnd(c *NextAndContext)

	// ExitNot_expr is called when exiting the not_expr production.
	ExitNot_expr(c *Not_exprContext)

	// ExitBool_term is called when exiting the bool_term production.
	ExitBool_term(c *Bool_termContext)

	// ExitRel_expr is called when exiting the rel_expr production.
	ExitRel_expr(c *Rel_exprContext)

	// ExitOpRel is called when exiting the opRel production.
	ExitOpRel(c *OpRelContext)

	// ExitRelOp is called when exiting the relOp production.
	ExitRelOp(c *RelOpContext)

	// ExitNum_expr is called when exiting the num_expr production.
	ExitNum_expr(c *Num_exprContext)

	// ExitNextSum is called when exiting the nextSum production.
	ExitNextSum(c *NextSumContext)

	// ExitProd_expr is called when exiting the prod_expr production.
	ExitProd_expr(c *Prod_exprContext)

	// ExitNextProd is called when exiting the nextProd production.
	ExitNextProd(c *NextProdContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitProc_call is called when exiting the proc_call production.
	ExitProc_call(c *Proc_callContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitNextParam is called when exiting the nextParam production.
	ExitNextParam(c *NextParamContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStmts is called when exiting the stmts production.
	ExitStmts(c *StmtsContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitBodyCond is called when exiting the bodyCond production.
	ExitBodyCond(c *BodyCondContext)

	// ExitAlter is called when exiting the alter production.
	ExitAlter(c *AlterContext)

	// ExitLoop_stmt is called when exiting the loop_stmt production.
	ExitLoop_stmt(c *Loop_stmtContext)

	// ExitForNotation is called when exiting the forNotation production.
	ExitForNotation(c *ForNotationContext)

	// ExitCtrl_flow is called when exiting the ctrl_flow production.
	ExitCtrl_flow(c *Ctrl_flowContext)

	// ExitRet_stmt is called when exiting the ret_stmt production.
	ExitRet_stmt(c *Ret_stmtContext)
}
