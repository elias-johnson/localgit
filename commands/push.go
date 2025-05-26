package commands

import (
    "fmt"
    "os"
    "path/filepath"
)

/**
 * Pushes changes from the cwd to a pseudo-remote repo.
 *
 * Creates the pseudo-remote localgit directory if it does not exist.
 */
func Push() {
    // Retrieves home directory
    home, err := os.UserHomeDir()
    if err != nil {
        fmt.Println("An error occurred while accessing home directory:", err)
        return
    }
    target := home + "/localgit"

    // Initializes remote directory if it does not exist
    _, err = os.Stat(target)
    if err != nil {
        err = os.Mkdir(target, 0755)
        if err != nil {
            fmt.Println("An error occurred while initializing remote directory:", err)
            return
        }
    }

    // Extracts the working directory
    wd, err := os.Getwd()
    if err != nil {
        fmt.Println("An error occurred while identifying working directory:", err)
        return
    }
    source := filepath.Base(wd)


    // Checks if repo already exists within remote directory
    fullTarget := target + "/" + source
    _, err = os.Stat(fullTarget)
    if err != nil {
        err = os.Mkdir(fullTarget, 0755)
        if err != nil {
            fmt.Println("An error occurred while pushing repo to remote directory:", err)
            return
        }

        // TODO - recursively copy contents from repo to remote directory
    } else {
        // TODO - check contents in remote directory before copying over changes
    }
}
