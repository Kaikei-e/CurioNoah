package parser

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func HTMLToDoc(inputHTMLLikeString string) (string, error) {
	doc, err := html.Parse(strings.NewReader(inputHTMLLikeString))
	if err != nil {
		return inputHTMLLikeString, fmt.Errorf("failed to parse html: %v", err)
	}

	txt := extractText(doc)

	return txt, nil
}

func extractText(n *html.Node) string {
	var text string
	if n.Type == html.TextNode {
		text = strings.TrimSpace(n.Data)
	}
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return ""
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += extractText(c)
	}
	text = strings.ReplaceAll(text, "\\n", " ")

	return text
}
