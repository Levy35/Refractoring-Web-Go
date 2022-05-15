//Done by: James Faber & Levi Coc

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"unicode/utf8"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/index.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
}

func (app *application) createMusicianInfo(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/music_artist.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}

}

func (app *application) createMusician(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/info", http.StatusSeeOther)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	full_name := r.PostForm.Get("full_name")
	album := r.PostForm.Get("album")
	genre := r.PostForm.Get("genre")
	date_released := r.PostForm.Get("date_released")
	artist := r.PostForm.Get("artist")

	//check the web form fields to validity
	errors := make(map[string]string)
	//check each field
	if strings.TrimSpace(genre) == "" {
		errors["genre"] = "This field cannot be left blank"
	} else if utf8.RuneCountInString(genre) > 20 {
		errors["genre"] = "This field cannot is to large(maximum is 20 characters)"
	}

	if strings.TrimSpace(full_name) == "" {
		errors["full_name"] = "This field cannot be left blank"
	} else if utf8.RuneCountInString(full_name) > 25 {
		errors["full_name"] = "This field cannot is to large(maximum is 25 characters)"
	}

	if strings.TrimSpace(album) == "" {
		errors["album"] = "This field cannot be left blank"
	} else if utf8.RuneCountInString(album) > 50 {
		errors["album"] = "This field cannot is to large(maximum is 50 characters)"
	}

	if strings.TrimSpace(date_released) == "" {
		errors["date_released"] = "This field cannot be left blank"
	} else if utf8.RuneCountInString(date_released) > 50 {
		errors["date_released"] = "This field cannot is to large(maximum is 50 characters)"
	}

	if strings.TrimSpace(artist) == "" {
		errors["artist"] = "This field cannot be left blank"
	} else if utf8.RuneCountInString(artist) > 50 {
		errors["artist"] = "This field cannot is to large(maximum is 50 characters)"
	}
	//check if there are any errrors in the map
	if len(errors) > 0 {
		ts, err := template.ParseFiles("./ui/html/music_artist.tmpl")
		if err != nil {
			log.Println(err.Error())
			http.Error(w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}
		err = ts.Execute(w, &templateData{
			ErrorsFromForm: errors,
			FormData:       r.PostForm,
		})

		if err != nil {
			log.Println(err.Error())
			http.Error(w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}

		return
	}

	//Insert music
	id, err := app.music.Insert(full_name, album, date_released, genre, artist)

	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "row ith id: %d has been inserted.", id)
}

func (app *application) displayMusician(w http.ResponseWriter, r *http.Request) {
	q, err := app.music.Read()

	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	// an instance of template data
	data := &templateData{
		MusicInfo: q,
	}

	//display the quotes using a template
	ts, err := template.ParseFiles("./ui/html/input.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, data)

	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
}
