// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comment/comment.proto

package comment

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type CommentRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	PostId               int64    `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id"`
	TextComment          string   `protobuf:"bytes,3,opt,name=textComment,proto3" json:"textComment"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentRequest) Reset()         { *m = CommentRequest{} }
func (m *CommentRequest) String() string { return proto.CompactTextString(m) }
func (*CommentRequest) ProtoMessage()    {}
func (*CommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_885638bbfd25b68b, []int{0}
}
func (m *CommentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentRequest.Merge(m, src)
}
func (m *CommentRequest) XXX_Size() int {
	return m.Size()
}
func (m *CommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommentRequest proto.InternalMessageInfo

func (m *CommentRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CommentRequest) GetPostId() int64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func (m *CommentRequest) GetTextComment() string {
	if m != nil {
		return m.TextComment
	}
	return ""
}

type CommentResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	PostId               int64    `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id"`
	UserId               int64    `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id"`
	UserName             string   `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name"`
	TextComment          string   `protobuf:"bytes,5,opt,name=textComment,proto3" json:"textComment"`
	PostTitle            string   `protobuf:"bytes,6,opt,name=post_title,json=postTitle,proto3" json:"post_title"`
	CreatedAt            string   `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommentResponse) Reset()         { *m = CommentResponse{} }
func (m *CommentResponse) String() string { return proto.CompactTextString(m) }
func (*CommentResponse) ProtoMessage()    {}
func (*CommentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_885638bbfd25b68b, []int{1}
}
func (m *CommentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommentResponse.Merge(m, src)
}
func (m *CommentResponse) XXX_Size() int {
	return m.Size()
}
func (m *CommentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommentResponse proto.InternalMessageInfo

func (m *CommentResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CommentResponse) GetPostId() int64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func (m *CommentResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CommentResponse) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *CommentResponse) GetTextComment() string {
	if m != nil {
		return m.TextComment
	}
	return ""
}

func (m *CommentResponse) GetPostTitle() string {
	if m != nil {
		return m.PostTitle
	}
	return ""
}

func (m *CommentResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*CommentRequest)(nil), "comment.CommentRequest")
	proto.RegisterType((*CommentResponse)(nil), "comment.CommentResponse")
}

func init() { proto.RegisterFile("comment/comment.proto", fileDescriptor_885638bbfd25b68b) }

var fileDescriptor_885638bbfd25b68b = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xeb, 0x06, 0x12, 0x72, 0x48, 0x05, 0x59, 0x42, 0xb1, 0x40, 0x44, 0x51, 0xa6, 0x4e,
	0x45, 0x82, 0x27, 0x28, 0x4c, 0x5d, 0x18, 0x52, 0x36, 0x86, 0xca, 0xc4, 0x37, 0x58, 0x22, 0x71,
	0x88, 0xaf, 0x88, 0x47, 0xe1, 0x91, 0x18, 0xfb, 0x08, 0x28, 0xbc, 0x08, 0xb2, 0x93, 0x56, 0x40,
	0x17, 0xa6, 0xe4, 0xbe, 0xcf, 0xbe, 0x3b, 0xfd, 0x86, 0xb3, 0xd2, 0x54, 0x15, 0xd6, 0x74, 0x35,
	0x7c, 0x67, 0x4d, 0x6b, 0xc8, 0xf0, 0x68, 0x28, 0xf3, 0x47, 0x98, 0xdc, 0xf5, 0xbf, 0x05, 0xbe,
	0xac, 0xd1, 0x12, 0x9f, 0xc0, 0x58, 0x2b, 0xc1, 0x32, 0x36, 0x0d, 0x8a, 0xb1, 0x56, 0x3c, 0x81,
	0xa8, 0x31, 0x96, 0x56, 0x5a, 0x89, 0xb1, 0x87, 0xa1, 0x2b, 0x17, 0x8a, 0x67, 0x70, 0x4c, 0xf8,
	0x46, 0xc3, 0x75, 0x11, 0x64, 0x6c, 0x1a, 0x17, 0x3f, 0x51, 0xbe, 0x61, 0x70, 0xb2, 0xeb, 0x6e,
	0x1b, 0x53, 0x5b, 0xfc, 0x7f, 0xfb, 0x04, 0xa2, 0xb5, 0xc5, 0xd6, 0x89, 0xa0, 0x17, 0xae, 0x5c,
	0x28, 0x7e, 0x01, 0xb1, 0x17, 0xb5, 0xac, 0x50, 0x1c, 0xf8, 0xa9, 0x47, 0x0e, 0xdc, 0xcb, 0x0a,
	0xff, 0x2e, 0x75, 0xb8, 0xb7, 0x14, 0xbf, 0x04, 0xf0, 0x03, 0x49, 0xd3, 0x33, 0x8a, 0xd0, 0x1f,
	0x88, 0x1d, 0x79, 0x70, 0xc0, 0xe9, 0xb2, 0x45, 0x49, 0xa8, 0x56, 0x92, 0x44, 0xd4, 0xeb, 0x81,
	0xcc, 0xe9, 0x7a, 0xb9, 0xcb, 0x6b, 0x89, 0xed, 0xab, 0x2e, 0x91, 0xcf, 0x01, 0xa4, 0x52, 0xdb,
	0xee, 0xc9, 0x6c, 0x1b, 0xf4, 0xef, 0x58, 0xcf, 0xc5, 0xbe, 0xe8, 0x13, 0xc9, 0x47, 0xb7, 0xa7,
	0x1f, 0x5d, 0xca, 0x36, 0x5d, 0xca, 0x3e, 0xbb, 0x94, 0xbd, 0x7f, 0xa5, 0xa3, 0xa7, 0xd0, 0x3f,
	0xd3, 0xcd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x91, 0x0c, 0x76, 0x58, 0xbf, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommentServiceClient interface {
	AddComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error)
}

type commentServiceClient struct {
	cc *grpc.ClientConn
}

func NewCommentServiceClient(cc *grpc.ClientConn) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) AddComment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, "/comment.CommentService/addComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
type CommentServiceServer interface {
	AddComment(context.Context, *CommentRequest) (*CommentResponse, error)
}

// UnimplementedCommentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCommentServiceServer struct {
}

func (*UnimplementedCommentServiceServer) AddComment(ctx context.Context, req *CommentRequest) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}

func RegisterCommentServiceServer(s *grpc.Server, srv CommentServiceServer) {
	s.RegisterService(&_CommentService_serviceDesc, srv)
}

func _CommentService_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.CommentService/AddComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).AddComment(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comment.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addComment",
			Handler:    _CommentService_AddComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment/comment.proto",
}

func (m *CommentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.TextComment) > 0 {
		i -= len(m.TextComment)
		copy(dAtA[i:], m.TextComment)
		i = encodeVarintComment(dAtA, i, uint64(len(m.TextComment)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PostId != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.PostId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CommentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.CreatedAt) > 0 {
		i -= len(m.CreatedAt)
		copy(dAtA[i:], m.CreatedAt)
		i = encodeVarintComment(dAtA, i, uint64(len(m.CreatedAt)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.PostTitle) > 0 {
		i -= len(m.PostTitle)
		copy(dAtA[i:], m.PostTitle)
		i = encodeVarintComment(dAtA, i, uint64(len(m.PostTitle)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.TextComment) > 0 {
		i -= len(m.TextComment)
		copy(dAtA[i:], m.TextComment)
		i = encodeVarintComment(dAtA, i, uint64(len(m.TextComment)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.UserName) > 0 {
		i -= len(m.UserName)
		copy(dAtA[i:], m.UserName)
		i = encodeVarintComment(dAtA, i, uint64(len(m.UserName)))
		i--
		dAtA[i] = 0x22
	}
	if m.UserId != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.UserId))
		i--
		dAtA[i] = 0x18
	}
	if m.PostId != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.PostId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintComment(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintComment(dAtA []byte, offset int, v uint64) int {
	offset -= sovComment(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CommentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovComment(uint64(m.Id))
	}
	if m.PostId != 0 {
		n += 1 + sovComment(uint64(m.PostId))
	}
	l = len(m.TextComment)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CommentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovComment(uint64(m.Id))
	}
	if m.PostId != 0 {
		n += 1 + sovComment(uint64(m.PostId))
	}
	if m.UserId != 0 {
		n += 1 + sovComment(uint64(m.UserId))
	}
	l = len(m.UserName)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	l = len(m.TextComment)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	l = len(m.PostTitle)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	l = len(m.CreatedAt)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovComment(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozComment(x uint64) (n int) {
	return sovComment(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CommentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComment
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CommentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostId", wireType)
			}
			m.PostId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PostId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TextComment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TextComment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipComment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComment
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CommentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComment
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CommentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostId", wireType)
			}
			m.PostId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PostId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TextComment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TextComment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostTitle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostTitle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipComment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComment
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipComment(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowComment
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
					return 0, ErrIntOverflowComment
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowComment
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
			if length < 0 {
				return 0, ErrInvalidLengthComment
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupComment
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthComment
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthComment        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowComment          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupComment = fmt.Errorf("proto: unexpected end of group")
)
