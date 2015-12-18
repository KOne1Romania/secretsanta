package main

func stop(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "stopping!")
}