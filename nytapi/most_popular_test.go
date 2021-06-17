package nytapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorstenpfister/gonyt/nytapi"
)

func Test_MostPopularCategory_ShouldBeReflectingCategory_WithValue(t *testing.T) {
	var cases = []struct {
		category                    nytapi.MostPopularCategory
		expectedMostPopularCategory nytapi.MostPopularCategory
	}{
		{nytapi.MostPopularCategory("emailed"), nytapi.Emailed},
		{nytapi.MostPopularCategory("shared"), nytapi.Shared},
		{nytapi.MostPopularCategory("viewed"), nytapi.Viewed},
	}

	for _, tt := range cases {
		assert.Equal(t, tt.category, tt.expectedMostPopularCategory)
	}
}

func Test_MostPopularCategory_ShouldBeReflectingCategoryPlain_WithValue(t *testing.T) {
	var cases = []struct {
		category                    nytapi.MostPopularCategory
		expectedMostPopularCategory nytapi.MostPopularCategory
	}{
		{"emailed", nytapi.Emailed},
		{"shared", nytapi.Shared},
		{"viewed", nytapi.Viewed},
	}

	for _, tt := range cases {
		assert.Equal(t, tt.category, tt.expectedMostPopularCategory)
	}
}

func Test_MostPopularCategory_ShouldBeValid(t *testing.T) {
	var cases = []struct {
		category nytapi.MostPopularCategory
	}{
		{nytapi.Emailed},
		{nytapi.Shared},
		{nytapi.Viewed},
	}

	for _, tt := range cases {
		assert.Nil(t, tt.category.IsValid())
	}
}

func Test_MostPopularCategory_ShouldBeInvalid(t *testing.T) {
	var cases = []struct {
		category nytapi.MostPopularCategory
	}{
		{nytapi.MostPopularCategory("not-valid")},
		{"also-not-valid"},
		{"Emailed"},
	}

	for _, tt := range cases {
		assert.Error(t, tt.category.IsValid())
	}
}

func Test_MostPopularPeriod_ShouldBeReflectingPeriod_WithValue(t *testing.T) {
	var cases = []struct {
		period                    nytapi.MostPopularPeriod
		expectedMostPopularPeriod nytapi.MostPopularPeriod
	}{
		{nytapi.MostPopularPeriod(1), nytapi.Day},
		{nytapi.MostPopularPeriod(7), nytapi.Week},
		{nytapi.MostPopularPeriod(30), nytapi.Month},
	}

	for _, tt := range cases {
		assert.Equal(t, tt.period, tt.expectedMostPopularPeriod)
	}
}

func Test_MostPopularPeriod_ShouldBeReflectingPeriodPlain_WithValue(t *testing.T) {
	var cases = []struct {
		period                    nytapi.MostPopularPeriod
		expectedMostPopularPeriod nytapi.MostPopularPeriod
	}{
		{1, nytapi.Day},
		{7, nytapi.Week},
		{30, nytapi.Month},
	}

	for _, tt := range cases {
		assert.Equal(t, tt.period, tt.expectedMostPopularPeriod)
	}
}

func Test_MostPopularPeriod_ShouldBeValid(t *testing.T) {
	var cases = []struct {
		period nytapi.MostPopularPeriod
	}{
		{nytapi.Day},
		{nytapi.Week},
		{nytapi.Month},
	}

	for _, tt := range cases {
		assert.Nil(t, tt.period.IsValid())
	}
}

func Test_MostPopularPeriod_ShouldBeInvalid(t *testing.T) {
	var cases = []struct {
		period nytapi.MostPopularPeriod
	}{
		{nytapi.MostPopularPeriod(100)},
	}

	for _, tt := range cases {
		assert.Error(t, tt.period.IsValid())
	}
}
