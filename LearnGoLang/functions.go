package main

import "fmt"

func printName(name string) bool {
    // defer makes sure the line executes right before
    // any return statements.
    defer fmt.Println("Done!")

    fmt.Println("Processing name")

    if name == "" {
        fmt.Println("No name was given")
        return false
    }
    fmt.Println(name)
    return true;
}
