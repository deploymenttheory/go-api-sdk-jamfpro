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
type DisplayField struct {
	Name string `xml:"name"`
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

// SharedSubsetCriteriaJamfProAPI represents the criteria for an Search item for jamfprom api
type SharedSubsetCriteriaJamfProAPI struct {
	Name         string `json:"name"`
	Priority     int    `json:"priority"`
	AndOr        string `json:"andOr"`
	SearchType   string `json:"searchType"`
	Value        string `json:"value"`
	OpeningParen bool   `json:"openingParen,omitempty"`
	ClosingParen bool   `json:"closingParen,omitempty"`
}

type SharedResourceLdapServer struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// ResponseError represents the structure of the error response when the force stop request fails.
type SharedResourcResponseError struct {
	HTTPStatus int                         `json:"httpStatus"`
	Errors     []SharedResourceErrorDetail `json:"errors"`
}

// ErrorDetail represents the details of an error in the response
// Used by jamfpro api MDM and managed software updates
type SharedResourceErrorDetail struct {
	Device      int    `json:"device"`
	Group       int    `json:"group"`
	Reason      string `json:"reason"`
	Code        string `json:"code"`
	Field       string `json:"field"`
	Description string `json:"description"`
	ID          string `json:"id"`
}
