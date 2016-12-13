package base

import (
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

// BinSchedule structures the information returned from the external service
type BinSchedule struct {
	Postcode      string
	RecyclingDate string
	RubbishDate   string
}

// Maintain a splice of the data from the response html
var headers, em []string

// traverse will walk the node graph extracting the information we want from the repsonse
func traverse(n *html.Node) {

	if isH2Tag(n) {
		headers = append(headers, extractType(n))
	}

	// the date information is held between the <em> tags
	if isEmTag(n) {
		em = append(em, extractDate(n))
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(c)
	}
}

func isH2Tag(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "h2"
}

func isEmTag(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "em"
}

// extractDate is a function which extracts the date from between the <em> tags in the response
func extractDate(n *html.Node) string {
	return strings.Join([]string{n.FirstChild.Data, n.LastChild.Data}, " ")
}

// extractType is a function which extracts the type of bin from between the <h2> tags in the response
func extractType(n *html.Node) string {
	return n.FirstChild.Data
}

// ResultParser takes a io.Reader and parses the dates for both Rubbish and Recycling
func ResultParser(r io.Reader) BinSchedule {
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	traverse(doc)

	output := BinSchedule{Postcode: "EX5 3DX", RecyclingDate: em[0], RubbishDate: em[1]}

	return output
}
