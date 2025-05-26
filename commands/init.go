package commands

import (
    "fmt"
    "os"
)

/**
 * Initializes instance of localgit within the cwd.
 *
 * Attempts to create a localgit folder within the cwd. If one does
 * not already exist, it will create one and initialize starter files within it.
 *
 * Init is limited in scope to the cwd.
 * Upon initialization, no changes will be made to the remote directory.
 */
func Init() {
    // Skips reinitialization if instance already exists
    _, err := os.Stat(".lit")
    if err == nil {
        fmt.Println("localgit has already been initialized in this directory.")
        return
    }

    // Creates localgit folder
    err = os.Mkdir(".lit", 0755)
    if err != nil {
        fmt.Println("An error occurred while trying to initialize localgit:", err)
        return
    }

    // Creates starter folders
    os.Mkdir(".lit/branches", 0777)
    os.Mkdir(".lit/info", 0777)

    // Creates starter files
    os.Create(".lit/branches/master")
    os.Create(".lit/info/exclude")
    os.Create(".lit/config")
    os.Create(".lit/description")
}
