// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MTAMessengerProto.proto

/*
Package MTAMessenger is a generated protocol buffer package.

It is generated from these files:
	MTAMessengerProto.proto

It has these top-level messages:
	Empty
	UserLogin
	User
	Response
	Request
	Message
	Conversation
	AllConversation
	WaittingMessage
	AllMessages
	AllInfoUser
	ConversationDetail
*/
package MTAMessenger

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// login
type UserLogin struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *UserLogin) Reset()                    { *m = UserLogin{} }
func (m *UserLogin) String() string            { return proto.CompactTextString(m) }
func (*UserLogin) ProtoMessage()               {}
func (*UserLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserLogin) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserLogin) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type User struct {
	Uid         uint64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Username    string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password    string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Phone       string `protobuf:"bytes,4,opt,name=phone" json:"phone,omitempty"`
	Email       string `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	CreatedTime string `protobuf:"bytes,6,opt,name=createdTime" json:"createdTime,omitempty"`
	Active      bool   `protobuf:"varint,7,opt,name=active" json:"active,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *User) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetCreatedTime() string {
	if m != nil {
		return m.CreatedTime
	}
	return ""
}

func (m *User) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

//
type Response struct {
	Response string `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	Check    bool   `protobuf:"varint,2,opt,name=check" json:"check,omitempty"`
	Id       string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func (m *Response) GetCheck() bool {
	if m != nil {
		return m.Check
	}
	return false
}

func (m *Response) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
type Request struct {
	Sessionkey string `protobuf:"bytes,1,opt,name=sessionkey" json:"sessionkey,omitempty"`
	Request    string `protobuf:"bytes,2,opt,name=request" json:"request,omitempty"`
	Id         string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Request) GetSessionkey() string {
	if m != nil {
		return m.Sessionkey
	}
	return ""
}

func (m *Request) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//
type Message struct {
	Mid         uint64 `protobuf:"varint,1,opt,name=mid" json:"mid,omitempty"`
	Sessionkey  string `protobuf:"bytes,2,opt,name=sessionkey" json:"sessionkey,omitempty"`
	Cid         string `protobuf:"bytes,3,opt,name=cid" json:"cid,omitempty"`
	Content     string `protobuf:"bytes,5,opt,name=content" json:"content,omitempty"`
	CreatedTime string `protobuf:"bytes,4,opt,name=createdTime" json:"createdTime,omitempty"`
	Check       bool   `protobuf:"varint,6,opt,name=check" json:"check,omitempty"`
	FromName    string `protobuf:"bytes,7,opt,name=from_name,json=fromName" json:"from_name,omitempty"`
	ToUid       string `protobuf:"bytes,8,opt,name=toUid" json:"toUid,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Message) GetMid() uint64 {
	if m != nil {
		return m.Mid
	}
	return 0
}

func (m *Message) GetSessionkey() string {
	if m != nil {
		return m.Sessionkey
	}
	return ""
}

func (m *Message) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *Message) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Message) GetCreatedTime() string {
	if m != nil {
		return m.CreatedTime
	}
	return ""
}

func (m *Message) GetCheck() bool {
	if m != nil {
		return m.Check
	}
	return false
}

func (m *Message) GetFromName() string {
	if m != nil {
		return m.FromName
	}
	return ""
}

func (m *Message) GetToUid() string {
	if m != nil {
		return m.ToUid
	}
	return ""
}

// tra ve cho client khi client chua co
type Conversation struct {
	Cid              string `protobuf:"bytes,1,opt,name=cid" json:"cid,omitempty"`
	ConversationName string `protobuf:"bytes,2,opt,name=ConversationName" json:"ConversationName,omitempty"`
	LastedTime       string `protobuf:"bytes,4,opt,name=LastedTime" json:"LastedTime,omitempty"`
	Sessionkey       string `protobuf:"bytes,5,opt,name=sessionkey" json:"sessionkey,omitempty"`
}

func (m *Conversation) Reset()                    { *m = Conversation{} }
func (m *Conversation) String() string            { return proto.CompactTextString(m) }
func (*Conversation) ProtoMessage()               {}
func (*Conversation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Conversation) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *Conversation) GetConversationName() string {
	if m != nil {
		return m.ConversationName
	}
	return ""
}

