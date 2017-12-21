// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/proto/heimdall.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	common/proto/heimdall.proto

It has these top-level messages:
	RegistRequest
	RegistResponse
	UnRegistRequest
	UnRegistResponse
	LogRequest
	LogResponse
	HeartbeatRequest
	HeartbeatResponse
	RuleRequest
	RuleResponse
	Agent
	Log
	Heartbeat
	Rule
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type RegistRequest struct {
	Agent *Agent `protobuf:"bytes,1,opt,name=agent" json:"agent,omitempty"`
}

func (m *RegistRequest) Reset()                    { *m = RegistRequest{} }
func (m *RegistRequest) String() string            { return proto1.CompactTextString(m) }
func (*RegistRequest) ProtoMessage()               {}
func (*RegistRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegistRequest) GetAgent() *Agent {
	if m != nil {
		return m.Agent
	}
	return nil
}

type RegistResponse struct {
	AgentId string `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
}

func (m *RegistResponse) Reset()                    { *m = RegistResponse{} }
func (m *RegistResponse) String() string            { return proto1.CompactTextString(m) }
func (*RegistResponse) ProtoMessage()               {}
func (*RegistResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegistResponse) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

type UnRegistRequest struct {
	AgentId string `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
}

func (m *UnRegistRequest) Reset()                    { *m = UnRegistRequest{} }
func (m *UnRegistRequest) String() string            { return proto1.CompactTextString(m) }
func (*UnRegistRequest) ProtoMessage()               {}
func (*UnRegistRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UnRegistRequest) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

type UnRegistResponse struct {
}

