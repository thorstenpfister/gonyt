package nytapi

// Book review as delivered by the New York Times API.
type BookReview struct {
	URL           string   `json:"url,omitempty"`
	PublicationDt string   `json:"publication_dt,omitempty"`
	Byline        string   `json:"byline,omitempty"`
	BookTitle     string   `json:"book_title,omitempty"`
	BookAuthor    string   `json:"book_author,omitempty"`
	Summary       string   `json:"summary,omitempty"`
	UUID          string   `json:"uuid,omitempty"`
	URI           string   `json:"uri,omitempty"`
	Isbn13        []string `json:"isbn13,omitempty"`
}
