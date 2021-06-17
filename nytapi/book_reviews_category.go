package nytapi

import "fmt"

// BookReviewsCategory as defined in the New York Times API.
type BookReviewsCategory string

// Valid book reviews category as defined by the New York Times API.
const (
	Author BookReviewsCategory = "author"
	Isbn   BookReviewsCategory = "isbn"
	Title  BookReviewsCategory = "title"
)

// IsValid checks the validity of a book reviews category.
func (category BookReviewsCategory) IsValid() error {
	switch category {
	case Author, Isbn, Title:
		return nil
	}
	return fmt.Errorf("invalid book reviews category: %v", category)
}
