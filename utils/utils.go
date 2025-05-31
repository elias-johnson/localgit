package utils

import (
    "log"
    "os"
    "path/filepath"
    "strings"
)

func GetWorkingDirectory() string {
    wd, err := os.Getwd()
    if err != nil {
        log.Fatalf("An error occurred while getting working directory: %v", err)
    }

    return wd
}

func CheckIfPathExistsWithinWD(path string) {
    if filepath.IsAbs(path) {
        wd := GetLocalGitDirectory()
        if !strings.HasPrefix(path, wd) {
            log.Fatalf("fatal: '%v' is outside repository at '%v'", path, wd)
        }
    }

    _, err := os.Stat(path)
    if err != nil {
        log.Fatalf("fatal: pathspec '%v' did not match any files in the localgit directory", path)
    }
}

func CheckIfLitIsInitialized() {
    if GetLocalGitDirectory() == "" {
        log.Fatal("fatal: not a localgit repository (or any of the parent directories)")
    }
}

func GetLocalGitDirectory() string {
    wd := GetWorkingDirectory()

    for {
        path := filepath.Join(wd, ".lit")
        info, err := os.Stat(path)

        if err == nil && info.IsDir() {
            return wd
        }

        if wd == "/" {
            return ""
        }

        parent := filepath.Dir(wd)
        wd = parent
    }
}
