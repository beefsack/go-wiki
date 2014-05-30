package wiki

import (
	"html/template"
	"io"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

type Wiki struct {
	config    Config
	templates *template.Template
}

func NewWiki(c Config) *Wiki {
	return &Wiki{
		config: c,
	}
}

func (w *Wiki) Templates() *template.Template {
	if w.templates == nil {
		w.templates = template.Must(template.ParseGlob(
			path.Join(w.config.templatesDir(), "*")))
	}
	return w.templates
}

func (w *Wiki) RenderPage(id string, wr io.Writer) error {
	p, err := w.config.Persist.ReadPage(id)
	if err != nil {
		return err
	}
	return w.Templates().ExecuteTemplate(wr, "pageRead.html", map[string]interface{}{
		"page": p,
	})
}

func (w *Wiki) Handler(prefix string) http.Handler {
	router := mux.NewRouter().PathPrefix(prefix).Subrouter()
	router.HandleFunc("/new", w.HandlePageNew).Methods("GET")
	router.HandleFunc("/{title}/edit", w.HandlePageEdit).Methods("GET")
	router.HandleFunc("/{title}", w.HandlePageRead).Methods("GET")
	router.HandleFunc("/", w.HandleHome).Methods("GET")
	return router
}

func (w *Wiki) HandlePageNew(wr http.ResponseWriter, r *http.Request) {
	wr.WriteHeader(http.StatusOK)
	wr.Write([]byte("New wiki page"))
}

func (w *Wiki) HandlePageRead(wr http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	w.config.Persist.CreatePage(Page{
		Title: title,
		Body:  "This is a fart",
	})
	wr.WriteHeader(http.StatusOK)
	if err := w.RenderPage(title, wr); err != nil {
		panic(err)
	}
}

func (w *Wiki) HandlePageEdit(wr http.ResponseWriter, r *http.Request) {
	wr.WriteHeader(http.StatusOK)
	wr.Write([]byte("Edit wiki page"))
}

func (w *Wiki) HandleHome(wr http.ResponseWriter, r *http.Request) {
	wr.WriteHeader(http.StatusOK)
	wr.Write([]byte("Wiki"))
}
