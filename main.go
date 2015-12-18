package main 

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/join", join)
    http.HandleFunc("/start", start)
    http.HandleFunc("/stop", stop)
}

