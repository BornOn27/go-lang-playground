package main

import "fmt"

func incrementor() func() int {
	//Here x is available to entire scope of the "incrementor-function"
	var count int
	return func() int {
		count++
		return count
	}
}

func main() {
	x := incrementor()
	y := incrementor()


	/*
		Since incrementor() is called and reference to the function returned is stored
		in "x". Now if we call function attached to "x", it will call only the function returned by the function.
		No new memory will be allotted to the content of incrementor()

		incrementor() ===> 	*****************************
						    **  var count int		   **
							**	func <anonymous>{      **
							**		count++			   **
							**	}				       **
						    *****************************

		So since incrementor() is called once for x, memory allotted to will have "count variable" available
		in entire scope. Now since the reference to the memory is stored in x. When we call the inner function
		only that closure/block will be executed.

		"It is similar to calling a function over an object in java. Since the memory is already allotted to the object"
			"Every data in it will remain the same and shared across the context/scope"
	*/

	fmt.Println("count in x = ", x())
	fmt.Println("count in x = ", x())
	fmt.Println("count in x = ", x())
	fmt.Println("count in x = ", x())


	/*
		Here we have again called incrementor(), so a new memory block will be created with value of count=0
		And called inner function on y will increment this memory block variable
	*/
	fmt.Println("count in y = ", y())
	fmt.Println("count in y = ", y())
}