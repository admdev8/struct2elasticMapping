/*
struct2elasticMapping enables developers to quickly build a
ElasticSearch Mapping from Go types.

Currently the Mapping is broken down to basic types (int,
byte, string, etc.) and support for more complex types like
time.Time (date), IP-Addresses, etc. is missing.

Using the special Tag "elastic" for structure members it
is possible to determine the type, analyzer and indexer
for the field.

	// Format: "Type,Analyzer,Index"
	type TestStruct struct {
		Field1 int64 `json:"title,omitempty" elastic:"date,standard,analyzed"`
		Subject string `json:"subject" elastic:",whitespace,not_analyzed"`
		Body string `json:"body" elastic:",german"`
	}

Currently it enables developers to quickly build a working
Mapping that can then be refined by the developer during
testing.

	package main

	import (
		"fmt"

		//nmap "github.com/lair-framework/go-nmap"

		s2m "github.com/marpie/struct2elasticMapping"
	)

	// Format: "Type,Analyzer,Index"
	type TestStruct struct {
		Field1  int64  `json:"title,omitempty" elastic:"date,standard,analyzed"`
		Subject string `json:"subject" elastic:",whitespace,not_analyzed"`
		Body    string `json:"body" elastic:",german"`
	}

	func main() {
		//name, mapping, err := s2m.Analyze(nmap.NmapRun{}, "json")
		name, mapping, err := s2m.Analyze(TestStruct{}, "json")
		if err != nil {
			panic(err.Error())
		}
		s, err := s2m.MappingAsJson(name, mapping)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(string(s))
	}

*/
package struct2elasticMapping
