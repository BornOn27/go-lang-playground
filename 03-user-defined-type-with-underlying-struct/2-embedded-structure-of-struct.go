package main

import "fmt"

/*
	This is an example of Embedded/Tree-Structure of Struct
	Where a struct can contain another struct within itself
 */


/*
==> "KEY-TAKE-AWAY"
	The fields in embedded struct will get promoted to super struct,
	in case there is no collision of fields between embedded struct
	and super struct

 */
type thisIsEmbeddedStruct struct {
	field1 int
	field2 bool
	field3 string
}

type thisIsSuperStructWithEmbeddedStructInside struct {
	/*
	********************* Important *****************************
		While creating a embedded struct we dont need to define the variable name
		for underlying struct, by default the variable name of
		struct is its name itself
	*************************************************************
	 */
	thisIsEmbeddedStruct	//Example of Anonymous Type
	justAnotherField string
}

type thisIsSuperStructWithEmbeddedStructInsideWithFieldCollision struct {
	/*
		********************* Important *****************************
		This is a case in which field is present in embedded struct
		as well as field of super struct itself
		*************************************************************
	*/
	thisIsEmbeddedStruct
	field3 string
	justAnotherField string
}

func main() {

	//While defining the value of embedded struct,
	//we will use the struct name itself for reference

	superStruct := thisIsSuperStructWithEmbeddedStructInside{
		thisIsEmbeddedStruct : thisIsEmbeddedStruct{
			field1: 100,
			field2: true,
			field3: "This is embedded struct string field",
		},
		justAnotherField: "This is just another string",
	}

	//For referring to the value of embedded type, we don't have to use the reference
	//We can access the field with the super struct
	fmt.Println()
	fmt.Println("This is example of Embedded Struct")
	fmt.Printf("superStruct.field1 : %v, superStruct.field2 : %v, superStruct.field3 : %v", superStruct.field1,superStruct.field2,superStruct.field3)
	fmt.Printf("superStruct.thisIsEmbeddedStruct.field1 : %v, superStruct.thisIsEmbeddedStruct.field2 : %v, superStruct.thisIsEmbeddedStruct.field3 : %v", superStruct.thisIsEmbeddedStruct.field1,superStruct.thisIsEmbeddedStruct.field2,superStruct.thisIsEmbeddedStruct.field3)


	fmt.Println()
	fmt.Println()
	fmt.Println("This is example of Embedded Struct with field collision")

	superStructWithCollision := thisIsSuperStructWithEmbeddedStructInsideWithFieldCollision{
		thisIsEmbeddedStruct : thisIsEmbeddedStruct{
			field1: 100,
			field2: true,
			field3: "This is embedded struct string field",
		},
		field3: "This is collided field with Embedded Struct",
		justAnotherField: "This is just another string",
	}

	fmt.Printf("superStructWithCollision.field3 = %v", superStructWithCollision.field3)
	fmt.Println()
	fmt.Printf("superStructWithCollision.thisIsEmbeddedStruct.field3 = %v", superStructWithCollision.thisIsEmbeddedStruct.field3)
}
