package main

import "fmt"

type Creature2 struct {
	Name     string
	Greeting string
}

func (c Creature2) Greet() {
	fmt.Printf("%s says %s", c.Name, c.Greeting)
}

func main() {
	sammy := Creature2{
		Name:     "Sammy",
		Greeting: "Hello!",
	}
	sammy.Greet()
}
