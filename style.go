package pdf

import (
	"image"
	"image/color"
)

type Style struct {
	Flow                    Flow
	BackgroundColor         color.Color
	BackgroundImage         image.Image
	Font                    string
	Size                    uint8
	Color                   color.Color
	Bold, Italic, Underline bool
	Alignment               Align
}

type Flow uint8

const (
	FlowNormal Flow = iota
	FlowAround
	FlowOver
	FlowUnder
)

type Align uint8

const (
	AlignLeft Align = iota
	AlighCenter
	AlignRight
	AlighFull
)
