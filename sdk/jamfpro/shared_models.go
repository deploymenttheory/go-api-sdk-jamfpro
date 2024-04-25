// shared_models.go
package jamfpro

// Shared Standalone Resources

type SharedResourceSite struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

type SharedResourceSiteProAPI struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SharedResourceCategory struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

type SharedResourceSelfServiceIcon struct {
	ID       int    `json:"id,omitempty" xml:"id,omitempty"`
	URI      string `json:"uri,omitempty" xml:"uri,omitempty"`
	Data     string `json:"data,omitempty" xml:"data,omitempty"`
	Filename string `json:"filename,omitempty" xml:"filename,omitempty"`
}

type SharedResourceSelfServiceCategories struct {
	Category []SharedResourceSelfServiceCategory
}

type SharedResourceSelfServiceCategory struct {
	ID       int    `json:"id,omitempty" xml:"id,omitempty"`
	Name     string `json:"name,omitempty" xml:"name,omitempty"`
	Priority int    `json:"priority,omitempty" xml:"priority,omitempty"`
}

// Advanced Search DisplayField
type SharedAdvancedSearchContainerDisplayField struct {
	Size         int                                      `json:"size,omitempty" xml:"size,omitempty"`
	DisplayField []SharedAdvancedSearchSubsetDisplayField `json:"display_field,omitempty" xml:"display_field,omitempty"`
}

type SharedAdvancedSearchSubsetDisplayField struct {
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// Criteria

type SharedContainerCriteria struct {
	Size      int                    `json:"size,omitempty" xml:"size,omitempty"`
	Criterion []SharedSubsetCriteria `json:"criterion,omitempty" xml:"criterion,omitempty"`
}

type SharedSubsetCriteria struct {
	Name         string `json:"name,omitempty" xml:"name,omitempty"`
	Priority     int    `json:"priority,omitempty" xml:"priority,omitempty"`
	AndOr        string `json:"and_or,omitempty" xml:"and_or,omitempty"`
	SearchType   string `json:"search_type,omitempty" xml:"search_type,omitempty"`
	Value        string `json:"value,omitempty" xml:"value,omitempty"`
	OpeningParen bool   `json:"opening_paren,omitempty" xml:"opening_paren,omitempty"`
	ClosingParen bool   `json:"closing_paren,omitempty" xml:"closing_paren,omitempty"`
}
