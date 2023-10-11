// classicapi_ibeacons.go
// Jamf Pro Classic Api - iBeacons
// api reference: https://developer.jamf.com/jamf-pro/reference/ibeacons
// Classic API requires the structs to support an XML data structure.

package jamfpro

const uriIbeacons = "/JSSResource/ibeacons"

type ResponseIbeacon struct {
	ID    int    `json:"id" xml:"id"`
	Name  string `json:"name" xml:"name"`
	UUID  string `json:"uuid" xml:"uuid"`
	Major int    `json:"major" xml:"major"`
	Minor int    `json:"minor" xml:"minor"`
}

type ResponseIbeaconList struct {
	Ibeacons []IbeaconListItem `json:"ibeacon" xml:"ibeacon"`
}

type IbeaconListItem struct {
	ID    int    `json:"id" xml:"id"`
	Name  string `json:"name" xml:"name"`
	UUID  string `json:"uuid" xml:"uuid"`
	Major int    `json:"major" xml:"major"`
	Minor int    `json:"minor" xml:"minor"`
}

type IBeaconScope struct {
	Id   int    `xml:"id"`
	Name string `xml:"name"`
}

// Functions TODO
