package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var (
	mu    sync.Mutex
	count int
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer, cycles float64) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0 //発振器yの相対周波数
	//	fmt.Println(freq)
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //位相差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var cycles float64
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	cycles = 20
	for k, v := range r.Form {
		if k == "cycles" {
			c, _ := strconv.Atoi(v[0])
			fmt.Println(v[0])
			cycles = float64(c)
		}
	}
	lissajous(w, cycles)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
