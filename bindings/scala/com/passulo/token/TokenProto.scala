// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO3

package com.passulo.token

object TokenProto extends _root_.scalapb.GeneratedFileObject {
  lazy val dependencies: Seq[_root_.scalapb.GeneratedFileObject] = Seq(
    com.google.protobuf.timestamp.TimestampProto
  )
  lazy val messagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_ <: _root_.scalapb.GeneratedMessage]] =
    Seq[_root_.scalapb.GeneratedMessageCompanion[_ <: _root_.scalapb.GeneratedMessage]](
      com.passulo.token.Token
    )
  private lazy val ProtoBytes: _root_.scala.Array[Byte] =
      scalapb.Encoding.fromBase64(scala.collection.immutable.Seq(
  """Chdjb20vcGFzc3Vsby90b2tlbi5wcm90bxIOY29tLnBhc3N1bG8udjEaH2dvb2dsZS9wcm90b2J1Zi90aW1lc3RhbXAucHJvd
  G8i4gQKBVRva2VuEhcKAmlkGAEgASgJQgfiPwQSAmlkUgJpZBIsCglmaXJzdE5hbWUYAiABKAlCDuI/CxIJZmlyc3ROYW1lUglma
  XJzdE5hbWUSLwoKbWlkZGxlTmFtZRgDIAEoCUIP4j8MEgptaWRkbGVOYW1lUgptaWRkbGVOYW1lEikKCGxhc3ROYW1lGAQgASgJQ
  g3iPwoSCGxhc3ROYW1lUghsYXN0TmFtZRIjCgZnZW5kZXIYBSABKAlCC+I/CBIGZ2VuZGVyUgZnZW5kZXISIwoGbnVtYmVyGAYgA
  SgJQgviPwgSBm51bWJlclIGbnVtYmVyEiMKBnN0YXR1cxgHIAEoCUIL4j8IEgZzdGF0dXNSBnN0YXR1cxImCgdjb21wYW55GAggA
  SgJQgziPwkSB2NvbXBhbnlSB2NvbXBhbnkSIAoFZW1haWwYCSABKAlCCuI/BxIFZW1haWxSBWVtYWlsEiwKCXRlbGVwaG9uZRgKI
  AEoCUIO4j8LEgl0ZWxlcGhvbmVSCXRlbGVwaG9uZRIyCgthc3NvY2lhdGlvbhgLIAEoCUIQ4j8NEgthc3NvY2lhdGlvblILYXNzb
  2NpYXRpb24SSwoKdmFsaWRVbnRpbBgMIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBCD+I/DBIKdmFsaWRVbnRpbFIKd
  mFsaWRVbnRpbBJOCgttZW1iZXJTaW5jZRgNIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBCEOI/DRILbWVtYmVyU2luY
  2VSC21lbWJlclNpbmNlQg0KC2NvbS5wYXNzdWxvYgZwcm90bzM="""
      ).mkString)
  lazy val scalaDescriptor: _root_.scalapb.descriptors.FileDescriptor = {
    val scalaProto = com.google.protobuf.descriptor.FileDescriptorProto.parseFrom(ProtoBytes)
    _root_.scalapb.descriptors.FileDescriptor.buildFrom(scalaProto, dependencies.map(_.scalaDescriptor))
  }
  lazy val javaDescriptor: com.google.protobuf.Descriptors.FileDescriptor = {
    val javaProto = com.google.protobuf.DescriptorProtos.FileDescriptorProto.parseFrom(ProtoBytes)
    com.google.protobuf.Descriptors.FileDescriptor.buildFrom(javaProto, _root_.scala.Array(
      com.google.protobuf.timestamp.TimestampProto.javaDescriptor
    ))
  }
  @deprecated("Use javaDescriptor instead. In a future version this will refer to scalaDescriptor.", "ScalaPB 0.5.47")
  def descriptor: com.google.protobuf.Descriptors.FileDescriptor = javaDescriptor
}