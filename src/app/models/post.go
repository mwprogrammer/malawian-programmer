package models

type PostMetaData struct {
	Essay bool
	Series bool
	SeriesName *string
}

type Post struct {
	Id int64
	Meta PostMetaData
	Title string
	Body []Paragraph
}

type Paragraph struct {
	Text *string
	Asset *Media
}

type Media struct {
	Type string
	AltText string
	Data string	
}