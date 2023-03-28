package main

import "fmt"

type Creature3 struct {
	Species string
}

func main() {
	var creature *Creature3

	fmt.Printf("1) %+v\n", creature)
	changeCreature3(creature)
	fmt.Printf("3) %+v\n", creature)
}

func changeCreature3(creature *Creature3) {
	if creature == nil {
		fmt.Println("creature is nil")
		return
	}

	creature.Species = "jellyfish"
	fmt.Printf("2) %+v\n", creature)
}
