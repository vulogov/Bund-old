// Code generated from Bund.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Bund

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBundListener is a complete listener for a parse tree produced by BundParser.
type BaseBundListener struct{}

var _ BundListener = &BaseBundListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBundListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBundListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBundListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBundListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpressions is called when production expressions is entered.
func (s *BaseBundListener) EnterExpressions(ctx *ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *BaseBundListener) ExitExpressions(ctx *ExpressionsContext) {}

// EnterRoot_term is called when production root_term is entered.
func (s *BaseBundListener) EnterRoot_term(ctx *Root_termContext) {}

// ExitRoot_term is called when production root_term is exited.
func (s *BaseBundListener) ExitRoot_term(ctx *Root_termContext) {}

// EnterNs is called when production ns is entered.
func (s *BaseBundListener) EnterNs(ctx *NsContext) {}

// ExitNs is called when production ns is exited.
func (s *BaseBundListener) ExitNs(ctx *NsContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseBundListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseBundListener) ExitBlock(ctx *BlockContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseBundListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseBundListener) ExitTerm(ctx *TermContext) {}

// EnterData is called when production data is entered.
func (s *BaseBundListener) EnterData(ctx *DataContext) {}

// ExitData is called when production data is exited.
func (s *BaseBundListener) ExitData(ctx *DataContext) {}

// EnterCall_term is called when production call_term is entered.
func (s *BaseBundListener) EnterCall_term(ctx *Call_termContext) {}

// ExitCall_term is called when production call_term is exited.
func (s *BaseBundListener) ExitCall_term(ctx *Call_termContext) {}

// EnterOperator_term is called when production operator_term is entered.
func (s *BaseBundListener) EnterOperator_term(ctx *Operator_termContext) {}

// ExitOperator_term is called when production operator_term is exited.
func (s *BaseBundListener) ExitOperator_term(ctx *Operator_termContext) {}

// EnterRef_call_term is called when production ref_call_term is entered.
func (s *BaseBundListener) EnterRef_call_term(ctx *Ref_call_termContext) {}

// ExitRef_call_term is called when production ref_call_term is exited.
func (s *BaseBundListener) ExitRef_call_term(ctx *Ref_call_termContext) {}

// EnterRef_operator_term is called when production ref_operator_term is entered.
func (s *BaseBundListener) EnterRef_operator_term(ctx *Ref_operator_termContext) {}

// ExitRef_operator_term is called when production ref_operator_term is exited.
func (s *BaseBundListener) ExitRef_operator_term(ctx *Ref_operator_termContext) {}

// EnterBoolean_term is called when production boolean_term is entered.
func (s *BaseBundListener) EnterBoolean_term(ctx *Boolean_termContext) {}

// ExitBoolean_term is called when production boolean_term is exited.
func (s *BaseBundListener) ExitBoolean_term(ctx *Boolean_termContext) {}

// EnterInteger_term is called when production integer_term is entered.
func (s *BaseBundListener) EnterInteger_term(ctx *Integer_termContext) {}

// ExitInteger_term is called when production integer_term is exited.
func (s *BaseBundListener) ExitInteger_term(ctx *Integer_termContext) {}

// EnterFloat_term is called when production float_term is entered.
func (s *BaseBundListener) EnterFloat_term(ctx *Float_termContext) {}

// ExitFloat_term is called when production float_term is exited.
func (s *BaseBundListener) ExitFloat_term(ctx *Float_termContext) {}

// EnterString_term is called when production string_term is entered.
func (s *BaseBundListener) EnterString_term(ctx *String_termContext) {}

// ExitString_term is called when production string_term is exited.
func (s *BaseBundListener) ExitString_term(ctx *String_termContext) {}

// EnterComplex_term is called when production complex_term is entered.
func (s *BaseBundListener) EnterComplex_term(ctx *Complex_termContext) {}

// ExitComplex_term is called when production complex_term is exited.
func (s *BaseBundListener) ExitComplex_term(ctx *Complex_termContext) {}

// EnterGlob_term is called when production glob_term is entered.
func (s *BaseBundListener) EnterGlob_term(ctx *Glob_termContext) {}

// ExitGlob_term is called when production glob_term is exited.
func (s *BaseBundListener) ExitGlob_term(ctx *Glob_termContext) {}

// EnterMode_term is called when production mode_term is entered.
func (s *BaseBundListener) EnterMode_term(ctx *Mode_termContext) {}

// ExitMode_term is called when production mode_term is exited.
func (s *BaseBundListener) ExitMode_term(ctx *Mode_termContext) {}

// EnterSeparate_term is called when production separate_term is entered.
func (s *BaseBundListener) EnterSeparate_term(ctx *Separate_termContext) {}

// ExitSeparate_term is called when production separate_term is exited.
func (s *BaseBundListener) ExitSeparate_term(ctx *Separate_termContext) {}

// EnterDatablock_term is called when production datablock_term is entered.
func (s *BaseBundListener) EnterDatablock_term(ctx *Datablock_termContext) {}

// ExitDatablock_term is called when production datablock_term is exited.
func (s *BaseBundListener) ExitDatablock_term(ctx *Datablock_termContext) {}

// EnterMatchblock_term is called when production matchblock_term is entered.
func (s *BaseBundListener) EnterMatchblock_term(ctx *Matchblock_termContext) {}

// ExitMatchblock_term is called when production matchblock_term is exited.
func (s *BaseBundListener) ExitMatchblock_term(ctx *Matchblock_termContext) {}

// EnterLogicblock_term is called when production logicblock_term is entered.
func (s *BaseBundListener) EnterLogicblock_term(ctx *Logicblock_termContext) {}

// ExitLogicblock_term is called when production logicblock_term is exited.
func (s *BaseBundListener) ExitLogicblock_term(ctx *Logicblock_termContext) {}

// EnterFunction_term is called when production function_term is entered.
func (s *BaseBundListener) EnterFunction_term(ctx *Function_termContext) {}

// ExitFunction_term is called when production function_term is exited.
func (s *BaseBundListener) ExitFunction_term(ctx *Function_termContext) {}

// EnterLambda_term is called when production lambda_term is entered.
func (s *BaseBundListener) EnterLambda_term(ctx *Lambda_termContext) {}

// ExitLambda_term is called when production lambda_term is exited.
func (s *BaseBundListener) ExitLambda_term(ctx *Lambda_termContext) {}

// EnterOperation_term is called when production operation_term is entered.
func (s *BaseBundListener) EnterOperation_term(ctx *Operation_termContext) {}

// ExitOperation_term is called when production operation_term is exited.
func (s *BaseBundListener) ExitOperation_term(ctx *Operation_termContext) {}

// EnterThing_term is called when production thing_term is entered.
func (s *BaseBundListener) EnterThing_term(ctx *Thing_termContext) {}

// ExitThing_term is called when production thing_term is exited.
func (s *BaseBundListener) ExitThing_term(ctx *Thing_termContext) {}
