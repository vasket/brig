// Code generated by protoc-gen-gogo.
// source: protocol.proto
// DO NOT EDIT!

/*
	Package wire is a generated protocol buffer package.

	It is generated from these files:
		protocol.proto

	It has these top-level messages:
		Request
		StoreVersionResponse
		FetchResponse
		Response
*/
package wire

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import brig_store "github.com/disorganizer/brig/store/wire"

import github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type RequestType int32

const (
	RequestType_INVALID       RequestType = 0
	RequestType_FETCH         RequestType = 1
	RequestType_STORE_VERSION RequestType = 2
	RequestType_UPDATE_FILE   RequestType = 3
)

var RequestType_name = map[int32]string{
	0: "INVALID",
	1: "FETCH",
	2: "STORE_VERSION",
	3: "UPDATE_FILE",
}
var RequestType_value = map[string]int32{
	"INVALID":       0,
	"FETCH":         1,
	"STORE_VERSION": 2,
	"UPDATE_FILE":   3,
}

func (x RequestType) Enum() *RequestType {
	p := new(RequestType)
	*p = x
	return p
}
func (x RequestType) String() string {
	return proto.EnumName(RequestType_name, int32(x))
}
func (x *RequestType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RequestType_value, data, "RequestType")
	if err != nil {
		return err
	}
	*x = RequestType(value)
	return nil
}
func (RequestType) EnumDescriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{0} }

type Request struct {
	ReqType          *RequestType `protobuf:"varint,1,req,name=req_type,enum=brig.transfer.RequestType" json:"req_type,omitempty"`
	ID               *int64       `protobuf:"varint,2,req,name=ID" json:"ID,omitempty"`
	Nonce            *int64       `protobuf:"varint,3,req,name=nonce" json:"nonce,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{0} }

func (m *Request) GetReqType() RequestType {
	if m != nil && m.ReqType != nil {
		return *m.ReqType
	}
	return RequestType_INVALID
}

func (m *Request) GetID() int64 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *Request) GetNonce() int64 {
	if m != nil && m.Nonce != nil {
		return *m.Nonce
	}
	return 0
}

type StoreVersionResponse struct {
	Version          *int32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *StoreVersionResponse) Reset()                    { *m = StoreVersionResponse{} }
func (m *StoreVersionResponse) String() string            { return proto.CompactTextString(m) }
func (*StoreVersionResponse) ProtoMessage()               {}
func (*StoreVersionResponse) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{1} }

func (m *StoreVersionResponse) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

type FetchResponse struct {
	Store            *brig_store.Store `protobuf:"bytes,1,req,name=store" json:"store,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *FetchResponse) Reset()                    { *m = FetchResponse{} }
func (m *FetchResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchResponse) ProtoMessage()               {}
func (*FetchResponse) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{2} }

func (m *FetchResponse) GetStore() *brig_store.Store {
	if m != nil {
		return m.Store
	}
	return nil
}

type Response struct {
	ReqType          *RequestType          `protobuf:"varint,1,req,name=req_type,enum=brig.transfer.RequestType" json:"req_type,omitempty"`
	ID               *int64                `protobuf:"varint,2,req,name=ID" json:"ID,omitempty"`
	Nonce            *int64                `protobuf:"varint,3,req,name=nonce" json:"nonce,omitempty"`
	Error            *string               `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
	StoreVersionResp *StoreVersionResponse `protobuf:"bytes,5,opt,name=store_version_resp" json:"store_version_resp,omitempty"`
	FetchResp        *FetchResponse        `protobuf:"bytes,6,opt,name=fetch_resp" json:"fetch_resp,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{3} }

func (m *Response) GetReqType() RequestType {
	if m != nil && m.ReqType != nil {
		return *m.ReqType
	}
	return RequestType_INVALID
}

func (m *Response) GetID() int64 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *Response) GetNonce() int64 {
	if m != nil && m.Nonce != nil {
		return *m.Nonce
	}
	return 0
}

func (m *Response) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func (m *Response) GetStoreVersionResp() *StoreVersionResponse {
	if m != nil {
		return m.StoreVersionResp
	}
	return nil
}

