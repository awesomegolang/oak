package render

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"io/ioutil"
	"path/filepath"

	"bitbucket.org/oakmoundstudio/plasticpiston/plastic/dlog"
)

var (
	fontdir = filepath.Join(
		filepath.Dir(wd),
		"assets",
		"font")

	d              *font.Drawer
	f              *truetype.Font
	defaultHinting = font.HintingNone
	defaultSize    = 12.0
	defaultDPI     = 72.0
	defaultColor   = image.White

	faceHinting     = defaultHinting
	faceSize        = defaultSize
	faceDPI         = defaultDPI
	loadedFonts     = make(map[string]*truetype.Font)
	defaultFontFile = "luxisr.ttf"
)

func InitFont(b_p *screen.Buffer) {
	b := *b_p
	LoadFont(defaultFontFile)
	f = loadedFonts[defaultFontFile]
	d = &font.Drawer{
		Dst: b.RGBA(),
		Src: defaultColor,
		Face: truetype.NewFace(f, &truetype.Options{
			Size:    defaultSize,
			DPI:     defaultDPI,
			Hinting: defaultHinting,
		}),
	}
}

func setFace() {
	d.Face = truetype.NewFace(f, &truetype.Options{
		Size:    faceSize,
		DPI:     faceDPI,
		Hinting: faceHinting,
	})
}

func DrawText(str string, x, y int) {
	d.Dot = fixed.P(x, y)
	d.DrawString(str)
}

func SetFontColor(im image.Image) {
	d.Src = im
}
func SetFontSize(fontSize float64) {
	faceSize = fontSize
	setFace()
}
func SetFontDPI(dpi float64) {
	faceDPI = dpi
	setFace()
}
func SetFontHinting(hintType string) {
	switch hintType {
	case "none":
		faceHinting = font.HintingNone
	case "vertical":
		faceHinting = font.HintingVertical
	case "full":
		faceHinting = font.HintingFull
	}
	setFace()
}
func ResetFontFormat() {
	faceHinting = defaultHinting
	faceSize = defaultSize
	faceDPI = defaultDPI
	setFace()
	d.Src = defaultColor
}

func LoadFont(fontFile string) {
	fontBytes, err := ioutil.ReadFile(filepath.Join(fontdir, fontFile))
	if err != nil {
		dlog.Error(err.Error())
		return
	}
	font, err := truetype.Parse(fontBytes)
	if err != nil {
		dlog.Error(err.Error())
		return
	}
	loadedFonts[fontFile] = font

}