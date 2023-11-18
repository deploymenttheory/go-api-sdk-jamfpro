## Getting Started with `go-api-sdk-jamfpro`

This guide will help you get started with `go-api-sdk-jamfpro`, a Go SDK for interfacing with Jamf Pro.

### Prerequisites

Ensure you have Go installed and set up on your system. If not, follow the instructions on the [official Go website](https://golang.org/doc/install).

### Installation

Install the `go-api-sdk-jamfpro` package using `go get`:

```bash
go get github.com/deploymenttheory/go-api-sdk-jamfpro
```

## Usage

sample code: [examples](https://github.com/deploymenttheory/go-jamfpro-api/tree/main/examples)

## Configuring the HTTP Client

To effectively use the `go-api-sdk-jamfpro` SDK, you'll need to set up and configure the HTTP client. Here's a step-by-step guide:

### 1. Setting Constants

At the start of your main program, you can optionally define define some constants that will be used to configure the client for http client. If you don't set any then defaults from `shared_api_client.go` will be used.

```go
const (
	maxConcurrentRequestsAllowed = 5 // Maximum allowed concurrent requests.
	defaultTokenLifespan         = 30 * time.Minute
	defaultBufferPeriod          = 5 * time.Minute
)
```

These constants are used to set the maximum number of concurrent requests the client can make, the lifespan of the token, and a buffer period.

2. Loading OAuth Credentials
OAuth credentials are essential for authenticating with the Jamf Pro API. Store these credentials in a JSON file for secure and easy access. The structure of the clientauth.json should be:

```json
{
  "instanceName": "your_jamf_instance_name", 
  "clientID": "your_client_id",
  "clientSecret": "your_client_secret"
}
```

Replace your_jamf_instance_name, with your jamf pro instance name. e.g for mycompany.jamfcloud.com , use "mycompany".
Replace your_client_id, and your_client_secret with your actual credentials.

In your Go program, load these credentials using:

```go
configFilePath := "path_to_your/clientauth.json"
authConfig, err := http_client.LoadClientAuthConfig(configFilePath)
if err != nil {
	log.Fatalf("Failed to load client OAuth configuration: %v", err)
}
```

3. Configuring the HTTP Client
With the OAuth credentials loaded, you can now configure the HTTP client:

```go
// Initialize a new default logger
logger := http_client.NewDefaultLogger()

// Set the desired log level on the logger
logger.SetLevel(http_client.LogLevelInfo) // LogLevel can be None, Warning, Info, or Debug

// Create the configuration for the HTTP client with the logger
config := http_client.Config{
	Logger: logger,
}
```

The Logger uses the SDK's default logger.

4. Initializing the Jamf Pro Client
Once the HTTP client is configured, initialize the Jamf Pro client:

```go
client := jamfpro.NewClient(authConfig.InstanceName, config)
```

Then, set the OAuth credentials for the client's HTTP client:

```go
oAuthCreds := http_client.OAuthCredentials{
	ClientID:     authConfig.ClientID,
	ClientSecret: authConfig.ClientSecret,
}
client.HTTP.SetOAuthCredentials(oAuthCreds)
```

With these steps, the HTTP client will be fully set up and ready to make requests to the Jamf Pro API. You can then proceed to use the client to perform various actions as demonstrated in the sample code provided.

Note: Remember to always keep your OAuth credentials confidential and never expose them in your code or public repositories. 

---

### URL Construction in the Client

The `go-api-sdk-jamfpro` SDK constructs URLs in a structured manner to ensure consistent and correct API endpoint accesses. Here's a breakdown of how it's done:

#### **Instance Name:**
The primary identifier for constructing URLs in the client is the `InstanceName` which represents the Jamf Pro instance. For example, for the URL `mycompany.jamfcloud.com`, the instance name would be `mycompany`.

#### **Base Domain:**
The SDK uses a constant base domain: `jamfcloud.com`. This domain is appended to the `InstanceName` to form the full domain for the API calls. This can be modifed within the http_client package in jamfpro_api_handler.go if you don't use jamf cloud hosting for your jamf instance.

```go
const (
	BaseDomain     = ".jamfcloud.com"
)
```

#### **Endpoint Path:**
Each API function in the SDK corresponds to a specific Jamf Pro API endpoint. The SDK appends this endpoint path to the constructed domain to derive the full URL.

#### **URL Construction Example:**
Given the `InstanceName` as `mycompany` and an endpoint path `/JSSResource/accounts/userid/{id}`, the full URL constructed by the client would be:
```
https://mycompany.jamfcloud.com/JSSResource/accounts/userid/{id}
```

#### **Customizability:**
The SDK is designed to be flexible. While it uses the `jamfcloud.com` domain by default, this can be updated to meet your requiremesnt for your environment.

### **Note:**
Always ensure that the `InstanceName` is correctly set when initializing the client. Avoid including the full domain (e.g., `.jamfcloud.com`) in the `InstanceName` as the SDK will automatically append it.

---
Putting it all together

```go
package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/path/to/your/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelInfo // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
```

## Go SDK for Jamf Pro API Progress Tracker

### API Coverage Progress

Date: Oct-2023 
Maintainer: [ShocOne]

## Overview

This document tracks the progress of API endpoint coverage tests. As endpoints are tested, they will be marked as covered.

## Coverage Legend

- ✅ - Covered
- ❌ - Not Covered
- ⚠️ - Partially Covered

## Endpoints

### Accounts - /JSSResource/accounts

- [x] ✅ GET `/userid/{id}` - GetAccountByID retrieves the Account by its ID
- [x] ✅ GET `/username/{username}` - GetAccountByName retrieves the Account by its name
- [x] ✅ GET `/` - GetAccounts retrieves all user accounts
- [x] ✅ GET `/groupid/{id}` - GetAccountGroupByID retrieves the Account Group by its ID
- [x] ✅ GET `/groupname/{id}` - GetAccountGroupByName retrieves the Account Group by its name
- [x] ✅ POST `/` - CreateAccount creates a new Jamf Pro Account.
- [x] ✅ POST `/groupid/0` - CreateAccountGroup creates a new Jamf Pro Account Group.
- [x] ✅ PUT `/userid/{id}` - UpdateAccountByID updates an existing Jamf Pro Account by ID
- [x] ✅ PUT `/username/{id}` - UpdateAccountByName updates an existing Jamf Pro Account by Name
- [x] ✅ PUT `/groupid/{id}` - UpdateAccountGroupByID updates an existing Jamf Pro Account Group by ID
- [x] ✅ PUT `/groupname/{id}` - UpdateAccountGroupByName updates an existing Jamf Pro Account Group by Name
- [x] ✅ DELETE `/userid/{id}` - DeleteAccountByID deletes an existing Jamf Pro Account by ID
- [x] ✅ DELETE `/username/{username}` - DeleteAccountByName deletes an existing Jamf Pro Account by Name
- [x] ✅ DELETE `/groupid/{id}` - DeleteAccountGroupByID deletes an existing Jamf Pro Account Group by ID
- [x] ✅ DELETE `/groupname/{username}` - DeleteAccountGroupByName deletes an existing Jamf Pro Account Group by Name

### Activation Code - /JSSResource/activationcode

- [x] ✅ GET `/JSSResource/activationcode` - GetActivationCode retrieves the current activation code and organization name.
- [x] ✅ PUT `/JSSResource/activationcode` - UpdateActivationCode updates the activation code with a new organization name and code.

### Jamf Pro API Integrations - /api/v1/api-integrations

- [x] ✅ GET `/api/v1/api-integrations` - GetApiIntegrations fetches all API integrations.
- [x] ✅ GET `/api/v1/api-integrations/{id}` - GetApiIntegrationByID fetches an API integration by its ID.
- [x] ✅ GET `/api/v1/api-integrations` followed by searching by name - GetApiIntegrationNameByID fetches an API integration by its display name and then retrieves its details using its ID.
- [x] ✅ POST `/api/v1/api-integrations` - CreateApiIntegration creates a new API integration.
- [x] ✅ POST `/api/v1/api-integrations/{id}/client-credentials` - CreateClientCredentialsByApiRoleID creates new client credentials for an API integration by its ID.
- [x] ✅ PUT `/api/v1/api-integrations/{id}` - UpdateApiIntegrationByID updates an API integration by its ID.
- [x] ✅ PUT `/api/v1/api-integrations` followed by searching by name - UpdateApiIntegrationByName updates an API integration based on its display name.
- [x] ✅ POST `/api/v1/api-integrations/{id}/client-credentials` (Used for updating) - UpdateClientCredentialsByApiIntegrationID updates client credentials for an API integration by its ID.
- [x] ✅ DELETE `/api/v1/api-integrations/{id}` - DeleteApiIntegrationByID deletes an API integration by its ID.
- [x] ✅ DELETE `/api/v1/api-integrations` followed by searching by name - DeleteApiIntegrationByName deletes an API integration by its display name.

### Jamf Pro API Role Privileges - /api/v1/api-role-privileges

- [x] ✅ GET `/api/v1/api-role-privileges` - `GetJamfAPIPrivileges` fetches a list of Jamf API role privileges.
- [x] ✅ GET `/api/v1/api-role-privileges/search?name={name}&limit={limit}` - `GetJamfAPIPrivilegesByName` fetches a Jamf API role privileges by name.


### Jamf Pro API Roles - /api/v1/api-roles

- [x] ✅ GET `/api/v1/api-roles` - GetJamfAPIRoles fetches all API roles.
- [x] ✅ GET `/api/v1/api-roles/{id}` - GetJamfApiRolesByID fetches a Jamf API role by its ID.
- [x] ✅ GET `/api/v1/api-roles` followed by searching by name - GetJamfApiRolesNameById fetches a Jamf API role by its display name and then retrieves its details using its ID.
- [x] ✅ POST `/api/v1/api-roles` - CreateJamfApiRole creates a new Jamf API role.
- [x] ✅ PUT `/api/v1/api-roles/{id}` - UpdateJamfApiRoleByID updates a Jamf API role by its ID.
- [x] ✅ PUT `/api/v1/api-roles` followed by searching by name - UpdateJamfApiRoleByName updates a Jamf API role based on its display name.
- [x] ✅ DELETE `/api/v1/api-roles/{id}` - DeleteJamfApiRoleByID deletes a Jamf API role by its ID.
- [x] ✅ DELETE `/api/v1/api-roles` followed by searching by name - DeleteJamfApiRoleByName deletes a Jamf API role by its display name.

### Jamf Pro Classic API - Advanced Computer Searches

- [x] ✅ GET `/JSSResource/advancedcomputersearches` - GetAdvancedComputerSearches fetches all advanced computer searches.
- [x] ✅ GET `/JSSResource/advancedcomputersearches/id/{id}` - GetAdvancedComputerSearchByID fetches an advanced computer search by its ID.
- [x] ✅ GET `/JSSResource/advancedcomputersearches/name/{name}` - GetAdvancedComputerSearchesByName fetches advanced computer searches by their name.
- [x] ✅ POST `/JSSResource/advancedcomputersearches` - CreateAdvancedComputerSearch creates a new advanced computer search.
- [x] ✅ PUT `/JSSResource/advancedcomputersearches/id/{id}` - UpdateAdvancedComputerSearchByID updates an existing advanced computer search by its ID.
- [x] ✅ PUT `/JSSResource/advancedcomputersearches/name/{name}` - UpdateAdvancedComputerSearchByName updates an advanced computer search by its name.
- [x] ✅ DELETE `/JSSResource/advancedcomputersearches/id/{id}` - DeleteAdvancedComputerSearchByID deletes an advanced computer search by its ID.
- [x] ✅ DELETE `/JSSResource/advancedcomputersearches/name/{name}` - DeleteAdvancedComputerSearchByName deletes an advanced computer search by its name.

### Jamf Pro Classic API - Advanced Mobile Device Searches

- [x] ✅ GET `/JSSResource/advancedmobiledevicesearches` - GetAdvancedMobileDeviceSearches fetches all advanced mobile device searches.
- [x] ✅ GET `/JSSResource/advancedmobiledevicesearches/id/{id}` - GetAdvancedMobileDeviceSearchByID fetches an advanced mobile device search by its ID.
- [x] ✅ GET `/JSSResource/advancedmobiledevicesearches/name/{name}` - GetAdvancedMobileDeviceSearchByName fetches advanced mobile device searches by their name.
- [x] ✅ POST `/JSSResource/advancedmobiledevicesearches` - CreateAdvancedMobileDeviceSearch creates a new advanced mobile device search.
- [x] ✅ PUT `/JSSResource/advancedmobiledevicesearches/id/{id}` - UpdateAdvancedMobileDeviceSearchByID updates an existing advanced mobile device search by its ID.
- [x] ✅ PUT `/JSSResource/advancedmobiledevicesearches/name/{name}` - UpdateAdvancedMobileDeviceSearchByName updates an advanced mobile device search by its name.
- [x] ✅ DELETE `/JSSResource/advancedmobiledevicesearches/id/{id}` - DeleteAdvancedMobileDeviceSearchByID deletes an advanced mobile device search by its ID.
- [x] ✅ DELETE `/JSSResource/advancedmobiledevicesearches/name/{name}` - DeleteAdvancedMobileDeviceSearchByName deletes an advanced mobile device search by its name.


### Jamf Pro Classic API - Advanced User Searches

- [x] ✅ GET `/JSSResource/advancedusersearches` - GetAdvancedUserSearches fetches all advanced user searches.
- [x] ✅ GET `/JSSResource/advancedusersearches/id/{id}` - GetAdvancedUserSearchByID fetches an advanced user search by its ID.
- [x] ✅ GET `/JSSResource/advancedusersearches/name/{name}` - GetAdvancedUserSearchesByName fetches advanced user searches by their name.
- [x] ✅ POST `/JSSResource/advancedusersearches` - CreateAdvancedUserSearch creates a new advanced user search.
- [x] ✅ PUT `/JSSResource/advancedusersearches/id/{id}` - UpdateAdvancedUserSearchByID updates an existing advanced user search by its ID.
- [x] ✅ PUT `/JSSResource/advancedusersearches/name/{name}` - UpdateAdvancedUserSearchByName updates an advanced user search by its name.
- [x] ✅ DELETE `/JSSResource/advancedusersearches/id/{id}` - DeleteAdvancedUserSearchByID deletes an advanced user search by its ID.
- [x] ✅ DELETE `/JSSResource/advancedusersearches/name/{name}` - DeleteAdvancedUserSearchByName deletes an advanced user search by its name.

### Allowed File Extensions - /JSSResource/allowedfileextensions

- [x] ✅ GET `/JSSResource/allowedfileextensions` - GetAllowedFileExtensions retrieves all allowed file extensions
- [x] ✅ GET `/JSSResource/allowedfileextensions/id/{id}` - GetAllowedFileExtensionByID retrieves the allowed file extension by its ID
- [x] ✅ GET `/JSSResource/allowedfileextensions/extension/{extensionName}` - GetAllowedFileExtensionByName retrieves the allowed file extension by its name
- [x] ✅ POST `/JSSResource/allowedfileextensions/id/0` - CreateAllowedFileExtension creates a new allowed file extension
- [] ❌ PUT `/JSSResource/allowedfileextensions/id/{id}` - UpdateAllowedFileExtensionByID (API doesn't support update)
- [x] ✅ DELETE `/JSSResource/allowedfileextensions/id/{id}` - DeleteAllowedFileExtensionByID deletes an existing allowed file extension by ID
- [x] ✅ DELETE `/JSSResource/allowedfileextensions/extension/{extensionName}` - DeleteAllowedFileExtensionByNameByID deletes an existing allowed file extension by resolving its name to an ID

### BYO Profiles - `/JSSResource/byoprofiles`

- [x] ✅ GET `/JSSResource/byoprofiles` - `GetBYOProfiles` retrieves all BYO profiles.
- [x] ✅ GET `/JSSResource/byoprofiles/id/{id}` - `GetBYOProfileByID` retrieves a BYO profile by its ID.
- [x] ✅ GET `/JSSResource/byoprofiles/name/{name}` - `GetBYOProfileByName` retrieves a BYO profile by its name.
- [x] ✅ POST `/JSSResource/byoprofiles/id/0` - `CreateBYOProfile` creates a new BYO profile.
- [x] ✅ PUT `/JSSResource/byoprofiles/id/{id}` - `UpdateBYOProfileByID` updates an existing BYO profile by its ID.
- [x] ✅ PUT `/JSSResource/byoprofiles/name/{oldName}` - `UpdateBYOProfileByName` updates an existing BYO profile by its name.
- [x] ✅ DELETE `/JSSResource/byoprofiles/id/{id}` - `DeleteBYOProfileByID` deletes an existing BYO profile by its ID.
- [x] ✅ DELETE `/JSSResource/byoprofiles/name/{name}` - `DeleteBYOProfileByName` deletes an existing BYO profile by its name.


### Jamf Pro API - Categories

- [x] ✅ GET `/api/v1/categories` - `GetCategories` retrieves categories based on query parameters.
- [x] ✅ GET `/api/v1/categories/{id}` - `GetCategoryByID` retrieves a category by its ID.
- [x] ✅ GET `/api/v1/categories/name/{name}` - `GetCategoryNameByID` retrieves a category by its name and then retrieves its details using its ID.
- [x] ✅ POST `/api/v1/categories` - `CreateCategory` creates a new category.
- [x] ✅ PUT `/api/v1/categories/{id}` - `UpdateCategoryByID` updates an existing category by its ID.
- [x] ✅ PUT `UpdateCategoryByNameByID` updates a category by its name and then updates its details using its ID.
- [x] ✅ DELETE `/api/v1/categories/{id}` - `DeleteCategoryByID` deletes a category by its ID.
- [x] ✅ DELETE `DeleteCategoryByNameByID` deletes a category by its name after inferring its ID.
- [x] ✅ POST `/api/v1/categories/delete-multiple` - `DeleteMultipleCategoriesByID` deletes multiple categories by their IDs.

### Jamf Pro Classic API - Computer Groups

- [x] ✅ GET `/JSSResource/computergroups` - GetComputerGroups fetches all computer groups.
- [x] ✅ GET `/JSSResource/computergroups/id/{id}` - GetComputerGroupByID fetches a computer group by its ID.
- [x] ✅ GET `/JSSResource/computergroups/name/{name}` - GetComputerGroupByName fetches a computer group by its name.
- [x] ✅ POST `/JSSResource/computergroups/id/0` - CreateComputerGroup creates a new computer group.
- [x] ✅ PUT `/JSSResource/computergroups/id/{id}` - UpdateComputerGroupByID updates an existing computer group by its ID.
- [x] ✅ PUT `/JSSResource/computergroups/name/{name}` - UpdateComputerGroupByName updates a computer group by its name.
- [x] ✅ DELETE `/JSSResource/computergroups/id/{id}` - DeleteComputerGroupByID deletes a computer group by its ID.
- [x] ✅ DELETE `/JSSResource/computergroups/name/{name}` - DeleteComputerGroupByName deletes a computer group by its name.


### Jamf Pro Classic API - Computer Extension Attributes

- [x] ✅ GET `/JSSResource/computerextensionattributes` - GetComputerExtensionAttributes gets a list of all computer extension attributes.
- [x] ✅ GET `/JSSResource/computerextensionattributes/id/{id}` - GetComputerExtensionAttributeByID retrieves a computer extension attribute by its ID.
- [x] ✅ GET `/JSSResource/computerextensionattributes/name/{name}` - GetComputerExtensionAttributeByName retrieves a computer extension attribute by its name.
- [x] ✅ POST `/JSSResource/computerextensionattributes/id/0` - CreateComputerExtensionAttribute creates a new computer extension attribute.
- [x] ✅ PUT `/JSSResource/computerextensionattributes/id/{id}` - UpdateComputerExtensionAttributeByID updates an existing computer extension attribute by its ID.
- [x] ✅ PUT `/JSSResource/computerextensionattributes/name/{name}` - UpdateComputerExtensionAttributeByName updates a computer extension attribute by its name.
- [x] ✅ DELETE `/JSSResource/computerextensionattributes/id/{id}` - DeleteComputerExtensionAttributeByID deletes a computer extension attribute by its ID.
- [x] ⚠️ DELETE (Complex Operation) - `DeleteComputerExtensionAttributeByNameByID` deletes a computer extension attribute by its name (involves fetching ID by name first). 


### Departments - /JSSResource/departments

- [x] ✅ GET `/JSSResource/departments` - GetDepartments retrieves all departments
- [x] ✅ GET `/JSSResource/departments/id/{id}` - GetDepartmentByID retrieves the department by its ID
- [x] ✅ GET `/JSSResource/departments/name/{name}` - GetDepartmentByName retrieves the department by its name
- [x] ✅ POST `/JSSResource/departments/id/0` - CreateDepartment creates a new department
- [x] ✅ PUT `/JSSResource/departments/id/{id}` - UpdateDepartmentByID updates an existing department
- [x] ✅ PUT `/JSSResource/departments/name/{oldName}` - UpdateDepartmentByName updates an existing department by its name
- [x] ✅ DELETE `/JSSResource/departments/id/{id}` - DeleteDepartmentByID deletes an existing department by its ID
- [x] ✅ DELETE `/JSSResource/departments/name/{name}` - DeleteDepartmentByName deletes an existing department by its name

### macOS Configuration Profiles - /JSSResource/osxconfigurationprofiles

- [x] ✅ GET `/JSSResource/osxconfigurationprofiles` - GetMacOSConfigurationProfiles retrieves all macOS configuration profiles.
- [x] ✅ GET `/JSSResource/osxconfigurationprofiles/id/{id}` - GetMacOSConfigurationProfileByID retrieves the macOS configuration profile by its ID.
- [x] ✅ GET `/JSSResource/osxconfigurationprofiles/name/{name}` - GetMacOSConfigurationProfileByName retrieves the macOS configuration profile by its name.
- [x] ✅ POST `/JSSResource/osxconfigurationprofiles/id/0` - CreateMacOSConfigurationProfile creates a new macOS configuration profile.
- [x] ✅ PUT `/JSSResource/osxconfigurationprofiles/id/{id}` - UpdateMacOSConfigurationProfileByID updates an existing macOS configuration profile by ID.
- [x] ✅ PUT `/JSSResource/osxconfigurationprofiles/name/{name}` - UpdateMacOSConfigurationProfileByName updates an existing macOS configuration profile by its name.
- [x] ✅ DELETE `/JSSResource/osxconfigurationprofiles/id/{id}` - DeleteMacOSConfigurationProfileByID deletes an existing macOS configuration profile by ID.
- [x] ✅ DELETE `/JSSResource/osxconfigurationprofiles/name/{name}` - DeleteMacOSConfigurationProfileByName deletes an existing macOS configuration profile by its name.

### Policies - /JSSResource/policies

- [x] ✅ GET `/JSSResource/policies` - GetPolicies retrieves a list of all policies
- [x] ✅ GET `/JSSResource/policies/id/{id}` - GetPolicyByID retrieves the details of a policy by its ID
- [x] ✅ GET `/JSSResource/policies/name/{name}` - GetPolicyByName retrieves a policy by its name
- [x] ✅ GET `/JSSResource/policies/category/{category}` - GetPolicyByCategory retrieves policies by their category
- [x] ✅ GET `/JSSResource/policies/createdBy/{createdBy}` - GetPoliciesByType retrieves policies by the type of entity that created them
- [x] ✅ POST `/JSSResource/policies/id/0` - CreatePolicy creates a new policy
- [x] ✅ PUT `/JSSResource/policies/id/{id}` - UpdatePolicyByID updates an existing policy by its ID
- [x] ✅ PUT `/JSSResource/policies/name/{name}` - UpdatePolicyByName updates an existing policy by its name
- [x] ✅ DELETE `/JSSResource/policies/id/{id}` - DeletePolicyByID deletes a policy by its ID
- [x] ✅ DELETE `/JSSResource/policies/name/{name}` - DeletePolicyByName deletes a policy by its name

### Jamf Pro API - Self Service Branding macOS

- [x] ✅ GET `/api/v1/self-service/branding/macos` - `GetSelfServiceBrandingMacOS` fetches all self-service branding configurations for macOS.
- [x] ✅ GET `/api/v1/self-service/branding/macos/{id}` - `GetSelfServiceBrandingMacOSByID` fetches a self-service branding configuration for macOS by its ID.
- [x] ✅ GET `/api/v1/self-service/branding/macos/name/{name}` - `GetSelfServiceBrandingMacOSByNameByID` fetches a self-service branding configuration for macOS by its name.
- [x] ✅ POST `/api/v1/self-service/branding/macos` - `CreateSelfServiceBrandingMacOS` creates a new self-service branding configuration for macOS.
- [x] ✅ PUT `/api/v1/self-service/branding/macos/{id}` - `UpdateSelfServiceBrandingMacOSByID` updates an existing self-service branding configuration for macOS by its ID.
- [x] ✅ PUT - `UpdateSelfServiceBrandingMacOSByName` updates a self-service branding configuration for macOS by its name.
- [x] ✅ DELETE `/api/v1/self-service/branding/macos/{id}` - `DeleteSelfServiceBrandingMacOSByID` deletes a self-service branding configuration for macOS by its ID.
- [x] ✅ DELETE - `DeleteSelfServiceBrandingMacOSByName` deletes a self-service branding configuration for macOS by its name.


### Scripts - /JSSResource/scripts

- [x] ✅ GET `/JSSResource/scripts` - GetScripts retrieves all scripts.
- [x] ✅ GET `/JSSResource/scripts/id/{id}` - GetScriptsByID retrieves the script details by its ID.
- [x] ✅ GET `/JSSResource/scripts/name/{name}` - GetScriptsByName retrieves the script details by its name.
- [x] ✅ POST `/JSSResource/scripts/id/0` - CreateScriptByID creates a new script.
- [x] ✅ PUT `/JSSResource/scripts/id/{id}` - UpdateScriptByID updates an existing script by its ID.
- [x] ✅ PUT `/JSSResource/scripts/name/{name}` - UpdateScriptByName updates an existing script by its name.
- [x] ✅ DELETE `/JSSResource/scripts/id/{id}` - DeleteScriptByID deletes an existing script by its ID.
- [x] ✅ DELETE `/JSSResource/scripts/name/{name}` - DeleteScriptByName deletes an existing script by its name.

### Jamf Pro Classic API - Sites

- [x] ✅ GET `/JSSResource/sites` - GetSites fetches all sites.
- [x] ✅ GET `/JSSResource/sites/id/{id}` - GetSiteByID fetches a site by its ID.
- [x] ✅ GET `/JSSResource/sites/name/{name}` - GetSiteByName fetches a site by its name.
- [x] ✅ POST `/JSSResource/sites/id/0` - CreateSite creates a new site.
- [x] ✅ PUT `/JSSResource/sites/id/{id}` - UpdateSiteByID updates an existing site by its ID.
- [x] ✅ PUT `/JSSResource/sites/name/{name}` - UpdateSiteByName updates a site by its name.
- [x] ✅ DELETE `/JSSResource/sites/id/{id}` - DeleteSiteByID deletes a site by its ID.
- [x] ✅ DELETE `/JSSResource/sites/name/{name}` - DeleteSiteByName deletes a site by its name.

### SSO Failover - /api/v1/sso/failover/generate

- [x] ✅ GET `/api/v1/sso/failover` - GetSSOFailoverSettings retrieves the current failover settings
- [x] ✅ PUT `/api/v1/sso/failover/generate` - UpdateFailoverUrl updates failover url, by changing failover key to new one, and returns new failover settings

### Jamf Pro API - Volume Purchasing Subscriptions

This documentation provides details on the API endpoints available for managing Volume Purchasing Subscriptions within Jamf Pro.

#### Endpoints

- [x] ✅ **GET** `/api/v1/volume-purchasing-subscriptions`  
  `GetVolumePurchasingSubscriptions` retrieves all volume purchasing subscriptions.

- [x] ✅ **GET** `/api/v1/volume-purchasing-subscriptions/{id}`  
  `GetVolumePurchasingSubscriptionByID` fetches a single volume purchasing subscription by its ID.

- [x] ✅ **POST** `/api/v1/volume-purchasing-subscriptions`  
  `CreateVolumePurchasingSubscription` creates a new volume purchasing subscription. If `siteId` is not included in the request, it defaults to `siteId: "-1"`.

- [x] ✅ **PUT** `/api/v1/volume-purchasing-subscriptions/{id}`  
  `UpdateVolumePurchasingSubscriptionByID` updates a volume purchasing subscription by its ID.

- [x] ✅ **DELETE** `/api/v1/volume-purchasing-subscriptions/{id}`  
  `DeleteVolumePurchasingSubscriptionByID` deletes a volume purchasing subscription by its ID.

- [x] ✅ **Custom Function**  
  `GetVolumePurchasingSubscriptionByNameByID` fetches a volume purchasing subscription by its display name and retrieves its details using its ID.

- [x] ✅ **Custom Function**  
  `UpdateVolumePurchasingSubscriptionByNameByID` updates a volume purchasing subscription by its display name.

- [x] ✅ **Custom Function**  
  `DeleteVolumePurchasingSubscriptionByName` deletes a volume purchasing subscription by its display name after resolving the name to an ID.

### Jamf Pro API - Computer Inventory Collection Settings

This documentation outlines the API endpoints available for managing Computer Inventory Collection Settings in Jamf Pro.

#### Endpoints

- [x] ✅ **GET** `/api/v1/computer-inventory-collection-settings`  
  `GetComputerInventoryCollectionSettings` retrieves the current computer inventory collection preferences and custom paths.

- [x] ✅ **PATCH** `/api/v1/computer-inventory-collection-settings`  
  `UpdateComputerInventoryCollectionSettings` updates the computer inventory collection preferences.

- [x] ✅ **POST** `/api/v1/computer-inventory-collection-settings/custom-path`  
  `CreateComputerInventoryCollectionSettingsCustomPath` creates a new custom path for the computer inventory collection settings.

- [x] ✅ **DELETE** `/api/v1/computer-inventory-collection-settings/custom-path/{id}`  
  `DeleteComputerInventoryCollectionSettingsCustomPathByID` deletes a custom path by its ID.

### Jamf Pro API - Jamf Pro Information

This documentation covers the API endpoints available for retrieving information about the Jamf Pro server.

#### Endpoints

- [x] ✅ **GET** `/api/v2/jamf-pro-information`  
  `GetJamfProInformation` retrieves information about various services enabled on the Jamf Pro server, like VPP token, DEP account status, BYOD, and more.

	
### Jamf Pro Classic API - Classes

This documentation provides details on the API endpoints available for managing classes within Jamf Pro using the Classic API which requires XML data structure support.

## Endpoints

- [x] ✅ **GET** `/JSSResource/classes`  
  `GetClasses` retrieves a list of all classes.

- [x] ✅ **GET** `/JSSResource/classes/id/{id}`  
  `GetClassesByID` fetches a single class by its ID.

- [x] ✅ **GET** `/JSSResource/classes/name/{name}`  
  `GetClassesByName` retrieves a class by its name.

- [x] ✅ **POST** `/JSSResource/classes/id/0`  
  `CreateClassesByID` creates a new class with the provided details. Using ID `0` indicates creation as per API pattern. If `siteId` is not included, it defaults to `siteId: "-1"`.

- [x] ✅ **PUT** `/JSSResource/classes/id/{id}`  
  `UpdateClassesByID` updates an existing class with the given ID.

- [x] ✅ **PUT** `/JSSResource/classes/name/{name}`  
  `UpdateClassesByName` updates an existing class with the given name.

- [x] ✅ **DELETE** `/JSSResource/classes/id/{id}`  
  `DeleteClassByID` deletes a class by its ID.

- [x] ✅ **DELETE** `/JSSResource/classes/name/{name}`  
  `DeleteClassByName` deletes a class by its name.

###  Jamf Pro Classic API - Computer Invitations
This documentation outlines the API endpoints available for managing computer invitations within Jamf Pro using the Classic API, which relies on XML data structures.

Endpoints
- [x] ✅ GET `/JSSResource/computerinvitations`
GetComputerInvitations retrieves a list of all computer invitations.

- [x] ✅ GET `/JSSResource/computerinvitations/id/{id}`
GetComputerInvitationByID fetches a single computer invitation by its ID.

- [x] ✅ GET `/JSSResource/computerinvitations/invitation/{invitation}`
GetComputerInvitationsByInvitationID retrieves a computer invitation by its invitation ID.

- [x] ✅ POST `/JSSResource/computerinvitations/id/0`
CreateComputerInvitation creates a new computer invitation. Using ID 0 indicates creation as per API pattern. If siteId is not included, it defaults to using a siteId of -1, implying no specific site association.

- [] ❌ PUT `/JSSResource/computerinvitations/invitation/{invitation}`
There is no documented endpoint for updating a computer invitation by its invitation ID.

- [x] ✅ DELETE `/JSSResource/computerinvitations/id/{id}`
DeleteComputerInvitationByID deletes a computer invitation by its ID.

- [] ❌ DELETE `/JSSResource/computerinvitations/invitation/{invitation}`
There is currently no SDK coverage for deleting an invitation by invitation ID

### Jamf Pro Classic API - Disk Encryption Configurations

This documentation provides details on the API endpoints available for managing disk encryption configurations within Jamf Pro using the Classic API which requires XML data structure support.

## Endpoints

- [x] ✅ **GET** `/JSSResource/diskencryptionconfigurations`  
  `GetDiskEncryptionConfigurations` retrieves a serialized list of all disk encryption configurations.

- [x] ✅ **GET** `/JSSResource/diskencryptionconfigurations/id/{id}`  
  `GetDiskEncryptionConfigurationByID` fetches a single disk encryption configuration by its ID.

- [x] ✅ **GET** `/JSSResource/diskencryptionconfigurations/name/{name}`  
  `GetDiskEncryptionConfigurationByName` retrieves a disk encryption configuration by its name.

- [x] ✅ **POST** `/JSSResource/diskencryptionconfigurations/id/0`  
  `CreateDiskEncryptionConfiguration` creates a new disk encryption configuration with the provided details. Using ID `0` indicates creation as per API pattern.

- [x] ✅ **PUT** `/JSSResource/diskencryptionconfigurations/id/{id}`  
  `UpdateDiskEncryptionConfigurationByID` updates an existing disk encryption configuration with the given ID.

- [x] ✅ **PUT** `/JSSResource/diskencryptionconfigurations/name/{name}`  
  `UpdateDiskEncryptionConfigurationByName` updates an existing disk encryption configuration with the given name.

- [x] ✅ **DELETE** `/JSSResource/diskencryptionconfigurations/id/{id}`  
  `DeleteDiskEncryptionConfigurationByID` deletes a disk encryption configuration by its ID.

- [x] ✅ **DELETE** `/JSSResource/diskencryptionconfigurations/name/{name}`  
  `DeleteDiskEncryptionConfigurationByName` deletes a disk encryption configuration by its name.

### Jamf Pro Classic API - Distribution Points

This documentation provides details on the API endpoints available for managing distribution points within Jamf Pro using the Classic API, which requires XML data structure support.

## Endpoints

- [x] ✅ **GET** `/JSSResource/distributionpoints`  
  `GetDistributionPoints` retrieves a serialized list of all distribution points.

- [x] ✅ **GET** `/JSSResource/distributionpoints/id/{id}`  
  `GetDistributionPointByID` fetches a single distribution point by its ID.

- [x] ✅ **GET** `/JSSResource/distributionpoints/name/{name}`  
  `GetDistributionPointByName` retrieves a distribution point by its name.

- [x] ✅ **POST** `/JSSResource/distributionpoints/id/0`  
  `CreateDistributionPoint` creates a new distribution point with the provided details. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/distributionpoints/id/{id}`  
  `UpdateDistributionPointByID` updates an existing distribution point by its ID.

- [x] ✅ **PUT** `/JSSResource/distributionpoints/name/{name}`  
  `UpdateDistributionPointByName` updates an existing distribution point by its name.

- [x] ✅ **DELETE** `/JSSResource/distributionpoints/id/{id}`  
  `DeleteDistributionPointByID` deletes a distribution point by its ID.

- [x] ✅ **DELETE** `/JSSResource/distributionpoints/name/{name}`  
  `DeleteDistributionPointByName` deletes a distribution point by its name.

### Jamf Pro Classic API - Directory Bindings

This documentation provides details on the API endpoints available for managing directory bindings within Jamf Pro using the Classic API, which requires XML data structure support.

## Endpoints

- [x] ✅ **GET** `/JSSResource/directorybindings`
`GetDirectoryBindings` retrieves a serialized list of all directory bindings.

- [x] ✅ **GET** `/JSSResource/directorybindings/id/{id}`
`GetDirectoryBindingByID` fetches a single directory binding by its ID.

- [x] ✅ **GET** `/JSSResource/directorybindings/name/{name}`
`GetDirectoryBindingByName` retrieves a directory binding by its name.

- [x] ✅ **POST** `/JSSResource/directorybindings/id/0`
`CreateDirectoryBinding` creates a new directory binding with the provided details. The ID 0 in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/directorybindings/id/{id}`
`UpdateDirectoryBindingByID` updates an existing directory binding by its ID.

- [x] ✅ **PUT** `/JSSResource/directorybindings/name/{name}`
`UpdateDirectoryBindingByName updates an existing directory binding by its name.

- [x] ✅ **DELETE** `/JSSResource/directorybindings/id/{id}`
`DeleteDirectoryBindingByID deletes a directory binding by its ID.

- [x] ✅ **DELETE** `/JSSResource/directorybindings/name/{name}`
`DeleteDirectoryBindingByName` deletes a directory binding by its name.

### Jamf Pro Classic API - Computers

This documentation provides details on the API endpoints available for managing computers within Jamf Pro using the Classic API, which requires XML data structure support.

## Endpoints

- [x] ✅ **GET** `/JSSResource/computers`
`GetComputers` retrieves a serialized list of all computers.

- [x] ✅ **GET** `/JSSResource/computers/id/{id}`
`GetComputerByID` fetches a single computer by its ID.

- [x] ✅ **GET** `/JSSResource/computers/name/{name}`
`GetComputerByName` retrieves a computer by its name.

- [x] ✅ **POST** `/JSSResource/computers/id/0`
CreateComputer creates a new computer with the provided details. The ID 0 in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/computers/id/{id}`
`UpdateComputerByID` updates an existing computer by its ID.

- [x] ✅ **PUT** `/JSSResource/computers/name/{name}`
`UpdateComputerByName` updates an existing computer by its name.

- [x] ✅ **DELETE** `/JSSResource/computers/id/{id}`
`DeleteComputerByID` deletes a computer by its ID.

- [x] ✅ **DELETE** `/JSSResource/computers/name/{name}`
`DeleteComputerByName` deletes a computer by its name.

### Jamf Pro Classic API - Dock Items

This documentation provides details on the API endpoints available for managing dock items within Jamf Pro using the Classic API, which requires XML data structure support.

## Endpoints

- [x] ✅ **GET** `/JSSResource/dockitems`
  `GetDockItems` retrieves a serialized list of all dock items.

- [x] ✅ **GET** `/JSSResource/dockitems/id/{id}`
  `GetDockItemsByID` fetches a single dock item by its ID.

- [x] ✅ **GET** `/JSSResource/dockitems/name/{name}`
  `GetDockItemsByName` retrieves a dock item by its name.

- [x] ✅ **POST** `/JSSResource/dockitems/id/0`
  `CreateDockItems` creates a new dock item with the provided details. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/dockitems/id/{id}`
  `UpdateDockItemsByID` updates an existing dock item by its ID.

- [x] ✅ **PUT** `/JSSResource/dockitems/name/{name}`
  `UpdateDockItemsByName` updates an existing dock item by its name.

- [x] ✅ **DELETE** `/JSSResource/dockitems/id/{id}`
  `DeleteDockItemsByID` deletes a dock item by its ID.

- [x] ✅ **DELETE** `/JSSResource/dockitems/name/{name}`
  `DeleteDockItemsByName` deletes a dock item by its name.

### Jamf Pro Classic API - eBooks

This documentation provides details on the API endpoints available for managing dock items within Jamf Pro using the Classic API, which requires XML data structure support.

## Endpoints

- [x] ✅ **GET** `/JSSResource/ebooks`
  `GetEbooks` retrieves a serialized list of all ebooks.

- [x] ✅ **GET** `/JSSResource/ebooks/id/{id}`
  `GetEbooksByID` fetches a single ebook by its ID.

- [x] ✅ **GET** `/JSSResource/ebooks/name/{name}`
  `GetEbooksByName` retrieves an ebook by its name.

- [x] ✅ **GET** `/JSSResource/ebooks/name/{name}/subset/{subset}`
  `GetEbooksByNameAndDataSubset` retrieves a specific subset (General, Scope, or SelfService) of an ebook by its name.

- [x] ✅ **POST** `/JSSResource/ebooks/id/0`
  `CreateEbook` creates a new ebook with the provided details. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/ebooks/id/{id}`
  `UpdateEbookByID` updates an existing ebook by its ID.

- [x] ✅ **PUT** `/JSSResource/ebooks/name/{name}`
  `UpdateEbookByName` updates an existing ebook by its name.

- [x] ✅ **DELETE** `/JSSResource/ebooks/id/{id}`
  `DeleteEbookByID` deletes an ebook by its ID.

- [x] ✅ **DELETE** `/JSSResource/ebooks/name/{name}`
  `DeleteEbookByName` deletes an ebook by its name.

### Jamf Pro Classic API - VPP Mac Applications

This documentation outlines the API endpoints available for managing VPP Mac applications within Jamf Pro using the Classic API, which supports XML data structures.

## Endpoints

- [x] ✅ **GET** `/JSSResource/macapplications`
  `GetMacApplications` retrieves a serialized list of all VPP Mac applications.

- [x] ✅ **GET** `/JSSResource/macapplications/id/{id}`
  `GetMacApplicationByID` fetches a single Mac application by its ID.

- [x] ✅ **GET** `/JSSResource/macapplications/name/{name}`
  `GetMacApplicationByName` retrieves a Mac application by its name.

- [x] ✅ **GET** `/JSSResource/macapplications/name/{name}/subset/{subset}`
  `GetMacApplicationByNameAndDataSubset` retrieves a specific subset (General, Scope, SelfService, VPPCodes, and VPP) of a Mac application by its name.

- [x] ✅ **GET** `/JSSResource/macapplications/id/{id}/subset/{subset}`
  `GetMacApplicationByIDAndDataSubset` retrieves a specific subset (General, Scope, SelfService, VPPCodes, and VPP) of a Mac application by its ID.

- [x] ✅ **POST** `/JSSResource/macapplications/id/0`
  `CreateMacApplication` creates a new Mac application with the provided details. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/macapplications/id/{id}`
  `UpdateMacApplicationByID` updates an existing Mac application by its ID.

- [x] ✅ **PUT** `/JSSResource/macapplications/name/{name}`
  `UpdateMacApplicationByName` updates an existing Mac application by its name.

- [x] ✅ **DELETE** `/JSSResource/macapplications/id/{id}`
  `DeleteMacApplicationByID` deletes a Mac application by its ID.

- [x] ✅ **DELETE** `/JSSResource/macapplications/name/{name}`
  `DeleteMacApplicationByName` deletes a Mac application by its name.


### Jamf Pro Classic API - iBeacons

This documentation outlines the API endpoints available for managing iBeacons within Jamf Pro using the Classic API, which supports XML data structures

## Endpoints

- [x] ✅ **GET** `/JSSResource/ibeacons`
  `GetIBeacons` retrieves a serialized list of all iBeacons.

- [x] ✅ **GET** `/JSSResource/ibeacons/id/{id}`
  `GetIBeaconByID` fetches a single iBeacon by its ID.

- [x] ✅ **GET** `/JSSResource/ibeacons/name/{name}`
  `GetIBeaconByName` retrieves an iBeacon by its name.

- [x] ✅ **POST** `/JSSResource/ibeacons/id/0`
  `CreateIBeacon` creates a new iBeacon with the provided details. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/ibeacons/id/{id}`
  `UpdateIBeaconByID` updates an existing iBeacon by its ID.

- [x] ✅ **PUT** `/JSSResource/ibeacons/name/{name}`
  `UpdateIBeaconByName` updates an existing iBeacon by its name.

- [x] ✅ **DELETE** `/JSSResource/ibeacons/id/{id}`
  `DeleteIBeaconByID` deletes an iBeacon by its ID.

- [x] ✅ **DELETE** `/JSSResource/ibeacons/name/{name}`
  `DeleteIBeaconByName` deletes an iBeacon by its name.

### Jamf Pro Classic API - LDAP Servers

This documentation outlines the API endpoints available for managing LDAP servers within Jamf Pro using the Classic API, which supports XML data structures.

## Endpoints

- [x] ✅ **GET** `/JSSResource/ldapservers`
  `GetLDAPServers` retrieves a serialized list of all LDAP servers.

- [x] ✅ **GET** `/JSSResource/ldapservers/id/{id}`
  `GetLDAPServerByID` fetches a single LDAP server by its ID.

- [x] ✅ **GET** `/JSSResource/ldapservers/name/{name}`
  `GetLDAPServerByName` retrieves a LDAP server by its name.

- [x] ✅ **GET** `/JSSResource/ldapservers/id/{id}/user/{user}`
  `GetLDAPServerByIDAndUserDataSubset` retrieves user data for a specific LDAP server by its ID.

- [x] ✅ **GET** `/JSSResource/ldapservers/id/{id}/group/{group}`
  `GetLDAPServerByIDAndGroupDataSubset` retrieves group data for a specific LDAP server by its ID.

- [x] ✅ **GET** `/JSSResource/ldapservers/id/{id}/group/{group}/user/{user}`
  `GetLDAPServerByIDAndUserMembershipInGroupDataSubset` retrieves user group membership details for a specific LDAP server by its ID.

- [x] ✅ **GET** `/JSSResource/ldapservers/name/{name}/user/{user}`
  `GetLDAPServerByNameAndUserDataSubset` retrieves user data for a specific LDAP server by its name.

- [x] ✅ **GET** `/JSSResource/ldapservers/name/{name}/group/{group}`
  `GetLDAPServerByNameAndGroupDataSubset` retrieves group data for a specific LDAP server by its name.

- [x] ✅ **GET** `/JSSResource/ldapservers/name/{name}/group/{group}/user/{user}`
  `GetLDAPServerByNameAndUserMembershipInGroupDataSubset` retrieves user group membership details for a specific LDAP server by its name.

- [x] ✅ **POST** `/JSSResource/ldapservers/id/0`
  `CreateLDAPServer` creates a new LDAP server with the provided details.

- [x] ✅ **PUT** `/JSSResource/ldapservers/id/{id}`
  `UpdateLDAPServerByID` updates an existing LDAP server by its ID.

- [x] ✅ **PUT** `/JSSResource/ldapservers/name/{name}`
  `UpdateLDAPServerByName` updates an existing LDAP server by its name.

- [x] ✅ **DELETE** `/JSSResource/ldapservers/id/{id}`
  `DeleteLDAPServerByID` deletes an LDAP server by its ID.

- [x] ✅ **DELETE** `/JSSResource/ldapservers/name/{name}`
  `DeleteLDAPServerByName` deletes an LDAP server by its name.

### Jamf Pro Classic API - Licensed Software

This documentation outlines the API endpoints available for managing Licensed Software within Jamf Pro using the Classic API, which supports XML data structures.

## Endpoints

- [x] ✅ **GET** `/JSSResource/licensedsoftware`
  `GetLicensedSoftware` retrieves a serialized list of all Licensed Software.

- [x] ✅ **GET** `/JSSResource/licensedsoftware/id/{id}`
  `GetLicensedSoftwareByID` fetches details of a single Licensed Software item by its ID.

- [x] ✅ **GET** `/JSSResource/licensedsoftware/name/{name}`
  `GetLicensedSoftwareByName` retrieves details of a Licensed Software item by its name.

- [x] ✅ **POST** `/JSSResource/licensedsoftware/id/0`
  `CreateLicensedSoftware` creates a new Licensed Software item. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/licensedsoftware/id/{id}`
  `UpdateLicensedSoftwareByID` updates an existing Licensed Software item by its ID.

- [x] ✅ **PUT** `/JSSResource/licensedsoftware/name/{name}`
  `UpdateLicensedSoftwareByName` updates an existing Licensed Software item by its name.

- [x] ✅ **DELETE** `/JSSResource/licensedsoftware/id/{id}`
  `DeleteLicensedSoftwareByID` deletes a Licensed Software item by its ID.

- [x] ✅ **DELETE** `/JSSResource/licensedsoftware/name/{name}`
  `DeleteLicensedSoftwareByName` deletes a Licensed Software item by its name.

### Jamf Pro Classic API - Mobile Device Applications

This documentation outlines the API endpoints available for managing Mobile Device Applications within Jamf Pro using the Classic API, which supports XML data structures.

## Endpoints

- [x] ✅ **GET** `/JSSResource/mobiledeviceapplications`
  `GetMobileDeviceApplications` retrieves a serialized list of all Mobile Device Applications.

- [x] ✅ **GET** `/JSSResource/mobiledeviceapplications/id/{id}`
  `GetMobileDeviceApplicationByID` fetches details of a single Mobile Device Application by its ID.

- [x] ✅ **GET** `/JSSResource/mobiledeviceapplications/name/{name}`
  `GetMobileDeviceApplicationByName` retrieves details of a Mobile Device Application by its name.

- [x] ✅ **GET** `/JSSResource/mobiledeviceapplications/bundleid/{bundleid}`
  `GetMobileDeviceApplicationByAppBundleID` fetches details of a Mobile Device Application by its Bundle ID.

- [x] ✅ **GET** `/JSSResource/mobiledeviceapplications/bundleid/{bundleid}/version/{version}`
  `GetMobileDeviceApplicationByAppBundleIDAndVersion` fetches details of a Mobile Device Application by its Bundle ID and specific version.

- [x] ✅ **GET** `/JSSResource/mobiledeviceapplications/id/{id}/subset/{subset}`
  `GetMobileDeviceApplicationByIDAndDataSubset` fetches a Mobile Device Application by its ID and a specified data subset.

- [x] ✅ **GET** `/JSSResource/mobiledeviceapplications/name/{name}/subset/{subset}`
  `GetMobileDeviceApplicationByNameAndDataSubset` fetches a Mobile Device Application by its name and a specified data subset.

- [x] ✅ **POST** `/JSSResource/mobiledeviceapplications/id/0`
  `CreateMobileDeviceApplication` creates a new Mobile Device Application. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceapplications/id/{id}`
  `UpdateMobileDeviceApplicationByID` updates an existing Mobile Device Application by its ID.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceapplications/name/{name}`
  `UpdateMobileDeviceApplicationByName` updates an existing Mobile Device Application by its name.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceapplications/bundleid/{bundleid}`
  `UpdateMobileDeviceApplicationByApplicationBundleID` updates an existing Mobile Device Application by its Bundle ID.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceapplications/bundleid/{bundleid}/version/{version}`
  `UpdateMobileDeviceApplicationByIDAndAppVersion` updates an existing Mobile Device Application by its ID and specific version.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceapplications/id/{id}`
  `DeleteMobileDeviceApplicationByID` deletes a Mobile Device Application by its ID.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceapplications/name/{name}`
  `DeleteMobileDeviceApplicationByName` deletes a Mobile Device Application by its name.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceapplications/bundleid/{bundleid}`
  `DeleteMobileDeviceApplicationByBundleID` deletes a Mobile Device Application by its Bundle ID.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceapplications/bundleid/{bundleid}/version/{version}`
  `DeleteMobileDeviceApplicationByBundleIDAndVersion` deletes a Mobile Device Application by its Bundle ID and specific version.

### Jamf Pro Classic API - Mobile Device Configuration Profiles

This documentation outlines the API endpoints available for managing Mobile Device Configuration Profiles within Jamf Pro using the Classic API, which supports XML data structures.

## Endpoints

- [x] ✅ **GET** `/JSSResource/mobiledeviceconfigurationprofiles`
  `GetMobileDeviceConfigurationProfiles` retrieves a serialized list of all Mobile Device Configuration Profiles.

- [x] ✅ **GET** `/JSSResource/mobiledeviceconfigurationprofiles/id/{id}`
  `GetMobileDeviceConfigurationProfileByID` fetches details of a single Mobile Device Configuration Profile by its ID.

- [x] ✅ **GET** `/JSSResource/mobiledeviceconfigurationprofiles/name/{name}`
  `GetMobileDeviceConfigurationProfileByName` retrieves details of a Mobile Device Configuration Profile by its name.

- [x] ✅ **GET** `/JSSResource/mobiledeviceconfigurationprofiles/id/{id}/subset/{subset}`
  `GetMobileDeviceConfigurationProfileByIDBySubset` fetches a specific Mobile Device Configuration Profile by its ID and a specified subset.

- [x] ✅ **GET** `/JSSResource/mobiledeviceconfigurationprofiles/name/{name}/subset/{subset}`
  `GetMobileDeviceConfigurationProfileByNameBySubset` fetches a specific Mobile Device Configuration Profile by its name and a specified subset.

- [x] ✅ **POST** `/JSSResource/mobiledeviceconfigurationprofiles/id/0`
  `CreateMobileDeviceConfigurationProfile` creates a new Mobile Device Configuration Profile. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceconfigurationprofiles/id/{id}`
  `UpdateMobileDeviceConfigurationProfileByID` updates an existing Mobile Device Configuration Profile by its ID.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceconfigurationprofiles/name/{name}`
  `UpdateMobileDeviceConfigurationProfileByName` updates an existing Mobile Device Configuration Profile by its name.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceconfigurationprofiles/id/{id}`
  `DeleteMobileDeviceConfigurationProfileByID` deletes a Mobile Device Configuration Profile by its ID.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceconfigurationprofiles/name/{name}`
  `DeleteMobileDeviceConfigurationProfileByName` deletes a Mobile Device Configuration Profile by its name.

### Jamf Pro Classic API - Mobile Extension Attributes

This documentation outlines the API endpoints available for managing Mobile Extension Attributes within Jamf Pro using the Classic API, which supports XML data structures.

## Endpoints

- [x] ✅ **GET** `/JSSResource/mobiledeviceextensionattributes`
  `GetMobileExtensionAttributes` retrieves a serialized list of all Mobile Extension Attributes.

- [x] ✅ **GET** `/JSSResource/mobiledeviceextensionattributes/id/{id}`
  `GetMobileExtensionAttributeByID` fetches details of a single Mobile Extension Attribute by its ID.

- [x] ✅ **GET** `/JSSResource/mobiledeviceextensionattributes/name/{name}`
  `GetMobileExtensionAttributeByName` retrieves details of a Mobile Extension Attribute by its name.

- [x] ✅ **POST** `/JSSResource/mobiledeviceextensionattributes/id/0`
  `CreateMobileExtensionAttribute` creates a new Mobile Extension Attribute. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceextensionattributes/id/{id}`
  `UpdateMobileExtensionAttributeByID` updates an existing Mobile Extension Attribute by its ID.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceextensionattributes/name/{name}`
  `UpdateMobileExtensionAttributeByName` updates an existing Mobile Extension Attribute by its name.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceextensionattributes/id/{id}`
  `DeleteMobileExtensionAttributeByID` deletes a Mobile Extension Attribute by its ID.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceextensionattributes/name/{name}`
  `DeleteMobileExtensionAttributeByName` deletes a Mobile Extension Attribute by its name.

### Jamf Pro Classic API - Mobile Device Enrollment Profiles

This documentation outlines the API endpoints available for managing Mobile Device Enrollment Profiles within Jamf Pro using the Classic API, which supports XML data structures.

## Endpoints

- [x] ✅ **GET** `/JSSResource/mobiledeviceenrollmentprofiles`
  `GetMobileDeviceEnrollmentProfiles` retrieves a serialized list of all Mobile Device Enrollment Profiles.

- [x] ✅ **GET** `/JSSResource/mobiledeviceenrollmentprofiles/id/{id}`
  `GetMobileDeviceEnrollmentProfileByID` fetches details of a single Mobile Device Enrollment Profile by its ID.

- [x] ✅ **GET** `/JSSResource/mobiledeviceenrollmentprofiles/name/{name}`
  `GetMobileDeviceEnrollmentProfileByName` retrieves details of a Mobile Device Enrollment Profile by its name.

- [x] ✅ **GET** `/JSSResource/mobiledeviceenrollmentprofiles/invitation/{invitation}`
  `GetProfileByInvitation` fetches a Mobile Device Enrollment Profile by its invitation.

- [x] ✅ **GET** `/JSSResource/mobiledeviceenrollmentprofiles/id/{id}/subset/{subset}`
  `GetMobileDeviceEnrollmentProfileByIDBySubset` fetches a specific Mobile Device Enrollment Profile by its ID and a specified subset.

- [x] ✅ **GET** `/JSSResource/mobiledeviceenrollmentprofiles/name/{name}/subset/{subset}`
  `GetMobileDeviceEnrollmentProfileByNameBySubset` fetches a specific Mobile Device Enrollment Profile by its name and a specified subset.

- [x] ✅ **POST** `/JSSResource/mobiledeviceenrollmentprofiles/id/0`
  `CreateMobileDeviceEnrollmentProfile` creates a new Mobile Device Enrollment Profile. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceenrollmentprofiles/id/{id}`
  `UpdateMobileDeviceEnrollmentProfileByID` updates an existing Mobile Device Enrollment Profile by its ID.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceenrollmentprofiles/name/{name}`
  `UpdateMobileDeviceEnrollmentProfileByName` updates an existing Mobile Device Enrollment Profile by its name.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceenrollmentprofiles/invitation/{invitation}`
  `UpdateMobileDeviceEnrollmentProfileByInvitation` updates an existing Mobile Device Enrollment Profile by its invitation.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceenrollmentprofiles/id/{id}`
  `DeleteMobileDeviceEnrollmentProfileByID` deletes a Mobile Device Enrollment Profile by its ID.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceenrollmentprofiles/name/{name}`
  `DeleteMobileDeviceEnrollmentProfileByName` deletes a Mobile Device Enrollment Profile by its name.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceenrollmentprofiles/invitation/{invitation}`
  `DeleteMobileDeviceEnrollmentProfileByInvitation` deletes a Mobile Device Enrollment Profile by its invitation.

### Jamf Pro Classic API - Printers

This documentation outlines the API endpoints available for managing printers within Jamf Pro using the Classic API, which supports XML data structures.

## Endpoints

- [x] ✅ **GET** `/JSSResource/printers`
  `GetPrinters` retrieves a serialized list of all printers.

- [x] ✅ **GET** `/JSSResource/printers/id/{id}`
  `GetPrinterByID` fetches details of a single printer by its ID.

- [x] ✅ **GET** `/JSSResource/printers/name/{name}`
  `GetPrinterByName` retrieves details of a printer by its name.

- [x] ✅ **POST** `/JSSResource/printers/id/0`
  `CreatePrinters` creates a new printer. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/printers/id/{id}`
  `UpdatePrinterByID` updates an existing printer by its ID.

- [x] ✅ **PUT** `/JSSResource/printers/name/{name}`
  `UpdatePrinterByName` updates an existing printer by its name.

- [x] ✅ **DELETE** `/JSSResource/printers/id/{id}`
  `DeletePrinterByID` deletes a printer by its ID.

- [x] ✅ **DELETE** `/JSSResource/printers/name/{name}`
  `DeletePrinterByName` deletes a printer by its name.

### Jamf Pro Classic API - Network Segments

This documentation outlines the API endpoints available for managing Network Segments within Jamf Pro using the Classic API, which supports XML data structures.

## Endpoints

- [x] ✅ **GET** `/JSSResource/networksegments`
  `GetNetworkSegments` retrieves a serialized list of all Network Segments.

- [x] ✅ **GET** `/JSSResource/networksegments/id/{id}`
  `GetNetworkSegmentByID` fetches details of a single Network Segment by its ID.

- [x] ✅ **GET** `/JSSResource/networksegments/name/{name}`
  `GetNetworkSegmentByName` retrieves details of a Network Segment by its name.

- [x] ✅ **POST** `/JSSResource/networksegments/id/0`
  `CreateNetworkSegment` creates a new Network Segment. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/networksegments/id/{id}`
  `UpdateNetworkSegmentByID` updates an existing Network Segment by its ID.

- [x] ✅ **PUT** `/JSSResource/networksegments/name/{name}`
  `UpdateNetworkSegmentByName` updates an existing Network Segment by its name.

- [x] ✅ **DELETE** `/JSSResource/networksegments/id/{id}`
  `DeleteNetworkSegmentByID` deletes a Network Segment by its ID.

- [x] ✅ **DELETE** `/JSSResource/networksegments/name/{name}`
  `DeleteNetworkSegmentByName` deletes a Network Segment by its name.


## Progress Summary

- Total Endpoints: 347
- Covered: 328
- Not Covered: 19
- Partially Covered: 0


## Notes

- No preview api endpoints will be covered by this sdk. Only generally available endpoints will be covered.

