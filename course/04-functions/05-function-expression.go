package main

import "fmt"

/*
	Just like other types in Go-Lang, a function expression can be stored in a variable
	And can be called via it, just like the function call

	It can be both i.e, with and w/o Parameters

*/

func main() {
	fmt.Println()
	functionExpression := func() {
		fmt.Println("I am a functionExpression")
	}

	functionExpression()

	functionExpressionWithParameters := func(x int, y string) {
		fmt.Println("Well I am a functionExpressionWithParameters with parameters as - ", x, y)
	}

	functionExpressionWithParameters(100, "String Parameter")
}
