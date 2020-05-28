package main

import "fmt"

// causes errors since the variables are not used and fmt is not used.
func varsEx() {
    // auto assignment of type:
    amountOfStonks := 5
    var anotherStonk = 8
    // manual assignment of type
    var moreStonks int8 = 9
    // multi var declaration
    var str1, str2 string
}

// when using println, spaces are inserted between the items printed
// when using print, no spaces are printed. - doesnt go to new line either
