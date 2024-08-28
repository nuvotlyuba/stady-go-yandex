package models

type URL string

func (u URL) Point() *URL {
	return &u
}

type URLData map[URL]URL

type JsonURLRequest struct {
	URL string `json:"url"`
}

type JsonURLResponse struct {
	Result string `json:"result"`
}
