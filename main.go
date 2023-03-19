package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(printArray("cbac"))

	cycleNames([]string{"Sean", "Programmer", "Cold"}, sayMyName)

	showRadius(3.5)
}

func sayMyName(name string) {
	if len(name) == 0 {
		fmt.Println("Nothin to say...")
	} else {
		fmt.Printf("Wuz hannen, %v! \n", name)
	}
}

func printArray(text string) string {
	if len(text) == 0 {
		return fmt.Sprintln("bitch ass nigga")
	}

	textSlice := strings.Split(text, "")
	myMap := make(map[string]uint8)
	var result string

	for i := 0; i < len(textSlice); i++ {
		if myMap[textSlice[i]] == 0 {
			myMap[textSlice[i]] = 1
			continue
		}

		if myMap[textSlice[i]] == 1 {
			result = fmt.Sprintf("the first repeated value is: %v", textSlice[i])
			break
		}
	}

	return result
}

func cycleNames(names []string, f func(string)) {
	for i := 0; i < len(names); i++ {
		f(names[i])
	}
}

func showRadius(r float32) {
	fmt.Println("Radius is: ", math.Pi*r*r)
}
