package main

import (
    "fmt"
    "net/http"
    "redis"
    "log"
    "encoding/json"
    "io"
)

type KeyVal struct {
    Name string
}

type Message struct {
	Name, Text string
}

func start(w http.ResponseWriter, req *http.Request) {
    fmt.Fprint(w, "starting")
    spec := redis.DefaultSpec().Db(13).Password("")
    client, e := redis.NewSynchClientWithSpec(spec)

    if e != nil {
        log.Println("failed to create the client", e)
        return
    }

    client.Flushall()
    client.Set("asd", []byte("val"))

    value, _ := client.Get("asd")

    key := string(value[:])

    log.Println(key)

    fmt.Fprint(w, value)

    log.Println("GET params:", req.URL.Query());

    decoder := json.NewDecoder(req.Body)

	for {
		var m Message
		if err := decoder.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}


}