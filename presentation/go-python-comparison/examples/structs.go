package examples

import "fmt"

type Car struct {
	Brand string
	Model string
	Year  int
	Engine
	da DriveAxle
}

func (c Car) GetMileage() int {
	return c.Mileage
}

func (c *Car) ReduceMileagePointer(kilometers int) {
	c.Mileage = c.Mileage - kilometers
	fmt.Printf("[Reference] Reduced mileage by %d, new mileage is %d\n", kilometers, c.Mileage)
}

func (c Car) ReduceMileageCopy(kilometers int) Car {
	c.Mileage = c.Mileage - kilometers
	fmt.Printf("[Copy] Reduced mileage by %d, new mileage is %d\n", kilometers, c.Mileage)
	fmt.Println("Returning new car with reduced mileage")
	return c
}

func ReduceMileageViaReference(c *Car, kilometers int) {
	c.Mileage = c.Mileage - kilometers
	fmt.Printf("[Reference] Reduced mileage by %d, new mileage is %d\n", kilometers, c.Mileage)
}

func (c Car) drive() string {
	return fmt.Sprintf("Holy sh*t it's %s %s driving!", c.Brand, c.Model)
}
func (c Car) isIvecoDaily() bool {
	return true
}

type DriveAxle struct {
	Condition float64
}

func (da DriveAxle) doesDriveAxleMakeNoise() {
	if da.Condition < 0.1 {
		fmt.Println("Panie, Most Ci napierdala!")
		return
	}
	fmt.Println("Drive Axle is fine")
}

type Engine struct {
	Capacity int
	Power    int
	Mileage  int
}

func (e Engine) IsPowerful() {
	if e.Power > 300 {
		fmt.Println("Engine is powerful")
		return
	}
	fmt.Println("Engine is weak")
}

type coalCapable interface {
	drive() string
	isIvecoDaily() bool
}

func MoveCoal(c coalCapable) {
	if c.isIvecoDaily() {
		fmt.Printf("Moving coal with... %s, O Boze O K ", c.drive())
	} else {
		fmt.Println("This vehicle is not capable of moving coal")
	}
}

func StructsExample() {
	carVariable := Car{
		Brand: "Iveco",
		Model: "Daily",
		Year:  2003,
		Engine: Engine{
			Capacity: 2300,
			Power:    120,
			Mileage:  200000,
		},
		da: DriveAxle{
			Condition: 0.001,
		},
	}

	carVariable.IsPowerful()
	MoveCoal(carVariable)

	fmt.Printf("Car is: %v\n", carVariable)
	fmt.Printf("[Attribute] Car mileage is: %d\n", carVariable.Mileage)
	fmt.Printf("[Attribute] Car Drive Axle condition is: %f\n", carVariable.da.Condition)
	carVariable.da.doesDriveAxleMakeNoise()

	fmt.Printf("[ReduceMileageCopy] Car address before: %p\n", &carVariable)
	carVariable = carVariable.ReduceMileageCopy(10000)
	fmt.Printf("[ReduceMileageCopy] Car address after: %p\n", &carVariable)

	fmt.Printf("[ReduceMileagePointer] Car address before: %p\n", &carVariable)
	carVariable.ReduceMileagePointer(5)
	fmt.Printf("[ReduceMileagePointer] Car address after: %p\n", &carVariable)

	fmt.Printf("[ReduceMileageViaReference] Car address before: %p\n", &carVariable)
	ReduceMileageViaReference(&carVariable, 5)
	fmt.Printf("[ReduceMileageViaReference] Car address after: %p\n", &carVariable)
}

type MyCustomList[T any] struct {
	elements []T
}

func (l *MyCustomList[T]) Append(element T) {
	l.elements = append(l.elements, element)
}

func (l *MyCustomList[T]) GetElements() []T {
	return l.elements
}

func GenericExample() {
	floatList := MyCustomList[float64]{}
	floatList.Append(1.0)
	floatList.Append(2.0)
	// floatList.Append("abc") // Error: cannot use "xd" (untyped string constant) as float64 value in argument to floatList.Append
	fmt.Printf("Float list is: %v\n", floatList.GetElements())

	carList := MyCustomList[Car]{}
	carList.Append(Car{Brand: "Iveco", Model: "Daily", Year: 2003, Engine: Engine{Capacity: 2300, Power: 120, Mileage: 200000}, da: DriveAxle{Condition: 0.001}})
	carList.Append(Car{Brand: "Porsche", Model: "911", Year: 2023, Engine: Engine{Capacity: 3500, Power: 600, Mileage: 100}, da: DriveAxle{Condition: 0.99}})
	fmt.Printf("Car list is: %v\n", carList.GetElements())
}
