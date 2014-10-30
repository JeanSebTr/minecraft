package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

const (
	packageName = "protocol"
)

func main() {

	fset := token.NewFileSet()

	src, err := parser.ParseFile(fset, "packets.go", os.Stdin, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	decls := make([]ast.Decl, 0)

	f := ast.File{
		Doc:        nil,
		Name:       src.Name,
		Decls:      nil, // top-level declarations; or nil
		Scope:      nil, //   *Scope          // package scope (this file only)
		Imports:    nil, //   []*ImportSpec   // imports in this file
		Unresolved: nil, //[]*Ident        // unresolved identifiers in this file
		Comments:   nil, //[]*CommentGroup // list of all comments in the source file
	}

	ast.Inspect(src, func(n ast.Node) bool {
		if n == nil {
			return false
		}
		// root node
		if _, ok := n.(*ast.File); ok {
			return true
		}
		// declaration node
		if genDec, ok := n.(*ast.GenDecl); ok && genDec.Tok == token.TYPE {
			return true
		}
		spec, ok := n.(*ast.TypeSpec)
		// we just want struct declaration
		if !ok {
			return false
		}
		if st, ok := spec.Type.(*ast.StructType); ok {
			fd := createWriteMethod(spec.Name, st)
			decls = append(decls, fd)

			fd = createReadMethod(spec.Name, st)
			decls = append(decls, fd)
		}
		return false
	})

	f.Decls = decls

	err = printer.Fprint(os.Stdout, fset, &f)
	if err != nil {
		panic(err)
	}
}

var ( // variables
	i_cvar  = ast.NewIdent("c")
	i_pkt   = ast.NewIdent("pkt")
	i_err   = ast.NewIdent("err")
	i_nil   = ast.NewIdent("nil")
	i_read  = ast.NewIdent("Read")
	i_write = ast.NewIdent("Write")
	i_vers  = ast.NewIdent("version")
)

var ( // ast parts
	e_cond = &ast.BinaryExpr{
		X:  i_err,
		Op: token.NEQ,
		Y:  i_nil,
	}
	e_retErr = &ast.ReturnStmt{
		Results: []ast.Expr{i_err},
	}
	e_body = &ast.BlockStmt{
		List: []ast.Stmt{e_retErr},
	}
	e_fntype = &ast.FuncType{
		Params:  createFieldList("c", "*Conn", "version", "McVersion"),
		Results: createFieldList("err", "error"),
	}
)

func createReadMethod(name *ast.Ident, st *ast.StructType) *ast.FuncDecl {
	body := &ast.BlockStmt{
		List: make([]ast.Stmt, 0),
	}
	fd := &ast.FuncDecl{
		Doc:  nil,
		Recv: createFieldList("pkt", "*"+name.Name),
		Name: i_read,
		Type: e_fntype,
		Body: body,
	}
	for _, field := range st.Fields.List {
		for _, name := range field.Names {
			st := createReadFieldStatement(name, field.Type, field.Tag)
			body.List = append(body.List, st)
		}
	}
	body.List = append(body.List, e_retErr)
	return fd
}

func createReadFieldStatement(fieldName *ast.Ident, fType ast.Expr, tag *ast.BasicLit) ast.Stmt {
	typeName, ok := fType.(*ast.Ident)
	if !ok {
		panic(fmt.Errorf("Field type of ?.%s should be an identifier. Got %T", fieldName.Name, fType))
	}
	cond := &ast.IfStmt{
		Init: &ast.AssignStmt{
			Lhs: []ast.Expr{&ast.SelectorExpr{
				X:   i_pkt,
				Sel: ast.NewIdent(fieldName.Name),
			}, i_err},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{&ast.CallExpr{
				Fun:  ast.NewIdent("Read" + typeName.Name),
				Args: []ast.Expr{i_cvar, i_vers},
			}},
		},
		Cond: e_cond,
		Body: e_body,
	}

	return cond
}

func createWriteMethod(name *ast.Ident, st *ast.StructType) *ast.FuncDecl {
	body := &ast.BlockStmt{
		List: make([]ast.Stmt, 0),
	}
	fd := &ast.FuncDecl{
		Doc:  nil,
		Recv: createFieldList("pkt", "*"+name.Name),
		Name: i_write,
		Type: e_fntype,
		Body: body,
	}
	for _, field := range st.Fields.List {
		for _, name := range field.Names {
			st := createWriteFieldStatement(name, field.Type, field.Tag)
			body.List = append(body.List, st)
		}
	}
	body.List = append(body.List, e_retErr)
	return fd
}

func createWriteFieldStatement(fieldName *ast.Ident, fType ast.Expr, tag *ast.BasicLit) ast.Stmt {

	cond := &ast.IfStmt{
		Init: &ast.AssignStmt{
			Lhs: []ast.Expr{i_err},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.SelectorExpr{
						X:   i_pkt,
						Sel: ast.NewIdent(fieldName.Name),
					},
					Sel: i_write,
				},
				Args: []ast.Expr{i_cvar, i_vers},
			}},
		},
		Cond: e_cond,
		Body: e_body,
	}

	return cond
}

func createFieldList(names ...string) *ast.FieldList {
	l := len(names)
	if l%2 != 0 {
		panic(fmt.Errorf("Need even number of arguments. Got %d [%v]", l, names))
	}
	fields := make([]*ast.Field, l/2)

	for i := 0; i < l/2; i++ {
		fields[i] = &ast.Field{
			Names: identsOrNil(names[i*2]),
			Type:  ast.NewIdent(names[i*2+1]),
		}
	}

	return &ast.FieldList{
		List: fields,
	}
}

func identsOrNil(name string) []*ast.Ident {
	if name == "" {
		return nil
	}
	return []*ast.Ident{ast.NewIdent(name)}
}
