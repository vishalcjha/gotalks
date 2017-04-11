package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var sameError = "Why Error Has Error"

func giveError() error {
	if rand.Int63n(int64(time.Now().Second()))%2 == 0 {
		return errors.New(sameError)
	}
	return errors.New(sameError)
}
func main() {
	if err := giveError(); err != nil {
		fmt.Println(err.Error())
	}
}
