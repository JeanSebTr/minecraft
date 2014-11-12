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

// required, but not useful
var fset = token.NewFileSet()

func main() {
	// parse source from Stdin
	src, err := parser.ParseFile(fset, "packets.go", os.Stdin, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	// scan all types
	ast.Inspect(src, func(n ast.Node) (cont bool) {
		sp, st, cont := scanTypes(n)
		if sp != nil && st != nil {
			types[sp.Name.Name] = &locSpec{}
		}
		return cont
	})

	// will store every generated methods
	var declarations []ast.Decl

	// imports
	declarations = append(declarations, &ast.GenDecl{
		Tok:    token.IMPORT,
		Lparen: 1,
		Specs: []ast.Spec{
			// binary
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"encoding/binary"`,
				},
			},
			// io
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"io"`,
				},
			},
		},
	})

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

func createMethod(name *ast.Ident, dir tMode, st *ast.StructType) *ast.FuncDecl {
	body := &ast.BlockStmt{
		List: make([]ast.Stmt, 0),
	}
	var method *ast.Ident
	if dir == READ {
		method = ident_Read
	} else {
		method = ident_Write
	}
	fd := &ast.FuncDecl{
		Recv: createFieldList("pkt", "*"+name.Name),
		Name: method,
		Type: fType_pktMethod,
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
			fmt.Fprintf(os.Stderr, "[]%+v @ %v\n", tIdent.Elt, fset.Position(field.Type.Pos()))
		default:
			panic(fmt.Errorf("Not *ast.Ident or *ast.ArrayType! %T %+v", field.Type, field.Type))
		}
	}
	body.List = append(body.List, stmt_return)
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
			X:   ident_pkt,
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
		}
		fmt.Fprintf(os.Stderr, "spec = %T %v\n", sp, sp)
	} else {
		prev = markTODO(prev, expr)
	}
	return prev
}

func createArrayStatement(prev []ast.Stmt, expr ast.Expr, sp spec, dir tMode, tag reflect.StructTag) []ast.Stmt {
	tLenName := tag.Get("ltype")
	lenType, ok := types[tLenName]
	prev = markTODO(prev, expr)
	if !ok {
		if tLenName == "nil" {
			return prev
		}
		panic(fmt.Errorf("Unknown ltype:%s of %v", tLenName, expr))
	}
	if dir == WRITE {
		lenCall := cast(ident_len, expr)
		valExpr := cast(ast.NewIdent(tLenName), lenCall)
		fmt.Fprintf(os.Stderr, "%v %v\n", lenType, tLenName)
		stmt := lenType.parse(valExpr, dir)
		fmt.Fprintf(os.Stderr, "%v %v\n", prev, stmt)
		prev = append(prev, stmt)
	}

	return prev
}

var todo = ast.NewIdent("TODO")

func markTODO(prev []ast.Stmt, expr ast.Expr) []ast.Stmt {
	return append(prev, &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun:  todo,
			Args: []ast.Expr{expr},
		},
	})
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
