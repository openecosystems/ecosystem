package shared

import (
	"sort"
	"strings"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

// DomainSystemName extracts and returns the system-specific domain name from the provided file's path.
func (fns Functions) DomainSystemName(file pgs.File) string {
	split := strings.Split(file.Name().String(), "/")
	return split[len(split)-3]
}

// DomainSystemName2 extracts the domain system name from the given file's name by splitting its UpperCamelCase path.
// It returns the third-to-last element as the domain system name.
func (fns Functions) DomainSystemName2(file pgs.File) pgs.Name {
	split := strings.Split(file.Name().UpperCamelCase().String(), "/")
	return pgs.Name(split[len(split)-3])
}

// ALlSystemNames generates a sorted, unique list of system names from the provided targets map.
func (fns Functions) ALlSystemNames(targets map[string]pgs.File) []pgs.Name {
	_systemNames := make([]string, 0)
	for k := range targets {
		systemName := fns.DomainSystemName2(targets[k])
		_systemNames = append(_systemNames, systemName.LowerCamelCase().String())
	}

	systemNames := make([]string, 0, len(_systemNames))
	m1 := make(map[string]bool)

	for _, val := range _systemNames {
		if _, ok := m1[val]; !ok {
			m1[val] = true
			systemNames = append(systemNames, val)
		}
	}
	sort.Strings(systemNames)

	sNames := make([]pgs.Name, 0, len(systemNames))
	for _, s := range systemNames {
		sNames = append(sNames, pgs.Name(s))
	}

	return sNames
}
