package jamfpro

import (
	"fmt"
)

const uriEnrollmentCustomizations = "/api/v2/enrollment-customizations"

type ResponseEnrollmentCustomization struct {
	TotalCount int                       `json:"totalCount"`
	Results    []EnrollmentCustomization `json:"results"`
}

type EnrollmentCustomization struct {
	ID                                      string                                            `json:"id"`
	SiteId                                  string                                            `json:"siteId"`
	DisplayName                             string                                            `json:"displayName"`
	Description                             string                                            `json:"description"`
	EnrollmentCustomizationBrandingSettings EnrollmentCustomizationDataSubsetBrandingSettings `json:"enrollmentCustomizationBrandingSettings"`
}

type EnrollmentCustomizationDataSubsetBrandingSettings struct {
	TextColor       string `json:"textColor"`
	ButtonColor     string `json:"buttonColor"`
	ButtonTextColor string `json:"buttonTextColor"`
	BackgroundColor string `json:"backgroundColor"`
	IconUrl         string `json:"iconUrl"`
}

func (c *Client) GetEnrollmentCustomizationIdByName(name string) (string, error) {
	var id string
	customizations, err := c.GetEnrollmentCustomizations()
	if err != nil {
		return "", err
	}

	for _, v := range customizations.Results {
		if v.DisplayName == name {
			id = v.ID
			break
		}
	}
	return id, err
}

func (c *Client) GetEnrollmentCustomizationByName(name string) (*EnrollmentCustomization, error) {
	allCustomizationsResponse, err := c.GetEnrollmentCustomizations()
	if err != nil {
		return nil, err
	}

	for _, customization := range allCustomizationsResponse.Results {
		if customization.DisplayName == name {
			return &customization, nil
		}
	}

	return nil, fmt.Errorf("enrollment customization with name '%s' not found", name)
}

func (c *Client) GetEnrollmentCustomizations() (*ResponseEnrollmentCustomization, error) {
	uri := fmt.Sprintf("%s?page=0&page-size=100&sort=id%%3Aasc", uriEnrollmentCustomizations)

	var out ResponseEnrollmentCustomization
	err := c.DoRequest("GET", uri, nil, nil, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf("failed to get enrollment customizations: %v", err)
	}

	return &out, nil
}

func (c *Client) GetEnrollmentCustomizationByID(customizationID string) (*EnrollmentCustomization, error) {
	uri := fmt.Sprintf("%s/%s", uriEnrollmentCustomizations, customizationID)

	var out EnrollmentCustomization
	err := c.DoRequest("GET", uri, nil, nil, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf("failed to get enrollment customization by ID: %v", err)
	}

	return &out, nil
}
