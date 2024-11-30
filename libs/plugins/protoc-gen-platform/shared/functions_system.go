package shared

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	"sort"
	"strings"
)

func (fns Functions) DomainSystemName(file pgs.File) string {
	split := strings.Split(file.Name().String(), "/")
	return split[len(split)-3]
}

func (fns Functions) DomainSystemName2(file pgs.File) pgs.Name {
	split := strings.Split(file.Name().UpperCamelCase().String(), "/")
	return pgs.Name(split[len(split)-3])
}

func (fns Functions) ALlSystemNames(targets map[string]pgs.File) []pgs.Name {

	_systemNames := make([]string, 0)
	for k, _ := range targets {
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
