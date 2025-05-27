package commands

import (
	"localgit/utils"
	"log"
)

/**
 * Adds changes to staging area.
 */
func Add(args []string) {
	utils.CheckForInitialization()

	// Ensures an additional argument was provided
    if len(args) == 0 {
		log.Fatalf("Nothing specified, nothing added.")
    }

	// Ensures each file/directory arg exists within cwd
	for _, arg := range args {
		if !utils.PathExists(arg) {
			log.Fatalf("fatal: pathspec '%v' did not match any files", arg)
		}

		wd := utils.GetWorkingDirectory()
		log.Printf("%v", wd)
		if !utils.PathInWD(wd, arg) {
			log.Fatalf("fatal: '%v' is outside repository at '%v'", arg, wd)
		}
	}

	// TODO go through each arg remaining in string, making sure it exists
	// If one does not exist, drop everything
	// Needs to be able to support adding directories --> recursive add
}
