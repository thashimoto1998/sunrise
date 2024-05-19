// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package liquiditypool

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Position            protoreflect.MessageDescriptor
	fd_Position_id         protoreflect.FieldDescriptor
	fd_Position_sender     protoreflect.FieldDescriptor
	fd_Position_lower_tick protoreflect.FieldDescriptor
	fd_Position_upper_tick protoreflect.FieldDescriptor
)

func init() {
	file_sunrise_liquiditypool_position_proto_init()
	md_Position = File_sunrise_liquiditypool_position_proto.Messages().ByName("Position")
	fd_Position_id = md_Position.Fields().ByName("id")
	fd_Position_sender = md_Position.Fields().ByName("sender")
	fd_Position_lower_tick = md_Position.Fields().ByName("lower_tick")
	fd_Position_upper_tick = md_Position.Fields().ByName("upper_tick")
}

var _ protoreflect.Message = (*fastReflection_Position)(nil)

type fastReflection_Position Position

func (x *Position) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Position)(x)
}

func (x *Position) slowProtoReflect() protoreflect.Message {
	mi := &file_sunrise_liquiditypool_position_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Position_messageType fastReflection_Position_messageType
var _ protoreflect.MessageType = fastReflection_Position_messageType{}

type fastReflection_Position_messageType struct{}

func (x fastReflection_Position_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Position)(nil)
}
func (x fastReflection_Position_messageType) New() protoreflect.Message {
	return new(fastReflection_Position)
}
func (x fastReflection_Position_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Position
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Position) Descriptor() protoreflect.MessageDescriptor {
	return md_Position
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Position) Type() protoreflect.MessageType {
	return _fastReflection_Position_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Position) New() protoreflect.Message {
	return new(fastReflection_Position)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Position) Interface() protoreflect.ProtoMessage {
	return (*Position)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Position) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_Position_id, value) {
			return
		}
	}
	if x.Sender != "" {
		value := protoreflect.ValueOfString(x.Sender)
		if !f(fd_Position_sender, value) {
			return
		}
	}
	if x.LowerTick != int64(0) {
		value := protoreflect.ValueOfInt64(x.LowerTick)
		if !f(fd_Position_lower_tick, value) {
			return
		}
	}
	if x.UpperTick != int64(0) {
		value := protoreflect.ValueOfInt64(x.UpperTick)
		if !f(fd_Position_upper_tick, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Position) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "sunrise.liquiditypool.Position.id":
		return x.Id != uint64(0)
	case "sunrise.liquiditypool.Position.sender":
		return x.Sender != ""
	case "sunrise.liquiditypool.Position.lower_tick":
		return x.LowerTick != int64(0)
	case "sunrise.liquiditypool.Position.upper_tick":
		return x.UpperTick != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.liquiditypool.Position"))
		}
		panic(fmt.Errorf("message sunrise.liquiditypool.Position does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Position) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "sunrise.liquiditypool.Position.id":
		x.Id = uint64(0)
	case "sunrise.liquiditypool.Position.sender":
		x.Sender = ""
	case "sunrise.liquiditypool.Position.lower_tick":
		x.LowerTick = int64(0)
	case "sunrise.liquiditypool.Position.upper_tick":
		x.UpperTick = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.liquiditypool.Position"))
		}
		panic(fmt.Errorf("message sunrise.liquiditypool.Position does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Position) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "sunrise.liquiditypool.Position.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "sunrise.liquiditypool.Position.sender":
		value := x.Sender
		return protoreflect.ValueOfString(value)
	case "sunrise.liquiditypool.Position.lower_tick":
		value := x.LowerTick
		return protoreflect.ValueOfInt64(value)
	case "sunrise.liquiditypool.Position.upper_tick":
		value := x.UpperTick
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.liquiditypool.Position"))
		}
		panic(fmt.Errorf("message sunrise.liquiditypool.Position does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Position) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "sunrise.liquiditypool.Position.id":
		x.Id = value.Uint()
	case "sunrise.liquiditypool.Position.sender":
		x.Sender = value.Interface().(string)
	case "sunrise.liquiditypool.Position.lower_tick":
		x.LowerTick = value.Int()
	case "sunrise.liquiditypool.Position.upper_tick":
		x.UpperTick = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.liquiditypool.Position"))
		}
		panic(fmt.Errorf("message sunrise.liquiditypool.Position does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Position) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "sunrise.liquiditypool.Position.id":
		panic(fmt.Errorf("field id of message sunrise.liquiditypool.Position is not mutable"))
	case "sunrise.liquiditypool.Position.sender":
		panic(fmt.Errorf("field sender of message sunrise.liquiditypool.Position is not mutable"))
	case "sunrise.liquiditypool.Position.lower_tick":
		panic(fmt.Errorf("field lower_tick of message sunrise.liquiditypool.Position is not mutable"))
	case "sunrise.liquiditypool.Position.upper_tick":
		panic(fmt.Errorf("field upper_tick of message sunrise.liquiditypool.Position is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.liquiditypool.Position"))
		}
		panic(fmt.Errorf("message sunrise.liquiditypool.Position does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Position) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "sunrise.liquiditypool.Position.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "sunrise.liquiditypool.Position.sender":
		return protoreflect.ValueOfString("")
	case "sunrise.liquiditypool.Position.lower_tick":
		return protoreflect.ValueOfInt64(int64(0))
	case "sunrise.liquiditypool.Position.upper_tick":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.liquiditypool.Position"))
		}
		panic(fmt.Errorf("message sunrise.liquiditypool.Position does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Position) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in sunrise.liquiditypool.Position", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Position) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Position) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Position) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Position) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Position)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		l = len(x.Sender)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.LowerTick != 0 {
			n += 1 + runtime.Sov(uint64(x.LowerTick))
		}
		if x.UpperTick != 0 {
			n += 1 + runtime.Sov(uint64(x.UpperTick))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Position)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.UpperTick != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.UpperTick))
			i--
			dAtA[i] = 0x20
		}
		if x.LowerTick != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.LowerTick))
			i--
			dAtA[i] = 0x18
		}
		if len(x.Sender) > 0 {
			i -= len(x.Sender)
			copy(dAtA[i:], x.Sender)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Sender)))
			i--
			dAtA[i] = 0x12
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Position)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Position: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Position: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Sender = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LowerTick", wireType)
				}
				x.LowerTick = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.LowerTick |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field UpperTick", wireType)
				}
				x.UpperTick = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.UpperTick |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: sunrise/liquiditypool/position.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender    string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	LowerTick int64  `protobuf:"varint,3,opt,name=lower_tick,json=lowerTick,proto3" json:"lower_tick,omitempty"`
	UpperTick int64  `protobuf:"varint,4,opt,name=upper_tick,json=upperTick,proto3" json:"upper_tick,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sunrise_liquiditypool_position_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_sunrise_liquiditypool_position_proto_rawDescGZIP(), []int{0}
}

func (x *Position) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Position) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Position) GetLowerTick() int64 {
	if x != nil {
		return x.LowerTick
	}
	return 0
}

func (x *Position) GetUpperTick() int64 {
	if x != nil {
		return x.UpperTick
	}
	return 0
}

var File_sunrise_liquiditypool_position_proto protoreflect.FileDescriptor

var file_sunrise_liquiditypool_position_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2e,
	0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x22, 0x70, 0x0a,
	0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x70, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x42,
	0xc7, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2e,
	0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x42, 0x0d, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0xa2, 0x02, 0x03, 0x53, 0x4c, 0x58, 0xaa, 0x02, 0x15, 0x53,
	0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x70, 0x6f, 0x6f, 0x6c, 0xca, 0x02, 0x15, 0x53, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x5c, 0x4c,
	0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0xe2, 0x02, 0x21, 0x53,
	0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x5c, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x70, 0x6f, 0x6f, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x16, 0x53, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x3a, 0x3a, 0x4c, 0x69, 0x71, 0x75,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_sunrise_liquiditypool_position_proto_rawDescOnce sync.Once
	file_sunrise_liquiditypool_position_proto_rawDescData = file_sunrise_liquiditypool_position_proto_rawDesc
)

func file_sunrise_liquiditypool_position_proto_rawDescGZIP() []byte {
	file_sunrise_liquiditypool_position_proto_rawDescOnce.Do(func() {
		file_sunrise_liquiditypool_position_proto_rawDescData = protoimpl.X.CompressGZIP(file_sunrise_liquiditypool_position_proto_rawDescData)
	})
	return file_sunrise_liquiditypool_position_proto_rawDescData
}

var file_sunrise_liquiditypool_position_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sunrise_liquiditypool_position_proto_goTypes = []interface{}{
	(*Position)(nil), // 0: sunrise.liquiditypool.Position
}
var file_sunrise_liquiditypool_position_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sunrise_liquiditypool_position_proto_init() }
func file_sunrise_liquiditypool_position_proto_init() {
	if File_sunrise_liquiditypool_position_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sunrise_liquiditypool_position_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sunrise_liquiditypool_position_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sunrise_liquiditypool_position_proto_goTypes,
		DependencyIndexes: file_sunrise_liquiditypool_position_proto_depIdxs,
		MessageInfos:      file_sunrise_liquiditypool_position_proto_msgTypes,
	}.Build()
	File_sunrise_liquiditypool_position_proto = out.File
	file_sunrise_liquiditypool_position_proto_rawDesc = nil
	file_sunrise_liquiditypool_position_proto_goTypes = nil
	file_sunrise_liquiditypool_position_proto_depIdxs = nil
}
