// Code generated by protoc-gen-gogo.
// source: file.proto
// DO NOT EDIT!

/*
	Package files is a generated protocol buffer package.

	It is generated from these files:
		file.proto

	It has these top-level messages:
		File
*/
package files

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type File struct {
	Filename string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Path     string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Filesize uint64   `protobuf:"varint,3,opt,name=filesize,proto3" json:"filesize,omitempty"`
	Mime     string   `protobuf:"bytes,4,opt,name=mime,proto3" json:"mime,omitempty"`
	Ctime    int64    `protobuf:"varint,5,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Mtime    int64    `protobuf:"varint,6,opt,name=mtime,proto3" json:"mtime,omitempty"`
	Uid      []byte   `protobuf:"bytes,7,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid      []byte   `protobuf:"bytes,8,opt,name=gid,proto3" json:"gid,omitempty"`
	Perm     FileMode `protobuf:"bytes,9,opt,name=perm,proto3,customtype=FileMode" json:"perm"`
}

func (m *File) Reset()                    { *m = File{} }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptorFile, []int{0} }

func init() {
	proto.RegisterType((*File)(nil), "files.File")
}
func (this *File) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*File)
	if !ok {
		that2, ok := that.(File)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Filename != that1.Filename {
		return false
	}
	if this.Path != that1.Path {
		return false
	}
	if this.Filesize != that1.Filesize {
		return false
	}
	if this.Mime != that1.Mime {
		return false
	}
	if this.Ctime != that1.Ctime {
		return false
	}
	if this.Mtime != that1.Mtime {
		return false
	}
	if !bytes.Equal(this.Uid, that1.Uid) {
		return false
	}
	if !bytes.Equal(this.Gid, that1.Gid) {
		return false
	}
	if !this.Perm.Equal(that1.Perm) {
		return false
	}
	return true
}
func (this *File) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 13)
	s = append(s, "&files.File{")
	s = append(s, "Filename: "+fmt.Sprintf("%#v", this.Filename)+",\n")
	s = append(s, "Path: "+fmt.Sprintf("%#v", this.Path)+",\n")
	s = append(s, "Filesize: "+fmt.Sprintf("%#v", this.Filesize)+",\n")
	s = append(s, "Mime: "+fmt.Sprintf("%#v", this.Mime)+",\n")
	s = append(s, "Ctime: "+fmt.Sprintf("%#v", this.Ctime)+",\n")
	s = append(s, "Mtime: "+fmt.Sprintf("%#v", this.Mtime)+",\n")
	s = append(s, "Uid: "+fmt.Sprintf("%#v", this.Uid)+",\n")
	s = append(s, "Gid: "+fmt.Sprintf("%#v", this.Gid)+",\n")
	s = append(s, "Perm: "+fmt.Sprintf("%#v", this.Perm)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringFile(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringFile(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func (m *File) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *File) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Filename) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintFile(data, i, uint64(len(m.Filename)))
		i += copy(data[i:], m.Filename)
	}
	if len(m.Path) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintFile(data, i, uint64(len(m.Path)))
		i += copy(data[i:], m.Path)
	}
	if m.Filesize != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintFile(data, i, uint64(m.Filesize))
	}
	if len(m.Mime) > 0 {
		data[i] = 0x22
		i++
		i = encodeVarintFile(data, i, uint64(len(m.Mime)))
		i += copy(data[i:], m.Mime)
	}
	if m.Ctime != 0 {
		data[i] = 0x28
		i++
		i = encodeVarintFile(data, i, uint64(m.Ctime))
	}
	if m.Mtime != 0 {
		data[i] = 0x30
		i++
		i = encodeVarintFile(data, i, uint64(m.Mtime))
	}
	if len(m.Uid) > 0 {
		data[i] = 0x3a
		i++
		i = encodeVarintFile(data, i, uint64(len(m.Uid)))
		i += copy(data[i:], m.Uid)
	}
	if len(m.Gid) > 0 {
		data[i] = 0x42
		i++
		i = encodeVarintFile(data, i, uint64(len(m.Gid)))
		i += copy(data[i:], m.Gid)
	}
	data[i] = 0x4a
	i++
	i = encodeVarintFile(data, i, uint64(m.Perm.Size()))
	n1, err := m.Perm.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func encodeFixed64File(data []byte, offset int, v uint64) int {
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
func encodeFixed32File(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintFile(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *File) Size() (n int) {
	var l int
	_ = l
	l = len(m.Filename)
	if l > 0 {
		n += 1 + l + sovFile(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovFile(uint64(l))
	}
	if m.Filesize != 0 {
		n += 1 + sovFile(uint64(m.Filesize))
	}
	l = len(m.Mime)
	if l > 0 {
		n += 1 + l + sovFile(uint64(l))
	}
	if m.Ctime != 0 {
		n += 1 + sovFile(uint64(m.Ctime))
	}
	if m.Mtime != 0 {
		n += 1 + sovFile(uint64(m.Mtime))
	}
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovFile(uint64(l))
	}
	l = len(m.Gid)
	if l > 0 {
		n += 1 + l + sovFile(uint64(l))
	}
	l = m.Perm.Size()
	n += 1 + l + sovFile(uint64(l))
	return n
}

func sovFile(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFile(x uint64) (n int) {
	return sovFile(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *File) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&File{`,
		`Filename:` + fmt.Sprintf("%v", this.Filename) + `,`,
		`Path:` + fmt.Sprintf("%v", this.Path) + `,`,
		`Filesize:` + fmt.Sprintf("%v", this.Filesize) + `,`,
		`Mime:` + fmt.Sprintf("%v", this.Mime) + `,`,
		`Ctime:` + fmt.Sprintf("%v", this.Ctime) + `,`,
		`Mtime:` + fmt.Sprintf("%v", this.Mtime) + `,`,
		`Uid:` + fmt.Sprintf("%v", this.Uid) + `,`,
		`Gid:` + fmt.Sprintf("%v", this.Gid) + `,`,
		`Perm:` + fmt.Sprintf("%v", this.Perm) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringFile(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *File) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFile
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
			return fmt.Errorf("proto: File: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: File: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filename", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
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
				return ErrInvalidLengthFile
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filename = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
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
				return ErrInvalidLengthFile
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filesize", wireType)
			}
			m.Filesize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Filesize |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
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
				return ErrInvalidLengthFile
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mime = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ctime", wireType)
			}
			m.Ctime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Ctime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mtime", wireType)
			}
			m.Mtime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Mtime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFile
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid = append(m.Uid[:0], data[iNdEx:postIndex]...)
			if m.Uid == nil {
				m.Uid = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gid", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFile
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gid = append(m.Gid[:0], data[iNdEx:postIndex]...)
			if m.Gid == nil {
				m.Gid = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Perm", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFile
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Perm.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFile(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFile
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
func skipFile(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFile
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
					return 0, ErrIntOverflowFile
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
					return 0, ErrIntOverflowFile
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
				return 0, ErrInvalidLengthFile
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFile
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
				next, err := skipFile(data[start:])
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
	ErrInvalidLengthFile = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFile   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorFile = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xcb, 0xcc, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0xb1, 0x8b, 0xa5, 0x74, 0xd3, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0xb2, 0x49,
	0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0x74, 0x29, 0xdd, 0x61, 0xe4, 0x62, 0x71, 0x03,
	0x6a, 0x14, 0x92, 0xe2, 0xe2, 0x00, 0x19, 0x90, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x19, 0x04, 0xe7, 0x0b, 0x09, 0x71, 0xb1, 0x14, 0x24, 0x96, 0x64, 0x48, 0x30, 0x81, 0xc5,
	0xc1, 0x6c, 0x98, 0xfa, 0xe2, 0xcc, 0xaa, 0x54, 0x09, 0x66, 0xa0, 0x38, 0x4b, 0x10, 0x9c, 0x0f,
	0x52, 0x9f, 0x9b, 0x09, 0x34, 0x87, 0x05, 0xa2, 0x1e, 0xc4, 0x16, 0x12, 0xe1, 0x62, 0x4d, 0x2e,
	0x01, 0x09, 0xb2, 0x02, 0x05, 0x99, 0x83, 0x20, 0x1c, 0x90, 0x68, 0x2e, 0x58, 0x94, 0x0d, 0x22,
	0x0a, 0xe6, 0x08, 0x09, 0x70, 0x31, 0x97, 0x66, 0xa6, 0x48, 0xb0, 0x03, 0xc5, 0x78, 0x82, 0x40,
	0x4c, 0x90, 0x48, 0x3a, 0x50, 0x84, 0x03, 0x22, 0x02, 0x64, 0x0a, 0xa9, 0x00, 0xdd, 0x94, 0x5a,
	0x94, 0x2b, 0xc1, 0x09, 0x12, 0x72, 0x12, 0x38, 0x71, 0x4f, 0x9e, 0xe1, 0xd6, 0x3d, 0x79, 0x0e,
	0x90, 0x5f, 0x7c, 0xf3, 0x53, 0x52, 0x83, 0xc0, 0xb2, 0x4e, 0x3a, 0x17, 0x1e, 0xca, 0x31, 0xdc,
	0x00, 0xe2, 0x0f, 0x0f, 0xe5, 0x18, 0x1b, 0x1e, 0xc9, 0x31, 0xae, 0x00, 0xe2, 0x13, 0x40, 0x7c,
	0x01, 0x88, 0x1f, 0x00, 0xf1, 0x8b, 0x47, 0x40, 0x39, 0x20, 0x3d, 0xe1, 0xb1, 0x1c, 0x43, 0x12,
	0x1b, 0x38, 0x4c, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x47, 0x7f, 0x64, 0x57, 0x01,
	0x00, 0x00,
}
