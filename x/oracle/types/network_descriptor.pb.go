// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sifnode/oracle/v1/network_descriptor.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// NetworkDescriptor is a unique identifier for all chains that Sifchain
// enables.
type NetworkDescriptor int32

const (
	// Not currently in use
	NetworkDescriptor_NETWORK_DESCRIPTOR_UNSPECIFIED NetworkDescriptor = 0
	// https://ethereum.org
	NetworkDescriptor_NETWORK_DESCRIPTOR_ETHEREUM NetworkDescriptor = 1
	// Bitcoin mainnet
	NetworkDescriptor_NETWORK_DESCRIPTOR_BITCOIN NetworkDescriptor = 2
	// https://github.com/ethereum/ropsten
	NetworkDescriptor_NETWORK_DESCRIPTOR_ETHEREUM_TESTNET_ROPSTEN NetworkDescriptor = 3
	// https://www.binance.org
	NetworkDescriptor_NETWORK_DESCRIPTOR_BINANCE_SMART_CHAIN NetworkDescriptor = 56
	// https://testnet.binance.org/
	NetworkDescriptor_NETWORK_DESCRIPTOR_BINANCE_SMART_CHAIN_TESTNET NetworkDescriptor = 97
	// Ganache local testnet
	NetworkDescriptor_NETWORK_DESCRIPTOR_GANACHE NetworkDescriptor = 5777
)

var NetworkDescriptor_name = map[int32]string{
	0:    "NETWORK_DESCRIPTOR_UNSPECIFIED",
	1:    "NETWORK_DESCRIPTOR_ETHEREUM",
	2:    "NETWORK_DESCRIPTOR_BITCOIN",
	3:    "NETWORK_DESCRIPTOR_ETHEREUM_TESTNET_ROPSTEN",
	56:   "NETWORK_DESCRIPTOR_BINANCE_SMART_CHAIN",
	97:   "NETWORK_DESCRIPTOR_BINANCE_SMART_CHAIN_TESTNET",
	5777: "NETWORK_DESCRIPTOR_GANACHE",
}

var NetworkDescriptor_value = map[string]int32{
	"NETWORK_DESCRIPTOR_UNSPECIFIED":                 0,
	"NETWORK_DESCRIPTOR_ETHEREUM":                    1,
	"NETWORK_DESCRIPTOR_BITCOIN":                     2,
	"NETWORK_DESCRIPTOR_ETHEREUM_TESTNET_ROPSTEN":    3,
	"NETWORK_DESCRIPTOR_BINANCE_SMART_CHAIN":         56,
	"NETWORK_DESCRIPTOR_BINANCE_SMART_CHAIN_TESTNET": 97,
	"NETWORK_DESCRIPTOR_GANACHE":                     5777,
}

func (x NetworkDescriptor) String() string {
	return proto.EnumName(NetworkDescriptor_name, int32(x))
}

func (NetworkDescriptor) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_721e8ae3af4d5f0a, []int{0}
}

func init() {
	proto.RegisterEnum("sifnode.oracle.v1.NetworkDescriptor", NetworkDescriptor_name, NetworkDescriptor_value)
}

func init() {
	proto.RegisterFile("sifnode/oracle/v1/network_descriptor.proto", fileDescriptor_721e8ae3af4d5f0a)
}

var fileDescriptor_721e8ae3af4d5f0a = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x93, 0x5e, 0xb8, 0x8b, 0x59, 0x4d, 0x07, 0x57, 0x15, 0xa6, 0xe0, 0xc2, 0x45, 0xc4,
	0x0c, 0xd5, 0x8d, 0xdb, 0x34, 0x39, 0xda, 0x41, 0x3a, 0x29, 0x33, 0x53, 0x04, 0x37, 0x43, 0x9b,
	0xa6, 0x6d, 0x50, 0x3b, 0x25, 0x8d, 0x55, 0xdf, 0x42, 0x5f, 0xc3, 0x27, 0x71, 0xd9, 0xa5, 0x4b,
	0x69, 0x5f, 0x44, 0x6c, 0x8d, 0xab, 0x20, 0xee, 0x0e, 0x9c, 0x8f, 0xef, 0x1c, 0xfe, 0x1f, 0x79,
	0x8b, 0x6c, 0x3c, 0xb3, 0xa3, 0x94, 0xd9, 0x7c, 0x90, 0xdc, 0xa6, 0x6c, 0xd9, 0x62, 0xb3, 0xb4,
	0x78, 0xb0, 0xf9, 0x8d, 0x19, 0xa5, 0x8b, 0x24, 0xcf, 0xe6, 0x85, 0xcd, 0xfd, 0x79, 0x6e, 0x0b,
	0x4b, 0xea, 0xdf, 0xac, 0xbf, 0x63, 0xfd, 0x65, 0xab, 0xb1, 0x37, 0xb1, 0x13, 0xbb, 0xdd, 0xb2,
	0xaf, 0x69, 0x07, 0x7a, 0xaf, 0x35, 0x54, 0x17, 0x3b, 0x4b, 0xf4, 0x23, 0x21, 0x07, 0x88, 0x0a,
	0xd0, 0x57, 0xb1, 0xbc, 0x34, 0x11, 0xa8, 0x50, 0xf2, 0x9e, 0x8e, 0xa5, 0xe9, 0x0b, 0xd5, 0x83,
	0x90, 0x9f, 0x73, 0x88, 0xb0, 0x43, 0x9a, 0x68, 0xbf, 0x82, 0x01, 0xdd, 0x01, 0x09, 0xfd, 0x2e,
	0x76, 0x09, 0x45, 0x8d, 0x0a, 0xa0, 0xcd, 0x75, 0x18, 0x73, 0x81, 0x6b, 0x84, 0xa1, 0xa3, 0x5f,
	0x04, 0x46, 0x83, 0xd2, 0x02, 0xb4, 0x91, 0x71, 0x4f, 0x69, 0x10, 0xf8, 0x1f, 0xf1, 0xd0, 0x61,
	0xa5, 0x50, 0x04, 0x22, 0x04, 0xa3, 0xba, 0x81, 0xd4, 0x26, 0xec, 0x04, 0x5c, 0xe0, 0x33, 0x72,
	0x82, 0xfc, 0xbf, 0xb1, 0xe5, 0x1d, 0x3c, 0x20, 0xcd, 0xca, 0x87, 0x2f, 0x02, 0x11, 0x84, 0x1d,
	0xc0, 0x2f, 0xc7, 0xed, 0xe8, 0x6d, 0x4d, 0xdd, 0xd5, 0x9a, 0xba, 0x1f, 0x6b, 0xea, 0x3e, 0x6f,
	0xa8, 0xb3, 0xda, 0x50, 0xe7, 0x7d, 0x43, 0x9d, 0x6b, 0x6f, 0x92, 0x15, 0xd3, 0xfb, 0xa1, 0x9f,
	0xd8, 0x3b, 0xa6, 0xb2, 0x71, 0x32, 0x1d, 0x64, 0x33, 0x56, 0xf6, 0xf5, 0x58, 0x36, 0x56, 0x3c,
	0xcd, 0xd3, 0xc5, 0xf0, 0xff, 0x36, 0xf9, 0xd3, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0xbc,
	0x0a, 0xeb, 0xd0, 0x01, 0x00, 0x00,
}
