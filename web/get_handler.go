package web

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

func handleGet(w http.ResponseWriter, req *http.Request) error {
	var err error
	u, _ := url.Parse(req.RequestURI)
	log.Println("GET requested ", u)
	pagectx := PageCtx{
		Buildnr: buildnr,
	}
	uristr := u.String()
	if uristr == "/main/posts" {
		return handleGetPosts(w, req)
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

func handleGetPosts(w http.ResponseWriter, req *http.Request) error {
	log.Println("providing posts")
	pagectx := PageCtx{
		Buildnr: buildnr,
	}
	templNameBase := "templates/base.htm"
	templNamePage := "templates/posts.htm"
	tmplIndex := template.Must(template.New("App").ParseFiles(templNameBase, templNamePage))
	err := tmplIndex.ExecuteTemplate(w, "base", pagectx)
	if err != nil {
		return err
	}
	err = tmplIndex.ExecuteTemplate(w, "body", pagectx)
	if err != nil {
		return err
	}
	return nil
}
