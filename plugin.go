package main

import (

	// "github.com/bifrostcloud/protoc-gen-serialization/pkg/tags"
	"strings"

	"github.com/fatih/camelcase"

	pgs "github.com/lyft/protoc-gen-star"
)

type base struct {
	Package  string
	Imports  []pkg
	Services []service
}
type service struct {
	UpperCamelCaseServiceName string
	LowerCamelCaseServiceName string
	Auth                      string
	Methods                   []rpc
}
type fieldImport struct {
	Name string
	Tag  string
}
type fields struct {
	Tmp []string
	// Struct Name
	TypeName string
	// StructType
	Type string
	// Field Structs
	FieldImport []fieldImport
	// Proto Message Name
	UpperCamelCase string
	// Keeps the fields name as a lowercase of message field in proto file
	Base []string
	// turns the field name into a single lowercase string without any seperators
	Lowercase []string
	// turns the field name into a single lowercase with dot seperators
	DotNotation []string
	// turns the field name into a lowercase, paramcase string
	ParamCase []string
}
type rpc struct {
	UpperCamelCaseMethodName  string
	UpperCamelCaseServiceName string
	InputType                 string
	InputFields               fields
	OutputType                string
}

// Input -
type Input struct {
	FieldName  string
	FieldValue string
}
type pkg struct {
	PackageName string
	PackagePath string
}

func (p *plugin) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	p.genGoTemplate(targets, packages)
	// lang := p.Parameters().Str("target_ext")
	// switch strings.ToLower(lang) {
	// case "go":
	// 	{
	// 		p.genGoTemplate(targets, packages)
	// 		break
	// 	}
	// default:
	// 	{
	// 		p.genGoTemplate(targets, packages)
	// 		break
	// 	}
	// }
	return p.Artifacts()
}
func (p *plugin) genGoTemplate(targets map[string]pgs.File, packages map[string]pgs.Package) {
	messages := map[string]bool{}

	for _, file := range targets {
		b := base{Package: p.Context.PackageName(file).String()}
		imports := map[string]pkg{
			"json": pkg{
				PackagePath: "encoding/json",
			},
			"xml": pkg{
				PackagePath: "encoding/xml",
			},
			"reflect": pkg{
				PackagePath: "reflect",
			},
			"strings": pkg{
				PackagePath: "strings",
			},
			"stacktrace": pkg{
				PackageName: "stacktrace",
				PackagePath: "github.com/palantir/stacktrace",
			},
			"mapstructure": pkg{
				PackageName: "mapstructure",
				PackagePath: "github.com/mitchellh/mapstructure",
			},
		}
		for _, srv := range file.Services() {
			s := service{}
			s.UpperCamelCaseServiceName = srv.Name().UpperCamelCase().String()
			s.LowerCamelCaseServiceName = srv.Name().LowerCamelCase().String()
			for _, method := range srv.Methods() {
				upperCamelCaseMethodName := p.Context.Name(method).UpperCamelCase().String()
				r := rpc{}
				r.UpperCamelCaseServiceName = srv.Name().UpperCamelCase().String()
				r.UpperCamelCaseMethodName = upperCamelCaseMethodName
				inputKey := p.Context.Name(method.Input()).UpperCamelCase().String()
				outputKey := p.Context.Name(method.Output()).UpperCamelCase().String()

				if !messages[inputKey] {
					messages[inputKey] = true

					r.InputType = p.Context.Name(method.Input()).String()
					if !method.Input().BuildTarget() {
						path := p.Context.ImportPath(method.Input()).String()
						imports[path] = pkg{
							PackageName: p.Context.PackageName(method.Input()).String(),
							PackagePath: path,
						}
						r.InputType = p.Context.PackageName(method.Input()).String() + "." + p.Context.Name(method.Input()).String()
					}
					r.InputFields.TypeName = p.Context.Name(method.Input()).UpperCamelCase().String()
					r.InputFields.Type = r.InputType
					for _, field := range method.Input().Fields() {
						r.InputFields.FieldImport = append(r.InputFields.FieldImport, fieldImport{
							Name: field.Name().UpperCamelCase().String(),
							Tag:  field.Name().String(),
						})
						r.InputFields.Base = append(r.InputFields.Base, strings.ToLower(field.Name().String()))
						r.InputFields.Lowercase = append(r.InputFields.Lowercase, strings.ToLower(field.Name().LowerCamelCase().String()))
						r.InputFields.DotNotation = append(r.InputFields.DotNotation, strings.ToLower(field.Name().LowerDotNotation().String()))
						spl := camelcase.Split(field.Name().UpperCamelCase().String())
						r.InputFields.ParamCase = append(r.InputFields.ParamCase, strings.ToLower(strings.Join(spl, "-")))
					}

				}
				if !messages[outputKey] {
					messages[outputKey] = true
					r.OutputType = p.Context.Name(method.Output()).String()
					if !method.Output().BuildTarget() {
						path := p.Context.ImportPath(method.Output()).String()
						imports[path] = pkg{
							PackagePath: path,
							PackageName: p.Context.PackageName(method.Output()).String(),
						}
						r.OutputType = p.Context.PackageName(method.Output()).String() + "." + p.Context.Name(method.Output()).String()
					}
				}

				s.Methods = append(s.Methods, r)
			}
			b.Services = append(b.Services, s)
		}
		if len(b.Services) == 0 {
			continue
		}
		for _, pkg := range imports {
			b.Imports = append(b.Imports, pkg)
		}
		pname := p.Context.OutputPath(file).SetExt(".serialization.go").String()
		p.OverwriteGeneratorTemplateFile(
			pname,
			template.Lookup("Base_go"),
			&b,
		)

	}
}
