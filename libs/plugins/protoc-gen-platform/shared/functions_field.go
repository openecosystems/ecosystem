package shared

import (
	"fmt"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

// StringField represents a string field type.
// IntegerField represents an integer field type.
// LongField represents a long field type.
// DoubleField represents a double field type.
// FloatField represents a float field type.
// BooleanField represents a boolean field type.
// ObjectField represents an object field type.
// MapField represents a map field type.
// MapEnumField represents a map enum field type.
// RepeatedField represents a repeated field type.
// EnumField represents an enumerated field type.
// ByteStringField represents a ByteString field type from the protobuf library.
// DurationField represents a Duration field type from the protobuf library.
// TimestampField represents a Timestamp field type from the protobuf library.
// LabelField represents a Label field type from the protobuf library.
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

// IsString determines if the provided field is of type string based on the field's type information.
func (fns Functions) IsString(field pgs.Field) bool {
	if fns.GetFieldType(field) == StringField {
		return true
	}

	return false
}

// IsDuration determines if the given field is of type `com.google.protobuf.Duration`. Returns true if it matches.
func (fns Functions) IsDuration(field pgs.Field) bool {
	if fns.GetFieldType(field) == DurationField {
		return true
	}

	return false
}

// IsTimestamp checks if the given field is of type `com.google.protobuf.Timestamp` and returns a boolean result.
func (fns Functions) IsTimestamp(field pgs.Field) bool {
	if fns.GetFieldType(field) == TimestampField {
		return true
	}

	return false
}

// IsLabel checks if the given field is a repeated field containing an embedded LabelDescriptor type.
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

// IsInteger checks if the given field has a type equivalent to IntegerField and returns true if it does, otherwise false.
func (fns Functions) IsInteger(field pgs.Field) bool {
	if fns.GetFieldType(field) == IntegerField {
		return true
	}

	return false
}

// IsLong checks if the provided field is of type LongField and returns true if it is, otherwise false.
func (fns Functions) IsLong(field pgs.Field) bool {
	if fns.GetFieldType(field) == LongField {
		return true
	}

	return false
}

// IsDouble checks if the given field has a type matching DoubleField and returns true if it does, false otherwise.
func (fns Functions) IsDouble(field pgs.Field) bool {
	if fns.GetFieldType(field) == DoubleField {
		return true
	}

	return false
}

// IsFloat checks if the specified field is of type Float and returns true if it is, otherwise returns false.
func (fns Functions) IsFloat(field pgs.Field) bool {
	if fns.GetFieldType(field) == FloatField {
		return true
	}

	return false
}

// IsByte checks if the provided field is of the type ByteStringField and returns true if it is, otherwise false.
func (fns Functions) IsByte(field pgs.Field) bool {
	if fns.GetFieldType(field) == ByteStringField {
		return true
	}

	return false
}

// IsBoolean determines if the provided protobuf field has a boolean data type. Returns true if the field is of type boolean.
func (fns Functions) IsBoolean(field pgs.Field) bool {
	if fns.GetFieldType(field) == BooleanField {
		return true
	}

	return false
}

// IsObject determines if a given field type should be treated as an object by checking against predefined field types.
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

// IsMap determines whether the provided field is either a MapField or MapEnumField by evaluating its type.
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

// IsMapEnum checks if the provided field is of the type MapEnumField and returns true if it matches, otherwise false.
func (fns Functions) IsMapEnum(field pgs.Field) bool {
	if fns.GetFieldType(field) == MapEnumField {
		return true
	}

	return false
}

// IsEnum checks if the provided field is of type Enum and returns true if it is, otherwise returns false.
func (fns Functions) IsEnum(field pgs.Field) bool {
	if fns.GetFieldType(field) == EnumField {
		return true
	}

	return false
}

// IsRepeated identifies if the given field has a "repeated" label by checking its type against RepeatedField constant.
func (fns Functions) IsRepeated(field pgs.Field) bool {
	if fns.GetFieldType(field) == RepeatedField {
		return true
	}

	return false
}

// GetObjectValueType determines and returns the type of the value in a field, handling embedded, map, enum, and repeated types.
func (fns Functions) GetObjectValueType(field pgs.Field) string {
	value := fns.GetFieldType(field)
	// value := fns.GetFieldProtoType(field.Type().ProtoType())

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

		// return field.Type().Element().Embed().Name().String()
	}

	return value
}

// GetListValueType determines the type of elements within a repeated field and returns it as a string representation.
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

// GetMapKeyType returns the proto type of the key for a given map field. Panics if the field is not a map.
func (fns Functions) GetMapKeyType(field pgs.Field) string {
	if !field.Type().IsMap() {
		panic("Field must be a map to determine map key")
	}
	return fns.GetFieldProtoType(field.Type().Key().ProtoType())
}

// GetMapValueType returns the type of values in the map field as a string. Panics if the provided field is not a map.
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

// GetRepeatedFieldType determines the type of a repeated field and returns its name if it is an enum or embedded message.
// If the field type is not repeated, it returns "not repeated".
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

// GetFieldType determines the type of a given field and returns its corresponding type as a string.
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

// Unwrap extracts the name of an embedded or enumerated field type or derives the field's proto type as a string.
func (fns Functions) Unwrap(ft pgs.FieldTypeElem) string {
	if ft.IsEmbed() {
		return ft.Embed().Name().String()
	}

	if ft.IsEnum() {
		return ft.Enum().Name().String()
	}

	return fns.GetFieldProtoType(ft.ProtoType())
}

// FullFieldType determines the full type string of a Protobuf field, including handling for repeated and map fields.
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

// GetFieldProtoType maps a ProtoType to its corresponding field type as a string representation.
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

// ShouldAddImport determines whether an import for a specific type should be added based on the provided message and type.
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

// HasField checks if the given message contains a field with the specified name and returns true if it exists, false otherwise.
func (fns Functions) HasField(msg pgs.Message, name string) bool {
	for _, f := range msg.Fields() {
		if f.Name().String() == name {
			return true
		}
	}
	return false
}

// IsFieldType checks if a field in the given message matches the specified name and type.
func (fns Functions) IsFieldType(msg pgs.Message, fieldName, fieldType string) bool {
	for _, f := range msg.Fields() {
		if f.Name().String() == fieldName && f.Descriptor().TypeName != nil {
			return *f.Descriptor().TypeName == fieldType
		}
	}
	return false
}
