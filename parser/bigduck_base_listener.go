// Code generated from BigDuck.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // BigDuck

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBigDuckListener is a complete listener for a parse tree produced by BigDuckParser.
type BaseBigDuckListener struct{}

var _ BigDuckListener = &BaseBigDuckListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBigDuckListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBigDuckListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBigDuckListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBigDuckListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseBigDuckListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseBigDuckListener) ExitProgram(ctx *ProgramContext) {}

// EnterVars_decl is called when production vars_decl is entered.
func (s *BaseBigDuckListener) EnterVars_decl(ctx *Vars_declContext) {}

// ExitVars_decl is called when production vars_decl is exited.
func (s *BaseBigDuckListener) ExitVars_decl(ctx *Vars_declContext) {}

// EnterVar_decl is called when production var_decl is entered.
func (s *BaseBigDuckListener) EnterVar_decl(ctx *Var_declContext) {}

// ExitVar_decl is called when production var_decl is exited.
func (s *BaseBigDuckListener) ExitVar_decl(ctx *Var_declContext) {}

// EnterNextVar is called when production nextVar is entered.
func (s *BaseBigDuckListener) EnterNextVar(ctx *NextVarContext) {}

// ExitNextVar is called when production nextVar is exited.
func (s *BaseBigDuckListener) ExitNextVar(ctx *NextVarContext) {}

// EnterNextVarDecl is called when production nextVarDecl is entered.
func (s *BaseBigDuckListener) EnterNextVarDecl(ctx *NextVarDeclContext) {}

// ExitNextVarDecl is called when production nextVarDecl is exited.
func (s *BaseBigDuckListener) ExitNextVarDecl(ctx *NextVarDeclContext) {}

// EnterVar_type is called when production var_type is entered.
func (s *BaseBigDuckListener) EnterVar_type(ctx *Var_typeContext) {}

// ExitVar_type is called when production var_type is exited.
func (s *BaseBigDuckListener) ExitVar_type(ctx *Var_typeContext) {}

// EnterScalar is called when production scalar is entered.
func (s *BaseBigDuckListener) EnterScalar(ctx *ScalarContext) {}

// ExitScalar is called when production scalar is exited.
func (s *BaseBigDuckListener) ExitScalar(ctx *ScalarContext) {}

// EnterTensor is called when production tensor is entered.
func (s *BaseBigDuckListener) EnterTensor(ctx *TensorContext) {}

// ExitTensor is called when production tensor is exited.
func (s *BaseBigDuckListener) ExitTensor(ctx *TensorContext) {}

// EnterDim is called when production dim is entered.
func (s *BaseBigDuckListener) EnterDim(ctx *DimContext) {}

// ExitDim is called when production dim is exited.
func (s *BaseBigDuckListener) ExitDim(ctx *DimContext) {}

// EnterNextDim is called when production nextDim is entered.
func (s *BaseBigDuckListener) EnterNextDim(ctx *NextDimContext) {}

// ExitNextDim is called when production nextDim is exited.
func (s *BaseBigDuckListener) ExitNextDim(ctx *NextDimContext) {}

// EnterProcs_decl is called when production procs_decl is entered.
func (s *BaseBigDuckListener) EnterProcs_decl(ctx *Procs_declContext) {}

// ExitProcs_decl is called when production procs_decl is exited.
func (s *BaseBigDuckListener) ExitProcs_decl(ctx *Procs_declContext) {}

// EnterProc_decl is called when production proc_decl is entered.
func (s *BaseBigDuckListener) EnterProc_decl(ctx *Proc_declContext) {}

// ExitProc_decl is called when production proc_decl is exited.
func (s *BaseBigDuckListener) ExitProc_decl(ctx *Proc_declContext) {}

// EnterSign is called when production sign is entered.
func (s *BaseBigDuckListener) EnterSign(ctx *SignContext) {}

// ExitSign is called when production sign is exited.
func (s *BaseBigDuckListener) ExitSign(ctx *SignContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseBigDuckListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseBigDuckListener) ExitArgs(ctx *ArgsContext) {}

// EnterNextTypes is called when production nextTypes is entered.
func (s *BaseBigDuckListener) EnterNextTypes(ctx *NextTypesContext) {}

// ExitNextTypes is called when production nextTypes is exited.
func (s *BaseBigDuckListener) ExitNextTypes(ctx *NextTypesContext) {}

