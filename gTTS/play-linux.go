// +build linux

package gTTS

import (
	"fmt"
	"os"
	"os/exec"
)

func init() {
	if _, err := os.Stat("/usr/bin/mpg123"); os.IsNotExist(err) {
		fmt.Println("mpg123 is not installed, please run: sudo apt-get install mpg123")
	}
}

// play mp3 file with mpg123 on linux
func (speech *Speech) play(fileName string) {
	cmd := "/usr/bin/mpg123 \"" + fileName + "\""
	_, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		panic(err)
	}
}
