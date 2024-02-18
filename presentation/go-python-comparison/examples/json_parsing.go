package examples

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Age       int    `json:"age"`
}

func JsonExample() {
	// Parsed as struct
	person := Person{}
	jsonLiteral := `{"first-name":"John","last-name":"Doe","age":21}`
	json.Unmarshal([]byte(jsonLiteral), &person)
	fmt.Println("Unmarshaled person: ", person)

	// Parsed as map of any type
	person_any := make(map[string]interface{})
	json.Unmarshal([]byte(jsonLiteral), &person_any)
	fmt.Println("Unmarshaled person: ", person_any)
}
