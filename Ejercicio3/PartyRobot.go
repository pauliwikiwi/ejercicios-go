package main

import "fmt"

func main() {
	var name = "Paula"
	var age = 24
	var nextTo = "Lucia"
	var zone = "on the left"
	var meters = 23.7834298

	fmt.Println(Welcome(name))
	fmt.Println(HappyBirthday(name, age))
	fmt.Println(AssignTable(name, age, nextTo, zone, meters))
}

func Welcome(name string) string {
	return "Welcome to my party " + name
	//return fmt.Sprintf("Welcome to my party, %s!",name)
}
func HappyBirthday(name string, age int) string {
	//Sprintf -> funcion que returnea cualquier cosa en un string
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

func AssignTable(name string, age int, nextTo string, zone string, meters float64) string {
	var welcome = fmt.Sprintf("Welcome to my party %s!\n", name)
	welcome += fmt.Sprintf("You have been assigned to table %d. Your table is %s, exactly %f meters from here\n", age, zone, meters)
	welcome += fmt.Sprintf("You will be sitting next to %s", nextTo)
	return welcome
}
