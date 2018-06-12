package pdf // import "vimagination.zapto.org/pdf"

import "image"

type Page struct {
	Style
	Header, Footer Block
	Contents       []interface{}
}

type Block struct {
	Style
	Contents []interface{}
}

type Inline struct {
	Style
	Contents []interface{}
}

type Text string

type Image struct {
	image.Image
}

type Link struct {
	Reference string
	Contents  interface{}
}

type Anchor struct {
	ID       string
	Contents interface{}
}
