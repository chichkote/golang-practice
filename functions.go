package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _,v := range n {
		f(v)
	}
}

func circleArea(r float64) float64{
	return math.Pi * r *r
}

// return multiple values
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _,v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	// sayGreeting("mario")
	// sayBye("luigi")
	// sayBye("mario")
	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	cycleNames([]string{"cloud", "tifa", "barret"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)
	fmt.Println(a1, a2)

	fn1, sn1 := getInitials("tifa lockhart")
	fmt.Println(fn1,sn1)

	fn2, sn2 := getInitials("cloud")
	fmt.Println(fn2,sn2)
}