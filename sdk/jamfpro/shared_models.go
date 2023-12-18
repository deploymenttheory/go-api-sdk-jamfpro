package jamfpro

// type SelfServiceIcon struct {
// 	ID       int    `xml:"id,omitempty"`
// 	Filename string `xml:"filename,omitempty"`
// 	URI      string `xml:"uri,omitempty"`
// }

// type SelfServiceCategory struct {
// 	ID        int    `xml:"id,omitempty"`
// 	Name      string `xml:"name,omitempty"`
// 	DisplayIn bool   `xml:"display_in,omitempty"`
// 	FeatureIn bool   `xml:"feature_in,omitempty"`
// }

type SharedResourceSite struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// type Building struct {
// 	ID   int    `xml:"id,omitempty"`
// 	Name string `xml:"name,omitempty"`
// }
