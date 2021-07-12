package main

import (
	"fmt"
	"strings"
)

/**
Go 에서는 두 개의 변수를 반환할 수 있다. 반환할 타입에 소괄호 () 로 타입을 정해주면 되고,
반환할 타입에 미리 변수 명을 붙여줄 수도 있다. 그럼 그 변수가 반환할 타입으로 초기화 된다.
*/
func lenAndUpperWithInitializeVariable(name string) (length int, uppercase string) {
	/**
	defer : 해당 메소드가 실행이 다 되고 난 이후에 수행되는 기능. 잘 활용하면 무궁무진하게 쓸 수 있을 것 같다.
	*/
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func lenAndUpperWithoutInitializeVariable(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLength, uppercase := lenAndUpperWithInitializeVariable("twosom")
	fmt.Print(totalLength, uppercase)

}
