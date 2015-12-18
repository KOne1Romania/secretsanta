package main

func start(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "starting")
}