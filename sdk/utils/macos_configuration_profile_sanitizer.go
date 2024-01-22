package utils

import (
	"fmt"
	"io"
	"os"

	"github.com/mitchellh/mapstructure"
	"howett.net/plist"
)

// Struct to mirror MacOS .plist conifguration profile data with bucket for unexpected values
type ConfigurationProfile struct {
	PayloadContent     []PayloadContentListItem
	PayloadDisplayName string
	PayloadIdentifier  string
	PayloadType        string
	PayloadUuid        string
	PayloadVersion     int
	UnexpectedValues   map[string]interface{} `mapstructure:",remain"`
}

// Struct to mirror xml payload item with key for all dynamic values
type PayloadContentListItem struct {
	PayloadDisplayName    string
	PayloadIdentifier     string
	PayloadType           string
	PayloadUuid           string
	PayloadVersion        int
	PayloadSpecificValues map[string]interface{} `mapstructure:",remain"`
}

// ConfigurationFilePlistToStruct takes filepath of MacOS Configuration Profile .plist file and returns &ConfigurationProfile
func ConfigurationFilePlistToStruct(filepath string) (*ConfigurationProfile, error) {
	plistFile, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer plistFile.Close()

	xmlData, err := io.ReadAll(plistFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read plist/xml file: %v", err)
	}

	var unmarshalledPlist map[string]interface{}
	_, err = plist.Unmarshal(xmlData, &unmarshalledPlist)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal plist/xml: %v", err)
	}

	var out ConfigurationProfile
	err = mapstructure.Decode(unmarshalledPlist, &out)
	if err != nil {
		return nil, fmt.Errorf("(mapstructure) failed to map unmarshaled configuration profile to struct: %v", err)
	}

	return &out, nil
}
