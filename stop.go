package main

import (
    "fmt"
    "net/http"
    "redis"
    "log"
    // "encoding/json"
    // "bytes"
)


type slackMsg struct {
    Text     string `json:"text"`
    Username string `json:"username"` // Anonymous animal sender
    Channel  string `json:"channel"`  // Recipient
    Token    string `json:"token"`
}

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
    	
  //   	form := url.Values{}
  //   	form.Add("channel", string(keys[i]))
  //   	form.Add("text", string(val[:]))
  //   	form.Add("token", "xoxb-17007351888-ytHkTD0IrKyAkjM45Mu3U0Sq")
  //   	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		// payload, _ := json.Marshal(slackMsg{
	 //        Text:     string(val[:]),
	 //        Channel:  string(keys[i]),
	 //        Username: "secretsanta",
	 //        Token: "xoxb-17007351888-ytHkTD0IrKyAkjM45Mu3U0Sq",
	 //    })

  //   	http.Post("https://slack.com/api/chat.postMessage", "application.json", bytes.NewBuffer(payload))

    	fmt.Fprint(w, string(keys[i]))
    	fmt.Fprint(w," - ")
    	fmt.Fprintln(w, string(val[:]))
    }

    client.Flushall()

}