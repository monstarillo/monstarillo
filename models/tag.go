package models

type Tag struct {
	TagName string `json:"tagName"`
	Value   string `json:"value"`
}

func NewTag(tagName, value string) Tag {
	var t Tag
	t.TagName = tagName
	t.Value = value
	return t
}
