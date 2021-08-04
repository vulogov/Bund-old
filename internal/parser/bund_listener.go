// Code generated from Bund.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Bund

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BundListener is a complete listener for a parse tree produced by BundParser.
type BundListener interface {
	antlr.ParseTreeListener

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

	// EnterRoot_term is called when entering the root_term production.
	EnterRoot_term(c *Root_termContext)

	// EnterNs is called when entering the ns production.
	EnterNs(c *NsContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterData is called when entering the data production.
	EnterData(c *DataContext)

	// EnterCall_term is called when entering the call_term production.
	EnterCall_term(c *Call_termContext)

	// EnterOperator_term is called when entering the operator_term production.
	EnterOperator_term(c *Operator_termContext)

	// EnterRef_call_term is called when entering the ref_call_term production.
	EnterRef_call_term(c *Ref_call_termContext)

	// EnterRef_operator_term is called when entering the ref_operator_term production.
	EnterRef_operator_term(c *Ref_operator_termContext)

	// EnterBoolean_term is called when entering the boolean_term production.
	EnterBoolean_term(c *Boolean_termContext)

	// EnterInteger_term is called when entering the integer_term production.
	EnterInteger_term(c *Integer_termContext)

	// EnterFloat_term is called when entering the float_term production.
	EnterFloat_term(c *Float_termContext)

	// EnterString_term is called when entering the string_term production.
	EnterString_term(c *String_termContext)

	// EnterRef_term is called when entering the ref_term production.
	EnterRef_term(c *Ref_termContext)

	// EnterComplex_term is called when entering the complex_term production.
	EnterComplex_term(c *Complex_termContext)

	// EnterGlob_term is called when entering the glob_term production.
	EnterGlob_term(c *Glob_termContext)

	// EnterMode_term is called when entering the mode_term production.
	EnterMode_term(c *Mode_termContext)

	// EnterSeparate_term is called when entering the separate_term production.
	EnterSeparate_term(c *Separate_termContext)

	// EnterDatablock_term is called when entering the datablock_term production.
	EnterDatablock_term(c *Datablock_termContext)

	// EnterMatchblock_term is called when entering the matchblock_term production.
	EnterMatchblock_term(c *Matchblock_termContext)

	// EnterLogicblock_term is called when entering the logicblock_term production.
	EnterLogicblock_term(c *Logicblock_termContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitRoot_term is called when exiting the root_term production.
	ExitRoot_term(c *Root_termContext)

	// ExitNs is called when exiting the ns production.
	ExitNs(c *NsContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitData is called when exiting the data production.
	ExitData(c *DataContext)

	// ExitCall_term is called when exiting the call_term production.
	ExitCall_term(c *Call_termContext)

	// ExitOperator_term is called when exiting the operator_term production.
	ExitOperator_term(c *Operator_termContext)

	// ExitRef_call_term is called when exiting the ref_call_term production.
	ExitRef_call_term(c *Ref_call_termContext)

	// ExitRef_operator_term is called when exiting the ref_operator_term production.
	ExitRef_operator_term(c *Ref_operator_termContext)

	// ExitBoolean_term is called when exiting the boolean_term production.
	ExitBoolean_term(c *Boolean_termContext)

	// ExitInteger_term is called when exiting the integer_term production.
	ExitInteger_term(c *Integer_termContext)

	// ExitFloat_term is called when exiting the float_term production.
	ExitFloat_term(c *Float_termContext)

	// ExitString_term is called when exiting the string_term production.
	ExitString_term(c *String_termContext)

	// ExitRef_term is called when exiting the ref_term production.
	ExitRef_term(c *Ref_termContext)

	// ExitComplex_term is called when exiting the complex_term production.
	ExitComplex_term(c *Complex_termContext)

	// ExitGlob_term is called when exiting the glob_term production.
	ExitGlob_term(c *Glob_termContext)

	// ExitMode_term is called when exiting the mode_term production.
	ExitMode_term(c *Mode_termContext)

	// ExitSeparate_term is called when exiting the separate_term production.
	ExitSeparate_term(c *Separate_termContext)

	// ExitDatablock_term is called when exiting the datablock_term production.
	ExitDatablock_term(c *Datablock_termContext)

	// ExitMatchblock_term is called when exiting the matchblock_term production.
	ExitMatchblock_term(c *Matchblock_termContext)

	// ExitLogicblock_term is called when exiting the logicblock_term production.
	ExitLogicblock_term(c *Logicblock_termContext)
}
