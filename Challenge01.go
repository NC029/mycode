
package main

import "fmt"

type Astro struct {
     name string   //name
     age  int      // age
     mission string //project mission
     isneeded bool  //true or false
}

func main() {

// create 3 structs for Astro

  Astro1 :=Astro{"Alyce", 30,"MSS", true}
  Astro2 :=Astro{"Dave", 40, "CSS", false}
  Astro3 :=Astro{"Milan",43, "LSS", true}

  fmt.Println(Astro1)
  fmt.Println(Astro2)
  fmt.Println(Astro3)

  // slice named people made up of Astro
    p := []Astro{Astro1, Astro2, Astro3}

    // display the slice
    fmt.Println(p)

    // select data from the struct
    fmt.Println(p[2].mission)
}
