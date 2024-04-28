package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}
type Contact struct {
	Email string
	Phone string
}
type Adress struct {
	House int
	Area  string
	State string
}

type Employeee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Adress
}

func main() {
	var prince Person
	// fmt.Println("Prince person : ", prince)
	prince.firstName = "Aniket"
	prince.lastName = "Jha"
	prince.age = 20
	// fmt.Println("Prince person : ", prince)

	var employee1 Employeee
	employee1.Person_Details = Person{
		firstName: "aniket",
		lastName:  "Jha",
		age:       23,
	}

	employee1.Person_Contact.Email = "contactkar@gmail.com"
	employee1.Person_Contact.Phone = "923u29334343434"
	employee1.Person_Address = Adress{
		House: 23,
		Area:  "Bhopal",
		State: "MadhyaPradesh",
	}
	fmt.Println("Employee1 \n", employee1.Person_Contact.Email)
}
