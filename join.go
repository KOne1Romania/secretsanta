package main

func join(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "joining")
}