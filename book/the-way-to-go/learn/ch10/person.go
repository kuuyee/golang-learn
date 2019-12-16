package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)

}

func main() {
	//1. 结构作为值类型
	var p1 Person
	p1.firstName = "Jiang"
	p1.lastName = "Xuguang"
	upPerson(&p1)
	fmt.Printf("The name of the person is %s %s \n", p1.firstName, p1.lastName)

	//2. 结构作为指针
	p2 := new(Person)
	fmt.Printf("p2的类型是 %T , *p2的类型是 %T \n", p2, (*p2))
	p2.firstName = "Kuu"
	p2.lastName = "Yee"
	(*p2).lastName = "Gz" //这是合法的
	p2.firstName = "Jiang"
	upPerson(p2)
	fmt.Printf("The name of the person is %s %s \n", p2.firstName, p2.lastName)

	//3. 结构作为字面量
	p3 := &Person{"Kuu", "Yee"}
	upPerson(p3)
	fmt.Printf("The name of the person is %s %s \n", p3.firstName, p3.lastName)

}
