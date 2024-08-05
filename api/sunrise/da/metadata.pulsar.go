// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package da

import (
	_ "cosmossdk.io/api/amino"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_Metadata_3_list)(nil)

type _Metadata_3_list struct {
	list *[]string
}

func (x *_Metadata_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Metadata_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_Metadata_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_Metadata_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_Metadata_3_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message Metadata at list field ShardUris as it is not of Message kind"))
}

func (x *_Metadata_3_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_Metadata_3_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_Metadata_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Metadata                     protoreflect.MessageDescriptor
	fd_Metadata_shard_size          protoreflect.FieldDescriptor
	fd_Metadata_shard_count         protoreflect.FieldDescriptor
	fd_Metadata_shard_uris          protoreflect.FieldDescriptor
	fd_Metadata_recovered_data_hash protoreflect.FieldDescriptor
)

func init() {
	file_sunrise_da_metadata_proto_init()
	md_Metadata = File_sunrise_da_metadata_proto.Messages().ByName("Metadata")
	fd_Metadata_shard_size = md_Metadata.Fields().ByName("shard_size")
	fd_Metadata_shard_count = md_Metadata.Fields().ByName("shard_count")
	fd_Metadata_shard_uris = md_Metadata.Fields().ByName("shard_uris")
	fd_Metadata_recovered_data_hash = md_Metadata.Fields().ByName("recovered_data_hash")
}

var _ protoreflect.Message = (*fastReflection_Metadata)(nil)

type fastReflection_Metadata Metadata

func (x *Metadata) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Metadata)(x)
}

