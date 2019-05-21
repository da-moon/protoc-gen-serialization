package tags

import (
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
	"strings"
)

var (
	rTags = regexp.MustCompile(`[\w_]+:"[^"]+"`)
)

// Tag -
type Tag struct {
	Key string

	Value string
}

// type Tag map[string]map[string]*structtag.Tags

// Extract -
func Extract(inputPath string) ([]Tag, error) {

	// result := make(Tag)
	result := make([]Tag, 0)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, inputPath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	for _, decl := range f.Decls {
		// check if is generic declaration
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		var typeSpec *ast.TypeSpec
		for _, spec := range genDecl.Specs {
			if ts, tsOK := spec.(*ast.TypeSpec); tsOK {
				typeSpec = ts
				break
			}
		}
		if typeSpec == nil {
			continue
		}
		structDecl, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			continue
		}
		for _, field := range structDecl.Fields.List {
			if len(field.Names) > 0 {
				if !strings.HasPrefix(field.Names[0].Name, "XXX") {
					t := field.Tag.Value
					area := Tag{
						Key:   field.Names[0].Name,
						Value: t[1 : len(t)-1],
					}
					// result[field.Names[0].Name] = t[1 : len(t)-1]
					result = append(result, area)
				}
			}
			if field.Doc == nil {
				continue
			}
		}
	}
	return result, nil
}
