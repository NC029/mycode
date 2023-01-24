package main

import (
    "fmt"
    "math/rand"
    "time"
)

// seed some entropy into the program
func init(){
    // rand.Seed(1)  // randomness has this "seed" value by default
    rand.Seed(time.Now().Unix()) 
}


func main(){

    // create a slice of strings
    instructorSlice := []string{"Ani", "Aya", "Ano", "Ari", "Ali"}

    // take a random selection from our instructorSlice
    // x := instructorSlice[rand.Intn(len(instructorSlice))]

    // x is local in scope to the if-else-if-else block
    // when we declare and initialize in the same line as the "if" keyword
    if x := instructorSlice[rand.Intn(len(instructorSlice))]; x == "Ari" {
        fmt.Println("Ari is Math instructor")
    } else if x == "Ani" { 
        fmt.Println("Ani teaches science chapters")
    } else {
        fmt.Println("Hmm, I don't know much about the instructor,", x)
    }

    // this should error, x is local to the if-else-if-else block
    //fmt.Println(x)

}
