package nytapi

import "time"

// Article as delivered by the New York Times API.
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

// Multimedia asset belonging to an article from the New York Times API.
type Multimedia struct {
	Url       string `json:"url,omitempty"`
	Format    string `json:"format,omitempty"`
	Height    int32  `json:"height,omitempty"`
	Width     int32  `json:"width,omitempty"`
	Mediatype string `json:"type,omitempty"`
	Subtype   string `json:"subtype,omitempty"`
	Caption   string `json:"caption,omitempty"`
	Copyright string `json:"copyright,omitempty"`
}

// PopularArticle as delivered by the New York Times API.
type PopularArticle struct {
	URI           string   `json:"uri,omitempty"`
	URL           string   `json:"url,omitempty"`
	ID            int64    `json:"id,omitempty"`
	AssetID       int64    `json:"asset_id,omitempty"`
	Source        string   `json:"source,omitempty"`
	PublishedDate string   `json:"published_date,omitempty"`
	Updated       string   `json:"updated,omitempty"`
	Section       string   `json:"section,omitempty"`
	Subsection    string   `json:"subsection,omitempty"`
	Nytdsection   string   `json:"nytdsection,omitempty"`
	AdxKeywords   string   `json:"adx_keywords,omitempty"`
	Byline        string   `json:"byline,omitempty"`
	Type          string   `json:"type,omitempty"`
	Title         string   `json:"title,omitempty"`
	Abstract      string   `json:"abstract,omitempty"`
	DesFacet      []string `json:"des_facet,omitempty"`
	OrgFacet      []string `json:"org_facet,omitempty"`
	PerFacet      []string `json:"per_facet,omitempty"`
	GeoFacet      []string `json:"geo_facet,omitempty"`
	Media         []struct {
		Type                   string `json:"type,omitempty"`
		Subtype                string `json:"subtype,omitempty"`
		Caption                string `json:"caption,omitempty"`
		Copyright              string `json:"copyright,omitempty"`
		ApprovedForSyndication int    `json:"approved_for_syndication,omitempty"`
		MediaMetadata          []struct {
			URL    string `json:"url,omitempty"`
			Format string `json:"format,omitempty"`
			Height int    `json:"height,omitempty"`
			Width  int    `json:"width,omitempty"`
		} `json:"media-metadata,omitempty"`
	} `json:"media,omitempty"`
	EtaID int `json:"eta_id,omitempty"`
}
