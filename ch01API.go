/*Neelima C: CHALLENGE 01 - Make an API request
https://pokeapi.co/api/v2/pokemon/pikachu  */


package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

func main() {

    resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/pikachu")

    
    // check for errors
    if err != nil {
        log.Fatal(err)
    }

    // the client must close the response body when finished
    defer resp.Body.Close()

    // read the content of the body with "ReadAll()"
    body, err := io.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    // print received data to the console
    fmt.Println(string(body))
}
