package commands

import (
    "path/filepath"
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
        if entry != ".lit" {
            addEntry(entry)
        }
    }
}

func addEntry(entry string) {
    println(entry)
    info, _ := os.Stat(entry)

    if info.IsDir() {
        entries, err := os.ReadDir(entry)
        if err != nil {
            log.Fatal(err)
        }

        for _, currentEntry := range entries {
            fullEntry := filepath.Join(entry, currentEntry.Name())
            addEntry(fullEntry)
        }
    } else {
        // make a copy of the file
        // calculate its byte count
        // prepend "blob <byte count>\0 to the copy"
        // calculate SHA1 hash of copy contents
        // compress the copy
        // put the copy in .lit/objects
        // something in .lit/index
    }
}
