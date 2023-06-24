package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":  4.99,
		"salad": 5.98,
		"pie":   6.98,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//looping maps
	for k,v := range menu {
		fmt.Println(k, "-", v)
	}

	//ints as key type
	phonebook := map[int]string{
		123456789: "mario",
		987654321: "luigi",
		456789124: "yoshi",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[987654321])

	phonebook[987654321] = "bowser"
	fmt.Println(phonebook)
}