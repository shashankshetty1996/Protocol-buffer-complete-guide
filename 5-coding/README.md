# installation

This are the following steps to install protobuf in you're system.

### For Mac system use

use homebrew to install protobuf

```shell
brew install protobuf
```

### For Linux

- Please refer [Installation](https://github.com/google/protobuf/releases)

### For Windows

- Please download from [list](https://github.com/google/protobuf/releases)
- Extract all to `C:\Proto3`
- Create two more director as `C:\Proto3\bin` and `C:\Proto3\include`
- Add `C:\Proto\bin` to _path_

### Usage

When we execute above code and it installs we'll get `protoc` as one of command executable.

```shell
protoc -I=proto --python_out=python proto/*.proto
```

Above command will create _python_ code for the same proto.

- Here `protoc` is the executable function.
- `-I=proto` is the director in which proto files are present.
- `--python_out=python` this to inform create python code and dump it inside _python_ directory.
- `proto/*.proto` perform action on all the proto files (or can specify any specific proto file).

Same above can be used for Java code as well. Where we'll us `--java_out` flag and dump compiled code to _java_ directory.

```shell
protoc -I=proto --java_out=java proto/*.proto
```

and for javascript we can use `-js_out` flag

```shell
protoc -I=proto --js_out=javascript proto/*.proto
```
