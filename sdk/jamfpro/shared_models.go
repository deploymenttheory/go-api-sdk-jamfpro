package jamfpro

// Shared Standalone Resources

type SharedResourceSite struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

type SharedResourceCategory struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type SharedResourceSelfServiceIcon struct {
	ID   int    `xml:"id"`
	URI  string `xml:"uri"`
	Data string `xml:"data"`
}

type SharedResourceSelfServiceCategories struct {
	Category []SharedResourceSelfServiceCategory
}

type SharedResourceSelfServiceCategory struct {
	ID       int    `xml:"id"`
	Name     string `xml:"name"`
	Priority int    `xml:"priority,omitempty"`
}

// Advanced Search DisplayField

type SharedAdvancedSearchContainerDisplayField struct {
	Size         int                                    `xml:"size"`
	DisplayField SharedAdvancedSearchSubsetDisplayField `xml:"display_field"`
}

type SharedAdvancedSearchSubsetDisplayField struct {
	Name string `xml:"name"`
}

// Criteria

type SharedContainerCriteria struct {
	Size     int
	Criteria SharedSubsetCriteria
}

type SharedSubsetCriteria struct {
	Name         string `xml:"name"`
	Priority     int    `xml:"priority"`
	AndOr        string `xml:"and_or"`
	SearchType   string `xml:"search_type"`
	Value        int    `xml:"value"`
	OpeningParen bool   `xml:"opening_paren,omitempty"`
	ClosingParen bool   `xml:"closing_paren,omitempty"`
}
