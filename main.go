package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func readUserInput() string {
	return strings.ToLower(os.Args[1])
}

func formatInput(input *string) *string {
	formatedInput := ""
	for _, char := range *input {
		formatedInput = formatedInput + strconv.FormatInt(int64(char), 10)
	}
	return &formatedInput
}

func silnia(n int64) *big.Int {
	value := new(big.Int).SetInt64(1)
	for i := n; i > 0; i-- {
		value = value.Mul(value, new(big.Int).SetInt64(i))
	}
	return value
}

func findRecursive(s string, searchedS string) bool {
	if len(searchedS) == 0 {
		return false
	}
	char := rune(searchedS[0])
	idx, err := findAsciiInStr(s, char)
	if err {
		return true
	} else {
		fmt.Println("index for ", char, ": ", idx)
		return findRecursive(s[idx+3:], searchedS[1:])
	}
}

func findAsciiInStr(s string, char rune) (int, bool) {
	chars := strconv.FormatInt(int64(char), 10)
	i := strings.Index(s, chars)
	err := false
	if i == -1 {
		err = true
	}
	return i, err
}

func main() {
	input := readUserInput()
	fmt.Println(formatInput(&input))
	for i := int64(0); i < 10000; i++ {
		fmt.Println("for ", i, "!:")
		s := silnia(i).String()
		if !findRecursive(s, input) {
			fmt.Println("Found it!")
			break
		}
	}
}
