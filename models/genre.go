package models


func (b *Book)SelectGenre(g string){

	var genre string
		switch g {
		case "Fantasy":
			genre =  "Fantasy"
		case "ScienceFiction":
			 genre = "Science Fiction"
		case "Drama":
			genre = "Drama"
		case "Romance":
			genre = "Romance"
		case "Adventure":
			genre = "Adventure"
		default:
			genre = "Unknown"
		}

		b.Genre = genre
}