func (m *Response) GetFetchResp() *FetchResponse {
	if m != nil {
		return m.FetchResp
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "brig.transfer.Request")
	proto.RegisterType((*StoreVersionResponse)(nil), "brig.transfer.StoreVersionResponse")
	proto.RegisterType((*FetchResponse)(nil), "brig.transfer.FetchResponse")
	proto.RegisterType((*Response)(nil), "brig.transfer.Response")
	proto.RegisterEnum("brig.transfer.RequestType", RequestType_name, RequestType_value)
}
func (m *Request) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Request) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ReqType == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		data[i] = 0x8
		i++
		i = encodeVarintProtocol(data, i, uint64(*m.ReqType))
	}
	if m.ID == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		data[i] = 0x10
		i++
		i = encodeVarintProtocol(data, i, uint64(*m.ID))
	}
	if m.Nonce == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		data[i] = 0x18
		i++
		i = encodeVarintProtocol(data, i, uint64(*m.Nonce))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *StoreVersionResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StoreVersionResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Version == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		data[i] = 0x8
		i++
		i = encodeVarintProtocol(data, i, uint64(*m.Version))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *FetchResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *FetchResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Store == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		data[i] = 0xa
		i++
		i = encodeVarintProtocol(data, i, uint64(m.Store.Size()))
		n1, err := m.Store.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Response) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Response) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ReqType == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		data[i] = 0x8
		i++
		i = encodeVarintProtocol(data, i, uint64(*m.ReqType))
	}
	if m.ID == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		data[i] = 0x10
		i++
		i = encodeVarintProtocol(data, i, uint64(*m.ID))
	}
	if m.Nonce == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		data[i] = 0x18
		i++
		i = encodeVarintProtocol(data, i, uint64(*m.Nonce))
	}
	if m.Error != nil {
		data[i] = 0x22
		i++
		i = encodeVarintProtocol(data, i, uint64(len(*m.Error)))
		i += copy(data[i:], *m.Error)
	}
	if m.StoreVersionResp != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintProtocol(data, i, uint64(m.StoreVersionResp.Size()))
		n2, err := m.StoreVersionResp.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.FetchResp != nil {
		data[i] = 0x32
		i++
		i = encodeVarintProtocol(data, i, uint64(m.FetchResp.Size()))
		n3, err := m.FetchResp.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Protocol(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Protocol(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintProtocol(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Request) Size() (n int) {
	var l int
	_ = l
	if m.ReqType != nil {
		n += 1 + sovProtocol(uint64(*m.ReqType))
	}
	if m.ID != nil {
		n += 1 + sovProtocol(uint64(*m.ID))
	}
	if m.Nonce != nil {
		n += 1 + sovProtocol(uint64(*m.Nonce))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StoreVersionResponse) Size() (n int) {
	var l int
	_ = l
	if m.Version != nil {
		n += 1 + sovProtocol(uint64(*m.Version))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FetchResponse) Size() (n int) {
	var l int
	_ = l
	if m.Store != nil {
		l = m.Store.Size()
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Response) Size() (n int) {
	var l int
	_ = l
	if m.ReqType != nil {
		n += 1 + sovProtocol(uint64(*m.ReqType))
	}
	if m.ID != nil {
		n += 1 + sovProtocol(uint64(*m.ID))
	}
	if m.Nonce != nil {
		n += 1 + sovProtocol(uint64(*m.Nonce))
	}
	if m.Error != nil {
		l = len(*m.Error)
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.StoreVersionResp != nil {
		l = m.StoreVersionResp.Size()
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.FetchResp != nil {
		l = m.FetchResp.Size()
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProtocol(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProtocol(x uint64) (n int) {
	return sovProtocol(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReqType", wireType)
			}
			var v RequestType
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (RequestType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ReqType = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ID = &v
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Nonce = &v
			hasFields[0] |= uint64(0x00000004)
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StoreVersionResponse) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StoreVersionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreVersionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Version = &v
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FetchResponse) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FetchResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FetchResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Store", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Store == nil {
				m.Store = &brig_store.Store{}
			}
			if err := m.Store.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Response) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReqType", wireType)
			}
			var v RequestType
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (RequestType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ReqType = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ID = &v
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Nonce = &v
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[iNdEx:postIndex])
			m.Error = &s
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreVersionResp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StoreVersionResp == nil {
				m.StoreVersionResp = &StoreVersionResponse{}
			}
			if err := m.StoreVersionResp.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FetchResp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FetchResp == nil {
				m.FetchResp = &FetchResponse{}
			}
			if err := m.FetchResp.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProtocol(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthProtocol
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProtocol
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipProtocol(data[start:])
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
	ErrInvalidLengthProtocol = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtocol   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorProtocol = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x50, 0x51, 0x4b, 0xc2, 0x60,
	0x14, 0x75, 0xd3, 0xa5, 0xde, 0x31, 0x9b, 0x1f, 0x11, 0x43, 0x42, 0x64, 0x3d, 0x24, 0x11, 0xa3,
	0xfc, 0x03, 0x61, 0x39, 0x69, 0x20, 0x1a, 0x73, 0xf9, 0xd0, 0xcb, 0x28, 0xb9, 0x96, 0x10, 0xfb,
	0xe6, 0xdd, 0x2a, 0xfa, 0x27, 0xfd, 0xa4, 0x1e, 0x7b, 0xe9, 0x3d, 0xea, 0x8f, 0xf4, 0xed, 0x9b,
	0x48, 0x4a, 0x8f, 0x3d, 0x5c, 0xd8, 0xdd, 0x39, 0xe7, 0x9e, 0x73, 0x3e, 0xa8, 0xc5, 0xc4, 0x53,
	0x3e, 0xe5, 0x0f, 0x8e, 0xfc, 0x60, 0xc6, 0x2d, 0xcd, 0xef, 0x9c, 0x94, 0x6e, 0xa2, 0x64, 0x86,
	0xd4, 0xd0, 0x93, 0x94, 0x13, 0xe6, 0x98, 0xed, 0x43, 0xd9, 0xc7, 0xc5, 0x23, 0x26, 0x29, 0x3b,
	0x82, 0x0a, 0xe1, 0x22, 0x4c, 0x5f, 0x62, 0xb4, 0x94, 0x96, 0xda, 0xae, 0x75, 0x1a, 0xce, 0x9a,
	0xd2, 0x59, 0x32, 0x03, 0xc1, 0x60, 0x00, 0xaa, 0xd7, 0xb3, 0x54, 0xc1, 0x2b, 0x32, 0x03, 0xb4,
	0x88, 0x47, 0x53, 0xb4, 0x8a, 0xd9, 0x6a, 0x1f, 0xc0, 0xce, 0x38, 0xb3, 0x98, 0x20, 0x25, 0x73,
	0x1e, 0xf9, 0x98, 0xc4, 0x3c, 0x4a, 0x90, 0x6d, 0x43, 0xf9, 0x29, 0xff, 0x25, 0xef, 0x6b, 0xf6,
	0x09, 0x18, 0x7d, 0x4c, 0xa7, 0xf7, 0x2b, 0x46, 0x0b, 0x34, 0x19, 0x4e, 0xe2, 0x7a, 0xa7, 0x9e,
	0xfb, 0xe7, 0x79, 0xe5, 0x49, 0xfb, 0x43, 0x81, 0xca, 0x8a, 0xfe, 0x5f, 0x89, 0xb3, 0x15, 0x89,
	0x38, 0x59, 0xa5, 0x96, 0xd2, 0xae, 0xb2, 0x53, 0x60, 0xd2, 0x33, 0x5c, 0xc6, 0x0d, 0x49, 0x38,
	0x5a, 0x9a, 0xc0, 0xf4, 0xce, 0xfe, 0x86, 0xc3, 0x9f, 0x4d, 0x8f, 0x01, 0x66, 0x59, 0xb1, 0x5c,
	0xb8, 0x25, 0x85, 0x7b, 0x1b, 0xc2, 0xb5, 0xe6, 0x87, 0x1e, 0xe8, 0xbf, 0xb3, 0xea, 0x50, 0xf6,
	0x86, 0x93, 0xee, 0xc0, 0xeb, 0x99, 0x05, 0x56, 0x05, 0xad, 0xef, 0x06, 0xe7, 0x17, 0xa6, 0xc2,
	0xea, 0x60, 0x8c, 0x83, 0x91, 0xef, 0x86, 0x13, 0xd7, 0x1f, 0x7b, 0xa3, 0xa1, 0xa9, 0x8a, 0x57,
	0xd5, 0xaf, 0x2e, 0x7b, 0xdd, 0xc0, 0x0d, 0xfb, 0xde, 0xc0, 0x35, 0x8b, 0x67, 0xbb, 0x6f, 0x5f,
	0x4d, 0xe5, 0x5d, 0xcc, 0xa7, 0x98, 0xd7, 0xef, 0x66, 0xe1, 0xba, 0xf4, 0x3c, 0x27, 0xfc, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x31, 0xf2, 0x57, 0x8e, 0x17, 0x02, 0x00, 0x00,
}
