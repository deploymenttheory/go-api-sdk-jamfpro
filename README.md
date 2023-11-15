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

Here, the DebugMode is set to true, which means the client will print debug information. The Logger uses the SDK's default logger.

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

Note: Remember to always keep your OAuth credentials confidential and never expose them in your code or public repositories. Using configuration files like clientauth.json and .gitignore-ing them is a good practice to ensure they're not accidentally committed.

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

- [ ] ✅ GET `/userid/{id}` - GetAccountByID retrieves the Account by its ID
- [ ] ✅ GET `/username/{username}` - GetAccountByName retrieves the Account by its name
- [ ] ✅ GET `/` - GetAccounts retrieves all user accounts
- [ ] ✅ GET `/groupid/{id}` - GetAccountGroupByID retrieves the Account Group by its ID
- [ ] ✅ GET `/groupname/{id}` - GetAccountGroupByName retrieves the Account Group by its name
- [ ] ✅ POST `/` - CreateAccount creates a new Jamf Pro Account.
- [ ] ✅ POST `/groupid/0` - CreateAccountGroup creates a new Jamf Pro Account Group.
- [ ] ✅ PUT `/userid/{id}` - UpdateAccountByID updates an existing Jamf Pro Account by ID
- [ ] ✅ PUT `/username/{id}` - UpdateAccountByName updates an existing Jamf Pro Account by Name
- [ ] ✅ PUT `/groupid/{id}` - UpdateAccountGroupByID updates an existing Jamf Pro Account Group by ID
- [ ] ✅ PUT `/groupname/{id}` - UpdateAccountGroupByName updates an existing Jamf Pro Account Group by Name
- [ ] ✅ DELETE `/userid/{id}` - DeleteAccountByID deletes an existing Jamf Pro Account by ID
- [ ] ✅ DELETE `/username/{username}` - DeleteAccountByName deletes an existing Jamf Pro Account by Name
- [ ] ✅ DELETE `/groupid/{id}` - DeleteAccountGroupByID deletes an existing Jamf Pro Account Group by ID
- [ ] ✅ DELETE `/groupname/{username}` - DeleteAccountGroupByName deletes an existing Jamf Pro Account Group by Name

### Activation Code - /JSSResource/activationcode

- [ ] ✅ GET `/JSSResource/activationcode` - GetActivationCode retrieves the current activation code and organization name.
- [ ] ✅ PUT `/JSSResource/activationcode` - UpdateActivationCode updates the activation code with a new organization name and code.

### Jamf Pro API Integrations - /api/v1/api-integrations

- [ ] ✅ GET `/api/v1/api-integrations` - GetApiIntegrations fetches all API integrations.
- [ ] ✅ GET `/api/v1/api-integrations/{id}` - GetApiIntegrationByID fetches an API integration by its ID.
- [ ] ✅ GET `/api/v1/api-integrations` followed by searching by name - GetApiIntegrationNameByID fetches an API integration by its display name and then retrieves its details using its ID.
- [ ] ✅ POST `/api/v1/api-integrations` - CreateApiIntegration creates a new API integration.
- [ ] ✅ POST `/api/v1/api-integrations/{id}/client-credentials` - CreateClientCredentialsByApiRoleID creates new client credentials for an API integration by its ID.
- [ ] ✅ PUT `/api/v1/api-integrations/{id}` - UpdateApiIntegrationByID updates an API integration by its ID.
- [ ] ✅ PUT `/api/v1/api-integrations` followed by searching by name - UpdateApiIntegrationByName updates an API integration based on its display name.
- [ ] ✅ POST `/api/v1/api-integrations/{id}/client-credentials` (Used for updating) - UpdateClientCredentialsByApiIntegrationID updates client credentials for an API integration by its ID.
- [ ] ✅ DELETE `/api/v1/api-integrations/{id}` - DeleteApiIntegrationByID deletes an API integration by its ID.
- [ ] ✅ DELETE `/api/v1/api-integrations` followed by searching by name - DeleteApiIntegrationByName deletes an API integration by its display name.

### Jamf Pro API Role Privileges - /api/v1/api-role-privileges

