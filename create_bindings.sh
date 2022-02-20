#!/usr/bin/env bash

set -Eeuo pipefail

msg() {
  echo >&2 -e "${1-}"
}

# make sure protoc is installed
if ! command -v protoc &> /dev/null
then
    msg "protoc could not be found - install it via brew install protobuf"
    exit
fi

msg "Testing protobuf definition…"
protoc --encode com.passulo.v1.Token com/passulo/token.proto < test/payload.txt > test/proto.bin
protoc --decode_raw < test/proto.bin &> /dev/null
protoc --decode com.passulo.v1.Token com/passulo/token.proto < test/proto.bin &> /dev/null
msg "✅"

msg "Creating Bindings for Obj-C"
protoc com/passulo/token.proto --objc_out=bindings/objc

msg "Creating Bindings for Java"
protoc com/passulo/token.proto --java_out=bindings/java

msg "Creating Bindings for C++"
protoc com/passulo/token.proto --cpp_out=bindings/cpp

msg "Creating Bindings for Kotlin"
protoc com/passulo/token.proto --kotlin_out=bindings/kotlin

msg "Creating Bindings for JavaScript"
protoc com/passulo/token.proto --js_out=bindings/javascript

msg "Creating Bindings for Python"
protoc com/passulo/token.proto --python_out=bindings/python

msg "Creating Bindings for PHP"
protoc com/passulo/token.proto --php_out=bindings/php

msg "Creating Bindings for Ruby"
protoc com/passulo/token.proto --ruby_out=bindings/ruby

msg "Creating Bindings for Swift"
protoc com/passulo/token.proto --swift_out=bindings/swift

msg "Creating Bindings for Scala"
protoc com/passulo/token.proto --plugin=bindings/scala/protoc-gen-scala --scala_out=bindings/scala

