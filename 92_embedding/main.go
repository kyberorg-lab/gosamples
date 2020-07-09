package main

import "fmt"

type Person struct {
	Name string
	code string
}

type Stuff struct {
	code int
}

type Agent struct {
	Person
	Stuff
	AbleToAct bool
}

func (p Person) GetName() string {
	return p.Name
}

func main() {
	agent := Agent{Person: Person{
		Name: "AgentOne",
		code: "123456",
	}, AbleToAct: false}

	fmt.Println("Secret code", agent.GetName())
}
