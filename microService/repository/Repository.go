package repository

import "fmt"

func (b *repository) Create() {
	fmt.Println("Creating a new record")
}

func (b *repository) Update() {
	fmt.Println("Updating an existing record")
}

func (b *repository) Delete() {
	fmt.Println("Deleting a record")
}

func (b *repository) Read() {
	fmt.Println("Reading a record")
}

func (b *repository) Filter() {
	fmt.Println("Filtering records")
}

func (b *repository) GetAll() {
	fmt.Println("Getting all records")
}
