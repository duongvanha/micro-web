// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transfers.proto

package movie

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

type Page struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Page) Reset()         { *m = Page{} }
func (m *Page) String() string { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()    {}
func (*Page) Descriptor() ([]byte, []int) {
	return fileDescriptor_transfers_0d024657fb779c7e, []int{0}
}
func (m *Page) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Page.Unmarshal(m, b)
}
func (m *Page) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Page.Marshal(b, m, deterministic)
}
func (dst *Page) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Page.Merge(dst, src)
}
func (m *Page) XXX_Size() int {
	return xxx_messageInfo_Page.Size(m)
}
func (m *Page) XXX_DiscardUnknown() {
	xxx_messageInfo_Page.DiscardUnknown(m)
}

var xxx_messageInfo_Page proto.InternalMessageInfo

func (m *Page) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type Chap struct {
	SeverName            string   `protobuf:"bytes,1,opt,name=SeverName" json:"SeverName,omitempty"`
	Urls                 []string `protobuf:"bytes,2,rep,name=Urls" json:"Urls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chap) Reset()         { *m = Chap{} }
func (m *Chap) String() string { return proto.CompactTextString(m) }
func (*Chap) ProtoMessage()    {}
func (*Chap) Descriptor() ([]byte, []int) {
	return fileDescriptor_transfers_0d024657fb779c7e, []int{1}
}
func (m *Chap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chap.Unmarshal(m, b)
}
func (m *Chap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chap.Marshal(b, m, deterministic)
}
func (dst *Chap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chap.Merge(dst, src)
}
func (m *Chap) XXX_Size() int {
	return xxx_messageInfo_Chap.Size(m)
}
func (m *Chap) XXX_DiscardUnknown() {
	xxx_messageInfo_Chap.DiscardUnknown(m)
}

var xxx_messageInfo_Chap proto.InternalMessageInfo

func (m *Chap) GetSeverName() string {
	if m != nil {
		return m.SeverName
	}
	return ""
}

func (m *Chap) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

type Movie struct {
	NameVi               string   `protobuf:"bytes,1,opt,name=NameVi" json:"NameVi,omitempty"`
	NameEn               string   `protobuf:"bytes,2,opt,name=NameEn" json:"NameEn,omitempty"`
	Image                string   `protobuf:"bytes,3,opt,name=Image" json:"Image,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=Status" json:"Status,omitempty"`
	IMDb                 string   `protobuf:"bytes,5,opt,name=IMDb" json:"IMDb,omitempty"`
	Director             string   `protobuf:"bytes,6,opt,name=Director" json:"Director,omitempty"`
	Country              string   `protobuf:"bytes,7,opt,name=Country" json:"Country,omitempty"`
	Year                 string   `protobuf:"bytes,8,opt,name=Year" json:"Year,omitempty"`
	Date                 string   `protobuf:"bytes,9,opt,name=Date" json:"Date,omitempty"`
	Time                 string   `protobuf:"bytes,10,opt,name=Time" json:"Time,omitempty"`
	Cam                  string   `protobuf:"bytes,11,opt,name=Cam" json:"Cam,omitempty"`
	Quality              string   `protobuf:"bytes,12,opt,name=Quality" json:"Quality,omitempty"`
	Sub                  string   `protobuf:"bytes,13,opt,name=Sub" json:"Sub,omitempty"`
	Type                 string   `protobuf:"bytes,14,opt,name=Type" json:"Type,omitempty"`
	Manufacturer         string   `protobuf:"bytes,15,opt,name=Manufacturer" json:"Manufacturer,omitempty"`
	View                 string   `protobuf:"bytes,16,opt,name=View" json:"View,omitempty"`
	Url                  string   `protobuf:"bytes,17,opt,name=Url" json:"Url,omitempty"`
	Movies               string   `protobuf:"bytes,18,opt,name=Movies" json:"Movies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Movie) Reset()         { *m = Movie{} }
func (m *Movie) String() string { return proto.CompactTextString(m) }
func (*Movie) ProtoMessage()    {}
func (*Movie) Descriptor() ([]byte, []int) {
	return fileDescriptor_transfers_0d024657fb779c7e, []int{2}
}
func (m *Movie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Movie.Unmarshal(m, b)
}
func (m *Movie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Movie.Marshal(b, m, deterministic)
}
func (dst *Movie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Movie.Merge(dst, src)
}
func (m *Movie) XXX_Size() int {
	return xxx_messageInfo_Movie.Size(m)
}
func (m *Movie) XXX_DiscardUnknown() {
	xxx_messageInfo_Movie.DiscardUnknown(m)
}

var xxx_messageInfo_Movie proto.InternalMessageInfo

func (m *Movie) GetNameVi() string {
	if m != nil {
		return m.NameVi
	}
	return ""
}

func (m *Movie) GetNameEn() string {
	if m != nil {
		return m.NameEn
	}
	return ""
}

func (m *Movie) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Movie) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Movie) GetIMDb() string {
	if m != nil {
		return m.IMDb
	}
	return ""
}

