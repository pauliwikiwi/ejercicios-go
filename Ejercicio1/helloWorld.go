package main

import "fmt"

const ovenTime = 40
const minPerLayer = 2

func main() {
	remainMinutes := RemainOverTime(30)
	fmt.Println(remainMinutes)
	preparationTime := PreparationTime(6)
	fmt.Println(preparationTime)
	timeInOven := ElapsedTime(5, 10)
	fmt.Println(timeInOven)
}
func RemainOverTime(actual int) int {

	return ovenTime - actual
}
func PreparationTime(numberOfLayers int) int {
	// Por cada capa tarda 2min
	return numberOfLayers * minPerLayer
}
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	// Cuantos minutos lleva en el horno y cuantas capas, y saber cuanto le queda
	return (numberOfLayers * minPerLayer) + actualMinutesInOven

}
