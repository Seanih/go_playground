package main

import (
	"fmt"
	"strings"
)

func main() {
	sayMyName("")
	fmt.Println(printArray("yesyouare"))
}

func sayMyName(name string) {
	if len(name) == 0 {
		fmt.Println("Nothin to say...")
	} else {
		fmt.Printf("Wuz hannen, %v!", name)
	}
}

func printArray(text string) string {
	if len(text) == 0 {
		return fmt.Sprintln("bitch ass nigga")
	}

	textSlice := strings.Split(text, "")
	myMap := make(map[string]uint8)

	for i := 0; i < len(textSlice); i++ {
		if myMap[textSlice[i]] == 0 {
			myMap[textSlice[i]] = 1
		} else {
			return fmt.Sprintf("the repeated value is: %v", textSlice[i])
		}
	}

	return ""
}
