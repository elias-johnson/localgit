package commands

import (
	"localgit/utils"
	"log"
)

/**
 *
 */
func Add(args []string) {
	// Ensures localgit has been initialized
	utils.CheckForInitialization()

	// Ensures an additional argument was provided
    if len(args) == 0 {
		log.Fatalf("Nothing specified, nothing added.")
    }
}
