package main

import "fmt"

/*
	This feature is mainly used to to call function at the end of the enclosing function
	It basically helps in closing files and similar operation,
	at the time of opening it.

	It huge improves the readability, as we are closing the resource at the time of opening
	not just of scroll up-down to ensure it is closed and resources released

*/

func foo() {
	fmt.Print("foo")
}

func bar() {
	fmt.Print("bar")
}


func main() {

	//OUTPUT - foobar
	foo()
	bar()

	fmt.Println()

	//OUTPUT - barfoo
	defer foo()
	bar()
}
