package main

import (
    "fmt"
    "net/http"
)

func join(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "joining")
}