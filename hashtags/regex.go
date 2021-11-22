package hashtags

import "regexp"

// SearchingHashtags regex to extract hashtags
var searchingHashtags = regexp.MustCompile(`#(\w+)`)
