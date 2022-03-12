package main

import (
	"fmt"
	"os"
	"strings"
)

func checkContains(slc []string, str string) (bool, string) {
	for _, v := range slc {
		if strings.ToLower(v) == strings.ToLower(str) {
			return true, v
		}
	}
	return false, ""
}

func checkIfOrderValid(orderValid *bool) bool {
	if len(os.Args) < 2 {
		*orderValid = false
		return false
	}
	return true
}

func readArguments() string {
	argumentsString := ""
	for _, v := range os.Args[1:] {
		argumentsString += v + " "
	}
	argumentsString = argumentsString[:len(argumentsString)-1]
	argumentsString = strings.ToLower(argumentsString)
	return argumentsString
}

func main() {

	var listOfBooks = []string{
		"The Paris Apartment",
		"House of Sky and Breath",
		"The Lincoln Highway",
		"Lord of The Rings: The Fellowship of the Ring",
		"Lord of The Rings: The Two Towers",
		"Lord of The Rings: The Return of The King",
	}

	orderValid := true

	// check if any argument given. If not, change the orderValid variable via sending the memory address
	if !checkIfOrderValid(&orderValid) {
		fmt.Println("Please give me a job to do")
	}

	if orderValid {

		// arguments in the lower case string format
		argumentsString := readArguments()

		switch argumentsString {
		case "list":
			fmt.Println("--- LIST OF ALL BOOKS ---")
			for _, v := range listOfBooks {
				fmt.Println(v)
			}

		default:
			// check if the slice contains the given book name
			b, s := checkContains(listOfBooks, argumentsString)
			if b {
				fmt.Println(s)
			} else {
				fmt.Println("Sorry,", argumentsString, "doesn't exist in our archive. Come later")
			}

		}
	}
}
