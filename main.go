package main

import "fmt"

func main() {
	names := []string{"one", "two", "three"}
	names = append(names, "flynn")
	fmt.Print(names)
}