func (m *Conversation) GetLastedTime() string {
	if m != nil {
		return m.LastedTime
	}
	return ""
}

func (m *Conversation) GetSessionkey() string {
	if m != nil {
		return m.Sessionkey
	}
	return ""
}

type AllConversation struct {
	ListConversation []*Conversation `protobuf:"bytes,1,rep,name=listConversation" json:"listConversation,omitempty"`
}

func (m *AllConversation) Reset()                    { *m = AllConversation{} }
func (m *AllConversation) String() string            { return proto.CompactTextString(m) }
func (*AllConversation) ProtoMessage()               {}
func (*AllConversation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AllConversation) GetListConversation() []*Conversation {
	if m != nil {
		return m.ListConversation
	}
	return nil
}

// tin nhan chua duoc nhan
type WaittingMessage struct {
	Waittingmess []*Message `protobuf:"bytes,1,rep,name=waittingmess" json:"waittingmess,omitempty"`
}

func (m *WaittingMessage) Reset()                    { *m = WaittingMessage{} }
func (m *WaittingMessage) String() string            { return proto.CompactTextString(m) }
func (*WaittingMessage) ProtoMessage()               {}
func (*WaittingMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *WaittingMessage) GetWaittingmess() []*Message {
	if m != nil {
		return m.Waittingmess
	}
	return nil
}

// tat ca tin nhan
type AllMessages struct {
	Allmess []*Message `protobuf:"bytes,1,rep,name=allmess" json:"allmess,omitempty"`
}

func (m *AllMessages) Reset()                    { *m = AllMessages{} }
func (m *AllMessages) String() string            { return proto.CompactTextString(m) }
func (*AllMessages) ProtoMessage()               {}
func (*AllMessages) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AllMessages) GetAllmess() []*Message {
	if m != nil {
		return m.Allmess
	}
	return nil
}

// tra ve tat ca User theo dieu kien
type AllInfoUser struct {
	Alluser []*User `protobuf:"bytes,1,rep,name=alluser" json:"alluser,omitempty"`
}

func (m *AllInfoUser) Reset()                    { *m = AllInfoUser{} }
func (m *AllInfoUser) String() string            { return proto.CompactTextString(m) }
func (*AllInfoUser) ProtoMessage()               {}
func (*AllInfoUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *AllInfoUser) GetAlluser() []*User {
	if m != nil {
		return m.Alluser
	}
	return nil
}

type ConversationDetail struct {
	Sessionkey string   `protobuf:"bytes,3,opt,name=sessionkey" json:"sessionkey,omitempty"`
	Cid        string   `protobuf:"bytes,1,opt,name=cid" json:"cid,omitempty"`
	Uid        []string `protobuf:"bytes,2,rep,name=uid" json:"uid,omitempty"`
}

func (m *ConversationDetail) Reset()                    { *m = ConversationDetail{} }
func (m *ConversationDetail) String() string            { return proto.CompactTextString(m) }
func (*ConversationDetail) ProtoMessage()               {}
func (*ConversationDetail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ConversationDetail) GetSessionkey() string {
	if m != nil {
		return m.Sessionkey
	}
	return ""
}

