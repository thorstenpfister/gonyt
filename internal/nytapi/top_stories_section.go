package nytapi

import "fmt"

type TopStoriesSection string

const (
	Arts         TopStoriesSection = "arts"
	Automobiles                    = "automobiles"
	Books                          = "books"
	Business                       = "business"
	Fashion                        = "fashion"
	Food                           = "food"
	Health                         = "health"
	Home                           = "home"
	Insider                        = "insider"
	Magazine                       = "magazine"
	Movies                         = "movies"
	Nyregion                       = "nyregion"
	Obituaries                     = "obituaries"
	Opinion                        = "opinion"
	Politics                       = "politics"
	Realestate                     = "realestate"
	Science                        = "science"
	Sports                         = "sports"
	Sundayreview                   = "sundayreview"
	Technology                     = "technology"
	Theater                        = "theater"
	Tmagazine                      = "t-magazine"
	Travel                         = "travel"
	Upshot                         = "upshot"
	Us                             = "us"
	World                          = "world"
)

func (section TopStoriesSection) IsValid() error {
	switch section {
	case Arts, Automobiles, Books, Business, Fashion, Food, Health, Home, Insider, Magazine, Movies, Nyregion, Obituaries, Opinion, Politics, Realestate, Science, Sports, Sundayreview, Technology, Theater, Tmagazine, Travel, Upshot, Us, World:
		return nil
	}
	return fmt.Errorf("Invalid 'Top Stories' section: %v", section)
}