// EnterNextArg is called when production nextArg is entered.
func (s *BaseBigDuckListener) EnterNextArg(ctx *NextArgContext) {}

// ExitNextArg is called when production nextArg is exited.
func (s *BaseBigDuckListener) ExitNextArg(ctx *NextArgContext) {}

// EnterRet_type is called when production ret_type is entered.
func (s *BaseBigDuckListener) EnterRet_type(ctx *Ret_typeContext) {}

// ExitRet_type is called when production ret_type is exited.
func (s *BaseBigDuckListener) ExitRet_type(ctx *Ret_typeContext) {}

// EnterBool_expr is called when production bool_expr is entered.
func (s *BaseBigDuckListener) EnterBool_expr(ctx *Bool_exprContext) {}

// ExitBool_expr is called when production bool_expr is exited.
func (s *BaseBigDuckListener) ExitBool_expr(ctx *Bool_exprContext) {}

// EnterNextBool is called when production nextBool is entered.
func (s *BaseBigDuckListener) EnterNextBool(ctx *NextBoolContext) {}

// ExitNextBool is called when production nextBool is exited.
func (s *BaseBigDuckListener) ExitNextBool(ctx *NextBoolContext) {}

// EnterAnd_expr is called when production and_expr is entered.
func (s *BaseBigDuckListener) EnterAnd_expr(ctx *And_exprContext) {}

// ExitAnd_expr is called when production and_expr is exited.
func (s *BaseBigDuckListener) ExitAnd_expr(ctx *And_exprContext) {}

// EnterNextAnd is called when production nextAnd is entered.
func (s *BaseBigDuckListener) EnterNextAnd(ctx *NextAndContext) {}

// ExitNextAnd is called when production nextAnd is exited.
func (s *BaseBigDuckListener) ExitNextAnd(ctx *NextAndContext) {}

// EnterNot_expr is called when production not_expr is entered.
func (s *BaseBigDuckListener) EnterNot_expr(ctx *Not_exprContext) {}

// ExitNot_expr is called when production not_expr is exited.
func (s *BaseBigDuckListener) ExitNot_expr(ctx *Not_exprContext) {}

// EnterBool_term is called when production bool_term is entered.
func (s *BaseBigDuckListener) EnterBool_term(ctx *Bool_termContext) {}

// ExitBool_term is called when production bool_term is exited.
func (s *BaseBigDuckListener) ExitBool_term(ctx *Bool_termContext) {}

// EnterRel_expr is called when production rel_expr is entered.
func (s *BaseBigDuckListener) EnterRel_expr(ctx *Rel_exprContext) {}

// ExitRel_expr is called when production rel_expr is exited.
func (s *BaseBigDuckListener) ExitRel_expr(ctx *Rel_exprContext) {}

// EnterOpRel is called when production opRel is entered.
func (s *BaseBigDuckListener) EnterOpRel(ctx *OpRelContext) {}

// ExitOpRel is called when production opRel is exited.
func (s *BaseBigDuckListener) ExitOpRel(ctx *OpRelContext) {}

// EnterRelOp is called when production relOp is entered.
func (s *BaseBigDuckListener) EnterRelOp(ctx *RelOpContext) {}

// ExitRelOp is called when production relOp is exited.
func (s *BaseBigDuckListener) ExitRelOp(ctx *RelOpContext) {}

// EnterNum_expr is called when production num_expr is entered.
func (s *BaseBigDuckListener) EnterNum_expr(ctx *Num_exprContext) {}

// ExitNum_expr is called when production num_expr is exited.
func (s *BaseBigDuckListener) ExitNum_expr(ctx *Num_exprContext) {}

// EnterNextSum is called when production nextSum is entered.
func (s *BaseBigDuckListener) EnterNextSum(ctx *NextSumContext) {}

// ExitNextSum is called when production nextSum is exited.
func (s *BaseBigDuckListener) ExitNextSum(ctx *NextSumContext) {}

// EnterProd_expr is called when production prod_expr is entered.
func (s *BaseBigDuckListener) EnterProd_expr(ctx *Prod_exprContext) {}

// ExitProd_expr is called when production prod_expr is exited.
func (s *BaseBigDuckListener) ExitProd_expr(ctx *Prod_exprContext) {}

// EnterNextProd is called when production nextProd is entered.
func (s *BaseBigDuckListener) EnterNextProd(ctx *NextProdContext) {}

