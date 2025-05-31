package utils

import (
    "bytes"
    "compress/zlib"
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

func GetByteCount(file string) int64 {
    info, err := os.Stat(file)

    if err != nil {
        log.Fatalf("An error occurred while reading a file: %v", err)
    }

    return info.Size()
}

func ReadFileContents(file string) []byte {
    fileContent, err := os.ReadFile(file)

    if err != nil {
        log.Fatalf("An error occurred while reading a file: %v", err)
    }

    return fileContent
}

func CompressFile(file []byte) {
    var compressed bytes.Buffer
    writer := zlib.NewWriter(&compressed)

    _, err := writer.Write(file)
    if err != nil {
        log.Fatalf("An error occurred while compressing a file: %v", err)
    }
    writer.Close()
}

func CreateBlobObject(fileName []byte, hash [20]byte) {
    if BlobObjectExists(hash) {
        return
    }

    // TODO strip off first two chars from hash
    // TODO then the next part

    wd := GetLocalGitDirectory()
    fullWD := filepath.Join(wd, ".lit", "objects", string(hash[:]))

    file, err := os.Create(fullWD)
    if err != nil {
        log.Fatalf("An error occurred while trying to create object: %v", err)
    }
    defer file.Close()

    _, err = file.Write(fileName)
    if err != nil {
        log.Fatalf("An error occurred while trying to write to an object: %v", err)
    }
}

func BlobObjectExists(hash [20]byte) bool {
    // TODO strip off first two chars from hash
    // TODO then the next part

    wd := GetLocalGitDirectory()
    fullWD := filepath.Join(wd, ".lit", "objects", string(hash[:]))
    _, err := os.Stat(fullWD)

    return err == nil
}
