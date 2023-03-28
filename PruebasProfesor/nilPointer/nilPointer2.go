package main

import "fmt"

type Creature2 struct {
	Name     string
	Greeting string
}

func (c Creature2) Greet() Creature2 {
	fmt.Printf("%s says %s!\n", c.Name, c.Greeting)
	return c
}

func (c Creature2) SayGoodbye(name string) {
	fmt.Println("Farewell", name, "!")
}

func main() {
	sammy := Creature2{
		Name:     "Sammy",
		Greeting: "Hello!",
	}
	sammy.Greet().SayGoodbye("gophers")

	Creature2.SayGoodbye(Creature2.Greet(sammy), "gophers")
}
