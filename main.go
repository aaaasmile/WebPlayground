package main

import (
	"flag"
	"fmt"
	"log"
	"mime"
	"time"

	"App/web"
)

func main() {
	start := time.Now()
	var ver = flag.Bool("version", false, "program version")

	flag.Parse()
	if *ver {
		log.Println(web.AppVersion())
		return
	}
	stopch := make(chan struct{})
	msgch := make(chan string)

	go func() {
		web.StartServer(stopch)
		log.Println("server exit")
		msgch <- "OK"
	}()

	msg := <-msgch
	log.Println("terminate with: ", msg)
	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("That's all folks. (elapsed time %v)\n", elapsed)
}

func init() {
	_ = mime.AddExtensionType(".js", "text/javascript")
	_ = mime.AddExtensionType(".css", "text/css")
	fmt.Println("Mime overide")
}
