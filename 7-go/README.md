# Protocol buffers in go

### Installation

```shell
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/golang/protobuf/proto
```

### usage

create a `src` director which will contain all proto files

### generator

create `generate.sh` (`generate.bat` for windows). Which will build proto files. Paste this command inside that script

```shell
protoc -I=src/ --go_out=src/ src/simple/simple.proto
```

to rename package name in generated go file add an option in `.proto` file, Given example of `simple/simple.proto` file. Here when we add this entire after package name, in the auto generated file _simplepb_ will be it's package name

```proto
option go_package = "simple;simplepb";
```
