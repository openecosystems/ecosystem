package shared

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	options "libs/protobuf/go/protobuf/gen/platform/options/v2"
)

func (fns Functions) EntityOptions(file pgs.File) options.EntityOptions {

	var entity options.EntityOptions

	_, err := file.Extension(options.E_Entity, &entity)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	return entity
}

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

func (fns Functions) ParentEntity(method pgs.Method) pgs.Message {

	file := method.File()
	return fns.Entity(file)

}

func (fns Functions) EntityName(msg pgs.Message) pgs.Name {

	var entity options.EntityOptions

	_, err := msg.File().Extension(options.E_Entity, &entity)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}
	return pgs.Name(entity.Entity)
}

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

func (fns Functions) EntityNamePlural(msg pgs.Message) pgs.Name {

	var entity options.EntityOptions

	_, err := msg.File().Extension(options.E_Entity, &entity)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}
	return pgs.Name(entity.EntityPlural)
}

func (fns Functions) EntityNamespace(msg pgs.Message) pgs.Name {
	var entity options.EntityOptions

	_, err := msg.File().Extension(options.E_Entity, &entity)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}
	return pgs.Name(entity.Namespace)
}

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

func (fns Functions) IsWorkspaceEntityHierarchy(msg pgs.Message) bool {
	var entity options.EntityOptions

	_, err := msg.File().Extension(options.E_Entity, &entity)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	if entity.Hierarchy == options.EntityHierarchy_ENTITY_HIERARCHY_WORKSPACE {
		return true
	}

	return false
}

func (fns Functions) EntityType(msg pgs.Message) pgs.Name {

	return fns.EntityTypeFromFile(msg.File())
}

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

func (fns Functions) EntityNameFromField(field pgs.Field) pgs.Name {
	return fns.EntityName(field.Message())
}

func (fns Functions) BinName(field pgs.Field) pgs.Name {
	return pgs.Name(truncate(field.Name().LowerCamelCase().String(), 14))
}

func truncate(s string, i int) string {
	runes := []rune(s)
	if len(runes) > i {
		return string(runes[:i])
	}
	return s
}
