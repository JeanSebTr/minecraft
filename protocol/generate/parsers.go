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

// func (i Uint16) Write(c *Conn, v McVersion) (err error) {
// 	bs := c.wb[:2]
// 	binary.BigEndian.PutUint16(bs, uint16(i))
// 	_, err = c.Out.Write(bs)
// 	return
// }

// func (i *Uint16) Read(c *Conn, v McVersion) (err error) {
// 	bs := c.rb[:2]
// 	_, err = io.ReadFull(c.In, bs)
// 	if err != nil {
// 		return
// 	}
// 	*i = Uint16(binary.BigEndian.Uint16(bs))
// 	return
// }

func (sp *intSpec) generate() []ast.Decl {
	var iType *ast.Ident
	strLen := strconv.Itoa(sp.size * 8)

	if sp.signed {
		iType = ast.NewIdent("int" + strLen)
	} else {
		iType = ast.NewIdent("uint" + strLen)
	}
	// encode
	encodeMethod := createFuncDecl(sp.encode, iType, ident_i, WRITE)
	encodeMethod.Body.List = []ast.Stmt{
		createBufferInit(sp.size, WRITE), // buf := c.wb[:size]
		// putint // binary.BigEndian.PutUint16(bs, uint16(i))
		stmt_writeBuf,
		stmt_return,
	}

	// decode
	decodeMethod := createFuncDecl(sp.decode, iType, ident_i, READ)
	decodeMethod.Body.List = []ast.Stmt{
		createBufferInit(sp.size, READ),
		stmt_readFull, // _, err = io.ReadFull(c.In, buf)
		// put // i = Uint16(binary.BigEndian.Uint16(bs))
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
		Lhs: []ast.Expr{ident_buf},
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
