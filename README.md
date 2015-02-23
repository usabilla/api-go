# Gobilla

Gobilla is a Go client for Usabilla API.

# Getting Started

After installing Go and setting up your [GOPATH](http://golang.org/doc/code.html#GOPATH), create a main.go file.

    package main

    import (
        "fmt"
        "os"

        "github.com/usabilla/gobilla"
    )

    func main() {
        args := os.Args[1:]
        gb := gobilla.New(args[0], args[1])
        b := gb.Buttons()
        buttons, _ := b.Get(nil)
        for _, button := range buttons.Items {
            feedback, _ := b.Feedback().Get(button.ID, nil)
            fmt.Printf("Feedback for button: %s\n%s\n", button.ID, feedback.Items)
        }
    }

Then install Gobilla package:

    go get github.com/usabilla/gobilla

Then run the file using the key and secret as command line arguments:

    go run main.go key secret

You will get the feedback items for all the buttons.

You can find a more detailed example [here](gobilla/gobilla.go).
