package main

import (
	"net/http"

	"github.com/beefsack/go-wiki"
)

func main() {
	w := wiki.NewWiki(wiki.Config{
		Persist: wiki.NewInMemory(),
	})
	http.Handle("/wiki/", w.Handler("/wiki/"))
	http.ListenAndServe(":9000", nil)
}
