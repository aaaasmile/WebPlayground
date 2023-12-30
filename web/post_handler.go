package web

import (
	"log"
	"net/http"
	"net/url"
)

func handlePost(w http.ResponseWriter, req *http.Request) error {

	u, _ := url.Parse(req.RequestURI)
	log.Println("POST requested ", u)
	w.Write([]byte("<h1> Hello from Server</h1>"))
	return nil
}
