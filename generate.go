package captcha

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func getRandomString() string {
	min, max := 7, 8
	targetLength := rand.Intn(max-min) + min
	b := make([]rune, targetLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func randomPolygon(iteration int) (n int, x, y, r, rotation float64) {
	n = rand.Intn(7-3) + 3 // defines shape (num points) of polygon
	x = rand.Float64() * float64(iteration*rand.Intn(24-8)+8)
	y = rand.Float64() * float64(iteration*rand.Intn(24-8)+8)
	r = rand.Float64() * float64(iteration*rand.Intn(24-8)+8)
	rotation = float64(rand.Intn(128-64) + 64)
	return n, x, y, r, rotation
}

func randomRGBA() (r, g, b, a int) {
	r = rand.Intn(255)
	g = rand.Intn(255)
	b = rand.Intn(255)
	a = 255
	return r, g, b, a
}

func randomFontColor() color.RGBA {
	r := uint8(rand.Intn(255-240) + 240)
	g := uint8(rand.Intn(255-240) + 240)
	b := uint8(rand.Intn(255-240) + 240)
	a := uint8(255)
	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

func randomColor() color.RGBA {
	r, g, b, a := randomRGBA()
	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}
}

func randomLinearGradient() gg.Gradient {
	x0 := float64(rand.Intn(75-25) + 25)
	y0 := float64(rand.Intn(175-125) + 125)
	x1 := float64(rand.Intn(325-275) + 275)
	y1 := float64(rand.Intn(50-10) + 10)
	grad := gg.NewLinearGradient(x0, y0, x1, y1)
	grad.AddColorStop(0, randomColor()) // colors inside gradient
	grad.AddColorStop(1, randomColor())
	grad.AddColorStop(.5, randomColor())
	return grad
}

func (c *Captcha) Generate() {
	s := time.Now()
	dc := gg.NewContext(imgWidth, imgHeight)
	grad := randomLinearGradient()

	dc.SetColor(randomFontColor()) // font color
	dc.DrawRectangle(0, 0, imgWidth, imgHeight)
	dc.Stroke()
	dc.SetFillStyle(grad)
	dc.MoveTo(0, 0)
	dc.LineTo(0, 400)
	dc.LineTo(0, 400)
	dc.LineTo(400, 0)
	dc.ClosePath()
	dc.Fill()

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size: 40,
	})
	dc.SetFontFace(face)
	w, h := dc.MeasureString(c.text)
	dc.DrawRectangle(10, 20, w, h)
	dc.Stroke()
	dc.DrawString(c.text, 25, 65)
	for i := 1; i < (rand.Intn(64-32) + 32); i++ {
		//dc.SetRGBA255(8*i, 16*1, 24*i, 255)
		//dc.DrawRegularPolygon(5, float64(i*24), float64(i*36), 300, 256)
		dc.SetRGBA255(randomRGBA())
		dc.DrawRegularPolygon(randomPolygon(i))
		dc.Stroke()
	}
	dc.SavePNG(c.File())
	dc.SavePNG("out.png")
	log.Println("finished in", time.Since(s))
}
