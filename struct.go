package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

//receiver function
func (b *bill) format() string {
	fs := "Bill brakdown: \n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	//tip
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip)

	//total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)
	return fs
} 

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

//add item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill -", b.name)

	return b
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("choose option (a -add item, s - save bill, t - add tip) ", reader)
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("item added - ", name, p)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip added - ", t)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("you saved the bill", b.name)
	default:
		fmt.Println("that was not a valid option...")
		promptOptions(b)
	}
}

// save bill
func (b *bill) save() {
		data := []byte(b.format())

		err :=os.WriteFile("bills/"+b.name+".txt", data, 0644)
		if err != nil {
			panic(err)
		}
		fmt.Println("bill was saved to file")
}

func main() {
	// mybill := newBill("mario's bill")
	// mybill.updateTip(10)
	// mybill.addItem("soup", 3.5)
	// mybill.addItem("coffe", 4.5)
	// fmt.Println(mybill.format())

	mybill := createBill()
	promptOptions(mybill)
	fmt.Println(mybill)

}

