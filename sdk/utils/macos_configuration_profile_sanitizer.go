package utils

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type ConfigurationProfileRoot struct {
	XMLName xml.Name `xml:"plist"`
	Dict    Dict     `xml:"dict"`
}

type Dict struct {
	PayloadContent     []PayloadEntry `xml:"array>dict"`
	PayloadDisplayName string         `xml:"key>PayloadDisplayName"`
	PayloadIdentifier  string         `xml:"key>PayloadIdentifier"`
	PayloadType        string         `xml:"key>PayloadType"`
	PayloadUUID        string         `xml:"key>PayloadUUID"`
	PayloadVersion     int            `xml:"key>PayloadVersion"`
}

type PayloadEntry struct {
	IfLostReturnToMessage string `xml:"key>IfLostReturnToMessage"`
	LockScreenFootnote    string `xml:"key>LockScreenFootnote"`
	PayloadDisplayName    string `xml:"key>PayloadDisplayName"`
	PayloadIdentifier     string `xml:"key>PayloadIdentifier"`
	PayloadType           string `xml:"key>PayloadType"`
	PayloadUUID           string `xml:"key>PayloadUUID"`
	PayloadVersion        int    `xml:"key>PayloadVersion"`
}

func SanitiseMacOsConfigurationProfile(filepath string) (*ConfigurationProfileRoot, error) {
	xmlFile, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer xmlFile.Close()

	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read xml file: %v", err)
	}

	// fmt.Println(string(xmlData))

	var plist ConfigurationProfileRoot
	err = xml.Unmarshal(xmlData, &plist)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal xml: %v", err)
	}

	// Sanitise the plist here

	return &plist, nil
}
