package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"text/template"
	"time"
)

const (
	buildnr = "00.20231227.01.00"
)

type PageCtx struct {
	Buildnr string
}

func AppVersion() string {
	return fmt.Sprintf("Program version: %s", buildnr)
}

func handleGet(w http.ResponseWriter, req *http.Request) error {
	var err error
	u, _ := url.Parse(req.RequestURI)
	log.Println("GET requested ", u)
	pagectx := PageCtx{
		Buildnr: buildnr,
	}

	templNameBase := "templates/base.htm"
	templNamePage := "templates/index.htm"
	//templNameBase := "templates/base-lit.htm"
	//templNamePage := "templates/index-lit.htm"
	//templNameBase := "templates/base-pre.htm"
	//templNamePage := "templates/index-pre.htm"

	tmplIndex := template.Must(template.New("App").ParseFiles(templNameBase, templNamePage))

	err = tmplIndex.ExecuteTemplate(w, "base", pagectx)
	if err != nil {
		return err
	}
	err = tmplIndex.ExecuteTemplate(w, "body", pagectx)
	if err != nil {
		return err
	}

	return nil
}

func apiHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		if err := handleGet(w, req); err != nil {
			log.Println("Error on process request: ", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	case "POST":
		if err := handlePost(w, req); err != nil {
			log.Println("Error on process request: ", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func StartServer(cr <-chan struct{}) {
	serverurl := "127.0.0.1:5904"
	rootURLPattern := "/main/"
	finalServURL := fmt.Sprintf("http://%s%s", strings.Replace(serverurl, "0.0.0.0", "localhost", 1), rootURLPattern)

	finalServURL = strings.Replace(finalServURL, "127.0.0.1", "localhost", 1)
	log.Println("Server started with URL ", serverurl)
	log.Println("Try this url: ", finalServURL)

	http.Handle(rootURLPattern+"static/", http.StripPrefix(rootURLPattern+"static", http.FileServer(http.Dir("static"))))
	http.HandleFunc(rootURLPattern, apiHandler)

	srv := &http.Server{
		Addr:         serverurl,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      nil,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println("Server is not listening anymore, error: ", err)
		}
	}()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt) //We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	log.Println("Enter in server loop")
loop:
	for {
		select {
		case <-sig:
			log.Println("stop because interrupt")
			break loop
		case <-cr:
			log.Println("stop because service shutdown")
			break loop
		}
	}
	var wait time.Duration
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Bye, service")
}
