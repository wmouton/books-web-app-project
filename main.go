package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    log.Println("server running on port 8080")
    http.HandleFunc("/", home)
    log.Fatalln(http.ListenAndServe(":8080", nil))
}

func home(response http.ResponseWriter, request *http.Request) {
    fmt.Fprint(response, "Welcome to the Booksite\n\n", request, string("\n"))
}
