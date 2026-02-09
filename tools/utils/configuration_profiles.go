package utils

import (
	"fmt"
	"io"
	"os"

	"github.com/mitchellh/mapstructure"
	"howett.net/plist"
)

// Struct to mirror MacOS .plist configuration profile data with bucket for unexpected values
type ConfigurationProfile struct {
	PayloadContent     []PayloadContentListItem
	PayloadDisplayName string
	PayloadIdentifier  string
	PayloadType        string
	PayloadUuid        string
	PayloadVersion     int
	UnexpectedValues   map[string]any `mapstructure:",remain"`
}

// Struct to mirror xml payload item with key for all dynamic values
type PayloadContentListItem struct {
	PayloadDisplayName    string
	PayloadIdentifier     string
	PayloadType           string
	PayloadUuid           string
	PayloadVersion        int
	PayloadSpecificValues map[string]any `mapstructure:",remain"`
}

// ConfigurationFilePlistToStructFromFile takes filepath of MacOS Configuration Profile .plist file and returns &ConfigurationProfile
func ConfigurationFilePlistToStructFromFile(filepath string) (*ConfigurationProfile, error) {
	plistFile, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer plistFile.Close()

	xmlData, err := io.ReadAll(plistFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read plist/xml file: %v", err)
	}

	return plistDataToStruct(xmlData)
}

// ConfigurationProfilePlistToStructFromString takes xml of MacOS Configuration Profile .plist file and returns &ConfigurationProfile
func ConfigurationProfilePlistToStructFromString(plistData string) (*ConfigurationProfile, error) {
	return plistDataToStruct([]byte(plistData))
}

// plistDataToStruct takes xml .plist bytes data and returns ConfigurationProfile
func plistDataToStruct(plistBytes []byte) (*ConfigurationProfile, error) {
	var unmarshalledPlist map[string]any
	_, err := plist.Unmarshal(plistBytes, &unmarshalledPlist)
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

// FilterPayloadSpecificFields extracts and returns only the payload-specific fields from the profile
func FilterPayloadSpecificFields(profile *ConfigurationProfile) []map[string]any {
	var filteredPayloads []map[string]any
	for _, payload := range profile.PayloadContent {
		filteredPayload := map[string]any{}
		for key, value := range payload.PayloadSpecificValues {
			// Add only the relevant payload-specific fields, ignoring MDM-specific fields
			if key != "PayloadUUID" && key != "PayloadIdentifier" && key != "PayloadType" && key != "PayloadVersion" {
				filteredPayload[key] = value
			}
		}
		filteredPayloads = append(filteredPayloads, filteredPayload)
	}
	return filteredPayloads
}

// ComparePayloads compares two sets of payload-specific fields and returns true if they are equal
func ComparePayloads(payloads1, payloads2 []map[string]any) bool {
	if len(payloads1) != len(payloads2) {
		return false
	}

	for i := range payloads1 {
		if !compareMaps(payloads1[i], payloads2[i]) {
			return false
		}
	}

	return true
}

// compareMaps compares two maps and returns true if they are equal
func compareMaps(map1, map2 map[string]any) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, val1 := range map1 {
		if val2, ok := map2[key]; !ok || val1 != val2 {
			return false
		}
	}

	return true
}
