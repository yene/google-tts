/**
 * gTTS allows to use google tts on linux and osx
 *
 * g := gTTS.Speech{Path: "/tmp/gTTS", Language: "en"}
 * g.Speak("Hi")
 */

package gTTS

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

type Speech struct {
	Path     string
	Language string
}

func (speech *Speech) Speak(text string) {
	fileName := speech.Path + "/" + alphanumericOnly(text) + ".mp3"
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		speech.createFolderIfNotExists(speech.Path)
		res := speech.downloadIfNotExists(fileName, text)
		if res != nil {
			fmt.Println(res)
			return
		}
	}
	speech.play(fileName)
}

// Create the folder if does not exists.
func (speech *Speech) createFolderIfNotExists(folder string) {
	dir, err := os.Open(folder)
	if os.IsNotExist(err) {
		_ = os.MkdirAll(folder, 0700)
	}
	defer dir.Close()
}

// Download the voice file if does not exists.
func (speech *Speech) downloadIfNotExists(fileName string, text string) error {
	// Create the file
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	url := "http://translate.google.com/translate_tts?ie=UTF-8&total=1&idx=0&textlen=32&client=tw-ob&q=" + url.QueryEscape(text) + "&tl=" + speech.Language
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func alphanumericOnly(in string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	out := reg.ReplaceAllString(in, "")
	return out
}
