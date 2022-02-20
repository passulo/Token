# Protobuf setup

This folder contains the protobuf-definition for the token.

## Generating source files

### C++, C#, Java, JavaScript, Kotlin, Obj-C, PHP, Python, Ruby

These languages are built into the `protoc` compiler. Install the compiler via

```shell
brew install protobuf
```

and then run

```shell
protoc src/main/protobuf/token.proto --objc_out=. --java_out=. # for more languages see protoc --help
```

### Swift

Install the [official plugin](https://github.com/apple/swift-protobuf) via

```shell
brew install swift-protobuf
```

and run

```shell
protoc src/main/protobuf/token.proto --swift_out=.
```

### Scala

Download the latest release from [github.com/scalapb/ScalaPB/](https://github.com/scalapb/ScalaPB/releases) and run

```shell
protoc src/main/protobuf/token.proto --plugin=/path/to/download/protoc-gen-scala --scala_out=src/main/scala
```

(For more info see [this guide](https://scalapb.github.io/docs/scalapbc/#using-scalapb-as-a-proper-protoc-plugin).)
