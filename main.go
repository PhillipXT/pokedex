package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    for {
        scanner := bufio.NewScanner(os.Stdin)
        fmt.Print("Pokedex > ")
        scanner.Scan()
        text := cleanInput(scanner.Text())
        fmt.Printf("Your command was: %v\n", text[0])
    }
}
