package main

import "fmt"

/*
	Syntax -
		func <function-name>(<parameter[s]>) func(<parameter[s]>) <return-type>
		where <return-type> is optional
*/


func returningFunctionWithNoParamsAndNoReturnType() func() {
	return func() {
		fmt.Println(`I am returned function with "no parameters and return-type"`)
	}
}

func returningFunctionWithNoParamsAndWithReturnType() func() bool{
	return func() bool{
		fmt.Println(`I am returned function with "no parameters and with return-type"`)
		return true
	}
}


func returningFunctionWithParamsAndNoReturnType() func(x int, y int) {
	return func(x int, y int) {
		fmt.Println(`I am returned function "with parameters and return-type"`, x, y)
	}
}

func returningFunctionWithParamsAndWithReturnType() func(x int, y int) (int, int){
	return func(x int, y int) (int, int){
		fmt.Println(`I am returned function "with parameters and return-type"`, x, y)
		return y, x
	}
}



func main() {
	fmt.Println("")

	functionTypes1 := returningFunctionWithNoParamsAndNoReturnType()
	fmt.Println("Returned function format - " + fmt.Sprintf("%T", functionTypes1))
	fmt.Println("Calling underlying function - ")
	functionTypes1()
	fmt.Println("")

	functionTypes2 := returningFunctionWithNoParamsAndWithReturnType()
	fmt.Println("Returned function format - " + fmt.Sprintf("%T", functionTypes2))
	fmt.Println("Calling underlying function - ")
	functionTypes2()
	fmt.Println("")

	functionTypes3 := returningFunctionWithParamsAndNoReturnType()
	fmt.Println("Returned function format - " + fmt.Sprintf("%T", functionTypes3))
	fmt.Println("Calling underlying function - ")
	functionTypes3(100, 100)
	fmt.Println("")

	functionTypes4 := returningFunctionWithParamsAndWithReturnType()
	fmt.Println("Returned function format - " + fmt.Sprintf("%T", functionTypes4))
	fmt.Println("Calling underlying function - ")
	functionTypes4(1000, 1000)
	fmt.Println("")

}
