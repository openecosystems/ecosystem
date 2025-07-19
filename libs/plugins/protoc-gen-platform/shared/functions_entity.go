package shared

import (
	options "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/options/v2"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

// EntityOptions extracts and returns the EntityOptions extension from the provided proto file descriptor.
func (fns Functions) EntityOptions(file pgs.File) options.EntityOptions {
	var entity options.EntityOptions

	_, err := file.Extension(options.E_Entity, &entity)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	return entity
}

// Entity retrieves a specific message from the provided file based on the EntityOptions extension configuration.
// If the extension exists and matches a message name, that message is returned. Otherwise, it returns nil.
func (fns Functions) Entity(file pgs.File) pgs.Message {
	var entity options.EntityOptions

	_, err := file.Extension(options.E_Entity, &entity)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	if entity.Entity != "" {
		for _, msg := range file.AllMessages() {
			if msg.Name().String() == pgs.Name(entity.Entity).UpperCamelCase().String() {
				return msg
			}
		}
	}

	return nil
}

// ParentEntity retrieves the message type of the parent entity associated with the provided RPC method.
func (fns Functions) ParentEntity(method pgs.Method) pgs.Message {
	file := method.File()
	return fns.Entity(file)
}

// EntityName retrieves the entity name by extracting the EntityOptions extension from the provided protobuf message.
func (fns Functions) EntityName(msg pgs.Message) pgs.Name {
	var entity options.EntityOptions

	_, err := msg.File().Extension(options.E_Entity, &entity)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}
	return pgs.Name(entity.Entity)
}

// EntityKeyName extracts the primary key field name from the provided protobuf message.
// It panics if no key field is marked in the entity definition.
func (fns Functions) EntityKeyName(msg pgs.Message) pgs.Name {
	key := "none"

	for _, field := range msg.Fields() {
		var entityField options.EntityFieldOptions

		_, err := field.Extension(options.E_EntityField, &entityField)
		if err != nil {
			panic(err.Error() + "unable to read extension from proto")
		}

		if entityField.Key {
			key = field.Name().String()
		}
	}

	if key == "none" {
		panic("entities require a key. No key was found in the entity")
	}

	return pgs.Name(key)
}

// EntityNamePlural retrieves the plural name of an entity from the protobuf message using the EntityOptions extension.
func (fns Functions) EntityNamePlural(msg pgs.Message) pgs.Name {
	var entity options.EntityOptions

	_, err := msg.File().Extension(options.E_Entity, &entity)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}
	return pgs.Name(entity.EntityPlural)
}

// EntityNamespace retrieves the namespace for an entity defined in the protobuf message using custom extension options.
func (fns Functions) EntityNamespace(msg pgs.Message) pgs.Name {
	var entity options.EntityOptions

	_, err := msg.File().Extension(options.E_Entity, &entity)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}
	return pgs.Name(entity.Namespace)
}

// IsGlobalEntityNamespace determines if the namespace of the given message's EntityOptions is set to "global".
func (fns Functions) IsGlobalEntityNamespace(msg pgs.Message) bool {
	var entity options.EntityOptions

	_, err := msg.File().Extension(options.E_Entity, &entity)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	if entity.Namespace == "global" {
		return true
	}

	return false
}

// IsPlatformEntityHierarchy checks if the given message has a platform-level entity hierarchy defined in its options.
func (fns Functions) IsPlatformEntityHierarchy(msg pgs.Message) bool {
	var entity options.EntityOptions

	_, err := msg.File().Extension(options.E_Entity, &entity)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	if entity.Hierarchy == options.EntityHierarchy_ENTITY_HIERARCHY_PLATFORM {
		return true
	}

	return false
}

// IsOrganizationEntityHierarchy checks if a protobuf message represents an organization entity hierarchy based on its options.
func (fns Functions) IsOrganizationEntityHierarchy(msg pgs.Message) bool {
	var entity options.EntityOptions

	_, err := msg.File().Extension(options.E_Entity, &entity)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	if entity.Hierarchy == options.EntityHierarchy_ENTITY_HIERARCHY_ORGANIZATION {
		return true
	}

	return false
}

