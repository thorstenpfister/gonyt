package nytapi

import "time"

type Article struct {
	Section           string       `json:"section,omitempty"`
	Subsection        string       `json:"subsection,omitempty"`
	Title             string       `json:"title,omitempty"`
	Abstract          string       `json:"abstract,omitempty"`
	Url               string       `json:"url,omitempty"`
	Uri               string       `json:"uri,omitempty"`
	Byline            string       `json:"byline,omitempty"`
	ItemType          string       `json:"item_type,omitempty"`
	UpdatedDate       time.Time    `json:"updated_date,omitempty"`
	CreatedDate       time.Time    `json:"created_date,omitempty"`
	PublishedDate     time.Time    `json:"published_date,omitempty"`
	MaterialTypeFacet string       `json:"material_type_facet,omitempty"`
	Kicker            string       `json:"kicker,omitempty"`
	DesFacet          []string     `json:"des_facet,omitempty"`
	OrgFacet          []string     `json:"org_facet,omitempty"`
	PerFacet          []string     `json:"per_facet,omitempty"`
	GeoFacet          []string     `json:"geo_facet,omitempty"`
	Multimedia        []Multimedia `json:"multimedia,omitempty"`
	ShortUrl          string       `json:"short_url,omitempty"`
}

type Multimedia struct {
	Url       string `json:"url,omitempty"`
	Format    string `json:"format,omitempty"`
	Height    int32  `json:"height,omitempty"`
	Width     int32  `json:"width,omitempty"`
	Type_     string `json:"type,omitempty"`
	Subtype   string `json:"subtype,omitempty"`
	Caption   string `json:"caption,omitempty"`
	Copyright string `json:"copyright,omitempty"`
}
