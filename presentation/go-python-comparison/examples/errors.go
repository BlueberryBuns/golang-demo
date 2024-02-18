package examples

import (
	"fmt"
	"io"
	"os"
)

func ErrorsExampleNormal() {
	//File Exists
	file, openError := os.Open("examples/file.txt")
	if openError != nil {
		fmt.Printf("Error appeared during opening file.txt: %v\n", openError)
		panic(openError)
	}

	defer file.Close()
	content, readError := io.ReadAll(file)
	if readError != nil {
		fmt.Printf("Error appeared during reading file.txt: %v\n", readError)
		panic(readError)
	}

	fmt.Printf("File content: %s\n", content)

	// File does not exist
	_, err := os.Open("file1.txt")
	if err != nil {
		fmt.Printf("Error appeared during opening file.txt: %v\n", err)
		// panic(err)
	}
}

func Divide[Numeric int | float64](dividend, divisor Numeric) (float64, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("division by zero, divider is %v", divisor)
	}
	return float64(dividend) / float64(divisor), nil
}

func ErrorsExampleCustom() {
	// Division by zero
	if result, err := Divide(10, 0); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Division Result: %v\n", result)
	}

	// Normal division
	if result, err := Divide(10.0, 3.0); err != nil {
		fmt.Printf("Error: %v\n", err)
		panic(err)
	} else {
		fmt.Printf("Division Result: %v\n", result)
	}
}
