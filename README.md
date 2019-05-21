# Protoc Plugin Boilerplate

## Overview

This library serves as a boilerplate to build code generators with protobuf compiler plugins.  
At Bifrost, we use quite a few code generators which are all protobuf compiler plugins and this library is the boilerplate that we use to build them. 
Using protobufs as a single source of truth has the added benefit of letting us efficiently use different programming languages so our developers can use the language they fancy to work on their projects. 
`github.com/lyft/protoc-gen-star` is the library that is in this boilerplate

## Google WellKnown Types

Most often we found ourselves using Google Well-Known Types as defined in [`https://developers.google.com/protocol-buffers/docs/reference/google.protobuf`](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf). the following table is the filename (proto import) to package (go import) for `github.com/gogo/protobuf`
 
| File Name                               | Package                                             |
|-----------------------------------------|-----------------------------------------------------|
| `google/protobuf/any.proto`             | github.com/gogo/protobuf/types                      |
| `google/protobuf/api.proto`             | google.golang.org/genproto/protobuf/api             |
| `google/protobuf/compiler/plugin.proto` | github.com/gogo/protobuf/protoc-gen-gogo/plugin     |
| `google/protobuf/descriptor.proto`      | github.com/gogo/protobuf/protoc-gen-gogo/descriptor |
| `google/protobuf/duration.proto`        | github.com/gogo/protobuf/types                      |
| `google/protobuf/empty.proto`           | github.com/gogo/protobuf/types                      |
| `google/protobuf/field_mask.proto`      | github.com/gogo/protobuf/types                      |
| `google/protobuf/source_context.proto`  | google.golang.org/genproto/protobuf/source_context  |
| `google/protobuf/struct.proto`          | github.com/gogo/protobuf/types                      |
| `google/protobuf/timestamp.proto`       | github.com/gogo/protobuf/types                      |
| `google/protobuf/type.proto`            | google.golang.org/genproto/protobuf/ptype           |
| `google/protobuf/wrappers.proto`        | github.com/gogo/protobuf/types                      |

the following table is the filename (proto import) to package (go import) for `github.com/gogo/protobuf`
keep in mind that when using `protoc-gen-star` library to crete a new plugin, you must use `github.com/golang/protobuf/` packages.

| File Name                               | Package                                             |
|-----------------------------------------|-----------------------------------------------------|
| `google/protobuf/any.proto`             | github.com/golang/protobuf/ptypes/any               |
| `google/protobuf/api.proto`             | google.golang.org/genproto/protobuf/api             |
| `google/protobuf/compiler/plugin.proto` | github.com/golang/protobuf/protoc-gen-go/plugin     |
| `google/protobuf/descriptor.proto`      | github.com/golang/protobuf/protoc-gen-go/descriptor |
| `google/protobuf/duration.proto`        | github.com/golang/protobuf/ptypes/duration          |
| `google/protobuf/empty.proto`           | github.com/golang/protobuf/ptypes/empty             |
| `google/protobuf/field_mask.proto`      | google.golang.org/genproto/protobuf/field_mask      |
| `google/protobuf/source_context.proto`  | google.golang.org/genproto/protobuf/source_context  |
| `google/protobuf/struct.proto`          | github.com/golang/protobuf/ptypes/struct            |
| `google/protobuf/timestamp.proto`       | github.com/golang/protobuf/ptypes/timestamp         |
| `google/protobuf/type.proto`            | google.golang.org/genproto/protobuf/ptype           |
| `google/protobuf/wrappers.proto`        | github.com/golang/protobuf/ptypes/wrappers          |

## Build 

### Dependencies 

We use [`mage`](https://github.com/magefile/mage) as our main build tool and [`packr`](https://github.com/gobuffalo/packr/) to embed static files (the templates).

install `mage` by running  :

```shell
go get -u -d -v github.com/magefile/mage
cd $GOPATH/src/github.com/magefile/mage
go run bootstrap.go
```

install `packr` by running :

```shell
go get -u -v github.com/gobuffalo/packr/v2/...
go get -u -v github.com/gobuffalo/packr/v2/packr2
```

### Commands

- `mage build` : rebuild the plugin and generate examples in `/example`
- `mage clean` : remove generated go files in example folder and clean up the files packer generates 