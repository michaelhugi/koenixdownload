package mediaplayer

import (
	"embed"
	"net/http"
	"strings"
)

//go:embed audio/*
var audioFS embed.FS

//go:embed background/*
var staticFS embed.FS

func RegisterTracks() {

	// Serve embedded static files
	http.Handle("/background/", http.FileServer(http.FS(staticFS)))

	// Serve embedded audio files
	http.Handle("/audio/", http.FileServer(http.FS(audioFS)))

	http.HandleFunc("/listen/", func(w http.ResponseWriter, r *http.Request) {
		slug := strings.TrimPrefix(r.URL.Path, "/listen/")

		var selected *track
		for _, t := range tracks {
			if t.urlSlug == slug {
				selected = &t
				break
			}
		}
		if selected == nil {
			http.NotFound(w, r)
			return
		}
		if err := Page(*selected).Render(r.Context(), w); err != nil {
			println(err)
		}
	})
}
