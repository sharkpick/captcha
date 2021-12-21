package captcha

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	workspace = "/dev/shm/"
	imgWidth  = 250
	imgHeight = 120
	filename  = "%s%s.png"
)

type Captcha struct {
	text string
}

func (c *Captcha) Text() string {
	return c.text
}

func (c *Captcha) File() string {
	return fmt.Sprintf(filename, workspace, c.text)
}

func New() *Captcha {
	return &Captcha{text: getRandomString()}
}

func (c *Captcha) Cleanup() {
	if err := os.Remove(c.File()); err != nil {
		log.Println("Error cleaning up captcha -", err)
	}
}
