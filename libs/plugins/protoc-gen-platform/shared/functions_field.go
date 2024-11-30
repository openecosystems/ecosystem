package shared

import (
	"fmt"
	pgs "github.com/lyft/protoc-gen-star/v2"
)

const (
	StringField     = "String"
	IntegerField    = "Integer"
	LongField       = "Long"
	DoubleField     = "Double"
	FloatField      = "Float"
	BooleanField    = "Boolean"
	ObjectField     = "Object"
	MapField        = "Map"
	MapEnumField    = "Map.Enum"
	RepeatedField   = "Repeated"
	EnumField       = "Enum"
	ByteStringField = "com.google.protobuf.ByteString"
	DurationField   = "com.google.protobuf.Duration"
	TimestampField  = "com.google.protobuf.Timestamp"
	LabelField      = "com.google.protobuf.Label"
)

func (fns Functions) IsString(field pgs.Field) bool {

	if fns.GetFieldType(field) == StringField {
		return true
	}

	return false
}

func (fns Functions) IsDuration(field pgs.Field) bool {

	if fns.GetFieldType(field) == DurationField {
		return true
	}

	return false
}

func (fns Functions) IsTimestamp(field pgs.Field) bool {

	if fns.GetFieldType(field) == TimestampField {
		return true
	}

	return false
}

func (fns Functions) IsLabel(field pgs.Field) bool {

	if fns.IsRepeated(field) {
		if field.Type().Element().IsEmbed() {
			if field.Type().Element().Embed().Name().String() == "LabelDescriptor" {
				return true
			}
		}
	}

	return false
}

func (fns Functions) IsInteger(field pgs.Field) bool {

	if fns.GetFieldType(field) == IntegerField {
		return true
	}

	return false
}

func (fns Functions) IsLong(field pgs.Field) bool {

	if fns.GetFieldType(field) == LongField {
		return true
	}

	return false
}

func (fns Functions) IsDouble(field pgs.Field) bool {

	if fns.GetFieldType(field) == DoubleField {
		return true
	}

	return false
}

func (fns Functions) IsFloat(field pgs.Field) bool {

	if fns.GetFieldType(field) == FloatField {
		return true
	}

	return false
}

func (fns Functions) IsByte(field pgs.Field) bool {

	if fns.GetFieldType(field) == ByteStringField {
		return true
	}

	return false
}

func (fns Functions) IsBoolean(field pgs.Field) bool {

	if fns.GetFieldType(field) == BooleanField {
		return true
	}

	return false
}

func (fns Functions) IsObject(field pgs.Field) bool {

	value := fns.GetFieldType(field)

	if value == ObjectField {
		return true
	}

	if value == IntegerField {
		return true
	}

	if value == LongField {
		return true
	}

	if value == DoubleField {
		return true
	}

	if value == FloatField {
		return true
	}

	return false
}

func (fns Functions) IsMap(field pgs.Field) bool {

	value := fns.GetFieldType(field)

	if value == MapField {
		return true
	}

	if value == MapEnumField {
		return true
	}

	return false
}

func (fns Functions) IsMapEnum(field pgs.Field) bool {

	if fns.GetFieldType(field) == MapEnumField {
		return true
	}

	return false
}

func (fns Functions) IsEnum(field pgs.Field) bool {

	if fns.GetFieldType(field) == EnumField {
		return true
	}

	return false
}

func (fns Functions) IsRepeated(field pgs.Field) bool {

	if fns.GetFieldType(field) == RepeatedField {
		return true
	}

	return false
}

func (fns Functions) GetObjectValueType(field pgs.Field) string {

	value := fns.GetFieldType(field)
	//value := fns.GetFieldProtoType(field.Type().ProtoType())

	if value == ObjectField {

		if field.Type().IsEmbed() {
			if field.Type().Embed().IsMapEntry() {
				return "Map.Entry"
			}

			return field.Type().Embed().Name().String()
		}

		if field.Type().IsEnum() {
			return field.Type().Enum().Name().String()
		}

		if field.Type().IsRepeated() {
			if field.Type().Element().IsEmbed() {
				return field.Type().Element().Embed().Name().String()

			}
		}

		if field.Type().Element().IsEnum() {
			return field.Type().Element().Enum().Name().String()
		}

		//return field.Type().Element().Embed().Name().String()
	}

	return value

}

func (fns Functions) GetListValueType(field pgs.Field) string {

	if !field.Type().IsRepeated() {
		panic("Field must be a list to determine list value")
	}

	value := fns.GetFieldProtoType(field.Type().Element().ProtoType())

	if value == ObjectField {
		if field.Type().Element().IsEmbed() {
			return field.Type().Element().Embed().Name().String()
		}

		if field.Type().Element().IsEnum() {
			return field.Type().Element().Enum().Name().String()
		}
	}

	return value

}

