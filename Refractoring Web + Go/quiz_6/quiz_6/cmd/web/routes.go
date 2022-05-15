package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/info", app.createMusicianInfo)
	mux.HandleFunc("/info-add", app.createMusician)
	mux.HandleFunc("/display", app.displayMusician)
	//create a file server to serve our static content
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mux
}
