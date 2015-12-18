package main

import (
    "fmt"
    "net/http"
    "redis"
    "log"

)

func stop(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Secret Santa assignments:")

    spec := redis.DefaultSpec().Db(13).Password("")
    client, e := redis.NewSynchClientWithSpec(spec)
    if e != nil {
        log.Println("failed to create the client", e)
        return
	}

    keys, e := client.AllKeys()
    for i := 0; i < len(keys) - 1; i++ {
    	client.Set(keys[i+1],[]byte(keys[i]))
    }
    client.Set(keys[0], []byte(keys[len(keys) - 1]))

    for i := 0; i < len(keys); i++ {
    	val, _ := client.Get(keys[i])
    	fmt.Fprint(w,string(keys[i]))
    	fmt.Fprint(w," - ")
    	fmt.Fprintln(w,string(val[:]))
    }

    client.Flushall()

}