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
 *
 * Returns the path that the localgit folder is within.
 * If localgit has not been initialized within the folder, returns an empty string.
 */
func CheckForInitialization() string {
    wd := GetWorkingDirectory()
    for {
        path := filepath.Join(wd, "/.lit")
        info, err := os.Stat(path)

        if err == nil && info.IsDir() {
            return wd
        } else if wd == "/" {
            break
        }

        parent := filepath.Dir(wd)
        wd = parent
    }

    return ""
}

func GetWorkingDirectory() string {
    wd, err := os.Getwd()
    if err != nil {
        log.Fatalf("An error occurred while getting working directory: %v", err)
    }

    return wd
}

// TODO potentially unexpected output when `path` is an absolute path
func PathExistsWithinWD(wd, path string) bool {
    fullPath := filepath.Join(wd, path)
    _, err := os.Stat(fullPath)

    return !os.IsNotExist(err)
}
