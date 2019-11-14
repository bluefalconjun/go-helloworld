package main

import "fmt"

const helloPrefix = "Hello, "

// Hello is to 
func Hello(in string) string {
    if in == "" {
        in = "World"
    }
    return helloPrefix + in
}

func main() {
    fmt.Println(Hello("world"))
}