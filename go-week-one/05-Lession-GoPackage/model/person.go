package model

type Person struct {
	Name string
	Age  string
}

func (p Person) GetFirstName() string {
	return p.Name
}

func (p *Person) ChangeName(name string) {
	p.Name = name
}
