package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func main() {
	mod := pgs.Init(pgs.DebugEnv("DEBUG"))

	mod.RegisterModule(&plugin{ModuleBase: pgs.ModuleBase{}})

	mod.RegisterPostProcessor(pgsgo.GoFmt())
	mod.Render()

}
