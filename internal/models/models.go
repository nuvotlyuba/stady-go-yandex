package models

type URL string

// type ID string

// type URLData struct {
// 	ID       *string
// 	LongURL  *URL
// 	ShortURL *URL
// }

func (u URL) Point() *URL {
	return &u
}

type URLData map[URL]URL
