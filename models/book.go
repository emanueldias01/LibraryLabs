package models

type Genre int

const (
	Fantasy Genre = iota
	ScienceFiction
	Drama
	Romance
	Adventure
)

func (g Genre) String() string {
	switch g {
	case Fantasy:
		return "Fantasy"
	case ScienceFiction:
		return "Science Fiction"
	case Drama:
		return "Drama"
	case Romance:
		return "Romance"
	case Adventure:
		return "Adventure"
	default:
		return "Unknown"
	}
}

type Book struct {
	Name string
	Author string
	Genre Genre
	Description string
	Avaliable bool
}