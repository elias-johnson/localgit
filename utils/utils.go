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
        path := filepath.Join(wd, "/.lit")
        info, err := os.Stat(path)

        if err == nil && info.IsDir() {
            return true
        } else if wd == "/" {
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

func PathExists(path string) bool {
    _, err := os.Stat(path)

    return !os.IsNotExist(err)
}

func PathInWD(wd, path string) bool {
    // append path to wd and see if it exists
    // what happens if the input is like '/tmp'?
    // does that make the total path like /mnt/c/Users/xxx/repo//tmp? with double //
    return true
}
