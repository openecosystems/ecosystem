package sdkv2alphalib

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
	"libs/protobuf/go/protobuf/gen/platform/options/v2"
	"libs/protobuf/go/protobuf/gen/platform/spec/v2"
)

var (
	globalMutex   sync.RWMutex
	GlobalSystems = new(Systems)
)

type System struct {
	specv2pb.SpecSystem
	Connectors []*Connector
	Dependency *Dependency
}

type Systems struct {
	systemsByName    map[FullSystemName]*System
	connectorsByPath map[string][]protoreflect.FileDescriptor
	numSystems       int
	fileSystem       *FileSystem
	settings         *specv2pb.SpecSettings
}

type FullSystemName string // e.g., "oeco.public.platform.Configuration"

func (s FullSystemName) IsValid() bool {
	i := len(s)
	if i < 0 {
		return false
	}
	return true
}

func (s *Systems) RegisterSystems(settingsProvider SpecSettingsProvider) error {
	if s == GlobalSystems {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	if s.systemsByName == nil {
		s.systemsByName = map[FullSystemName]*System{}
		s.connectorsByPath = make(map[string][]protoreflect.FileDescriptor)
	}

	s.fileSystem = NewFileSystem()
	s.settings = settingsProvider.GetSettings()
	if err := s.processSystems(); err != nil {
		return err
	}

	return nil
}

func (s *Systems) GetSystems() map[FullSystemName]*System {
	return s.systemsByName
}

func (s *Systems) GetSystemByName(systemName string) (*System, error) {
	if system, ok := s.systemsByName[FullSystemName(systemName)]; ok {
		return system, nil
	}

	return nil, fmt.Errorf("system '%s' not found", systemName)
}

func (s *Systems) processSystems() error {
	for _, ss := range s.settings.Systems2 {

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

// loadSystemDescriptors loads protobuf descriptors from a FileDescriptorSet and registers them
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

		sds := fd.Services()
		for i := 0; i < sds.Len(); i++ {
			// connectors = append(connectors, connectorv2alphalib.NewConnectorWithSchema(sds.Get(i)))
		}

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

// parseFields dynamically parses and displays information about each field in a MessageDescriptor
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
