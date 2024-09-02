package models

type URL string

func (u URL) Point() *URL {
	return &u
}

type URLData map[URL]URL

type JSONURLRequest struct {
	URL string `json:"url"`
}

type JSONURLResponse struct {
	Result string `json:"result"`
}

type ObjURL struct {
	UUID        string `json:"uuid"`
	ShortURL    URL    `json:"short_url"`
	OriginalURL URL    `json:"original_url"`
}
