package main

import "fmt"

type Article4 struct {
	Title  string
	Author string
}

func (a Article4) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) String() string {
	return fmt.Sprintf("The %q book was written by %s.", b.Title, b.Author)
}

type Stringer interface {
	String() string
}

func main() {
	a := Article4{
		Title:  "Understanding Interfaces in Go",
		Author: "Sammy Shark",
	}
	Print2(a)

	b := Book{
		Title:  "All About Go",
		Author: "Jenny Dolphin",
		Pages:  25,
	}
	Print2(b)
}

func Print2(s Stringer) {
	fmt.Println(s.String())
}
