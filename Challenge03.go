/*Neelima C: CHALLEGE 03 - Create a struct named Astro */

package main

import "fmt"

type Astro struct {
 name string   //name
 age  int      // age
 mission string //project mission
 isneeded bool  //true or false
}
// this is our new struct
type nasaMission struct {
        people []Astro
        number int
        message string
}

func main() {

// create 3 structs for Astro

  Astro1 :=Astro{"Alyce", 30,"MSS", true}
  Astro2 :=Astro{"Dave", 40, "CSS", false}
  Astro3 :=Astro{"Milan",43, "LSS", true}

  fmt.Println(Astro1)
  fmt.Println(Astro2)
  fmt.Println(Astro3)

      //   name   slice of Astro
    p := []Astro{Astro1, Astro2, Astro3}
    
    fmt.Println(p)
    
    // select data from the struct
    fmt.Println(p[1].mission)

    // initialize a nasaMission struct, "s"
    s := nasaMission{p, 3, "success"}
    
    // display "s" without fields
    fmt.Println(s)
    
    // display "s" with fields
    fmt.Printf("%+v", s)
}
