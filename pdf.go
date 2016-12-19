package pdf

import (
	"image"
	"image/color"
)

type Page struct {
	Width, Height   int
	Header, Footer  Block
	BackgroundColor color.Color
	BackgroundImage image.Image
}

type Flow uint8

const (
	FlowNormal Flow = iota
	FlowAround
	FlowOver
	FlowUnder
)

type Block struct {
	Flow            Flow
	BackgroundColor color.Color
	BackgroundImage image.Image
	Contents        []interface{}
}

type Align uint8

const (
	AlignLeft Align = iota
	AlighCenter
	AlignRight
	AlighFull
)

type Text struct {
	Font                    string
	Size                    uint8
	Color                   color.Color
	Bold, Italic, Underline bool
	Alignment               Align
	Data                    string
}

type Link struct {
	Reference string
	Contents  interface{}
}

type Anchor struct {
	ID       string
	Contents interface{}
}