- [ ] ✅ GET `/api/v1/api-role-privileges` - `GetJamfAPIPrivileges` fetches a list of Jamf API role privileges.
- [ ] ✅ GET `/api/v1/api-role-privileges/search?name={name}&limit={limit}` - `GetJamfAPIPrivilegesByName` fetches a Jamf API role privileges by name.


### Jamf Pro API Roles - /api/v1/api-roles

- [ ] ✅ GET `/api/v1/api-roles` - GetJamfAPIRoles fetches all API roles.
- [ ] ✅ GET `/api/v1/api-roles/{id}` - GetJamfApiRolesByID fetches a Jamf API role by its ID.
- [ ] ✅ GET `/api/v1/api-roles` followed by searching by name - GetJamfApiRolesNameById fetches a Jamf API role by its display name and then retrieves its details using its ID.
- [ ] ✅ POST `/api/v1/api-roles` - CreateJamfApiRole creates a new Jamf API role.
- [ ] ✅ PUT `/api/v1/api-roles/{id}` - UpdateJamfApiRoleByID updates a Jamf API role by its ID.
- [ ] ✅ PUT `/api/v1/api-roles` followed by searching by name - UpdateJamfApiRoleByName updates a Jamf API role based on its display name.
- [ ] ✅ DELETE `/api/v1/api-roles/{id}` - DeleteJamfApiRoleByID deletes a Jamf API role by its ID.
- [ ] ✅ DELETE `/api/v1/api-roles` followed by searching by name - DeleteJamfApiRoleByName deletes a Jamf API role by its display name.

### Jamf Pro Classic API - Advanced Computer Searches

- [ ] ✅ GET `/JSSResource/advancedcomputersearches` - GetAdvancedComputerSearches fetches all advanced computer searches.
- [ ] ✅ GET `/JSSResource/advancedcomputersearches/id/{id}` - GetAdvancedComputerSearchByID fetches an advanced computer search by its ID.
- [ ] ✅ GET `/JSSResource/advancedcomputersearches/name/{name}` - GetAdvancedComputerSearchesByName fetches advanced computer searches by their name.
- [ ] ✅ POST `/JSSResource/advancedcomputersearches` - CreateAdvancedComputerSearch creates a new advanced computer search.
- [ ] ✅ PUT `/JSSResource/advancedcomputersearches/id/{id}` - UpdateAdvancedComputerSearchByID updates an existing advanced computer search by its ID.
- [ ] ✅ PUT `/JSSResource/advancedcomputersearches/name/{name}` - UpdateAdvancedComputerSearchByName updates an advanced computer search by its name.
- [ ] ✅ DELETE `/JSSResource/advancedcomputersearches/id/{id}` - DeleteAdvancedComputerSearchByID deletes an advanced computer search by its ID.
- [ ] ✅ DELETE `/JSSResource/advancedcomputersearches/name/{name}` - DeleteAdvancedComputerSearchByName deletes an advanced computer search by its name.

### Jamf Pro Classic API - Advanced Mobile Device Searches

- [ ] ✅ GET `/JSSResource/advancedmobiledevicesearches` - GetAdvancedMobileDeviceSearches fetches all advanced mobile device searches.
- [ ] ✅ GET `/JSSResource/advancedmobiledevicesearches/id/{id}` - GetAdvancedMobileDeviceSearchByID fetches an advanced mobile device search by its ID.
- [ ] ✅ GET `/JSSResource/advancedmobiledevicesearches/name/{name}` - GetAdvancedMobileDeviceSearchByName fetches advanced mobile device searches by their name.
- [ ] ✅ POST `/JSSResource/advancedmobiledevicesearches` - CreateAdvancedMobileDeviceSearch creates a new advanced mobile device search.
- [ ] ✅ PUT `/JSSResource/advancedmobiledevicesearches/id/{id}` - UpdateAdvancedMobileDeviceSearchByID updates an existing advanced mobile device search by its ID.
- [ ] ✅ PUT `/JSSResource/advancedmobiledevicesearches/name/{name}` - UpdateAdvancedMobileDeviceSearchByName updates an advanced mobile device search by its name.
- [ ] ✅ DELETE `/JSSResource/advancedmobiledevicesearches/id/{id}` - DeleteAdvancedMobileDeviceSearchByID deletes an advanced mobile device search by its ID.
- [ ] ✅ DELETE `/JSSResource/advancedmobiledevicesearches/name/{name}` - DeleteAdvancedMobileDeviceSearchByName deletes an advanced mobile device search by its name.


