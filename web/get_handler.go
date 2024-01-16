package web

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

func handleGet(w http.ResponseWriter, req *http.Request) error {
	u, _ := url.Parse(req.RequestURI)
	log.Println("GET requested ", u)
	pagectx := PageCtx{
		Buildnr: buildnr,
	}
	uristr := u.String()
	if uristr == "/main/posts" {
		return handleGetPosts(w, req)
	}

	templNameHeader := "templates/header.htm"
	templNameBody := "templates/index.htm"
	templNameFooter := "templates/footer.htm"
	//templNameBase := "templates/base-lit.htm"
	//templNamePage := "templates/index-lit.htm"
	//templNameBase := "templates/base-pre.htm"
	//templNamePage := "templates/index-pre.htm"

	if err := render(templNameHeader, templNameBody, templNameFooter, w, &pagectx); err != nil {
		return err
	}
	return nil
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
