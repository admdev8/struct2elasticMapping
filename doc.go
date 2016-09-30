/*
struct2elasticMapping enables developers to quickly build a
ElasticSearch Mapping from Go types.

Currently the Mapping is broken down to basic types (int,
byte, string, etc.) and support for more complex types like
time.Time (date), IP-Addresses, etc. is missing.

Currently it enables developers to quickly build a working
Mapping that can then be refined by the developer during
testing.

	package main

	import (
		"fmt"

		nmap "github.com/lair-framework/go-nmap"

		s2m "github.com/marpie/struct2elasticMapping"
	)

	func main() {
		name, mapping, err := s2m.Analyze(nmap.NmapRun{}, "json")
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