func (m *Movie) GetDirector() string {
	if m != nil {
		return m.Director
	}
	return ""
}

func (m *Movie) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Movie) GetYear() string {
	if m != nil {
		return m.Year
	}
	return ""
}

func (m *Movie) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Movie) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *Movie) GetCam() string {
	if m != nil {
		return m.Cam
	}
	return ""
}

func (m *Movie) GetQuality() string {
	if m != nil {
		return m.Quality
	}
	return ""
}

func (m *Movie) GetSub() string {
	if m != nil {
		return m.Sub
	}
	return ""
}

func (m *Movie) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Movie) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *Movie) GetView() string {
	if m != nil {
		return m.View
	}
	return ""
}

func (m *Movie) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Movie) GetMovies() string {
	if m != nil {
		return m.Movies
	}
	return ""
}

type ListMovie struct {
	Movies               []*Movie `protobuf:"bytes,1,rep,name=movies" json:"movies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMovie) Reset()         { *m = ListMovie{} }
func (m *ListMovie) String() string { return proto.CompactTextString(m) }
func (*ListMovie) ProtoMessage()    {}
func (*ListMovie) Descriptor() ([]byte, []int) {
	return fileDescriptor_transfers_0d024657fb779c7e, []int{3}
}
func (m *ListMovie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMovie.Unmarshal(m, b)
}
func (m *ListMovie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMovie.Marshal(b, m, deterministic)
}
func (dst *ListMovie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMovie.Merge(dst, src)
}
func (m *ListMovie) XXX_Size() int {
	return xxx_messageInfo_ListMovie.Size(m)
}
func (m *ListMovie) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMovie.DiscardUnknown(m)
}

var xxx_messageInfo_ListMovie proto.InternalMessageInfo

func (m *ListMovie) GetMovies() []*Movie {
	if m != nil {
		return m.Movies
	}
	return nil
}

func init() {
	proto.RegisterType((*Page)(nil), "movie.page")
	proto.RegisterType((*Chap)(nil), "movie.Chap")
	proto.RegisterType((*Movie)(nil), "movie.Movie")
	proto.RegisterType((*ListMovie)(nil), "movie.ListMovie")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MovieRepository service

type MovieRepositoryClient interface {
	GetByPage(ctx context.Context, in *Page, opts ...grpc.CallOption) (*ListMovie, error)
}

type movieRepositoryClient struct {
	cc *grpc.ClientConn
}

func NewMovieRepositoryClient(cc *grpc.ClientConn) MovieRepositoryClient {
	return &movieRepositoryClient{cc}
}

func (c *movieRepositoryClient) GetByPage(ctx context.Context, in *Page, opts ...grpc.CallOption) (*ListMovie, error) {
	out := new(ListMovie)
	err := grpc.Invoke(ctx, "/movie.MovieRepository/GetByPage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MovieRepository service

type MovieRepositoryServer interface {
	GetByPage(context.Context, *Page) (*ListMovie, error)
}

func RegisterMovieRepositoryServer(s *grpc.Server, srv MovieRepositoryServer) {
	s.RegisterService(&_MovieRepository_serviceDesc, srv)
}

func _MovieRepository_GetByPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Page)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieRepositoryServer).GetByPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movie.MovieRepository/GetByPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieRepositoryServer).GetByPage(ctx, req.(*Page))
	}
	return interceptor(ctx, in, info, handler)
}

var _MovieRepository_serviceDesc = grpc.ServiceDesc{
	ServiceName: "movie.MovieRepository",
	HandlerType: (*MovieRepositoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByPage",
			Handler:    _MovieRepository_GetByPage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transfers.proto",
}

func init() { proto.RegisterFile("transfers.proto", fileDescriptor_transfers_0d024657fb779c7e) }

var fileDescriptor_transfers_0d024657fb779c7e = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x4f, 0xcb, 0xda, 0x40,
	0x10, 0xc6, 0xd1, 0x98, 0x68, 0x46, 0x5b, 0xed, 0x52, 0xca, 0x20, 0x3d, 0x48, 0xe8, 0x41, 0x7a,
	0x10, 0x6a, 0x2f, 0xbd, 0xf4, 0x52, 0x95, 0x22, 0xd4, 0xd2, 0x26, 0x55, 0xe8, 0x71, 0x23, 0x6b,
	0xbb, 0x90, 0x7f, 0x6c, 0x36, 0x96, 0x7c, 0x9a, 0x7e, 0xd5, 0x97, 0xd9, 0x89, 0xca, 0x7b, 0xca,
	0xf3, 0xfc, 0x66, 0xe6, 0x99, 0xb0, 0xbb, 0x30, 0xb5, 0x46, 0x16, 0xf5, 0x45, 0x99, 0x7a, 0x55,
	0x99, 0xd2, 0x96, 0xc2, 0xcf, 0xcb, 0xab, 0x56, 0xd1, 0x1c, 0x06, 0x95, 0xfc, 0xa3, 0x84, 0xe0,
	0x2f, 0xf6, 0x16, 0xbd, 0xa5, 0x1f, 0x3b, 0x1d, 0x7d, 0x82, 0xc1, 0xe6, 0xaf, 0xac, 0xc4, 0x5b,
	0x08, 0x13, 0x75, 0x55, 0xe6, 0xbb, 0xcc, 0xb9, 0x21, 0x8c, 0x1f, 0x80, 0x26, 0x8f, 0x26, 0xab,
	0xb1, 0xbf, 0xf0, 0x96, 0x61, 0xec, 0x74, 0xf4, 0xdf, 0x03, 0xff, 0x40, 0xf9, 0xe2, 0x0d, 0x04,
	0xd4, 0x75, 0xd2, 0xdd, 0x60, 0xe7, 0x6e, 0x7c, 0x57, 0x60, 0xff, 0xc1, 0x77, 0x85, 0x78, 0x0d,
	0xfe, 0x3e, 0xa7, 0x1f, 0xf1, 0x1c, 0x66, 0x43, 0xdd, 0x89, 0x95, 0xb6, 0xa9, 0x71, 0xc0, 0xdd,
	0xec, 0x68, 0xf7, 0xfe, 0xb0, 0x4d, 0xd1, 0x77, 0xd4, 0x69, 0x31, 0x87, 0xd1, 0x56, 0x1b, 0x75,
	0xb6, 0xa5, 0xc1, 0xc0, 0xf1, 0xbb, 0x17, 0x08, 0xc3, 0x4d, 0xd9, 0x14, 0xd6, 0xb4, 0x38, 0x74,
	0xa5, 0x9b, 0xa5, 0xa4, 0xdf, 0x4a, 0x1a, 0x1c, 0x71, 0x12, 0x69, 0x62, 0x5b, 0x69, 0x15, 0x86,
	0xcc, 0x48, 0x13, 0xfb, 0xa5, 0x73, 0x85, 0xc0, 0x8c, 0xb4, 0x98, 0x81, 0xb7, 0x91, 0x39, 0x8e,
	0x1d, 0x22, 0x49, 0x7b, 0x7e, 0x36, 0x32, 0xd3, 0xb6, 0xc5, 0x09, 0xef, 0xe9, 0x2c, 0xf5, 0x26,
	0x4d, 0x8a, 0x2f, 0xb8, 0x37, 0x69, 0x52, 0x97, 0xd8, 0x56, 0x0a, 0x5f, 0x76, 0x89, 0x6d, 0xa5,
	0x44, 0x04, 0x93, 0x83, 0x2c, 0x9a, 0x8b, 0x3c, 0xdb, 0xc6, 0x28, 0x83, 0x53, 0x57, 0x7b, 0xc6,
	0x68, 0xee, 0xa4, 0xd5, 0x3f, 0x9c, 0xf1, 0x1c, 0x69, 0x4a, 0x3f, 0x9a, 0x0c, 0x5f, 0x71, 0xfa,
	0xd1, 0x64, 0x74, 0x72, 0xee, 0x22, 0x6a, 0x14, 0x7c, 0x72, 0xec, 0xa2, 0x0f, 0x10, 0x7e, 0xd3,
	0xb5, 0xe5, 0x4b, 0x7a, 0x07, 0x41, 0xce, 0x4d, 0xbd, 0x85, 0xb7, 0x1c, 0xaf, 0x27, 0x2b, 0x67,
	0x57, 0xae, 0x1a, 0x77, 0xb5, 0xf5, 0x67, 0x98, 0x32, 0x50, 0x55, 0x59, 0x6b, 0x5b, 0x9a, 0x56,
	0xbc, 0x87, 0xf0, 0xab, 0xb2, 0x5f, 0xda, 0x1f, 0x74, 0x49, 0xe3, 0x6e, 0x8a, 0xde, 0xce, 0x7c,
	0xd6, 0x99, 0xfb, 0x92, 0x34, 0x70, 0xef, 0xee, 0xe3, 0x53, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19,
	0xbb, 0x93, 0x35, 0x8a, 0x02, 0x00, 0x00,
}
