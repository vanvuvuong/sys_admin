package basics

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

func zero(interval int) {
	interval = 0 // dereference the pointer & assign value to 0
}

func zeropointer(interval *int) {
	// *int mean take `interval` as pointer
	*interval = 0 // dereference the pointer & assign value to 0
}

func AdvancedDefine() {
	// pointer: use with `&` before variable, point to memory reference
	interval := 1
	p("Pointer reference value", &interval)
	p("Pointer value", interval)
	zero(interval)
	zeropointer(&interval)
	p("Pointer reference value", &interval)
	p("Pointer value", interval)

	// struct implement
	exampleStruct := new(exampleStructType)
	exampleStruct.Id = 1
	exampleStruct.Name = "hello motherfucker"
	p("Example struct", exampleStruct)

	// method implement
	c := circle{1.2}
	pf("Area of %f circle is %f\n", c.Radius, c.area())

	// interface: NOT NEED FOR NOW
	// enum: NOT NEED FOR NOW
	// struct embedding: NOT NEED FOR NOW
	// generic: NOT NEED FOR NOW
}
