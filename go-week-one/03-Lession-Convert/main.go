package main

import "fmt"

type Person struct {
	Name string
	Age  string
}

type Male struct {
	Name string
	Age  int
}

type Student struct {
	Firstname string
	Lastname  string
	Age       int
}

func printDISPLAY(p Person) {
	fmt.Println(p.Name)
}

func main() {
	fmt.Println("start golang")
	aMale := Male{Name: "Teo", Age: 34}
	aPerson := Person{Name: "Ty Person", Age: "40"}

	printDISPLAY(aPerson)
	fmt.Println(aMale)

}
