# Key-generator

[![Build Status](https://travis-ci.org/katena-chain/key-generator.svg?branch=master)](https://travis-ci.org/katena-chain/key-generator)

## Requirements

This project uses [golang](https://golang.org/).

**To compile on Ubuntu/Debian, you may need to install the ```libgl1-mesa-dev``` and ```xorg-dev``` packages, and have GCC installed.**


### Tested versions

In order to run the project properly, some tools are required:

- [golang](https://golang.org/) (Tested: v1.12.6)

## Installation

Refer to the cross-compiling section below.

## Using the tool

For the GUI :

Install go-bindata:

```bash
go get -u github.com/go-bindata/go-bindata
```

Generate assets:

```bash
go generate gui/main.go
```

```bash
build/key-generator-gui-[linux-amd64] //[replace with your OS]
```

For the CLI :
```bash
build/goreleaser/cli-build_[linux_amd64]/key-generator-cli //[replace with your OS]
```

### Commands

| **Command** | **Usage** | **Possible flag** | **Example** | 
|--|--|--|--|
**gen-ed** | *Generates an ED25519 key pair.* | ```--save filePath``` to save the keys to a given file | ```bash build/key-generator genEd --save ~/filePath```|
|**gen-x** | *Generates an ED25519 key pair.* | ```--save filePath``` to save the keys to a given file | ```bash build/key-generator genX --save ~/filePath```|


## Releases

You'll find the release binaries under the ``build`` folder. Run it like the above using the path corresponding path.

### Cross-compiling

To cross-compile  for Windows and MacOS, build the Dockerfile and use the corresponding image.

For example, from the root, run :
```bash
docker run -v ${PWD}:/app transchain/golang-crosscompile:v1.0.0 goreleaser --rm-dist --skip-publish --snapshot
```
