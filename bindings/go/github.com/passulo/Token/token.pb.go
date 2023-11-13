// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: com/passulo/token.proto

package Token

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Token_Gender int32

const (
	Token_undefined Token_Gender = 0
	Token_female    Token_Gender = 1
	Token_male      Token_Gender = 2
	Token_diverse   Token_Gender = 3
)

// Enum value maps for Token_Gender.
var (
	Token_Gender_name = map[int32]string{
		0: "undefined",
		1: "female",
		2: "male",
		3: "diverse",
	}
	Token_Gender_value = map[string]int32{
		"undefined": 0,
		"female":    1,
		"male":      2,
		"diverse":   3,
	}
)

func (x Token_Gender) Enum() *Token_Gender {
	p := new(Token_Gender)
	*p = x
	return p
}

func (x Token_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Token_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_com_passulo_token_proto_enumTypes[0].Descriptor()
}

func (Token_Gender) Type() protoreflect.EnumType {
	return &file_com_passulo_token_proto_enumTypes[0]
}

func (x Token_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Token_Gender.Descriptor instead.
func (Token_Gender) EnumDescriptor() ([]byte, []int) {
	return file_com_passulo_token_proto_rawDescGZIP(), []int{0, 0}
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // identifier for this token only
	FirstName   string                 `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	MiddleName  string                 `protobuf:"bytes,3,opt,name=middleName,proto3" json:"middleName,omitempty"`
	LastName    string                 `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Gender      Token_Gender           `protobuf:"varint,5,opt,name=gender,proto3,enum=com.passulo.v1.Token_Gender" json:"gender,omitempty"`
	Number      string                 `protobuf:"bytes,6,opt,name=number,proto3" json:"number,omitempty"`
	Status      string                 `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Company     string                 `protobuf:"bytes,8,opt,name=company,proto3" json:"company,omitempty"`
	Email       string                 `protobuf:"bytes,9,opt,name=email,proto3" json:"email,omitempty"`
	Telephone   string                 `protobuf:"bytes,10,opt,name=telephone,proto3" json:"telephone,omitempty"`
	Association string                 `protobuf:"bytes,11,opt,name=association,proto3" json:"association,omitempty"`
	ValidUntil  *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=validUntil,proto3" json:"validUntil,omitempty"`
	MemberSince *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=memberSince,proto3" json:"memberSince,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_passulo_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_com_passulo_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_com_passulo_token_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Token) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Token) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *Token) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Token) GetGender() Token_Gender {
	if x != nil {
		return x.Gender
	}
	return Token_undefined
}

func (x *Token) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Token) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Token) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *Token) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Token) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

func (x *Token) GetAssociation() string {
	if x != nil {
		return x.Association
	}
	return ""
}

func (x *Token) GetValidUntil() *timestamppb.Timestamp {
	if x != nil {
		return x.ValidUntil
	}
	return nil
}

func (x *Token) GetMemberSince() *timestamppb.Timestamp {
	if x != nil {
		return x.MemberSince
	}
	return nil
}

var File_com_passulo_token_proto protoreflect.FileDescriptor

var file_com_passulo_token_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x75, 0x6c, 0x6f, 0x2f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x70,
	0x61, 0x73, 0x73, 0x75, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x03, 0x0a, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x75, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x6e, 0x74,
	0x69, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c,
	0x12, 0x3c, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x22, 0x3a,
	0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0d, 0x0a, 0x09, 0x75, 0x6e, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x66, 0x65, 0x6d, 0x61, 0x6c,
	0x65, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x6d, 0x61, 0x6c, 0x65, 0x10, 0x02, 0x12, 0x0b, 0x0a,
	0x07, 0x64, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x10, 0x03, 0x42, 0x27, 0x0a, 0x0b, 0x63, 0x6f,
	0x6d, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x75, 0x6c, 0x6f, 0x5a, 0x18, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x75, 0x6c, 0x6f, 0x2f, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_passulo_token_proto_rawDescOnce sync.Once
	file_com_passulo_token_proto_rawDescData = file_com_passulo_token_proto_rawDesc
)

func file_com_passulo_token_proto_rawDescGZIP() []byte {
	file_com_passulo_token_proto_rawDescOnce.Do(func() {
		file_com_passulo_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_passulo_token_proto_rawDescData)
	})
	return file_com_passulo_token_proto_rawDescData
}

var file_com_passulo_token_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_passulo_token_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_passulo_token_proto_goTypes = []interface{}{
	(Token_Gender)(0),             // 0: com.passulo.v1.Token.Gender
	(*Token)(nil),                 // 1: com.passulo.v1.Token
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_com_passulo_token_proto_depIdxs = []int32{
	0, // 0: com.passulo.v1.Token.gender:type_name -> com.passulo.v1.Token.Gender
	2, // 1: com.passulo.v1.Token.validUntil:type_name -> google.protobuf.Timestamp
	2, // 2: com.passulo.v1.Token.memberSince:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_com_passulo_token_proto_init() }
func file_com_passulo_token_proto_init() {
	if File_com_passulo_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_passulo_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_passulo_token_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_passulo_token_proto_goTypes,
		DependencyIndexes: file_com_passulo_token_proto_depIdxs,
		EnumInfos:         file_com_passulo_token_proto_enumTypes,
		MessageInfos:      file_com_passulo_token_proto_msgTypes,
	}.Build()
	File_com_passulo_token_proto = out.File
	file_com_passulo_token_proto_rawDesc = nil
	file_com_passulo_token_proto_goTypes = nil
	file_com_passulo_token_proto_depIdxs = nil
}
