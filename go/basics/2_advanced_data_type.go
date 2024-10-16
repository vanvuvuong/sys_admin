package basics

// struct: typed collections of fields
type exampleStructType struct {
	Id   int32
	Name string
}

// method
type Circle struct {
	Radius float32
}

func (c Circle) area() float32 {
	return 3.14 * c.Radius * c.Radius
}

func (c *Circle) bigger() {
	c.Radius = 10 * c.Radius
}

func zero(interval int) {
	interval = 0 // assign value to 0
}

func zeropointer(interval *int) {
	// *int mean take `interval` as pointer
	*interval = 0 // dereference the pointer & assign value to 0
}

func AdvancedDefine() {
	// pointer: use with `&` before variable, point to memory reference
	interval := 1
	zero(interval)
	P("Pointer without pointer symbol reference value", &interval)
	P("Pointer without pointer symbol value", interval)
	zeropointer(&interval)
	P("Pointer without pointer symbol reference value", &interval)
	P("Pointer without pointer symbol value", interval)

	// struct implement
	exampleStruct := new(exampleStructType)
	exampleStruct.Id = 1
	exampleStruct.Name = "hello motherfucker"
	P("Example struct", exampleStruct)

	// method implement
	c := Circle{1.2}
	Pf("Area of %f circle is %f\n", c.Radius, c.area())

	// method used pointer
	// bigC := &Circle{1.5} // same as bigC := new(Circle)
	bigC := new(Circle)
	bigC.Radius = 1.5
	bigC.bigger()
	Pf("Big circle radius: %f\n", bigC.Radius)

	// interface: NOT NEED FOR NOW
	// enum: NOT NEED FOR NOW
	// struct embedding: NOT NEED FOR NOW
	// generic: NOT NEED FOR NOW
}
