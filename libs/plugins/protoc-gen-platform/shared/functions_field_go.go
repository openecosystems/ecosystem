package shared

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
)

// GoStringField represents a string type in Go.
// GoInteger32Field represents a 32-bit integer type in Go.
// GoInteger64Field represents a 64-bit integer type in Go.
// GoUnsignedInteger32Field represents a 32-bit unsigned integer type in Go.
// GoUnsignedInteger64Field represents a 64-bit unsigned integer type in Go.
// GoFloat64Field represents a 64-bit floating-point type in Go.
// GoFloat32Field represents a 32-bit floating-point type in Go.
// GoBooleanField represents a boolean type in Go.
// GoMapField represents a map type in Go.
// GoStructField represents a struct type in Go.
// GoStructPBField represents a Protobuf struct type in Go.
// GoSliceField represents a slice type in Go.
// GoBytesField represents a byte slice type in Go.
// GoEnumField represents an enum type in Go.
// GoDurationField represents a Protobuf duration type in Go.
// GoTimestampField represents a Protobuf timestamp type in Go.
const (
	GoStringField            = "string"
	GoInteger32Field         = "int32"
	GoInteger64Field         = "int64"
	GoUnsignedInteger32Field = "uint32"
	GoUnsignedInteger64Field = "uint64"
	GoFloat64Field           = "float64"
	GoFloat32Field           = "float32"
	GoBooleanField           = "bool"
	GoMapField               = "map"
	GoStructField            = "struct"
	GoStructPBField          = "structpb"
	GoSliceField             = "[]"
	GoBytesField             = "[]byte"
	GoEnumField              = "enum"
	GoDurationField          = "durationpb.Duration"
	GoTimestampField         = "timestamppb.Timestamp"
)

// IsGoString checks if the provided pgs.Field has a Go type equivalent to GoStringField and returns true if it does.
func (fns Functions) IsGoString(field pgs.Field) bool {
	if fns.GetGoFieldType(field) == GoStringField {
		return true
	}

	return false
}

// IsGoDuration returns true if the provided field has a Go type of "durationpb.Duration", otherwise returns false.
func (fns Functions) IsGoDuration(field pgs.Field) bool {
	if fns.GetGoFieldType(field) == GoDurationField {
		return true
	}

	return false
}

// IsGoTimestamp determines if the provided field is of type GoTimestampField (`timestamppb.Timestamp`).
func (fns Functions) IsGoTimestamp(field pgs.Field) bool {
	if fns.GetGoFieldType(field) == GoTimestampField {
		return true
	}

	return false
}

// IsGoInteger32 checks if the given field has a Go type of "int32". It returns true if the type matches, otherwise false.
func (fns Functions) IsGoInteger32(field pgs.Field) bool {
	if fns.GetGoFieldType(field) == GoInteger32Field {
		return true
	}

	return false
}

// IsGoUnsignedInteger32 checks if the field has the Go type "uint32" and returns true if so, otherwise false.
func (fns Functions) IsGoUnsignedInteger32(field pgs.Field) bool {
	if fns.GetGoFieldType(field) == GoUnsignedInteger32Field {
		return true
	}

	return false
}

// IsGoInteger64 determines if the given field's type is "int64" in Go. Returns true if it matches, otherwise false.
func (fns Functions) IsGoInteger64(field pgs.Field) bool {
	if fns.GetGoFieldType(field) == GoInteger64Field {
		return true
	}

	return false
}

// IsGoUnsignedInteger64 checks if the given field is of type GoUnsignedInteger64Field and returns true if it matches.
func (fns Functions) IsGoUnsignedInteger64(field pgs.Field) bool {
	if fns.GetGoFieldType(field) == GoUnsignedInteger64Field {
		return true
	}

	return false
}

// IsGoFloat32 determines whether the specified field has a Go type of float32 and returns true if it matches.
func (fns Functions) IsGoFloat32(field pgs.Field) bool {
	if fns.GetGoFieldType(field) == GoFloat32Field {
		return true
	}

	return false
}

// IsGoFloat64 determines if the provided field has a Go type of "float64".
func (fns Functions) IsGoFloat64(field pgs.Field) bool {
	if fns.GetGoFieldType(field) == GoFloat64Field {
		return true
	}

	return false
}

// IsGoByte checks if the given field's Go type is a byte slice ([]byte). Returns true if it matches, otherwise false.
func (fns Functions) IsGoByte(field pgs.Field) bool {
	if fns.GetGoFieldType(field) == GoBytesField {
		return true
	}

	return false
}

// IsGoBoolean checks if the given field has a Go type of "bool".
func (fns Functions) IsGoBoolean(field pgs.Field) bool {
	if fns.GetGoFieldType(field) == GoBooleanField {
		return true
	}

	return false
}

// IsGoMap checks if the provided field is of type GoMapField and returns true if it is, otherwise returns false.
func (fns Functions) IsGoMap(field pgs.Field) bool {
	value := fns.GetGoFieldType(field)

	if value == GoMapField {
		return true
	}

	return false
}

// IsGoStruct determines if the provided field is of type Go struct based on its Go field type. Returns true if it is.
func (fns Functions) IsGoStruct(field pgs.Field) bool {
	value := fns.GetGoFieldType(field)

	if value == GoStructField {
		return true
	}

	return false
}

