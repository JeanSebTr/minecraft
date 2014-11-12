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

	ident_io    = ast.NewIdent("io")
	ident_rb    = ast.NewIdent("rb")
	ident_wb    = ast.NewIdent("wb")
	ident_buf   = ast.NewIdent("buf")
	ident_nil   = ast.NewIdent("nil")
	ident_pkt   = ast.NewIdent("pkt")
	ident_len   = ast.NewIdent("len")
	ident_err   = ast.NewIdent("err")
	ident_error = ast.NewIdent("error")

	ident_Conn      = ast.NewIdent("Conn")
	ident_input     = ast.NewIdent("In")
	ident_output    = ast.NewIdent("Out")
	ident_Read      = ast.NewIdent("Read")
	ident_Write     = ast.NewIdent("Write")
	ident_ReadFull  = ast.NewIdent("ReadFull")
	ident_McVersion = ast.NewIdent("McVersion")
	ident_binary    = ast.NewIdent("binary")
	ident_bigEndian = ast.NewIdent("BigEndian")
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
	sel_input = &ast.SelectorExpr{
		X:   ident_c,
		Sel: ident_input,
	}
	sel_output = &ast.SelectorExpr{
		X:   ident_c,
		Sel: ident_output,
	}
	sel_writeFn = &ast.SelectorExpr{
		X:   sel_output,
		Sel: ident_Write,
	}
	sel_readfull = &ast.SelectorExpr{
		X:   ident_io,
		Sel: ident_ReadFull,
	}
	sel_bigEndian = &ast.SelectorExpr{
		X:   ident_binary,
		Sel: ident_bigEndian,
	}
)

// expressions
var (
	expr_errNotNil = &ast.BinaryExpr{
		X:  ident_err,
		Op: token.NEQ,
		Y:  ident_nil,
	}
	exprL_ident_i     = []ast.Expr{ident_i}
	exprL_ident_buf   = []ast.Expr{ident_buf}
	exprL_ident_err   = []ast.Expr{ident_err}
	exprL_connAndVer  = []ast.Expr{ident_c, ident_v}
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
	stmt_return      = &ast.ReturnStmt{}
	stmt_blockReturn = &ast.BlockStmt{
		List: []ast.Stmt{stmt_return},
	}
	stmt_retOnErr = &ast.IfStmt{
		Cond: expr_errNotNil,
		Body: stmt_blockReturn,
	}
	// _, err = io.ReadFull(c.In, bs)
	stmt_readFull = &ast.AssignStmt{
		Lhs: exprL_blankAndErr,
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{&ast.CallExpr{
			Fun:  sel_readfull,
			Args: []ast.Expr{sel_input, ident_buf},
		}},
	}
	// _, err = c.Out.Write(buf)
	stmt_writeBuf = &ast.AssignStmt{
		Lhs: exprL_blankAndErr,
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{&ast.CallExpr{
			Fun:  sel_writeFn,
			Args: []ast.Expr{ident_buf},
		}},
	}
)

// misc
var (
	fType_pktMethod = &ast.FuncType{
		Params:  fList_paramsRead,
		Results: fList_returnWrite,
	}
)
