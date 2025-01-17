package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func UpdateAge(p *Person, newAge int) {
	p.Age = newAge // Thay đổi giá trị qua con trỏ
}

func (p Person) String() string {
	return fmt.Sprintf("{name: %s, age: %d}", p.Name, p.Age)
}

func main() {
	person := Person{Name: "Đạt", Age: 25}

	fmt.Println("Tuổi ban đầu:", person) // {name: Đạt age: 25}
	UpdateAge(&person, 30)
	fmt.Println("Tuổi sau cập nhật:", person) // {name: Đạt age: 30}
}
