package main

import "fmt"


var package_level_x int
var package_level_y string
var package_level_z bool

func main(){
	/*

	Exercise 1:
	1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the
	IDENTIFIERS “x” and “y” and “z”
		a. 42
		b. “James Bond”
		c. true

	2. Now print the values stored in those variables using
		a. a single print statement
		b. multiple print statements

	*/

	x := 42
	y := "James Bond"
	z := true

	fmt.Println("::::::::::: Exercise-1 :::::::::::")
	fmt.Println(x, y, z)

	//Since "fmt.Println" doesn't allow concatenation of different type (string + int), we need to use "fmt.Sprintf"
	fmt.Printf("x=%d	y=%s	z=%t  \n", x, y, z)
	fmt.Println(fmt.Sprintf("x=%d	y=%s	z=%t", x, y, z))




	/*

	Exercise 2:
	1. Use var to DECLARE three VARIABLES. The variables should have package level
	scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
		a. identifier “x” type int
		b. identifier “y” type string
		c. identifier “z” type bool

	2. in func main
		a. print out the values for each identifier
		b. The compiler assigned values to the variables. What are these values called?

	*/
	fmt.Println()
	fmt.Println("::::::::::: Exercise-2 :::::::::::")
	fmt.Printf("Zero value of package_level_x=%d	package_level_y=%s	package_level_z=%t  \n", package_level_x, package_level_y, package_level_z)
	fmt.Println(`Default values assigned by compiler are called "Zero Values"`)


	/*

	Exercise-3
	Using the code from the previous exercise,
	1. At the package level scope, assign the following values to the three variables
		a. for x assign 42
		b. for y assign “James Bond”
		c. for z assign true

	2. in func main Todd McLeod - Learn To Code Go on Udemy - Part 1 - Page 26
		a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
		b. print out the value stored by variable “s”

	 */
	fmt.Println()
	fmt.Println("::::::::::: Exercise-3 :::::::::::")
	package_level_x = 42
	package_level_y = "James Bond"
	package_level_z = true

	s := fmt.Sprintf("package_level_x=%d\tpackage_level_y=%s\tpackage_level_z=%t", package_level_x, package_level_y, package_level_z)
	fmt.Println(s)


	/*

	Exercise-4
	For this exercise
	1. Create your own type. Have the underlying type be an int.

	2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR”
	keyword

	3. in func main
		a. print out the value of the variable “x”
		b. print out the type of the variable “x”
		c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
		d. print out the value of the variable “x”

	*/

	fmt.Println()
	fmt.Println("::::::::::: Exercise-4 :::::::::::")

	type VariableWrapper bool

	var variableOfCustomType VariableWrapper
	fmt.Print("variableOfCustomType = ")
	fmt.Print(variableOfCustomType)
	fmt.Printf(" with type=%T", variableOfCustomType)
	fmt.Println()
	fmt.Println("The default or Zero value of Custom Defined Variables " +
		"are Default/Zero value of underlying type")
	fmt.Println()
	variableOfCustomType = true
	fmt.Printf("After assigning value, variableOfCustomType = %t with type=%T", variableOfCustomType, variableOfCustomType)


	/*

	Enlightenment on Type Usage and Underlying Type

	type T1 int
	type T2 T1
	type T3 []T2
	type T4 T3

	All have same underlying type as "Int"
	 */


	/*

	Exercise-5

	Building on the code from the previous example
	1. at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
	a. eg:
			type hotdog int
			var x hotdog
			var y int
	2. in func main
		a. this should already be done
			i. print out the value of the variable “x”
			ii. print out the type of the variable “x”
			iii. assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
			Todd McLeod - Learn To Code Go on Udemy - Part 1 - Page 27
			iv. print out the value of the variable “x”
		b. now do this
			i. now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
				1. then use the “=” operator to ASSIGN that value to “y”
				2. print out the value stored in “y”
				3. print out the type of “y”
	 */
	fmt.Println()
	fmt.Println()
	fmt.Println("::::::::::: Exercise-5 :::::::::::")
	type conversionExample_X int

	var conversionExample_Y conversionExample_X
	var conversionExample_Z int


	conversionExample_Y = 20
	fmt.Printf("conversion_example_Y = %d", conversionExample_Y)
	fmt.Println()

	conversionExample_Z = 250
	fmt.Printf("conversion_example_Z = %d", conversionExample_Z)
	fmt.Println()

	conversionExample_Z = int(conversionExample_Y)
	fmt.Printf("Value after conversion from conversionExample_Y is conversion_example_Z = %d", conversionExample_Z)
	fmt.Println()
}
