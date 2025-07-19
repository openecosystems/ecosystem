package sdkv2betalib

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"

	optionv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/options/v2"
	specv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/spec/v2"
)

// globalMutex is a read-write mutex used to synchronize access to global resources.
// GlobalSystems is a globally accessible instance of the Systems type, initialized with default values.
var (
	globalMutex   sync.RWMutex
	GlobalSystems = new(Systems)
)

// System describes a system entity composed of a SpecSystem, a collection of Connectors, and an optional Dependency.
type System struct {
	specv2pb.SpecSystem
	Connectors []*Connector
	Dependency *Dependency
}

// Systems represents a collection of systems and their associated connectors.
// It maintains mappings by system name and connector path while tracking other configurations and dependencies.
// A file system and specification settings are also managed within this type.
type Systems struct {
	systemsByName    map[FullSystemName]*System
	connectorsByPath map[string][]protoreflect.FileDescriptor
	// numSystems       int
	fileSystem *FileSystem
	settings   *specv2pb.SpecSettings
}

// FullSystemName represents the name identifier for a system.
type FullSystemName string

// IsValid checks if the FullSystemName has a non-zero length and returns true if valid, otherwise false.
func (s FullSystemName) IsValid() bool {
	i := len(s)
	return i > 0
}

// RegisterSystems initializes the Systems instance with provided settings and processes system definitions.
// Returns an error if system registration fails.
func (s *Systems) RegisterSystems(provider BaseSpecConfigurationProvider) error {
	if s == GlobalSystems {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	if s.systemsByName == nil {
		s.systemsByName = map[FullSystemName]*System{}
		s.connectorsByPath = make(map[string][]protoreflect.FileDescriptor)
	}

	s.fileSystem = NewFileSystem()

	bytes, err := provider.GetConfigurationBytes()
	if err != nil {
		return nil
	}

	var settings specv2pb.SpecSettings
	err = proto.Unmarshal(bytes, &settings)
	if err != nil {
		return nil
	}

	s.settings = &settings
	if err := s.processSystems(); err != nil {
		return err
	}

	return nil
}

// GetSystems retrieves a map of all registered systems, keyed by their FullSystemName.
func (s *Systems) GetSystems() map[FullSystemName]*System {
	return s.systemsByName
}

// GetSystemByName retrieves a system by its name from the systemsByName map and returns it. Returns an error if not found.
func (s *Systems) GetSystemByName(systemName string) (*System, error) {
	if system, ok := s.systemsByName[FullSystemName(systemName)]; ok {
		return system, nil
	}

	return nil, fmt.Errorf("system '%s' not found", systemName)
}

// processSystems initializes, validates, and registers systems using the provided configuration and dependencies.
func (s *Systems) processSystems() error {
	for _, ss := range s.settings.Systems {
		// Validate system

		//if prev := s.systemsByName[FullSystemName(ss.Name)]; len(prev) > 0 {
		//	fmt.Printf("system %q is already registered\n", ss.Name)
		//}

		dependencyProvider := NewDynamicDependencyProvider(ss)
		dependency, err := dependencyProvider.registry.GetDependency()
		if err != nil {
			fmt.Println("dynamic dependency error: ", err)
			return err
		}

		system, err2 := s.loadSystemDescriptors(ss, dependency.data)
		if err2 != nil {
			fmt.Println("load system descriptors error: ")
			fmt.Println(err2)
			return nil
		}

		system.Dependency = dependency

		s.systemsByName[FullSystemName(system.Name)] = system
	}

	return nil
}

// loadSystemDescriptors processes a SpecSystem and its data to create and return a System object or an error if it fails.
func (s *Systems) loadSystemDescriptors(ss *specv2pb.SpecSystem, data []byte) (*System, error) {
	if err := os.Setenv("GOLANG_PROTOBUF_REGISTRATION_CONFLICT", "ignore"); err != nil {
		fmt.Println("Error setting environment variable:", err)
		return nil, fmt.Errorf("error setting environment variable: %v", err)
	}

	fdSet := &descriptorpb.FileDescriptorSet{}
	if err := proto.Unmarshal(data, fdSet); err != nil {
		return nil, fmt.Errorf("failed to unmarshal file descriptor set: %v", err)
	}

	files, err := protodesc.NewFiles(fdSet)
	if err != nil {
		return nil, err
	}

	var connectors []*Connector
	files.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		if !strings.HasPrefix(fd.Path(), "platform/") {
			return true
		}

		if err := protoregistry.GlobalFiles.RegisterFile(fd); err != nil {
			fmt.Println(fmt.Errorf("failed to register file descriptor: %v", err))
			return true
		}

		// fmt.Println("Registered file descriptor:", fd.Path())

		options := fd.Options().(*descriptorpb.FileOptions)
		if !proto.HasExtension(options, optionv2pb.E_Entity) {
			return true
		}

		customValue := proto.GetExtension(options, optionv2pb.E_Entity)
		_, ok := customValue.(*optionv2pb.EntityOptions)
		if !ok {
			fmt.Println("Type assertion failed")
			return true
		}

		_ = fd.Services()
		//for i := 0; i < sds.Len(); i++ {
		// connectors = append(connectors, connectorv2alphalib.NewConnectorWithSchema(sds.Get(i)))
		//}

		return true
	})

	return &System{
		SpecSystem: specv2pb.SpecSystem{
			Name:      ss.Name,
			Version:   ss.Version,
			Protocols: ss.Protocols,
		},
		Connectors: connectors,
	}, nil
}

// parseFields prints detailed information about the fields of a provided protoreflect.MessageDescriptor.
// It outputs the field name, type, cardinality (repeated, optional, or required), and any nested message or enum details.
//
//nolint:unused
func parseFields(messageDesc protoreflect.MessageDescriptor) {
	fields := messageDesc.Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)

		// Print the field name and type
		fmt.Printf("      Field Name: %s\n", field.Name())
		fmt.Printf("      Field Type: %s\n", field.Kind()) // Field type (e.g., int32, string)

		// Check if the field is repeated, optional, or required
		if field.Cardinality() == protoreflect.Repeated {
			fmt.Println("      Cardinality: Repeated")
		} else if field.HasOptionalKeyword() {
			fmt.Println("      Cardinality: Optional")
		} else {
			fmt.Println("      Cardinality: Required")
		}

		// For message types, print the nested message type
		if field.Kind() == protoreflect.MessageKind {
			fmt.Printf("      Nested Message Type: %s\n", field.Message().FullName())
		}

		// For enums, print the enum name and values
		if field.Kind() == protoreflect.EnumKind {
			enum := field.Enum()
			fmt.Printf("      Enum Type: %s\n", enum.FullName())

			// Print enum values
			values := enum.Values()
			for k := 0; k < values.Len(); k++ {
				fmt.Printf("        Enum Value: %s = %d\n", values.Get(k).Name(), values.Get(k).Number())
			}
		}
		fmt.Println() // Line break between fields
	}
}
