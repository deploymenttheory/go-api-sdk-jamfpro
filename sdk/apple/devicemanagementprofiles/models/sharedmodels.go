package models

// XEAPClientConfiguration represents the EAPClientConfiguration dictionary within 802.1X Global Ethernet configuration
type XEAPClientConfiguration struct {
	AcceptEAPTypes []int  `xml:"AcceptEAPTypes>integer"`
	UserName       string `xml:"UserName,omitempty"`
	UserPassword   string `xml:"UserPassword,omitempty"`
}
