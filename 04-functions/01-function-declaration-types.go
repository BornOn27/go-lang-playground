package main

import "fmt"

/*
	Basic structure of function -

	func (r receiver) identifier(parameter[s]) (return[s]) {code}

	func			==> 	Mandatory	==> Keyword to write a function
	(r receiver) 	==> 	Optional	==>	To be discussed further
	identifier 		==> 	Mandatory 	==> Function Name
	parameter[s] 	==> 	Optional	==> Parameters passed to the function
	return[s]		==> 	Optional 	==> One or multiple values returned by function

*/

func functionType1(){
	fmt.Println("I am simplest function")
}

func functionType2(x int, y int){
	fmt.Println("I am a function with parameters", x, y)
}

func functionType3(x, y int){
	fmt.Println("I am a variation of previous function with parameters", x, y)
	fmt.Println("I am not recommended since because of BAD READABILITY")
	fmt.Println(`"x, y int" is not encouraged but "x int, y int"`)
}

/*
	One thing to note here is the argument of function is called "Variadic (Variable) parameter"
	Although it is going to store the data in "Slice/Array",
		"It doesn't mean we can pass Slice/Array."
		"Since GO-Lang is static language,
			It is not allowed as it is of different type"

	i.e, It is asking "Unlimited Number of Ints", not Slice of Int

	[Important] ==> "But there is a way to do that"

	xi := int[]{1,2,3,4,5,6,7}
	Not-Allowed		==> 	functionType4(xi)
	Allowed			==> 	functionType4(xi...)

	It is allowed in 2nd case because
		"Compiler finds User is aware of Function Parameter/Argument,
			And since the INPUT provided will be anyways converted to slice,
			Let's just allow that and Compiler will convert into slice and
			pass to the function"

		It is generally Called ==> "Unfurling of Slice"

*/
func functionType4(x ...int)  {
	fmt.Println("I am a function with variable Parameters")
	fmt.Println("I can accept any number of argument")
	fmt.Println("I always store the values in SLICE i.e Arrays")

	fmt.Printf("%T", x)
}

/*

	func functionTypeWithInvalidVariadicParameterDeclaration(x ...int, y ...string)

	This is not allowed as "Variadic Parameter must always be at last
	That means we can't use more than one "Variadic Parameter"

	This is not allowed because in case of Variadic Parameter
	"We can pass 0 argument, and that could mess-up further argument/parameter"

*/

func functionTypeWithMixedVariadicParameterDeclaration(x int, y ...int)  {
	sum := 0
	fmt.Println("I am going to add all the parameters")

	for values := range y{
		sum += values
	}

	fmt.Printf("Total Sum of Integers %d", sum)
}


func main() {
	functionType1()

	functionType2(10, 12)

	functionType3(100, 300)

	functionType4(1,2,3,4,5,6,7,8,9)

	functionTypeWithMixedVariadicParameterDeclaration(1, 2,3,4,5,6,7,8,9)
}
