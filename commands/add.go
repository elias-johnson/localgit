package commands

import (
    "fmt"
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
        byteCount := utils.GetByteCount(entry)
        fileContent := utils.ReadFileContents(entry)

        // Prepends metadata to the file content
        prefix := fmt.Sprintf("blob %v\x00", byteCount)
        modifiedFileContent := append([]byte(prefix), fileContent...)

        // Calculates the SHA1 hash of the file content
        // TODO

        // Compresses the file content
        // TODO

        // Creates new blob object if it does not already exist
        // TODO

        // Updates staging file
        // TODO
        // something in .lit/index
    }
}
