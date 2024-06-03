// jamfproapi_gsx_connection.go
// Jamf Pro Api - GSX connection
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-gsx-connection
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

const uriGSXConnection = "/api/v1/gsx-connection"

// Responses

type ResponseGSXConnectionHistoryList struct {
	TotalCount *int                           `json:"totalCount,omitempty"`
	Results    []ResponseGSXConnectionHistory `json:"results,omitempty"`
}

type ResponseGSXConnectionHistory struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

// Resource
type ResourceGSXConnection struct {
	Enabled          bool        `json:"enabled"`
	Username         string      `json:"username"`
	ServiceAccountNo string      `json:"serviceAccountNo"`
	ShipToNo         string      `json:"shipToNo"`
	GsxKeystore      GsxKeystore `json:"gsxKeystore"`
}

type GsxKeystore struct {
	Name            string `json:"name"`
	ExpirationEpoch int64  `json:"expirationEpoch"`
	ErrorMessage    string `json:"errorMessage"`
}

// CRUD

/*
Function: GetGSXConnectionInformation
Method: GET
Path: /api/v1/gsx-connection
Description: Gets the GSX connection settings.
Parameters: None
Returns: ResourceGSXConnection - The GSX connection settings.
Errors: Returns an error if the request fails.
Example:

	connection, err := client.GetGSXConnectionInformation()
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(connection)
*/
func (c *Client) GetGSXConnectionInformation() (*ResourceGSXConnection, error) {
	endpoint := uriGSXConnection

	var gsxConnectionSettings ResourceGSXConnection
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &gsxConnectionSettings)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "gsx connection information", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &gsxConnectionSettings, nil
}

/*
Function: UpdateGSXConnectionInformation
Method: PATCH
Path: /api/v1/gsx-connection
Description: Updates the GSX connection settings.
Parameters:
  - gsxConnectionSettingsUpdate (*ResourceGSXConnection): The updated GSX connection settings.

Returns: ResourceGSXConnection - The updated GSX connection settings.
Errors: Returns an error if the request fails or if the resource cannot be updated.
Example:

	updatedSettings := &jamfpro.ResourceGSXConnection{
	    Enabled: false,
	    Username: "",
	    ServiceAccountNo: "0",
	    ShipToNo: "0",
	    GsxKeystore: jamfpro.GsxKeystore{
	        Name: "certificate.p12",
	        ExpirationEpoch: 169195490000,
	        ErrorMessage: "Certificate error",
	    },
	}
	updated, err := client.UpdateGSXConnectionInformation(updatedSettings)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(updated)
*/
func (c *Client) UpdateGSXConnectionInformation(gsxConnectionSettingsUpdate *ResourceGSXConnection) (*ResourceGSXConnection, error) {
	endpoint := uriGSXConnection

	var out ResourceGSXConnection
	resp, err := c.HTTP.DoRequest("PATCH", endpoint, gsxConnectionSettingsUpdate, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "gsx connection information", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

/*
Function: GetGSXConnectionHistory
Method: GET
Path: /api/v1/gsx-connection/history
Description: Retrieves all GSX connection history.
Parameters:
  - sort_filter (string): A string specifying the sorting criteria.

Returns: ResponseGSXConnectionHistoryList - A list of GSX connection history.
Errors: Returns an error if the request fails.
Example:

	history, err := client.GetGSXConnectionHistory("sort=id:desc")
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(history)
*/
func (c *Client) GetGSXConnectionHistory(sort_filter string) (*ResponseGSXConnectionHistoryList, error) {
	resp, err := c.DoPaginatedGet(
		uriGSXConnection,
		maxPageSize,
		startingPageNumber,
		sort_filter,
	)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "gsx connection history", err)
	}

	var out ResponseGSXConnectionHistoryList
	out.TotalCount = &resp.Size

	for _, value := range resp.Results {
		var newObj ResponseGSXConnectionHistory
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "gsx connection history", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}
