package generic

import (
	"fmt"
	"net/http"
	"slices"
	"strings"
)

func Register(album *Album) {
	http.HandleFunc(album.Slug, func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		if code != "" {
			http.Redirect(w, r, album.Slug+"/"+code, http.StatusMovedPermanently)
			return
		}
		if err := page(album, r.URL.Query().Get("err")).Render(r.Context(), w); err != nil {
			println(err)
		}
	})
}

func RegisterDownload(album *Album, codes []string) {
	slug := fmt.Sprintf("%s/", album.Slug)
	http.HandleFunc(slug, func(w http.ResponseWriter, r *http.Request) {
		code := strings.TrimPrefix(r.URL.Path, slug)
		if code == "bg.jpg" {
			album.ServeBackground(w, r)
			return
		}
		if strings.HasPrefix(code, "mp3/") {
			code = strings.TrimPrefix(code, "mp3/")

			if !slices.Contains(codes, code) {
				http.Redirect(w, r, album.Slug+"?err=Invalid code", http.StatusMovedPermanently)
				return
			}
			album.ServeMp3(w, r)
			return
		}
		if strings.HasPrefix(code, "wav/") {
			code = strings.TrimPrefix(code, "wav/")
			if !slices.Contains(codes, code) {
				http.Redirect(w, r, album.Slug+"?err=Invalid code", http.StatusMovedPermanently)
				return
			}
			album.ServeWav(w, r)
			return
		}
		if !slices.Contains(codes, code) {
			http.Redirect(w, r, album.Slug+"?err=Invalid code", http.StatusMovedPermanently)
			return
		}
		if err := formatSelection(album, code).Render(r.Context(), w); err != nil {
			println(err)
		}

	})
}
