package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

//link represents a link in a HTML document
type Link struct {
	Href string
	Text string
}

//Parse will take an HTML document and return slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {

	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(doc)
	return nil, nil
}

func dfs(n *html.Node, padding string) {
	fmt.Println(padding, n.Data)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
