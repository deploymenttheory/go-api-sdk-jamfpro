// jamfproapi_categories.go
// Jamf Pro Api - osx configuration profiles
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-categories
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

const uriCategories = "/api/v1/categories"

type ResponseCategories struct {
	TotalCount *int       `json:"totalCount,omitempty"`
	Results    []Category `json:"results,omitempty"`
}

type Category struct {
	Id       *string `json:"id,omitempty"` // The response type to be returned is a string
	Name     *string `json:"name,omitempty"`
	Priority *int    `json:"priority,omitempty"`
	Href     *string `json:"href,omitempty"`
}

type GeneralCategory struct {
	ID   string `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

//TODO
