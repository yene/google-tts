package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yene/google-tts/gTTS"
)

func handler(w http.ResponseWriter, r *http.Request) {
	say, ok := r.URL.Query()["say"]
	if !ok || len(say) != 1 || len(say[0]) == 0 {
		http.Error(w, "No ?say parameter provided. Example /?say=Hello+World", http.StatusBadRequest)
		return
	}
	g.Speak(say[0])
}

var g gTTS.Speech

func main() {
	g = gTTS.Speech{Path: "/tmp/gTTS", Language: "en"}
	http.HandleFunc("/", handler)
	fmt.Println("http://localhost:8080/?say=Hello+World")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
