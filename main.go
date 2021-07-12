package main

import "fmt"

func canIDrink(age int) bool {

	/**
	Go 에서는 if 문 안에 바로 변수를 만들 수 있다. (Variable Expression)
	if else 절 안에서만 사용한다는 의미이다.
	그리고 ;(세미콜론) 뒤에 만든 변수를 곧바로 사용할 수 있다.
	*/
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
