package main

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/nfnt/resize"

	"github.com/golang/freetype"
)

//
// DrawText 用于在图片照片那个绘制文字
func DrawText() {
	// 读字体数据
	fontBytes, err := ioutil.ReadFile("./image/mb.ttc")
	if err != nil {
		log.Println(err)
		return
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}
	pngFile, err := os.Open("./image/bg.png")
	if err != nil {
		fmt.Println(err)
	}
	defer pngFile.Close()
	pngImage, err := png.Decode(pngFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	pngImage = resize.Resize(600, 800, pngImage, resize.Lanczos3)
	dstImage := image.NewNRGBA(image.Rect(0, 0, 600, 800))
	draw.Draw(dstImage, dstImage.Bounds(), pngImage, image.ZP, draw.Src)
	fc := freetype.NewContext()
	fc.SetDPI(72)
	fc.SetFontSize(24)
	fc.SetFont(font)
	// //fc.SetClip(image.Rect(0, 0, 200, 200))
	// _, err = fc.DrawString("人生若如初见,何事秋风悲画扇", freetype.Pt(10, 10))
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	bio := bytes.NewBuffer(make([]byte, 10240))
	err = png.Encode(bio, dstImage)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("./image/111.png", bio.Bytes(), 0666)
}
