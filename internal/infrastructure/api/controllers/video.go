package controllers

type IndexVideoPayload struct {
	Url         string `json:"link"`
	Description string `json:"description"`
}

type VectorsPayload struct {
	Vectors []string `json:"vectors"`
}
