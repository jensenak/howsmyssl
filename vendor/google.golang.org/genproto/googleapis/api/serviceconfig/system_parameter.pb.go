// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/serviceconfig/system_parameter.proto
// DO NOT EDIT!

package google_api // import "google.golang.org/genproto/googleapis/api/serviceconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ### System parameter configuration
//
// A system parameter is a special kind of parameter defined by the API
// system, not by an individual API. It is typically mapped to an HTTP header
// and/or a URL query parameter. This configuration specifies which methods
// change the names of the system parameters.
type SystemParameters struct {
	// Define system parameters.
	//
	// The parameters defined here will override the default parameters
	// implemented by the system. If this field is missing from the service
	// config, default system parameters will be used. Default system parameters
	// and names is implementation-dependent.
	//
	// Example: define api key and alt name for all methods
	//
	// SystemParameters
	//   rules:
	//     - selector: "*"
	//       parameters:
	//         - name: api_key
	//           url_query_parameter: api_key
	//         - name: alt
	//           http_header: Response-Content-Type
	//
	// Example: define 2 api key names for a specific method.
	//
	// SystemParameters
	//   rules:
	//     - selector: "/ListShelves"
	//       parameters:
	//         - name: api_key
	//           http_header: Api-Key1
	//         - name: api_key
	//           http_header: Api-Key2
	Rules []*SystemParameterRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *SystemParameters) Reset()                    { *m = SystemParameters{} }
func (m *SystemParameters) String() string            { return proto.CompactTextString(m) }
func (*SystemParameters) ProtoMessage()               {}
func (*SystemParameters) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *SystemParameters) GetRules() []*SystemParameterRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Define a system parameter rule mapping system parameter definitions to
// methods.
type SystemParameterRule struct {
	// Selects the methods to which this rule applies. Use '*' to indicate all
	// methods in all APIs.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	// Define parameters. Multiple names may be defined for a parameter.
	// For a given method call, only one of them should be used. If multiple
	// names are used the behavior is implementation-dependent.
	// If none of the specified names are present the behavior is
	// parameter-dependent.
	Parameters []*SystemParameter `protobuf:"bytes,2,rep,name=parameters" json:"parameters,omitempty"`
}

func (m *SystemParameterRule) Reset()                    { *m = SystemParameterRule{} }
func (m *SystemParameterRule) String() string            { return proto.CompactTextString(m) }
func (*SystemParameterRule) ProtoMessage()               {}
func (*SystemParameterRule) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *SystemParameterRule) GetParameters() []*SystemParameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

// Define a parameter's name and location. The parameter may be passed as either
// an HTTP header or a URL query parameter, and if both are passed the behavior
// is implementation-dependent.
type SystemParameter struct {
	// Define the name of the parameter, such as "api_key", "alt", "callback",
	// and etc. It is case sensitive.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Define the HTTP header name to use for the parameter. It is case
	// insensitive.
	HttpHeader string `protobuf:"bytes,2,opt,name=http_header,json=httpHeader" json:"http_header,omitempty"`
	// Define the URL query parameter name to use for the parameter. It is case
	// sensitive.
	UrlQueryParameter string `protobuf:"bytes,3,opt,name=url_query_parameter,json=urlQueryParameter" json:"url_query_parameter,omitempty"`
}

func (m *SystemParameter) Reset()                    { *m = SystemParameter{} }
func (m *SystemParameter) String() string            { return proto.CompactTextString(m) }
func (*SystemParameter) ProtoMessage()               {}
func (*SystemParameter) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func init() {
	proto.RegisterType((*SystemParameters)(nil), "google.api.SystemParameters")
	proto.RegisterType((*SystemParameterRule)(nil), "google.api.SystemParameterRule")
	proto.RegisterType((*SystemParameter)(nil), "google.api.SystemParameter")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/serviceconfig/system_parameter.proto", fileDescriptor13)
}

var fileDescriptor13 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x15, 0x0a, 0x08, 0xae, 0x12, 0x7f, 0x5c, 0x86, 0x08, 0x86, 0xa2, 0x4c, 0x99, 0x1c,
	0x09, 0xc4, 0xc4, 0xd6, 0x09, 0x16, 0x14, 0xc2, 0x07, 0x88, 0x4c, 0x38, 0xdc, 0x48, 0x8e, 0x1d,
	0xce, 0x4e, 0xa5, 0x7e, 0x7b, 0x1c, 0xb7, 0xa4, 0x55, 0x84, 0xba, 0x58, 0xe7, 0x7b, 0xbf, 0xbb,
	0x77, 0x7a, 0xf0, 0x26, 0x8d, 0x91, 0x0a, 0xb9, 0x34, 0x4a, 0x68, 0xc9, 0x0d, 0xc9, 0x4c, 0xa2,
	0x6e, 0xc9, 0x38, 0x93, 0x6d, 0x24, 0xd1, 0xd6, 0x36, 0xf3, 0x4f, 0x66, 0x91, 0x56, 0x75, 0x85,
	0x95, 0xd1, 0xdf, 0xb5, 0xcc, 0xec, 0xda, 0x3a, 0x6c, 0xca, 0x56, 0x90, 0x68, 0xd0, 0x21, 0xf1,
	0x30, 0xc3, 0x60, 0xbb, 0xcf, 0x0f, 0x24, 0xaf, 0x70, 0xf5, 0x11, 0xa8, 0xfc, 0x0f, 0xb2, 0xec,
	0x09, 0x4e, 0xa8, 0x53, 0x68, 0xe3, 0xe8, 0x7e, 0x92, 0x4e, 0x1f, 0xe6, 0x7c, 0xc7, 0xf3, 0x11,
	0x5c, 0x78, 0xae, 0xd8, 0xd0, 0x89, 0x86, 0xd9, 0x3f, 0x2a, 0xbb, 0x85, 0x33, 0x8b, 0x0a, 0x2b,
	0x67, 0xc8, 0x2f, 0x8c, 0xd2, 0xf3, 0x62, 0xf8, 0xb3, 0x67, 0x80, 0xe1, 0x38, 0x1b, 0x1f, 0x05,
	0xbb, 0xbb, 0x43, 0x76, 0x7b, 0x78, 0xb2, 0x82, 0xcb, 0x91, 0xcc, 0x18, 0x1c, 0x6b, 0x5f, 0x6e,
	0x7d, 0x42, 0xcd, 0xe6, 0x30, 0x5d, 0x3a, 0xd7, 0x96, 0x4b, 0x14, 0x5f, 0x48, 0xde, 0xa4, 0x97,
	0xa0, 0x6f, 0xbd, 0x84, 0x0e, 0xe3, 0x30, 0xeb, 0x48, 0x95, 0x3f, 0x1d, 0xd2, 0x7a, 0x97, 0x55,
	0x3c, 0x09, 0xe0, 0xb5, 0x97, 0xde, 0x7b, 0x65, 0x30, 0x59, 0xa4, 0x70, 0x51, 0x99, 0x66, 0xef,
	0xca, 0xc5, 0xcd, 0xe8, 0x8e, 0xbc, 0x8f, 0x39, 0x8f, 0x3e, 0x4f, 0x43, 0xde, 0x8f, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xc6, 0x7c, 0x98, 0x54, 0xc1, 0x01, 0x00, 0x00,
}
