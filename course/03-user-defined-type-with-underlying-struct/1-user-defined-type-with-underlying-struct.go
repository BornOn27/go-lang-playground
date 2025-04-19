package main

import "fmt"

/*
	In go-lang we dont call/have class and object concept
	We call it as - A User Defined Type with underlying data-structure as "struct"
 */

type userDefinedTypeWithUnderlyingAsStruct struct {
	field_1 int
	field_2 bool
	field_3 string
}

func main() {

	//Now I will be creating a value of type "UserDefinedTypeWithUnderlyingAsStruct"
	value_1 := userDefinedTypeWithUnderlyingAsStruct{
		field_1: 100,
		field_2: true,
		field_3: "This is string for field_3",
	}
	fmt.Println()
	fmt.Println(value_1)

	value_2 := userDefinedTypeWithUnderlyingAsStruct{
		field_1: 1000,
		field_2: false,
		field_3: "This is another string for field_3",
	}

	fmt.Println(value_2)


	// This is generally not encouraged as a standard practise
	fmt.Println()
	fmt.Println("Another way of initializing the values")
	value_3 := userDefinedTypeWithUnderlyingAsStruct{
		1000,
		false,
		"This is another string for field_3",
	}
	fmt.Println(value_3)
}