package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func int63n(n int64) (int64, error) {
	if n < 0 {
		return 0, errors.New("Negative seed not supported")
	}
	return rand.Int63n(n), nil
}
func main() {
	//func Int63n(n int64) int64
	if newRand, err := int63n(-10); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("New Random Walue Is ", newRand)
	}
}
