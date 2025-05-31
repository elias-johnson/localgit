package tests

import (
    "testing"
    "localgit/commands"
)

func TestAdd(t *testing.T) {
    commands.Init()

	// TODO add a txt file or smth

	commands.Add([]string{})
}
