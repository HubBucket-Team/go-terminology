// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package snomed

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SctID struct {
	Identifier           int64    `protobuf:"varint,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SctID) Reset()         { *m = SctID{} }
func (m *SctID) String() string { return proto.CompactTextString(m) }
func (*SctID) ProtoMessage()    {}
func (*SctID) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_12d48f34f0349cdf, []int{0}
}
func (m *SctID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SctID.Unmarshal(m, b)
}
func (m *SctID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SctID.Marshal(b, m, deterministic)
}
func (dst *SctID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SctID.Merge(dst, src)
}
func (m *SctID) XXX_Size() int {
	return xxx_messageInfo_SctID.Size(m)
}
func (m *SctID) XXX_DiscardUnknown() {
	xxx_messageInfo_SctID.DiscardUnknown(m)
}

var xxx_messageInfo_SctID proto.InternalMessageInfo

func (m *SctID) GetIdentifier() int64 {
	if m != nil {
		return m.Identifier
	}
	return 0
}

type ReferenceSetItemID struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReferenceSetItemID) Reset()         { *m = ReferenceSetItemID{} }
func (m *ReferenceSetItemID) String() string { return proto.CompactTextString(m) }
func (*ReferenceSetItemID) ProtoMessage()    {}
func (*ReferenceSetItemID) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_12d48f34f0349cdf, []int{1}
}
func (m *ReferenceSetItemID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReferenceSetItemID.Unmarshal(m, b)
}
func (m *ReferenceSetItemID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReferenceSetItemID.Marshal(b, m, deterministic)
}
func (dst *ReferenceSetItemID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReferenceSetItemID.Merge(dst, src)
}
func (m *ReferenceSetItemID) XXX_Size() int {
	return xxx_messageInfo_ReferenceSetItemID.Size(m)
}
func (m *ReferenceSetItemID) XXX_DiscardUnknown() {
	xxx_messageInfo_ReferenceSetItemID.DiscardUnknown(m)
}

var xxx_messageInfo_ReferenceSetItemID proto.InternalMessageInfo

func (m *ReferenceSetItemID) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func init() {
	proto.RegisterType((*SctID)(nil), "snomed.SctID")
	proto.RegisterType((*ReferenceSetItemID)(nil), "snomed.ReferenceSetItemID")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SnomedCTClient is the client API for SnomedCT service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SnomedCTClient interface {
	GetConcept(ctx context.Context, in *SctID, opts ...grpc.CallOption) (*Concept, error)
	// GetExtendedConcept returns the concept with the specified identifier.
	// The preferred description will be determined by language preferences
	// defined at runtime.
	// For example, the header accept-language may be used to define language preferences
	// using tags as per format defined by IETF (http://www.ietf.org/rfc/rfc2616.txt)
	// or by setting at a server-wide basis.
	GetExtendedConcept(ctx context.Context, in *SctID, opts ...grpc.CallOption) (*ExtendedConcept, error)
	// GetDescriptions returns descriptions for a given concept.
	GetDescriptions(ctx context.Context, in *SctID, opts ...grpc.CallOption) (*ConceptDescriptions, error)
	// GetReferenceSets returns the reference sets to which this concept is a member
	GetReferenceSets(ctx context.Context, in *SctID, opts ...grpc.CallOption) (SnomedCT_GetReferenceSetsClient, error)
	// GetAllChildren returns all children of the specified concept
	GetAllChildren(ctx context.Context, in *SctID, opts ...grpc.CallOption) (SnomedCT_GetAllChildrenClient, error)
	// GetDescription returns a single description, by identifier
	GetDescription(ctx context.Context, in *SctID, opts ...grpc.CallOption) (*Description, error)
	// GetReferenceSetItem returns a single item from a reference set, by identifier
	GetReferenceSetItem(ctx context.Context, in *ReferenceSetItemID, opts ...grpc.CallOption) (*ReferenceSetItem, error)
	// CrossMap translates from SNOMED CT to an alternative coding system via a map reference set
	CrossMap(ctx context.Context, in *CrossMapRequest, opts ...grpc.CallOption) (SnomedCT_CrossMapClient, error)
	// FromCrossMap translates from an external coding system to SNOMED-CT.
	FromCrossMap(ctx context.Context, in *TranslateFromRequest, opts ...grpc.CallOption) (*TranslateFromResponse, error)
	// Map translates a SNOMED CT concept into the best match within the specified reference set
	Map(ctx context.Context, in *MapRequest, opts ...grpc.CallOption) (*MapResponse, error)
	// Subsumes determines whether one concept subsumes another
	// This is an implementation of the HL7 FHIR terminology service subsumes method
	// (https://www.hl7.org/fhir/terminology-service.html)
	Subsumes(ctx context.Context, in *SubsumptionRequest, opts ...grpc.CallOption) (*SubsumptionResponse, error)
	// Parse parses a SNOMED expression (compositional grammar)
	Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*Expression, error)
	// Refinements returns the appropriate refinements for this specified concept
	Refinements(ctx context.Context, in *RefinementRequest, opts ...grpc.CallOption) (*RefinementResponse, error)
}

type snomedCTClient struct {
	cc *grpc.ClientConn
}

func NewSnomedCTClient(cc *grpc.ClientConn) SnomedCTClient {
	return &snomedCTClient{cc}
}

func (c *snomedCTClient) GetConcept(ctx context.Context, in *SctID, opts ...grpc.CallOption) (*Concept, error) {
	out := new(Concept)
	err := c.cc.Invoke(ctx, "/snomed.SnomedCT/GetConcept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snomedCTClient) GetExtendedConcept(ctx context.Context, in *SctID, opts ...grpc.CallOption) (*ExtendedConcept, error) {
	out := new(ExtendedConcept)
	err := c.cc.Invoke(ctx, "/snomed.SnomedCT/GetExtendedConcept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snomedCTClient) GetDescriptions(ctx context.Context, in *SctID, opts ...grpc.CallOption) (*ConceptDescriptions, error) {
	out := new(ConceptDescriptions)
	err := c.cc.Invoke(ctx, "/snomed.SnomedCT/GetDescriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snomedCTClient) GetReferenceSets(ctx context.Context, in *SctID, opts ...grpc.CallOption) (SnomedCT_GetReferenceSetsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SnomedCT_serviceDesc.Streams[0], "/snomed.SnomedCT/GetReferenceSets", opts...)
	if err != nil {
		return nil, err
	}
	x := &snomedCTGetReferenceSetsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SnomedCT_GetReferenceSetsClient interface {
	Recv() (*ReferenceSetItem, error)
	grpc.ClientStream
}

type snomedCTGetReferenceSetsClient struct {
	grpc.ClientStream
}

func (x *snomedCTGetReferenceSetsClient) Recv() (*ReferenceSetItem, error) {
	m := new(ReferenceSetItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *snomedCTClient) GetAllChildren(ctx context.Context, in *SctID, opts ...grpc.CallOption) (SnomedCT_GetAllChildrenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SnomedCT_serviceDesc.Streams[1], "/snomed.SnomedCT/GetAllChildren", opts...)
	if err != nil {
		return nil, err
	}
	x := &snomedCTGetAllChildrenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SnomedCT_GetAllChildrenClient interface {
	Recv() (*ConceptReference, error)
	grpc.ClientStream
}

type snomedCTGetAllChildrenClient struct {
	grpc.ClientStream
}

func (x *snomedCTGetAllChildrenClient) Recv() (*ConceptReference, error) {
	m := new(ConceptReference)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *snomedCTClient) GetDescription(ctx context.Context, in *SctID, opts ...grpc.CallOption) (*Description, error) {
	out := new(Description)
	err := c.cc.Invoke(ctx, "/snomed.SnomedCT/GetDescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snomedCTClient) GetReferenceSetItem(ctx context.Context, in *ReferenceSetItemID, opts ...grpc.CallOption) (*ReferenceSetItem, error) {
	out := new(ReferenceSetItem)
	err := c.cc.Invoke(ctx, "/snomed.SnomedCT/GetReferenceSetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snomedCTClient) CrossMap(ctx context.Context, in *CrossMapRequest, opts ...grpc.CallOption) (SnomedCT_CrossMapClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SnomedCT_serviceDesc.Streams[2], "/snomed.SnomedCT/CrossMap", opts...)
	if err != nil {
		return nil, err
	}
	x := &snomedCTCrossMapClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SnomedCT_CrossMapClient interface {
	Recv() (*ReferenceSetItem, error)
	grpc.ClientStream
}

type snomedCTCrossMapClient struct {
	grpc.ClientStream
}

func (x *snomedCTCrossMapClient) Recv() (*ReferenceSetItem, error) {
	m := new(ReferenceSetItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *snomedCTClient) FromCrossMap(ctx context.Context, in *TranslateFromRequest, opts ...grpc.CallOption) (*TranslateFromResponse, error) {
	out := new(TranslateFromResponse)
	err := c.cc.Invoke(ctx, "/snomed.SnomedCT/FromCrossMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snomedCTClient) Map(ctx context.Context, in *MapRequest, opts ...grpc.CallOption) (*MapResponse, error) {
	out := new(MapResponse)
	err := c.cc.Invoke(ctx, "/snomed.SnomedCT/Map", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snomedCTClient) Subsumes(ctx context.Context, in *SubsumptionRequest, opts ...grpc.CallOption) (*SubsumptionResponse, error) {
	out := new(SubsumptionResponse)
	err := c.cc.Invoke(ctx, "/snomed.SnomedCT/Subsumes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snomedCTClient) Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*Expression, error) {
	out := new(Expression)
	err := c.cc.Invoke(ctx, "/snomed.SnomedCT/Parse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snomedCTClient) Refinements(ctx context.Context, in *RefinementRequest, opts ...grpc.CallOption) (*RefinementResponse, error) {
	out := new(RefinementResponse)
	err := c.cc.Invoke(ctx, "/snomed.SnomedCT/Refinements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnomedCTServer is the server API for SnomedCT service.
type SnomedCTServer interface {
	GetConcept(context.Context, *SctID) (*Concept, error)
	// GetExtendedConcept returns the concept with the specified identifier.
	// The preferred description will be determined by language preferences
	// defined at runtime.
	// For example, the header accept-language may be used to define language preferences
	// using tags as per format defined by IETF (http://www.ietf.org/rfc/rfc2616.txt)
	// or by setting at a server-wide basis.
	GetExtendedConcept(context.Context, *SctID) (*ExtendedConcept, error)
	// GetDescriptions returns descriptions for a given concept.
	GetDescriptions(context.Context, *SctID) (*ConceptDescriptions, error)
	// GetReferenceSets returns the reference sets to which this concept is a member
	GetReferenceSets(*SctID, SnomedCT_GetReferenceSetsServer) error
	// GetAllChildren returns all children of the specified concept
	GetAllChildren(*SctID, SnomedCT_GetAllChildrenServer) error
	// GetDescription returns a single description, by identifier
	GetDescription(context.Context, *SctID) (*Description, error)
	// GetReferenceSetItem returns a single item from a reference set, by identifier
	GetReferenceSetItem(context.Context, *ReferenceSetItemID) (*ReferenceSetItem, error)
	// CrossMap translates from SNOMED CT to an alternative coding system via a map reference set
	CrossMap(*CrossMapRequest, SnomedCT_CrossMapServer) error
	// FromCrossMap translates from an external coding system to SNOMED-CT.
	FromCrossMap(context.Context, *TranslateFromRequest) (*TranslateFromResponse, error)
	// Map translates a SNOMED CT concept into the best match within the specified reference set
	Map(context.Context, *MapRequest) (*MapResponse, error)
	// Subsumes determines whether one concept subsumes another
	// This is an implementation of the HL7 FHIR terminology service subsumes method
	// (https://www.hl7.org/fhir/terminology-service.html)
	Subsumes(context.Context, *SubsumptionRequest) (*SubsumptionResponse, error)
	// Parse parses a SNOMED expression (compositional grammar)
	Parse(context.Context, *ParseRequest) (*Expression, error)
	// Refinements returns the appropriate refinements for this specified concept
	Refinements(context.Context, *RefinementRequest) (*RefinementResponse, error)
}

func RegisterSnomedCTServer(s *grpc.Server, srv SnomedCTServer) {
	s.RegisterService(&_SnomedCT_serviceDesc, srv)
}

func _SnomedCT_GetConcept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SctID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnomedCTServer).GetConcept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snomed.SnomedCT/GetConcept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnomedCTServer).GetConcept(ctx, req.(*SctID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnomedCT_GetExtendedConcept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SctID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnomedCTServer).GetExtendedConcept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snomed.SnomedCT/GetExtendedConcept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnomedCTServer).GetExtendedConcept(ctx, req.(*SctID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnomedCT_GetDescriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SctID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnomedCTServer).GetDescriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snomed.SnomedCT/GetDescriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnomedCTServer).GetDescriptions(ctx, req.(*SctID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnomedCT_GetReferenceSets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SctID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SnomedCTServer).GetReferenceSets(m, &snomedCTGetReferenceSetsServer{stream})
}

type SnomedCT_GetReferenceSetsServer interface {
	Send(*ReferenceSetItem) error
	grpc.ServerStream
}

type snomedCTGetReferenceSetsServer struct {
	grpc.ServerStream
}

func (x *snomedCTGetReferenceSetsServer) Send(m *ReferenceSetItem) error {
	return x.ServerStream.SendMsg(m)
}

func _SnomedCT_GetAllChildren_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SctID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SnomedCTServer).GetAllChildren(m, &snomedCTGetAllChildrenServer{stream})
}

type SnomedCT_GetAllChildrenServer interface {
	Send(*ConceptReference) error
	grpc.ServerStream
}

type snomedCTGetAllChildrenServer struct {
	grpc.ServerStream
}

func (x *snomedCTGetAllChildrenServer) Send(m *ConceptReference) error {
	return x.ServerStream.SendMsg(m)
}

func _SnomedCT_GetDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SctID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnomedCTServer).GetDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snomed.SnomedCT/GetDescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnomedCTServer).GetDescription(ctx, req.(*SctID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnomedCT_GetReferenceSetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReferenceSetItemID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnomedCTServer).GetReferenceSetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snomed.SnomedCT/GetReferenceSetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnomedCTServer).GetReferenceSetItem(ctx, req.(*ReferenceSetItemID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnomedCT_CrossMap_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CrossMapRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SnomedCTServer).CrossMap(m, &snomedCTCrossMapServer{stream})
}

type SnomedCT_CrossMapServer interface {
	Send(*ReferenceSetItem) error
	grpc.ServerStream
}

type snomedCTCrossMapServer struct {
	grpc.ServerStream
}

func (x *snomedCTCrossMapServer) Send(m *ReferenceSetItem) error {
	return x.ServerStream.SendMsg(m)
}

func _SnomedCT_FromCrossMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateFromRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnomedCTServer).FromCrossMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snomed.SnomedCT/FromCrossMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnomedCTServer).FromCrossMap(ctx, req.(*TranslateFromRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnomedCT_Map_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnomedCTServer).Map(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snomed.SnomedCT/Map",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnomedCTServer).Map(ctx, req.(*MapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnomedCT_Subsumes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubsumptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnomedCTServer).Subsumes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snomed.SnomedCT/Subsumes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnomedCTServer).Subsumes(ctx, req.(*SubsumptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnomedCT_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnomedCTServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snomed.SnomedCT/Parse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnomedCTServer).Parse(ctx, req.(*ParseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnomedCT_Refinements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefinementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnomedCTServer).Refinements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snomed.SnomedCT/Refinements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnomedCTServer).Refinements(ctx, req.(*RefinementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SnomedCT_serviceDesc = grpc.ServiceDesc{
	ServiceName: "snomed.SnomedCT",
	HandlerType: (*SnomedCTServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConcept",
			Handler:    _SnomedCT_GetConcept_Handler,
		},
		{
			MethodName: "GetExtendedConcept",
			Handler:    _SnomedCT_GetExtendedConcept_Handler,
		},
		{
			MethodName: "GetDescriptions",
			Handler:    _SnomedCT_GetDescriptions_Handler,
		},
		{
			MethodName: "GetDescription",
			Handler:    _SnomedCT_GetDescription_Handler,
		},
		{
			MethodName: "GetReferenceSetItem",
			Handler:    _SnomedCT_GetReferenceSetItem_Handler,
		},
		{
			MethodName: "FromCrossMap",
			Handler:    _SnomedCT_FromCrossMap_Handler,
		},
		{
			MethodName: "Map",
			Handler:    _SnomedCT_Map_Handler,
		},
		{
			MethodName: "Subsumes",
			Handler:    _SnomedCT_Subsumes_Handler,
		},
		{
			MethodName: "Parse",
			Handler:    _SnomedCT_Parse_Handler,
		},
		{
			MethodName: "Refinements",
			Handler:    _SnomedCT_Refinements_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetReferenceSets",
			Handler:       _SnomedCT_GetReferenceSets_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAllChildren",
			Handler:       _SnomedCT_GetAllChildren_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CrossMap",
			Handler:       _SnomedCT_CrossMap_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server.proto",
}

// SearchClient is the client API for Search service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	Extract(ctx context.Context, in *ExtractRequest, opts ...grpc.CallOption) (*ExtractResponse, error)
	Synonyms(ctx context.Context, in *SynonymRequest, opts ...grpc.CallOption) (Search_SynonymsClient, error)
}

type searchClient struct {
	cc *grpc.ClientConn
}

func NewSearchClient(cc *grpc.ClientConn) SearchClient {
	return &searchClient{cc}
}

func (c *searchClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/snomed.Search/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) Extract(ctx context.Context, in *ExtractRequest, opts ...grpc.CallOption) (*ExtractResponse, error) {
	out := new(ExtractResponse)
	err := c.cc.Invoke(ctx, "/snomed.Search/Extract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchClient) Synonyms(ctx context.Context, in *SynonymRequest, opts ...grpc.CallOption) (Search_SynonymsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Search_serviceDesc.Streams[0], "/snomed.Search/Synonyms", opts...)
	if err != nil {
		return nil, err
	}
	x := &searchSynonymsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Search_SynonymsClient interface {
	Recv() (*SynonymResponseItem, error)
	grpc.ClientStream
}

type searchSynonymsClient struct {
	grpc.ClientStream
}

func (x *searchSynonymsClient) Recv() (*SynonymResponseItem, error) {
	m := new(SynonymResponseItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SearchServer is the server API for Search service.
type SearchServer interface {
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	Extract(context.Context, *ExtractRequest) (*ExtractResponse, error)
	Synonyms(*SynonymRequest, Search_SynonymsServer) error
}

func RegisterSearchServer(s *grpc.Server, srv SearchServer) {
	s.RegisterService(&_Search_serviceDesc, srv)
}

func _Search_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snomed.Search/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_Extract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).Extract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snomed.Search/Extract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).Extract(ctx, req.(*ExtractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Search_Synonyms_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SynonymRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SearchServer).Synonyms(m, &searchSynonymsServer{stream})
}

type Search_SynonymsServer interface {
	Send(*SynonymResponseItem) error
	grpc.ServerStream
}

type searchSynonymsServer struct {
	grpc.ServerStream
}

func (x *searchSynonymsServer) Send(m *SynonymResponseItem) error {
	return x.ServerStream.SendMsg(m)
}

var _Search_serviceDesc = grpc.ServiceDesc{
	ServiceName: "snomed.Search",
	HandlerType: (*SearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Search_Search_Handler,
		},
		{
			MethodName: "Extract",
			Handler:    _Search_Extract_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Synonyms",
			Handler:       _Search_Synonyms_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server.proto",
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_server_12d48f34f0349cdf) }

var fileDescriptor_server_12d48f34f0349cdf = []byte{
	// 714 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x95, 0xdd, 0x52, 0xd3, 0x40,
	0x14, 0xc7, 0x27, 0x30, 0xd4, 0xba, 0xa0, 0xe8, 0xa9, 0x7c, 0xb5, 0xa0, 0x18, 0x1d, 0x45, 0xc5,
	0xa6, 0x7c, 0x78, 0xe3, 0x9d, 0x14, 0xec, 0x70, 0xe1, 0x0c, 0xd3, 0x32, 0x5c, 0x38, 0x2a, 0x86,
	0xe4, 0x14, 0x32, 0x93, 0xec, 0x86, 0xdd, 0x85, 0xa1, 0x32, 0xbd, 0xf1, 0x05, 0xbc, 0xf0, 0xd2,
	0xc7, 0xf2, 0x15, 0x7c, 0x10, 0x27, 0x9b, 0x6c, 0xbb, 0x84, 0x7e, 0x5c, 0xb5, 0xe7, 0x9c, 0xcd,
	0xff, 0x97, 0xf3, 0x3f, 0x9b, 0x5d, 0x32, 0x23, 0x90, 0x5f, 0x22, 0xaf, 0xc6, 0x9c, 0x49, 0x06,
	0x05, 0x41, 0x59, 0x84, 0x7e, 0x79, 0x26, 0xfd, 0x4d, 0xb3, 0xe5, 0xe5, 0x53, 0xc6, 0x4e, 0x43,
	0x74, 0xdc, 0x38, 0x70, 0x5c, 0x4a, 0x99, 0x74, 0x65, 0xc0, 0xa8, 0x48, 0xab, 0xf6, 0x4b, 0x32,
	0xd5, 0xf2, 0xe4, 0xfe, 0x2e, 0x3c, 0x26, 0x24, 0xf0, 0x91, 0xca, 0xa0, 0x1d, 0x20, 0x5f, 0xb4,
	0x56, 0xad, 0xb5, 0xc9, 0xa6, 0x91, 0xb1, 0xb7, 0x09, 0x34, 0xb1, 0x8d, 0x1c, 0xa9, 0x87, 0x2d,
	0x94, 0xfb, 0x12, 0xa3, 0x81, 0x4f, 0xdd, 0x35, 0x9f, 0xda, 0xfc, 0x35, 0x4d, 0x8a, 0x2d, 0xf5,
	0x36, 0xf5, 0x43, 0x38, 0x22, 0xa4, 0x81, 0xb2, 0xce, 0xa8, 0x87, 0xb1, 0x84, 0x7b, 0xd5, 0xec,
	0x35, 0x15, 0xbf, 0x3c, 0xab, 0xc3, 0xac, 0x6e, 0xaf, 0xfd, 0xfc, 0xfb, 0xef, 0xf7, 0x84, 0x0d,
	0xab, 0xce, 0xe5, 0x86, 0x93, 0xd6, 0x1c, 0x2f, 0xad, 0x09, 0xe7, 0xba, 0xcf, 0xe8, 0x02, 0x23,
	0xd0, 0x40, 0xb9, 0x77, 0x25, 0x91, 0xfa, 0xe8, 0x0f, 0xd1, 0x5f, 0xd0, 0x61, 0x6e, 0x9d, 0xbd,
	0xa1, 0x38, 0x6f, 0xe0, 0xd5, 0x38, 0x8e, 0x83, 0xd9, 0x93, 0x20, 0xc9, 0x6c, 0x03, 0xe5, 0x2e,
	0x0a, 0x8f, 0x07, 0xb1, 0x72, 0x33, 0x4f, 0xab, 0xe4, 0xba, 0x31, 0xd7, 0xda, 0xef, 0x14, 0xd1,
	0x81, 0xb7, 0x63, 0x89, 0xbe, 0x89, 0x60, 0xe4, 0x41, 0x03, 0xa5, 0x39, 0x84, 0x5b, 0xd8, 0x45,
	0x1d, 0xe6, 0x47, 0x65, 0xd7, 0x14, 0xf3, 0x35, 0xac, 0x8d, 0x65, 0x72, 0x6c, 0x0b, 0x94, 0xa2,
	0x66, 0xc1, 0x39, 0xb9, 0xdf, 0x40, 0xf9, 0x21, 0x0c, 0xeb, 0x67, 0x41, 0xe8, 0x73, 0xa4, 0x43,
	0x71, 0x59, 0x97, 0x3d, 0xaa, 0xbd, 0xad, 0x70, 0x55, 0x58, 0x1f, 0x8b, 0x73, 0xfb, 0xf2, 0x35,
	0x0b, 0x4e, 0x14, 0xd2, 0x70, 0x2b, 0x8f, 0x2c, 0xe9, 0xd0, 0x58, 0x63, 0xaf, 0x2b, 0xda, 0x0b,
	0x78, 0x6e, 0xd0, 0x4c, 0xeb, 0x6e, 0x6e, 0x97, 0x0e, 0x29, 0xe5, 0x7c, 0x4c, 0x1c, 0x82, 0xf2,
	0x30, 0xef, 0x46, 0xfa, 0x3a, 0x08, 0x9d, 0x3a, 0x78, 0x1c, 0x48, 0x8c, 0x72, 0x68, 0x4e, 0x8a,
	0x75, 0xce, 0x84, 0xf8, 0xe4, 0xc6, 0xd0, 0xdb, 0x90, 0x3a, 0xd3, 0xc4, 0xf3, 0x0b, 0x14, 0x72,
	0x04, 0x6c, 0xf4, 0x56, 0xcd, 0xfe, 0x1d, 0x07, 0x7e, 0xd7, 0xf1, 0x12, 0xcd, 0xc8, 0x8d, 0x6b,
	0x16, 0x74, 0xc8, 0xcc, 0x47, 0xce, 0xa2, 0x1e, 0x77, 0x59, 0xcb, 0x1f, 0x72, 0x97, 0x8a, 0xd0,
	0x95, 0x98, 0x94, 0x35, 0x7c, 0x65, 0x48, 0x55, 0xc4, 0x8c, 0x0a, 0x1c, 0xd8, 0xae, 0x66, 0x09,
	0xe7, 0x5a, 0x77, 0xee, 0x77, 0x9d, 0x6b, 0xd1, 0x85, 0x2f, 0x64, 0x32, 0x21, 0x82, 0xd6, 0x34,
	0x9a, 0x2c, 0xdd, 0xc8, 0x8d, 0x52, 0x1f, 0xd8, 0x5f, 0xe4, 0xc6, 0xf0, 0x9d, 0x14, 0x5b, 0x17,
	0x27, 0xe2, 0x22, 0x42, 0xd1, 0x1f, 0x5e, 0x9a, 0x51, 0x63, 0xd7, 0xa8, 0xca, 0xc0, 0x5a, 0x86,
	0xac, 0x28, 0xe4, 0x1c, 0x94, 0x0c, 0xa4, 0xd0, 0xaa, 0x47, 0x64, 0xea, 0xc0, 0xe5, 0x02, 0xe1,
	0x91, 0x96, 0x50, 0xa1, 0x16, 0x86, 0xfe, 0x91, 0x12, 0x73, 0x14, 0x22, 0xd9, 0x8a, 0xcf, 0x94,
	0xde, 0x0a, 0x54, 0x0c, 0x3d, 0xec, 0x95, 0x9d, 0x58, 0xc9, 0xfd, 0x20, 0xd3, 0x4d, 0x6c, 0x07,
	0x14, 0x23, 0xa4, 0x52, 0xc0, 0x92, 0x31, 0xf0, 0x2c, 0xa9, 0x11, 0xe5, 0x41, 0xa5, 0xec, 0xd5,
	0x47, 0x7f, 0x63, 0xa6, 0x5b, 0xbc, 0x0f, 0xdb, 0xfc, 0x33, 0x41, 0x0a, 0x2d, 0x74, 0xb9, 0x77,
	0x06, 0xcd, 0xde, 0xbf, 0xb9, 0x9e, 0x45, 0x2a, 0xd6, 0xf4, 0xf9, 0x7c, 0x3a, 0x23, 0x2f, 0x29,
	0x72, 0x09, 0x1e, 0x9a, 0xa6, 0xa5, 0x4a, 0x5f, 0xc9, 0x9d, 0xbd, 0x2b, 0xc9, 0x5d, 0x4f, 0xc2,
	0xbc, 0x71, 0xe2, 0x26, 0x09, 0xad, 0xba, 0x70, 0x2b, 0x9f, 0xc9, 0x3e, 0x55, 0xb2, 0x15, 0x7b,
	0xde, 0x90, 0xa5, 0x61, 0x9c, 0x9c, 0xbb, 0xc9, 0xba, 0xf7, 0x96, 0x80, 0x6f, 0xa4, 0xd8, 0xea,
	0x50, 0x46, 0x3b, 0x91, 0xe8, 0xeb, 0x67, 0x99, 0xdb, 0xf3, 0xd6, 0xf9, 0x54, 0x5f, 0x7d, 0x42,
	0x03, 0xe7, 0x9d, 0x29, 0xd6, 0xac, 0x9d, 0x2d, 0xf2, 0xc4, 0x63, 0x51, 0x15, 0x43, 0x9f, 0x07,
	0x57, 0x55, 0x89, 0x3c, 0x0a, 0x28, 0x0b, 0xd9, 0x69, 0x27, 0xd3, 0xf4, 0xe4, 0x4e, 0xa1, 0xa5,
	0xee, 0xdc, 0x03, 0xeb, 0x73, 0x76, 0xdf, 0x9e, 0x14, 0xd4, 0x55, 0xba, 0xf5, 0x3f, 0x00, 0x00,
	0xff, 0xff, 0xd1, 0x02, 0xae, 0x6e, 0x8e, 0x07, 0x00, 0x00,
}
