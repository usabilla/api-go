# Usabilla

This is a Go client for [Usabilla API](https://usabilla.com/api).

[![CircleCI](https://circleci.com/gh/usabilla/api-go.svg?style=svg)](https://circleci.com/gh/usabilla/api-go)
[![GoDoc](http://godoc.org/github.com/usabilla/api-go?status.svg)](http://godoc.org/github.com/usabilla/api-go)

## Getting Started

After installing Go and setting up your [GOPATH](http://golang.org/doc/code.html#GOPATH), create a `main.go` file

```go
package main

import (
    "os"
    "fmt"

    "github.com/usabilla/api-go"
)

func main() {
    key := os.Getenv("USABILLA_API_KEY")
    secret := os.Getenv("USABILLA_API_SECRET")

    // Pass the key and secret which should be defined as ENV vars
    usabilla := usabilla.New(key, secret, nil)

    resource := usabilla.Buttons()

    // Get the first ten buttons
    params := map[string]string{"limit": "10"}
    buttons, _ := resource.Get(params)

    // Print all feedback items for each button
    for _, button := range buttons.Items {
        feedback, _ := resource.Feedback().Get(button.ID, nil)
        fmt.Printf("Feedback for button: %s\n%v\n", button.ID, feedback.Items)
    }
}
```

Then install usabilla package

    go get github.com/usabilla/api-go

Run the file

    go run main.go

And you will get all feedback items for each button.

The project includes a more detailed [example](example/main.go), which you can run from the root directory of the project

    go run example/main.go
