package nytapi

import "fmt"

// TopStoriesSection as defined in the New York Times API.
type TopStoriesSection string

// Valid 'Top stories' section as defined by the New York Times API.
const (
	Arts         TopStoriesSection = "arts"
	Automobiles  TopStoriesSection = "automobiles"
	Books        TopStoriesSection = "books"
	Business     TopStoriesSection = "business"
	Fashion      TopStoriesSection = "fashion"
	Food         TopStoriesSection = "food"
	Health       TopStoriesSection = "health"
	Home         TopStoriesSection = "home"
	Insider      TopStoriesSection = "insider"
	Magazine     TopStoriesSection = "magazine"
	Movies       TopStoriesSection = "movies"
	Nyregion     TopStoriesSection = "nyregion"
	Obituaries   TopStoriesSection = "obituaries"
	Opinion      TopStoriesSection = "opinion"
	Politics     TopStoriesSection = "politics"
	Realestate   TopStoriesSection = "realestate"
	Science      TopStoriesSection = "science"
	Sports       TopStoriesSection = "sports"
	Sundayreview TopStoriesSection = "sundayreview"
	Technology   TopStoriesSection = "technology"
	Theater      TopStoriesSection = "theater"
	Tmagazine    TopStoriesSection = "t-magazine"
	Travel       TopStoriesSection = "travel"
	Upshot       TopStoriesSection = "upshot"
	Us           TopStoriesSection = "us"
	World        TopStoriesSection = "world"
)

// IsValid checks the validity of a 'Top stories' section.
func (section TopStoriesSection) IsValid() error {
	switch section {
	case Arts, Automobiles, Books, Business, Fashion, Food, Health, Home, Insider, Magazine, Movies, Nyregion, Obituaries, Opinion, Politics, Realestate, Science, Sports, Sundayreview, Technology, Theater, Tmagazine, Travel, Upshot, Us, World:
		return nil
	}
	return fmt.Errorf("invalid 'Top Stories' section: %v", section)
}
