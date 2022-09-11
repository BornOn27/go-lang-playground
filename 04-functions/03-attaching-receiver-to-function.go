package main

import "fmt"

/*

	********************************* IMPORTANT ***************************************************
	** 	function-syntax																			 **
	***********************************************************************************************
	**	func (r receiver) identifier(parameter[s]) (return[s]) {code}							 **
	***********************************************************************************************
	**																							 **
	** 	This is one of the main concept of GO-Lang, which is uncommon from other languages		 **
	**	Since in GO-Lang we don't have the concept of class, it only have struct,				 **
	**	In other language, we need an object to call a function, similarly GO-Lang				 **
	**	provides the way we can do that.														 **
	**																							 **
	**	i.e																						 **
	**																							 **
	**	We can bind a function to a "Value of certain type of struct", which means that function **
	**	can be called from a value of certain struct, which we call it as a Receiver			 **
	***********************************************************************************************

*/

type person struct {
	firstName string
	lastName string
}

type secretAgent struct {
	person
	ltk bool
}

func (agent secretAgent) orderToKill(target person) bool {
	fmt.Printf("Agent %s %s loading ...", agent.firstName, agent.lastName)
	fmt.Println()
	fmt.Printf("Order To Neutralise Target [%s %s] recieved", target.firstName, target.lastName)
	return true
}

func main() {
	//Not a recommended way of assigning the value, Always use the variable name
	target1 := person{
		"Mr. Underworld",
		", The NYs' Druglord",
	}

	agent1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	fmt.Println()
	fmt.Println("Sending Order to neutralize a target ...")
	fmt.Printf("Did Order Recieved?")
	fmt.Println()
	//Value of SecretAgent is used to call the function
	agent1.orderToKill(target1)
	fmt.Println()

}


