package person

type Person struct {
	Name string
	Age int
}

func New(name string, age int) Person {
	if age < 0 || age > 100 {
		panic("Invalid age.")
	}
	return Person {
		Name: name,
		Age: age,
	}
}