package something

import (
	"fmt"
	"strings"
)

/**
Go 언어에서는 method 앞글자가 소문자면 private, 대문자면 public 으로 간주
*/
func sayBye() {
	fmt.Println("Bye")
}

func SayHello() {
	fmt.Println("Hello")
}

/**
함수를 작성할 때는 파라미터에 변수 명과 타입을 적고 파라미터 뒤에는 반환형을 적는다.
만약 파라미터 모두 같은 타입이라면 맨 마지막 파라미터에만 타입을 적어주면 된다.
*/
func multiply(a, b int) int {
	return a * b
}

func printMyName(firstName, lastName string) {
	fmt.Println(firstName + lastName)
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

/**
여러 값을 한번에 받아오고 싶으면 타입 앞에 (...) 을 붙여준다.
*/
func repeatMe(words ...string) {
	fmt.Println(words)
}
