package main

import "fmt"

// IntegerWrapper - Here I am creating a "Type" which internally
//will be of any supported Data-Type of Go-Lang
//But it is hidden (in a way behind a type) from but hidden since it is static language
//And while assigning values, we would know what data it will accept
type IntegerWrapper int

var misleadingInteger IntegerWrapper

//This is a variable of "Type" int
var integer int


func main() {
	misleadingInteger = 10
	integer = 10

	fmt.Println(fmt.Sprintf("integer variable is definitely a %T with value %d", integer, integer))

	fmt.Println(fmt.Sprintf("Whereas misleadingInteger seems to be an Integer but it is %T with value %d", misleadingInteger, misleadingInteger))

	/*

		Since both seems to be an integer but they are not
		So value assigning wont be allowed directly via variable
		Since both are of different "Type"

		integer = misleadingInteger
		misleadingInteger = integer

	*/
}
