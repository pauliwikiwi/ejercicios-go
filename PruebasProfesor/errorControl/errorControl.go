package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	err := errors.New("barnacles")
	fmt.Println("Sammy says:", err)

	err2 := fmt.Errorf("error occurred at: %v", time.Now())
	fmt.Println("An error happened:", err2)

}
