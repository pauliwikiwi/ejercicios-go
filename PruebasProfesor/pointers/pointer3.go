package main

import "fmt"

type Creature2 struct {
	Species string
}

func main() {
	var creature Creature2 = Creature2{Species: "shark"}

	fmt.Printf("1) %+v\n", creature)
	changeCreature2(&creature)
	fmt.Printf("3) %+v\n", creature)
}

func changeCreature2(creature *Creature2) {
	creature.Species = "jellyfish"
	fmt.Printf("2) %+v\n", creature)
}
