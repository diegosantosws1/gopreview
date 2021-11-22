package hashtags

import (
	"errors"
	"strings"
)

// ErrMsgEmpty message error
var ErrMsgEmpty = errors.New("the message is empty")

// ErrHashtagNotFound hashtag was not found
var ErrHashtagNotFound = errors.New("the hastag was not found")

// GetHashtags get the hashtag
func GetHashtags(msg string) (hashtags []string, err error) {
	if len(msg) == 0 {
		err = ErrMsgEmpty
		return
	}

	hashtags = searchingHashtags.FindAllString(msg, -1)
	if len(hashtags) == 0 {
		err = ErrHashtagNotFound
	}
	return
}

// GetSentenceWithoutHashtag get the hashtag
func GetSentenceWithoutHashtag(msg string) (tags []string, err error) {
	if len(msg) == 0 {
		err = ErrMsgEmpty
		return
	}

	hashtags := searchingHashtags.FindAllString(msg, -1)
	if len(hashtags) == 0 {
		err = ErrHashtagNotFound
		return
	}

	tags = make([]string, len(hashtags))
	for i, htg := range hashtags {
		tag := strings.Split(htg, "#")
		tags[i] = tag[1]
	}
	return
}
