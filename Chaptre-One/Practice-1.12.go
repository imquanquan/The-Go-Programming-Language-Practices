package main

import (
	"net/http"
	"fmt"
	"log"
	"sync"
	"io"
	"math/rand"
	"image/gif"
	"image"
	"math"
	"image/color"
	"strconv"
)

var mu sync.Mutex
var count int
var palette = []color.Color{ color.RGBA{255,255,255,math.MaxUint8}, color.Black, color.RGBA{0,0,128,math.MaxUint8},color.RGBA{30,144,255,math.MaxUint8},color.RGBA{72,61,139,math.MaxUint8},color.RGBA{255,0,255,math.MaxUint8},color.RGBA{124,252,0,math.MaxUint8}}

func lissajous(out io.Writer, cycles float64) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i, c := 0, 1; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),uint8(c))
		}
		c++
		if c >= len(palette) {
			c = 1
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func Handle(w http.ResponseWriter, r *http.Request)  {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	if len(r.Form["cycles"]) != 0 {
		cycles, err := strconv.ParseFloat(r.Form["cycles"][0], 64)
		if err != nil {
			log.Println(err)
		}
		lissajous(w, cycles)
	}else {
		lissajous(w, 5.0)
	}
	count++
}

func Counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func main()  {
	http.HandleFunc("/count", Counter)
	http.HandleFunc("/", Handle)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}