func (m *ConversationDetail) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *ConversationDetail) GetUid() []string {
	if m != nil {
		return m.Uid
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "MTAMessenger.Empty")
	proto.RegisterType((*UserLogin)(nil), "MTAMessenger.UserLogin")
	proto.RegisterType((*User)(nil), "MTAMessenger.User")
	proto.RegisterType((*Response)(nil), "MTAMessenger.Response")
	proto.RegisterType((*Request)(nil), "MTAMessenger.Request")
	proto.RegisterType((*Message)(nil), "MTAMessenger.Message")
	proto.RegisterType((*Conversation)(nil), "MTAMessenger.Conversation")
	proto.RegisterType((*AllConversation)(nil), "MTAMessenger.AllConversation")
	proto.RegisterType((*WaittingMessage)(nil), "MTAMessenger.WaittingMessage")
	proto.RegisterType((*AllMessages)(nil), "MTAMessenger.AllMessages")
	proto.RegisterType((*AllInfoUser)(nil), "MTAMessenger.AllInfoUser")
	proto.RegisterType((*ConversationDetail)(nil), "MTAMessenger.ConversationDetail")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ChatgRPC service

type ChatgRPCClient interface {
	// rpc LoadMess(Request) returns(WaittingMessage){}
	// dieu huong chat
	RouteChat(ctx context.Context, opts ...grpc.CallOption) (ChatgRPC_RouteChatClient, error)
	// dang ki thanh vien
	Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	// dang nhap
	Login(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*Response, error)
	// logout
	Logout(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// kiem tra co user tren database hay ko, truyen vao la 1 sessionkey
	CheckUser(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// load tin nhan chua duoc nhan, truyen vao la 1 sessionkey
	LoadWaittingMess(ctx context.Context, in *Request, opts ...grpc.CallOption) (*WaittingMessage, error)
	// return uid
	GetId(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// tra ve cid
	CreateConversation(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// dau vao la uid tra ve danh sach Conversation
	GetAllConversation(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AllConversation, error)
	// add n uid to conversation
	AddUidToConversation(ctx context.Context, in *ConversationDetail, opts ...grpc.CallOption) (*Response, error)
	// load tat ca tin nhan
	LoadAllMessOnCid(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AllMessages, error)
	// ket ban
	AddFriend(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// lay danh sach tat ca user
	GetListUser(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AllInfoUser, error)
	// lay danh sach ban be
	GetListFriend(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AllInfoUser, error)
	// lay thong tin 1 user
	GetInfoUser(ctx context.Context, in *Request, opts ...grpc.CallOption) (*User, error)
}

type chatgRPCClient struct {
	cc *grpc.ClientConn
}

func NewChatgRPCClient(cc *grpc.ClientConn) ChatgRPCClient {
	return &chatgRPCClient{cc}
}

func (c *chatgRPCClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (ChatgRPC_RouteChatClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ChatgRPC_serviceDesc.Streams[0], c.cc, "/MTAMessenger.ChatgRPC/RouteChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatgRPCRouteChatClient{stream}
	return x, nil
}

type ChatgRPC_RouteChatClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatgRPCRouteChatClient struct {
	grpc.ClientStream
}

func (x *chatgRPCRouteChatClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatgRPCRouteChatClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatgRPCClient) Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/MTAMessenger.ChatgRPC/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatgRPCClient) Login(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/MTAMessenger.ChatgRPC/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatgRPCClient) Logout(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/MTAMessenger.ChatgRPC/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatgRPCClient) CheckUser(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/MTAMessenger.ChatgRPC/CheckUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatgRPCClient) LoadWaittingMess(ctx context.Context, in *Request, opts ...grpc.CallOption) (*WaittingMessage, error) {
	out := new(WaittingMessage)
	err := grpc.Invoke(ctx, "/MTAMessenger.ChatgRPC/LoadWaittingMess", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatgRPCClient) GetId(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/MTAMessenger.ChatgRPC/GetId", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatgRPCClient) CreateConversation(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/MTAMessenger.ChatgRPC/CreateConversation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatgRPCClient) GetAllConversation(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AllConversation, error) {
	out := new(AllConversation)
	err := grpc.Invoke(ctx, "/MTAMessenger.ChatgRPC/GetAllConversation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatgRPCClient) AddUidToConversation(ctx context.Context, in *ConversationDetail, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/MTAMessenger.ChatgRPC/AddUidToConversation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatgRPCClient) LoadAllMessOnCid(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AllMessages, error) {
	out := new(AllMessages)
	err := grpc.Invoke(ctx, "/MTAMessenger.ChatgRPC/LoadAllMessOnCid", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatgRPCClient) AddFriend(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/MTAMessenger.ChatgRPC/AddFriend", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatgRPCClient) GetListUser(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AllInfoUser, error) {
	out := new(AllInfoUser)
	err := grpc.Invoke(ctx, "/MTAMessenger.ChatgRPC/GetListUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatgRPCClient) GetListFriend(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AllInfoUser, error) {
	out := new(AllInfoUser)
	err := grpc.Invoke(ctx, "/MTAMessenger.ChatgRPC/GetListFriend", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatgRPCClient) GetInfoUser(ctx context.Context, in *Request, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/MTAMessenger.ChatgRPC/GetInfoUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ChatgRPC service

type ChatgRPCServer interface {
	// rpc LoadMess(Request) returns(WaittingMessage){}
	// dieu huong chat
	RouteChat(ChatgRPC_RouteChatServer) error
	// dang ki thanh vien
	Register(context.Context, *User) (*Response, error)
	// dang nhap
	Login(context.Context, *UserLogin) (*Response, error)
	// logout
	Logout(context.Context, *Request) (*Response, error)
	// kiem tra co user tren database hay ko, truyen vao la 1 sessionkey
	CheckUser(context.Context, *Request) (*Response, error)
	// load tin nhan chua duoc nhan, truyen vao la 1 sessionkey
	LoadWaittingMess(context.Context, *Request) (*WaittingMessage, error)
	// return uid
	GetId(context.Context, *Request) (*Response, error)
	// tra ve cid
	CreateConversation(context.Context, *Request) (*Response, error)
	// dau vao la uid tra ve danh sach Conversation
	GetAllConversation(context.Context, *Request) (*AllConversation, error)
	// add n uid to conversation
	AddUidToConversation(context.Context, *ConversationDetail) (*Response, error)
	// load tat ca tin nhan
	LoadAllMessOnCid(context.Context, *Request) (*AllMessages, error)
	// ket ban
	AddFriend(context.Context, *Request) (*Response, error)
	// lay danh sach tat ca user
	GetListUser(context.Context, *Request) (*AllInfoUser, error)
	// lay danh sach ban be
	GetListFriend(context.Context, *Request) (*AllInfoUser, error)
	// lay thong tin 1 user
	GetInfoUser(context.Context, *Request) (*User, error)
}

func RegisterChatgRPCServer(s *grpc.Server, srv ChatgRPCServer) {
	s.RegisterService(&_ChatgRPC_serviceDesc, srv)
}

func _ChatgRPC_RouteChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatgRPCServer).RouteChat(&chatgRPCRouteChatServer{stream})
}

type ChatgRPC_RouteChatServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatgRPCRouteChatServer struct {
	grpc.ServerStream
}

func (x *chatgRPCRouteChatServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatgRPCRouteChatServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatgRPC_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatgRPCServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MTAMessenger.ChatgRPC/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatgRPCServer).Register(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatgRPC_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatgRPCServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MTAMessenger.ChatgRPC/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatgRPCServer).Login(ctx, req.(*UserLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatgRPC_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatgRPCServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MTAMessenger.ChatgRPC/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatgRPCServer).Logout(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatgRPC_CheckUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatgRPCServer).CheckUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MTAMessenger.ChatgRPC/CheckUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatgRPCServer).CheckUser(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatgRPC_LoadWaittingMess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatgRPCServer).LoadWaittingMess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MTAMessenger.ChatgRPC/LoadWaittingMess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatgRPCServer).LoadWaittingMess(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatgRPC_GetId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatgRPCServer).GetId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MTAMessenger.ChatgRPC/GetId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatgRPCServer).GetId(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatgRPC_CreateConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatgRPCServer).CreateConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MTAMessenger.ChatgRPC/CreateConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatgRPCServer).CreateConversation(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatgRPC_GetAllConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatgRPCServer).GetAllConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MTAMessenger.ChatgRPC/GetAllConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatgRPCServer).GetAllConversation(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatgRPC_AddUidToConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConversationDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatgRPCServer).AddUidToConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MTAMessenger.ChatgRPC/AddUidToConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatgRPCServer).AddUidToConversation(ctx, req.(*ConversationDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatgRPC_LoadAllMessOnCid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatgRPCServer).LoadAllMessOnCid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MTAMessenger.ChatgRPC/LoadAllMessOnCid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatgRPCServer).LoadAllMessOnCid(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatgRPC_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatgRPCServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MTAMessenger.ChatgRPC/AddFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatgRPCServer).AddFriend(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatgRPC_GetListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatgRPCServer).GetListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MTAMessenger.ChatgRPC/GetListUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatgRPCServer).GetListUser(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatgRPC_GetListFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatgRPCServer).GetListFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MTAMessenger.ChatgRPC/GetListFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatgRPCServer).GetListFriend(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatgRPC_GetInfoUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatgRPCServer).GetInfoUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MTAMessenger.ChatgRPC/GetInfoUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatgRPCServer).GetInfoUser(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatgRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MTAMessenger.ChatgRPC",
	HandlerType: (*ChatgRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ChatgRPC_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _ChatgRPC_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _ChatgRPC_Logout_Handler,
		},
		{
			MethodName: "CheckUser",
			Handler:    _ChatgRPC_CheckUser_Handler,
		},
		{
			MethodName: "LoadWaittingMess",
			Handler:    _ChatgRPC_LoadWaittingMess_Handler,
		},
		{
			MethodName: "GetId",
			Handler:    _ChatgRPC_GetId_Handler,
		},
		{
			MethodName: "CreateConversation",
			Handler:    _ChatgRPC_CreateConversation_Handler,
		},
		{
			MethodName: "GetAllConversation",
			Handler:    _ChatgRPC_GetAllConversation_Handler,
		},
		{
			MethodName: "AddUidToConversation",
			Handler:    _ChatgRPC_AddUidToConversation_Handler,
		},
		{
			MethodName: "LoadAllMessOnCid",
			Handler:    _ChatgRPC_LoadAllMessOnCid_Handler,
		},
		{
			MethodName: "AddFriend",
			Handler:    _ChatgRPC_AddFriend_Handler,
		},
		{
			MethodName: "GetListUser",
			Handler:    _ChatgRPC_GetListUser_Handler,
		},
		{
			MethodName: "GetListFriend",
			Handler:    _ChatgRPC_GetListFriend_Handler,
		},
		{
			MethodName: "GetInfoUser",
			Handler:    _ChatgRPC_GetInfoUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RouteChat",
			Handler:       _ChatgRPC_RouteChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "MTAMessengerProto.proto",
}

func init() { proto.RegisterFile("MTAMessengerProto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 763 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdd, 0x4e, 0xdb, 0x48,
	0x14, 0x8e, 0xf3, 0xe7, 0xe4, 0x84, 0x5d, 0xa2, 0x11, 0x0b, 0xde, 0xac, 0x76, 0x15, 0xcd, 0x55,
	0xb4, 0xaa, 0x68, 0x45, 0x6f, 0x80, 0x56, 0xad, 0x22, 0x17, 0x10, 0x95, 0xa1, 0xc8, 0x05, 0xb5,
	0xbd, 0xaa, 0xdc, 0xf8, 0x10, 0x46, 0xd8, 0x9e, 0xd4, 0x33, 0x01, 0xf1, 0x0e, 0x7d, 0x9e, 0x3e,
	0x48, 0x5f, 0xa5, 0x2f, 0x50, 0xcd, 0xd8, 0x0e, 0x76, 0x9c, 0xa4, 0xa5, 0x37, 0x91, 0xcf, 0x37,
	0xe7, 0x7c, 0x73, 0xe6, 0x3b, 0xf3, 0x8d, 0x02, 0x5b, 0x27, 0xe7, 0xc3, 0x13, 0x14, 0x02, 0xa3,
	0x31, 0xc6, 0x67, 0x31, 0x97, 0x7c, 0x7b, 0xa2, 0x7e, 0xc9, 0x5a, 0x7e, 0x81, 0x9a, 0xd0, 0x38,
	0x08, 0x27, 0xf2, 0x8e, 0xda, 0xd0, 0xbe, 0x10, 0x18, 0x3b, 0x7c, 0xcc, 0x22, 0xd2, 0x83, 0xd6,
	0x54, 0x60, 0x1c, 0x79, 0x21, 0x5a, 0x46, 0xdf, 0x18, 0xb4, 0xdd, 0x59, 0xac, 0xd6, 0x26, 0x9e,
	0x10, 0xb7, 0x3c, 0xf6, 0xad, 0x6a, 0xb2, 0x96, 0xc5, 0xf4, 0xab, 0x01, 0x75, 0xc5, 0x42, 0xba,
	0x50, 0x9b, 0x32, 0x5f, 0xd7, 0xd6, 0x5d, 0xf5, 0x59, 0xa0, 0xac, 0xae, 0xa0, 0xac, 0x15, 0x29,
	0xc9, 0x06, 0x34, 0x26, 0x57, 0x3c, 0x42, 0xab, 0xae, 0x17, 0x92, 0x40, 0xa1, 0x18, 0x7a, 0x2c,
	0xb0, 0x1a, 0x09, 0xaa, 0x03, 0xd2, 0x87, 0xce, 0x28, 0x46, 0x4f, 0xa2, 0x7f, 0xce, 0x42, 0xb4,
	0x9a, 0x7a, 0x2d, 0x0f, 0x91, 0x4d, 0x68, 0x7a, 0x23, 0xc9, 0x6e, 0xd0, 0x32, 0xfb, 0xc6, 0xa0,
	0xe5, 0xa6, 0x11, 0x75, 0xa0, 0xe5, 0xa2, 0x98, 0xf0, 0x48, 0xe8, 0x6e, 0xe2, 0xf4, 0x3b, 0x3b,
	0x7c, 0x16, 0xab, 0x7d, 0x47, 0x57, 0x38, 0xba, 0xd6, 0x47, 0x68, 0xb9, 0x49, 0x40, 0xfe, 0x84,
	0x2a, 0xcb, 0x3a, 0xaf, 0x32, 0x9f, 0xbe, 0x05, 0xd3, 0xc5, 0xcf, 0x53, 0x14, 0x92, 0xfc, 0x07,
	0x20, 0x50, 0x08, 0xc6, 0xa3, 0x6b, 0xbc, 0x4b, 0xe9, 0x72, 0x08, 0xb1, 0xc0, 0x8c, 0x93, 0xd4,
	0x54, 0x95, 0x2c, 0x2c, 0x91, 0x7e, 0x33, 0xc0, 0x54, 0x73, 0xf3, 0xc6, 0xa8, 0xe4, 0x0d, 0xef,
	0xe5, 0x0d, 0x99, 0x3f, 0xb7, 0x4f, 0xb5, 0xb4, 0x4f, 0x17, 0x6a, 0xa3, 0x19, 0x9d, 0xfa, 0x54,
	0x3b, 0x8f, 0x78, 0x24, 0x31, 0x92, 0xa9, 0x88, 0x59, 0x38, 0x2f, 0x63, 0xbd, 0x2c, 0xe3, 0x4c,
	0x86, 0x66, 0x5e, 0x86, 0x7f, 0xa0, 0x7d, 0x19, 0xf3, 0xf0, 0xa3, 0x9e, 0xb1, 0x99, 0x28, 0xa7,
	0x80, 0x53, 0x2f, 0x29, 0x91, 0xfc, 0x82, 0xf9, 0x56, 0x2b, 0x99, 0x98, 0x0e, 0xe8, 0x17, 0x03,
	0xd6, 0x6c, 0x1e, 0xdd, 0x60, 0x2c, 0x3c, 0xc9, 0x78, 0x94, 0xf5, 0x69, 0xdc, 0xf7, 0xf9, 0x3f,
	0x74, 0xf3, 0x19, 0xa7, 0xf7, 0x17, 0xa8, 0x84, 0x2b, 0x15, 0x1c, 0x4f, 0x14, 0x1b, 0xcf, 0x21,
	0x73, 0x2a, 0x35, 0xe6, 0x55, 0xa2, 0x1f, 0x60, 0x7d, 0x18, 0x04, 0x85, 0x86, 0x0e, 0xa1, 0x1b,
	0x30, 0x21, 0xf3, 0x98, 0x65, 0xf4, 0x6b, 0x83, 0xce, 0x4e, 0x6f, 0x3b, 0xef, 0xa4, 0xed, 0x7c,
	0x86, 0x5b, 0xaa, 0xa1, 0x0e, 0xac, 0xbf, 0xf3, 0x98, 0x94, 0x2c, 0x1a, 0x67, 0x53, 0xdc, 0x83,
	0xb5, 0xdb, 0x14, 0x0a, 0x51, 0x88, 0x94, 0xf6, 0xaf, 0x22, 0x6d, 0x9a, 0xec, 0x16, 0x52, 0xe9,
	0x0b, 0xe8, 0x0c, 0x83, 0x20, 0x5d, 0x13, 0xe4, 0x31, 0x98, 0x5e, 0x10, 0xfc, 0x9c, 0x24, 0xcb,
	0xa2, 0xcf, 0x74, 0xfd, 0x71, 0x74, 0xc9, 0xb5, 0x5d, 0x1f, 0xe9, 0x7a, 0xe5, 0xc7, 0xb4, 0x9e,
	0x14, 0xeb, 0x55, 0x92, 0x9b, 0xa5, 0xd0, 0xf7, 0x40, 0xf2, 0x47, 0x7b, 0x85, 0x52, 0x99, 0xaf,
	0xa8, 0x6d, 0x6d, 0xd9, 0x0d, 0xcc, 0x4d, 0x36, 0x7d, 0x24, 0xaa, 0xfd, 0x9a, 0x42, 0xa6, 0xcc,
	0xdf, 0xf9, 0x6e, 0x42, 0xcb, 0xbe, 0xf2, 0xe4, 0xd8, 0x3d, 0xb3, 0xc9, 0x4b, 0x68, 0xbb, 0x7c,
	0x2a, 0x51, 0x01, 0x64, 0xf1, 0x81, 0x7a, 0x8b, 0x61, 0x5a, 0x19, 0x18, 0x4f, 0x0c, 0xb2, 0xab,
	0x4c, 0x3d, 0x66, 0x42, 0x62, 0x4c, 0x16, 0x1c, 0xa8, 0xb7, 0x59, 0xc4, 0xb2, 0x07, 0x80, 0x56,
	0xc8, 0x3e, 0x34, 0x92, 0x87, 0x70, 0xab, 0x5c, 0xa6, 0x17, 0x56, 0xd4, 0xee, 0x41, 0xd3, 0xe1,
	0x63, 0x3e, 0x2d, 0xf5, 0x9c, 0x3e, 0x09, 0x2b, 0x4a, 0x9f, 0x43, 0xdb, 0x56, 0x4e, 0xd2, 0x33,
	0x79, 0x70, 0xf5, 0x6b, 0xe8, 0x3a, 0xdc, 0xf3, 0xf3, 0xb7, 0x6c, 0x19, 0xc9, 0xbf, 0x45, 0x78,
	0xee, 0x62, 0xd2, 0x0a, 0xd9, 0x85, 0xc6, 0x11, 0xca, 0x63, 0xff, 0xe1, 0x5d, 0x1c, 0x00, 0xb1,
	0xf5, 0x4b, 0x51, 0x70, 0xd1, 0x83, 0x69, 0x1c, 0x20, 0x47, 0x28, 0xe7, 0xcd, 0xf8, 0x6b, 0xc7,
	0x99, 0xab, 0xa2, 0x15, 0xe2, 0xc2, 0xc6, 0xd0, 0xf7, 0x2f, 0x98, 0x7f, 0xce, 0x0b, 0x7c, 0xfd,
	0xe5, 0x16, 0x4e, 0x6e, 0xf5, 0x8a, 0x0e, 0x0f, 0x13, 0xb9, 0x53, 0x1b, 0xbe, 0x89, 0x6c, 0xb6,
	0x54, 0xad, 0xbf, 0x4b, 0xfd, 0x65, 0xce, 0x4d, 0x86, 0x3e, 0xf4, 0xfd, 0xc3, 0x98, 0x61, 0xf4,
	0x1b, 0x72, 0x0f, 0xa1, 0x73, 0x84, 0xd2, 0x61, 0x42, 0xae, 0xba, 0x34, 0xe5, 0x06, 0x32, 0xeb,
	0xd3, 0x0a, 0xb1, 0xe1, 0x8f, 0x94, 0x62, 0x75, 0x13, 0x2b, 0x49, 0xf6, 0x75, 0x1f, 0xb3, 0x07,
	0x65, 0x09, 0xc5, 0x02, 0x17, 0xd2, 0xca, 0xa7, 0xa6, 0xfe, 0x63, 0xf2, 0xf4, 0x47, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x6b, 0x7f, 0xaa, 0x08, 0xb3, 0x08, 0x00, 0x00,
}
