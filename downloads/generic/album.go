package generic

import (
	"bytes"
	"net/http"
	"time"
)

type Album struct {
	Name           string
	Slug           string
	BackgroundFile string
	Mp3Data        []byte
	WavData        []byte
	ModTime        time.Time
}

func (a *Album) ServeMp3(w http.ResponseWriter, r *http.Request) {
	fileName := a.Name + "_mp3.zip"

	// Set header so the browser triggers a download
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)

	// http.ServeContent requires a ReadSeeker.
	// bytes.NewReader converts our []byte into a Seeker.
	content := bytes.NewReader(a.Mp3Data)

	http.ServeContent(w, r, fileName, a.ModTime, content)
}

func (a *Album) ServeWav(w http.ResponseWriter, r *http.Request) {
	fileName := a.Name + "_wav.zip"
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)

	content := bytes.NewReader(a.WavData)

	http.ServeContent(w, r, fileName, a.ModTime, content)
}
