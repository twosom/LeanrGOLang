package main

import "fmt"

func main() {
	const NAME = "name"
	const AGE = "age"
	twosom := map[string]string{
		NAME: "twosom",
		AGE:  "27"}

	fmt.Println(twosom[NAME])
	fmt.Println(twosom[AGE])
}
