package examples

import "fmt"

func SlicesExample() []int {
	var slice []int
	_ = []int{1, 2, 3, 4, 5}
	_ = make([]int, 0, 10)
	slice = append(slice, 10, 11, 12)
	slice = append(slice, 20)
	slice = append(slice, 30)
	slice = append(slice, 40)
	slice = append(slice, 50)
	fmt.Printf("This is the full slice %v\n", slice)
	fmt.Printf("This is the subset 1 of slice %v\n", slice[2:])
	fmt.Printf("This is the subset 2 of slice %v\n", slice[2:])
	fmt.Printf("This is the subset 3 of slice %v\n", slice[2:])

	slice_additional := []int{1, 2, 3, 4, 5}
	slice = append(slice, slice_additional...)
	fmt.Println("This is the slice with additional elements: ", slice)
	return slice
}

func ArraysExample() [5]int {
	var array [5]int
	_ = [5]int{1, 2, 3, 4, 5}
	_ = make([]int, 0, 10)
	array[0] = 100
	array[1] = 200
	array[2] = 300
	array[3] = 400
	array[4] = 500
	fmt.Printf("This is the full array %v\n", array)
	fmt.Printf("This is the subset 1 of array %v\n", array[2:])
	fmt.Printf("This is the subset 2 of array %v\n", array[:3])
	fmt.Printf("This is the subset 3 of array %v\n", array[1:3])
	// fmt.Printf("This is the subset 3 of array %v\n", array[20:]) // Error: invalid argument: index 20 out of bounds [0:6]

	return array
}
