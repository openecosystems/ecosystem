package shared

import (
	"strings"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

// GetPackageVersion extracts the version component from the package name in the format "scope.system.version".
// It will panic if the package name does not adhere to the expected format.
// Returns the version as a string (e.g., "v2alpha").
func (fns Functions) GetPackageVersion(file pgs.File) string {
	p := file.Descriptor().GetPackage()
	s := strings.Split(p, ".")

	if len(s) != 3 {
		panic("Invalid package name. Must be in the format scope.system.version, for example: platform.solution.v2alpha")
	}

	return s[len(s)-1]
}

// GetPackageVersionName returns the version name of the package based on its descriptor in the format scope.system.version.
// It extracts the last segment of the package name. Panics if the package name format is invalid.
func (fns Functions) GetPackageVersionName(file pgs.File) pgs.Name {
	p := file.Descriptor().GetPackage()
	s := strings.Split(p, ".")

	if len(s) != 3 {
		panic("Invalid package name. Must be in the format scope.system.version, for example: platform.solution.v2alpha")
	}

	return pgs.Name(s[len(s)-1])
}

// GetAllVersions extracts and returns a unique list of version identifiers from the provided map of file targets.
func (fns Functions) GetAllVersions(targets map[string]pgs.File) []string {
	versions := make([]string, 0, len(targets))
	for t := range targets {
		p := targets[t].Descriptor().GetPackage()
		s := strings.Split(p, ".")

		if len(s) != 3 {
			panic("Invalid package name. Must be in the format scope.system.version, for example: platform.solution.v2alpha")
		}

		versions = append(versions, s[len(s)-1])
	}

	uniqueVersions := make([]string, 0, len(versions))
	m := make(map[string]bool)
	for _, val := range versions {
		if _, ok := m[val]; !ok {
			m[val] = true
			uniqueVersions = append(uniqueVersions, val)
		}
	}

	return uniqueVersions
}