// ExitNextProd is called when production nextProd is exited.
func (s *BaseBigDuckListener) ExitNextProd(ctx *NextProdContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseBigDuckListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseBigDuckListener) ExitFactor(ctx *FactorContext) {}

// EnterProc_call is called when production proc_call is entered.
func (s *BaseBigDuckListener) EnterProc_call(ctx *Proc_callContext) {}

// ExitProc_call is called when production proc_call is exited.
func (s *BaseBigDuckListener) ExitProc_call(ctx *Proc_callContext) {}

// EnterParam is called when production param is entered.
func (s *BaseBigDuckListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseBigDuckListener) ExitParam(ctx *ParamContext) {}

// EnterNextParam is called when production nextParam is entered.
func (s *BaseBigDuckListener) EnterNextParam(ctx *NextParamContext) {}

// ExitNextParam is called when production nextParam is exited.
func (s *BaseBigDuckListener) ExitNextParam(ctx *NextParamContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseBigDuckListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseBigDuckListener) ExitBlock(ctx *BlockContext) {}

// EnterStmts is called when production stmts is entered.
func (s *BaseBigDuckListener) EnterStmts(ctx *StmtsContext) {}

// ExitStmts is called when production stmts is exited.
func (s *BaseBigDuckListener) ExitStmts(ctx *StmtsContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseBigDuckListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseBigDuckListener) ExitStmt(ctx *StmtContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseBigDuckListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseBigDuckListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseBigDuckListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseBigDuckListener) ExitCondition(ctx *ConditionContext) {}

// EnterBodyCond is called when production bodyCond is entered.
func (s *BaseBigDuckListener) EnterBodyCond(ctx *BodyCondContext) {}

// ExitBodyCond is called when production bodyCond is exited.
func (s *BaseBigDuckListener) ExitBodyCond(ctx *BodyCondContext) {}

// EnterEndIfBlock is called when production endIfBlock is entered.
func (s *BaseBigDuckListener) EnterEndIfBlock(ctx *EndIfBlockContext) {}

// ExitEndIfBlock is called when production endIfBlock is exited.
func (s *BaseBigDuckListener) ExitEndIfBlock(ctx *EndIfBlockContext) {}

// EnterAlter is called when production alter is entered.
func (s *BaseBigDuckListener) EnterAlter(ctx *AlterContext) {}

// ExitAlter is called when production alter is exited.
func (s *BaseBigDuckListener) ExitAlter(ctx *AlterContext) {}

// EnterLoop_stmt is called when production loop_stmt is entered.
func (s *BaseBigDuckListener) EnterLoop_stmt(ctx *Loop_stmtContext) {}

// ExitLoop_stmt is called when production loop_stmt is exited.
func (s *BaseBigDuckListener) ExitLoop_stmt(ctx *Loop_stmtContext) {}

// EnterForStyle is called when production forStyle is entered.
func (s *BaseBigDuckListener) EnterForStyle(ctx *ForStyleContext) {}

// ExitForStyle is called when production forStyle is exited.
func (s *BaseBigDuckListener) ExitForStyle(ctx *ForStyleContext) {}

// EnterWhileStyle is called when production whileStyle is entered.
func (s *BaseBigDuckListener) EnterWhileStyle(ctx *WhileStyleContext) {}

// ExitWhileStyle is called when production whileStyle is exited.
func (s *BaseBigDuckListener) ExitWhileStyle(ctx *WhileStyleContext) {}

// EnterInfLoop is called when production infLoop is entered.
func (s *BaseBigDuckListener) EnterInfLoop(ctx *InfLoopContext) {}

// ExitInfLoop is called when production infLoop is exited.
func (s *BaseBigDuckListener) ExitInfLoop(ctx *InfLoopContext) {}

// EnterCtrl_flow is called when production ctrl_flow is entered.
func (s *BaseBigDuckListener) EnterCtrl_flow(ctx *Ctrl_flowContext) {}

// ExitCtrl_flow is called when production ctrl_flow is exited.
func (s *BaseBigDuckListener) ExitCtrl_flow(ctx *Ctrl_flowContext) {}

// EnterRet_stmt is called when production ret_stmt is entered.
func (s *BaseBigDuckListener) EnterRet_stmt(ctx *Ret_stmtContext) {}

// ExitRet_stmt is called when production ret_stmt is exited.
func (s *BaseBigDuckListener) ExitRet_stmt(ctx *Ret_stmtContext) {}
