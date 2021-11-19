package gopreview

import "github.com/DiegoSantosWS/gopreview/metadata"

// Parser html parser to tag meta
func Parser(url string) (res metadata.HTMLMeta, err error) {
	read, err := metadata.NewRequest(url)
	if err != nil {
		return
	}

	return metadata.ExtractMatadata(read, "meta")
}
