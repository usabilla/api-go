# Gobilla

Gobilla is a Go client for [Usabilla API](https://usabilla.com/api).

[![Build Status](https://travis-ci.org/usabilla/gobilla.svg?branch=master)](https://travis-ci.org/usabilla/gobilla)

## Getting Started

After installing Go and setting up your [GOPATH](http://golang.org/doc/code.html#GOPATH), create a `main.go` file

```go
    package main

    import (
        "fmt"
        "os"

        "github.com/usabilla/gobilla"
    )

    func main() {
        args := os.Args[1:]

        // Pass the key and secret from command line arguments
        gb := gobilla.New(args[0], args[1])
        
        b := gb.Buttons()

        // Get the first ten buttons
        params := map[string]string{"limit": "10"}
        buttons, _ := b.Get(params)
        
        // Print all feedback items for each button
        for _, button := range buttons.Items {
            feedback, _ := b.Feedback().Get(button.ID, nil)
            fmt.Printf("Feedback for button: %s\n%v\n", button.ID, feedback.Items)
        }
    }
```

Then install Gobilla package

    go get github.com/usabilla/gobilla

Then run the file using your API `key` and `secret` as command line arguments

    go run main.go <your_api_key> <your_api_secret>

You will get all feedback items for each button.

The project includes a more detailed [example](example/main.go), which you can run form the root directory of the project

    go run example/main.go <your_api_key> <your_api_secret>
