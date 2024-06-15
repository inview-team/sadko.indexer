package controllers

type IndexVideoPayload struct {
	Url         string `json:"url"`
	Description string `json:"description"`
}

type VectorsPayload struct {
	Vectors []string `json:"vectors"`
}
