package commands

import (
	"localgit/utils"
	"log"
	"os"
)

/**
 * Adds files to staging area.
 */
func Add(entries []string) {
	utils.CheckIfLitIsInitialized()

	// Ensures an additional argument was provided
    if len(entries) == 0 {
		log.Fatalf("Nothing specified, nothing added.")
    }

	// Ensures each entry exists within this directory
	for _, entry := range entries {
		utils.CheckIfPathExistsWithinWD(entry)
	}

	// Recursively adds all entries
	for _, entry := range entries {
		addEntry(entry)
	}
}

func addEntry(entry string) {
	info, _ := os.Stat(entry)
	print(info)

	// if info.IsDir() {
	// 	entries, err := os.ReadDir(entry)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	for _, entry := range entries {
	// 		addEntry(entry.Name())
	// 	}
	// } else {

	// }

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
