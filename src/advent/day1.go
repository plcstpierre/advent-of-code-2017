package main

import (
	"fmt"
	"commons"
)

func main() {
	fmt.Println("hello world")
	commons.ReadFile("input/day1/a.txt", printTheString)
}

func printTheString(value string) {
	fmt.Println(value)
}
