package commands

import (
    "log"
    "os"
)

/**
 * Initializes instance of localgit within the cwd.
 */
func Init() {
    // Skips reinitialization if instance already exists
    _, err := os.Stat(".lit")
    if err == nil {
        log.Fatal("localgit has already been initialized in this directory.")
    }

    // Creates localgit folder
    err = os.Mkdir(".lit", 0755)
    if err != nil {
        log.Fatalf("An error occurred while trying to initialize localgit:", err)
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
