// classes.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"fmt"
)

const uriAPIClasses = "/JSSResource/classes"

type ResponseClass struct {
	ID                  int                             `json:"id" xml:"id"`
	Source              string                          `json:"source" xml:"source"`
	Name                string                          `json:"name" xml:"name"`
	Description         string                          `json:"description" xml:"description"`
	Site                ClassesDataSubsetSite           `json:"site" xml:"site"`
	MobileDeviceGroup   ClassesDataSubsetGroup          `json:"mobile_device_group" xml:"mobile_device_group"`
	Students            []ClassesDataSubsetName         `json:"students>student" xml:"students>student"`
	Teachers            []ClassesDataSubsetName         `json:"teachers>teacher" xml:"teachers>teacher"`
	TeacherIDs          []ClassesDataSubsetID           `json:"teacher_ids>id" xml:"teacher_ids>id"`
	StudentGroupIDs     []ClassesDataSubsetID           `json:"student_group_ids>id" xml:"student_group_ids>id"`
	TeacherGroupIDs     []ClassesDataSubsetID           `json:"teacher_group_ids>id" xml:"teacher_group_ids>id"`
	MobileDevices       []ClassesDataSubsetMobileDevice `json:"mobile_devices>mobile_device" xml:"mobile_devices>mobile_device"`
	MobileDeviceGroupID []ClassesDataSubsetID           `json:"mobile_device_group_id>id" xml:"mobile_device_group_id>id"`
	MeetingTimes        ClassesDataSubsetMeetingTimes   `json:"meeting_times" xml:"meeting_times"`
	AppleTVs            []ClassesDataSubsetAppleTV      `json:"apple_tvs>apple_tv" xml:"apple_tvs>apple_tv"`
}

type ClassesDataSubsetSite struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

type ClassesDataSubsetGroup struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

type ClassesDataSubsetName struct {
	Name string `json:"name" xml:",chardata"`
}

type ClassesDataSubsetID struct {
	ID int `json:"id" xml:"id"`
}

type ClassesDataSubsetMobileDevice struct {
	Name           string `json:"name" xml:"name"`
	UDID           string `json:"udid" xml:"udid"`
	WifiMacAddress string `json:"wifi_mac_address" xml:"wifi_mac_address"`
	DeviceID       string `json:"device_id,omitempty" xml:"device_id,omitempty"`
}

type ClassesDataSubsetMeetingTimes struct {
	MeetingTime ClassesDataSubsetMeetingTime `json:"meeting_time" xml:"meeting_time"`
}

type ClassesDataSubsetMeetingTime struct {
	Days      string `json:"days" xml:"days"`
	StartTime int    `json:"start_time" xml:"start_time"`
	EndTime   int    `json:"end_time" xml:"end_time"`
}

type ClassesDataSubsetAppleTV struct {
	Name            string `json:"name" xml:"name"`
	UDID            string `json:"udid" xml:"udid"`
	WifiMacAddress  string `json:"wifi_mac_address" xml:"wifi_mac_address"`
	DeviceID        string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	AirplayPassword string `json:"airplay_password,omitempty" xml:"airplay_password,omitempty"`
}

// GetClassByID retrieves the Class by its ID
func (c *Client) GetClassByID(id int) (*ResponseClass, error) {
	url := fmt.Sprintf("%s/id/%d", uriAPIClasses, id)

	var class ResponseClass
	if err := c.DoRequestDebug("GET", url, nil, nil, &class); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &class, nil
}

// GetClasses retrieves a list of all Classes
func (c *Client) GetClasses() (*ResponseClass, error) {
	url := uriAPIClasses

	var classes ResponseClass
	if err := c.DoRequest("GET", url, nil, nil, &classes); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &classes, nil
}

// GetClassByName retrieves the Class by its name
func (c *Client) GetClassByName(className string) (*ResponseClass, error) {
	url := fmt.Sprintf("%s/name/%s", uriAPIClasses, className)

	var class ResponseClass
	if err := c.DoRequest("GET", url, nil, nil, &class); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &class, nil
}

// CreateClass creates a new Class
func (c *Client) CreateClass(class *ResponseClass) error {
	url := fmt.Sprintf("%s/id/0", uriAPIClasses) // Set ID to 0 to let Jamf assign an ID

	if err := c.DoRequest("POST", url, class, nil, nil); err != nil {
		return fmt.Errorf("failed to create class: %v", err)
	}

	return nil
}

// UpdateClassByID updates an existing Class by ID
func (c *Client) UpdateClassByID(id int, class *ResponseClass) error {
	url := fmt.Sprintf("%s/id/%d", uriAPIClasses, id)

	if err := c.DoRequest("PUT", url, class, nil, nil); err != nil {
		return fmt.Errorf("failed to update class by ID: %v", err)
	}

	return nil
}

// UpdateClassByName updates an existing Class by Name
func (c *Client) UpdateClassByName(className string, class *ResponseClass) error {
	url := fmt.Sprintf("%s/name/%s", uriAPIClasses, className)

	if err := c.DoRequest("PUT", url, class, nil, nil); err != nil {
		return fmt.Errorf("failed to update class by Name: %v", err)
	}

	return nil
}

// DeleteClassByID deletes a Class by its ID
func (c *Client) DeleteClassByID(id int) error {
	url := fmt.Sprintf("%s/id/%d", uriAPIClasses, id)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete class by ID: %v", err)
	}

	return nil
}

// DeleteClassByName deletes a Class by its name
func (c *Client) DeleteClassByName(className string) error {
	url := fmt.Sprintf("%s/name/%s", uriAPIClasses, className)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete class named '%s': %v", className, err)
	}

	return nil
}
