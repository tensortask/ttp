// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: type.proto

package ttp

/*
Package Name:
*/

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Type int32

const (
	// Not a legal value for DataType.  Used to indicate a DataType field
	// has not been set.
	Type_INVALID Type = 0
	// Data types that all computation devices are expected to be
	// capable to support.
	Type_FLOAT      Type = 1
	Type_DOUBLE     Type = 2
	Type_INT32      Type = 3
	Type_UINT8      Type = 4
	Type_INT16      Type = 5
	Type_INT8       Type = 6
	Type_STRING     Type = 7
	Type_COMPLEX64  Type = 8
	Type_INT64      Type = 9
	Type_BOOL       Type = 10
	Type_QINT8      Type = 11
	Type_QUINT8     Type = 12
	Type_QINT32     Type = 13
	Type_BFLOAT16   Type = 14
	Type_QINT16     Type = 15
	Type_QUINT16    Type = 16
	Type_UINT16     Type = 17
	Type_COMPLEX128 Type = 18
	Type_HALF       Type = 19
	Type_UINT32     Type = 22
	Type_UINT64     Type = 23
)

var Type_name = map[int32]string{
	0:  "INVALID",
	1:  "FLOAT",
	2:  "DOUBLE",
	3:  "INT32",
	4:  "UINT8",
	5:  "INT16",
	6:  "INT8",
	7:  "STRING",
	8:  "COMPLEX64",
	9:  "INT64",
	10: "BOOL",
	11: "QINT8",
	12: "QUINT8",
	13: "QINT32",
	14: "BFLOAT16",
	15: "QINT16",
	16: "QUINT16",
	17: "UINT16",
	18: "COMPLEX128",
	19: "HALF",
	22: "UINT32",
	23: "UINT64",
}
var Type_value = map[string]int32{
	"INVALID":    0,
	"FLOAT":      1,
	"DOUBLE":     2,
	"INT32":      3,
	"UINT8":      4,
	"INT16":      5,
	"INT8":       6,
	"STRING":     7,
	"COMPLEX64":  8,
	"INT64":      9,
	"BOOL":       10,
	"QINT8":      11,
	"QUINT8":     12,
	"QINT32":     13,
	"BFLOAT16":   14,
	"QINT16":     15,
	"QUINT16":    16,
	"UINT16":     17,
	"COMPLEX128": 18,
	"HALF":       19,
	"UINT32":     22,
	"UINT64":     23,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_type_8fbe5ade4163f197, []int{0}
}

func init() {
	proto.RegisterEnum("ttp.Type", Type_name, Type_value)
}

func init() { proto.RegisterFile("type.proto", fileDescriptor_type_8fbe5ade4163f197) }

var fileDescriptor_type_8fbe5ade4163f197 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0x3d, 0x4e, 0xc3, 0x40,
	0x10, 0x85, 0xbd, 0xf9, 0xcf, 0xe4, 0x87, 0x61, 0x91, 0x40, 0xa2, 0x98, 0x82, 0x92, 0x82, 0xc8,
	0x89, 0xb5, 0x4a, 0x1b, 0x93, 0x04, 0x2c, 0x2d, 0x36, 0x01, 0x07, 0x51, 0x47, 0x0a, 0x29, 0x6d,
	0x21, 0x53, 0xa4, 0xe3, 0x02, 0x48, 0x1c, 0x83, 0x23, 0x70, 0x04, 0xca, 0x94, 0x29, 0xe3, 0xf5,
	0x05, 0x28, 0x29, 0xd1, 0xae, 0xed, 0xee, 0xcd, 0x37, 0xef, 0xed, 0x1b, 0x2d, 0x40, 0xb2, 0x8d,
	0xd7, 0x57, 0xf1, 0x6b, 0x94, 0x44, 0xbc, 0x9a, 0x24, 0xf1, 0xf9, 0xc5, 0x26, 0xda, 0x44, 0x03,
	0x03, 0x56, 0x6f, 0x2f, 0x03, 0x3d, 0x99, 0xc1, 0xa8, 0xdc, 0x78, 0xf9, 0x51, 0x81, 0x5a, 0xb8,
	0x8d, 0xd7, 0xbc, 0x03, 0x4d, 0xcf, 0x7f, 0x9a, 0x48, 0x6f, 0x8a, 0x16, 0x6f, 0x43, 0x7d, 0x2e,
	0x83, 0x49, 0x88, 0x8c, 0x03, 0x34, 0xa6, 0xc1, 0xd2, 0x95, 0x33, 0xac, 0x68, 0xec, 0xf9, 0xe1,
	0x68, 0x88, 0x55, 0x2d, 0x97, 0x9e, 0x1f, 0x8e, 0xb1, 0x56, 0x50, 0x5b, 0x60, 0x9d, 0xb7, 0xa0,
	0x66, 0x60, 0x43, 0xc7, 0x1e, 0xc3, 0x07, 0xcf, 0xbf, 0xc1, 0x26, 0xef, 0x41, 0xfb, 0x3a, 0xb8,
	0xbb, 0x97, 0xb3, 0x67, 0xe1, 0x60, 0xab, 0xf0, 0x0b, 0x07, 0xdb, 0xda, 0xef, 0x06, 0x81, 0x44,
	0xd0, 0x70, 0x61, 0xa2, 0x1d, 0x1d, 0x5d, 0xe4, 0x6f, 0x77, 0x8d, 0xce, 0x2b, 0x7b, 0xbc, 0x0b,
	0x2d, 0xd7, 0x5c, 0x65, 0x0b, 0xec, 0x97, 0x1b, 0x5b, 0xe0, 0x91, 0xbe, 0xdd, 0x24, 0x6c, 0x81,
	0xa8, 0x17, 0x85, 0x3e, 0xe6, 0x7d, 0x80, 0xa2, 0xd9, 0x1e, 0x8e, 0x91, 0xeb, 0xbe, 0xdb, 0x89,
	0x9c, 0xe3, 0x49, 0xe9, 0x1a, 0x0d, 0xf1, 0xb4, 0xd4, 0xc2, 0xc1, 0x33, 0xd7, 0xd9, 0xa7, 0x64,
	0x1d, 0x52, 0x62, 0xbf, 0x29, 0xb1, 0xbf, 0x94, 0xd8, 0xbb, 0x22, 0xf6, 0xa5, 0x88, 0x7d, 0x2b,
	0x62, 0x3f, 0x8a, 0xd8, 0x4e, 0x11, 0x3b, 0x28, 0x62, 0x9f, 0x19, 0x59, 0xbb, 0x8c, 0xac, 0x7d,
	0x46, 0xd6, 0xaa, 0x61, 0x3e, 0x73, 0xf4, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x87, 0x8f, 0xea, 0xef,
	0x83, 0x01, 0x00, 0x00,
}
