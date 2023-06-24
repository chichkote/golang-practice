package main

import "fmt"

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	// group A types -> strings, ints, bools, floats, arrays, structs(non-pointer values)
	name := "tifa"

	name = updateName(name)
	fmt.Println(name)

	// group B types -> maps, slices, functions (pointer wrapper values)
	menu := map[string]float64 {
		"pie": 23.5,
		"ice cream": 10.5,
	}

	updateMenu(menu)
	fmt.Println(menu)

}