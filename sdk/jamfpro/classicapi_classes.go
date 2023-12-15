// classicapi_classes.go
// Jamf Pro Classic Api - Classes
// api reference: https://developer.jamf.com/jamf-pro/reference/classes
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

// Constants for the classes endpoint
const uriClasses = "/JSSResource/classes"

// ResponseClassesList represents the XML response for a list of classes.
type ResponseClassesList struct {
	Size    int             `xml:"size"`
	Classes []ClassListItem `xml:"class"`
}

// ClassItem represents a single class item in the list.
type ClassListItem struct {
	ID          int    `xml:"id"`
	Name        string `xml:"name"`
	Description string `xml:"description"`
}

// Structs for the Class response by ID
type ResourceClass struct {
	ID                  int                       `xml:"id,omitempty"`
	Source              string                    `xml:"source,omitempty"`
	Name                string                    `xml:"name"` // Required
	Description         string                    `xml:"description,omitempty"`
	Site                ClassSubsetSite           `xml:"site"`
	MobileDeviceGroup   ClassSubsetDeviceGroup    `xml:"mobile_device_group,omitempty"`
	Students            []ClassSubsetStudent      `xml:"students>student,omitempty"`
	Teachers            []ClassSubsetTeacher      `xml:"teachers>teacher,omitempty"`
	TeacherIDs          []ClassSubsetTeacherID    `xml:"teacher_ids>id,omitempty"`
	StudentGroupIDs     []ClassSubsetGroupID      `xml:"student_group_ids>id,omitempty"`
	TeacherGroupIDs     []ClassSubsetGroupID      `xml:"teacher_group_ids>id,omitempty"`
	MobileDevices       []ClassSubsetMobileDevice `xml:"mobile_devices>mobile_device,omitempty"`
	MobileDeviceGroupID []ClassSubsetGroupID      `xml:"mobile_device_group_id>id,omitempty"`
	MeetingTimes        ClassSubsetMeetingTimes   `xml:"meeting_times,omitempty"`
	AppleTVs            []ClassSubsetAppleTV      `xml:"apple_tvs>apple_tv,omitempty"`
}
type ClassSubsetSite struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name"` // Required
}

type ClassSubsetDeviceGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type ClassSubsetStudent struct {
	Student string `xml:"student,omitempty"`
}

type ClassSubsetTeacher struct {
	Teacher string `xml:"teacher,omitempty"`
}

type ClassSubsetTeacherID struct {
	ID int `xml:"id,omitempty"`
}

type ClassSubsetGroupID struct {
	ID int `xml:"id,omitempty"`
}

type ClassSubsetMobileDevice struct {
	Name           string `xml:"name,omitempty"`
	UDID           string `xml:"udid,omitempty"`
	WifiMacAddress string `xml:"wifi_mac_address,omitempty"`
}

type ClassSubsetMeetingTimes struct {
	MeetingTime ClassSubsetMeetingTime `xml:"meeting_time,omitempty"`
}

type ClassSubsetMeetingTime struct {
	Days      string `xml:"days,omitempty"`
	StartTime int    `xml:"start_time,omitempty"`
	EndTime   int    `xml:"end_time,omitempty"`
}

type ClassSubsetAppleTV struct {
	Name            string `xml:"name,omitempty"`
	UDID            string `xml:"udid,omitempty"`
	WifiMacAddress  string `xml:"wifi_mac_address,omitempty"`
	DeviceID        string `xml:"device_id,omitempty"`
	AirplayPassword string `xml:"airplay_password,omitempty"`
}

// GetClasses gets a list of all classes.
func (c *Client) GetClasses() (*ResponseClassesList, error) {
	endpoint := uriClasses

	var classes ResponseClassesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &classes)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all Classes: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &classes, nil
}

// GetClassesByID retrieves a class by its ID.
func (c *Client) GetClassesByID(id int) (*ResourceClass, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriClasses, id)

	var class ResourceClass
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &class)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Class by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &class, nil
}

// GetClassesByName retrieves a class by its name.
func (c *Client) GetClassesByName(name string) (*ResourceClass, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriClasses, name)

	var class ResourceClass
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &class)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Class by Name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &class, nil
}

// CreateClassesByID creates a new class with the given details.
func (c *Client) CreateClassesByID(class *ResourceClass) (*ResourceClass, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriClasses) // Using ID 0 for creation as per API pattern

	// If the site is not provided, set default values
	if class.Site.ID == 0 && class.Site.Name == "" {
		class.Site = ClassSubsetSite{
			ID:   -1,
			Name: "None",
		}
	}

	// Wrap the class request with the desired XML structure using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"class"`
		*ResourceClass
	}{
		ResourceClass: class,
	}

	var createdClass ResourceClass
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdClass)
	if err != nil {
		return nil, fmt.Errorf("failed to create Class: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdClass, nil
}

// UpdateClassByID updates an existing class with the given ID.
func (c *Client) UpdateClassesByID(id int, class *ResourceClass) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriClasses, id)

	// Wrap the class request with the desired XML structure using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"class"`
		*ResourceClass
	}{
		ResourceClass: class,
	}

	_, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, nil)
	if err != nil {
		return fmt.Errorf("failed to update Class by ID: %v", err)
	}

	return nil
}

// UpdateClassByName updates an existing class with the given name.
func (c *Client) UpdateClassesByName(name string, class *ResourceClass) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriClasses, name)

	// Wrap the class request with the desired XML structure using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"class"`
		*ResourceClass
	}{
		ResourceClass: class,
	}

	_, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, nil)
	if err != nil {
		return fmt.Errorf("failed to update Class by Name: %v", err)
	}

	return nil
}

// DeleteClassByID deletes an existing class with the given ID.
func (c *Client) DeleteClassByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriClasses, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Class by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteClassByName deletes a class by its name.
func (c *Client) DeleteClassByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriClasses, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Class by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
