package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		hander := func(w http.ResponseWriter, r *http.Request) {
			showGif(w)
		}
		http.HandleFunc("/", hander)
		http.ListenAndServe("localhost:8080", nil)
		return
	} else {
		showGif(os.Stdout)
	}

}

func showGif(out io.Writer) {
	const (
		cycles = 10    // 完整的X振荡器变化的个数
		res    = 0.001 // 角度分辨率
		size   = 100   // 图像画布包含 [-size..+size]
		nframs = 64    // 动画中的帧数
		delay  = 8     // 以10ms为单位的帧间延迟
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframs}
	phase := 0.0 // phase difference
	for i := 0; i < nframs; i++ {
		rect := image.Rect(0, 0, 2*size+4, 2*size+4)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Tan(t)
			y := math.Cos(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意 : 忽略编码错误
}
