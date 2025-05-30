package utils

import (
    "log"
    "os"
    "path/filepath"
)

/**
 * Returns the path of the lit directory if one exists.
 */
func CheckForLitInitialization() string {
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

func PathExistsWithinWD(wd, path string) bool {
    // TODO potentially unexpected output when `path` is an absolute path
    // TODO fix -> check if path starts with `/`, treat it differently
    fullPath := filepath.Join(wd, path)
    _, err := os.Stat(fullPath)

    return !os.IsNotExist(err)
}
