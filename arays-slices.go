package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{1,2,3}
	var ages = [3]int{1, 2, 3}
	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slices (use arrays as under hood)
	var scores = []int{10, 20, 30}
	scores[2] = 25
	scores = append(scores, 35)

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeone := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeone, rangeTwo, rangeThree)

}