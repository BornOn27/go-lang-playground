package main

import "fmt"


func _returningFunctionWithNoParamsAndNoReturnType() func() {
	return func() {
		fmt.Println(`I am returned function with "no parameters and return-type"`)
	}
}

func _returningFunctionWithNoParamsAndWithReturnType() func() bool{
	return func() bool{
		fmt.Println(`I am returned function with "no parameters and with return-type"`)
		return true
	}
}


func _returningFunctionWithParamsAndNoReturnType() func(x int, y int) {
	return func(x int, y int) {
		fmt.Println(`I am returned function "with parameters and return-type"`, x, y)
	}
}

func _returningFunctionWithParamsAndWithReturnType() func(x int, y int) (int, int){
	return func(x int, y int) (int, int){
		fmt.Println(`I am returned function "with parameters and return-type"`, x, y)
		return y, x
	}
}


func main() {
	fmt.Println("")

	fmt.Println(`Calling underlying function - "_returningFunctionWithNoParamsAndNoReturnType()()"`)
	_returningFunctionWithNoParamsAndNoReturnType()()
	fmt.Println("")

	fmt.Println(`Calling underlying function - "_returningFunctionWithNoParamsAndWithReturnType()()"`)
	_returningFunctionWithNoParamsAndWithReturnType()()
	fmt.Println("")

	fmt.Println(`Calling underlying function - "_returningFunctionWithParamsAndNoReturnType()(100, 100)"`)
	_returningFunctionWithParamsAndNoReturnType()(100, 100)
	fmt.Println("")

	fmt.Println(`Calling underlying function - "_returningFunctionWithParamsAndWithReturnType()(1000, 1000)"`)
	_returningFunctionWithParamsAndWithReturnType()(1000, 1000)
	fmt.Println("")

}
