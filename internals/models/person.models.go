package models

import "fmt"

type person struct {
	name    string
	address string
	email   string
}

func AddPerson(name string, address string, email string) *person {
	return &person{
		name:    name,
		address: address,
		email:   email,
	}
}

func (p *person) Greet() {
	fmt.Println("hallo", p.name, "im go")
}

func (p *person) SetPerson(name string) {
	p.name = name
}
