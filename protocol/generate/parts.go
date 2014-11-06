package main

import (
	"go/ast"
	"go/token"
)

// identifier
var (
	ident_  = ast.NewIdent("_")
	ident_c = ast.NewIdent("c")
	ident_i = ast.NewIdent("i")
	ident_v = ast.NewIdent("v")

	ident_rb        = ast.NewIdent("rb")
	ident_wb        = ast.NewIdent("wb")
	ident_buf       = ast.NewIdent("buf")
	ident_err       = ast.NewIdent("err")
	ident_error     = ast.NewIdent("error")
	ident_Conn      = ast.NewIdent("Conn")
	ident_McVersion = ast.NewIdent("McVersion")
)

// selector aka `struct.field`
var (
	sel_readbuf = &ast.SelectorExpr{
		X:   ident_c,
		Sel: ident_rb,
	}
	sel_writebuf = &ast.SelectorExpr{
		X:   ident_c,
		Sel: ident_wb,
	}
)

// expressions
var (
	exprL_blankAndErr = []ast.Expr{ident_, ident_err}
)

// fields
var (
	field_c = &ast.Field{
		Names: []*ast.Ident{ident_c},
		Type:  &ast.StarExpr{X: ident_Conn},
	}
	field_version = &ast.Field{
		Names: []*ast.Ident{ident_v},
		Type:  ident_McVersion,
	}
	field_err = &ast.Field{
		Names: []*ast.Ident{ident_err},
		Type:  ident_error,
	}
	fList_paramsRead = &ast.FieldList{
		List: []*ast.Field{
			field_c,
			field_version,
		},
	}
	fList_returnWrite = &ast.FieldList{
		List: []*ast.Field{field_err},
	}
)

// statements
var (
	stmt_return = &ast.ReturnStmt{}
	// _, err = io.ReadFull(c.In, bs)
	stmt_readFull = &ast.AssignStmt{
		Lhs: exprL_blankAndErr,
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{},
	}
	// _, err = c.Out.Write(buf)
	stmt_writeBuf = &ast.AssignStmt{
		Lhs: exprL_blankAndErr,
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{},
	}
)
