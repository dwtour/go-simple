package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        _, _ = fmt.Fprint(w, "Welcome to my website!")
    })

    fmt.Println("start")
    _ = http.ListenAndServe(":8080", nil)
}

