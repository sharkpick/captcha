package captcha

import (
	"testing"
)

var (
	captcha *Captcha
)

func TestCreate(t *testing.T) {
	captcha = New()
}

func TestGenerate(t *testing.T) {
	captcha.Generate()
	defer captcha.Cleanup()
}
