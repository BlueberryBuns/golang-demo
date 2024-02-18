package main

import (
	"fmt"

	alt "github.com/BlueberryBuns/presentation/go-python-comparison/alternative"
	"github.com/BlueberryBuns/presentation/go-python-comparison/examples"
)

const (
	RED = iota
	BLUE
	GREEN
	PURPLE
	YELLOW
)

func main() {
	// Variables section
	var privateVariable string = "Hello World private"
	fmt.Printf("This is the message %s\n", privateVariable)

	// int zmienna = 123 // private
	// public int zmienna_publiczna = 123
	// private int zmienna_prywatną = 123

	var PublicVariable string = "Hello World public"
	fmt.Println("This is the message from public: ", PublicVariable)

	anotherVariable := "Hello World another"
	fmt.Println("This is the message from another: ", anotherVariable)

	otherDeclaration := "Hello World other"
	fmt.Println("This is the message from other declaration: ", otherDeclaration)

	const constantVariable = "Hello World constant"
	fmt.Println("This is the message from constant: ", constantVariable)

	// var notUsedVariable = "Hello World not used" // Error: notUsedVariable declared and not used
	_ = "Hello World not used"

	alt.FooPublic()

	// Slices and Arrays section (Lists)
	slice := examples.SlicesExample()
	array := examples.ArraysExample()
	slice = append(slice, array[:]...)
	fmt.Println("This is the slice with additional elements from array: ", slice)

	// Maps section (Dictionaries/Hashmaps)
	examples.MapsExample()

	// Structs section (Classes, Interfaces)
	examples.StructsExample()
	examples.GenericExample()

	// Errors section (Exceptions)
	examples.ErrorsExampleCustom()
	examples.ErrorsExampleNormal()

	// Concurrency
	examples.SeqExample()
	examples.ChannelsExample()
	examples.ChannelsBufferingExample()
	examples.WaitGroupExample()

	// JSON
	examples.JsonExample()
}
