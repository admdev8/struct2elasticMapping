package struct2elasticMapping

import (
	"reflect"
	"strings"
)

// Analyze converts the supplied interface to a Mapping object.
// `nameFromStructTag` can be used to retrieve the name of
// struct members from a tag - e.g. in the following snippet
// 	type t struct {
//		F1 string `json:"fieldName1,omitempty"`
// 	}
// 	Analyze(t{}, "json")
// The function would use "fieldName1" as the name for the
// member "F1". If no Tag is present the members name is used.
func Analyze(i interface{}, nameFromStructTag string) (name string, mapping *Mapping, err error) {
	t := reflect.TypeOf(i)
	_, p, err := analyzeType(t, nameFromStructTag, make(map[string]bool))
	if err != nil {
		return
	}
	name = t.Name()
	mapping = &Mapping{
		Properties: p.Properties,
	}
	return
}

// analyzeType converts the type from Go to ElasticSearch. If a
// struct is found it recursively tries to determine convert everything
// to a ES 'object'.
func analyzeType(t reflect.Type, nameFromStructTag string, wt map[string]bool) (name string, p *Property, err error) {
	name = t.Name()
	p = &Property{}
	switch t.Kind() {
	case reflect.Array,
		reflect.Map,
		reflect.Slice:
		return analyzeType(t.Elem(), nameFromStructTag, wt)

	case reflect.Struct:
		p.Type = FieldTypeObject
		p.Properties = make(map[string]Property)
		if wt[t.Name()] {
			// Already walked this path, don't fall into loop...
			return
		}
		wt[t.Name()] = true
		for i := 0; i < t.NumField(); i++ {
			// Do sth. with the field
			f := t.Field(i)
			var sub *Property
			_, sub, err = analyzeType(f.Type, nameFromStructTag, wt)
			if err != nil {
				return
			}
			fName := f.Name
			if len(nameFromStructTag) > 0 {
				structTag := strings.Split(f.Tag.Get(nameFromStructTag), ",")[0]
				if len(structTag) > 0 {
					fName = structTag
				}
			}
			p.Properties[fName] = *sub
		}

	case reflect.Bool:
		p.Type = FieldTypeBoolean

	case reflect.Int,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint64,
		reflect.Uint32,
		reflect.Uintptr:
		p.Type = FieldTypeLong

	case reflect.Int32,
		reflect.Uint16:
		p.Type = FieldTypeInteger

	case reflect.Int16:
		p.Type = FieldTypeShort

	case reflect.Int8:
		p.Type = FieldTypeByte

	case reflect.Uint8:
		p.Type = FieldTypeShort

	case reflect.Float64:
		p.Type = FieldTypeDouble

	case reflect.Float32:
		p.Type = FieldTypeFloat

	case reflect.String:
		p.Type = FieldTypeString

	default:
		p.Type = FieldTypeBinary
	}

	return
}
