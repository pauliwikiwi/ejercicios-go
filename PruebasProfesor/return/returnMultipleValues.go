package main

import "fmt"

func main() {
	val, err := repeat2("Sammy", -1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}

func repeat2(word string, reps int) (string, error) {
	if reps <= 0 {
		return "", fmt.Errorf("invalid value of %d provided for reps. value must be greater than 0", reps)
	}
	var value string
	for i := 0; i < reps; i++ {
		value = value + word
	}
	return value, nil
}
