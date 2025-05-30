package commands

import (
	"localgit/utils"
	"log"
)

/**
 * Adds changes to staging area.
 */
func Add(args []string) {
	// Ensures localgit has been initialized within this directory
	wd := utils.CheckForLitInitialization()
	if wd == "" {
		log.Fatal("fatal: not a localgit repository (or any of the parent directories)")
	}

	// Ensures an additional argument was provided
    if len(args) == 0 {
		log.Fatalf("Nothing specified, nothing added.")
    }

	// Ensures each entry exists within this directory
	for _, arg := range args {
		if !utils.PathExistsWithinWD(wd, arg) {
			log.Fatalf("fatal: pathspec '%v' did not match any files in the localgit directory", arg)
		}
	}

	// Recursively adds all entries
	for _, arg := range args {
		addEntry(arg)
	}
}

func addEntry(entry string) {
	// if entry is a file
		// make a copy of the file
		// calculate its byte count
		// prepend "blob <byte count>\0 to the copy"
		// calculate SHA1 hash of copy contents
		// compress the copy
		// put the copy in .lit/objects
		// something in .lit/index
	// if entry is a directory
		// for each of its entries
			// addEntry(entry)
}
