package main

import (
	"github.com/avgalaida/library/infrastructure/utilits"
	"html/template"
	"net/http"
)

func bookListQueryHandler(w http.ResponseWriter, _ *http.Request) {
	books := GetBookList()
	util.ResponseOk(w, books)
}

func bookVersionQueryHandler(w http.ResponseWriter, r *http.Request) {
	id := template.HTMLEscapeString(r.FormValue("id"))
	version := template.HTMLEscapeString(r.FormValue("version"))

	book := GetBookWithVersion(id, version)
	util.ResponseOk(w, book)
}
