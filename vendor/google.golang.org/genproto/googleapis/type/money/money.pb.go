// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/type/money.proto

/*
Package money is a generated protocol buffer package.

It is generated from these files:
	google/type/money.proto

It has these top-level messages:
	Money
*/
package money

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Represents an amount of money with its currency type.
type Money struct {
	// The 3-letter currency code defined in ISO 4217.
	CurrencyCode string `protobuf:"bytes,1,opt,name=currency_code,json=currencyCode" json:"currency_code,omitempty"`
	// The whole units of the amount.
	// For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
	Units int64 `protobuf:"varint,2,opt,name=units" json:"units,omitempty"`
	// Number of nano (10^-9) units of the amount.
	// The value must be between -999,999,999 and +999,999,999 inclusive.
	// If `units` is positive, `nanos` must be positive or zero.
	// If `units` is zero, `nanos` can be positive, zero, or negative.
	// If `units` is negative, `nanos` must be negative or zero.
	// For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
	Nanos int32 `protobuf:"varint,3,opt,name=nanos" json:"nanos,omitempty"`
}

func (m *Money) Reset()                    { *m = Money{} }
func (m *Money) String() string            { return proto.CompactTextString(m) }
func (*Money) ProtoMessage()               {}
func (*Money) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Money) GetCurrencyCode() string {
	if m != nil {
		return m.CurrencyCode
	}
	return ""
}

func (m *Money) GetUnits() int64 {
	if m != nil {
		return m.Units
	}
	return 0
}

func (m *Money) GetNanos() int32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

func init() {
	proto.RegisterType((*Money)(nil), "google.type.Money")
}

func init() { proto.RegisterFile("google/type/money.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0xcd, 0xcf, 0x4b, 0xad, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0x48, 0xe8, 0x81, 0x24, 0x94, 0x22, 0xb8, 0x58, 0x7d, 0x41,
	0x72, 0x42, 0xca, 0x5c, 0xbc, 0xc9, 0xa5, 0x45, 0x45, 0xa9, 0x79, 0xc9, 0x95, 0xf1, 0xc9, 0xf9,
	0x29, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x3c, 0x30, 0x41, 0xe7, 0xfc, 0x94, 0x54,
	0x21, 0x11, 0x2e, 0xd6, 0xd2, 0xbc, 0xcc, 0x92, 0x62, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xe6, 0x20,
	0x08, 0x07, 0x24, 0x9a, 0x97, 0x98, 0x97, 0x5f, 0x2c, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x1a, 0x04,
	0xe1, 0x38, 0xc5, 0x72, 0xf1, 0x27, 0xe7, 0xe7, 0xea, 0x21, 0x59, 0xe6, 0xc4, 0x05, 0xb6, 0x2a,
	0x00, 0xe4, 0x8a, 0x00, 0xc6, 0x28, 0x33, 0xa8, 0x54, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e,
	0x7e, 0x51, 0xba, 0x7e, 0x7a, 0x6a, 0x1e, 0xd8, 0x8d, 0xfa, 0x10, 0xa9, 0xc4, 0x82, 0xcc, 0x62,
	0x24, 0xf7, 0x5b, 0x83, 0xc9, 0x45, 0x4c, 0xcc, 0xee, 0x21, 0x01, 0x49, 0x6c, 0x60, 0x85, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xc2, 0x5b, 0xa5, 0xe7, 0x00, 0x00, 0x00,
}
