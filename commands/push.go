package commands

import (
    "fmt"
    "os"
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

    // Checks if remote directory exists
    _, err = os.Stat(target)
    if err != nil {
        err = os.Mkdir(target, 0755)
        if err != nil {
            fmt.Println("An error occurred while initializing remote directory:", err)
            return
        }
    }
}