### Jamf Pro Classic API - Advanced User Searches

- [ ] ✅ GET `/JSSResource/advancedusersearches` - GetAdvancedUserSearches fetches all advanced user searches.
- [ ] ✅ GET `/JSSResource/advancedusersearches/id/{id}` - GetAdvancedUserSearchByID fetches an advanced user search by its ID.
- [ ] ✅ GET `/JSSResource/advancedusersearches/name/{name}` - GetAdvancedUserSearchesByName fetches advanced user searches by their name.
- [ ] ✅ POST `/JSSResource/advancedusersearches` - CreateAdvancedUserSearch creates a new advanced user search.
- [ ] ✅ PUT `/JSSResource/advancedusersearches/id/{id}` - UpdateAdvancedUserSearchByID updates an existing advanced user search by its ID.
- [ ] ✅ PUT `/JSSResource/advancedusersearches/name/{name}` - UpdateAdvancedUserSearchByName updates an advanced user search by its name.
- [ ] ✅ DELETE `/JSSResource/advancedusersearches/id/{id}` - DeleteAdvancedUserSearchByID deletes an advanced user search by its ID.
- [ ] ✅ DELETE `/JSSResource/advancedusersearches/name/{name}` - DeleteAdvancedUserSearchByName deletes an advanced user search by its name.

### Allowed File Extensions - /JSSResource/allowedfileextensions

