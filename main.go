package main

import (
	"encoding/base64"
	"image"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/nickalie/go-webpbin"
)

func main() {
	// reading the file
	fake_kirby, err := os.Open("fake_kirby.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer fake_kirby.Close()

	// io.Reader to byte
	content, _ := ioutil.ReadAll(fake_kirby)

	// into base64
	encoded := base64.StdEncoding.EncodeToString(content)

	// decoding
	by := base64.NewDecoder(base64.StdEncoding, strings.NewReader(encoded))

	imgSrc, _, err := image.Decode(by)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("image.webp")
	if err != nil {
		log.Fatal(err)
	}
	if err := webpbin.Encode(f, imgSrc); err != nil {
		f.Close()
		log.Fatal(err)
	}
	f.Close()
}
