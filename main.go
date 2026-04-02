package main

import (
	"koenixdownload/mediaplayer"
	"net/http"
)

// GO:EMBED directives - These lines pull the folders into the binary at compile time
//

func main() {
	mediaplayer.RegisterTracks()
	println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
