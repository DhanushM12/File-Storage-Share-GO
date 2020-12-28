package main

import (
    "fmt"
    "net/http"
)

func test(w http.ResponseWriter, r *http.Request){
    fmt.Printf(w, "Hello World!")
}

func main() {
    http.HandleFunc("/", test)
    http.ListenAndServe(":20201", nil)
}