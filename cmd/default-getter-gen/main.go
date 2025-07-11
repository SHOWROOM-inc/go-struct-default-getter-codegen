package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"unicode"
)

func main() {
	inputFile := flag.String("input", "", "Path to input Go file (required)")
	outputFile := flag.String("output", "getters.gen.go", "Path to output Go file")
	packageName := flag.String("package", "main", "Package name for generated code")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "This is tool to generate getters code from structs in specified --input file.\n")
		fmt.Fprintf(os.Stderr, "Usage: %s --input <path/to/model.go> --output <path/to/output.go> --package <package_name>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if *inputFile == "" {
		log.Fatal("❌ --input is required")
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, *inputFile, nil, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	process(node, *outputFile, *packageName)
}

func process(node *ast.File, outputFile, packageName string) {
	structTypes, aliasTypes := searchStructs(node)

	var buf bytes.Buffer
	buf.WriteString("// Code generated by generate_getters.go; DO NOT EDIT.\n")
	buf.WriteString("package " + packageName + "\n\n")
	buf.WriteString("// Getter methods generated from default tags and zero values.\n\n")

	// 構造体の初期値取得関数を生成
	for typeName, structType := range structTypes {
		generateGettersForStruct(&buf, typeName, structType)
	}

	// 型エイリアスに対しても初期値取得関数を生成
	for aliasName, baseName := range aliasTypes {
		if baseStruct, ok := structTypes[baseName]; ok {
			generateGettersForStruct(&buf, aliasName, baseStruct)
		}
	}

	// ファイル出力
	err := os.WriteFile(outputFile, buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Getter methods written to:", outputFile)
}

func searchStructs(node *ast.File) (structTypes map[string]*ast.StructType, aliasTypes map[string]string) {
	structTypes = make(map[string]*ast.StructType)
	aliasTypes = make(map[string]string)

	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec := spec.(*ast.TypeSpec)
			switch t := typeSpec.Type.(type) {
			case *ast.StructType:
				structTypes[typeSpec.Name.Name] = t
			case *ast.Ident:
				// aliasされた型も対象とする。ただし型エイリアス(type alias)は対象外、別型名(define alias)のみ対象とする。
				// - type alias: `type Struct2 = Struct1`　すなわち=で定義
				// - define alias: `type Struct2 Struct1`　すなわち=なしで定義、こちらは別の型となる
				if typeSpec.Assign == 0 {
					aliasTypes[typeSpec.Name.Name] = t.Name
				}
			}
		}
	}

	return
}

func generateGettersForStruct(buf *bytes.Buffer, typeName string, structType *ast.StructType) {
	var getters []string

	for _, field := range structType.Fields.List {
		if len(field.Names) == 0 {
			continue
		}
		fieldName := field.Names[0].Name

		ptrType, ok := field.Type.(*ast.StarExpr)
		if !ok {
			continue
		}
		ident, ok := ptrType.X.(*ast.Ident)
		if !ok {
			continue
		}
		fieldType := ident.Name

		getters = append(getters, generateGetter(typeName, fieldName, fieldType))
	}

	if len(getters) > 0 {
		buf.WriteString(fmt.Sprintf("// ======= Getter methods for %s =======\n", typeName))
		for _, g := range getters {
			buf.WriteString(g + "\n\n")
		}
	}
}

func generateGetter(structName, fieldName, fieldType string) string {
	methodName := "Get" + export(fieldName) + "OrDefault"
	return fmt.Sprintf(`func (r *%s) %s(defaultValue %s) %s {
	if r.%s != nil {
		return *r.%s
	}
	return defaultValue
}`, structName, methodName, fieldType, fieldType, fieldName, fieldName)
}

func export(s string) string {
	r := []rune(s)
	if len(r) > 0 {
		r[0] = unicode.ToUpper(r[0])
	}
	return string(r)
}
