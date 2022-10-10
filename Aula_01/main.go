package main

import "fmt"

func main() {
    name := "Go Developers"
    _, err := fmt.Println("Azure for", name)
    fmt.Println(err)
}