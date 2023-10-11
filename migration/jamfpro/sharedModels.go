package jamfpro

type SelfServiceIcon struct {
	ID       int    `xml:"id,omitempty"`
	Filename string `xml:"filename,omitempty"`
	URI      string `xml:"uri,omitempty"`
}

type SelfServiceCategory struct {
	ID        int    `xml:"id,omitempty"`
	Name      string `xml:"name,omitempty"`
	DisplayIn bool   `xml:"display_in,omitempty"`
	FeatureIn bool   `xml:"feature_in,omitempty"`
}
