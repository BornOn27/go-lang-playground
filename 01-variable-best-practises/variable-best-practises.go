package main

import "fmt"

/*

General Knowledge about variable in go-lang

- unused variable will throw compilation error in go-lang
- Underscore "_" can be used if in case value needs to be stored but not used anywhere
- Short-hand variables are not meant for global or package level visibility
- For package level, we should always use var-variant "var variable"

 */
func main() {
	//This is short-hand variable declaration
	thisVariableTypeDefinedAutomaticallyByItsValue := "I am string"
	fmt.Println(thisVariableTypeDefinedAutomaticallyByItsValue)

	//This is 2 step declaration, where we have to give Type
	var thisIsAnotherWayToDefineVariableWithTypeDefinedAfterVariable string
	thisIsAnotherWayToDefineVariableWithTypeDefinedAfterVariable = "I am another way of defining string"
	fmt.Println(thisIsAnotherWayToDefineVariableWithTypeDefinedAfterVariable)

	//This is variation of above where we can assign the value during declaration
	var variationOfAboveWhereValuesAssignedDuringDeclaration string = "I am variable with value at declaration"
	fmt.Println(variationOfAboveWhereValuesAssignedDuringDeclaration)
}
