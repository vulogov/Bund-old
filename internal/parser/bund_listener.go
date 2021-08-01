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

	// EnterBoolean_term is called when entering the boolean_term production.
	EnterBoolean_term(c *Boolean_termContext)

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

	// ExitBoolean_term is called when exiting the boolean_term production.
	ExitBoolean_term(c *Boolean_termContext)
}
