// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/TheThingsNetwork/ttn/api/rights.proto

package ttnpb

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Right is the enum that defines all the different rights to do something in
// the network.
type Right int32

const (
	// RIGHT_INVALID is the invalid right and should not be used.
	// It can denote a parsing error.
	RIGHT_INVALID Right = 0
	// RIGHT_USER_ACCOUNT_READ is the right to read the user's account.
	RIGHT_USER_ACCOUNT_READ Right = 1
	// RIGHT_USER_ACCOUNT_WRITE is the right to edit the user's account.
	RIGHT_USER_ACCOUNT_WRITE Right = 2
	// RIGHT_USER_DELETE is the right to delete the user's account.
	RIGHT_USER_DELETE Right = 3
	// RIGHT_USER_AUTHORIZEDCLIENTS is the right to view and manage the authorized
	// clients of the user: list, authorize and de-authorize clients.
	RIGHT_USER_AUTHORIZEDCLIENTS Right = 4
	// RIGHT_USER_APPLICATIONS_LIST is the right to list all of the applications
	// the user is a collaborator of.
	RIGHT_USER_APPLICATIONS_LIST Right = 5
	// RIGHT_USER_APPLICATIONS_CREATE is the right to create an application in
	// the users account.
	RIGHT_USER_APPLICATIONS_CREATE Right = 6
	// RIGHT_USER_APPLICATIONS is the right to manage applications in the users account.
	RIGHT_USER_APPLICATIONS Right = 8
	// RIGHT_USER_GATEWAYS_LIST is the right to list gateways the user is collaborator of.
	RIGHT_USER_GATEWAYS_LIST Right = 9
	// RIGHT_USER_GATEWAYS_CREATE is the right to register a gateway in the users account.
	RIGHT_USER_GATEWAYS_CREATE Right = 10
	// RIGHT_USER_APPLICATIONS is the right to manage applications in the users account.
	RIGHT_USER_GATEWAYS Right = 11
	// RIGHT_USER_CLIENTS_LIST is the right to list the OAuth clients of the user.
	RIGHT_USER_CLIENTS_LIST Right = 12
	// RIGHT_USER_CLIENTS_CREATE is the right to create a new OAuth client in the users account.
	RIGHT_USER_CLIENTS_CREATE Right = 13
	// RIGHT_USER_CLIENTS_MANAGE is the right to update and delete OAuth clients of the user.
	RIGHT_USER_CLIENTS_MANAGE Right = 14
	// RIGHT_APPLICATION_INFO is the right to see basic information about the application.
	// This does not include API keys.
	RIGHT_APPLICATION_INFO Right = 31
	// RIGHT_APPLICATION_SETTINGS_BASIC is the right to edit basic settings of the application.
	RIGHT_APPLICATION_SETTINGS_BASIC Right = 32
	// RIGHT_APPLICATION_SETTINGS_KEYS is the right to view and edit application
	// API keys.
	RIGHT_APPLICATION_SETTINGS_KEYS Right = 33
	// RIGHT_APPLICATION_SETTINGS_COLLABORATORS is the right to manage the
	// collaborators of the application.
	RIGHT_APPLICATION_SETTINGS_COLLABORATORS Right = 34
	// RIGHT_APPLICATION_DELETE is the right to delete the application.
	RIGHT_APPLICATION_DELETE Right = 35
	// RIGHT_APPLICATION_DEVICES_READ is the right to see the devices of an application.
	RIGHT_APPLICATION_DEVICES_READ Right = 36
	// RIGHT_APPLICATION_DEVICES_WRITE is the right to register devices to an application.
	RIGHT_APPLICATION_DEVICES_WRITE Right = 37
	// RIGHT_APPLICATION_TRAFFIC_READ is the right to read traffic of the
	// application (uplink and downlink).
	RIGHT_APPLICATION_TRAFFIC_READ Right = 38
	// RIGHT_APPLICATION_TRAFFIC_WRITE is the right to write traffic of the
	// application (uplink and downlink).
	RIGHT_APPLICATION_TRAFFIC_WRITE Right = 39
	// RIGHT_GATEWAY_INFO is the right to see basic information about the gateway.
	// This does not include API keys.
	RIGHT_GATEWAY_INFO Right = 51
	// RIGHT_GATEWAY_SETTINGS_BASIC is the right to manage basic settings of the gateway.
	RIGHT_GATEWAY_SETTINGS_BASIC Right = 52
	// RIGHT_GATEWAY_SETTINGS_KEYS is the right to view and edit API keys of the gateway.
	RIGHT_GATEWAY_SETTINGS_KEYS Right = 53
	// RIGHT_GATEWAY_SETTINGS_COLLABORATORS is the right to manage collaborators
	// of the gateway.
	RIGHT_GATEWAY_SETTINGS_COLLABORATORS Right = 54
	// RIGHT_GATEWAY_DELETE is the right to delete the gateway.
	RIGHT_GATEWAY_DELETE Right = 55
	// RIGHT_GATEWAY_TRAFFIC is the right to view traffic of the gateway.
	RIGHT_GATEWAY_TRAFFIC Right = 56
	// RIGHT_GATEWAY_STATUS is the right to view the status of the gateway.
	RIGHT_GATEWAY_STATUS Right = 57
	// RIGHT_GATEWAY_LOCATION is the right to view the location of the gateway.
	RIGHT_GATEWAY_LOCATION Right = 58
)

