package main

import (
	"fmt"
	"math"
)

// shape interface
type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

// circle methods
func (s circle) area() float64 {
	return math.Pi * s.radius * s.radius
}

func (s circle) circumf() float64 {
	return 2 * math.Pi * s.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumf())
}


func main() {

	shapes := []shape {
		square{length: 12.2},
		circle{radius: 5.5},
		circle{radius: 56.2},
		square{length: 23},
	}

	for _,v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}

}