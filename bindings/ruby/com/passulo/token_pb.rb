# frozen_string_literal: true
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/passulo/token.proto

require 'google/protobuf'

require 'google/protobuf/timestamp_pb'


descriptor_data = "\n\x17\x63om/passulo/token.proto\x12\x0e\x63om.passulo.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"\xff\x02\n\x05Token\x12\n\n\x02id\x18\x01 \x01(\t\x12\x11\n\tfirstName\x18\x02 \x01(\t\x12\x12\n\nmiddleName\x18\x03 \x01(\t\x12\x10\n\x08lastName\x18\x04 \x01(\t\x12,\n\x06gender\x18\x05 \x01(\x0e\x32\x1c.com.passulo.v1.Token.Gender\x12\x0e\n\x06number\x18\x06 \x01(\t\x12\x0e\n\x06status\x18\x07 \x01(\t\x12\x0f\n\x07\x63ompany\x18\x08 \x01(\t\x12\r\n\x05\x65mail\x18\t \x01(\t\x12\x11\n\ttelephone\x18\n \x01(\t\x12\x13\n\x0b\x61ssociation\x18\x0b \x01(\t\x12.\n\nvalidUntil\x18\x0c \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12/\n\x0bmemberSince\x18\r \x01(\x0b\x32\x1a.google.protobuf.Timestamp\":\n\x06Gender\x12\r\n\tundefined\x10\x00\x12\n\n\x06\x66\x65male\x10\x01\x12\x08\n\x04male\x10\x02\x12\x0b\n\x07\x64iverse\x10\x03\x42\'\n\x0b\x63om.passuloZ\x18github.com/passulo/Tokenb\x06proto3"

pool = Google::Protobuf::DescriptorPool.generated_pool

begin
  pool.add_serialized_file(descriptor_data)
rescue TypeError => e
  # Compatibility code: will be removed in the next major version.
  require 'google/protobuf/descriptor_pb'
  parsed = Google::Protobuf::FileDescriptorProto.decode(descriptor_data)
  parsed.clear_dependency
  serialized = parsed.class.encode(parsed)
  file = pool.add_serialized_file(serialized)
  warn "Warning: Protobuf detected an import path issue while loading generated file #{__FILE__}"
  imports = [
    ["google.protobuf.Timestamp", "google/protobuf/timestamp.proto"],
  ]
  imports.each do |type_name, expected_filename|
    import_file = pool.lookup(type_name).file_descriptor
    if import_file.name != expected_filename
      warn "- #{file.name} imports #{expected_filename}, but that import was loaded as #{import_file.name}"
    end
  end
  warn "Each proto file must use a consistent fully-qualified name."
  warn "This will become an error in the next major version."
end

module Com
  module Passulo
    module V1
      Token = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("com.passulo.v1.Token").msgclass
      Token::Gender = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("com.passulo.v1.Token.Gender").enummodule
    end
  end
end
