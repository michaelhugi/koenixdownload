package main

import (
	"embed"
	"koenixdownload/downloads"
	"koenixdownload/mediaplayer"
	"net/http"
)

// GO:EMBED directives - These lines pull the folders into the binary at compile time
//

//go:embed style/*
var styleFs embed.FS

func main() {
	http.Handle("/style/", http.FileServer(http.FS(styleFs)))

	mediaplayer.RegisterTracks()
	downloads.RegisterDownloads()
	println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
