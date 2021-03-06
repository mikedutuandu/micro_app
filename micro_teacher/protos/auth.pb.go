// Code generated by protoc-gen-go. DO NOT EDIT.
// source: micro_teacher/protos/auth.proto

package micro_teacher_pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Company              string   `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52e1fdbb028fae4, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52e1fdbb028fae4, []int{1}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type CreateUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52e1fdbb028fae4, []int{2}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateUserResponse struct {
	Token                *Token   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52e1fdbb028fae4, []int{3}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type AuthRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52e1fdbb028fae4, []int{4}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthResponse struct {
	Token                *Token   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52e1fdbb028fae4, []int{5}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type ValidateTokenRequest struct {
	Token                *Token   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateTokenRequest) Reset()         { *m = ValidateTokenRequest{} }
func (m *ValidateTokenRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateTokenRequest) ProtoMessage()    {}
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52e1fdbb028fae4, []int{6}
}

func (m *ValidateTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateTokenRequest.Unmarshal(m, b)
}
func (m *ValidateTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateTokenRequest.Marshal(b, m, deterministic)
}
func (m *ValidateTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateTokenRequest.Merge(m, src)
}
func (m *ValidateTokenRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateTokenRequest.Size(m)
}
func (m *ValidateTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateTokenRequest proto.InternalMessageInfo

func (m *ValidateTokenRequest) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type ValidateTokenResponse struct {
	Token                *Token   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateTokenResponse) Reset()         { *m = ValidateTokenResponse{} }
func (m *ValidateTokenResponse) String() string { return proto.CompactTextString(m) }
func (*ValidateTokenResponse) ProtoMessage()    {}
func (*ValidateTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52e1fdbb028fae4, []int{7}
}

func (m *ValidateTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateTokenResponse.Unmarshal(m, b)
}
func (m *ValidateTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateTokenResponse.Marshal(b, m, deterministic)
}
func (m *ValidateTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateTokenResponse.Merge(m, src)
}
func (m *ValidateTokenResponse) XXX_Size() int {
	return xxx_messageInfo_ValidateTokenResponse.Size(m)
}
func (m *ValidateTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateTokenResponse proto.InternalMessageInfo

func (m *ValidateTokenResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "micro_teacher.User")
	proto.RegisterType((*Token)(nil), "micro_teacher.Token")
	proto.RegisterType((*CreateUserRequest)(nil), "micro_teacher.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "micro_teacher.CreateUserResponse")
	proto.RegisterType((*AuthRequest)(nil), "micro_teacher.AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "micro_teacher.AuthResponse")
	proto.RegisterType((*ValidateTokenRequest)(nil), "micro_teacher.ValidateTokenRequest")
	proto.RegisterType((*ValidateTokenResponse)(nil), "micro_teacher.ValidateTokenResponse")
}

func init() { proto.RegisterFile("micro_teacher/protos/auth.proto", fileDescriptor_b52e1fdbb028fae4) }

var fileDescriptor_b52e1fdbb028fae4 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x5d, 0x4f, 0xea, 0x40,
	0x10, 0x4d, 0x7b, 0xcb, 0xbd, 0xdc, 0x41, 0x8c, 0x8e, 0x98, 0x34, 0xf5, 0x41, 0xac, 0x26, 0x1a,
	0x1e, 0x20, 0x81, 0x37, 0x63, 0x82, 0xc2, 0x2f, 0x10, 0x3f, 0x1e, 0x78, 0x21, 0x4b, 0x3b, 0x09,
	0x8d, 0xf4, 0xc3, 0xdd, 0x2d, 0xc6, 0x1f, 0x6f, 0x62, 0xba, 0x2d, 0x48, 0x57, 0x84, 0xc4, 0xb7,
	0x99, 0xd9, 0x33, 0xe7, 0x9c, 0x39, 0x4d, 0xe1, 0x34, 0x0c, 0x3c, 0x1e, 0x4f, 0x24, 0x31, 0x6f,
	0x46, 0xbc, 0x93, 0xf0, 0x58, 0xc6, 0xa2, 0xc3, 0x52, 0x39, 0x6b, 0xab, 0x1a, 0xeb, 0x25, 0x80,
	0xbb, 0x00, 0xeb, 0x49, 0x10, 0xc7, 0x7d, 0x30, 0x03, 0xdf, 0x36, 0x9a, 0xc6, 0xd5, 0xff, 0x91,
	0x19, 0xf8, 0x88, 0x60, 0x45, 0x2c, 0x24, 0xdb, 0x54, 0x13, 0x55, 0xa3, 0x0d, 0xff, 0xbc, 0x38,
	0x4c, 0x58, 0xf4, 0x6e, 0xff, 0x51, 0xe3, 0x65, 0x8b, 0x0d, 0xa8, 0x50, 0xc8, 0x82, 0xb9, 0x6d,
	0xa9, 0x79, 0xde, 0xa0, 0x03, 0xd5, 0x84, 0x09, 0xf1, 0x16, 0x73, 0xdf, 0xae, 0xa8, 0x87, 0x55,
	0xef, 0xf6, 0xa0, 0xf2, 0x18, 0xbf, 0x50, 0x94, 0xad, 0xca, 0xac, 0x28, 0xb4, 0xf3, 0x26, 0x9b,
	0x2e, 0xd8, 0x3c, 0xf0, 0x95, 0x7e, 0x75, 0x94, 0x37, 0xee, 0x0d, 0x1c, 0x0e, 0x39, 0x31, 0x49,
	0x99, 0xe5, 0x11, 0xbd, 0xa6, 0x24, 0x24, 0x5e, 0x82, 0x95, 0x0a, 0xe2, 0x6a, 0xbf, 0xd6, 0x3d,
	0x6a, 0x97, 0xee, 0x6b, 0x2b, 0xa4, 0x02, 0xb8, 0xb7, 0x80, 0xeb, 0xdb, 0x22, 0x89, 0x23, 0x41,
	0xd8, 0x5a, 0xea, 0x9b, 0x6a, 0xbf, 0xa1, 0xed, 0x2b, 0x93, 0x85, 0x2b, 0xb7, 0x0f, 0xb5, 0xbb,
	0x54, 0xce, 0x96, 0xca, 0xab, 0xab, 0x8d, 0x9f, 0xae, 0x36, 0xb5, 0xab, 0xaf, 0x61, 0x2f, 0x27,
	0xf8, 0x85, 0xf8, 0x00, 0x1a, 0xcf, 0x59, 0x0a, 0x4c, 0x52, 0x3e, 0x2f, 0x5c, 0xb4, 0xd6, 0x03,
	0xdc, 0xc1, 0x31, 0x84, 0x63, 0x8d, 0x43, 0x37, 0xb2, 0x9b, 0xa4, 0xfb, 0x61, 0xe4, 0x31, 0x3c,
	0x10, 0x5f, 0x04, 0x1e, 0xe1, 0x3d, 0xc0, 0x57, 0xae, 0xd8, 0xd4, 0x56, 0xbf, 0x7d, 0x30, 0xe7,
	0x6c, 0x0b, 0xa2, 0xb0, 0xd3, 0x07, 0x2b, 0x53, 0x40, 0x47, 0x83, 0xae, 0xa5, 0xef, 0x9c, 0x6c,
	0x7c, 0x2b, 0x08, 0xc6, 0x50, 0x2f, 0x1d, 0x8a, 0xe7, 0x1a, 0x7a, 0x53, 0x94, 0xce, 0xc5, 0x76,
	0x50, 0xce, 0x3d, 0xc0, 0xf1, 0x41, 0x09, 0x36, 0x49, 0xa6, 0xd3, 0xbf, 0xea, 0xe7, 0xea, 0x7d,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x1f, 0x55, 0xb6, 0x7f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/micro_teacher.AuthService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/micro_teacher.AuthService/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error) {
	out := new(ValidateTokenResponse)
	err := c.cc.Invoke(ctx, "/micro_teacher.AuthService/ValidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	Auth(context.Context, *AuthRequest) (*AuthResponse, error)
	ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedAuthServiceServer) Auth(ctx context.Context, req *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (*UnimplementedAuthServiceServer) ValidateToken(ctx context.Context, req *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/micro_teacher.AuthService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/micro_teacher.AuthService/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/micro_teacher.AuthService/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ValidateToken(ctx, req.(*ValidateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "micro_teacher.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _AuthService_CreateUser_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _AuthService_Auth_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _AuthService_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "micro_teacher/protos/auth.proto",
}
