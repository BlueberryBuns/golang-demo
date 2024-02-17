package examples

import (
	"fmt"
	"maps"
)

func MapsExample() {
	exampleMap := make(map[string]int)

	exampleMap["key_one"] = 1337
	exampleMap["key_two"] = 2115
	exampleMap["three"] = 1234

	fmt.Printf("Map is: %v\n", exampleMap)

	delete(exampleMap, "three")
	fmt.Printf("Map after delete is: %v\n", exampleMap)

	if _, ok := exampleMap["three"]; ok {
		fmt.Println("Key three exists")
	} else {
		fmt.Println("Key three does not exist")
	}

	map2 := map[string]int{"key_one": 1337, "key_two": 1337}
	fmt.Printf("Maps %v and %v are equal: %t", exampleMap, map2, maps.Equal(exampleMap, map2))
}
