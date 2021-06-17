package nytapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorstenpfister/gonyt/nytapi"
)

func Test_BookReviewsCategory_ShouldBeReflectingCategory_WithValue(t *testing.T) {
	var cases = []struct {
		category                    nytapi.BookReviewsCategory
		expectedBookReviewsCategory nytapi.BookReviewsCategory
	}{
		{nytapi.BookReviewsCategory("author"), nytapi.Author},
		{nytapi.BookReviewsCategory("isbn"), nytapi.Isbn},
		{nytapi.BookReviewsCategory("title"), nytapi.Title},
	}

	for _, tt := range cases {
		assert.Equal(t, tt.category, tt.expectedBookReviewsCategory)
	}
}

func Test_BookReviewsCategory_ShouldBeReflectingCategoryPlain_WithValue(t *testing.T) {
	var cases = []struct {
		category                    nytapi.BookReviewsCategory
		expectedBookReviewsCategory nytapi.BookReviewsCategory
	}{
		{"author", nytapi.Author},
		{"isbn", nytapi.Isbn},
		{"title", nytapi.Title},
	}

	for _, tt := range cases {
		assert.Equal(t, tt.category, tt.expectedBookReviewsCategory)
	}
}

func Test_BookReviewsCategory_ShouldBeValid(t *testing.T) {
	var cases = []struct {
		category nytapi.BookReviewsCategory
	}{
		{nytapi.Author},
		{nytapi.Isbn},
		{nytapi.Title},
	}

	for _, tt := range cases {
		assert.Nil(t, tt.category.IsValid())
	}
}

func Test_BookReviewsCategory_ShouldBeInvalid(t *testing.T) {
	var cases = []struct {
		category nytapi.BookReviewsCategory
	}{
		{nytapi.BookReviewsCategory("not-valid")},
		{"also-not-valid"},
		{"Author"},
	}

	for _, tt := range cases {
		assert.Error(t, tt.category.IsValid())
	}
}
