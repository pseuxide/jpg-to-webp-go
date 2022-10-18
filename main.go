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
	// 画像を読み込み
	fake_kirby, err := os.Open("fake_kirby.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer fake_kirby.Close()

	// io.Readerからbyteに変換
	content, _ := ioutil.ReadAll(fake_kirby)

	// base64に変換
	encoded := base64.StdEncoding.EncodeToString(content)

	// デコード
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
