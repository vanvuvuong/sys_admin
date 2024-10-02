package main

// struct: typed collections of fields
type exampleStructType struct {
	Id   int32
	Name string
}

// method:
type circle struct {
	Radius float32
}

func (c circle) area() float32 {
	return 3.14 * c.Radius * c.Radius
}

func AdvancedDefine() {
	// pointer: NOT NEED FOR NOW

	// struct implement
	exampleStruct := new(exampleStructType)
	p("Example struct", exampleStruct)

	// method implement
	c := circle{1.2}
	pf("Area of %f circle is %f\n", c.Radius, c.area())

	// interface: NOT NEED FOR NOW
	// enum: NOT NEED FOR NOW
	// struct embedding: NOT NEED FOR NOW
	// generic: NOT NEED FOR NOW
}
