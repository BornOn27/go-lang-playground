package main

import "fmt"

/*

	The term "Anonymous simply means without Name/Identifier"
	Just like "Anonymous struct", we have "Anonymous Function"

	Format -
			func(parameter[s]){
				<code>
			}(Argument[s])

*/

func main() {

	//Example of Anonymous Function
	func(){

	}()

	fmt.Println()

	func(){
		fmt.Println("I am Anonymous function without Argument")
	}()

	func(x int){
		fmt.Printf("I am Anonymous function with Argument %v", x)
	}(1000)
}
