package main

import (
    "net/http"
    "log"
    "redis"
    "fmt"
)

type KeyVal struct {
    Name string
}

func join(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "You just joined Secret Santa!")
	// create the Redis client
	spec := redis.DefaultSpec().Db(13).Password("")
	client, e := redis.NewSynchClientWithSpec(spec)
	if e != nil {
		log.Println("failed to create the client", e)
		return
	}
		
	client.Set(req.URL.Query().Get("user_name"), []byte(""))
	log.Println("Value set in Redis")
	 
}