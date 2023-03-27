package main

import "fmt"

func main() {
	var knightIsAwake = true
	fmt.Println(CanFastAttack(knightIsAwake))

	var knightIsAwake2 = false
	var archerIsAwake = true
	var prisonerIsAwake = false
	fmt.Println(CanSpy(knightIsAwake2, archerIsAwake, prisonerIsAwake))

	var archerIsAwake2 = false
	var prisonerIsAwake2 = true
	fmt.Println(CanSignalPrisoner(archerIsAwake2, prisonerIsAwake2))

	var knightIsAwake3 = false
	var archerIsAwake3 = true
	var prisonerIsAwake3 = false
	var petDogIsPresent = false
	fmt.Println(CanFreePrisoner(knightIsAwake3, archerIsAwake3, prisonerIsAwake3, petDogIsPresent))
}

func CanFastAttack(knightIsAwake bool) bool {
	if knightIsAwake == false {
		return true
	}
	return false
}
func CanSpy(knightIsAwake2 bool, archerIsAwake bool, prisonerIsAwake bool) bool {
	if !knightIsAwake2 && !archerIsAwake && !prisonerIsAwake {
		return false
	}
	return true
}

func CanSignalPrisoner(archerIsAwake2 bool, prisonerIsAwake2 bool) bool {
	if !archerIsAwake2 && prisonerIsAwake2 {
		return true
	}
	return false
}
func CanFreePrisoner(knightIsAwake3 bool, archerIsAwake3 bool, prisonerIsAwake3 bool, petDogIsPresent bool) bool {
	if petDogIsPresent && !archerIsAwake3 && prisonerIsAwake3 {
		return true
	} else if !petDogIsPresent && !archerIsAwake3 && !knightIsAwake3 && prisonerIsAwake3 {
		return true
	}
	return false
}
