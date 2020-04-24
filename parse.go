package link

import "io"

//link represents a link in a HTML document
type Link struct {
	Href string
	Text string
}

//Parse will take an HTML document and return slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
