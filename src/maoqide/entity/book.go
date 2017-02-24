package entity

type Response struct {
	Count int    `json:"count"`
	Start int    `json:"start"`
	Total int    `json:"total"`
	Books []Book `json:"books"`
}

type Book struct {
	Rating       Rate     `json:"rating"`
	Subtitle     string   `json:"subtitle"`
	Author       []string `json:"author"`
	Pubdate      string   `json:"pubdate"`
	Tags         []Tag    `json:"tags"`
	Origin_title string   `json:"origin_title"`
	Image        string   `json:"image"`
	Binding      string   `json:"binding"`
	Translator   []string `json:"translator"`
	Catalog      string   `json:"catalog"`
	Pages        string   `json:"pages"`
	Images       []Image  `json:"images"`
	Alt          string   `json:"alt"`
	Id           string   `json:"id"`
	Publisher    string   `json:"publisher"`
	Isbn10       string   `json:"isbn10"`
	Isbn13       string   `json:"isbn13"`
	Title        string   `json:"title"`
	Url          string   `json:"url"`
	Alt_title    string   `json:"alt_title"`
	Author_intro string   `json:"author_intro"`
	Summary      string   `json:"summary"`
	Series       Series   `json:"series"`
	Price        string   `json:"price"`
	Ebook_price  string   `json:"ebook_price"`
	Ebook_url    string   `json:"ebook_url"`
}

type Rate struct {
	Max       int    `json:"max"`
	NumRaters int    `json:"numRaters"`
	Average   string `json:"average"`
	Min       int    `json:"min"`
}

type Tag struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

type Image struct {
	Small  string `json:"small"`
	Large  string `json:"large"`
	Medium string `json:"medium"`
}

type Series struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}
