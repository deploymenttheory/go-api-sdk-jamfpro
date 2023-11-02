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
config := http_client.Config{
	DebugMode:             true,
	Logger:                http_client.NewDefaultLogger(),
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

Certainly! Here's the section about URL construction in Markdown format:

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


### Activation Code - /JSSResource/activationcode

- [ ] ✅ GET `/JSSResource/activationcode` - GetActivationCode retrieves the current activation code.
- [ ] ✅ PUT `/JSSResource/activationcode` - UpdateActivationCode updates the activation code.

### Jamf Pro Classic API - Advanced Computer Searches

- [ ] ✅ GET `/JSSResource/advancedcomputersearches` - GetAdvancedComputerSearches fetches all advanced computer searches.
- [ ] ✅ GET `/JSSResource/advancedcomputersearches/id/{id}` - GetAdvancedComputerSearchByID fetches an advanced computer search by its ID.
- [ ] ✅ GET `/JSSResource/advancedcomputersearches/name/{name}` - GetAdvancedComputerSearchesByName fetches advanced computer searches by their name.
- [ ] ✅ POST `/JSSResource/advancedcomputersearches` - CreateAdvancedComputerSearch creates a new advanced computer search.
- [ ] ✅ PUT `/JSSResource/advancedcomputersearches/id/{id}` - UpdateAdvancedComputerSearchByID updates an existing advanced computer search by its ID.
- [ ] ✅ PUT `/JSSResource/advancedcomputersearches/name/{name}` - UpdateAdvancedComputerSearchByName updates an advanced computer search by its name.
- [ ] ✅ DELETE `/JSSResource/advancedcomputersearches/id/{id}` - DeleteAdvancedComputerSearchByID deletes an advanced computer search by its ID.
- [ ] ✅ DELETE `/JSSResource/advancedcomputersearches/name/{name}` - DeleteAdvancedComputerSearchByName deletes an advanced computer search by its name.

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

## Progress Summary

- Total Endpoints: 61
- Covered: 60
- Not Covered: 1
- Partially Covered: 0


## Notes

- No preview api endpoints will be covered by this sdk. Only generally available endpoints will be covered.

