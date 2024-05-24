package vercode

import (
	bs64 "github.com/mojocn/base64Captcha"
	"image/color"
)

var captcha *bs64.Captcha
var (
	Height          = 50
	Width           = 100
	NoiseCount      = 4
	ShowLineOptions = bs64.OptionShowHollowLine
	Length          = 4
	Source          = bs64.TxtAlphabet + bs64.TxtNumbers
	BgColor         = color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
	FontsStorage bs64.FontsStorage
	Fonts        []string
)
