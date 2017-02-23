package main

import (
    "fmt"
    "io"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "OK")
}

func main() {
    fmt.Printf("[%s] Started serving 200 OK", time.Now().Format(time.RFC850))
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
