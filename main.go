package main

import "fmt"

func main() {
	bill := newBill("Sean")
	bill.tip = 20.00
	fmt.Println(bill.format())

}
