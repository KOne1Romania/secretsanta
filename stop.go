package main

import (
    "fmt"
    "net/http"
)

func stop(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "stopping!")
}