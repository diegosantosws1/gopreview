package metadata

// HTMLMeta represent the
type HTMLMeta struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	URL         string `json:"url"`
}

// SetValue set value
func (val *HTMLMeta) SetValue(value string) {
	val.Description = value
	val.Image = value
	val.Title = value
	val.URL = value
}
