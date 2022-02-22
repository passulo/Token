//Generated by the protocol buffer compiler. DO NOT EDIT!
// source: com/passulo/token.proto

package com.passulo;

@kotlin.jvm.JvmSynthetic
public inline fun token(block: com.passulo.TokenKt.Dsl.() -> kotlin.Unit): com.passulo.TokenOuterClass.Token =
  com.passulo.TokenKt.Dsl._create(com.passulo.TokenOuterClass.Token.newBuilder()).apply { block() }._build()
public object TokenKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: com.passulo.TokenOuterClass.Token.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: com.passulo.TokenOuterClass.Token.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): com.passulo.TokenOuterClass.Token = _builder.build()

    /**
     * <pre>
     * identifier for this token only
     * </pre>
     *
     * <code>string id = 1;</code>
     */
    public var id: kotlin.String
      @JvmName("getId")
      get() = _builder.getId()
      @JvmName("setId")
      set(value) {
        _builder.setId(value)
      }
    /**
     * <pre>
     * identifier for this token only
     * </pre>
     *
     * <code>string id = 1;</code>
     */
    public fun clearId() {
      _builder.clearId()
    }

    /**
     * <code>string firstName = 2;</code>
     */
    public var firstName: kotlin.String
      @JvmName("getFirstName")
      get() = _builder.getFirstName()
      @JvmName("setFirstName")
      set(value) {
        _builder.setFirstName(value)
      }
    /**
     * <code>string firstName = 2;</code>
     */
    public fun clearFirstName() {
      _builder.clearFirstName()
    }

    /**
     * <code>string middleName = 3;</code>
     */
    public var middleName: kotlin.String
      @JvmName("getMiddleName")
      get() = _builder.getMiddleName()
      @JvmName("setMiddleName")
      set(value) {
        _builder.setMiddleName(value)
      }
    /**
     * <code>string middleName = 3;</code>
     */
    public fun clearMiddleName() {
      _builder.clearMiddleName()
    }

    /**
     * <code>string lastName = 4;</code>
     */
    public var lastName: kotlin.String
      @JvmName("getLastName")
      get() = _builder.getLastName()
      @JvmName("setLastName")
      set(value) {
        _builder.setLastName(value)
      }
    /**
     * <code>string lastName = 4;</code>
     */
    public fun clearLastName() {
      _builder.clearLastName()
    }

    /**
     * <code>.com.passulo.v1.Token.Gender gender = 5;</code>
     */
    public var gender: com.passulo.TokenOuterClass.Token.Gender
      @JvmName("getGender")
      get() = _builder.getGender()
      @JvmName("setGender")
      set(value) {
        _builder.setGender(value)
      }
    /**
     * <code>.com.passulo.v1.Token.Gender gender = 5;</code>
     */
    public fun clearGender() {
      _builder.clearGender()
    }

    /**
     * <code>string number = 6;</code>
     */
    public var number: kotlin.String
      @JvmName("getNumber")
      get() = _builder.getNumber()
      @JvmName("setNumber")
      set(value) {
        _builder.setNumber(value)
      }
    /**
     * <code>string number = 6;</code>
     */
    public fun clearNumber() {
      _builder.clearNumber()
    }

    /**
     * <code>string status = 7;</code>
     */
    public var status: kotlin.String
      @JvmName("getStatus")
      get() = _builder.getStatus()
      @JvmName("setStatus")
      set(value) {
        _builder.setStatus(value)
      }
    /**
     * <code>string status = 7;</code>
     */
    public fun clearStatus() {
      _builder.clearStatus()
    }

    /**
     * <code>string company = 8;</code>
     */
    public var company: kotlin.String
      @JvmName("getCompany")
      get() = _builder.getCompany()
      @JvmName("setCompany")
      set(value) {
        _builder.setCompany(value)
      }
    /**
     * <code>string company = 8;</code>
     */
    public fun clearCompany() {
      _builder.clearCompany()
    }

    /**
     * <code>string email = 9;</code>
     */
    public var email: kotlin.String
      @JvmName("getEmail")
      get() = _builder.getEmail()
      @JvmName("setEmail")
      set(value) {
        _builder.setEmail(value)
      }
    /**
     * <code>string email = 9;</code>
     */
    public fun clearEmail() {
      _builder.clearEmail()
    }

    /**
     * <code>string telephone = 10;</code>
     */
    public var telephone: kotlin.String
      @JvmName("getTelephone")
      get() = _builder.getTelephone()
      @JvmName("setTelephone")
      set(value) {
        _builder.setTelephone(value)
      }
    /**
     * <code>string telephone = 10;</code>
     */
    public fun clearTelephone() {
      _builder.clearTelephone()
    }

    /**
     * <code>string association = 11;</code>
     */
    public var association: kotlin.String
      @JvmName("getAssociation")
      get() = _builder.getAssociation()
      @JvmName("setAssociation")
      set(value) {
        _builder.setAssociation(value)
      }
    /**
     * <code>string association = 11;</code>
     */
    public fun clearAssociation() {
      _builder.clearAssociation()
    }

    /**
     * <code>.google.protobuf.Timestamp validUntil = 12;</code>
     */
    public var validUntil: com.google.protobuf.Timestamp
      @JvmName("getValidUntil")
      get() = _builder.getValidUntil()
      @JvmName("setValidUntil")
      set(value) {
        _builder.setValidUntil(value)
      }
    /**
     * <code>.google.protobuf.Timestamp validUntil = 12;</code>
     */
    public fun clearValidUntil() {
      _builder.clearValidUntil()
    }
    /**
     * <code>.google.protobuf.Timestamp validUntil = 12;</code>
     * @return Whether the validUntil field is set.
     */
    public fun hasValidUntil(): kotlin.Boolean {
      return _builder.hasValidUntil()
    }

    /**
     * <code>.google.protobuf.Timestamp memberSince = 13;</code>
     */
    public var memberSince: com.google.protobuf.Timestamp
      @JvmName("getMemberSince")
      get() = _builder.getMemberSince()
      @JvmName("setMemberSince")
      set(value) {
        _builder.setMemberSince(value)
      }
    /**
     * <code>.google.protobuf.Timestamp memberSince = 13;</code>
     */
    public fun clearMemberSince() {
      _builder.clearMemberSince()
    }
    /**
     * <code>.google.protobuf.Timestamp memberSince = 13;</code>
     * @return Whether the memberSince field is set.
     */
    public fun hasMemberSince(): kotlin.Boolean {
      return _builder.hasMemberSince()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun com.passulo.TokenOuterClass.Token.copy(block: com.passulo.TokenKt.Dsl.() -> kotlin.Unit): com.passulo.TokenOuterClass.Token =
  com.passulo.TokenKt.Dsl._create(this.toBuilder()).apply { block() }._build()
