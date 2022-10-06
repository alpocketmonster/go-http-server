package subtest

type Person struct {
	age int
}

type person2 struct {
	age int
}

func (p *Person) SetAge(newAge int) {
	p.age = newAge
}

func (p *person2) SetAge(newAge int) {
	p.age = newAge
}

func CreatePerson2(newAge int) *person2 {
	p := person2{
		age: newAge,
	}
	return &p
}
