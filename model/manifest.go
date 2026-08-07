package model

type Manifest []ManifestItem

type ManifestItem struct {
	Id        string
	Href      string
	MediaType string
}

