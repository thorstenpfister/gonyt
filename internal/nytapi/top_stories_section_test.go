package nytapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorstenpfister/gonyt/internal/nytapi"
)

func Test_TopStoriesSection_ShouldBeReflectingSection_WithValue(t *testing.T) {
	var cases = []struct {
		section                   nytapi.TopStoriesSection
		expectedTopStoriesSection nytapi.TopStoriesSection
	}{
		{nytapi.TopStoriesSection("arts"), nytapi.Arts},
		{nytapi.TopStoriesSection("books"), nytapi.Books},
		{nytapi.TopStoriesSection("fashion"), nytapi.Fashion},
		{nytapi.TopStoriesSection("food"), nytapi.Food},
		{nytapi.TopStoriesSection("health"), nytapi.Health},
		{nytapi.TopStoriesSection("home"), nytapi.Home},
		{nytapi.TopStoriesSection("insider"), nytapi.Insider},
		{nytapi.TopStoriesSection("magazine"), nytapi.Magazine},
		{nytapi.TopStoriesSection("movies"), nytapi.Movies},
		{nytapi.TopStoriesSection("nyregion"), nytapi.Nyregion},
		{nytapi.TopStoriesSection("obituaries"), nytapi.Obituaries},
		{nytapi.TopStoriesSection("opinion"), nytapi.Opinion},
		{nytapi.TopStoriesSection("politics"), nytapi.Politics},
		{nytapi.TopStoriesSection("realestate"), nytapi.Realestate},
		{nytapi.TopStoriesSection("science"), nytapi.Science},
		{nytapi.TopStoriesSection("sports"), nytapi.Sports},
		{nytapi.TopStoriesSection("sundayreview"), nytapi.Sundayreview},
		{nytapi.TopStoriesSection("technology"), nytapi.Technology},
		{nytapi.TopStoriesSection("theater"), nytapi.Theater},
		{nytapi.TopStoriesSection("t-magazine"), nytapi.Tmagazine},
		{nytapi.TopStoriesSection("travel"), nytapi.Travel},
		{nytapi.TopStoriesSection("upshot"), nytapi.Upshot},
		{nytapi.TopStoriesSection("us"), nytapi.Us},
		{nytapi.TopStoriesSection("world"), nytapi.World},
	}

	for _, tt := range cases {
		assert.Equal(t, tt.section, tt.expectedTopStoriesSection)
	}
}

func Test_TopStoriesSection_ShouldBeReflectingSectionPlain_WithValue(t *testing.T) {
	var cases = []struct {
		section                   nytapi.TopStoriesSection
		expectedTopStoriesSection nytapi.TopStoriesSection
	}{
		{"arts", nytapi.Arts},
		{"books", nytapi.Books},
		{"fashion", nytapi.Fashion},
		{"food", nytapi.Food},
		{"health", nytapi.Health},
		{"home", nytapi.Home},
		{"insider", nytapi.Insider},
		{"magazine", nytapi.Magazine},
		{"movies", nytapi.Movies},
		{"nyregion", nytapi.Nyregion},
		{"obituaries", nytapi.Obituaries},
		{"opinion", nytapi.Opinion},
		{"politics", nytapi.Politics},
		{"realestate", nytapi.Realestate},
		{"science", nytapi.Science},
		{"sports", nytapi.Sports},
		{"sundayreview", nytapi.Sundayreview},
		{"technology", nytapi.Technology},
		{"theater", nytapi.Theater},
		{"t-magazine", nytapi.Tmagazine},
		{"travel", nytapi.Travel},
		{"upshot", nytapi.Upshot},
		{"us", nytapi.Us},
		{"world", nytapi.World},
	}

	for _, tt := range cases {
		assert.Equal(t, tt.section, tt.expectedTopStoriesSection)
	}
}

func Test_TopStoriesSection_ShouldBeValid(t *testing.T) {
	var cases = []struct {
		section nytapi.TopStoriesSection
	}{
		{nytapi.Arts},
		{nytapi.Books},
		{nytapi.Fashion},
		{nytapi.Food},
		{nytapi.Health},
		{nytapi.Home},
		{nytapi.Insider},
		{nytapi.Magazine},
		{nytapi.Movies},
		{nytapi.Nyregion},
		{nytapi.Obituaries},
		{nytapi.Opinion},
		{nytapi.Politics},
		{nytapi.Realestate},
		{nytapi.Science},
		{nytapi.Sports},
		{nytapi.Sundayreview},
		{nytapi.Technology},
		{nytapi.Theater},
		{nytapi.Tmagazine},
		{nytapi.Travel},
		{nytapi.Upshot},
		{nytapi.Us},
		{nytapi.World},
	}

	for _, tt := range cases {
		assert.Nil(t, tt.section.IsValid())
	}
}

func Test_TopStoriesSection_ShouldBeInvalid(t *testing.T) {
	var cases = []struct {
		section nytapi.TopStoriesSection
	}{
		{nytapi.TopStoriesSection("not-valid")},
		{"also-not-valid"},
		{"Arts"},
	}

	for _, tt := range cases {
		assert.Error(t, tt.section.IsValid())
	}
}
