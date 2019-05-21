package main

import (
	"io"
	"strings"
	templatex "text/template"

	"github.com/gobuffalo/packr/v2"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

// template -
var template *templatex.Template

type plugin struct {
	pgs.ModuleBase
	pgsgo.Context
}

func (*plugin) Name() string {
	return "serialization"
}

func (p *plugin) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	p.Context = pgsgo.InitContext(c.Parameters())
}
func init() {
	template = templatex.New("serialization")
	box := packr.New("serializationBox", "./templates")
	err := box.Walk(walkFN)
	if err != nil {
		panic(err)
	}
}

var walkFN = func(s string, file packr.File) error {
	var sb strings.Builder
	if _, err := io.Copy(&sb, file); err != nil {
		return err
	}
	var err error
	template, err = template.Parse(sb.String())
	return err
}