func (x *Metadata) slowProtoReflect() protoreflect.Message {
	mi := &file_sunrise_da_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Metadata_messageType fastReflection_Metadata_messageType
var _ protoreflect.MessageType = fastReflection_Metadata_messageType{}

type fastReflection_Metadata_messageType struct{}

func (x fastReflection_Metadata_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Metadata)(nil)
}
func (x fastReflection_Metadata_messageType) New() protoreflect.Message {
	return new(fastReflection_Metadata)
}
func (x fastReflection_Metadata_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Metadata
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Metadata) Descriptor() protoreflect.MessageDescriptor {
	return md_Metadata
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Metadata) Type() protoreflect.MessageType {
	return _fastReflection_Metadata_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Metadata) New() protoreflect.Message {
	return new(fastReflection_Metadata)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Metadata) Interface() protoreflect.ProtoMessage {
	return (*Metadata)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Metadata) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ShardSize != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ShardSize)
		if !f(fd_Metadata_shard_size, value) {
			return
		}
	}
	if x.ShardCount != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ShardCount)
		if !f(fd_Metadata_shard_count, value) {
			return
		}
	}
	if len(x.ShardUris) != 0 {
		value := protoreflect.ValueOfList(&_Metadata_3_list{list: &x.ShardUris})
		if !f(fd_Metadata_shard_uris, value) {
			return
		}
	}
	if len(x.RecoveredDataHash) != 0 {
		value := protoreflect.ValueOfBytes(x.RecoveredDataHash)
		if !f(fd_Metadata_recovered_data_hash, value) {
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
func (x *fastReflection_Metadata) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "sunrise.da.Metadata.shard_size":
		return x.ShardSize != uint64(0)
	case "sunrise.da.Metadata.shard_count":
		return x.ShardCount != uint64(0)
	case "sunrise.da.Metadata.shard_uris":
		return len(x.ShardUris) != 0
	case "sunrise.da.Metadata.recovered_data_hash":
		return len(x.RecoveredDataHash) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.da.Metadata"))
		}
		panic(fmt.Errorf("message sunrise.da.Metadata does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Metadata) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "sunrise.da.Metadata.shard_size":
		x.ShardSize = uint64(0)
	case "sunrise.da.Metadata.shard_count":
		x.ShardCount = uint64(0)
	case "sunrise.da.Metadata.shard_uris":
		x.ShardUris = nil
	case "sunrise.da.Metadata.recovered_data_hash":
		x.RecoveredDataHash = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.da.Metadata"))
		}
		panic(fmt.Errorf("message sunrise.da.Metadata does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Metadata) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "sunrise.da.Metadata.shard_size":
		value := x.ShardSize
		return protoreflect.ValueOfUint64(value)
	case "sunrise.da.Metadata.shard_count":
		value := x.ShardCount
		return protoreflect.ValueOfUint64(value)
	case "sunrise.da.Metadata.shard_uris":
		if len(x.ShardUris) == 0 {
			return protoreflect.ValueOfList(&_Metadata_3_list{})
		}
		listValue := &_Metadata_3_list{list: &x.ShardUris}
		return protoreflect.ValueOfList(listValue)
	case "sunrise.da.Metadata.recovered_data_hash":
		value := x.RecoveredDataHash
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.da.Metadata"))
		}
		panic(fmt.Errorf("message sunrise.da.Metadata does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Metadata) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "sunrise.da.Metadata.shard_size":
		x.ShardSize = value.Uint()
	case "sunrise.da.Metadata.shard_count":
		x.ShardCount = value.Uint()
	case "sunrise.da.Metadata.shard_uris":
		lv := value.List()
		clv := lv.(*_Metadata_3_list)
		x.ShardUris = *clv.list
	case "sunrise.da.Metadata.recovered_data_hash":
		x.RecoveredDataHash = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.da.Metadata"))
		}
		panic(fmt.Errorf("message sunrise.da.Metadata does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Metadata) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "sunrise.da.Metadata.shard_uris":
		if x.ShardUris == nil {
			x.ShardUris = []string{}
		}
		value := &_Metadata_3_list{list: &x.ShardUris}
		return protoreflect.ValueOfList(value)
	case "sunrise.da.Metadata.shard_size":
		panic(fmt.Errorf("field shard_size of message sunrise.da.Metadata is not mutable"))
	case "sunrise.da.Metadata.shard_count":
		panic(fmt.Errorf("field shard_count of message sunrise.da.Metadata is not mutable"))
	case "sunrise.da.Metadata.recovered_data_hash":
		panic(fmt.Errorf("field recovered_data_hash of message sunrise.da.Metadata is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.da.Metadata"))
		}
		panic(fmt.Errorf("message sunrise.da.Metadata does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Metadata) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "sunrise.da.Metadata.shard_size":
		return protoreflect.ValueOfUint64(uint64(0))
	case "sunrise.da.Metadata.shard_count":
		return protoreflect.ValueOfUint64(uint64(0))
	case "sunrise.da.Metadata.shard_uris":
		list := []string{}
		return protoreflect.ValueOfList(&_Metadata_3_list{list: &list})
	case "sunrise.da.Metadata.recovered_data_hash":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sunrise.da.Metadata"))
		}
		panic(fmt.Errorf("message sunrise.da.Metadata does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Metadata) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in sunrise.da.Metadata", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Metadata) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Metadata) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Metadata) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Metadata) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Metadata)
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
		if x.ShardSize != 0 {
			n += 1 + runtime.Sov(uint64(x.ShardSize))
		}
		if x.ShardCount != 0 {
			n += 1 + runtime.Sov(uint64(x.ShardCount))
		}
		if len(x.ShardUris) > 0 {
			for _, s := range x.ShardUris {
				l = len(s)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		l = len(x.RecoveredDataHash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
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
		x := input.Message.Interface().(*Metadata)
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
		if len(x.RecoveredDataHash) > 0 {
			i -= len(x.RecoveredDataHash)
			copy(dAtA[i:], x.RecoveredDataHash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RecoveredDataHash)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.ShardUris) > 0 {
			for iNdEx := len(x.ShardUris) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.ShardUris[iNdEx])
				copy(dAtA[i:], x.ShardUris[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ShardUris[iNdEx])))
				i--
				dAtA[i] = 0x1a
			}
		}
		if x.ShardCount != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ShardCount))
			i--
			dAtA[i] = 0x10
		}
		if x.ShardSize != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ShardSize))
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
		x := input.Message.Interface().(*Metadata)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Metadata: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ShardSize", wireType)
				}
				x.ShardSize = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ShardSize |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ShardCount", wireType)
				}
				x.ShardCount = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ShardCount |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ShardUris", wireType)
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
				x.ShardUris = append(x.ShardUris, string(dAtA[iNdEx:postIndex]))
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RecoveredDataHash", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.RecoveredDataHash = append(x.RecoveredDataHash[:0], dAtA[iNdEx:postIndex]...)
				if x.RecoveredDataHash == nil {
					x.RecoveredDataHash = []byte{}
				}
				iNdEx = postIndex
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
// source: sunrise/da/metadata.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShardSize         uint64   `protobuf:"varint,1,opt,name=shard_size,json=shardSize,proto3" json:"shard_size,omitempty"`
	ShardCount        uint64   `protobuf:"varint,2,opt,name=shard_count,json=shardCount,proto3" json:"shard_count,omitempty"`
	ShardUris         []string `protobuf:"bytes,3,rep,name=shard_uris,json=shardUris,proto3" json:"shard_uris,omitempty"`
	RecoveredDataHash []byte   `protobuf:"bytes,4,opt,name=recovered_data_hash,json=recoveredDataHash,proto3" json:"recovered_data_hash,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sunrise_da_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_sunrise_da_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetShardSize() uint64 {
	if x != nil {
		return x.ShardSize
	}
	return 0
}

func (x *Metadata) GetShardCount() uint64 {
	if x != nil {
		return x.ShardCount
	}
	return 0
}

func (x *Metadata) GetShardUris() []string {
	if x != nil {
		return x.ShardUris
	}
	return nil
}

func (x *Metadata) GetRecoveredDataHash() []byte {
	if x != nil {
		return x.RecoveredDataHash
	}
	return nil
}

var File_sunrise_da_metadata_proto protoreflect.FileDescriptor

var file_sunrise_da_metadata_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x64, 0x61, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x75, 0x6e,
	0x72, 0x69, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x1a, 0x11, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x61,
	0x6d, 0x69, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x08,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x68,
	0x61, 0x72, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x68,
	0x61, 0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x5f, 0x75, 0x72, 0x69, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68,
	0x61, 0x72, 0x64, 0x55, 0x72, 0x69, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x48, 0x61, 0x73, 0x68, 0x42, 0x85, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e,
	0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x64, 0x61, 0x42, 0x0d, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1b, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75,
	0x6e, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x64, 0x61, 0xa2, 0x02, 0x03, 0x53, 0x44, 0x58, 0xaa, 0x02,
	0x0a, 0x53, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x44, 0x61, 0xca, 0x02, 0x0a, 0x53, 0x75,
	0x6e, 0x72, 0x69, 0x73, 0x65, 0x5c, 0x44, 0x61, 0xe2, 0x02, 0x16, 0x53, 0x75, 0x6e, 0x72, 0x69,
	0x73, 0x65, 0x5c, 0x44, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0b, 0x53, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x3a, 0x3a, 0x44, 0x61, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sunrise_da_metadata_proto_rawDescOnce sync.Once
	file_sunrise_da_metadata_proto_rawDescData = file_sunrise_da_metadata_proto_rawDesc
)

func file_sunrise_da_metadata_proto_rawDescGZIP() []byte {
	file_sunrise_da_metadata_proto_rawDescOnce.Do(func() {
		file_sunrise_da_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_sunrise_da_metadata_proto_rawDescData)
	})
	return file_sunrise_da_metadata_proto_rawDescData
}

var file_sunrise_da_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sunrise_da_metadata_proto_goTypes = []interface{}{
	(*Metadata)(nil), // 0: sunrise.da.Metadata
}
var file_sunrise_da_metadata_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sunrise_da_metadata_proto_init() }
func file_sunrise_da_metadata_proto_init() {
	if File_sunrise_da_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sunrise_da_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
			RawDescriptor: file_sunrise_da_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sunrise_da_metadata_proto_goTypes,
		DependencyIndexes: file_sunrise_da_metadata_proto_depIdxs,
		MessageInfos:      file_sunrise_da_metadata_proto_msgTypes,
	}.Build()
	File_sunrise_da_metadata_proto = out.File
	file_sunrise_da_metadata_proto_rawDesc = nil
	file_sunrise_da_metadata_proto_goTypes = nil
	file_sunrise_da_metadata_proto_depIdxs = nil
}
