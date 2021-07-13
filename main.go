package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"Chicken", "Noodle"}
	twosom := person{name: "twosom", age: 27, favFood: favFood}
	fmt.Println(twosom.age)

}
