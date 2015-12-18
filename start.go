package main

import (
    "fmt"
    "net/http"
    "redis"
    "log"
 //   "encoding/json"
  //  "net/url"
)

// type SlackStruct struct {
// 	User_name string
// 	Token string
// 	Team_id string
// 	Team_domain string
// 	Channel_id string
// 	Channel_name string
// 	User_id string
// 	Command string
// 	Text string
// 	Response_url string
// }

func start(w http.ResponseWriter, req *http.Request) {
    fmt.Fprint(w, "Secret Santa is starting now!")
    spec := redis.DefaultSpec().Db(13).Password("")
    client, e := redis.NewSynchClientWithSpec(spec)

    if e != nil {
        log.Println("failed to create the client", e)
        return
    }

    client.Flushall()

    log.Println("GET params:", req.URL.Query());

    log.Println(req.URL.Query().Get("user_name"))

}