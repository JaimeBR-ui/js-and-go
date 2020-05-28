package main

import "fmt"

func alternateWelcome() {
    fmt.Println("Saying hi from aux.go!")
}

// You don't need to import a package if they are part of the same package.
// However, make sure you build with all the needed go files. Ex:
// jaime$ go build main.go aux.go
