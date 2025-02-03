package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	color.Red(input())
}

func input() string {
	fmt.Print("введите что-нибудь: ")
	var input string
	fmt.Scan(&input)
	return input
}
