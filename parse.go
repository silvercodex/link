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
	nodes := linkNodes(doc)
	for _, node := range nodes {
		fmt.Println(node)
	}
	return nil, nil
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}

	return ret
}

func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