func (m *UnRegistResponse) Reset()                    { *m = UnRegistResponse{} }
func (m *UnRegistResponse) String() string            { return proto1.CompactTextString(m) }
func (*UnRegistResponse) ProtoMessage()               {}
func (*UnRegistResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type LogRequest struct {
	Log *Log `protobuf:"bytes,1,opt,name=log" json:"log,omitempty"`
}

func (m *LogRequest) Reset()                    { *m = LogRequest{} }
func (m *LogRequest) String() string            { return proto1.CompactTextString(m) }
func (*LogRequest) ProtoMessage()               {}
func (*LogRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LogRequest) GetLog() *Log {
	if m != nil {
		return m.Log
	}
	return nil
}

type LogResponse struct {
}

func (m *LogResponse) Reset()                    { *m = LogResponse{} }
func (m *LogResponse) String() string            { return proto1.CompactTextString(m) }
func (*LogResponse) ProtoMessage()               {}
func (*LogResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type HeartbeatRequest struct {
	Heartbeat *Heartbeat `protobuf:"bytes,1,opt,name=heartbeat" json:"heartbeat,omitempty"`
}

func (m *HeartbeatRequest) Reset()                    { *m = HeartbeatRequest{} }
func (m *HeartbeatRequest) String() string            { return proto1.CompactTextString(m) }
func (*HeartbeatRequest) ProtoMessage()               {}
func (*HeartbeatRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *HeartbeatRequest) GetHeartbeat() *Heartbeat {
	if m != nil {
		return m.Heartbeat
	}
	return nil
}

type HeartbeatResponse struct {
	Lastupdate int32 `protobuf:"varint,1,opt,name=lastupdate" json:"lastupdate,omitempty"`
}

func (m *HeartbeatResponse) Reset()                    { *m = HeartbeatResponse{} }
func (m *HeartbeatResponse) String() string            { return proto1.CompactTextString(m) }
func (*HeartbeatResponse) ProtoMessage()               {}
func (*HeartbeatResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *HeartbeatResponse) GetLastupdate() int32 {
	if m != nil {
		return m.Lastupdate
	}
	return 0
}

type RuleRequest struct {
}

func (m *RuleRequest) Reset()                    { *m = RuleRequest{} }
func (m *RuleRequest) String() string            { return proto1.CompactTextString(m) }
func (*RuleRequest) ProtoMessage()               {}
func (*RuleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type RuleResponse struct {
	Rules []*Rule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *RuleResponse) Reset()                    { *m = RuleResponse{} }
func (m *RuleResponse) String() string            { return proto1.CompactTextString(m) }
func (*RuleResponse) ProtoMessage()               {}
func (*RuleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RuleResponse) GetRules() []*Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

type Agent struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Hostname    string `protobuf:"bytes,2,opt,name=hostname" json:"hostname,omitempty"`
	Version     string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	Hostaddress string `protobuf:"bytes,4,opt,name=hostaddress" json:"hostaddress,omitempty"`
	Deviceinfo  string `protobuf:"bytes,5,opt,name=deviceinfo" json:"deviceinfo,omitempty"`
	Status      int32  `protobuf:"varint,6,opt,name=status" json:"status,omitempty"`
	Health      int32  `protobuf:"varint,7,opt,name=health" json:"health,omitempty"`
	CreatedAt   int32  `protobuf:"varint,8,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	LastcheckAt int32  `protobuf:"varint,9,opt,name=lastcheck_at,json=lastcheckAt" json:"lastcheck_at,omitempty"`
}

func (m *Agent) Reset()                    { *m = Agent{} }
func (m *Agent) String() string            { return proto1.CompactTextString(m) }
func (*Agent) ProtoMessage()               {}
func (*Agent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Agent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Agent) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Agent) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Agent) GetHostaddress() string {
	if m != nil {
		return m.Hostaddress
	}
	return ""
}

func (m *Agent) GetDeviceinfo() string {
	if m != nil {
		return m.Deviceinfo
	}
	return ""
}

func (m *Agent) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Agent) GetHealth() int32 {
	if m != nil {
		return m.Health
	}
	return 0
}

func (m *Agent) GetCreatedAt() int32 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Agent) GetLastcheckAt() int32 {
	if m != nil {
		return m.LastcheckAt
	}
	return 0
}

type Log struct {
	AgentId     string `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	Hostname    string `protobuf:"bytes,2,opt,name=hostname" json:"hostname,omitempty"`
	Version     string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	Hostaddress string `protobuf:"bytes,4,opt,name=hostaddress" json:"hostaddress,omitempty"`
	Start       int64  `protobuf:"varint,5,opt,name=start" json:"start,omitempty"`
	Stop        int64  `protobuf:"varint,6,opt,name=stop" json:"stop,omitempty"`
	Match       int32  `protobuf:"varint,7,opt,name=match" json:"match,omitempty"`
	Ruleid      string `protobuf:"bytes,8,opt,name=ruleid" json:"ruleid,omitempty"`
	Request     []byte `protobuf:"bytes,9,opt,name=request,proto3" json:"request,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto1.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Log) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *Log) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Log) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Log) GetHostaddress() string {
	if m != nil {
		return m.Hostaddress
	}
	return ""
}

func (m *Log) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Log) GetStop() int64 {
	if m != nil {
		return m.Stop
	}
	return 0
}

func (m *Log) GetMatch() int32 {
	if m != nil {
		return m.Match
	}
	return 0
}

func (m *Log) GetRuleid() string {
	if m != nil {
		return m.Ruleid
	}
	return ""
}

func (m *Log) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

type Heartbeat struct {
	AgentId      string `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	Hostname     string `protobuf:"bytes,2,opt,name=hostname" json:"hostname,omitempty"`
	Version      string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	Hostaddress  string `protobuf:"bytes,4,opt,name=hostaddress" json:"hostaddress,omitempty"`
	CountRequest int32  `protobuf:"varint,5,opt,name=count_request,json=countRequest" json:"count_request,omitempty"`
	CountAlert   int32  `protobuf:"varint,6,opt,name=count_alert,json=countAlert" json:"count_alert,omitempty"`
	CountDoubt   int32  `protobuf:"varint,7,opt,name=count_doubt,json=countDoubt" json:"count_doubt,omitempty"`
	Uptime       int32  `protobuf:"varint,8,opt,name=uptime" json:"uptime,omitempty"`
	Load1        string `protobuf:"bytes,9,opt,name=load1" json:"load1,omitempty"`
	Load5        string `protobuf:"bytes,10,opt,name=load5" json:"load5,omitempty"`
	Load15       string `protobuf:"bytes,11,opt,name=load15" json:"load15,omitempty"`
	Memusage     int32  `protobuf:"varint,12,opt,name=memusage" json:"memusage,omitempty"`
	Cpuusage     int32  `protobuf:"varint,13,opt,name=cpuusage" json:"cpuusage,omitempty"`
	Gorutines    int32  `protobuf:"varint,14,opt,name=gorutines" json:"gorutines,omitempty"`
	GcCount      int32  `protobuf:"varint,15,opt,name=gc_count,json=gcCount" json:"gc_count,omitempty"`
	GcTime       int32  `protobuf:"varint,16,opt,name=gc_time,json=gcTime" json:"gc_time,omitempty"`
}

func (m *Heartbeat) Reset()                    { *m = Heartbeat{} }
func (m *Heartbeat) String() string            { return proto1.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()               {}
func (*Heartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Heartbeat) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *Heartbeat) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Heartbeat) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Heartbeat) GetHostaddress() string {
	if m != nil {
		return m.Hostaddress
	}
	return ""
}