// IsGoStructPB checks if the given field is of type "structpb" in the generated Go code and returns true if matched.
func (fns Functions) IsGoStructPB(field pgs.Field) bool {
	value := fns.GetGoFieldType(field)

	if value == GoStructPBField {
		return true
	}

	return false
}

// GetStructType returns the name of the struct type of the given field. It panics if the field is not a struct.
func (fns Functions) GetStructType(field pgs.Field) string {
	if !fns.IsGoStruct(field) {
		panic("Field must be a struct to determine struct type")
	}

	return field.Type().Embed().Name().String()
}

// IsGoSlice checks if the given field in the protobuf structure is represented as a Go slice type.
func (fns Functions) IsGoSlice(field pgs.Field) bool {
	if fns.GetGoFieldType(field) == GoSliceField {
		return true
	}

	return false
}

// GetGoSliceValueType determines the element type of a repeated field and returns its Go representation as a string.
// If the field is not repeated, the method panics.
// It handles struct and enum types explicitly by returning their names.
func (fns Functions) GetGoSliceValueType(field pgs.Field) string {
	if !field.Type().IsRepeated() {
		panic("Field must be a list to determine list value")
	}

	value := fns.GetGoFieldProtoType(field.Type().Element().ProtoType())

	if value == GoStructField {
		return field.Type().Element().Embed().Name().String()
	} else if value == GoEnumField {
		return field.Type().Element().Enum().Name().String()
	}

	return value
}

// GetGoMapKeyType returns the Go type of the key for a protobuf map field based on its definition and key type.
func (fns Functions) GetGoMapKeyType(field pgs.Field) string {
	value := fns.GetGoFieldProtoType(fns.GetGoMapKeyFieldTypeElem(field).ProtoType())

	if value == GoStructField {
		return field.Type().Key().Embed().Name().String()
	} else if value == GoEnumField {
		return field.Type().Key().Enum().Name().String()
	}

	return value
}

// GetGoMapKeyFieldTypeElem retrieves the key type of a map field.
// Panics if the provided field is not of map type.
func (fns Functions) GetGoMapKeyFieldTypeElem(field pgs.Field) pgs.FieldTypeElem {
	if !field.Type().IsMap() {
		panic("Field must be a map to determine map key")
	}

	return field.Type().Key()
}

// GetGoMapValueFieldTypeElem returns the element type of a map field's value. Panics if the field is not a map.
func (fns Functions) GetGoMapValueFieldTypeElem(field pgs.Field) pgs.FieldTypeElem {
	if !field.Type().IsMap() {
		panic("Field must be a map to determine map value")
	}

	return field.Type().Element()
}

// GetGoMapValueType determines the Go type of a map's value from a given protobuf field.
// It returns the type as a string, including struct or enum names when applicable.
func (fns Functions) GetGoMapValueType(field pgs.Field) string {
	value := fns.GetGoFieldProtoType(fns.GetGoMapValueFieldTypeElem(field).ProtoType())

	if value == GoStructField {
		return field.Type().Element().Embed().Name().String()
	} else if value == GoEnumField {
		return field.Type().Element().Enum().Name().String()
	}

	return value
}

// GetGoFieldType determines and returns the Go field type for the provided Protobuf field based on its characteristics.
func (fns Functions) GetGoFieldType(field pgs.Field) string {
	if field.Type().IsMap() {
		return GoMapField
	}

	if field.Type().IsRepeated() {
		return GoSliceField
	}

	if field.Type().IsEmbed() {
		if embed := field.Type().Embed(); embed.IsWellKnown() {
			switch embed.WellKnownType() {
			case pgs.AnyWKT:
				return GoStringField
			case pgs.DurationWKT:
				return GoDurationField
			case pgs.TimestampWKT:
				return GoTimestampField
			case pgs.Int32ValueWKT:
				return GoInteger32Field
			case pgs.UInt32ValueWKT:
				return GoUnsignedInteger32Field
			case pgs.Int64ValueWKT:
				return GoInteger64Field
			case pgs.UInt64ValueWKT:
				return GoUnsignedInteger64Field
			case pgs.DoubleValueWKT:
				return GoFloat64Field
			case pgs.FloatValueWKT:
				return GoFloat32Field
			case pgs.BoolValueWKT:
				return GoBooleanField
			}
		}
	}

	if field.Descriptor().TypeName != nil && *field.Descriptor().TypeName == ".google.protobuf.Struct" {
		return GoStructPBField
	}
	return fns.GetGoFieldProtoType(field.Type().ProtoType())
}

// GetGoFieldProtoType maps a protobuf field type to its corresponding Go field type as a string representation.
func (fns Functions) GetGoFieldProtoType(t pgs.ProtoType) string {
	switch t {
	case pgs.Int32T, pgs.SInt32, pgs.SFixed32:
		return GoInteger32Field
	case pgs.UInt32T, pgs.Fixed32T:
		return GoUnsignedInteger32Field
	case pgs.Int64T, pgs.SInt64, pgs.SFixed64:
		return GoInteger64Field
	case pgs.UInt64T, pgs.Fixed64T:
		return GoUnsignedInteger64Field
	case pgs.DoubleT:
		return GoFloat64Field
	case pgs.FloatT:
		return GoFloat32Field
	case pgs.BoolT:
		return GoBooleanField
	case pgs.StringT:
		return GoStringField
	case pgs.BytesT:
		return GoBytesField
	case pgs.EnumT:
		return GoEnumField
	case pgs.MessageT:
		return GoStructField
	default:
		panic("unsupported proto type for go")
	}
}
