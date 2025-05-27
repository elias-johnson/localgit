package commands

import (
    "fmt"
)

/**
 *
 */
func Add(args []string) {
	// Ensures localgit has been initialized
	// TODO

	// Ensures an additional argument was provided
    if len(args) == 0 {
        fmt.Println("Nothing specified, nothing added.")
        return
    }
}
