// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: models.proto

package models

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Ecosystem int32

const (
	Ecosystem_UNKNOWN_ECOSYSTEM Ecosystem = 0
	Ecosystem_Maven             Ecosystem = 1
	Ecosystem_RubyGems          Ecosystem = 2
	Ecosystem_Go                Ecosystem = 3
	Ecosystem_Npm               Ecosystem = 4
	Ecosystem_PyPI              Ecosystem = 5
	Ecosystem_Cargo             Ecosystem = 6
	Ecosystem_NuGet             Ecosystem = 7
	Ecosystem_Packagist         Ecosystem = 8
	Ecosystem_Hex               Ecosystem = 9
	Ecosystem_Pub               Ecosystem = 10
	Ecosystem_CycloneDxSBOM     Ecosystem = 11
	Ecosystem_SpdxSBOM          Ecosystem = 12
)

// Enum value maps for Ecosystem.
var (
	Ecosystem_name = map[int32]string{
		0:  "UNKNOWN_ECOSYSTEM",
		1:  "Maven",
		2:  "RubyGems",
		3:  "Go",
		4:  "Npm",
		5:  "PyPI",
		6:  "Cargo",
		7:  "NuGet",
		8:  "Packagist",
		9:  "Hex",
		10: "Pub",
		11: "CycloneDxSBOM",
		12: "SpdxSBOM",
	}
	Ecosystem_value = map[string]int32{
		"UNKNOWN_ECOSYSTEM": 0,
		"Maven":             1,
		"RubyGems":          2,
		"Go":                3,
		"Npm":               4,
		"PyPI":              5,
		"Cargo":             6,
		"NuGet":             7,
		"Packagist":         8,
		"Hex":               9,
		"Pub":               10,
		"CycloneDxSBOM":     11,
		"SpdxSBOM":          12,
	}
)

func (x Ecosystem) Enum() *Ecosystem {
	p := new(Ecosystem)
	*p = x
	return p
}

func (x Ecosystem) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ecosystem) Descriptor() protoreflect.EnumDescriptor {
	return file_models_proto_enumTypes[0].Descriptor()
}

func (Ecosystem) Type() protoreflect.EnumType {
	return &file_models_proto_enumTypes[0]
}

func (x Ecosystem) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ecosystem.Descriptor instead.
func (Ecosystem) EnumDescriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{0}
}

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ecosystem Ecosystem `protobuf:"varint,1,opt,name=ecosystem,proto3,enum=Ecosystem" json:"ecosystem,omitempty"`
	Name      string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version   string    `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	mi := &file_models_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{0}
}

func (x *Package) GetEcosystem() Ecosystem {
	if x != nil {
		return x.Ecosystem
	}
	return Ecosystem_UNKNOWN_ECOSYSTEM
}

func (x *Package) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Package) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type PackageManifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ecosystem Ecosystem  `protobuf:"varint,1,opt,name=ecosystem,proto3,enum=Ecosystem" json:"ecosystem,omitempty"`
	Path      string     `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Packages  []*Package `protobuf:"bytes,3,rep,name=packages,proto3" json:"packages,omitempty"`
}

func (x *PackageManifest) Reset() {
	*x = PackageManifest{}
	mi := &file_models_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PackageManifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageManifest) ProtoMessage() {}

func (x *PackageManifest) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageManifest.ProtoReflect.Descriptor instead.
func (*PackageManifest) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{1}
}

func (x *PackageManifest) GetEcosystem() Ecosystem {
	if x != nil {
		return x.Ecosystem
	}
	return Ecosystem_UNKNOWN_ECOSYSTEM
}

func (x *PackageManifest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PackageManifest) GetPackages() []*Package {
	if x != nil {
		return x.Packages
	}
	return nil
}

var File_models_proto protoreflect.FileDescriptor

var file_models_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61,
	0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x65, 0x63, 0x6f,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45,
	0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x75, 0x0a, 0x0f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x6e, 0x69,
	0x66, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x09, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45, 0x63, 0x6f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x52, 0x09, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x24, 0x0a, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x08,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x2a, 0xae, 0x01, 0x0a, 0x09, 0x45, 0x63, 0x6f,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x45, 0x43, 0x4f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x4d, 0x61, 0x76, 0x65, 0x6e, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x75, 0x62, 0x79,
	0x47, 0x65, 0x6d, 0x73, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x6f, 0x10, 0x03, 0x12, 0x07,
	0x0a, 0x03, 0x4e, 0x70, 0x6d, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x79, 0x50, 0x49, 0x10,
	0x05, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05,
	0x4e, 0x75, 0x47, 0x65, 0x74, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x69, 0x73, 0x74, 0x10, 0x08, 0x12, 0x07, 0x0a, 0x03, 0x48, 0x65, 0x78, 0x10, 0x09, 0x12,
	0x07, 0x0a, 0x03, 0x50, 0x75, 0x62, 0x10, 0x0a, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x79, 0x63, 0x6c,
	0x6f, 0x6e, 0x65, 0x44, 0x78, 0x53, 0x42, 0x4f, 0x4d, 0x10, 0x0b, 0x12, 0x0c, 0x0a, 0x08, 0x53,
	0x70, 0x64, 0x78, 0x53, 0x42, 0x4f, 0x4d, 0x10, 0x0c, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x66, 0x65, 0x64, 0x65, 0x70, 0x2f,
	0x76, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_proto_rawDescOnce sync.Once
	file_models_proto_rawDescData = file_models_proto_rawDesc
)

func file_models_proto_rawDescGZIP() []byte {
	file_models_proto_rawDescOnce.Do(func() {
		file_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_proto_rawDescData)
	})
	return file_models_proto_rawDescData
}

var file_models_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_models_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_models_proto_goTypes = []any{
	(Ecosystem)(0),          // 0: Ecosystem
	(*Package)(nil),         // 1: Package
	(*PackageManifest)(nil), // 2: PackageManifest
}
var file_models_proto_depIdxs = []int32{
	0, // 0: Package.ecosystem:type_name -> Ecosystem
	0, // 1: PackageManifest.ecosystem:type_name -> Ecosystem
	1, // 2: PackageManifest.packages:type_name -> Package
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_models_proto_init() }
func file_models_proto_init() {
	if File_models_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_models_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_proto_goTypes,
		DependencyIndexes: file_models_proto_depIdxs,
		EnumInfos:         file_models_proto_enumTypes,
		MessageInfos:      file_models_proto_msgTypes,
	}.Build()
	File_models_proto = out.File
	file_models_proto_rawDesc = nil
	file_models_proto_goTypes = nil
	file_models_proto_depIdxs = nil
}
