package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"reflect"
)

type tMode uint8

const (
	READ tMode = iota
	WRITE
)

func main() {
	// required, but not useful
	fset := token.NewFileSet()

	// parse source from Stdin
	src, err := parser.ParseFile(fset, "packets.go", os.Stdin, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	// scan all types
	ast.Inspect(src, func(n ast.Node) (cont bool) {
		sp, st, cont := scanTypes(n)
		if sp != nil && st != nil {
			types[sp.Name.Name] = nil // just so we know this type exists
		}
		return cont
	})

	// will store every generated methods
	declarations := make([]ast.Decl, 0)

	// generate r/w methods for paquet types
	ast.Inspect(src, func(n ast.Node) (cont bool) {
		sp, st, cont := scanTypes(n)
		if sp != nil && st != nil {
			readMethod := createMethod(sp.Name, READ, st)
			declarations = append(declarations, readMethod)

			writeMethod := createMethod(sp.Name, WRITE, st)
			declarations = append(declarations, writeMethod)
		}
		return cont
	})

	// generate r/w methods for existing types
	for _, typeSpec := range types {
		if typeSpec != nil {
			for _, declaration := range typeSpec.generate() {
				declarations = append(declarations, declaration)
			}
		}
	}

	// create a new File to store the AST
	dst := &ast.File{
		Name:  src.Name,
		Decls: declarations,
	}

	err = printer.Fprint(os.Stdout, fset, dst)
	if err != nil {
		panic(err)
	}
}

func scanTypes(n ast.Node) (*ast.TypeSpec, *ast.StructType, bool) {
	if n == nil {
		return nil, nil, false
	}
	// root node
	if _, ok := n.(*ast.File); ok {
		return nil, nil, true
	}
	// declaration node
	if genDec, ok := n.(*ast.GenDecl); ok && genDec.Tok == token.TYPE {
		return nil, nil, true
	}
	spec, ok := n.(*ast.TypeSpec)
	// we just want struct declaration
	if !ok {
		return nil, nil, false
	}
	st, ok := spec.Type.(*ast.StructType)
	if !ok {
		return nil, nil, false
	}
	return spec, st, false
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

func createMethod(name *ast.Ident, dir tMode, st *ast.StructType) *ast.FuncDecl {
	body := &ast.BlockStmt{
		List: make([]ast.Stmt, 0),
	}
	var method *ast.Ident
	if dir == READ {
		method = i_read
	} else {
		method = i_write
	}
	fd := &ast.FuncDecl{
		Recv: createFieldList("pkt", "*"+name.Name),
		Name: method,
		Type: e_fntype,
		Body: body,
	}
	for _, field := range st.Fields.List {
		switch tIdent := field.Type.(type) {
		case *ast.Ident:
			body.List = createFieldsStatements(body.List, tIdent.Name, field, dir, createSimpleStatement)
		case *ast.ArrayType:
			if arrIdent, ok := tIdent.Elt.(*ast.Ident); ok {
				body.List = createFieldsStatements(body.List, arrIdent.Name, field, dir, createArrayStatement)
			} else {
				panic(fmt.Errorf("*ast.ArrayType! %T %+v", tIdent.Elt, tIdent.Elt))
			}
		default:
			panic(fmt.Errorf("Not *ast.Ident or *ast.ArrayType! %T %+v", field.Type, field.Type))
		}
	}
	body.List = append(body.List, e_retErr)
	return fd
}

type stmtFactory func([]ast.Stmt, ast.Expr, spec, tMode, reflect.StructTag) []ast.Stmt

func createFieldsStatements(prev []ast.Stmt, tName string, field *ast.Field, dir tMode, factory stmtFactory) []ast.Stmt {
	sp, ok := types[tName]
	if !ok {
		panic(fmt.Errorf("Undefined type: %s on fields %v", tName, field.Names))
	}
	var tag reflect.StructTag
	if field.Tag != nil {
		tag = reflect.StructTag(field.Tag.Value[1 : len(field.Tag.Value)-1])
	}
	for _, name := range field.Names {
		prev = factory(prev, &ast.SelectorExpr{
			X:   i_pkt,
			Sel: ast.NewIdent(name.Name),
		}, sp, dir, tag)
	}
	return prev
}

func createSimpleStatement(prev []ast.Stmt, expr ast.Expr, sp spec, dir tMode, tag reflect.StructTag) []ast.Stmt {
	if sp != nil {
		st := sp.parse(expr, dir)
		if st != nil {
			return append(prev, st)
		} else {
			fmt.Fprintf(os.Stderr, "spec = %T %v\n", sp, sp)
		}
	}
	return prev
}

func createArrayStatement(prev []ast.Stmt, expr ast.Expr, sp spec, dir tMode, tag reflect.StructTag) []ast.Stmt {
	tLenName := tag.Get("ltype")
	_, ok := types[tLenName]
	if !ok {
		if tLenName == "nil" {
			return prev
		}
		panic(fmt.Errorf("Unknown ltype:%s of %v", tLenName, expr))
	}
	fmt.Fprintf(os.Stderr, "[] %T, %+v %+v %s:%s\n", expr, expr, sp, tag, tag.Get("ltype"))
	return prev
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