- [ ] ✅ GET `/JSSResource/allowedfileextensions` - GetAllowedFileExtensions retrieves all allowed file extensions
- [ ] ✅ GET `/JSSResource/allowedfileextensions/id/{id}` - GetAllowedFileExtensionByID retrieves the allowed file extension by its ID
- [ ] ✅ GET `/JSSResource/allowedfileextensions/extension/{extensionName}` - GetAllowedFileExtensionByName retrieves the allowed file extension by its name
- [ ] ✅ POST `/JSSResource/allowedfileextensions/id/0` - CreateAllowedFileExtension creates a new allowed file extension
- [ ] ❌ PUT `/JSSResource/allowedfileextensions/id/{id}` - UpdateAllowedFileExtensionByID (API doesn't support update)
- [ ] ✅ DELETE `/JSSResource/allowedfileextensions/id/{id}` - DeleteAllowedFileExtensionByID deletes an existing allowed file extension by ID
- [ ] ✅ DELETE `/JSSResource/allowedfileextensions/extension/{extensionName}` - DeleteAllowedFileExtensionByNameByID deletes an existing allowed file extension by resolving its name to an ID

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

- [ ] ✅ GET `/api/v1/categories` - `GetCategories` retrieves categories based on query parameters.
- [ ] ✅ GET `/api/v1/categories/{id}` - `GetCategoryByID` retrieves a category by its ID.
- [ ] ✅ GET `/api/v1/categories/name/{name}` - `GetCategoryNameByID` retrieves a category by its name and then retrieves its details using its ID.
- [ ] ✅ POST `/api/v1/categories` - `CreateCategory` creates a new category.
- [ ] ✅ PUT `/api/v1/categories/{id}` - `UpdateCategoryByID` updates an existing category by its ID.
- [ ] ✅ PUT `UpdateCategoryByNameByID` updates a category by its name and then updates its details using its ID.
- [ ] ✅ DELETE `/api/v1/categories/{id}` - `DeleteCategoryByID` deletes a category by its ID.
- [ ] ✅ DELETE `DeleteCategoryByNameByID` deletes a category by its name after inferring its ID.
- [ ] ✅ POST `/api/v1/categories/delete-multiple` - `DeleteMultipleCategoriesByID` deletes multiple categories by their IDs.

### Jamf Pro Classic API - Computer Groups

- [ ] ✅ GET `/JSSResource/computergroups` - GetComputerGroups fetches all computer groups.
- [ ] ✅ GET `/JSSResource/computergroups/id/{id}` - GetComputerGroupByID fetches a computer group by its ID.
- [ ] ✅ GET `/JSSResource/computergroups/name/{name}` - GetComputerGroupByName fetches a computer group by its name.
- [ ] ✅ POST `/JSSResource/computergroups/id/0` - CreateComputerGroup creates a new computer group.
- [ ] ✅ PUT `/JSSResource/computergroups/id/{id}` - UpdateComputerGroupByID updates an existing computer group by its ID.
- [ ] ✅ PUT `/JSSResource/computergroups/name/{name}` - UpdateComputerGroupByName updates a computer group by its name.
- [ ] ✅ DELETE `/JSSResource/computergroups/id/{id}` - DeleteComputerGroupByID deletes a computer group by its ID.
- [ ] ✅ DELETE `/JSSResource/computergroups/name/{name}` - DeleteComputerGroupByName deletes a computer group by its name.


### Jamf Pro Classic API - Computer Extension Attributes

- [ ] ✅ GET `/JSSResource/computerextensionattributes` - GetComputerExtensionAttributes gets a list of all computer extension attributes.
- [ ] ✅ GET `/JSSResource/computerextensionattributes/id/{id}` - GetComputerExtensionAttributeByID retrieves a computer extension attribute by its ID.
- [ ] ✅ GET `/JSSResource/computerextensionattributes/name/{name}` - GetComputerExtensionAttributeByName retrieves a computer extension attribute by its name.
- [ ] ✅ POST `/JSSResource/computerextensionattributes/id/0` - CreateComputerExtensionAttribute creates a new computer extension attribute.
- [ ] ✅ PUT `/JSSResource/computerextensionattributes/id/{id}` - UpdateComputerExtensionAttributeByID updates an existing computer extension attribute by its ID.
- [ ] ✅ PUT `/JSSResource/computerextensionattributes/name/{name}` - UpdateComputerExtensionAttributeByName updates a computer extension attribute by its name.
- [ ] ✅ DELETE `/JSSResource/computerextensionattributes/id/{id}` - DeleteComputerExtensionAttributeByID deletes a computer extension attribute by its ID.
- [ ] ⚠️ DELETE (Complex Operation) - `DeleteComputerExtensionAttributeByNameByID` deletes a computer extension attribute by its name (involves fetching ID by name first). 


### Departments - /JSSResource/departments

- [ ] ✅ GET `/JSSResource/departments` - GetDepartments retrieves all departments
- [ ] ✅ GET `/JSSResource/departments/id/{id}` - GetDepartmentByID retrieves the department by its ID
- [ ] ✅ GET `/JSSResource/departments/name/{name}` - GetDepartmentByName retrieves the department by its name
- [ ] ✅ POST `/JSSResource/departments/id/0` - CreateDepartment creates a new department
- [ ] ✅ PUT `/JSSResource/departments/id/{id}` - UpdateDepartmentByID updates an existing department
- [ ] ✅ PUT `/JSSResource/departments/name/{oldName}` - UpdateDepartmentByName updates an existing department by its name
- [ ] ✅ DELETE `/JSSResource/departments/id/{id}` - DeleteDepartmentByID deletes an existing department by its ID
- [ ] ✅ DELETE `/JSSResource/departments/name/{name}` - DeleteDepartmentByName deletes an existing department by its name

### macOS Configuration Profiles - /JSSResource/osxconfigurationprofiles

- [ ] ✅ GET `/JSSResource/osxconfigurationprofiles` - GetMacOSConfigurationProfiles retrieves all macOS configuration profiles.
- [ ] ✅ GET `/JSSResource/osxconfigurationprofiles/id/{id}` - GetMacOSConfigurationProfileByID retrieves the macOS configuration profile by its ID.
- [ ] ✅ GET `/JSSResource/osxconfigurationprofiles/name/{name}` - GetMacOSConfigurationProfileByName retrieves the macOS configuration profile by its name.
- [ ] ✅ POST `/JSSResource/osxconfigurationprofiles/id/0` - CreateMacOSConfigurationProfile creates a new macOS configuration profile.
- [ ] ✅ PUT `/JSSResource/osxconfigurationprofiles/id/{id}` - UpdateMacOSConfigurationProfileByID updates an existing macOS configuration profile by ID.
- [ ] ✅ PUT `/JSSResource/osxconfigurationprofiles/name/{name}` - UpdateMacOSConfigurationProfileByName updates an existing macOS configuration profile by its name.
- [ ] ✅ DELETE `/JSSResource/osxconfigurationprofiles/id/{id}` - DeleteMacOSConfigurationProfileByID deletes an existing macOS configuration profile by ID.
- [ ] ✅ DELETE `/JSSResource/osxconfigurationprofiles/name/{name}` - DeleteMacOSConfigurationProfileByName deletes an existing macOS configuration profile by its name.

### Policies - /JSSResource/policies

- [ ] ✅ GET `/JSSResource/policies` - GetPolicies retrieves a list of all policies
- [ ] ✅ GET `/JSSResource/policies/id/{id}` - GetPolicyByID retrieves the details of a policy by its ID
- [ ] ✅ GET `/JSSResource/policies/name/{name}` - GetPolicyByName retrieves a policy by its name
- [ ] ✅ GET `/JSSResource/policies/category/{category}` - GetPolicyByCategory retrieves policies by their category
- [ ] ✅ GET `/JSSResource/policies/createdBy/{createdBy}` - GetPoliciesByType retrieves policies by the type of entity that created them
- [ ] ✅ POST `/JSSResource/policies/id/0` - CreatePolicy creates a new policy
- [ ] ✅ PUT `/JSSResource/policies/id/{id}` - UpdatePolicyByID updates an existing policy by its ID
- [ ] ✅ PUT `/JSSResource/policies/name/{name}` - UpdatePolicyByName updates an existing policy by its name
- [ ] ✅ DELETE `/JSSResource/policies/id/{id}` - DeletePolicyByID deletes a policy by its ID
- [ ] ✅ DELETE `/JSSResource/policies/name/{name}` - DeletePolicyByName deletes a policy by its name

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

- [ ] ✅ GET `/JSSResource/scripts` - GetScripts retrieves all scripts.
- [ ] ✅ GET `/JSSResource/scripts/id/{id}` - GetScriptsByID retrieves the script details by its ID.
- [ ] ✅ GET `/JSSResource/scripts/name/{name}` - GetScriptsByName retrieves the script details by its name.
- [ ] ✅ POST `/JSSResource/scripts/id/0` - CreateScriptByID creates a new script.
- [ ] ✅ PUT `/JSSResource/scripts/id/{id}` - UpdateScriptByID updates an existing script by its ID.
- [ ] ✅ PUT `/JSSResource/scripts/name/{name}` - UpdateScriptByName updates an existing script by its name.
- [ ] ✅ DELETE `/JSSResource/scripts/id/{id}` - DeleteScriptByID deletes an existing script by its ID.
- [ ] ✅ DELETE `/JSSResource/scripts/name/{name}` - DeleteScriptByName deletes an existing script by its name.

### Jamf Pro Classic API - Sites

- [ ] ✅ GET `/JSSResource/sites` - GetSites fetches all sites.
- [ ] ✅ GET `/JSSResource/sites/id/{id}` - GetSiteByID fetches a site by its ID.
- [ ] ✅ GET `/JSSResource/sites/name/{name}` - GetSiteByName fetches a site by its name.
- [ ] ✅ POST `/JSSResource/sites/id/0` - CreateSite creates a new site.
- [ ] ✅ PUT `/JSSResource/sites/id/{id}` - UpdateSiteByID updates an existing site by its ID.
- [ ] ✅ PUT `/JSSResource/sites/name/{name}` - UpdateSiteByName updates a site by its name.
- [ ] ✅ DELETE `/JSSResource/sites/id/{id}` - DeleteSiteByID deletes a site by its ID.
- [ ] ✅ DELETE `/JSSResource/sites/name/{name}` - DeleteSiteByName deletes a site by its name.

### SSO Failover - /api/v1/sso/failover/generate

- [ ] ✅ GET `/api/v1/sso/failover` - GetSSOFailoverSettings retrieves the current failover settings
- [ ] ✅ PUT `/api/v1/sso/failover/generate` - UpdateFailoverUrl updates failover url, by changing failover key to new one, and returns new failover settings

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

Jamf Pro Classic API - Directory Bindings
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

Jamf Pro Classic API - Computers
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

## Progress Summary

- Total Endpoints: 201
- Covered: 198
- Not Covered: 3
- Partially Covered: 0


## Notes

- No preview api endpoints will be covered by this sdk. Only generally available endpoints will be covered.

