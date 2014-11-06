package main

import (
	"go/ast"
	"go/token"
	"strconv"
)

type spec interface {
	parse(ast.Expr, tMode) ast.Stmt
	generate() []ast.Decl
}

type extSpec struct {
	encode, decode *ast.Ident
}

func ext(encode, decode string) *extSpec {
	return &extSpec{ast.NewIdent(encode), ast.NewIdent(decode)}
}

type intSpec struct {
	size   int
	signed bool
	*extSpec
}

func i(l int) *intSpec {
	return (&intSpec{l, true, nil}).init()
}

func ui(l int) *intSpec {
	return (&intSpec{l, false, nil}).init()
}

func (sp *extSpec) parse(expr ast.Expr, dir tMode) ast.Stmt {
	call := &ast.CallExpr{}
	assign := &ast.AssignStmt{
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{call},
	}
	cond := &ast.IfStmt{
		Init: assign,
		Cond: e_cond,
		Body: e_body,
	}
	if dir == READ {
		assign.Lhs = []ast.Expr{expr, i_err}
		call.Fun = sp.decode
		call.Args = []ast.Expr{i_cvar, i_vers}
	} else {
		assign.Lhs = []ast.Expr{i_err}
		call.Fun = sp.encode
		call.Args = []ast.Expr{expr, i_cvar, i_vers}
	}
	return cond
}

func (sp *extSpec) generate() []ast.Decl {
	return nil // external type!
}

func (sp *intSpec) init() *intSpec {
	strLen := strconv.Itoa(sp.size * 8)
	if sp.signed {
		sp.extSpec = &extSpec{
			ast.NewIdent("encodeInt" + strLen),
			ast.NewIdent("decodeInt" + strLen),
		}
	} else {
		sp.extSpec = &extSpec{
			ast.NewIdent("encodeUint" + strLen),
			ast.NewIdent("decodeUint" + strLen),
		}
	}
	return sp // just for chaining
}

func (sp *intSpec) generate() []ast.Decl {
	return nil
}
