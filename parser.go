package gopreview

import (
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

// ParserHashtags extract hashtag of text
func ParserHashtags(text string, sentence bool) (tags []string, err error) {
	if sentence {
		tags, err = hashtags.GetSentenceWithoutHashtag(text)
		return
	}

	tags, err = hashtags.GetHashtags(text)
	return
}
