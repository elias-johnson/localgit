package tests

import (
    "testing"
    "localgit/commands"
    "os"
)

func TestInit(t *testing.T) {
    commands.Init()

    _, err := os.Stat(".lit")
    if err != nil {
        t.Errorf("Init did not create a localgit directory.")
    }

    _, err = os.Stat(".lit/branches")
    if err != nil {
        t.Errorf("Init did not create a branches directory within the localgit directory.")
    }

    _, err = os.Stat(".lit/hooks")
    if err != nil {
        t.Errorf("Init did not create a hooks directory within the localgit directory.")
    }

    _, err = os.Stat(".lit/info")
    if err != nil {
        t.Errorf("Init did not create an info directory within the localgit directory.")
    }

    _, err = os.Stat(".lit/objects")
    if err != nil {
        t.Errorf("Init did not create an objects directory within the localgit directory.")
    }

    _, err = os.Stat(".lit/refs")
    if err != nil {
        t.Errorf("Init did not create a refs directory within the localgit directory.")
    }

    _, err = os.Stat(".lit/HEAD")
    if err != nil {
        t.Errorf("Init did not create a HEAD file within the localgit directory.")
    }

    _, err = os.Stat(".lit/info/exclude")
    if err != nil {
        t.Errorf("Init did not create an exclude file within the info directory.")
    }
}
