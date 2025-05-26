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
    // There are some folders intentionally omitted for simplicity
    os.Mkdir(".lit/branches", 0777)
    os.Mkdir(".lit/hooks", 0777)
    os.Mkdir(".lit/info", 0777)
    os.Mkdir(".lit/objects", 0777)
    os.Mkdir(".lit/refs", 0777)

    // Creates starter files
    // There are some files intentionally omitted for simplicity
    os.Create(".lit/HEAD")
    os.Create(".lit/info/exclude")
}
