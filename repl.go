package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
)

type sCommand struct {
    name        string
    description string
    callback    func() error
}

func startRepl() {

    commands := getCommands()

    reader := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ")
        reader.Scan()

        words := cleanInput(reader.Text())
        if len(words) == 0 {
            continue
        }

        commandName := words[0]

        command, exists := commands[commandName]
        if exists {
            err := command.callback()
            if err != nil {
                fmt.Println(err)
            }
        } else {
            fmt.Println("Unknown command")
        }
    }
}

func cleanInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
}

func getCommands() map[string]sCommand {
    return map[string]sCommand {
        "help": {
            name:           "help",
            description:    "Displays a help message",
            callback:       commandHelp,
        },
        "map": {
            name:           "map",
            description:    "Display a list of locations",
            callback:       commandMap,
        },
        "exit": {
            name:           "exit",
            description:    "Exit the Pokedex",
            callback:       commandExit,
        },
    }
}
