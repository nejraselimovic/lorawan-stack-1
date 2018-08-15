// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: go.thethings.network/lorawan-stack/api/gatewayserver.proto

package ttnpb

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"
import google_protobuf3 "github.com/gogo/protobuf/types"

import context "context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// GatewayUp may contain zero or more uplink messages and/or a status message for the gateway.
type GatewayUp struct {
	// UplinkMessages received by the gateway.
	UplinkMessages []*UplinkMessage `protobuf:"bytes,1,rep,name=uplink_messages,json=uplinkMessages" json:"uplink_messages,omitempty"`
	GatewayStatus  *GatewayStatus   `protobuf:"bytes,2,opt,name=gateway_status,json=gatewayStatus" json:"gateway_status,omitempty"`
}

func (m *GatewayUp) Reset()                    { *m = GatewayUp{} }
func (m *GatewayUp) String() string            { return proto.CompactTextString(m) }
func (*GatewayUp) ProtoMessage()               {}
func (*GatewayUp) Descriptor() ([]byte, []int) { return fileDescriptorGatewayserver, []int{0} }

func (m *GatewayUp) GetUplinkMessages() []*UplinkMessage {
	if m != nil {
		return m.UplinkMessages
	}
	return nil
}

func (m *GatewayUp) GetGatewayStatus() *GatewayStatus {
	if m != nil {
		return m.GatewayStatus
	}
	return nil
}

// GatewayDown contains downlink messages for the gateway.
type GatewayDown struct {
	// DownlinkMessage for the gateway.
	DownlinkMessage *DownlinkMessage `protobuf:"bytes,1,opt,name=downlink_message,json=downlinkMessage" json:"downlink_message,omitempty"`
}

func (m *GatewayDown) Reset()                    { *m = GatewayDown{} }
func (m *GatewayDown) String() string            { return proto.CompactTextString(m) }
func (*GatewayDown) ProtoMessage()               {}
func (*GatewayDown) Descriptor() ([]byte, []int) { return fileDescriptorGatewayserver, []int{1} }

func (m *GatewayDown) GetDownlinkMessage() *DownlinkMessage {
	if m != nil {
		return m.DownlinkMessage
	}
	return nil
}

type GetFrequencyPlanRequest struct {
	// ID of the frequency plan to fetch.
	FrequencyPlanID string `protobuf:"bytes,1,opt,name=frequency_plan_id,json=frequencyPlanId,proto3" json:"frequency_plan_id,omitempty"`
}

func (m *GetFrequencyPlanRequest) Reset()         { *m = GetFrequencyPlanRequest{} }
func (m *GetFrequencyPlanRequest) String() string { return proto.CompactTextString(m) }
func (*GetFrequencyPlanRequest) ProtoMessage()    {}
func (*GetFrequencyPlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorGatewayserver, []int{2}
}

func (m *GetFrequencyPlanRequest) GetFrequencyPlanID() string {
	if m != nil {
		return m.FrequencyPlanID
	}
	return ""
}

