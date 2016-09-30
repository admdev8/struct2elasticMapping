package struct2elasticMapping

import (
	"encoding/json"
)

func MappingAsJson(name string, m *Mapping) ([]byte, error) {
	v := make(map[string]Mapping)
	v[name] = *m
	return json.MarshalIndent(v, "", "\t")
}
