package shared

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
)

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

func (fns Functions) IsGoString(field pgs.Field) bool {

	if fns.GetGoFieldType(field) == GoStringField {
		return true
	}

	return false
}

func (fns Functions) IsGoDuration(field pgs.Field) bool {

	if fns.GetGoFieldType(field) == GoDurationField {
		return true
	}

	return false
}

func (fns Functions) IsGoTimestamp(field pgs.Field) bool {

	if fns.GetGoFieldType(field) == GoTimestampField {
		return true
	}

	return false
}

func (fns Functions) IsGoInteger32(field pgs.Field) bool {

	if fns.GetGoFieldType(field) == GoInteger32Field {
		return true
	}

	return false
}

func (fns Functions) IsGoUnsignedInteger32(field pgs.Field) bool {

	if fns.GetGoFieldType(field) == GoUnsignedInteger32Field {
		return true
	}

	return false
}

func (fns Functions) IsGoInteger64(field pgs.Field) bool {

	if fns.GetGoFieldType(field) == GoInteger64Field {
		return true
	}

	return false
}

func (fns Functions) IsGoUnsignedInteger64(field pgs.Field) bool {

	if fns.GetGoFieldType(field) == GoUnsignedInteger64Field {
		return true
	}

	return false
}

func (fns Functions) IsGoFloat32(field pgs.Field) bool {

	if fns.GetGoFieldType(field) == GoFloat32Field {
		return true
	}

	return false
}

func (fns Functions) IsGoFloat64(field pgs.Field) bool {

	if fns.GetGoFieldType(field) == GoFloat64Field {
		return true
	}

	return false
}

func (fns Functions) IsGoByte(field pgs.Field) bool {

	if fns.GetGoFieldType(field) == GoBytesField {
		return true
	}

	return false
}

func (fns Functions) IsGoBoolean(field pgs.Field) bool {

	if fns.GetGoFieldType(field) == GoBooleanField {
		return true
	}

	return false
}

func (fns Functions) IsGoMap(field pgs.Field) bool {

	value := fns.GetGoFieldType(field)

	if value == GoMapField {
		return true
	}

	return false
}

func (fns Functions) IsGoStruct(field pgs.Field) bool {

	value := fns.GetGoFieldType(field)

	if value == GoStructField {
		return true
	}

	return false
}

func (fns Functions) IsGoStructPB(field pgs.Field) bool {

	value := fns.GetGoFieldType(field)

	if value == GoStructPBField {
		return true
	}

	return false
}

func (fns Functions) GetStructType(field pgs.Field) string {

	if !fns.IsGoStruct(field) {
		panic("Field must be a struct to determine struct type")
	}

	return field.Type().Embed().Name().String()
}

func (fns Functions) IsGoSlice(field pgs.Field) bool {

	if fns.GetGoFieldType(field) == GoSliceField {
		return true
	}

	return false
}

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

func (fns Functions) GetGoMapKeyType(field pgs.Field) string {

	value := fns.GetGoFieldProtoType(fns.GetGoMapKeyFieldTypeElem(field).ProtoType())

	if value == GoStructField {
		return field.Type().Key().Embed().Name().String()
	} else if value == GoEnumField {
		return field.Type().Key().Enum().Name().String()
	}

	return value

}

func (fns Functions) GetGoMapKeyFieldTypeElem(field pgs.Field) pgs.FieldTypeElem {
	if !field.Type().IsMap() {
		panic("Field must be a map to determine map key")
	}

	return field.Type().Key()
}

func (fns Functions) GetGoMapValueFieldTypeElem(field pgs.Field) pgs.FieldTypeElem {
	if !field.Type().IsMap() {
		panic("Field must be a map to determine map value")
	}

	return field.Type().Element()
}

func (fns Functions) GetGoMapValueType(field pgs.Field) string {

	value := fns.GetGoFieldProtoType(fns.GetGoMapValueFieldTypeElem(field).ProtoType())

	if value == GoStructField {
		return field.Type().Element().Embed().Name().String()
	} else if value == GoEnumField {
		return field.Type().Element().Enum().Name().String()
	}

	return value

}

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
