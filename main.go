package main

import (
	"fmt"
	"strings"
)

func getInitials(name string) (string, string) {
	capped := strings.ToUpper(name)
	nameArr := strings.Split(capped, " ")

	var initials []string

	for _, value := range nameArr {
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	fName, lName := getInitials("Sean")
	fmt.Println(fName, lName)

	i := 0
	for i < len(myList) {
		fmt.Println(myList[i])
		i++
	}

	sayHello("Sean The Programmer")
}
