// +build mage

package main

import (
	"os"
	"path/filepath"

	"github.com/magefile/mage/sh"
)

func Base() error {
	err := sh.Run("packr2", "clean")
	if err != nil {
		return err
	}
	err = sh.Run(
		"protoc",
		"-I", filepath.Join("proto"),
		"-I", filepath.Join(os.Getenv("GOPATH"), "src"),
		"-I", filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "golang", "protobuf"),
		`--go_out=`+
			`Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor,`+
			`Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any,`+
			`:`+
			filepath.Join("proto"),
		filepath.Join("proto", "types.proto"),
	)
	return nil
}
func Build() error {
	err := Base()
	if err != nil {
		return err
	}
	err = sh.Run("packr2")
	if err != nil {
		return err
	}
	err = sh.Run("go", "install", ".")
	if err != nil {
		sh.Run("packr2", "clean")
		return err
	}
	err = sh.Run("packr2", "clean")
	if err != nil {
		return err
	}
	return nil
}

func Proto() error {
	err := Base()
	if err != nil {
		return err
	}
	// err = sh.Run("make", "-C", "example", "clean")
	// if err != nil {
	// 	return err
	// }
	err = sh.Run(
		"protoc",
		"-I", ".",
		"-I", filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "golang", "protobuf"),
		"-I", filepath.Join(os.Getenv("GOPATH"), "src"),
		"-I", filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "gogo", "protobuf", "protobuf"),
		`--gogo_out=`+
			`Mgithub.com/gogo/protobuf/gogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto,`+
			`Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,`+
			`:.`,
		filepath.Join("example", "example.proto"),
	)
	if err != nil {
		return err
	}
	err = sh.Run(
		"protoc",
		"-I", ".",
		"-I", filepath.Join(os.Getenv("GOPATH"), "src"),
		"-I", filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "gogo", "protobuf", "protobuf"),
		`--serialization_out=`+
			`target_ext="go"`+
			`Mgithub.com/gogo/protobuf/gogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto,`+
			`Mgithub.com/bifrostcloud/protoc-gen-serialization/proto/types.proto=github.com/bifrostcloud/protoc-gen-serialization/proto,`+
			`Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,`+
			`:.`,
		filepath.Join("example", "example.proto"),
	)
	if err != nil {
		return err
	}
	return nil
}
func Example() error {
	err := Build()
	if err != nil {
		return err
	}
	err = Proto()
	if err != nil {
		return err
	}
	err = sh.Run("packr2", "clean")
	if err != nil {
		return err
	}
	return nil
}

func Clean() error {
	err := sh.Run("make", "-C", "example", "clean")
	if err != nil {
		return err
	}
	err = sh.Run("packr2", "clean")
	if err != nil {
		return err
	}

	return nil
}
