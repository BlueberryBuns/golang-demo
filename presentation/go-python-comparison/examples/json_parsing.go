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
	person := Person{}
	jsonLiteral := `{"first-name":"John","last-name":"Doe","age":21}`
	json.Unmarshal([]byte(jsonLiteral), &person)
	fmt.Println("Unmarshaled person: ", person)
}
