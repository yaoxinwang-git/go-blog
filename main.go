package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "码神之路go博客"
	indexData.Desc = "码神之路go博客描述"
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
	//w.Write([]byte("hello world ,this is my blog app"))
}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "码神之路go博客"
	indexData.Desc = "码神之路go博客描述"
	t := template.New("index.html")

	path, _ := os.Getwd()
	t, _ = t.ParseFiles(path + "/template/index.html")
	t.Execute(w,indexData)
}
func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexHtml)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
