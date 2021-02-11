# Key-generator

[![Build Status](https://travis-ci.org/katena-chain/key-generator.svg?branch=master)](https://travis-ci.org/katena-chain/key-generator)

## Development requirements (Linux)

In order to run the project properly, some tools are required:

- [golang](https://golang.org/) (Tested: v1.12.6)
- xorg-dev (gui requirement)
- libgl1-mesa-dev (gui requirement)

Install go-bindata: 
```bash
go get -u github.com/go-bindata/go-bindata/...
```

Generate assets:
```bash
go generate gui/main.go
```

## Build GUI (Linux) DEPRECATED

Build the project:
```bash
CGO_ENABLED=1 go build -o build/local/gui_key-generator_Linux_x86_64 ./gui
```

## Build CLI (Linux)

Build the project:
```bash
go build -o build/local/cli_key-generator_Linux_x86_64 ./cli
```

## Cross-compiling DEPRECATED

To cross-compile for Windows and MacOS, you can use our docker image to build the project for different platforms:
```bash
docker run -v ${PWD}:/app transchain/golang-crosscompile:v1.0.0 goreleaser --rm-dist --skip-publish --snapshot
```

This command will create a build/goreleaser folder containing all the binaries.
> Note: the docker image will create the build/goreleaser folder with root permissions