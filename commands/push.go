package commands

import (
    "fmt"
    "io"
    "os"
    "path/filepath"
)

/**
 * Pushes changes from the cwd to a pseudo-remote repo.
 *
 * Creates the pseudo-remote localgit directory if it does not exist.
 */
func Push() {
    // Retrieves target directory
    home, err := os.UserHomeDir()
    if err != nil {
        fmt.Println("An error occurred while accessing home directory:", err)
        return
    }
    target := home + "/.litstore"

    // Initializes remote directory if it does not exist
    _, err = os.Stat(target)
    if err != nil {
        err = os.Mkdir(target, 0755)
        if err != nil {
            fmt.Println("An error occurred while initializing remote directory:", err)
            return
        }
    }

    // Extracts the working directory
    dirPath, err := os.Getwd()
    if err != nil {
        fmt.Println("An error occurred while identifying working directory:", err)
        return
    }
    source := filepath.Base(dirPath)

    // Checks if repo already exists within remote directory
    fullTarget := target + "/" + source
    _, err = os.Stat(fullTarget)
    if err != nil {
        err = os.Mkdir(fullTarget, 0755)
        if err != nil {
            fmt.Println("An error occurred while pushing repo to remote directory:", err)
            return
        }
    }

    copyRepo(dirPath, fullTarget)
}

func copyRepo(dirPath, fullTarget string) {
    entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("An error occurred while reading a directory:", err)
		return
	}

    for _, entry := range entries {
        if !entry.IsDir() {
            // Opens source file for reading
            sourcePath := filepath.Join(dirPath, entry.Name())
            source, err := os.Open(sourcePath)
            if err != nil {
                fmt.Println("An error occurred while reading source file:", err)
                return
            }

            // Creates destination file for writing
            destinationFile := fullTarget + "/" + entry.Name()
            destination, err := os.Create(destinationFile)
            if err != nil {
                fmt.Println("An error occurred while creating destination file:", err)
                return
            }

            // Copies contents of source to destination
            _, err = io.Copy(destination, source)
            if err != nil {
                fmt.Println("An error occurred while copying a file:", err)
                return
            }
        } else {
            // Makes new directory
            destinationDir := fullTarget + "/" + entry.Name()
            err = os.Mkdir(destinationDir, 0755)
            if err != nil {
                fmt.Println("An error occurred while copying over a directory:", err)
                return
            }

            updatedDirPath := dirPath + "/" + entry.Name()
            copyRepo(updatedDirPath, fullTarget)
        }
    }
}
