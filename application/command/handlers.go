package main

import (
	"github.com/avgalaida/library/infrastructure/utilits"
	"html/template"
	"net/http"
)

type response struct {
	ID string `json:"id"`
}

func createBookCommandHandler(w http.ResponseWriter, r *http.Request) {
	title := template.HTMLEscapeString(r.FormValue("title"))
	authors := template.HTMLEscapeString(r.FormValue("authors"))

	resp := CreateBook(title, authors, r.RemoteAddr)
	util.ResponseOk(w, response{ID: resp})
}

func deleteBookCommandHandler(w http.ResponseWriter, r *http.Request) {
	id := template.HTMLEscapeString(r.FormValue("id"))

	resp := DeleteBook(id, r.RemoteAddr)
	util.ResponseOk(w, response{ID: resp})
}

func restoreBookCommandHandler(w http.ResponseWriter, r *http.Request) {
	id := template.HTMLEscapeString(r.FormValue("id"))
	status := template.HTMLEscapeString(r.FormValue("status"))

	resp := RestoreBook(id, status, r.RemoteAddr)
	util.ResponseOk(w, response{ID: resp})
}

func changeBookTitleCommandHandler(w http.ResponseWriter, r *http.Request) {
	id := template.HTMLEscapeString(r.FormValue("id"))
	title := template.HTMLEscapeString(r.FormValue("title"))

	resp := ChangeBookTitle(id, title, r.RemoteAddr)
	util.ResponseOk(w, response{ID: resp})
}

func changeBookAuthorsCommandHandler(w http.ResponseWriter, r *http.Request) {
	id := template.HTMLEscapeString(r.FormValue("id"))
	authors := template.HTMLEscapeString(r.FormValue("authors"))

	resp := ChangeBookAuthors(id, authors, r.RemoteAddr)
	util.ResponseOk(w, response{ID: resp})
}

func rollbackBookCommandHandler(w http.ResponseWriter, r *http.Request) {
	id := template.HTMLEscapeString(r.FormValue("id"))
	title := template.HTMLEscapeString(r.FormValue("title"))
	authors := template.HTMLEscapeString(r.FormValue("authors"))
	status := template.HTMLEscapeString(r.FormValue("status"))

	resp := RollBackBook(id, status, title, authors, r.RemoteAddr)
	util.ResponseOk(w, response{ID: resp})
}
