package main

import (
	"fmt"
	"log"
	"net"
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
	fmt.Println("http://" + getMainIP() + ":8080/?say=Hello+World")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getMainIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "localhost"
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "localhost"
}
