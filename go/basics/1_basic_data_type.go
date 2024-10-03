package basics

func Define() {
	// char variable example
	exampleChar := 'a'
	p("Character:", exampleChar)

	// string variable example
	exampleString := "string"
	p("String:", exampleString)

	// number variable example
	exampleFloat := 1.11111
	p("Float:", exampleFloat)
	exampleInt := 1
	p("Integer:", exampleInt)

	// multiple variables declaration
	exampleInt1, exampleInt2 := 1, 2
	p("Integer:", exampleInt1, exampleInt2)

	// array: fixed type & length
	var exampleArray1 [5]int
	exampleArray2 := [...]int{1: 5}
	p("Array:", exampleArray1, exampleArray2)

	// two-dimentional array
	exampleArray3 := [2][3]int{
		{1: 3},
		{1: 3},
	}
	p("2-dimentional array:", exampleArray3)

	// slice: fixed type, dynamic length
	var exampleSlice []string
	p("Slice:", exampleSlice)

	// map: like dictionary in Python
	exampleMap1 := make(map[string]string)
	exampleMap1["key"] = "value"
	p("Map:", exampleMap1)
	exampleMap2 := map[string]int{"foo": 1, "bar": 2}
	p("Map:", exampleMap2)
}
