package main

import (
	"net/url"

	"levijames.net/test/pkg/models"
)

type templateData struct {
	MusicInfo      []*models.Music
	ErrorsFromForm map[string]string
	FormData       url.Values
}
