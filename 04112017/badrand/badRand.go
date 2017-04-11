package main

import (
	"fmt"
	"math/rand"
)

func main() {
        //func Int63n(n int64) int64
        newRand := rand.Int63n(-10)
	fmt.Println("New Random Walue Is ", newRand)
}

