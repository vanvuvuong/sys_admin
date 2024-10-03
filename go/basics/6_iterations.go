package basics

func RangeOverData() {
	// loop over array
	nums := [...]int{1: 10}
	for index, num := range nums {
		pf("Line %d: %-3d\n", index, num)
	}

	// loop over map
	exampleMap := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range exampleMap {
		pf("%s -> %6s\n", key, value)
	}
	for key := range exampleMap {
		pf("Key: %s\n", key)
	}
}
