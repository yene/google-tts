// +build darwin

package gTTS

import (
	"os/exec"
)

// play mp3 file with afplay on osx
func (speech *Speech) play(fileName string) {
	cmd := "/usr/bin/afplay \"" + fileName + "\""
	_, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		panic(err)
	}
}
