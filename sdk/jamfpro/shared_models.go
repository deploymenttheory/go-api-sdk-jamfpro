package jamfpro

// Shared Standalone Resources

type SharedResourceSite struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// Advanced Search DisplayField

type SharedAdvancedSearchContainerDisplayField struct {
	Size         int                                    `xml:"size"`
	DisplayField SharedAdvancedSearchSubsetDisplayField `xml:"display_field"`
}

type SharedAdvancedSearchSubsetDisplayField struct {
	Name string `xml:"name"`
}

// Advanced Search Criteria

type SharedAdvancedSearchContainerCriteria struct {
	Size     int
	Criteria SharedAdvancedSearchSubsetCriteria
}

type SharedAdvancedSearchSubsetCriteria struct {
	Name         string `xml:"name"`
	Priority     int    `xml:"priority"`
	AndOr        string `xml:"and_or"`
	SearchType   string `xml:"search_type"`
	Value        int    `xml:"value"`
	OpeningParen bool   `xml:"opening_paren,omitempty"`
	ClosingParen bool   `xml:"closing_paren,omitempty"`
}
