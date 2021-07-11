package main

import (
	"fmt"
	"strings"
)

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

func main() {
	fmt.Println(multiply(2, 2))
	printMyName("two", "som")

	totalLength, upperName := lenAndUpper("twosom")
	fmt.Println(totalLength, upperName)
	repeatMe("seoul", "incheon", "pusan", "bucheon", "daegu")

}
