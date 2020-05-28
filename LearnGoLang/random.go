package main

import (
    "time"
    "math/rand"
)

func genRandDie() int8 { // this defines the return type
    // Set the seed so we get a unique random value every time
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(5) + 1 // (0 - 5) + 1
}
