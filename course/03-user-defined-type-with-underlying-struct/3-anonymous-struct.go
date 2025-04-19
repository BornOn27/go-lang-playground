package main

import "fmt"

/*
	This covers the concept of Anonymous Struct
	Anonymous struct is defined as struct with no name
	It is used when we want to use the struct only once,
	So there is no need to define it separately

	Since it doesn't have any name, we can't reference it
	for initializing the value.
	"So value must be assigned when it is getting defined"
 */


func main() {
		anonymousStructExample := struct {
			first string
			second bool
			third int
		}{
			first:  "This is string",
			second: true,
			third:  1000,
		}

		fmt.Println()
		fmt.Println("This is example of Anonymous Struct")
		fmt.Println(anonymousStructExample)
}
