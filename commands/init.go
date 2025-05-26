package commands

import (
    "fmt"
    "os"
)

/**
 * Initializes instance of localgit within the current working directory.
 *
 * Attempts to create a localgit folder within the cwd. If one does
 * not already exist, it will create one and initialize starter files within it.
 *
 * Init is limited in scope to the cwd.
 * Upon initialization, no changes will be made to the target directory.
 */
func Init() {
    // Skips reinitialization if instance already exists
    _, err := os.Stat(".localgit")
    if err == nil {
        fmt.Println("localgit has already been initialized in this directory.")
        return
    }

    // Creates localgit folder
    err = os.Mkdir(".localgit", 0755)
    if err != nil {
        fmt.Println("An error occurred while trying to initialize localgit:", err)
        return
    }

    // Creates starter folders
    os.Mkdir(".localgit/branches", 0755)
    os.Mkdir(".localgit/info", 0755)

    // Creates starter files
    os.Create(".localgit/branches/master")
    os.Create(".localgit/info/exclude")
    os.Create(".localgit/config")
    os.Create(".localgit/description")
    os.Create(".localgit/target")
}
