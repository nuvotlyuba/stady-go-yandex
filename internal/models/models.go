package models

type URL string

func (u URL) Point() *URL {
	return &u
}

type URLData map[URL]URL
