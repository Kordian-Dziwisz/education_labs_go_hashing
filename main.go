package main

import (
	"fmt"
	"os"
)

func readUserInput() string {
	return os.Args[1]
}

func formatInput(input *string) *string {
	formatedInput := ""
	for _, char := range *input {
		fmt.Println(string(float32(int(char))))
		formatedInput = formatedInput + string(int(char))
	}
	return &formatedInput
}

func main() {
	input := readUserInput()
	fmt.Println(formatInput(&input))
}
