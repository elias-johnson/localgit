package main

import (
    "fmt"
    "localgit/commands"
    "log"
    "os"
)

func main() {
    log.SetFlags(0)
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
    case "add":
        commands.Add(args[1:])
    case "push":
        commands.Push()
    case "help":
        fmt.Println("[Display usage]")
    default:
        fmt.Printf("lit: '%s' is not a localgit command. See `lit help`\n", args[0])
    }
}
