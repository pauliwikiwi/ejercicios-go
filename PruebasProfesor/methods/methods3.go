package main

import "fmt"

type Creature3 struct {
	Name     string
	Greeting string
}

func (c Creature3) Greet() Creature3 {
	fmt.Printf("%s says %s!\n", c.Name, c.Greeting)
	return c
}

func (c Creature3) SayGoodbye(name string) {
	fmt.Println("Farewell", name, "!")
}

func main() {
	sammy := Creature3{
		Name:     "Sammy",
		Greeting: "Hello!",
	}
	sammy.Greet().SayGoodbye("gophers")

	Creature3.SayGoodbye(Creature3.Greet(sammy), "gophers")
}
