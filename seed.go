package main

func init() {
	people = append(people, Person{ID: "1", FirstName: "Nic", LastName: "Raboy", Address: &Address{City: "Dublin", State: "CA"}})
	people = append(people, Person{ID: "2", FirstName: "Maria", LastName: "Raboy"})
}