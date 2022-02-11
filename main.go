package main

import (
	"fmt"
	"strings"
)

type bill struct {
	item string
	qty  int
}

func main() {
	fmt.Println("S Sharanya Bharghavi - 11903126")
	var bill bill
	var dayTotal = 0
	var flag = 1
	for flag == 1 {
		fmt.Println("Hello Customer! ")
		fmt.Println("Welcome to XYZ Cafe! Kindly Check out our menu below: ")
		printMenu()
		fmt.Println("Enter your Choice")
		fmt.Scan(&bill.item)
		bill.item = strings.ToUpper(bill.item)
		if bill.item == "END" {
			fmt.Println("Total Earnings of the day is: ", dayTotal)
			flag = 0
		} else {
			fmt.Println("Enter Quantity:")
			fmt.Scan(&bill.qty)
			total := calcAmt(bill)
			generateBill(bill, total)
			dayTotal += total

		}
	}
}

func printMenu() {

	fmt.Println("Item Name\t Choice \t Price")
	fmt.Println("Idli\t\t I\t\t Rs.20")
	fmt.Println("Vada\t\t V\t\t Rs.25")
	fmt.Println("Dosa \t\t D\t\t Rs.80")
	fmt.Println("Coffee\t\t C\t\t Rs.40")
	fmt.Println("Jalebi\t\t J\t\t Rs.30")
	fmt.Println("Lemonade\t L\t\t Rs.15")
	fmt.Println("Manchurian\t M\t\t Rs.90")
	fmt.Println("Spring Roll\t S\t\t Rs.20")
	fmt.Println("Tomato Soup\t T\t\t Rs.20")
	fmt.Println("French Fries\t F\t\t Rs.30")
	fmt.Println("Hakka Noodle\t H\t\t Rs.70")
	fmt.Println("Paneer Pakoda\t P\t\t Rs.30")
	fmt.Println("Bhature & Chole\t B\t\t Rs.30")

}
func calcAmt(bill bill) int {
	menu := map[string]int{
		"C": 40,
		"D": 80,
		"T": 20,
		"I": 20,
		"V": 25,
		"B": 30,
		"P": 30,
		"M": 90,
		"H": 70,
		"F": 30,
		"J": 30,
		"L": 15,
		"S": 20,
	}
	return bill.qty * menu[bill.item]
}
func generateBill(bill bill, total int) {
	decode := map[string]string{
		"C": "Coffee",
		"D": "Dosa",
		"T": "Tomato Soup",
		"I": "Idli",
		"V": "Vada",
		"B": "Bhature & Chole",
		"P": "Paneer Pakoda",
		"M": "Manchurian",
		"H": "Hakka Noodle",
		"F": "French Fries",
		"J": "Jalebi",
		"L": "Lemonade",
		"S": "Spring Roll",
	}
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("\tCustomer Receipt")
	fmt.Println("Item\t\tItem Code\t Quantity\t Total")
	fmt.Println(decode[bill.item], "\t\t", bill.item, "\t\t", bill.qty, "\t\t", "Rs.", total)
	fmt.Println("Thank You. Visit us Again")
	fmt.Println("--------------------------------------------------------------------")

}
