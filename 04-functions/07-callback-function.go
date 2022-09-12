package main

import "fmt"

/*

	Callback function is the name given to the function when passed as a 
	Parameter to a function

*/

func toBePassedAsArgumentToAnotherFunction(x ...int)  int{
	for _, v := range x{
		fmt.Print(v,", ")
	}

	return 0
}

func functionToAcceptFunctionAsParameter(functionAsParameter func(x ...int) int, y ...int)  {
	fmt.Println("Calling function received as Parameter - ")
	returnedValue := functionAsParameter(y...)
	fmt.Println("Value Returned ", returnedValue)
}

func main() {
	fmt.Println()

	//We can call the function as usual
	fmt.Println("We can call the function as usual")
	xi := []int{1,2,3,4,5,6,7,}
	toBePassedAsArgumentToAnotherFunction(xi...)
	fmt.Println()
	fmt.Println()

	//We can pass the argument via Variable
	referenceToFunction := toBePassedAsArgumentToAnotherFunction
	fmt.Println("We can pass the argument via Variable")
	functionToAcceptFunctionAsParameter(referenceToFunction, xi...)
	fmt.Println()

	//Or We can pass the argument directly
	functionToAcceptFunctionAsParameter(toBePassedAsArgumentToAnotherFunction, xi...)

}