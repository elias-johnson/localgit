package utils

import (
    "log"
    "os"
    "path/filepath"
)

/**
 * Checks if the current directory is in a localgit repository.
 *
 * A directory is a localgit repository if it, or any of its parent directories,
 * contain the localgit folder.
 */
func CheckForInitialization() bool {
    wd := GetWorkingDirectory()
    for {
        // Checks if localgit directory exists
        path := filepath.Join(wd, "/.lit")
        info, err := os.Stat(path)
        if err == nil && info.IsDir() {
            return true
        }

        if wd == "/" {
            break
        }

        parent := filepath.Dir(wd)
        wd = parent
    }

    log.Fatal("fatal: not a localgit repository (or any of the parent directories)")
    return false
}

func GetWorkingDirectory() string {
    wd, err := os.Getwd()
    if err != nil {
        log.Fatalf("An error occurred while getting working directory: %v", err)
    }

    return wd
}
