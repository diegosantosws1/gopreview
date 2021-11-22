package gopreview

import (
	"io"

	"github.com/diegosantosws1/gopreview/hashtags"
	"github.com/diegosantosws1/gopreview/metadata"
)

// Parser html parser to tag meta
func Parser(url string) (res metadata.HTMLMeta, err error) {
	read, err := metadata.NewRequest(url)
	if err != nil {
		return
	}

	return metadata.ExtractMatadata(read, "meta")
}

// ParserPerReader receive the reader to extract information of metadata
func ParserPerReader(read io.Reader) (res metadata.HTMLMeta, err error) {
	return metadata.ExtractMatadata(read, "meta")
}

// ParserHashtags extract hashtag of text
func ParserHashtags(text string, sentence bool) (tags []string, err error) {
	if sentence {
		tags, err = hashtags.GetSentenceWithoutHashtag(text)
		return
	}

	tags, err = hashtags.GetHashtags(text)
	return
}
