package main

import "fmt"

/*

	********************************* IMPORTANT ***************************************************
	**						Polymorphism + Inheritance							 				 **
	***********************************************************************************************
	**	type Species interface{																	 **
	**			speak()																			 **
	**	}																						 **
	***********************************************************************************************
	**																							 **
	** 	This is one of the main concept of GO-Lang, which is uncommon from other languages		 **
	**	GO-Lang we don't have the concept of class, it only have struct.				 		 **
	**	This concept of Interface allows us to implement "Polymorphism & Abstraction"
	**
	**	"Key-Take-Away"
	**		- A value can be of multiple type
	**		- "Any type that has the method/function speak() is also a Species Type"
			- In general terms, If any type has a function which have me is also my type
			- Empty Interface [interface{}] used in Function parameter can accept anything
				because all the types in GO-Lang will have an empty

	***********************************************************************************************

*/

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	ltk bool
}

type terrorist struct {
	person
	ltk bool
}

func (agent secretAgent) haveLicenceToKill() {
	fmt.Println("I have licence to kill")
}

func (t terrorist) haveLicenceToKill() {
	fmt.Println("I don't have licence to kill")
}

//This means any type that has the method "haveLicenceToKill()" will of type "killers"
type killers interface {
	haveLicenceToKill()
}

func whoIsThisKiller(killer killers) {
	fmt.Println("Who are you?")
	fmt.Printf("I am a %v.", killer)
	fmt.Println()
	fmt.Println("Do you have licence to kill?")
	killer.haveLicenceToKill()
}

func main() {
	fmt.Println()
	agent := secretAgent{
		person{"James",
			"Bond",
		},
		true,
	}

	terrorist := terrorist{
		person{
			"Mr. Dark",
			"Soul",
		},
		false,
	}

	whoIsThisKiller(agent)
	fmt.Println()
	whoIsThisKiller(terrorist)
}
