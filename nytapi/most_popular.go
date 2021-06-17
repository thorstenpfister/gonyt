package nytapi

import "fmt"

// MostPopularCategory defined by the New York Times API as either: emailed, shared, viewed
type MostPopularCategory string

// Valid most popular category as defined by the New York Times API.
const (
	Emailed MostPopularCategory = "emailed"
	Shared  MostPopularCategory = "shared"
	Viewed  MostPopularCategory = "viewed"
)

// IsValid checks the validity of a most popular category.
func (category MostPopularCategory) IsValid() error {
	switch category {
	case Emailed, Shared, Viewed:
		return nil
	}
	return fmt.Errorf("invalid most popular category: %v", category)
}

// MostPopularPeriod defined by the New York Times API as either: 1, 7, 30 (days)
type MostPopularPeriod int

// Valid most popular period as defined by the New York Times API.
const (
	Day   MostPopularPeriod = 1
	Week  MostPopularPeriod = 7
	Month MostPopularPeriod = 30
)

func (period MostPopularPeriod) IsValid() error {
	switch period {
	case Day, Week, Month:
		return nil
	}
	return fmt.Errorf("invalid most popular period: %v", period)
}