var Right_name = map[int32]string{
	0:  "RIGHT_INVALID",
	1:  "RIGHT_USER_ACCOUNT_READ",
	2:  "RIGHT_USER_ACCOUNT_WRITE",
	3:  "RIGHT_USER_DELETE",
	4:  "RIGHT_USER_AUTHORIZEDCLIENTS",
	5:  "RIGHT_USER_APPLICATIONS_LIST",
	6:  "RIGHT_USER_APPLICATIONS_CREATE",
	8:  "RIGHT_USER_APPLICATIONS",
	9:  "RIGHT_USER_GATEWAYS_LIST",
	10: "RIGHT_USER_GATEWAYS_CREATE",
	11: "RIGHT_USER_GATEWAYS",
	12: "RIGHT_USER_CLIENTS_LIST",
	13: "RIGHT_USER_CLIENTS_CREATE",
	14: "RIGHT_USER_CLIENTS_MANAGE",
	31: "RIGHT_APPLICATION_INFO",
	32: "RIGHT_APPLICATION_SETTINGS_BASIC",
	33: "RIGHT_APPLICATION_SETTINGS_KEYS",
	34: "RIGHT_APPLICATION_SETTINGS_COLLABORATORS",
	35: "RIGHT_APPLICATION_DELETE",
	36: "RIGHT_APPLICATION_DEVICES_READ",
	37: "RIGHT_APPLICATION_DEVICES_WRITE",
	38: "RIGHT_APPLICATION_TRAFFIC_READ",
	39: "RIGHT_APPLICATION_TRAFFIC_WRITE",
	51: "RIGHT_GATEWAY_INFO",
	52: "RIGHT_GATEWAY_SETTINGS_BASIC",
	53: "RIGHT_GATEWAY_SETTINGS_KEYS",
	54: "RIGHT_GATEWAY_SETTINGS_COLLABORATORS",
	55: "RIGHT_GATEWAY_DELETE",
	56: "RIGHT_GATEWAY_TRAFFIC",
	57: "RIGHT_GATEWAY_STATUS",
	58: "RIGHT_GATEWAY_LOCATION",
}
var Right_value = map[string]int32{
	"RIGHT_INVALID":                            0,
	"RIGHT_USER_ACCOUNT_READ":                  1,
	"RIGHT_USER_ACCOUNT_WRITE":                 2,
	"RIGHT_USER_DELETE":                        3,
	"RIGHT_USER_AUTHORIZEDCLIENTS":             4,
	"RIGHT_USER_APPLICATIONS_LIST":             5,
	"RIGHT_USER_APPLICATIONS_CREATE":           6,
	"RIGHT_USER_APPLICATIONS":                  8,
	"RIGHT_USER_GATEWAYS_LIST":                 9,
	"RIGHT_USER_GATEWAYS_CREATE":               10,
	"RIGHT_USER_GATEWAYS":                      11,
	"RIGHT_USER_CLIENTS_LIST":                  12,
	"RIGHT_USER_CLIENTS_CREATE":                13,
	"RIGHT_USER_CLIENTS_MANAGE":                14,
	"RIGHT_APPLICATION_INFO":                   31,
	"RIGHT_APPLICATION_SETTINGS_BASIC":         32,
	"RIGHT_APPLICATION_SETTINGS_KEYS":          33,
	"RIGHT_APPLICATION_SETTINGS_COLLABORATORS": 34,
	"RIGHT_APPLICATION_DELETE":                 35,
	"RIGHT_APPLICATION_DEVICES_READ":           36,
	"RIGHT_APPLICATION_DEVICES_WRITE":          37,
	"RIGHT_APPLICATION_TRAFFIC_READ":           38,
	"RIGHT_APPLICATION_TRAFFIC_WRITE":          39,
	"RIGHT_GATEWAY_INFO":                       51,
	"RIGHT_GATEWAY_SETTINGS_BASIC":             52,
	"RIGHT_GATEWAY_SETTINGS_KEYS":              53,
	"RIGHT_GATEWAY_SETTINGS_COLLABORATORS":     54,
	"RIGHT_GATEWAY_DELETE":                     55,
	"RIGHT_GATEWAY_TRAFFIC":                    56,
	"RIGHT_GATEWAY_STATUS":                     57,
	"RIGHT_GATEWAY_LOCATION":                   58,
}

