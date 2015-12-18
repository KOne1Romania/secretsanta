package main

import (
    "net/http"
    "log"
    "encoding/json"
    "redis"
)

type KeyVal struct {
    Name string
}

func join(w http.ResponseWriter, req *http.Request) {
	// create the Redis client
	spec := redis.DefaultSpec().Db(13).Password("")
	client, e := redis.NewSynchClientWithSpec(spec)
	if e != nil {
		log.Println("failed to create the client", e)
		return
	}
		
	decoder := json.NewDecoder(req.Body)
	reqKeyVal := KeyVal{}
	err := decoder.Decode(&reqKeyVal)
	
	if err != nil {
		log.Println(err)
	}

	log.Println("reqKeyVal = " + reqKeyVal.Name)
	
	client.Set(reqKeyVal.Name, []byte("false"))
	log.Println("Value set in Redis")
	 
}