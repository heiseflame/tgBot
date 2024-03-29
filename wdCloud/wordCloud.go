package group

import (
	"bytes"
	"fmt"

	// "fmt"
	"image/color"
	"image/png"

	// "os"
	// "time"

	wordclouds "github.com/jizizr/WCloud"
)

var DefaultColors = []color.RGBA{
	{27, 27, 27, 255},
	{95, 174, 227, 255},
	{123, 104, 238, 255},
	{60, 179, 113, 255},
	// {0x70, 0xD6, 0xBF, 0xff},
}

var boxes = wordclouds.Mask(
	"source/mask.png",
	800,
	800,
	color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 0,
	})

var oarr []wordclouds.Option
var colors = make([]color.Color, 0)

func init() {
	for _, c := range DefaultColors {
		colors = append(colors, c)
	}
	oarr = []wordclouds.Option{
		wordclouds.FontFile("source/font.ttf"),
		wordclouds.FontMaxSize(150),
		wordclouds.FontMinSize(20),
		wordclouds.Colors(colors),
		wordclouds.MaskBoxes(boxes),
		wordclouds.Height(800),
		wordclouds.Width(800),
		wordclouds.RandomPlacement(false),
		wordclouds.WordSizeFunction("linear"),
		wordclouds.CopyrightFontSize(10),
		wordclouds.CopyrightString(""),
	}
}

func Rank(inputWords map[string]int, name string) []byte {

	// Load words
	// inputWords := map[string]int{"消息": 42, "是啊": 30, "中文": 15, "也是": 10, "而我": 5, "撒旦": 11, "落后": 11}

	// Load config

	// start := time.Now()
	oarr[10] = wordclouds.CopyrightString(fmt.Sprintf("by %s", name))
	w := wordclouds.NewWordcloud(inputWords,
		oarr...,
	)

	// outputFile, _ := os.Create("test.png")
	buf := new(bytes.Buffer)
	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(buf, w.Draw())

	// Don't forget to close files
	// outputFile.Close()
	// fmt.Printf("Done in %v\n", time.Since(start))
	return buf.Bytes()
}