func (m *Heartbeat) GetCountRequest() int32 {
	if m != nil {
		return m.CountRequest
	}
	return 0
}

func (m *Heartbeat) GetCountAlert() int32 {
	if m != nil {
		return m.CountAlert
	}
	return 0
}

func (m *Heartbeat) GetCountDoubt() int32 {
	if m != nil {
		return m.CountDoubt
	}
	return 0
}

func (m *Heartbeat) GetUptime() int32 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *Heartbeat) GetLoad1() string {
	if m != nil {
		return m.Load1
	}
	return ""
}

func (m *Heartbeat) GetLoad5() string {
	if m != nil {
		return m.Load5
	}
	return ""
}

func (m *Heartbeat) GetLoad15() string {
	if m != nil {
		return m.Load15
	}
	return ""
}

func (m *Heartbeat) GetMemusage() int32 {
	if m != nil {
		return m.Memusage
	}
	return 0
}

func (m *Heartbeat) GetCpuusage() int32 {
	if m != nil {
		return m.Cpuusage
	}
	return 0
}

func (m *Heartbeat) GetGorutines() int32 {
	if m != nil {
		return m.Gorutines
	}
	return 0
}

func (m *Heartbeat) GetGcCount() int32 {
	if m != nil {
		return m.GcCount
	}
	return 0
}

func (m *Heartbeat) GetGcTime() int32 {
	if m != nil {
		return m.GcTime
	}
	return 0
}

