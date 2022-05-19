package metadata

import (
	"errors"
	"io"

	"github.com/PuerkitoBio/goquery"
)

// ExtractMatadata extract
func ExtractMatadata(read io.Reader, where string) (hm HTMLMeta, err error) {
	return extractMetadata(read, where)
}

func extractMetadata(read io.Reader, where string) (hm HTMLMeta, err error) {
	doc, err := goquery.NewDocumentFromReader(read)
	if err != nil {
		return
	}

	doc.Find(where).Each(func(index int, n *goquery.Selection) {
		if name, _ := n.Attr("name"); name == "description" {
			val, _ := n.Attr("content")
			hm.SetValue(val)
		}

		name, ok := n.Attr("property")
		if !ok {
			return
		}

		switch name {
		case "og:title":
			val, _ := n.Attr("content")
			hm.Title = val
		case "og:image":
			val, _ := n.Attr("content")
			hm.Image = val
		case "og:url":
			val, _ := n.Attr("content")
			hm.URL = val
		}
	})

	if hm.Title == "" {
		return HTMLMeta{}, errors.New(ErrMetadataNotFound)
	}
	if hm.Title == hm.Image || hm.Title == hm.URL {
		return HTMLMeta{}, errors.New(ErrMetadataIsHazy)
	}

	return
}
