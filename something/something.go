package something

import "fmt"

/**
	Go 언어에서는 method 앞글자가 소문자면 private, 대문자면 public 으로 간주
 */
func sayBye() {
	fmt.Println("Bye")
}

func SayHello() {
	fmt.Println("Hello")
}