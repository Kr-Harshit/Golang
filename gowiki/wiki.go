package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {

	filename := filepath.Join("./Pages", p.Title+".txt")
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := filepath.Join("./Pages", title+".txt")
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// http handlers
var templates = template.Must(template.ParseGlob("./templates/*.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return " ", errors.New("invalid Page title")
	}
	return m[2], nil
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	pagePath := "./Pages/"
	var titles []string

	files, err := ioutil.ReadDir(pagePath)
	if err != nil {
		fmt.Println("no file present")
		log.Fatal(err)
	}
	for _, f := range files {
		titles = append(titles, strings.Replace(f.Name(), ".txt", "", 1))
	}

	err = templates.ExecuteTemplate(w, "index.html", titles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func viewhandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func edithandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func savehandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")

	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	fileServer := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fileServer))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/view/", makeHandler(viewhandler))
	http.HandleFunc("/edit/", makeHandler(edithandler))
	http.HandleFunc("/save/", makeHandler(savehandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
