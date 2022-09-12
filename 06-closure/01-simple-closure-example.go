package main

import "fmt"

func main() {
	fmt.Println()
	x := "This is available to entire main function"
	fmt.Println("Accessing main-function scope variable ")
	fmt.Println("x = ", x)

	{
		fmt.Println("I am a simple closure block. I can access everything at main-function scope" +
			". But main-function can't access my content")

		x := "This is variable x with closure level scope, any outer variable will be ignored"
		y := "This is just a simple variable"

		fmt.Println("Different closure level variables - ")
		fmt.Println("x = ", x)
		fmt.Println("y = ", y)

	}

	//fmt.Println("y = ", y) ==> y is not available here
}
