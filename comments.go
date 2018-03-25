package main

import (
	"net/http"
	"github.com/jackdreilly/db/db"
	"encoding/json"
	"log"
	"os"
	"github.com/rs/cors"
	"flag"
	"strconv"
)

type Comments struct {
	Comments []string
}

var (
	dbPort = flag.Int("db_port", 8083, "Port to connect to db instance.")
	webPort = flag.Int("web_port", 8092, "Port for web server.")
)

func main() {
	o := db.DefaultClientOptions()
	o.Port = int32(*dbPort)
	client, e := db.NewClient(o)
	check(e)
	log.SetOutput(os.Stdout)
	http.HandleFunc("/get", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		postId := request.URL.Query().Get("post_id")
		log.Println("get post_id", postId)
		comments, e := client.GetList(postId)
		if e != nil {
			log.Println("error:", e.Error())
		}
		if comments == nil {
			comments = []string{}
		}
		json.NewEncoder(writer).Encode(&Comments{Comments:comments})
	})
	http.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {
		post_id := request.URL.Query().Get("post_id")
		comment := request.URL.Query().Get("comment")
		log.Println("add post_id", post_id)
		if e := client.Append(post_id, comment); e != nil {
			log.Println("add error:", e)
		}
	})
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(*webPort), cors.Default().Handler(http.DefaultServeMux)))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
