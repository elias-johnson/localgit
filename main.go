package main

import (
	"fmt"
	"localgit/commands"
	"os"
)

func main() {
    args := os.Args[1:]

    // Displays usage if no arguments were supplied
    if len(args) == 0 {
        fmt.Println("[Display usage]")
        return
    }

    // Executes desired argument or displays usage if invalid command
    switch args[0] {
    case "init":
        commands.Init()
    case "push":
        commands.Push()
    case "help":
        // TODO
        fmt.Println("Filler")
    default:
        fmt.Printf("localgit: '%s' is not a localgit command. See `localgit help`\n", args[0])
    }
}
