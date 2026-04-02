package generic

import (
	"bytes"
	"net/http"
	"time"
)

type Album struct {
	Name       string
	Slug       string
	Background []byte
	Mp3Data    []byte
	WavData    []byte
	ModTime    time.Time
}

// Method to serve MP3 directly from the struct
func (a *Album) ServeMp3(w http.ResponseWriter, r *http.Request) {
	fileName := a.Name + "_mp3.zip"

	// Set header so the browser triggers a download
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)

	// http.ServeContent requires a ReadSeeker.
	// bytes.NewReader converts our []byte into a Seeker.
	content := bytes.NewReader(a.Mp3Data)

	http.ServeContent(w, r, fileName, a.ModTime, content)
}

// Method to serve WAV directly from the struct
func (a *Album) ServeWav(w http.ResponseWriter, r *http.Request) {
	fileName := a.Name + "_wav.zip"
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)

	content := bytes.NewReader(a.WavData)

	http.ServeContent(w, r, fileName, a.ModTime, content)
}

func (a *Album) ServeBackground(w http.ResponseWriter, r *http.Request) {
	// 1. Determine the correct MIME type (e.g., image/jpeg, image/png)
	// If you store the type in the struct, use that. Otherwise, detect it.
	contentType := "image/jpeg"
	w.Header().Set("Content-Type", contentType)

	// 2. Remove "Content-Disposition: attachment"
	// This ensures the browser displays the image instead of downloading it.
	w.Header().Set("Content-Disposition", "inline")

	content := bytes.NewReader(a.Background) // Assuming WavData now contains image bytes

	// 3. Serve the content
	// The filename here is used primarily for the ETag calculation
	http.ServeContent(w, r, "bg.jpg", a.ModTime, content)
}