func (Right) EnumDescriptor() ([]byte, []int) { return fileDescriptorRights, []int{0} }

func init() {
	proto.RegisterEnum("ttn.v3.Right", Right_name, Right_value)
	golang_proto.RegisterEnum("ttn.v3.Right", Right_name, Right_value)
}
func (x Right) String() string {
	s, ok := Right_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/rights.proto", fileDescriptorRights)
}
func init() {
	golang_proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/rights.proto", fileDescriptorRights)
}

var fileDescriptorRights = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x3d, 0x50, 0x13, 0x4f,
	0x18, 0xc6, 0x77, 0xff, 0x7f, 0x40, 0x5d, 0xc5, 0x59, 0x56, 0xf9, 0x0a, 0xf8, 0x82, 0x80, 0x8a,
	0x8e, 0x12, 0x67, 0xf0, 0xbb, 0x5b, 0x2e, 0x4b, 0xd8, 0xf1, 0xbc, 0x63, 0xee, 0x36, 0x30, 0xd0,
	0x64, 0x8c, 0x83, 0x49, 0x86, 0x91, 0x64, 0xe0, 0xd0, 0x96, 0x32, 0xa5, 0xa5, 0x9d, 0xce, 0xd8,
	0x50, 0x52, 0x52, 0x52, 0x52, 0x52, 0x52, 0x39, 0xdc, 0x5e, 0x43, 0x49, 0x49, 0xe9, 0x90, 0xbb,
	0x4c, 0x2e, 0xe1, 0x82, 0x55, 0xb2, 0xfb, 0xfc, 0xde, 0xe7, 0xde, 0x7d, 0xf6, 0x9d, 0x25, 0xcf,
	0x8b, 0x65, 0xaf, 0xb4, 0x5d, 0x98, 0xf9, 0x54, 0xf9, 0x92, 0x56, 0xa5, 0x35, 0x55, 0x2a, 0x6f,
	0x14, 0xb7, 0xac, 0x35, 0xef, 0x5b, 0x65, 0x73, 0x3d, 0xed, 0x79, 0x1b, 0xe9, 0x8f, 0xd5, 0x72,
	0x7a, 0xb3, 0x5c, 0x2c, 0x79, 0x5b, 0x33, 0xd5, 0xcd, 0x8a, 0x57, 0x61, 0x3d, 0x9e, 0xb7, 0x31,
	0xf3, 0x75, 0x36, 0xf5, 0x2c, 0x56, 0x59, 0xac, 0x14, 0x2b, 0xe9, 0xba, 0x5c, 0xd8, 0xfe, 0x5c,
	0x5f, 0xd5, 0x17, 0xf5, 0x7f, 0x61, 0xd9, 0x93, 0x3f, 0xd7, 0x48, 0xb7, 0x73, 0xe1, 0xc3, 0xfa,
	0x48, 0xaf, 0x23, 0xb3, 0x0b, 0x2a, 0x2f, 0xad, 0x25, 0x6e, 0xca, 0x0c, 0x45, 0x6c, 0x84, 0x0c,
	0x86, 0x5b, 0x39, 0x57, 0x38, 0x79, 0x6e, 0x18, 0x76, 0xce, 0x52, 0x79, 0x47, 0xf0, 0x0c, 0xc5,
	0x6c, 0x94, 0x0c, 0x25, 0x88, 0xcb, 0x8e, 0x54, 0x82, 0xfe, 0xc7, 0xfa, 0x49, 0x5f, 0x4c, 0xcd,
	0x08, 0x53, 0x28, 0x41, 0xff, 0x67, 0xe3, 0x64, 0x34, 0x5e, 0x94, 0x53, 0x0b, 0xb6, 0x23, 0x57,
	0x45, 0xc6, 0x30, 0xa5, 0xb0, 0x94, 0x4b, 0xbb, 0xda, 0x89, 0xc5, 0x45, 0x53, 0x1a, 0x5c, 0x49,
	0xdb, 0x72, 0xf3, 0xa6, 0x74, 0x15, 0xed, 0x66, 0x13, 0x04, 0x3a, 0x11, 0x86, 0x23, 0xb8, 0x12,
	0xb4, 0xa7, 0xbd, 0xf3, 0x18, 0x43, 0xaf, 0xb7, 0x75, 0x9e, 0xe5, 0x4a, 0x2c, 0xf3, 0x95, 0xc8,
	0xfe, 0x06, 0x03, 0x92, 0x4a, 0x52, 0x23, 0x6b, 0xc2, 0x06, 0xc9, 0x9d, 0x04, 0x9d, 0xde, 0x6c,
	0xfb, 0x66, 0x74, 0xa2, 0xd0, 0xf5, 0x16, 0xbb, 0x47, 0x86, 0x13, 0xc4, 0xc8, 0xb4, 0xb7, 0x83,
	0xfc, 0x81, 0x5b, 0x3c, 0x2b, 0xe8, 0x6d, 0x96, 0x22, 0x03, 0xa1, 0x1c, 0x3b, 0x49, 0x5e, 0x5a,
	0xf3, 0x36, 0x1d, 0x63, 0x53, 0x64, 0xfc, 0xb2, 0xe6, 0x0a, 0xa5, 0xa4, 0x95, 0x75, 0xf3, 0x73,
	0xdc, 0x95, 0x06, 0x1d, 0x67, 0x93, 0x64, 0xec, 0x0a, 0xea, 0xbd, 0x58, 0x71, 0xe9, 0x7d, 0xf6,
	0x94, 0x4c, 0x5f, 0x01, 0x19, 0xb6, 0x69, 0xf2, 0x39, 0xdb, 0xe1, 0xca, 0x76, 0x5c, 0x3a, 0xd1,
	0x8c, 0x31, 0x4e, 0x47, 0x37, 0x3d, 0xd9, 0xbc, 0xa5, 0x56, 0x75, 0x49, 0x1a, 0xc2, 0x0d, 0x47,
	0x68, 0x2a, 0xb9, 0xa9, 0x06, 0x13, 0x4e, 0xd2, 0x83, 0x64, 0x23, 0xe5, 0xf0, 0xf9, 0x79, 0x69,
	0x84, 0x46, 0x0f, 0x93, 0x8d, 0x1a, 0x4c, 0x68, 0xf4, 0x88, 0x0d, 0x10, 0x16, 0x42, 0xd1, 0x9d,
	0x85, 0x01, 0xce, 0x36, 0x27, 0xae, 0xb1, 0xdf, 0x16, 0xde, 0x0b, 0x36, 0x46, 0x46, 0x3a, 0x10,
	0xf5, 0xe0, 0x5e, 0xb2, 0x69, 0x32, 0xd5, 0x01, 0x68, 0x0d, 0xed, 0x15, 0x1b, 0x22, 0x77, 0x5b,
	0xc9, 0x28, 0xb0, 0xd7, 0x6c, 0x98, 0xf4, 0xb7, 0x2a, 0x51, 0xff, 0xf4, 0xcd, 0xe5, 0x22, 0x57,
	0x71, 0x95, 0x73, 0xe9, 0xdb, 0xe6, 0x60, 0x34, 0x14, 0xd3, 0x0e, 0x4f, 0x4f, 0xdf, 0xa5, 0xba,
	0x6a, 0xbf, 0x01, 0xcd, 0xfd, 0xc4, 0x87, 0x3e, 0xe0, 0x23, 0x1f, 0xf0, 0xb1, 0x0f, 0xf8, 0xc4,
	0x07, 0x7c, 0xea, 0x03, 0x3a, 0xf3, 0x01, 0x9d, 0xfb, 0x80, 0x77, 0x34, 0xa0, 0x9a, 0x06, 0xb4,
	0xab, 0x01, 0xef, 0x69, 0x40, 0xfb, 0x1a, 0xf0, 0x81, 0x06, 0x7c, 0xa8, 0x01, 0x1f, 0x69, 0xc0,
	0xc7, 0x1a, 0xd0, 0x89, 0x06, 0x7c, 0xaa, 0x01, 0x9d, 0x69, 0xc0, 0xe7, 0x1a, 0xd0, 0x4e, 0x00,
	0xa8, 0x16, 0x00, 0xfe, 0x1e, 0x00, 0xfa, 0x11, 0x00, 0xfe, 0x15, 0x00, 0xda, 0x0d, 0x00, 0xed,
	0x05, 0x80, 0xf7, 0x03, 0xc0, 0x07, 0x01, 0xe0, 0xd5, 0xc7, 0xff, 0x7a, 0xc0, 0xaa, 0xeb, 0xc5,
	0x8b, 0xdf, 0x6a, 0xa1, 0xd0, 0x53, 0x7f, 0x89, 0x66, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf9,
	0xfc, 0x4b, 0xb1, 0xf4, 0x04, 0x00, 0x00,
}
