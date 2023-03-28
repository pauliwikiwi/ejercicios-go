package main

import "fmt"

type Article2 struct {
	Title  string
	Author string
}

func (a Article2) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

/*type Stringer interface {
	String() string
}*/

func main() {
	a := Article2{
		Title:  "Understanding Interfaces in Go",
		Author: "Sammy Shark",
	}
	Print(a)
}

func Print(s Stringer) {
	fmt.Println(s.String())
}
