package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// string pkg
	greeting := "hello there freinds!"

	fmt.Println(strings.Contains(greeting, "hello!"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "th"))
	fmt.Println(strings.Split(greeting, " "))

	//the original value is unchanged
	fmt.Println("Original string value: ", greeting)


	// sort pkg
	ages := []int{45, 20, 34, 67,12, 38, 90,23}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 20)
	fmt.Println(index)

	names := []string{"yoshi", "luigi", "mario", "bowser", "peach"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser"))

}