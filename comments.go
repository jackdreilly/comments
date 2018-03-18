package main

import (
	"net/http"
	"github.com/jackdreilly/db/db"
	"encoding/json"
	"io"
	"log"
	"strings"
	"bytes"
	"os"
	"github.com/rs/cors"
)

type Comments struct {
	Comments []string
}

func main() {
	o := db.DefaultClientOptions()
	o.Port = 8083
	client, e := db.NewClient(o)
	check(e)
	log.SetOutput(os.Stdout)
	http.HandleFunc("/get", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		post_id := request.URL.Query().Get("post_id")
		log.Println("get post_id", post_id)
		comments, e := client.Get(post_id)
		if e == nil {
			io.WriteString(writer, comments)
		} else {
			json.NewEncoder(writer).Encode(&Comments{Comments:[]string{}})
		}
	})
	http.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {
		post_id := request.URL.Query().Get("post_id")
		comment := request.URL.Query().Get("comment")
		log.Println("add post_id", post_id)
		comments, e := client.Get(post_id)
		c := &Comments{}
		if e == nil {
			json.NewDecoder(strings.NewReader(comments)).Decode(c)
		}
		c.Comments = append(c.Comments, comment)
		log.Println("Num comments:", len(c.Comments))
		log.Println("Comments:", *c)
		var b bytes.Buffer
		json.NewEncoder(&b).Encode(&c)
		client.Set(post_id, b.String())
	})
	log.Fatal(http.ListenAndServe(":8092", cors.Default().Handler(http.DefaultServeMux)))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
