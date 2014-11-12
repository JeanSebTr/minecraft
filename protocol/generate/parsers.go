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

type locSpec struct {
	//
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

func (sp *locSpec) parse(expr ast.Expr, dir tMode) ast.Stmt {
	sel := &ast.SelectorExpr{
		X: expr,
	}
	if dir == READ {
		sel.Sel = ident_Read
	} else {
		sel.Sel = ident_Write
	}
	return &ast.IfStmt{
		Init: &ast.AssignStmt{
			Lhs: exprL_ident_err,
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun:  sel,
					Args: exprL_connAndVer,
				},
			},
		},
		Cond: expr_errNotNil,
		Body: stmt_blockReturn,
	}
}

func (sp *locSpec) generate() []ast.Decl {
	return nil // local type!
}

func (sp *extSpec) parse(expr ast.Expr, dir tMode) ast.Stmt {
	call := &ast.CallExpr{}
	assign := &ast.AssignStmt{
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{call},
	}
	cond := &ast.IfStmt{
		Init: assign,
		Cond: expr_errNotNil,
		Body: stmt_blockReturn,
	}
	if dir == READ {
		assign.Lhs = []ast.Expr{expr, ident_err}
		call.Fun = sp.decode
		call.Args = exprL_connAndVer
	} else {
		assign.Lhs = exprL_ident_err
		call.Fun = sp.encode
		call.Args = []ast.Expr{expr, ident_c, ident_v}
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
	strLen := strconv.Itoa(sp.size * 8)
	iType := ast.NewIdent("int" + strLen)
	iCast := ast.NewIdent("uint" + strLen)
	iMethod := "Uint" + strLen

	if !sp.signed {
		iType, iCast = iCast, iType
	}
	// encode
	encodeMethod := createFuncDecl(sp.encode, iType, ident_i, WRITE)
	encodeMethod.Body.List = []ast.Stmt{
		createBufferInit(sp.size, WRITE),
		// binary.BigEndian.PutUint16(buf, uint16(i))
		&ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   sel_bigEndian,
					Sel: ast.NewIdent("Put" + iMethod),
				},
				Args: []ast.Expr{
					ident_buf,
					cast(iCast, ident_i, sp.signed),
				},
			},
		},
		stmt_writeBuf,
		stmt_return,
	}

	// decode
	decodeMethod := createFuncDecl(sp.decode, iType, ident_i, READ)
	decodeMethod.Body.List = []ast.Stmt{
		createBufferInit(sp.size, READ),
		stmt_readFull,
		stmt_retOnErr,
		// i = Uint16(binary.BigEndian.Uint16(bs))
		&ast.AssignStmt{
			Lhs: exprL_ident_i,
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{
				cast(iType, &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   sel_bigEndian,
						Sel: ast.NewIdent(iMethod),
					},
					Args: exprL_ident_buf,
				}, sp.signed),
			},
		},
		stmt_return,
	}
	return []ast.Decl{encodeMethod, decodeMethod}
}

func createFuncDecl(fName, tName, pName *ast.Ident, dir tMode) *ast.FuncDecl {
	fn := &ast.FuncDecl{
		Name: fName,
		Type: &ast.FuncType{},
		Body: &ast.BlockStmt{},
	}
	field := &ast.Field{
		Names: []*ast.Ident{pName},
		Type:  tName,
	}
	if dir == READ {
		fn.Type.Params = fList_paramsRead
		fn.Type.Results = &ast.FieldList{
			List: []*ast.Field{
				field,
				field_err,
			},
		}
	} else {
		fn.Type.Params = &ast.FieldList{
			List: []*ast.Field{
				field,
				field_c,
				field_version,
			},
		}
		fn.Type.Results = fList_returnWrite
	}
	return fn
}

func createBufferInit(length int, dir tMode) ast.Stmt {
	var sel *ast.SelectorExpr
	if dir == READ {
		sel = sel_readbuf
	} else {
		sel = sel_writebuf
	}
	// bs := c.rb[:2]
	assign := &ast.AssignStmt{
		Lhs: exprL_ident_buf,
		Tok: token.DEFINE,
		Rhs: []ast.Expr{&ast.SliceExpr{
			X:   sel,
			Low: nil,
			High: &ast.BasicLit{
				Kind:  token.INT,
				Value: strconv.Itoa(length),
			},
		}},
	}
	return assign
}

func cast(typeIdent, varIdent ast.Expr, shouldCast ...bool) ast.Expr {
	if len(shouldCast) == 0 || shouldCast[0] {
		return &ast.CallExpr{
			Fun:  typeIdent,
			Args: []ast.Expr{varIdent},
		}
	}
	return varIdent
}
