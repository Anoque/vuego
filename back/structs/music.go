package structs

type Artist struct {
	Id int
	Name string
}

type Track struct {
	Id int
	ArtistId int
	Title string
	Description string
}

type Genre struct {
	 Id int
	 Name string
}

type ArtistGenre struct {
	Id int
	ArtistId int
	GenreId int
}