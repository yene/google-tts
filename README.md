# google-tts

## Steps
1. run binary with ./google-tts-arm
1. Play small texts by calling `http://IP:8080/?say=Hello+World`

## Raspberry Pi tips
* force always analog out with `amixer cset numid=3 1`
* Install mp3 player `sudo apt-get install mpg123`
* `env GOOS=linux GOARCH=arm GOARM=7 go build -o google-tts-arm`

## inspiration
https://github.com/pndurette/gTTS/blob/master/gtts/tts.py
sudo pip install gTTS, then gtts-cli "The good of mankind far outweighs the bad" | mpg123

## TODO
- [ ] Small web UI to create links
- [ ] port gTTS code from python over to the gTTS package
- [ ] add go releaser
- [ ] move params into cmd line params
