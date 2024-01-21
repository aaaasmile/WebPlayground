package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

func handleGet(w http.ResponseWriter, req *http.Request) error {
	u, _ := url.Parse(req.RequestURI)
	log.Println("GET requested ", u)

	uristr := u.String()
	switch uristr {
	case "/main/":
		return handleIndexGetIndex(w, req)
	case "/main/posts":
		return handleGetPosts(w, req)
	case "/main/admin01/edit":
		return handleGetEdit(w, req)
	}

	return fmt.Errorf("unsupported url: %s", uristr)
}

func handleGetEdit(w http.ResponseWriter, req *http.Request) error {
	pagectx := PageCtx{
		Buildnr: buildnr,
	}
	templNameHeader := "templates/header_editor.htm"
	templNameBody := "templates/editor.htm"
	templNameFooter := "templates/footer.htm"

	if err := render(templNameHeader, templNameBody, templNameFooter, w, &pagectx); err != nil {
		return err
	}
	return nil
}

func handleIndexGetIndex(w http.ResponseWriter, req *http.Request) error {
	pagectx := PageCtx{
		Buildnr: buildnr,
	}
	templNameHeader := "templates/header.htm"
	templNameBody := "templates/index.htm"
	templNameFooter := "templates/footer.htm"

	if err := render(templNameHeader, templNameBody, templNameFooter, w, &pagectx); err != nil {
		return err
	}
	return nil
}

func handleGetPosts(w http.ResponseWriter, req *http.Request) error {
	log.Println("providing posts")
	pagectx := PageCtx{
		Buildnr: buildnr,
	}
	templNameHeader := "templates/header.htm"
	templNameBody := "templates/posts.htm"
	templNameFooter := "templates/footer.htm"

	return render(templNameHeader, templNameBody, templNameFooter, w, &pagectx)
}

func render(templNameHeader string, templNameBody string, templNameFooter string, w http.ResponseWriter, pagectx *PageCtx) error {
	tmplIndex := template.Must(template.New("App").ParseFiles(templNameHeader,
		templNameBody,
		templNameFooter))

	err := tmplIndex.ExecuteTemplate(w, "header", pagectx)
	if err != nil {
		return err
	}
	err = tmplIndex.ExecuteTemplate(w, "body", pagectx)
	if err != nil {
		return err
	}
	err = tmplIndex.ExecuteTemplate(w, "footer", pagectx)
	if err != nil {
		return err
	}
	return nil
}
