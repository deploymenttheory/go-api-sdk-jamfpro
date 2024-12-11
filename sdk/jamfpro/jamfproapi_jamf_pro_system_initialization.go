// jamfproapi_jamf_pro_system_initialization.go
// Jamf Pro Api - JAMF system initialization
// api reference: https://developer.jamf.com/jamf-pro/reference/post_v1-system-initialize
// Jamf Pro API requires the structs to support an JSON data structure.

package jamfpro

import (
	"encoding/json"
	"fmt"
)

const uriSystemInitialize = "/api/v1/system/initialize"
const uriInitializeDatabaseConnection = "/api/v1/system/initialize-database-connection"

// Request struct for database password initialization
type ResourceDatabasePassword struct {
	Password string `json:"password"`
}

// Request Struct
type ResourceSystemInitialize struct {
	ActivationCode  string `json:"activationCode"`
	InstitutionName string `json:"institutionName"`
	EulaAccepted    bool   `json:"eulaAccepted"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Email           string `json:"email,omitempty"`
	JssUrl          string `json:"jssUrl"`
}

// InitializeJamfProServer initializes a fresh Jamf Pro Server installation
func (c *Client) InitializeJamfProServer(systemConfig *ResourceSystemInitialize) error {
	endpoint := uriSystemInitialize

	resp, err := c.HTTP.DoRequest("POST", endpoint, systemConfig, nil)
	if err != nil {
		return fmt.Errorf("failed to initialize Jamf Pro system: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	var errorResponse ResponseError
	if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err == nil {
		if errorResponse.HTTPStatus != 0 {
			return fmt.Errorf("system initialization failed with status %d: %v", errorResponse.HTTPStatus, errorResponse.Errors)
		}
	}

	return nil
}

// InitializeDatabaseConnection sets up the database password during startup
func (c *Client) InitializeDatabaseConnection(password string) error {
	endpoint := uriInitializeDatabaseConnection

	// Create request payload
	dbConfig := ResourceDatabasePassword{
		Password: password,
	}

	resp, err := c.HTTP.DoRequest("POST", endpoint, dbConfig, nil)
	if err != nil {
		return fmt.Errorf("failed to initialize database connection: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	var errorResponse ResponseError
	if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err == nil {
		if errorResponse.HTTPStatus != 0 {
			return fmt.Errorf("database initialization failed with status %d: %v", errorResponse.HTTPStatus, errorResponse.Errors)
		}
	}

	return nil
}
