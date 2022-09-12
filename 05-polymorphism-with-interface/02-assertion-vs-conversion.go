package main

import "fmt"

/*
	Since in GO-Lang, a value can be of "Multiple Type"
	So if "golang-interface" is used as function argument,
	and we need to find the type of the object,

	We use assertion and it helps us to take different action, according to type of value/object
	i.e,

	superType...(subType)...<field>

*/

type name struct {
	firstName string
	lastName string
}

type deer struct {
	name
	canHunt bool
}

type bear struct {
	name
	canHunt bool
}

func (d deer) canHuntAnimals() {
	fmt.Printf("Do I hunt Animals? %v", d.canHunt)
}

func (b bear) canHuntAnimals() {
	fmt.Printf("Do I hunt Animals? %v", b.canHunt)
}

type animal interface {
	canHuntAnimals()
}

func doesCageRequiredInZoo(animal animal) {
	//This is the example of "Assertion"

	switch animal.(type) {
	case bear:
		fmt.Printf("Does this animal require Cage? %v", animal.(bear).canHunt)
		fmt.Println()
	case deer:
		fmt.Printf("Does this animal require Cage? %v", animal.(deer).canHunt)
		fmt.Println()
	}
}

func main() {

	deer := deer{
		name{
			"I am ",
			"deer",
		},
		false,
	}

	bear := bear{
		name{
			"I am ",
			"bear",
		},
		true,
	}

	fmt.Println()
	fmt.Println("Deer Details ", deer)
	doesCageRequiredInZoo(deer)
	fmt.Println()

	fmt.Println("Bear Details ", bear)
	doesCageRequiredInZoo(bear)
}