func init() {
	proto.RegisterType((*GatewayUp)(nil), "ttn.lorawan.v3.GatewayUp")
	golang_proto.RegisterType((*GatewayUp)(nil), "ttn.lorawan.v3.GatewayUp")
	proto.RegisterType((*GatewayDown)(nil), "ttn.lorawan.v3.GatewayDown")
	golang_proto.RegisterType((*GatewayDown)(nil), "ttn.lorawan.v3.GatewayDown")
	proto.RegisterType((*GetFrequencyPlanRequest)(nil), "ttn.lorawan.v3.GetFrequencyPlanRequest")
	golang_proto.RegisterType((*GetFrequencyPlanRequest)(nil), "ttn.lorawan.v3.GetFrequencyPlanRequest")
}
func (this *GatewayUp) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*GatewayUp)
	if !ok {
		that2, ok := that.(GatewayUp)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *GatewayUp")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *GatewayUp but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *GatewayUp but is not nil && this == nil")
	}
	if len(this.UplinkMessages) != len(that1.UplinkMessages) {
		return fmt.Errorf("UplinkMessages this(%v) Not Equal that(%v)", len(this.UplinkMessages), len(that1.UplinkMessages))
	}
	for i := range this.UplinkMessages {
		if !this.UplinkMessages[i].Equal(that1.UplinkMessages[i]) {
			return fmt.Errorf("UplinkMessages this[%v](%v) Not Equal that[%v](%v)", i, this.UplinkMessages[i], i, that1.UplinkMessages[i])
		}
	}
	if !this.GatewayStatus.Equal(that1.GatewayStatus) {
		return fmt.Errorf("GatewayStatus this(%v) Not Equal that(%v)", this.GatewayStatus, that1.GatewayStatus)
	}
	return nil
}
func (this *GatewayUp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GatewayUp)
	if !ok {
		that2, ok := that.(GatewayUp)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.UplinkMessages) != len(that1.UplinkMessages) {
		return false
	}
	for i := range this.UplinkMessages {
		if !this.UplinkMessages[i].Equal(that1.UplinkMessages[i]) {
			return false
		}
	}
	if !this.GatewayStatus.Equal(that1.GatewayStatus) {
		return false
	}
	return true
}
func (this *GatewayDown) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*GatewayDown)
	if !ok {
		that2, ok := that.(GatewayDown)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *GatewayDown")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *GatewayDown but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *GatewayDown but is not nil && this == nil")
	}
	if !this.DownlinkMessage.Equal(that1.DownlinkMessage) {
		return fmt.Errorf("DownlinkMessage this(%v) Not Equal that(%v)", this.DownlinkMessage, that1.DownlinkMessage)
	}
	return nil
}
func (this *GatewayDown) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GatewayDown)
	if !ok {
		that2, ok := that.(GatewayDown)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.DownlinkMessage.Equal(that1.DownlinkMessage) {
		return false
	}
	return true
}
func (this *GetFrequencyPlanRequest) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*GetFrequencyPlanRequest)
	if !ok {
		that2, ok := that.(GetFrequencyPlanRequest)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *GetFrequencyPlanRequest")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *GetFrequencyPlanRequest but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *GetFrequencyPlanRequest but is not nil && this == nil")
	}
	if this.FrequencyPlanID != that1.FrequencyPlanID {
		return fmt.Errorf("FrequencyPlanID this(%v) Not Equal that(%v)", this.FrequencyPlanID, that1.FrequencyPlanID)
	}
	return nil
}
func (this *GetFrequencyPlanRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetFrequencyPlanRequest)
	if !ok {
		that2, ok := that.(GetFrequencyPlanRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.FrequencyPlanID != that1.FrequencyPlanID {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GtwGs service

type GtwGsClient interface {
	// LinkGateway links the gateway to the Gateway Server. The authentication information
	// will be used to determine the gateway ID. If no authentication information
	// is present, this gateway may not be used for downlink.
	LinkGateway(ctx context.Context, opts ...grpc.CallOption) (GtwGs_LinkGatewayClient, error)
	// GetFrequencyPlan associated to the gateway. The gateway is ID'd by its authentication token.
	GetFrequencyPlan(ctx context.Context, in *GetFrequencyPlanRequest, opts ...grpc.CallOption) (*FrequencyPlan, error)
}

type gtwGsClient struct {
	cc *grpc.ClientConn
}

func NewGtwGsClient(cc *grpc.ClientConn) GtwGsClient {
	return &gtwGsClient{cc}
}

func (c *gtwGsClient) LinkGateway(ctx context.Context, opts ...grpc.CallOption) (GtwGs_LinkGatewayClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GtwGs_serviceDesc.Streams[0], c.cc, "/ttn.lorawan.v3.GtwGs/LinkGateway", opts...)
	if err != nil {
		return nil, err
	}
	x := &gtwGsLinkGatewayClient{stream}
	return x, nil
}

type GtwGs_LinkGatewayClient interface {
	Send(*GatewayUp) error
	Recv() (*GatewayDown, error)
	grpc.ClientStream
}

type gtwGsLinkGatewayClient struct {
	grpc.ClientStream
}

func (x *gtwGsLinkGatewayClient) Send(m *GatewayUp) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gtwGsLinkGatewayClient) Recv() (*GatewayDown, error) {
	m := new(GatewayDown)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gtwGsClient) GetFrequencyPlan(ctx context.Context, in *GetFrequencyPlanRequest, opts ...grpc.CallOption) (*FrequencyPlan, error) {
	out := new(FrequencyPlan)
	err := grpc.Invoke(ctx, "/ttn.lorawan.v3.GtwGs/GetFrequencyPlan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GtwGs service

type GtwGsServer interface {
	// LinkGateway links the gateway to the Gateway Server. The authentication information
	// will be used to determine the gateway ID. If no authentication information
	// is present, this gateway may not be used for downlink.
	LinkGateway(GtwGs_LinkGatewayServer) error
	// GetFrequencyPlan associated to the gateway. The gateway is ID'd by its authentication token.
	GetFrequencyPlan(context.Context, *GetFrequencyPlanRequest) (*FrequencyPlan, error)
}

func RegisterGtwGsServer(s *grpc.Server, srv GtwGsServer) {
	s.RegisterService(&_GtwGs_serviceDesc, srv)
}

func _GtwGs_LinkGateway_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GtwGsServer).LinkGateway(&gtwGsLinkGatewayServer{stream})
}

type GtwGs_LinkGatewayServer interface {
	Send(*GatewayDown) error
	Recv() (*GatewayUp, error)
	grpc.ServerStream
}

type gtwGsLinkGatewayServer struct {
	grpc.ServerStream
}

func (x *gtwGsLinkGatewayServer) Send(m *GatewayDown) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gtwGsLinkGatewayServer) Recv() (*GatewayUp, error) {
	m := new(GatewayUp)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GtwGs_GetFrequencyPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFrequencyPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GtwGsServer).GetFrequencyPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GtwGs/GetFrequencyPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GtwGsServer).GetFrequencyPlan(ctx, req.(*GetFrequencyPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GtwGs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.GtwGs",
	HandlerType: (*GtwGsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFrequencyPlan",
			Handler:    _GtwGs_GetFrequencyPlan_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LinkGateway",
			Handler:       _GtwGs_LinkGateway_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "go.thethings.network/lorawan-stack/api/gatewayserver.proto",
}

// Client API for NsGs service

type NsGsClient interface {
	// ScheduleDownlink instructs the Gateway Server to schedule a downlink message.
	// The Gateway Server may refuse if there are any conflicts in the schedule or
	// if a duty cycle prevents the gateway from transmitting.
	ScheduleDownlink(ctx context.Context, in *DownlinkMessage, opts ...grpc.CallOption) (*google_protobuf3.Empty, error)
}

type nsGsClient struct {
	cc *grpc.ClientConn
}

func NewNsGsClient(cc *grpc.ClientConn) NsGsClient {
	return &nsGsClient{cc}
}

func (c *nsGsClient) ScheduleDownlink(ctx context.Context, in *DownlinkMessage, opts ...grpc.CallOption) (*google_protobuf3.Empty, error) {
	out := new(google_protobuf3.Empty)
	err := grpc.Invoke(ctx, "/ttn.lorawan.v3.NsGs/ScheduleDownlink", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NsGs service

type NsGsServer interface {
	// ScheduleDownlink instructs the Gateway Server to schedule a downlink message.
	// The Gateway Server may refuse if there are any conflicts in the schedule or
	// if a duty cycle prevents the gateway from transmitting.
	ScheduleDownlink(context.Context, *DownlinkMessage) (*google_protobuf3.Empty, error)
}

func RegisterNsGsServer(s *grpc.Server, srv NsGsServer) {
	s.RegisterService(&_NsGs_serviceDesc, srv)
}

func _NsGs_ScheduleDownlink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownlinkMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsGsServer).ScheduleDownlink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.NsGs/ScheduleDownlink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsGsServer).ScheduleDownlink(ctx, req.(*DownlinkMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _NsGs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.NsGs",
	HandlerType: (*NsGsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScheduleDownlink",
			Handler:    _NsGs_ScheduleDownlink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.thethings.network/lorawan-stack/api/gatewayserver.proto",
}

// Client API for Gs service

type GsClient interface {
	// Get the connection stats of a gateway
	GetGatewayConnectionStats(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*GatewayConnectionStats, error)
}

type gsClient struct {
	cc *grpc.ClientConn
}

func NewGsClient(cc *grpc.ClientConn) GsClient {
	return &gsClient{cc}
}

func (c *gsClient) GetGatewayConnectionStats(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*GatewayConnectionStats, error) {
	out := new(GatewayConnectionStats)
	err := grpc.Invoke(ctx, "/ttn.lorawan.v3.Gs/GetGatewayConnectionStats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Gs service

type GsServer interface {
	// Get the connection stats of a gateway
	GetGatewayConnectionStats(context.Context, *GatewayIdentifiers) (*GatewayConnectionStats, error)
}

func RegisterGsServer(s *grpc.Server, srv GsServer) {
	s.RegisterService(&_Gs_serviceDesc, srv)
}

func _Gs_GetGatewayConnectionStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GsServer).GetGatewayConnectionStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.Gs/GetGatewayConnectionStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GsServer).GetGatewayConnectionStats(ctx, req.(*GatewayIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.Gs",
	HandlerType: (*GsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGatewayConnectionStats",
			Handler:    _Gs_GetGatewayConnectionStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.thethings.network/lorawan-stack/api/gatewayserver.proto",
}

func (m *GatewayUp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GatewayUp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UplinkMessages) > 0 {
		for _, msg := range m.UplinkMessages {
			dAtA[i] = 0xa
			i++
			i = encodeVarintGatewayserver(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.GatewayStatus != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGatewayserver(dAtA, i, uint64(m.GatewayStatus.Size()))
		n1, err := m.GatewayStatus.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *GatewayDown) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GatewayDown) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DownlinkMessage != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGatewayserver(dAtA, i, uint64(m.DownlinkMessage.Size()))
		n2, err := m.DownlinkMessage.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *GetFrequencyPlanRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetFrequencyPlanRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FrequencyPlanID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGatewayserver(dAtA, i, uint64(len(m.FrequencyPlanID)))
		i += copy(dAtA[i:], m.FrequencyPlanID)
	}
	return i, nil
}

func encodeVarintGatewayserver(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedGatewayUp(r randyGatewayserver, easy bool) *GatewayUp {
	this := &GatewayUp{}
	if r.Intn(10) != 0 {
		v1 := r.Intn(5)
		this.UplinkMessages = make([]*UplinkMessage, v1)
		for i := 0; i < v1; i++ {
			this.UplinkMessages[i] = NewPopulatedUplinkMessage(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		this.GatewayStatus = NewPopulatedGatewayStatus(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedGatewayDown(r randyGatewayserver, easy bool) *GatewayDown {
	this := &GatewayDown{}
	if r.Intn(10) != 0 {
		this.DownlinkMessage = NewPopulatedDownlinkMessage(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedGetFrequencyPlanRequest(r randyGatewayserver, easy bool) *GetFrequencyPlanRequest {
	this := &GetFrequencyPlanRequest{}
	this.FrequencyPlanID = randStringGatewayserver(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyGatewayserver interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneGatewayserver(r randyGatewayserver) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringGatewayserver(r randyGatewayserver) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneGatewayserver(r)
	}
	return string(tmps)
}
func randUnrecognizedGatewayserver(r randyGatewayserver, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldGatewayserver(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldGatewayserver(dAtA []byte, r randyGatewayserver, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateGatewayserver(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateGatewayserver(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateGatewayserver(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateGatewayserver(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateGatewayserver(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateGatewayserver(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateGatewayserver(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *GatewayUp) Size() (n int) {
	var l int
	_ = l
	if len(m.UplinkMessages) > 0 {
		for _, e := range m.UplinkMessages {
			l = e.Size()
			n += 1 + l + sovGatewayserver(uint64(l))
		}
	}
	if m.GatewayStatus != nil {
		l = m.GatewayStatus.Size()
		n += 1 + l + sovGatewayserver(uint64(l))
	}
	return n
}

func (m *GatewayDown) Size() (n int) {
	var l int
	_ = l
	if m.DownlinkMessage != nil {
		l = m.DownlinkMessage.Size()
		n += 1 + l + sovGatewayserver(uint64(l))
	}
	return n
}

func (m *GetFrequencyPlanRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.FrequencyPlanID)
	if l > 0 {
		n += 1 + l + sovGatewayserver(uint64(l))
	}
	return n
}

func sovGatewayserver(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGatewayserver(x uint64) (n int) {
	return sovGatewayserver((x << 1) ^ uint64((int64(x) >> 63)))
}
func (m *GatewayUp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGatewayserver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GatewayUp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GatewayUp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UplinkMessages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatewayserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGatewayserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UplinkMessages = append(m.UplinkMessages, &UplinkMessage{})
			if err := m.UplinkMessages[len(m.UplinkMessages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GatewayStatus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatewayserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGatewayserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GatewayStatus == nil {
				m.GatewayStatus = &GatewayStatus{}
			}
			if err := m.GatewayStatus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGatewayserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGatewayserver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GatewayDown) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGatewayserver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GatewayDown: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GatewayDown: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DownlinkMessage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatewayserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGatewayserver
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DownlinkMessage == nil {
				m.DownlinkMessage = &DownlinkMessage{}
			}
			if err := m.DownlinkMessage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGatewayserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGatewayserver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetFrequencyPlanRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGatewayserver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetFrequencyPlanRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetFrequencyPlanRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FrequencyPlanID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatewayserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGatewayserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FrequencyPlanID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGatewayserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGatewayserver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGatewayserver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGatewayserver
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGatewayserver
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGatewayserver
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGatewayserver
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGatewayserver
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGatewayserver(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGatewayserver = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGatewayserver   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("go.thethings.network/lorawan-stack/api/gatewayserver.proto", fileDescriptorGatewayserver)
}
func init() {
	golang_proto.RegisterFile("go.thethings.network/lorawan-stack/api/gatewayserver.proto", fileDescriptorGatewayserver)
}

var fileDescriptorGatewayserver = []byte{
	// 626 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x3d, 0x4c, 0xdb, 0x40,
	0x14, 0xf6, 0xd1, 0x1f, 0x89, 0x43, 0x25, 0xd4, 0x95, 0x5a, 0x08, 0xea, 0x03, 0x65, 0x68, 0x19,
	0x8a, 0x53, 0x85, 0x56, 0xaa, 0xba, 0x54, 0xa2, 0x29, 0x16, 0xfd, 0x53, 0x15, 0xc4, 0x50, 0x96,
	0xe8, 0x12, 0x5f, 0x2e, 0x56, 0x9c, 0x3b, 0xd7, 0x77, 0x26, 0x62, 0x63, 0x64, 0xec, 0xd8, 0x6e,
	0x55, 0x27, 0xa6, 0x8a, 0x91, 0x91, 0x91, 0x91, 0x91, 0xa9, 0xc2, 0xe7, 0x85, 0x91, 0x91, 0xb1,
	0x8a, 0x63, 0x0b, 0x12, 0x88, 0x1a, 0xb6, 0x7b, 0xef, 0xbe, 0xef, 0xbb, 0xf7, 0xbd, 0xf7, 0x0e,
	0xbf, 0x66, 0xc2, 0x52, 0x4d, 0xaa, 0x9a, 0x2e, 0x67, 0xd2, 0xe2, 0x54, 0x75, 0x44, 0xd0, 0x2a,
	0x7a, 0x22, 0x20, 0x1d, 0xc2, 0x17, 0xa5, 0x22, 0xf5, 0x56, 0x91, 0xf8, 0x6e, 0x91, 0x11, 0x45,
	0x3b, 0x64, 0x4b, 0xd2, 0x60, 0x93, 0x06, 0x96, 0x1f, 0x08, 0x25, 0xcc, 0x49, 0xa5, 0xb8, 0x95,
	0x42, 0xad, 0xcd, 0xa5, 0xfc, 0x8b, 0x9b, 0x69, 0xf5, 0x54, 0xf2, 0xaf, 0x46, 0x64, 0xb9, 0x0e,
	0xe5, 0xca, 0x6d, 0xb8, 0x34, 0x90, 0x29, 0x73, 0xd4, 0xf7, 0xb2, 0x12, 0x7b, 0xac, 0x97, 0x23,
	0xb2, 0xda, 0x54, 0x4a, 0xc2, 0xa8, 0xbc, 0x31, 0x4d, 0x11, 0x87, 0x28, 0x72, 0x43, 0x5a, 0x40,
	0x99, 0x2b, 0x38, 0xf1, 0x52, 0xda, 0x22, 0x73, 0x55, 0x33, 0xac, 0x59, 0x75, 0xd1, 0x2e, 0x32,
	0xc1, 0x44, 0x31, 0x49, 0xd7, 0xc2, 0x46, 0x12, 0x25, 0x41, 0x72, 0x4a, 0xe1, 0xc0, 0x84, 0x60,
	0x1e, 0xbd, 0x40, 0x39, 0x61, 0x40, 0x94, 0x2b, 0x32, 0xcf, 0xb3, 0x83, 0xf7, 0xb4, 0xed, 0xab,
	0x74, 0x00, 0x85, 0x9f, 0x08, 0x8f, 0xdb, 0xbd, 0x91, 0xac, 0xfb, 0xe6, 0x0a, 0xce, 0x85, 0xbe,
	0xe7, 0xf2, 0x56, 0x35, 0x6b, 0xc0, 0x34, 0x9a, 0xbf, 0xb5, 0x30, 0x51, 0x7a, 0x6c, 0xf5, 0x8f,
	0xdb, 0x5a, 0x4f, 0x60, 0x9f, 0x7a, 0xa8, 0xca, 0x64, 0x78, 0x39, 0x94, 0x66, 0x19, 0x4f, 0xa6,
	0x73, 0xae, 0x4a, 0x45, 0x54, 0x28, 0xa7, 0xc7, 0xe6, 0xd1, 0x75, 0x32, 0xe9, 0xd3, 0x6b, 0x09,
	0xa8, 0x72, 0x8f, 0x5d, 0x0e, 0x0b, 0x5f, 0xf1, 0x44, 0x7a, 0x5f, 0x16, 0x1d, 0x6e, 0xbe, 0xc7,
	0x53, 0x8e, 0xe8, 0xf0, 0xcb, 0xe5, 0x4d, 0xa3, 0x44, 0x76, 0x6e, 0x50, 0xb6, 0x9c, 0xe2, 0xb2,
	0xfa, 0x72, 0x4e, 0x7f, 0xa2, 0xb0, 0x81, 0x1f, 0xd9, 0x54, 0xad, 0x04, 0xf4, 0x5b, 0x48, 0x79,
	0x7d, 0xeb, 0x8b, 0x47, 0x78, 0xa5, 0x7b, 0x96, 0xca, 0x7c, 0x83, 0xef, 0x37, 0xb2, 0x7c, 0xd5,
	0xf7, 0x08, 0xaf, 0xba, 0x4e, 0xf2, 0xce, 0xf8, 0xf2, 0x03, 0xfd, 0x77, 0x2e, 0xd7, 0x47, 0x5a,
	0x2d, 0x57, 0x72, 0x8d, 0xbe, 0x84, 0x53, 0xfa, 0x83, 0xf0, 0x1d, 0x5b, 0x75, 0x6c, 0x69, 0xae,
	0xe2, 0x89, 0x8f, 0x2e, 0x6f, 0xa5, 0x26, 0xcc, 0x99, 0x21, 0xee, 0xd7, 0xfd, 0xfc, 0xec, 0x90,
	0xab, 0xae, 0x91, 0x05, 0xf4, 0x1c, 0x99, 0x1b, 0x78, 0x6a, 0xb0, 0x60, 0xf3, 0xe9, 0x15, 0xd2,
	0xf5, 0x96, 0xf2, 0x57, 0xda, 0xde, 0x87, 0x2a, 0xad, 0xe1, 0xdb, 0x9f, 0xa5, 0x2d, 0xcd, 0x0f,
	0x78, 0x6a, 0xad, 0xde, 0xa4, 0x4e, 0xe8, 0xd1, 0xac, 0x81, 0xe6, 0xff, 0x5a, 0x9b, 0x7f, 0x68,
	0xf5, 0xd6, 0xcb, 0xca, 0xd6, 0xcb, 0x7a, 0xd7, 0x5d, 0xaf, 0x52, 0x1b, 0x8f, 0xd9, 0xd2, 0x64,
	0x78, 0xc6, 0xa6, 0x2a, 0x35, 0xf3, 0x56, 0x70, 0x4e, 0xeb, 0xdd, 0xcd, 0xec, 0x0e, 0x58, 0x9a,
	0x85, 0x21, 0xa6, 0x57, 0x2f, 0x3e, 0x7b, 0xfe, 0xc9, 0x10, 0xcc, 0x80, 0xd6, 0xf2, 0x6f, 0x74,
	0x18, 0x01, 0x3a, 0x8a, 0x00, 0x1d, 0x47, 0x80, 0x4e, 0x22, 0x40, 0xa7, 0x11, 0x18, 0x67, 0x11,
	0x18, 0xe7, 0x11, 0xa0, 0x6d, 0x0d, 0xc6, 0x8e, 0x06, 0x63, 0x57, 0x03, 0xda, 0xd3, 0x60, 0xec,
	0x6b, 0x40, 0x07, 0x1a, 0xd0, 0xa1, 0x06, 0x74, 0xa4, 0x01, 0x1d, 0x6b, 0x30, 0x4e, 0x34, 0xa0,
	0x53, 0x0d, 0xc6, 0x99, 0x06, 0x74, 0xae, 0xc1, 0xd8, 0x8e, 0xc1, 0xd8, 0x89, 0x01, 0x7d, 0x8f,
	0xc1, 0xf8, 0x11, 0x03, 0xfa, 0x15, 0x83, 0xb1, 0x1b, 0x83, 0xb1, 0x17, 0x03, 0xda, 0x8f, 0x01,
	0x1d, 0xc4, 0x80, 0x36, 0x9e, 0x8d, 0xf0, 0xc3, 0xfd, 0x16, 0x2b, 0x2a, 0xc5, 0xfd, 0x5a, 0xed,
	0x6e, 0xd2, 0xa3, 0xa5, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x7f, 0x84, 0xc3, 0x78, 0x05,
	0x00, 0x00,
}
