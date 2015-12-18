package main

import (
    "fmt"
    "net/http"
)

func start(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "starting")
}