type Rule struct {
	Content string `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
}

func (m *Rule) Reset()                    { *m = Rule{} }
func (m *Rule) String() string            { return proto1.CompactTextString(m) }
func (*Rule) ProtoMessage()               {}
func (*Rule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *Rule) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto1.RegisterType((*RegistRequest)(nil), "proto.RegistRequest")
	proto1.RegisterType((*RegistResponse)(nil), "proto.RegistResponse")
	proto1.RegisterType((*UnRegistRequest)(nil), "proto.UnRegistRequest")
	proto1.RegisterType((*UnRegistResponse)(nil), "proto.UnRegistResponse")
	proto1.RegisterType((*LogRequest)(nil), "proto.LogRequest")
	proto1.RegisterType((*LogResponse)(nil), "proto.LogResponse")
	proto1.RegisterType((*HeartbeatRequest)(nil), "proto.HeartbeatRequest")
	proto1.RegisterType((*HeartbeatResponse)(nil), "proto.HeartbeatResponse")
	proto1.RegisterType((*RuleRequest)(nil), "proto.RuleRequest")
	proto1.RegisterType((*RuleResponse)(nil), "proto.RuleResponse")
	proto1.RegisterType((*Agent)(nil), "proto.Agent")
	proto1.RegisterType((*Log)(nil), "proto.Log")
	proto1.RegisterType((*Heartbeat)(nil), "proto.Heartbeat")
	proto1.RegisterType((*Rule)(nil), "proto.Rule")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HeimdallService service

type HeimdallServiceClient interface {
	Regist(ctx context.Context, in *RegistRequest, opts ...grpc.CallOption) (*RegistResponse, error)
	UnRegist(ctx context.Context, in *UnRegistRequest, opts ...grpc.CallOption) (*UnRegistResponse, error)
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
	Rule(ctx context.Context, in *RuleRequest, opts ...grpc.CallOption) (*RuleResponse, error)
}

type heimdallServiceClient struct {
	cc *grpc.ClientConn
}

func NewHeimdallServiceClient(cc *grpc.ClientConn) HeimdallServiceClient {
	return &heimdallServiceClient{cc}
}

func (c *heimdallServiceClient) Regist(ctx context.Context, in *RegistRequest, opts ...grpc.CallOption) (*RegistResponse, error) {
	out := new(RegistResponse)
	err := grpc.Invoke(ctx, "/proto.HeimdallService/Regist", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heimdallServiceClient) UnRegist(ctx context.Context, in *UnRegistRequest, opts ...grpc.CallOption) (*UnRegistResponse, error) {
	out := new(UnRegistResponse)
	err := grpc.Invoke(ctx, "/proto.HeimdallService/UnRegist", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heimdallServiceClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := grpc.Invoke(ctx, "/proto.HeimdallService/Log", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heimdallServiceClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	out := new(HeartbeatResponse)
	err := grpc.Invoke(ctx, "/proto.HeimdallService/Heartbeat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heimdallServiceClient) Rule(ctx context.Context, in *RuleRequest, opts ...grpc.CallOption) (*RuleResponse, error) {
	out := new(RuleResponse)
	err := grpc.Invoke(ctx, "/proto.HeimdallService/Rule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HeimdallService service

type HeimdallServiceServer interface {
	Regist(context.Context, *RegistRequest) (*RegistResponse, error)
	UnRegist(context.Context, *UnRegistRequest) (*UnRegistResponse, error)
	Log(context.Context, *LogRequest) (*LogResponse, error)
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
	Rule(context.Context, *RuleRequest) (*RuleResponse, error)
}

func RegisterHeimdallServiceServer(s *grpc.Server, srv HeimdallServiceServer) {
	s.RegisterService(&_HeimdallService_serviceDesc, srv)
}

func _HeimdallService_Regist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeimdallServiceServer).Regist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HeimdallService/Regist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeimdallServiceServer).Regist(ctx, req.(*RegistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeimdallService_UnRegist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnRegistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeimdallServiceServer).UnRegist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HeimdallService/UnRegist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeimdallServiceServer).UnRegist(ctx, req.(*UnRegistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeimdallService_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeimdallServiceServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HeimdallService/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeimdallServiceServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeimdallService_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeimdallServiceServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HeimdallService/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeimdallServiceServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeimdallService_Rule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeimdallServiceServer).Rule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HeimdallService/Rule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeimdallServiceServer).Rule(ctx, req.(*RuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HeimdallService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HeimdallService",
	HandlerType: (*HeimdallServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Regist",
			Handler:    _HeimdallService_Regist_Handler,
		},
		{
			MethodName: "UnRegist",
			Handler:    _HeimdallService_UnRegist_Handler,
		},
		{
			MethodName: "Log",
			Handler:    _HeimdallService_Log_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _HeimdallService_Heartbeat_Handler,
		},
		{
			MethodName: "Rule",
			Handler:    _HeimdallService_Rule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/proto/heimdall.proto",
}

func init() { proto1.RegisterFile("common/proto/heimdall.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 717 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x25, 0x49, 0x9d, 0xc4, 0xe3, 0xa4, 0x4d, 0x97, 0xd2, 0x2e, 0xa5, 0x40, 0x6a, 0x2e, 0x15,
	0xa0, 0x54, 0x6d, 0x55, 0x71, 0x42, 0x22, 0xc0, 0xa1, 0x48, 0x3d, 0x19, 0x38, 0x47, 0x5b, 0x7b,
	0x71, 0x2c, 0x6c, 0x6f, 0xf0, 0xae, 0x7b, 0xe7, 0x87, 0xf8, 0x11, 0xfe, 0x84, 0xaf, 0x40, 0x3b,
	0xbb, 0xeb, 0xb8, 0xad, 0xd4, 0x1b, 0xe2, 0x94, 0xbc, 0x37, 0x6f, 0x66, 0x77, 0x9e, 0x67, 0x07,
	0x9e, 0xc4, 0xa2, 0x28, 0x44, 0x79, 0xbc, 0xaa, 0x84, 0x12, 0xc7, 0x4b, 0x9e, 0x15, 0x09, 0xcb,
	0xf3, 0x19, 0x42, 0xe2, 0xe1, 0x4f, 0x78, 0x06, 0xe3, 0x88, 0xa7, 0x99, 0x54, 0x11, 0xff, 0x51,
	0x73, 0xa9, 0x48, 0x08, 0x1e, 0x4b, 0x79, 0xa9, 0x68, 0x67, 0xda, 0x39, 0x0a, 0x4e, 0x47, 0x46,
	0x3e, 0x9b, 0x6b, 0x2e, 0x32, 0xa1, 0xf0, 0x15, 0x6c, 0xba, 0x24, 0xb9, 0x12, 0xa5, 0xe4, 0xe4,
	0x31, 0x0c, 0x31, 0xb4, 0xc8, 0x12, 0x4c, 0xf4, 0xa3, 0x01, 0xe2, 0x4f, 0x49, 0xf8, 0x1a, 0xb6,
	0xbe, 0x96, 0x37, 0xcf, 0xb8, 0x47, 0x4d, 0x60, 0xb2, 0x56, 0x9b, 0xe2, 0xe1, 0x4b, 0x80, 0x4b,
	0x91, 0xba, 0xe4, 0x03, 0xe8, 0xe5, 0x22, 0xb5, 0xd7, 0x03, 0x7b, 0x3d, 0x1d, 0xd7, 0x74, 0x38,
	0x86, 0x00, 0xb5, 0x36, 0xf5, 0x3d, 0x4c, 0x2e, 0x38, 0xab, 0xd4, 0x15, 0x67, 0xcd, 0xe9, 0x33,
	0xf0, 0x97, 0x8e, 0xb3, 0x65, 0x26, 0xb6, 0xcc, 0x5a, 0xbb, 0x96, 0x84, 0x67, 0xb0, 0xdd, 0xaa,
	0x61, 0x1b, 0x7e, 0x06, 0x90, 0x33, 0xa9, 0xea, 0x55, 0xc2, 0x14, 0xc7, 0x2a, 0x5e, 0xd4, 0x62,
	0xf4, 0x3d, 0xa2, 0x3a, 0xe7, 0xf6, 0xcc, 0xf0, 0x04, 0x46, 0x06, 0xda, 0xf4, 0x43, 0xf0, 0xaa,
	0x3a, 0xe7, 0x92, 0x76, 0xa6, 0xbd, 0xa3, 0xe0, 0x34, 0xb0, 0xe7, 0xa3, 0xc6, 0x44, 0xc2, 0x9f,
	0x5d, 0xf0, 0xd0, 0x75, 0xb2, 0x09, 0xdd, 0xc6, 0xa8, 0x6e, 0x96, 0x90, 0x7d, 0x18, 0x2e, 0x85,
	0x54, 0x25, 0x2b, 0x38, 0xed, 0x22, 0xdb, 0x60, 0x42, 0x61, 0x70, 0xcd, 0x2b, 0x99, 0x89, 0x92,
	0xf6, 0x8c, 0xb3, 0x16, 0x92, 0x29, 0x04, 0x5a, 0xc5, 0x92, 0xa4, 0xe2, 0x52, 0xd2, 0x0d, 0x8c,
	0xb6, 0x29, 0xdd, 0x53, 0xc2, 0xaf, 0xb3, 0x98, 0x67, 0xe5, 0x37, 0x41, 0x3d, 0x14, 0xb4, 0x18,
	0xb2, 0x0b, 0x7d, 0xa9, 0x98, 0xaa, 0x25, 0xed, 0x63, 0xbf, 0x16, 0x69, 0x7e, 0xc9, 0x59, 0xae,
	0x96, 0x74, 0x60, 0x78, 0x83, 0xc8, 0x53, 0x80, 0xb8, 0xe2, 0x4c, 0xf1, 0x64, 0xc1, 0x14, 0x1d,
	0x62, 0xcc, 0xb7, 0xcc, 0x5c, 0x91, 0x43, 0x18, 0x69, 0xc3, 0xe2, 0x25, 0x8f, 0xbf, 0x6b, 0x81,
	0x8f, 0x82, 0xa0, 0xe1, 0xe6, 0x2a, 0xfc, 0xd3, 0x81, 0xde, 0xa5, 0x48, 0xef, 0x19, 0x98, 0x7f,
	0x66, 0xc6, 0x0e, 0x78, 0x52, 0xb1, 0x4a, 0xa1, 0x0f, 0xbd, 0xc8, 0x00, 0x42, 0x60, 0x43, 0x2a,
	0xb1, 0x42, 0x03, 0x7a, 0x11, 0xfe, 0xd7, 0xca, 0x82, 0xa9, 0xd8, 0x75, 0x6f, 0x80, 0x36, 0x45,
	0x7f, 0xc7, 0x2c, 0xc1, 0xc6, 0xfd, 0xc8, 0x22, 0x7d, 0xa7, 0xca, 0x0c, 0x05, 0x36, 0x3c, 0x8a,
	0x1c, 0x0c, 0x7f, 0xf7, 0xc0, 0x6f, 0x06, 0xed, 0x7f, 0xb4, 0xfc, 0x02, 0xc6, 0xb1, 0xa8, 0x4b,
	0xb5, 0x70, 0x17, 0xf4, 0xb0, 0xa1, 0x11, 0x92, 0xee, 0xf5, 0x3c, 0x87, 0xc0, 0x88, 0x58, 0xce,
	0x2b, 0x65, 0x27, 0x01, 0x90, 0x9a, 0x6b, 0x66, 0x2d, 0x48, 0x44, 0x7d, 0xa5, 0xac, 0x29, 0x46,
	0xf0, 0x51, 0x33, 0xda, 0x99, 0x7a, 0xa5, 0xb2, 0x82, 0xdb, 0x91, 0xb0, 0x48, 0xfb, 0x98, 0x0b,
	0x96, 0x9c, 0xa0, 0x2f, 0x7e, 0x64, 0x80, 0x63, 0xcf, 0x29, 0xac, 0xd9, 0x73, 0x5d, 0x03, 0xc3,
	0xe7, 0x34, 0x30, 0xee, 0x1a, 0xa4, 0xad, 0x29, 0x78, 0x51, 0x4b, 0x96, 0x72, 0x3a, 0xc2, 0xea,
	0x0d, 0xd6, 0xb1, 0x78, 0x55, 0x9b, 0xd8, 0xd8, 0xc4, 0x1c, 0x26, 0x07, 0xe0, 0xa7, 0xa2, 0xaa,
	0x55, 0x56, 0x72, 0x49, 0x37, 0xcd, 0xa4, 0x36, 0x84, 0xfe, 0x16, 0x69, 0xbc, 0xc0, 0x16, 0xe8,
	0x16, 0x06, 0x07, 0x69, 0xfc, 0x41, 0x43, 0xb2, 0x07, 0x83, 0x34, 0x5e, 0x60, 0x37, 0x13, 0xd3,
	0x4d, 0x1a, 0x7f, 0xc9, 0x0a, 0x1e, 0x4e, 0x61, 0x43, 0xbf, 0x66, 0xfd, 0x41, 0x62, 0x51, 0x2a,
	0xb7, 0x51, 0xfd, 0xc8, 0xc1, 0xd3, 0x5f, 0x5d, 0xd8, 0xba, 0xb0, 0x4b, 0xf9, 0x33, 0xaf, 0xf4,
	0x33, 0x23, 0x6f, 0xa0, 0x6f, 0x96, 0x1f, 0xd9, 0x71, 0x2b, 0xa1, 0xbd, 0x39, 0xf7, 0x1f, 0xdd,
	0x62, 0xed, 0x9a, 0x7b, 0x40, 0xde, 0xc2, 0xd0, 0xed, 0x4d, 0xb2, 0x6b, 0x45, 0xb7, 0xd6, 0xee,
	0xfe, 0xde, 0x1d, 0xbe, 0x49, 0x9f, 0x99, 0x77, 0xb6, 0xdd, 0x5a, 0xa7, 0x36, 0x89, 0xb4, 0xa9,
	0x46, 0xff, 0xae, 0x3d, 0xaa, 0x7b, 0x77, 0xb6, 0xa7, 0xcd, 0xa5, 0x77, 0x03, 0x4d, 0x85, 0x13,
	0xeb, 0x0f, 0x69, 0xaf, 0x3e, 0x9b, 0xf7, 0xf0, 0x06, 0xe7, 0x52, 0xae, 0xfa, 0xc8, 0x9e, 0xfd,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x6a, 0x6c, 0xc0, 0xd8, 0x06, 0x00, 0x00,
}