func (fns Functions) GetMapKeyType(field pgs.Field) string {

	if !field.Type().IsMap() {
		panic("Field must be a map to determine map key")
	}
	return fns.GetFieldProtoType(field.Type().Key().ProtoType())

}

func (fns Functions) GetMapValueType(field pgs.Field) string {

	if !field.Type().IsMap() {
		panic("Field must be a map to determine map value")
	}

	value := fns.GetFieldProtoType(field.Type().Element().ProtoType())

	if value == ObjectField {
		if field.Type().Element().IsEmbed() {
			return field.Type().Element().Embed().Name().String()
		}

		if field.Type().Element().IsEnum() {
			return field.Type().Element().Enum().Name().String()
		}
	}

	return value

}

func (fns Functions) GetRepeatedFieldType(fieldType pgs.FieldType) string {
	if fieldType.IsRepeated() {
		if fieldType.Element().IsEnum() {
			return fieldType.Element().Enum().Descriptor().GetName()
		}

		if fieldType.Element().IsEmbed() {
			return fieldType.Element().Embed().Descriptor().GetName()
		}

		return fns.GetFieldProtoType(fieldType.ProtoType())
	}

	return "not repeated"
}

func (fns Functions) GetFieldType(field pgs.Field) string {

	if field.Type().IsMap() {
		if field.Type().Element().ProtoType() == pgs.EnumT {
			return MapEnumField
		}
		return MapField
	}

	if field.Type().IsRepeated() {
		return RepeatedField
	}

	if field.Type().IsEnum() {
		return EnumField
	}

	if field.Type().IsEmbed() {
		if embed := field.Type().Embed(); embed.IsWellKnown() {
			switch embed.WellKnownType() {
			case pgs.AnyWKT:
				return StringField
			case pgs.DurationWKT:
				return DurationField
			case pgs.TimestampWKT:
				return TimestampField
			case pgs.Int32ValueWKT, pgs.UInt32ValueWKT:
				return IntegerField
			case pgs.Int64ValueWKT, pgs.UInt64ValueWKT:
				return LongField
			case pgs.DoubleValueWKT:
				return DoubleField
			case pgs.FloatValueWKT:
				return FloatField
			case pgs.BoolValueWKT:
				return BooleanField
			}
		}
	}

	return fns.GetFieldProtoType(field.Type().ProtoType())
}

func (fns Functions) Unwrap(ft pgs.FieldTypeElem) string {
	if ft.IsEmbed() {
		return ft.Embed().Name().String()
	}

	if ft.IsEnum() {
		return ft.Enum().Name().String()
	}

	return fns.GetFieldProtoType(ft.ProtoType())
}

func (fns Functions) FullFieldType(field pgs.Field) string {

	value := fns.GetFieldType(field)

	if value == RepeatedField {
		return RepeatedField + " " + fns.GetRepeatedFieldType(field.Type())
	}

	if value == MapField || value == MapEnumField {
		kt := field.Type().Key()
		vt := field.Type().Element()

		k := fns.Unwrap(kt)
		v := fns.Unwrap(vt)

		return fmt.Sprintf("map<%s,%s>", k, v)
	}

	return value
}

func (fns Functions) GetFieldProtoType(t pgs.ProtoType) string {

	switch t {
	case pgs.Int32T, pgs.UInt32T, pgs.SInt32, pgs.Fixed32T, pgs.SFixed32:
		return IntegerField
	case pgs.Int64T, pgs.UInt64T, pgs.SInt64, pgs.Fixed64T, pgs.SFixed64:
		return LongField
	case pgs.DoubleT:
		return DoubleField
	case pgs.FloatT:
		return FloatField
	case pgs.BoolT:
		return BooleanField
	case pgs.StringT:
		return StringField
	case pgs.BytesT:
		return ByteStringField
	default:
		return ObjectField
	}
}

func (fns Functions) ShouldAddImport(msg pgs.Message, t string) bool {

	//for _, f := range msg.Fields() {
	//
	//	if f.Type().IsEmbed() {
	//		m.Log("Embed Name: " + f.Type().Embed().Name().String())
	//	}
	//
	//	if f.Type().IsEnum() {
	//		m.Log("Enum Name: " + f.Type().Enum().Name().String())
	//	}
	//}

	return false

}

func (fns Functions) HasField(msg pgs.Message, name string) bool {
	for _, f := range msg.Fields() {
		if f.Name().String() == name {
			return true
		}
	}
	return false
}

func (fns Functions) IsFieldType(msg pgs.Message, fieldName, fieldType string) bool {
	for _, f := range msg.Fields() {
		if f.Name().String() == fieldName && f.Descriptor().TypeName != nil {
			return *f.Descriptor().TypeName == fieldType
		}
	}
	return false
}