// IsWorkspaceEntityHierarchy checks if a given protobuf message represents an entity with workspace hierarchy.
func (fns Functions) IsWorkspaceEntityHierarchy(msg pgs.Message) bool {
	var entity options.EntityOptions

	_, err := msg.File().Extension(options.E_Entity, &entity)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	if entity.Hierarchy == options.EntityHierarchy_ENTITY_HIERARCHY_ECOSYSTEM {
		return true
	}

	return false
}

// EntityType resolves the type name of an entity from a given Protocol Buffers message by using its associated file context.
func (fns Functions) EntityType(msg pgs.Message) pgs.Name {
	return fns.EntityTypeFromFile(msg.File())
}

// EntityTypeFromFile retrieves the entity type for the given protobuf file based on the defined EntityOptions extension.
func (fns Functions) EntityTypeFromFile(file pgs.File) pgs.Name {
	var entity options.EntityOptions

	_, err := file.Extension(options.E_Entity, &entity)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}
	switch entity.Type {
	case options.EntityType_ENTITY_TYPE_AEROSPIKE:
		return pgs.Name("aerospike")
	case options.EntityType_ENTITY_TYPE_DGRAPH:
		return pgs.Name("dgraph")
	case options.EntityType_ENTITY_TYPE_MONGODB:
		return pgs.Name("mongo")
	case options.EntityType_ENTITY_TYPE_BIGQUERY:
		return pgs.Name("big_query")
	case options.EntityType_ENTITY_TYPE_REDIS:
		return pgs.Name("redis")
	case options.EntityType_ENTITY_TYPE_ROCKSDB:
		return pgs.Name("rocksdb")
	case options.EntityType_ENTITY_TYPE_COUCHBASE:
		return pgs.Name("couchbase")
	case options.EntityType_ENTITY_TYPE_UNSPECIFIED:
		fallthrough
	default:
		return pgs.Name("none")
	}
}

// IsEntity checks if the provided protobuf message matches the entity defined in its file's custom options.
// Returns true if the message name corresponds to the configured entity, otherwise false.
// Panics if the entity extension cannot be read from the file.
func (fns Functions) IsEntity(msg pgs.Message) bool {
	var entity options.EntityOptions

	_, err := msg.File().Extension(options.E_Entity, &entity)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	if msg.Name().String() == pgs.Name(entity.Entity).UpperCamelCase().String() {
		return true
	}

	return false
}

// GetEntityType extracts the entity type from the given file's options and returns its string representation.
func (fns Functions) GetEntityType(file pgs.File) string {
	var entity options.EntityOptions

	_, err := file.Extension(options.E_Entity, &entity)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}
	switch entity.Type {
	case options.EntityType_ENTITY_TYPE_UNSPECIFIED:
		return "none"
	case options.EntityType_ENTITY_TYPE_AEROSPIKE:
		return "aerospike"
	case options.EntityType_ENTITY_TYPE_DGRAPH:
		return "dgraph"
	case options.EntityType_ENTITY_TYPE_MONGODB:
		return "mongo"
	case options.EntityType_ENTITY_TYPE_BIGQUERY:
		return "bigquery"
	case options.EntityType_ENTITY_TYPE_COUCHBASE:
		return "couchbase"
	default:
		return "none"
	}
}

// EntityNameFromField retrieves the name of the entity associated with the given protobuf field.
func (fns Functions) EntityNameFromField(field pgs.Field) pgs.Name {
	return fns.EntityName(field.Message())
}

// BinName generates a truncated name for a field, converting it to lower camel case and limiting it to 14 characters.
func (fns Functions) BinName(field pgs.Field) pgs.Name {
	return pgs.Name(truncate(field.Name().LowerCamelCase().String(), 14))
}

// truncate truncates the input string s to the specified length i if it exceeds i.
// If the length of s is less than or equal to i, it returns s unchanged.
func truncate(s string, i int) string {
	runes := []rune(s)
	if len(runes) > i {
		return string(runes[:i])
	}
	return s
}
