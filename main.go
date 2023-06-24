package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	//String
	var nameOne string = "mario"
	var nameTwo = "Luigi"
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"

	fmt.Println(nameFour)

	//numbers
	//ints

	var age1 int = 20
	var age2 = 30
	age3 := 40
	fmt.Println(age1, age2, age3)

	//bits and memory
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint = 256

	//float
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 1234567898765432123456.45
	scoreThree := 1.5
	fmt.Println(scoreOne, scoreTwo, scoreThree)

	//print
	fmt.Print("hello, ")
	fmt.Print("World! \n")
	fmt.Print("new line \n")

	fmt.Println("hello, World!")
	fmt.Println("Goodbye")
	fmt.Println("my age is ", age1, " and my name is", nameOne)

	//Printf formatted stirngs, _% = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age1, nameOne)
	fmt.Printf("my age is %q and my name is %q \n", age1, nameOne)
	fmt.Printf("my age is %T", age1)
	fmt.Printf("you scored %f points \n", 255.55)
	fmt.Printf("you scored %0.1f points \n", 255.55)

	//Sprintf(save formatted string
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age1, nameOne)
	fmt.Println("the saved string is ", str)
}