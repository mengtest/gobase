package main

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/nfnt/resize"

	"golang.org/x/image/font"

	"github.com/golang/freetype"
)

const (
	dx       = 600              // 图片的大小 宽度
	dy       = 800              // 图片的大小 高度
	fontFile = "./image/mb.ttf" // 需要使用的字体文件
	fontSize = 14               // 字体尺寸
	fontDPI  = 72               // 屏幕每英寸的分辨率
	imgSize  = 80 * 1024        // 默认设置图片的大小为80KB
)

/*

 内存优化:
	1.原始png图片,只读不写, 可以在并发环境下,做为公共变量进行使用
	2.输出缓冲区, 不能作为公共变量, 原因是并发环境下,会导致数据错乱
	3.

*/
func main() {

	/************************* 可以作为公共变量,目的是为了减少内存的重复申请和销毁*******************************/
	bgfile, _ := os.Open("./image/bg.png")
	defer bgfile.Close()
	pngImg, _ := png.Decode(bgfile)
	pngImg = resize.Resize(dx, dy, pngImg, resize.Lanczos3)
	/************************* 可以作为公共变量,目的是为了减少内存的重复申请和销毁*******************************/

	// 新建一个 指定大小的 RGBA位图
	str := bytes.NewBufferString("人生若如初见,何事秋风悲画扇!")
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	draw.Draw(img, img.Bounds(), pngImg, image.ZP, draw.Src)
	bio := bytes.NewBuffer(make([]byte, 0, imgSize))
	// 读字体数据
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		log.Println(err)
		return
	}
	font1, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}
	c := freetype.NewContext()
	c.SetDPI(fontDPI)
	c.SetFont(font1)
	c.SetFontSize(fontSize)
	c.SetClip(image.Rect(0, 0, 190, 500))
	c.SetDst(img)
	c.SetSrc(image.Black)
	c.SetHinting(font.HintingNone)

	pt := freetype.Pt(10, 80) // 字出现的位置
	_, err = c.DrawString(str.String(), pt)
	if err != nil {
		log.Println(err)
		return
	}
	// 以PNG格式保存文件
	err = png.Encode(bio, img)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("./image/buffer.png", bio.Bytes(), 0666)
}
