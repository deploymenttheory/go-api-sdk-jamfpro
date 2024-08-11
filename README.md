## Getting Started with `go-api-sdk-jamfpro`

This guide will help you get started with `go-api-sdk-jamfpro`, a Go SDK for interfacing with Jamf Pro.

### Go Prerequisites

Ensure you have Go installed and set up on your system. If not, follow the instructions on the [official Go website](https://golang.org/doc/install).

### Installation

Install the `go-api-sdk-jamfpro` package using `go get`:

```bash
go get github.com/deploymenttheory/go-api-sdk-jamfpro
```

## Usage

It's highly recommended to use the [examples](https://github.com/deploymenttheory/go-api-sdk-jamfpro/tree/main/examples) library to get started with the SDK. Here you will find examples of how to use the SDK to perform various operations on your Jamf Pro instance.


### Authentication Pre-requisites: Obtaining Client ID and Client Secret

To securely interact with the Jamf Pro API, it's essential to obtain a unique client ID and client secret. These credentials are used to authenticate API requests and ensure that only authorized users and applications can access your Jamf Pro environment.

### API Roles and Clients in Jamf Pro

The API Roles and Clients functionality in Jamf Pro provides a dedicated interface for controlling access to both the Jamf Pro API and the Classic API. This feature allows you to create custom privilege sets, known as API roles, and assign them as needed to ensure that API clients possess only the necessary capabilities for their tasks. Roles can be shared between clients or assigned more than one to a client, offering a flexible way to manage and reuse privilege sets for various purposes with granular control.

### Creating API Clients

To create an API client and generate a client ID and secret:

1. Navigate to the **Settings** in your Jamf Pro dashboard.
2. Select **System Settings**.
3. Choose **API Roles and Clients** under the System Settings options.
4. Click on **New Client** to create a new API client.
5. Assign a name and description for your client, and select the API roles that define the permissions this client will have.
6. Once the client is created, Jamf Pro will generate a unique client ID and client secret. Make sure to securely store these credentials as they are required for authenticating your API requests.

For a detailed guide and best practices on creating and managing API clients and roles, refer to the official Jamf Pro documentation: [API Roles and Clients in Jamf Pro](https://learn.jamf.com/bundle/jamf-pro-documentation-current/page/API_Roles_and_Clients.html).

Remember to keep your client ID and secret confidential and secure, as these credentials provide access to your Jamf Pro environment's API.

## Configuring the Jamf Pro Client with the Go SDK

The `go-api-sdk-jamfpro` provides two convenient ways to build and configure your Jamf Pro client: using environment variables or a JSON configuration file. This flexibility allows for easy integration into different environments and deployment pipelines.

### Option 1: Building Client with Environment Variables

For scenarios where you prefer not to use configuration files (e.g., in containerized environments or CI/CD pipelines), you can configure the Jamf Pro client using environment variables.

1. **Set Environment Variables**: Define the necessary environment variables in your environment. This includes credentials (for OAuth or classic auth), instance details, and client options.

    ```shell
    export CLIENT_ID="your_client_id"
    export CLIENT_SECRET="your_client_secret"
    export INSTANCE_DOMAIN="https://your_instance.jamfcloud.com" # use the fqdn
    export AUTH_METHOD="oauth2" # or "basic"
    export BASIC_AUTH_USERNAME="your_basic_auth_username" # Required if using basic auth
    export BASIC_AUTH_PASSWORD="your_basic_auth_password" # Required if using basic auth
    export CLIENT_ID="your_client_id" # Required if using oauth2
    export CLIENT_SECRET="your_client_secret" # Required if using oauth2
    export LOG_LEVEL="info" # or "debug" / "info" / "warn" / "dpanic" / "error"
    export LOG_OUTPUT_FORMAT="pretty" # or "json" 
    export LOG_CONSOLE_SEPARATOR=" " # or any other separator
    export LOG_EXPORT_PATH="/your/log/path/" # optional, ensure permissions to file path
    export EXPORT_LOGS="true" # or "false"
    export HIDE_SENSITIVE_DATA="true" # or "false"
    export MAX_RETRY_ATTEMPTS="3" # optional  
    export MAX_CONCURRENT_REQUESTS="5" # optional
    export ENABLE_DYNAMIC_RATE_LIMITING="true" # or "false"
    export TOKEN_REFRESH_BUFFER_PERIOD_SECONDS="300" # optional, in seconds
    export TOTAL_RETRY_DURATION_SECONDS="300" # optional, in seconds
    export CUSTOM_TIMEOUT_SECONDS="60" # optional, in seconds
    export FOLLOW_REDIRECTS="true" # or "false"
    export MAX_REDIRECTS="5" # Sets the maximum number of redirects
    export ENABLE_CONCURRENCY_MANAGEMENT="true" # or "false"
    export JAMF_LOAD_BALANCER_LOCK="true" # or "false"
    export CUSTOM_COOKIES='[{"name": "jpro-ingress", "value": "your_cookie_value"}, {"name": "sessionToken", "value": "abc123"}, {"name": "userPref", "value": "lightMode"}]' # optional, JSON array of cookies
    ```

2. **Build the Client**: Use the `BuildClientWithEnv` function to build the Jamf Pro client using the environment variables.

    ```go
    client, err := jamfpro.BuildClientWithEnv()
    if err != nil {
        log.Fatalf("Failed to build Jamf Pro client with environment variables: %v", err)
    }
    ```

    This method will automatically read the configuration from the environment variables and initialize the Jamf Pro client.

### Option 2: Building Client with a Configuration File

For those who prefer using configuration files for setting up the client, the SDK supports loading configuration from a JSON file.

1. **Prepare the Configuration File**: Create a JSON file with the necessary configuration. This includes authentication credentials, environment settings, and client options.

    ```json
    {
      "log_level": "info", // or "debug" / "info" / "warn" / "dpanic" / "error"
      "log_output_format": "pretty", // or "json"
      "log_console_separator": "  ",
      "log_export_path": "/your/log/path/", // optional, ensure permissions to file path
      "export_logs": true, // or false
      "hide_sensitive_data": false, // redact sensitive data from logs
      "instance_domain": "https://lbgsandbox.jamfcloud.com",
      "auth_method": "oauth2", // or "basic"
      "client_id": "your_client_id", // Required if using oauth2
      "client_secret": "your_client_secret", // Required if using oauth2
      "basic_auth_username": "your_basic_auth_username", // Required if using basic auth
      "basic_auth_password": "your_basic_auth_password", // Required if using basic auth
      "jamf_load_balancer_lock": false, // or true
      "max_retry_attempts": 3,
      "enable_dynamic_rate_limiting": true,
      "max_concurrent_requests": 5, // optional
      "token_refresh_buffer_period_seconds": 300, // optional in seconds
      "total_retry_duration_seconds": 300, // optional in seconds
      "custom_timeout_seconds": 300, // optional in seconds
      "follow_redirects": true,
      "max_redirects": 5,
      "enable_concurrency_management": true,
      "custom_cookies": [
        {
          "name": "cookie1",
          "value": "value1"
        },
        {
          "name": "cookie2",
          "value": "value2"
        }
      ]
    }
    ```

    Replace placeholders with actual values as needed.

2. **Load Configuration and Build the Client**: Use the `BuildClientWithConfigFile` function to read the configuration from the file and initialize the Jamf Pro client.

    ```go
    configFilePath := "path_to_your/client_config.json"
    client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
    if err != nil {
        log.Fatalf("Failed to build Jamf Pro client with configuration file: %v", err)
    }
    ```

    This method will load the configuration from the specified file and use it to set up the Jamf Pro client.

### Summary

Both methods provide a flexible way to configure and initialize the Jamf Pro client, allowing you to choose the approach that best fits your deployment strategy and environment. Remember to handle credentials securely and avoid exposing sensitive information in your code or public repositories.


## Calling SDK Functions

Once the Jamf Pro client is configured and initialized, you can start making API calls to perform various operations on your Jamf Pro instance. This section provides examples of common SDK functions you might want to use.

### Fetching Device Details

To fetch details about a specific device, you can use the `GetComputerByID` function. You will need the device's unique identifier (such as a serial number) to retrieve its details.

```go
// Assuming 'client' is your initialized Jamf Pro client
deviceID := "your_jamf_computer_id"
deviceDetails, err := client.GetComputerByID(deviceID)
if err != nil {
    log.Fatalf("Failed to get device details: %v", err)
}

// Use 'deviceDetails' as needed
fmt.Printf("Device Name: %s\n", deviceDetails.General.DeviceName)
```


## Go SDK for Jamf Pro API Progress Tracker

### API Coverage Progress

Date: Feb-2024
Maintainer: [ShocOne]

## Overview

This document tracks the progress of API endpoint coverage tests. As endpoints are tested, they will be marked as covered.

## Coverage Legend

- ✅ - Covered
- ❌ - Not Covered
- ⚠️ - Information

## Endpoints

### Accounts - /JSSResource/accounts

This documentation outlines the operations available for Jamf Pro Accounts and Account Groups.

## Operations

- [x] ✅ **GET** `/JSSResource/accounts`
  - `GetAccounts` operation retrieves all user accounts.

- [x] ✅ **GET** `/JSSResource/accounts/userid/{id}`
  - `GetAccountByID` operation retrieves the Account by its ID.

- [x] ✅ **GET** `/JSSResource/accounts/username/{name}`
  - `GetAccountByName` operation retrieves the Account by its name.

- [x] ✅ **GET** `/JSSResource/accounts/groupid/{id}`
  - `GetAccountGroupByID` operation retrieves the Account Group by its ID.

- [x] ✅ **GET** `/JSSResource/accounts/groupname/{name}`
  - `GetAccountGroupByName` operation retrieves the Account Group by its name.

- [x] ✅ **POST** `/JSSResource/accounts/userid/{id}` 
  - `CreateAccount` operation creates a new Jamf Pro Account.

- [x] ✅ **POST** `/JSSResource/accounts/groupid/{id}`
  - `CreateAccountGroup` operation creates a new Jamf Pro Account Group.

- [x] ✅ **PUT** `/JSSResource/accounts/userid/{id}`
  - `UpdateAccountByID` operation updates an existing Jamf Pro Account by ID.

- [x] ✅ **PUT** `/JSSResource/accounts/username/{name}`
  - `UpdateAccountByName` operation updates an existing Jamf Pro Account by Name.

- [x] ✅ **PUT** `/JSSResource/accounts/groupid/{id}`
  - `UpdateAccountGroupByID` operation updates an existing Jamf Pro Account Group by ID.

- [x] ✅ **PUT** `/JSSResource/accounts/groupname/{name}`
  - `UpdateAccountGroupByName` operation updates an existing Jamf Pro Account Group by Name.

- [x] ✅ **DELETE** `/JSSResource/accounts/userid/{id}`
  - `DeleteAccountByID` operation deletes an existing Jamf Pro Account by ID.

- [x] ✅ **DELETE** `/JSSResource/accounts/username/{name}`
  - `DeleteAccountByName` operation deletes an existing Jamf Pro Account by Name.

- [x] ✅ **DELETE** `/JSSResource/accounts/groupid/{id}`
  - `DeleteAccountGroupByID` operation deletes an existing Jamf Pro Account Group by ID.

- [x] ✅ **DELETE** `/JSSResource/accounts/groupname/{name}`
  - `DeleteAccountGroupByName` operation deletes an existing Jamf Pro Account Group by Name.

## Summary

- Total Endpoints Covered: 5
  - `/JSSResource/accounts`
  - `/JSSResource/accounts/userid/{id}`
  - `/JSSResource/accounts/username/{name}`
  - `/JSSResource/accounts/groupid/{id}`
  - `/JSSResource/accounts/groupname/{name}`

- Total Operations Covered: 15


### Activation Code - /JSSResource/activationcode

This documentation outlines the operations available for Activation Code in Jamf Pro.

## Operations

- [x] ✅ **GET** `/JSSResource/activationcode`
  - `GetActivationCode` operation retrieves the current activation code and organization name.

- [x] ✅ **PUT** `/JSSResource/activationcode`
  - `UpdateActivationCode` operation updates the activation code with a new organization name and code.

## Summary

- Total Endpoints Covered: 2
  - `/JSSResource/activationcode`

- Total Operations Covered: 2


### Jamf Pro API Integrations - /api/v1/api-integrations

This documentation outlines the operations available for Jamf API Integrations.

## Operations

- [x] ✅ **GET** `/api/v1/api-integrations`
  - `GetApiIntegrations` operation fetches all API integrations.

- [x] ✅ **GET** `/api/v1/api-integrations/{id}`
  - `GetApiIntegrationByID` operation fetches an API integration by its ID.

- [x] ✅ **GET** `/api/v1/api-integrations` followed by searching by name
  - `GetApiIntegrationNameByID` operation fetches an API integration by its display name and then retrieves its details using its ID.

- [x] ✅ **POST**  `/api/v1/api-integrations`
  - `CreateApiIntegration` operation creates a new API integration.

- [x] ✅ **POST**  `/api/v1/api-integrations/{id}/client-credentials`
  - `CreateClientCredentialsByApiRoleID` operation creates new client credentials for an API integration by its ID.

- [x] ✅ **PUT** `/api/v1/api-integrations/{id}`
  - `UpdateApiIntegrationByID` operation updates an API integration by its ID.

- [x] ✅ **PUT** `/api/v1/api-integrations` followed by searching by name
  - `UpdateApiIntegrationByName` operation updates an API integration based on its display name.

- [x] ✅ **POST**  `/api/v1/api-integrations/{id}/client-credentials` (Used for updating)
  - `UpdateClientCredentialsByApiIntegrationID` operation updates client credentials for an API integration by its ID.

- [x] ✅ **DELETE** `/api/v1/api-integrations/{id}`
  - `DeleteApiIntegrationByID` operation deletes an API integration by its ID.

- [x] ✅ **DELETE** `/api/v1/api-integrations` followed by searching by name
  - `DeleteApiIntegrationByName` operation deletes an API integration by its display name.

## Summary

- Total Endpoints Covered: 3
  - `/api/v1/api-integrations`
  - `/api/v1/api-integrations/{id}`
  - `/api/v1/api-integrations` followed by searching by name

- Total Operations Covered: 8


### Jamf Pro API Role Privileges - /api/v1/api-role-privileges

This documentation outlines the operations available for Jamf API Role Privileges.

## Operations

- [x] ✅ **GET** `/api/v1/api-role-privileges`
  - `GetJamfAPIPrivileges` operation fetches a list of Jamf API role privileges.

- [x] ✅ **GET** `/api/v1/api-role-privileges/search?name={name}&limit={limit}`
  - `GetJamfAPIPrivilegesByName` operation fetches Jamf API role privileges by name.

## Summary

- Total Endpoints Covered: 2
  - `/api/v1/api-role-privileges`
  - `/api/v1/api-role-privileges/search?name={name}&limit={limit}`

- Total Operations Covered: 2



### Jamf Pro API Roles - /api/v1/api-roles

This documentation outlines the operations available for Jamf API Roles.

## Operations

- [x] ✅ **GET** `/api/v1/api-roles`
  - `GetJamfAPIRoles` operation fetches all API roles.

- [x] ✅ **GET** `/api/v1/api-roles/{id}`
  - `GetJamfApiRolesByID` operation fetches a Jamf API role by its ID.

- [x] ✅ **GET** `/api/v1/api-roles` followed by searching by name
  - `GetJamfApiRolesNameById` operation fetches a Jamf API role by its display name and then retrieves its details using its ID.

- [x] ✅ **POST**  `/api/v1/api-roles`
  - `CreateJamfApiRole` operation creates a new Jamf API role.

- [x] ✅ **PUT** `/api/v1/api-roles/{id}`
  - `UpdateJamfApiRoleByID` operation updates a Jamf API role by its ID.

- [x] ✅ **PUT** `/api/v1/api-roles` followed by searching by name
  - `UpdateJamfApiRoleByName` operation updates a Jamf API role based on its display name.

- [x] ✅ **DELETE** `/api/v1/api-roles/{id}`
  - `DeleteJamfApiRoleByID` operation deletes a Jamf API role by its ID.

- [x] ✅ **DELETE** `/api/v1/api-roles` followed by searching by name
  - `DeleteJamfApiRoleByName` operation deletes a Jamf API role by its display name.

## Summary

- Total Endpoints Covered: 3
  - `/api/v1/api-roles`
  - `/api/v1/api-roles/{id}`
  - `/api/v1/api-roles` followed by searching by name

- Total Operations Covered: 8


### Jamf Pro Classic API - Advanced Computer Searches

This documentation outlines the operations available for Advanced Computer Searches.

## Operations

- [x] ✅ **GET** `/JSSResource/advancedcomputersearches`
  - `GetAdvancedComputerSearches` operation fetches all advanced computer searches.

- [x] ✅ **GET** `/JSSResource/advancedcomputersearches/id/{id}`
  - `GetAdvancedComputerSearchByID` operation fetches an advanced computer search by its ID.

- [x] ✅ **GET** `/JSSResource/advancedcomputersearches/name/{name}`
  - `GetAdvancedComputerSearchesByName` operation fetches advanced computer searches by their name.

- [x] ✅ **POST**  `/JSSResource/advancedcomputersearches`
  - `CreateAdvancedComputerSearch` operation creates a new advanced computer search.

- [x] ✅ **PUT** `/JSSResource/advancedcomputersearches/id/{id}`
  - `UpdateAdvancedComputerSearchByID` operation updates an existing advanced computer search by its ID.

- [x] ✅ **PUT** `/JSSResource/advancedcomputersearches/name/{name}`
  - `UpdateAdvancedComputerSearchByName` operation updates an advanced computer search by its name.

- [x] ✅ **DELETE** `/JSSResource/advancedcomputersearches/id/{id}`
  - `DeleteAdvancedComputerSearchByID` operation deletes an advanced computer search by its ID.

- [x] ✅ **DELETE** `/JSSResource/advancedcomputersearches/name/{name}`
  - `DeleteAdvancedComputerSearchByName` operation deletes an advanced computer search by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/advancedcomputersearches`
  - `/JSSResource/advancedcomputersearches/id/{id}`
  - `/JSSResource/advancedcomputersearches/name/{name}`

- Total Operations Covered: 8


### Jamf Pro Classic API - Advanced Mobile Device Searches

This documentation outlines the operations available for Advanced Mobile Device Searches.

## Operations

- [x] ✅ **GET** `/JSSResource/advancedmobiledevicesearches`
  - `GetAdvancedMobileDeviceSearches` operation fetches all advanced mobile device searches.

- [x] ✅ **GET** `/JSSResource/advancedmobiledevicesearches/id/{id}`
  - `GetAdvancedMobileDeviceSearchByID` operation fetches an advanced mobile device search by its ID.

- [x] ✅ **GET** `/JSSResource/advancedmobiledevicesearches/name/{name}`
  - `GetAdvancedMobileDeviceSearchByName` operation fetches advanced mobile device searches by their name.

- [x] ✅ **POST**  `/JSSResource/advancedmobiledevicesearches`
  - `CreateAdvancedMobileDeviceSearch` operation creates a new advanced mobile device search.

- [x] ✅ **PUT** `/JSSResource/advancedmobiledevicesearches/id/{id}`
  - `UpdateAdvancedMobileDeviceSearchByID` operation updates an existing advanced mobile device search by its ID.

- [x] ✅ **PUT** `/JSSResource/advancedmobiledevicesearches/name/{name}`
  - `UpdateAdvancedMobileDeviceSearchByName` operation updates an advanced mobile device search by its name.

- [x] ✅ **DELETE** `/JSSResource/advancedmobiledevicesearches/id/{id}`
  - `DeleteAdvancedMobileDeviceSearchByID` operation deletes an advanced mobile device search by its ID.

- [x] ✅ **DELETE** `/JSSResource/advancedmobiledevicesearches/name/{name}`
  - `DeleteAdvancedMobileDeviceSearchByName` operation deletes an advanced mobile device search by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/advancedmobiledevicesearches`
  - `/JSSResource/advancedmobiledevicesearches/id/{id}`
  - `/JSSResource/advancedmobiledevicesearches/name/{name}`

- Total Operations Covered: 8


### Jamf Pro Classic API - Advanced User Searches

This documentation outlines the operations available for Advanced User Searches.

## Operations

- [x] ✅ **GET** `/JSSResource/advancedusersearches`
  - `GetAdvancedUserSearches` operation fetches all advanced user searches.

- [x] ✅ **GET** `/JSSResource/advancedusersearches/id/{id}`
  - `GetAdvancedUserSearchByID` operation fetches an advanced user search by its ID.

- [x] ✅ **GET** `/JSSResource/advancedusersearches/name/{name}`
  - `GetAdvancedUserSearchesByName` operation fetches advanced user searches by their name.

- [x] ✅ **POST**  `/JSSResource/advancedusersearches`
  - `CreateAdvancedUserSearch` operation creates a new advanced user search.

- [x] ✅ **PUT** `/JSSResource/advancedusersearches/id/{id}`
  - `UpdateAdvancedUserSearchByID` operation updates an existing advanced user search by its ID.

- [x] ✅ **PUT** `/JSSResource/advancedusersearches/name/{name}`
  - `UpdateAdvancedUserSearchByName` operation updates an advanced user search by its name.

- [x] ✅ **DELETE** `/JSSResource/advancedusersearches/id/{id}`
  - `DeleteAdvancedUserSearchByID` operation deletes an advanced user search by its ID.

- [x] ✅ **DELETE** `/JSSResource/advancedusersearches/name/{name}`
  - `DeleteAdvancedUserSearchByName` operation deletes an advanced user search by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/advancedusersearches`
  - `/JSSResource/advancedusersearches/id/{id}`
  - `/JSSResource/advancedusersearches/name/{name}`

- Total Operations Covered: 8


### Allowed File Extensions - /JSSResource/allowedfileextensions

This documentation outlines the operations available for Allowed File Extensions.

## Operations

- [x] ✅ **GET** `/JSSResource/allowedfileextensions`
  - `GetAllowedFileExtensions` operation retrieves all allowed file extensions.

- [x] ✅ **GET** `/JSSResource/allowedfileextensions/id/{id}`
  - `GetAllowedFileExtensionByID` operation retrieves the allowed file extension by its ID.

- [x] ✅ **GET** `/JSSResource/allowedfileextensions/extension/{extensionName}`
  - `GetAllowedFileExtensionByName` operation retrieves the allowed file extension by its name.

- [x] ✅ **POST**  `/JSSResource/allowedfileextensions/id/0`
  - `CreateAllowedFileExtension` operation creates a new allowed file extension.

- [] ⚠️ **PUT** `/JSSResource/allowedfileextensions/id/{id}`
  - `UpdateAllowedFileExtensionByID` (API doesn't support update).

- [x] ✅ **DELETE** `/JSSResource/allowedfileextensions/id/{id}`
  - `DeleteAllowedFileExtensionByID` operation deletes an existing allowed file extension by ID.

- [x] ✅ **DELETE** `/JSSResource/allowedfileextensions/extension/{extensionName}`
  - `DeleteAllowedFileExtensionByNameByID` operation deletes an existing allowed file extension by resolving its name to an ID.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/allowedfileextensions`
  - `/JSSResource/allowedfileextensions/id/{id}`
  - `/JSSResource/allowedfileextensions/extension/{extensionName}`

- Total Operations Covered: 6


### BYO Profiles - `/JSSResource/byoprofiles`

This documentation outlines the operations available for BYO profiles.

## Operations

- [x] ✅ **GET** `/JSSResource/byoprofiles`
  - `GetBYOProfiles` operation retrieves all BYO profiles.

- [x] ✅ **GET** `/JSSResource/byoprofiles/id/{id}`
  - `GetBYOProfileByID` operation retrieves a BYO profile by its ID.

- [x] ✅ **GET** `/JSSResource/byoprofiles/name/{name}`
  - `GetBYOProfileByName` operation retrieves a BYO profile by its name.

- [x] ✅ **POST**  `/JSSResource/byoprofiles/id/0`
  - `CreateBYOProfile` operation creates a new BYO profile.

- [x] ✅ **PUT** `/JSSResource/byoprofiles/id/{id}`
  - `UpdateBYOProfileByID` operation updates an existing BYO profile by its ID.

- [x] ✅ **PUT** `/JSSResource/byoprofiles/name/{oldName}`
  - `UpdateBYOProfileByName` operation updates an existing BYO profile by its name.

- [x] ✅ **DELETE** `/JSSResource/byoprofiles/id/{id}`
  - `DeleteBYOProfileByID` operation deletes an existing BYO profile by its ID.

- [x] ✅ **DELETE** `/JSSResource/byoprofiles/name/{name}`
  - `DeleteBYOProfileByName` operation deletes an existing BYO profile by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/byoprofiles`
  - `/JSSResource/byoprofiles/id/{id}`
  - `/JSSResource/byoprofiles/name/{name}`

- Total Operations Covered: 8



### Jamf Pro API - Categories

This documentation outlines the operations available for categories using the API.

## Operations

- [x] ✅ **GET** `/api/v1/categories`
  - `GetCategories` operation retrieves categories based on query parameters.

- [x] ✅ **GET** `/api/v1/categories/{id}`
  - `GetCategoryByID` operation retrieves a category by its ID.

- [x] ✅ **GET** `/api/v1/categories/name/{name}`
  - `GetCategoryNameByID` operation retrieves a category by its name and then retrieves its details using its ID.

- [x] ✅ **POST**  `/api/v1/categories`
  - `CreateCategory` operation creates a new category.

- [x] ✅ **PUT** `/api/v1/categories/{id}`
  - `UpdateCategoryByID` operation updates an existing category by its ID.

- [x] ✅ **PUT** `UpdateCategoryByNameByID`
  - `UpdateCategoryByNameByID` operation updates a category by its name and then updates its details using its ID.

- [x] ✅ **DELETE** `/api/v1/categories/{id}`
  - `DeleteCategoryByID` operation deletes a category by its ID.

- [x] ✅ **DELETE** `DeleteCategoryByNameByID`
  - `DeleteCategoryByNameByID` operation deletes a category by its name after inferring its ID.

- [x] ✅ **POST**  `/api/v1/categories/delete-multiple`
  - `DeleteMultipleCategoriesByID` operation deletes multiple categories by their IDs.

## Summary

- Total Endpoints Covered: 3
  - `/api/v1/categories`
  - `/api/v1/categories/{id}`
  - `/api/v1/categories/name/{name}`

- Total Operations Covered: 9


### Jamf Pro Classic API - Computer Groups

This documentation outlines the operations available for computer groups using the Classic API.

## Operations

- [x] ✅ **GET** `/JSSResource/computergroups`
  - `GetComputerGroups` operation fetches all computer groups.

- [x] ✅ **GET** `/JSSResource/computergroups/id/{id}`
  - `GetComputerGroupByID` operation fetches a computer group by its ID.

- [x] ✅ **GET** `/JSSResource/computergroups/name/{name}`
  - `GetComputerGroupByName` operation fetches a computer group by its name.

- [x] ✅ **POST**  `/JSSResource/computergroups/id/0`
  - `CreateComputerGroup` operation creates a new computer group.

- [x] ✅ **PUT** `/JSSResource/computergroups/id/{id}`
  - `UpdateComputerGroupByID` operation updates an existing computer group by its ID.

- [x] ✅ **PUT** `/JSSResource/computergroups/name/{name}`
  - `UpdateComputerGroupByName` operation updates a computer group by its name.

- [x] ✅ **DELETE** `/JSSResource/computergroups/id/{id}`
  - `DeleteComputerGroupByID` operation deletes a computer group by its ID.

- [x] ✅ **DELETE** `/JSSResource/computergroups/name/{name}`
  - `DeleteComputerGroupByName` operation deletes a computer group by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/computergroups`
  - `/JSSResource/computergroups/id/{id}`
  - `/JSSResource/computergroups/name/{name}`

- Total Operations Covered: 8



### macOS Configuration Profiles - /JSSResource/osxconfigurationprofiles

This documentation outlines the operations available for macOS configuration profiles using the API.

## Operations

- [x] ✅ **GET** `/JSSResource/osxconfigurationprofiles`
  - `GetMacOSConfigurationProfiles` operation retrieves all macOS configuration profiles.

- [x] ✅ **GET** `/JSSResource/osxconfigurationprofiles/id/{id}`
  - `GetMacOSConfigurationProfileByID` operation retrieves the macOS configuration profile by its ID.

- [x] ✅ **GET** `/JSSResource/osxconfigurationprofiles/name/{name}`
  - `GetMacOSConfigurationProfileByName` operation retrieves the macOS configuration profile by its name.

- [x] ✅ **POST**  `/JSSResource/osxconfigurationprofiles/id/0`
  - `CreateMacOSConfigurationProfile` operation creates a new macOS configuration profile.

- [x] ✅ **PUT** `/JSSResource/osxconfigurationprofiles/id/{id}`
  - `UpdateMacOSConfigurationProfileByID` operation updates an existing macOS configuration profile by ID.

- [x] ✅ **PUT** `/JSSResource/osxconfigurationprofiles/name/{name}`
  - `UpdateMacOSConfigurationProfileByName` operation updates an existing macOS configuration profile by its name.

- [x] ✅ **DELETE** `/JSSResource/osxconfigurationprofiles/id/{id}`
  - `DeleteMacOSConfigurationProfileByID` operation deletes an existing macOS configuration profile by ID.

- [x] ✅ **DELETE** `/JSSResource/osxconfigurationprofiles/name/{name}`
  - `DeleteMacOSConfigurationProfileByName` operation deletes an existing macOS configuration profile by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/osxconfigurationprofiles`
  - `/JSSResource/osxconfigurationprofiles/id/{id}`
  - `/JSSResource/osxconfigurationprofiles/name/{name}`

- Total Operations Covered: 8


### Departments - /JSSResource/departments

This documentation outlines the operations available for departments using the API.

## Operations

- [x] ✅ **GET** `/JSSResource/departments`
  - `GetDepartments` operation retrieves all departments.

- [x] ✅ **GET** `/JSSResource/departments/id/{id}`
  - `GetDepartmentByID` operation retrieves the department by its ID.

- [x] ✅ **GET** `/JSSResource/departments/name/{name}`
  - `GetDepartmentByName` operation retrieves the department by its name.

- [x] ✅ **POST**  `/JSSResource/departments/id/0`
  - `CreateDepartment` operation creates a new department.

- [x] ✅ **PUT** `/JSSResource/departments/id/{id}`
  - `UpdateDepartmentByID` operation updates an existing department.

- [x] ✅ **PUT** `/JSSResource/departments/name/{oldName}`
  - `UpdateDepartmentByName` operation updates an existing department by its name.

- [x] ✅ **DELETE** `/JSSResource/departments/id/{id}`
  - `DeleteDepartmentByID` operation deletes an existing department by its ID.

- [x] ✅ **DELETE** `/JSSResource/departments/name/{name}`
  - `DeleteDepartmentByName` operation deletes an existing department by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/departments`
  - `/JSSResource/departments/id/{id}`
  - `/JSSResource/departments/name/{name}`

- Total Operations Covered: 8

### Policies - /JSSResource/policies

This documentation outlines the operations available for policies using the API.

## Operations

- [x] ✅ **GET** `/JSSResource/policies`
  - `GetPolicies` operation retrieves a list of all policies.

- [x] ✅ **GET** `/JSSResource/policies/id/{id}`
  - `GetPolicyByID` operation retrieves the details of a policy by its ID.

- [x] ✅ **GET** `/JSSResource/policies/name/{name}`
  - `GetPolicyByName` operation retrieves a policy by its name.

- [x] ✅ **GET** `/JSSResource/policies/category/{category}`
  - `GetPolicyByCategory` operation retrieves policies by their category.

- [x] ✅ **GET** `/JSSResource/policies/createdBy/{createdBy}`
  - `GetPoliciesByType` operation retrieves policies by the type of entity that created them.

- [x] ✅ **POST**  `/JSSResource/policies/id/0`
  - `CreatePolicy` operation creates a new policy.

- [x] ✅ **PUT** `/JSSResource/policies/id/{id}`
  - `UpdatePolicyByID` operation updates an existing policy by its ID.

- [x] ✅ **PUT** `/JSSResource/policies/name/{name}`
  - `UpdatePolicyByName` operation updates an existing policy by its name.

- [x] ✅ **DELETE** `/JSSResource/policies/id/{id}`
  - `DeletePolicyByID` operation deletes a policy by its ID.

- [x] ✅ **DELETE** `/JSSResource/policies/name/{name}`
  - `DeletePolicyByName` operation deletes a policy by its name.

## Summary

- Total Endpoints Covered: 5
  - `/JSSResource/policies`
  - `/JSSResource/policies/id/{id}`
  - `/JSSResource/policies/name/{name}`
  - `/JSSResource/policies/category/{category}`
  - `/JSSResource/policies/createdBy/{createdBy}`

- Total Operations Covered: 10


### Jamf Pro API - Self Service Branding macOS

This documentation outlines the operations available for self-service branding configurations for macOS using the API.

## Operations

- [x] ✅ **GET** `/api/v1/self-service/branding/macos`
  - `GetSelfServiceBrandingMacOS` operation fetches all self-service branding configurations for macOS.

- [x] ✅ **GET** `/api/v1/self-service/branding/macos/{id}`
  - `GetSelfServiceBrandingMacOSByID` operation fetches a self-service branding configuration for macOS by its ID.

- [x] ✅ **GET** `/api/v1/self-service/branding/macos/name/{name}`
  - `GetSelfServiceBrandingMacOSByNameByID` operation fetches a self-service branding configuration for macOS by its name.

- [x] ✅ **POST**  `/api/v1/self-service/branding/macos`
  - `CreateSelfServiceBrandingMacOS` operation creates a new self-service branding configuration for macOS.

- [x] ✅ **PUT** `/api/v1/self-service/branding/macos/{id}`
  - `UpdateSelfServiceBrandingMacOSByID` operation updates an existing self-service branding configuration for macOS by its ID.

- [x] ✅ **PUT** - `UpdateSelfServiceBrandingMacOSByName` operation updates a self-service branding configuration for macOS by its name.

- [x] ✅ **DELETE** `/api/v1/self-service/branding/macos/{id}`
  - `DeleteSelfServiceBrandingMacOSByID` operation deletes a self-service branding configuration for macOS by its ID.

- [x] ✅ **DELETE** - `DeleteSelfServiceBrandingMacOSByName` operation deletes a self-service branding configuration for macOS by its name.

## Summary

- Total Endpoints Covered: 4
  - `/api/v1/self-service/branding/macos`
  - `/api/v1/self-service/branding/macos/{id}`
  - `/api/v1/self-service/branding/macos/name/{name}`
  - `/api/v1/self-service/branding/macos`

- Total Operations Covered: 8


### Jamf Pro Classic API - Scripts

This documentation outlines the operations available for scripts using the API.

## Operations

- [x] ✅ **GET** `/JSSResource/scripts`
  - `GetScripts` operation retrieves all scripts.

- [x] ✅ **GET** `/JSSResource/scripts/id/{id}`
  - `GetScriptsByID` operation retrieves the script details by its ID.

- [x] ✅ **GET** `/JSSResource/scripts/name/{name}`
  - `GetScriptsByName` operation retrieves the script details by its name.

- [x] ✅ **POST**  `/JSSResource/scripts/id/0`
  - `CreateScriptByID` operation creates a new script.

- [x] ✅ **PUT** `/JSSResource/scripts/id/{id}`
  - `UpdateScriptByID` operation updates an existing script by its ID.

- [x] ✅ **PUT** `/JSSResource/scripts/name/{name}`
  - `UpdateScriptByName` operation updates an existing script by its name.

- [x] ✅ **DELETE** `/JSSResource/scripts/id/{id}`
  - `DeleteScriptByID` operation deletes an existing script by its ID.

- [x] ✅ **DELETE** `/JSSResource/scripts/name/{name}`
  - `DeleteScriptByName` operation deletes an existing script by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/scripts`
  - `/JSSResource/scripts/id/{id}`
  - `/JSSResource/scripts/name/{name}`

- Total Operations Covered: 8

### Jamf Pro Classic API - Sites

This documentation outlines the operations available for sites using the API.

## Operations

- [x] ✅ **GET** `/JSSResource/sites`
  - `GetSites` operation fetches all sites.

- [x] ✅ **GET** `/JSSResource/sites/id/{id}`
  - `GetSiteByID` operation fetches a site by its ID.

- [x] ✅ **GET** `/JSSResource/sites/name/{name}`
  - `GetSiteByName` operation fetches a site by its name.

- [x] ✅ **POST**  `/JSSResource/sites/id/0`
  - `CreateSite` operation creates a new site.

- [x] ✅ **PUT** `/JSSResource/sites/id/{id}`
  - `UpdateSiteByID` operation updates an existing site by its ID.

- [x] ✅ **PUT** `/JSSResource/sites/name/{name}`
  - `UpdateSiteByName` operation updates a site by its name.

- [x] ✅ **DELETE** `/JSSResource/sites/id/{id}`
  - `DeleteSiteByID` operation deletes a site by its ID.

- [x] ✅ **DELETE** `/JSSResource/sites/name/{name}`
  - `DeleteSiteByName` operation deletes a site by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/sites`
  - `/JSSResource/sites/id/{id}`
  - `/JSSResource/sites/name/{name}`

- Total Operations Covered: 8


### Jamf Pro API - SSO Failover

This documentation outlines the operations available for SSO Failover using the API.

## Operations

- [x] ✅ **GET** `/api/v1/sso/failover`
  - `GetSSOFailoverSettings` operation retrieves the current failover settings.

- [x] ✅ **PUT** `/api/v1/sso/failover/generate`
  - `UpdateFailoverUrl` operation updates the failover URL by changing the failover key to a new one and returns new failover settings.

## Summary

- Total Endpoints Covered: 2
  - `/api/v1/sso/failover`
  - `/api/v1/sso/failover/generate`

- Total Operations Covered: 2


### Jamf Pro API - Volume Purchasing Subscriptions

This documentation provides details on the API endpoints available for managing Volume Purchasing Subscriptions within Jamf Pro.

## Operations

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

## Summary

- Total Endpoints Covered: 2
  - `/api/v1/volume-purchasing-subscriptions`
  - `/api/v1/volume-purchasing-subscriptions/{id}`

- Total Operations Covered: 5
- Total Custom Operations Covered: 3

### Jamf Pro API - Computer Inventory Collection Settings

This documentation outlines the API endpoints available for managing Computer Inventory Collection Settings in Jamf Pro.

## Operations

- [x] ✅ **GET** `/api/v1/computer-inventory-collection-settings`  
  `GetComputerInventoryCollectionSettings` retrieves the current computer inventory collection preferences and custom paths.

- [x] ✅ **PATCH** `/api/v1/computer-inventory-collection-settings`  
  `UpdateComputerInventoryCollectionSettings` updates the computer inventory collection preferences.

- [x] ✅ **POST** `/api/v1/computer-inventory-collection-settings/custom-path`  
  `CreateComputerInventoryCollectionSettingsCustomPath` creates a new custom path for the computer inventory collection settings.

- [x] ✅ **DELETE** `/api/v1/computer-inventory-collection-settings/custom-path/{id}`  
  `DeleteComputerInventoryCollectionSettingsCustomPathByID` deletes a custom path by its ID.

## Summary

- Total Endpoints Covered: 3
  - `/api/v1/computer-inventory-collection-settings`
  - `/api/v1/computer-inventory-collection-settings/custom-path`
  - `/api/v1/computer-inventory-collection-settings/custom-path/{id}`

- Total Operations Covered: 4

### Jamf Pro API - Jamf Pro Information

This documentation covers the API endpoints available for retrieving information about the Jamf Pro server.

## Operations

- [x] ✅ **GET** `/api/v2/jamf-pro-information`  
  `GetJamfProInformation` retrieves information about various services enabled on the Jamf Pro server, like VPP token, DEP account status, BYOD, and more.

## Summary

- Total Endpoints Covered: 1
  - `/api/v2/jamf-pro-information`

- Total Operations Covered: 1
	
### Jamf Pro Classic API - Classes

This documentation provides details on the API endpoints available for managing classes within Jamf Pro using the Classic API which requires XML data structure support.

## Operations

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

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/classes`
  - `/JSSResource/classes/id/{id}`
  - `/JSSResource/classes/name/{name}`

- Total Operations Covered: 8

### Jamf Pro Classic API - Computer Invitations

This documentation outlines the API endpoints available for managing computer invitations within Jamf Pro using the Classic API, which relies on XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/computerinvitations`
GetComputerInvitations retrieves a list of all computer invitations.

- [x] ✅ **GET** `/JSSResource/computerinvitations/id/{id}`
GetComputerInvitationByID fetches a single computer invitation by its ID.

- [x] ✅ **GET** `/JSSResource/computerinvitations/invitation/{invitation}`
GetComputerInvitationsByInvitationID retrieves a computer invitation by its invitation ID.

- [x] ✅ **POST**  `/JSSResource/computerinvitations/id/0`
CreateComputerInvitation creates a new computer invitation. Using ID 0 indicates creation as per API pattern. If siteId is not included, it defaults to using a siteId of -1, implying no specific site association.

- [] ❌ **PUT** `/JSSResource/computerinvitations/invitation/{invitation}`
There is no documented endpoint for updating a computer invitation by its invitation ID.

- [x] ✅ **DELETE** `/JSSResource/computerinvitations/id/{id}`
DeleteComputerInvitationByID deletes a computer invitation by its ID.

- [] ❌ **DELETE** `/JSSResource/computerinvitations/invitation/{invitation}`
There is currently no SDK coverage for deleting an invitation by invitation ID

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/computerinvitations`
  - `/JSSResource/computerinvitations/id/{id}`
  - `/JSSResource/computerinvitations/invitation/{invitation}`

- Total Operations Covered: 5
- Total Operations Not Covered: 3

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

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/diskencryptionconfigurations`
  - `/JSSResource/diskencryptionconfigurations/id/{id}`
  - `/JSSResource/diskencryptionconfigurations/name/{name}`

- Total Operations Covered: 8

### Jamf Pro Classic API - Distribution Points

This documentation outlines the operations available for managing Distribution Points within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/distributionpoints`
  - `GetDistributionPoints` operation retrieves a serialized list of all distribution points.

- [x] ✅ **GET** `/JSSResource/distributionpoints/id/{id}`
  - `GetDistributionPointByID` operation fetches a single distribution point by its ID.

- [x] ✅ **GET** `/JSSResource/distributionpoints/name/{name}`
  - `GetDistributionPointByName` operation retrieves a distribution point by its name.

- [x] ✅ **POST** `/JSSResource/distributionpoints/id/0`
  - `CreateDistributionPoint` operation creates a new distribution point with the provided details. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/distributionpoints/id/{id}`
  - `UpdateDistributionPointByID` operation updates an existing distribution point by its ID.

- [x] ✅ **PUT** `/JSSResource/distributionpoints/name/{name}`
  - `UpdateDistributionPointByName` operation updates an existing distribution point by its name.

- [x] ✅ **DELETE** `/JSSResource/distributionpoints/id/{id}`
  - `DeleteDistributionPointByID` operation deletes a distribution point by its ID.

- [x] ✅ **DELETE** `/JSSResource/distributionpoints/name/{name}`
  - `DeleteDistributionPointByName` operation deletes a distribution point by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/distributionpoints`
  - `/JSSResource/distributionpoints/id/{id}`
  - `/JSSResource/distributionpoints/name/{name}`

- Total Operations Covered: 8


### Jamf Pro Classic API - Directory Bindings

This documentation outlines the operations available for managing Directory Bindings within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/directorybindings`
  - `GetDirectoryBindings` operation retrieves a serialized list of all directory bindings.

- [x] ✅ **GET** `/JSSResource/directorybindings/id/{id}`
  - `GetDirectoryBindingByID` operation fetches a single directory binding by its ID.

- [x] ✅ **GET** `/JSSResource/directorybindings/name/{name}`
  - `GetDirectoryBindingByName` operation retrieves a directory binding by its name.

- [x] ✅ **POST** `/JSSResource/directorybindings/id/0`
  - `CreateDirectoryBinding` operation creates a new directory binding with the provided details. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/directorybindings/id/{id}`
  - `UpdateDirectoryBindingByID` operation updates an existing directory binding by its ID.

- [x] ✅ **PUT** `/JSSResource/directorybindings/name/{name}`
  - `UpdateDirectoryBindingByName` operation updates an existing directory binding by its name.

- [x] ✅ **DELETE** `/JSSResource/directorybindings/id/{id}`
  - `DeleteDirectoryBindingByID` operation deletes a directory binding by its ID.

- [x] ✅ **DELETE** `/JSSResource/directorybindings/name/{name}`
  - `DeleteDirectoryBindingByName` operation deletes a directory binding by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/directorybindings`
  - `/JSSResource/directorybindings/id/{id}`
  - `/JSSResource/directorybindings/name/{name}`

- Total Operations Covered: 8


### Jamf Pro Classic API - Computers

This documentation outlines the operations available for managing Computers within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/computers`
  - `GetComputers` operation retrieves a serialized list of all computers.

- [x] ✅ **GET** `/JSSResource/computers/id/{id}`
  - `GetComputerByID` operation fetches a single computer by its ID.

- [x] ✅ **GET** `/JSSResource/computers/name/{name}`
  - `GetComputerByName` operation retrieves a computer by its name.

- [] ❌ **GET** `/JSSResource/computers/subset/basic`
  - `GetComputerByBasicDataSubset` operation retrieves a basic data about a computer.

- [] ❌ **GET** `/JSSResource/computers/match/{match}`
  - `GetComputerBySearchTerm` operation retrieves a Match and performs the same function as a simple search in the GUI.

- [] ❌ **GET** `/JSSResource/computers/match/name/{matchname}`
  - `GetComputerByNameParameter` operation retrieves a Match and performs the same function as a simple search in the GUI.

- [] ❌ **GET** `/JSSResource/computers/id/{id}/subset/{subset}`
  - `GetComputerByIDAndDataSubset` Subset values can also be appended using an ampersand to return multiple subsets (e.g. /subsets/General&Location).

- [] ❌ **GET** `/JSSResource/computers/name/{name}/subset/{subset}`
  - `GetComputerByNameAndDataSubset` Subset values can also be appended using an ampersand to return multiple subsets (e.g. /subsets/General&Location).

- [] ❌ **GET** `/JSSResource/computers/udid/{udid}`
  - `GetComputerByUUID` operation retrieves a computer by its UUID.

- [] ❌ **GET** `/JSSResource/computers/udid/{udid}/subset/{subset}`
  - `GetComputerByUUIDAndDataSubset` operation retrieves a computer by its UUID and a data subset.

- [] ❌ **GET** `/JSSResource/computers/serialnumber/{serialnumber}`
  - `GetComputerBySerialNumber` operation retrieves a computer by its serial number.

- [] ❌ **GET** `/JSSResource/computers/serialnumber/{serialnumber}/subset/{subset}`
  - `GetComputerBySerialNumberAndDataSubset` operation retrieves a computer by its Serial Number and a data subset.

- [] ❌ **GET** `/JSSResource/computers/macaddress/{macaddress}`
  - `GetComputerByMACAddress` operation retrieves a computer by its MAC Address.

- [] ❌ **GET** `/JSSResource/computers/macaddress/{macaddress}/subset/{subset}`
  - `GetComputerByMACAddressAndDataSubset` operation retrieves a computer by its MAC Address and a data subset.

- [x] ✅ **POST** `/JSSResource/computers/id/0`
  - `CreateComputer` operation creates a new computer with the provided details. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/computers/id/{id}`
  - `UpdateComputerByID` operation updates an existing computer by its ID.

- [x] ✅ **PUT** `/JSSResource/computers/name/{name}`
  - `UpdateComputerByName` operation updates an existing computer by its name.

- [] ❌ **PUT** `/JSSResource/computers/udid/{udid}`
  - `UpdateComputerByUUID` operation updates an existing computer by its UUID.

- [] ❌ **PUT** `/JSSResource/computers/serialnumber/{serialnumber}`
  - `UpdateComputerBySerialNumber` operation updates an existing computer by its Serial Number.

- [] ❌ **PUT** `/JSSResource/computers/macaddress/{macaddress}`
  - `UpdateComputerByMacAddress` operation updates an existing computer by its Mac Address.

- [x] ✅ **DELETE** `/JSSResource/computers/id/{id}`
  - `DeleteComputerByID` operation deletes a computer by its ID.

- [x] ✅ **DELETE** `/JSSResource/computers/name/{name}`
  - `DeleteComputerByName` operation deletes a computer by its name.

- [] ❌ **DELETE** `/JSSResource/computers/udid/{udid}`
  - `DeleteComputerByUUID operation deletes a computer by its UUID.

- [] ❌ **DELETE** `/JSSResource/computers/serialnumber/{serialnumber}`
  - `DeleteComputerBySerialNumber` operation deletes a computer by its Serial Number.

- [] ❌ **DELETE** `/JSSResource/computers/macaddress/{macaddress}`
  - `DeleteComputerByMacAddress` operation deletes a computer by its Mac Address.

- [] ❌ **DELETE** `/JSSResource/computers/extensionattributedataflush/id/{id}`
  - `Deletes data collected by an extension attribute` operation Deletes data collected by an extension attribute.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/computers`
  - `/JSSResource/computers/id/{id}`
  - `/JSSResource/computers/name/{name}`

- Total Operations Covered: 8
- Total Operations Not Covered: 18


### Jamf Pro Classic API - Dock Items

This documentation outlines the operations available for managing Dock Items within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/dockitems`
  - `GetDockItems` operation retrieves a serialized list of all dock items.

- [x] ✅ **GET** `/JSSResource/dockitems/id/{id}`
  - `GetDockItemByID` operation fetches a single dock item by its ID.

- [x] ✅ **GET** `/JSSResource/dockitems/name/{name}`
  - `GetDockItemByName` operation retrieves a dock item by its name.

- [x] ✅ **POST** `/JSSResource/dockitems/id/0`
  - `CreateDockItem` operation creates a new dock item with the provided details. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/dockitems/id/{id}`
  - `UpdateDockItemByID` operation updates an existing dock item by its ID.

- [x] ✅ **PUT** `/JSSResource/dockitems/name/{name}`
  - `UpdateDockItemByName` operation updates an existing dock item by its name.

- [x] ✅ **DELETE** `/JSSResource/dockitems/id/{id}`
  - `DeleteDockItemByID` operation deletes a dock item by its ID.

- [x] ✅ **DELETE** `/JSSResource/dockitems/name/{name}`
  - `DeleteDockItemByName` operation deletes a dock item by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/dockitems`
  - `/JSSResource/dockitems/id/{id}`
  - `/JSSResource/dockitems/name/{name}`

- Total Operations Covered: 8


### Jamf Pro Classic API - eBooks

This documentation outlines the operations available for managing eBooks within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/ebooks`
  - `GetEbooks` operation retrieves a serialized list of all eBooks.

- [x] ✅ **GET** `/JSSResource/ebooks/id/{id}`
  - `GetEbookByID` operation fetches a single eBook by its ID.

- [x] ✅ **GET** `/JSSResource/ebooks/name/{name}`
  - `GetEbookByName` operation retrieves an eBook by its name.

- [x] ✅ **GET** `/JSSResource/ebooks/name/{name}/subset/{subset}`
  - `GetEbooksByNameAndDataSubset` operation retrieves a specific subset (General, Scope, or SelfService) of an eBook by its name.

- [x] ✅ **POST** `/JSSResource/ebooks/id/0`
  - `CreateEbook` operation creates a new eBook with the provided details. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/ebooks/id/{id}`
  - `UpdateEbookByID` operation updates an existing eBook by its ID.

- [x] ✅ **PUT** `/JSSResource/ebooks/name/{name}`
  - `UpdateEbookByName` operation updates an existing eBook by its name.

- [x] ✅ **DELETE** `/JSSResource/ebooks/id/{id}`
  - `DeleteEbookByID` operation deletes an eBook by its ID.

- [x] ✅ **DELETE** `/JSSResource/ebooks/name/{name}`
  - `DeleteEbookByName` operation deletes an eBook by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/ebooks`
  - `/JSSResource/ebooks/id/{id}`
  - `/JSSResource/ebooks/name/{name}`

- Total Operations Covered: 9


### Jamf Pro Classic API - VPP Mac Applications

This documentation outlines the operations available for managing VPP Mac applications within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/macapplications`
  - `GetMacApplications` operation retrieves a serialized list of all VPP Mac applications.

- [x] ✅ **GET** `/JSSResource/macapplications/id/{id}`
  - `GetMacApplicationByID` operation fetches a single Mac application by its ID.

- [x] ✅ **GET** `/JSSResource/macapplications/name/{name}`
  - `GetMacApplicationByName` operation retrieves a Mac application by its name.

- [x] ✅ **GET** `/JSSResource/macapplications/name/{name}/subset/{subset}`
  - `GetMacApplicationByNameAndDataSubset` operation retrieves a specific subset (General, Scope, SelfService, VPPCodes, and VPP) of a Mac application by its name.

- [x] ✅ **GET** `/JSSResource/macapplications/id/{id}/subset/{subset}`
  - `GetMacApplicationByIDAndDataSubset` operation retrieves a specific subset (General, Scope, SelfService, VPPCodes, and VPP) of a Mac application by its ID.

- [x] ✅ **POST** `/JSSResource/macapplications/id/0`
  - `CreateMacApplication` operation creates a new Mac application with the provided details. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/macapplications/id/{id}`
  - `UpdateMacApplicationByID` operation updates an existing Mac application by its ID.

- [x] ✅ **PUT** `/JSSResource/macapplications/name/{name}`
  - `UpdateMacApplicationByName` operation updates an existing Mac application by its name.

- [x] ✅ **DELETE** `/JSSResource/macapplications/id/{id}`
  - `DeleteMacApplicationByID` operation deletes a Mac application by its ID.

- [x] ✅ **DELETE** `/JSSResource/macapplications/name/{name}`
  - `DeleteMacApplicationByName` operation deletes a Mac application by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/macapplications`
  - `/JSSResource/macapplications/id/{id}`
  - `/JSSResource/macapplications/name/{name}`

- Total Operations Covered: 10

### Jamf Pro Classic API - iBeacons

This documentation outlines the operations available for managing iBeacons within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/ibeacons`
  - `GetIBeacons` operation retrieves a serialized list of all iBeacons.

- [x] ✅ **GET** `/JSSResource/ibeacons/id/{id}`
  - `GetIBeaconByID` operation fetches a single iBeacon by its ID.

- [x] ✅ **GET** `/JSSResource/ibeacons/name/{name}`
  - `GetIBeaconByName` operation retrieves an iBeacon by its name.

- [x] ✅ **POST** `/JSSResource/ibeacons/id/0`
  - `CreateIBeacon` operation creates a new iBeacon with the provided details. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/ibeacons/id/{id}`
  - `UpdateIBeaconByID` operation updates an existing iBeacon by its ID.

- [x] ✅ **PUT** `/JSSResource/ibeacons/name/{name}`
  - `UpdateIBeaconByName` operation updates an existing iBeacon by its name.

- [x] ✅ **DELETE** `/JSSResource/ibeacons/id/{id}`
  - `DeleteIBeaconByID` operation deletes an iBeacon by its ID.

- [x] ✅ **DELETE** `/JSSResource/ibeacons/name/{name}`
  - `DeleteIBeaconByName` operation deletes an iBeacon by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/ibeacons`
  - `/JSSResource/ibeacons/id/{id}`
  - `/JSSResource/ibeacons/name/{name}`

- Total Operations Covered: 8


### Jamf Pro Classic API - LDAP Servers

This documentation outlines the operations available for managing LDAP servers within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/ldapservers`
  - `GetLDAPServers` operation retrieves a serialized list of all LDAP servers.

- [x] ✅ **GET** `/JSSResource/ldapservers/id/{id}`
  - `GetLDAPServerByID` operation fetches a single LDAP server by its ID.

- [x] ✅ **GET** `/JSSResource/ldapservers/name/{name}`
  - `GetLDAPServerByName` operation retrieves an LDAP server by its name.

- [x] ✅ **GET** `/JSSResource/ldapservers/id/{id}/user/{user}`
  - `GetLDAPServerByIDAndUserDataSubset` operation retrieves user data for a specific LDAP server by its ID.

- [x] ✅ **GET** `/JSSResource/ldapservers/id/{id}/group/{group}`
  - `GetLDAPServerByIDAndGroupDataSubset` operation retrieves group data for a specific LDAP server by its ID.

- [x] ✅ **GET** `/JSSResource/ldapservers/id/{id}/group/{group}/user/{user}`
  - `GetLDAPServerByIDAndUserMembershipInGroupDataSubset` operation retrieves user group membership details for a specific LDAP server by its ID.

- [x] ✅ **GET** `/JSSResource/ldapservers/name/{name}/user/{user}`
  - `GetLDAPServerByNameAndUserDataSubset` operation retrieves user data for a specific LDAP server by its name.

- [x] ✅ **GET** `/JSSResource/ldapservers/name/{name}/group/{group}`
  - `GetLDAPServerByNameAndGroupDataSubset` operation retrieves group data for a specific LDAP server by its name.

- [x] ✅ **GET** `/JSSResource/ldapservers/name/{name}/group/{group}/user/{user}`
  - `GetLDAPServerByNameAndUserMembershipInGroupDataSubset` operation retrieves user group membership details for a specific LDAP server by its name.

- [x] ✅ **POST** `/JSSResource/ldapservers/id/0`
  - `CreateLDAPServer` operation creates a new LDAP server with the provided details.

- [x] ✅ **PUT** `/JSSResource/ldapservers/id/{id}`
  - `UpdateLDAPServerByID` operation updates an existing LDAP server by its ID.

- [x] ✅ **PUT** `/JSSResource/ldapservers/name/{name}`
  - `UpdateLDAPServerByName` operation updates an existing LDAP server by its name.

- [x] ✅ **DELETE** `/JSSResource/ldapservers/id/{id}`
  - `DeleteLDAPServerByID` operation deletes an LDAP server by its ID.

- [x] ✅ **DELETE** `/JSSResource/ldapservers/name/{name}`
  - `DeleteLDAPServerByName` operation deletes an LDAP server by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/ldapservers`
  - `/JSSResource/ldapservers/id/{id}`
  - `/JSSResource/ldapservers/name/{name}`

- Total Operations Covered: 14


### Jamf Pro Classic API - Licensed Software

This documentation outlines the operations available for managing Licensed Software within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/licensedsoftware`
  - `GetLicensedSoftware` operation retrieves a serialized list of all Licensed Software.

- [x] ✅ **GET** `/JSSResource/licensedsoftware/id/{id}`
  - `GetLicensedSoftwareByID` operation fetches details of a single Licensed Software item by its ID.

- [x] ✅ **GET** `/JSSResource/licensedsoftware/name/{name}`
  - `GetLicensedSoftwareByName` operation retrieves details of a Licensed Software item by its name.

- [x] ✅ **POST** `/JSSResource/licensedsoftware/id/0`
  - `CreateLicensedSoftware` operation creates a new Licensed Software item. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/licensedsoftware/id/{id}`
  - `UpdateLicensedSoftwareByID` operation updates an existing Licensed Software item by its ID.

- [x] ✅ **PUT** `/JSSResource/licensedsoftware/name/{name}`
  - `UpdateLicensedSoftwareByName` operation updates an existing Licensed Software item by its name.

- [x] ✅ **DELETE** `/JSSResource/licensedsoftware/id/{id}`
  - `DeleteLicensedSoftwareByID` operation deletes a Licensed Software item by its ID.

- [x] ✅ **DELETE** `/JSSResource/licensedsoftware/name/{name}`
  - `DeleteLicensedSoftwareByName` operation deletes a Licensed Software item by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/licensedsoftware`
  - `/JSSResource/licensedsoftware/id/{id}`
  - `/JSSResource/licensedsoftware/name/{name}`

- Total Operations Covered: 8


### Jamf Pro Classic API - Mobile Device Applications

This documentation outlines the operations available for managing Mobile Device Applications within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/mobiledeviceapplications`
  - `GetMobileDeviceApplications` operation retrieves a serialized list of all Mobile Device Applications.

- [x] ✅ **GET** `/JSSResource/mobiledeviceapplications/id/{id}`
  - `GetMobileDeviceApplicationByID` operation fetches details of a single Mobile Device Application by its ID.

- [x] ✅ **GET** `/JSSResource/mobiledeviceapplications/name/{name}`
  - `GetMobileDeviceApplicationByName` operation retrieves details of a Mobile Device Application by its name.

- [x] ✅ **GET** `/JSSResource/mobiledeviceapplications/bundleid/{bundleid}`
  - `GetMobileDeviceApplicationByAppBundleID` operation fetches details of a Mobile Device Application by its Bundle ID.

- [x] ✅ **GET** `/JSSResource/mobiledeviceapplications/bundleid/{bundleid}/version/{version}`
  - `GetMobileDeviceApplicationByAppBundleIDAndVersion` operation fetches details of a Mobile Device Application by its Bundle ID and specific version.

- [x] ✅ **GET** `/JSSResource/mobiledeviceapplications/id/{id}/subset/{subset}`
  - `GetMobileDeviceApplicationByIDAndDataSubset` operation fetches a Mobile Device Application by its ID and a specified data subset.

- [x] ✅ **GET** `/JSSResource/mobiledeviceapplications/name/{name}/subset/{subset}`
  - `GetMobileDeviceApplicationByNameAndDataSubset` operation fetches a Mobile Device Application by its name and a specified data subset.

- [x] ✅ **POST** `/JSSResource/mobiledeviceapplications/id/0`
  - `CreateMobileDeviceApplication` operation creates a new Mobile Device Application. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceapplications/id/{id}`
  - `UpdateMobileDeviceApplicationByID` operation updates an existing Mobile Device Application by its ID.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceapplications/name/{name}`
  - `UpdateMobileDeviceApplicationByName` operation updates an existing Mobile Device Application by its name.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceapplications/bundleid/{bundleid}`
  - `UpdateMobileDeviceApplicationByApplicationBundleID` operation updates an existing Mobile Device Application by its Bundle ID.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceapplications/bundleid/{bundleid}/version/{version}`
  - `UpdateMobileDeviceApplicationByIDAndAppVersion` operation updates an existing Mobile Device Application by its ID and specific version.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceapplications/id/{id}`
  - `DeleteMobileDeviceApplicationByID` operation deletes a Mobile Device Application by its ID.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceapplications/name/{name}`
  - `DeleteMobileDeviceApplicationByName` operation deletes a Mobile Device Application by its name.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceapplications/bundleid/{bundleid}`
  - `DeleteMobileDeviceApplicationByBundleID` operation deletes a Mobile Device Application by its Bundle ID.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceapplications/bundleid/{bundleid}/version/{version}`
  - `DeleteMobileDeviceApplicationByBundleIDAndVersion` operation deletes a Mobile Device Application by its Bundle ID and specific version.

## Summary

- Total Endpoints Covered: 4
  - `/JSSResource/mobiledeviceapplications`
  - `/JSSResource/mobiledeviceapplications/id/{id}`
  - `/JSSResource/mobiledeviceapplications/name/{name}`
  - `/JSSResource/mobiledeviceapplications/bundleid/{bundleid}`

- Total Operations Covered: 14


### Jamf Pro Classic API - Mobile Device Configuration Profiles

This documentation outlines the operations available for managing Mobile Device Configuration Profiles within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/mobiledeviceconfigurationprofiles`
  - `GetMobileDeviceConfigurationProfiles` operation retrieves a serialized list of all Mobile Device Configuration Profiles.

- [x] ✅ **GET** `/JSSResource/mobiledeviceconfigurationprofiles/id/{id}`
  - `GetMobileDeviceConfigurationProfileByID` operation fetches details of a single Mobile Device Configuration Profile by its ID.

- [x] ✅ **GET** `/JSSResource/mobiledeviceconfigurationprofiles/name/{name}`
  - `GetMobileDeviceConfigurationProfileByName` operation retrieves details of a Mobile Device Configuration Profile by its name.

- [x] ✅ **GET** `/JSSResource/mobiledeviceconfigurationprofiles/id/{id}/subset/{subset}`
  - `GetMobileDeviceConfigurationProfileByIDBySubset` operation fetches a specific Mobile Device Configuration Profile by its ID and a specified subset.

- [x] ✅ **GET** `/JSSResource/mobiledeviceconfigurationprofiles/name/{name}/subset/{subset}`
  - `GetMobileDeviceConfigurationProfileByNameBySubset` operation fetches a specific Mobile Device Configuration Profile by its name and a specified subset.

- [x] ✅ **POST** `/JSSResource/mobiledeviceconfigurationprofiles/id/0`
  - `CreateMobileDeviceConfigurationProfile` operation creates a new Mobile Device Configuration Profile. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceconfigurationprofiles/id/{id}`
  - `UpdateMobileDeviceConfigurationProfileByID` operation updates an existing Mobile Device Configuration Profile by its ID.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceconfigurationprofiles/name/{name}`
  - `UpdateMobileDeviceConfigurationProfileByName` operation updates an existing Mobile Device Configuration Profile by its name.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceconfigurationprofiles/id/{id}`
  - `DeleteMobileDeviceConfigurationProfileByID` operation deletes a Mobile Device Configuration Profile by its ID.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceconfigurationprofiles/name/{name}`
  - `DeleteMobileDeviceConfigurationProfileByName` operation deletes a Mobile Device Configuration Profile by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/mobiledeviceconfigurationprofiles`
  - `/JSSResource/mobiledeviceconfigurationprofiles/id/{id}`
  - `/JSSResource/mobiledeviceconfigurationprofiles/name/{name}`

- Total Operations Covered: 10


### Jamf Pro Classic API - Mobile Extension Attributes

This documentation outlines the operations available for managing Mobile Extension Attributes within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/mobiledeviceextensionattributes`
  - `GetMobileExtensionAttributes` operation retrieves a serialized list of all Mobile Extension Attributes.

- [x] ✅ **GET** `/JSSResource/mobiledeviceextensionattributes/id/{id}`
  - `GetMobileExtensionAttributeByID` operation fetches details of a single Mobile Extension Attribute by its ID.

- [x] ✅ **GET** `/JSSResource/mobiledeviceextensionattributes/name/{name}`
  - `GetMobileExtensionAttributeByName` operation retrieves details of a Mobile Extension Attribute by its name.

- [x] ✅ **POST** `/JSSResource/mobiledeviceextensionattributes/id/0`
  - `CreateMobileExtensionAttribute` operation creates a new Mobile Extension Attribute. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceextensionattributes/id/{id}`
  - `UpdateMobileExtensionAttributeByID` operation updates an existing Mobile Extension Attribute by its ID.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceextensionattributes/name/{name}`
  - `UpdateMobileExtensionAttributeByName` operation updates an existing Mobile Extension Attribute by its name.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceextensionattributes/id/{id}`
  - `DeleteMobileExtensionAttributeByID` operation deletes a Mobile Extension Attribute by its ID.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceextensionattributes/name/{name}`
  - `DeleteMobileExtensionAttributeByName` operation deletes a Mobile Extension Attribute by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/mobiledeviceextensionattributes`
  - `/JSSResource/mobiledeviceextensionattributes/id/{id}`
  - `/JSSResource/mobiledeviceextensionattributes/name/{name}`

- Total Operations Covered: 8


### Jamf Pro Classic API - Mobile Device Enrollment Profiles

This documentation outlines the operations available for managing Mobile Device Enrollment Profiles within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/mobiledeviceenrollmentprofiles`
  - `GetMobileDeviceEnrollmentProfiles` operation retrieves a serialized list of all Mobile Device Enrollment Profiles.

- [x] ✅ **GET** `/JSSResource/mobiledeviceenrollmentprofiles/id/{id}`
  - `GetMobileDeviceEnrollmentProfileByID` operation fetches details of a single Mobile Device Enrollment Profile by its ID.

- [x] ✅ **GET** `/JSSResource/mobiledeviceenrollmentprofiles/name/{name}`
  - `GetMobileDeviceEnrollmentProfileByName` operation retrieves details of a Mobile Device Enrollment Profile by its name.

- [x] ✅ **GET** `/JSSResource/mobiledeviceenrollmentprofiles/invitation/{invitation}`
  - `GetProfileByInvitation` operation fetches a Mobile Device Enrollment Profile by its invitation.

- [x] ✅ **GET** `/JSSResource/mobiledeviceenrollmentprofiles/id/{id}/subset/{subset}`
  - `GetMobileDeviceEnrollmentProfileByIDBySubset` operation fetches a specific Mobile Device Enrollment Profile by its ID and a specified subset.

- [x] ✅ **GET** `/JSSResource/mobiledeviceenrollmentprofiles/name/{name}/subset/{subset}`
  - `GetMobileDeviceEnrollmentProfileByNameBySubset` operation fetches a specific Mobile Device Enrollment Profile by its name and a specified subset.

- [x] ✅ **POST** `/JSSResource/mobiledeviceenrollmentprofiles/id/0`
  - `CreateMobileDeviceEnrollmentProfile` operation creates a new Mobile Device Enrollment Profile. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceenrollmentprofiles/id/{id}`
  - `UpdateMobileDeviceEnrollmentProfileByID` operation updates an existing Mobile Device Enrollment Profile by its ID.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceenrollmentprofiles/name/{name}`
  - `UpdateMobileDeviceEnrollmentProfileByName` operation updates an existing Mobile Device Enrollment Profile by its name.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceenrollmentprofiles/invitation/{invitation}`
  - `UpdateMobileDeviceEnrollmentProfileByInvitation` operation updates an existing Mobile Device Enrollment Profile by its invitation.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceenrollmentprofiles/id/{id}`
  - `DeleteMobileDeviceEnrollmentProfileByID` operation deletes a Mobile Device Enrollment Profile by its ID.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceenrollmentprofiles/name/{name}`
  - `DeleteMobileDeviceEnrollmentProfileByName` operation deletes a Mobile Device Enrollment Profile by its name.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceenrollmentprofiles/invitation/{invitation}`
  - `DeleteMobileDeviceEnrollmentProfileByInvitation` operation deletes a Mobile Device Enrollment Profile by its invitation.

## Summary

- Total Endpoints Covered: 4
  - `/JSSResource/mobiledeviceenrollmentprofiles`
  - `/JSSResource/mobiledeviceenrollmentprofiles/id/{id}`
  - `/JSSResource/mobiledeviceenrollmentprofiles/name/{name}`
  - `/JSSResource/mobiledeviceenrollmentprofiles/invitation/{invitation}`

- Total Operations Covered: 12


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

This documentation outlines the operations available for managing Network Segments within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/networksegments`
  - `GetNetworkSegments` operation retrieves a serialized list of all Network Segments.

- [x] ✅ **GET** `/JSSResource/networksegments/id/{id}`
  - `GetNetworkSegmentByID` operation fetches details of a single Network Segment by its ID.

- [x] ✅ **GET** `/JSSResource/networksegments/name/{name}`
  - `GetNetworkSegmentByName` operation retrieves details of a Network Segment by its name.

- [x] ✅ **POST** `/JSSResource/networksegments/id/0`
  - `CreateNetworkSegment` operation creates a new Network Segment. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/networksegments/id/{id}`
  - `UpdateNetworkSegmentByID` operation updates an existing Network Segment by its ID.

- [x] ✅ **PUT** `/JSSResource/networksegments/name/{name}`
  - `UpdateNetworkSegmentByName` operation updates an existing Network Segment by its name.

- [x] ✅ **DELETE** `/JSSResource/networksegments/id/{id}`
  - `DeleteNetworkSegmentByID` operation deletes a Network Segment by its ID.

- [x] ✅ **DELETE** `/JSSResource/networksegments/name/{name}`
  - `DeleteNetworkSegmentByName` operation deletes a Network Segment by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/networksegments`
  - `/JSSResource/networksegments/id/{id}`
  - `/JSSResource/networksegments/name/{name}`

- Total Operations Covered: 8

### Jamf Pro Classic API - Mobile Device Groups

This documentation outlines the operations available for managing Mobile Device Groups within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/mobiledevicegroups`
  - `GetMobileDeviceGroups` operation retrieves a serialized list of all Mobile Device Groups.

- [x] ✅ **GET** `/JSSResource/mobiledevicegroups/id/{id}`
  - `GetMobileDeviceGroupByID` operation fetches details of a single Mobile Device Group by its ID.

- [x] ✅ **GET** `/JSSResource/mobiledevicegroups/name/{name}`
  - `GetMobileDeviceGroupByName` operation retrieves details of a Mobile Device Group by its name.

- [x] ✅ **POST** `/JSSResource/mobiledevicegroups/id/0`
  - `CreateMobileDeviceGroup` operation creates a new Mobile Device Group. The ID `0` in the endpoint indicates creation.

- [x] ✅ **PUT** `/JSSResource/mobiledevicegroups/id/{id}`
  - `UpdateMobileDeviceGroupByID` operation updates an existing Mobile Device Group by its ID.

- [x] ✅ **PUT** `/JSSResource/mobiledevicegroups/name/{name}`
  - `UpdateMobileDeviceGroupByName` operation updates an existing Mobile Device Group by its name.

- [x] ✅ **DELETE** `/JSSResource/mobiledevicegroups/id/{id}`
  - `DeleteMobileDeviceGroupByID` operation deletes a Mobile Device Group by its ID.

- [x] ✅ **DELETE** `/JSSResource/mobiledevicegroups/name/{name}`
  - `DeleteMobileDeviceGroupByName` operation deletes a Mobile Device Group by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/mobiledevicegroups`
  - `/JSSResource/mobiledevicegroups/id/{id}`
  - `/JSSResource/mobiledevicegroups/name/{name}`

- Total Operations Covered: 8


### Jamf Pro Classic API - Mobile Device Provisioning Profiles

This documentation outlines the operations available for managing Mobile Device Provisioning Profiles within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/mobiledeviceprovisioningprofiles`
  - `GetMobileDeviceProvisioningProfiles` operation retrieves a serialized list of all Mobile Device Provisioning Profiles.

- [x] ✅ **GET** `/JSSResource/mobiledeviceprovisioningprofiles/id/{id}`
  - `GetMobileDeviceProvisioningProfileByID` operation fetches a specific Mobile Device Provisioning Profile by its ID.

- [x] ✅ **GET** `/JSSResource/mobiledeviceprovisioningprofiles/name/{name}`
  - `GetMobileDeviceProvisioningProfileByName` operation fetches a specific Mobile Device Provisioning Profile by its name.

- [x] ✅ **GET** `/JSSResource/mobiledeviceprovisioningprofiles/uuid/{uuid}`
  - `GetMobileDeviceProvisioningProfileByUUID` operation fetches a specific Mobile Device Provisioning Profile by its UUID.

- [x] ✅ **POST** `/JSSResource/mobiledeviceprovisioningprofiles/id/{id}`
  - `CreateMobileDeviceProvisioningProfileByID` operation creates a new Mobile Device Provisioning Profile by its ID.

- [x] ✅ **POST** `/JSSResource/mobiledeviceprovisioningprofiles/name/{name}`
  - `CreateMobileDeviceProvisioningProfileByName` operation creates a new Mobile Device Provisioning Profile by its name.

- [x] ✅ **POST** `/JSSResource/mobiledeviceprovisioningprofiles/uuid/{uuid}`
  - `CreateMobileDeviceProvisioningProfileByUUID` operation creates a new Mobile Device Provisioning Profile by its UUID.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceprovisioningprofiles/id/{id}`
  - `UpdateMobileDeviceProvisioningProfileByID` operation updates an existing Mobile Device Provisioning Profile by its ID.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceprovisioningprofiles/name/{name}`
  - `UpdateMobileDeviceProvisioningProfileByName` operation updates an existing Mobile Device Provisioning Profile by its name.

- [x] ✅ **PUT** `/JSSResource/mobiledeviceprovisioningprofiles/uuid/{uuid}`
  - `UpdateMobileDeviceProvisioningProfileByUUID` operation updates an existing Mobile Device Provisioning Profile by its UUID.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceprovisioningprofiles/id/{id}`
  - `DeleteMobileDeviceProvisioningProfileByID` operation deletes a Mobile Device Provisioning Profile by its ID.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceprovisioningprofiles/name/{name}`
  - `DeleteMobileDeviceProvisioningProfileByName` operation deletes a Mobile Device Provisioning Profile by its name.

- [x] ✅ **DELETE** `/JSSResource/mobiledeviceprovisioningprofiles/uuid/{uuid}`
  - `DeleteMobileDeviceProvisioningProfileByUUID` operation deletes a Mobile Device Provisioning Profile by its UUID.

## Summary

- Total Endpoints Covered: 4
  - `/JSSResource/mobiledeviceprovisioningprofiles`
  - `/JSSResource/mobiledeviceprovisioningprofiles/id/{id}`
  - `/JSSResource/mobiledeviceprovisioningprofiles/name/{name}`
  - `/JSSResource/mobiledeviceprovisioningprofiles/uuid/{uuid}`

- Total Operations Covered: 12

### Jamf Pro API - Buildings

This documentation outlines the operations available for managing Buildings within Jamf Pro using the API, which supports JSON data structures.

## Operations

- [x] ✅ **GET** `/api/v1/buildings`
  - `GetBuildings` operation retrieves a serialized list of all buildings.

- [x] ✅ **GET** `/api/v1/buildings/{id}`
  - `GetBuildingByID` operation fetches a specific building by its ID.

- [x] ✅ **GET** `/api/v1/buildings/{id}/history`
  - `GetBuildingResourceHistoryByID` operation retrieves the resource history of a specific building by its ID.

- [x] ✅ **POST** `/api/v1/buildings`
  - `CreateBuilding` operation creates a new building.

- [x] ✅ **PUT** `/api/v1/buildings/{id}`
  - `UpdateBuildingByID` operation updates an existing building by its ID.

- [x] ✅ **POST** `/api/v1/buildings/{id}/history`
  - `CreateBuildingResourceHistoryByID` operation updates the resource history of a building by its ID.

- [x] ✅ **DELETE** `/api/v1/buildings/{id}`
  - `DeleteBuildingByID` operation deletes a building by its ID.

- [x] ✅ **POST** `/api/v1/buildings/delete-multiple`
  - `DeleteMultipleBuildingsByID` operation deletes multiple buildings by their IDs.

- [] ❌ **POST** `/api/v1/buildings/{id}/history/export`
  - `ExportBuildingResourceHistoryByID` operation (not implemented/available).

## Summary

- Total Endpoints Covered: 4
  - `/api/v1/buildings`
  - `/api/v1/buildings/{id}`
  - `/api/v1/buildings/{id}/history`
  - `/api/v1/buildings/delete-multiple`

- Total Operations Covered: 8
- Total Operations Not Covered: 1

### Jamf Pro Classic API - Users

This documentation outlines the operations available for managing Users within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/users`
  - `GetUsers` operation retrieves a serialized list of all Users.

- [x] ✅ **GET** `/JSSResource/users/id/{id}`
  - `GetUserByID` operation fetches a specific User by their ID.

- [x] ✅ **GET** `/JSSResource/users/name/{name}`
  - `GetUserByName` operation fetches a specific User by their name.

- [x] ✅ **GET** `/JSSResource/users/email/{email}`
  - `GetUserByEmail` operation fetches a specific User by their email.

- [x] ✅ **POST** `/JSSResource/users/id/0`
  - `CreateUser` operation creates a new User.

- [x] ✅ **PUT** `/JSSResource/users/id/{id}`
  - `UpdateUserByID` operation updates an existing User by their ID.

- [x] ✅ **PUT** `/JSSResource/users/name/{name}`
  - `UpdateUserByName` operation updates an existing User by their name.

- [x] ✅ **PUT** `/JSSResource/users/email/{email}`
  - `UpdateUserByEmail` operation updates an existing User by their email.

- [x] ✅ **DELETE** `/JSSResource/users/id/{id}`
  - `DeleteUserByID` operation deletes a User by their ID.

- [x] ✅ **DELETE** `/JSSResource/users/name/{name}`
  - `DeleteUserByName` operation deletes a User by their name.

- [x] ✅ **DELETE** `/JSSResource/users/email/{email}`
  - `DeleteUserByEmail` operation deletes a User by their email.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/users`
  - `/JSSResource/users/id/{id}`
  - `/JSSResource/users/name/{name}`
  - `/JSSResource/users/email/{email}`

- Total Operations Covered: 11


### Jamf Pro Classic API - User Groups

This documentation outlines the operations available for managing User Groups within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/usergroups`
  - `GetUserGroups` operation retrieves a serialized list of all User Groups.

- [x] ✅ **GET** `/JSSResource/usergroups/id/{id}`
  - `GetUserGroupsByID` operation fetches a specific User Group by its ID.

- [x] ✅ **GET** `/JSSResource/usergroups/name/{name}`
  - `GetUserGroupsByName` operation fetches a specific User Group by its name.

- [x] ✅ **POST** `/JSSResource/usergroups/id/0`
  - `CreateUserGroup` operation creates a new User Group.

- [x] ✅ **PUT** `/JSSResource/usergroups/id/{id}`
  - `UpdateUserGroupByID` operation updates an existing User Group by its ID.

- [x] ✅ **PUT** `/JSSResource/usergroups/name/{name}`
  - `UpdateUserGroupByName` operation updates an existing User Group by its name.

- [x] ✅ **DELETE** `/JSSResource/usergroups/id/{id}`
  - `DeleteUserGroupByID` operation deletes a User Group by its ID.

- [x] ✅ **DELETE** `/JSSResource/usergroups/name/{name}`
  - `DeleteUserGroupByName` operation deletes a User Group by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/usergroups`
  - `/JSSResource/usergroups/id/{id}`
  - `/JSSResource/usergroups/name/{name}`

- Total Operations Covered: 8

### Jamf Pro Classic API - User Extension Attributes

This documentation outlines the operations available for managing User Extension Attributes within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/userextensionattributes`
  - `GetUserExtensionAttributes` operation retrieves a serialized list of all User Extension Attributes.

- [x] ✅ **GET** `/JSSResource/userextensionattributes/id/{id}`
  - `GetUserExtensionAttributeByID` operation fetches a specific User Extension Attribute by its ID.

- [x] ✅ **GET** `/JSSResource/userextensionattributes/name/{name}`
  - `GetUserExtensionAttributeByName` operation fetches a specific User Extension Attribute by its name.

- [x] ✅ **POST** `/JSSResource/userextensionattributes/id/0`
  - `CreateUserExtensionAttribute` operation creates a new User Extension Attribute.

- [x] ✅ **PUT** `/JSSResource/userextensionattributes/id/{id}`
  - `UpdateUserExtensionAttributeByID` operation updates an existing User Extension Attribute by its ID.

- [x] ✅ **PUT** `/JSSResource/userextensionattributes/name/{name}`
  - `UpdateUserExtensionAttributeByName` operation updates an existing User Extension Attribute by its name.

- [x] ✅ **DELETE** `/JSSResource/userextensionattributes/id/{id}`
  - `DeleteUserExtensionAttributeByID` operation deletes a User Extension Attribute by its ID.

- [x] ✅ **DELETE** `/JSSResource/userextensionattributes/name/{name}`
  - `DeleteUserExtensionAttributeByName` operation deletes a User Extension Attribute by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/userextensionattributes`
  - `/JSSResource/userextensionattributes/id/{id}`
  - `/JSSResource/userextensionattributes/name/{name}`

- Total Operations Covered: 8

### Jamf Pro Classic API - Mobile Devices

This documentation details the operations available for managing Mobile Devices within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/mobiledevices`
  - `GetMobileDevices` operation retrieves a serialized list of all mobile devices.

- [x] ✅ **GET** `/JSSResource/mobiledevices/id/{id}`
  - `GetMobileDeviceByID` operation fetches a specific mobile device by its ID.

- [x] ✅ **GET** `/JSSResource/mobiledevices/name/{name}`
  - `GetMobileDeviceByName` operation fetches a specific mobile device by its name.

- [x] ✅ **GET** `/JSSResource/mobiledevices/id/{id}/subset/{subset}`
  - `GetMobileDeviceByIDAndDataSubset` operation retrieves a specific subset of data for a mobile device by its ID.

- [x] ✅ **GET** `/JSSResource/mobiledevices/name/{name}/subset/{subset}`
  - `GetMobileDeviceByNameAndDataSubset` operation retrieves a specific subset of data for a mobile device by its name.

- [] ❌ **GET** `/JSSResource/mobiledevices/match/{match}`
  - `GetMobileDeviceBySearchTerm` operation retrieves a Match and performs the same function as a simple search in the GUI.

- [] ❌ **GET** `/JSSResource/mobiledevices/udid/{udid}`
  - `GetMobileDeviceByUUID` operation retrieves a mobile device by its UUID.

- [] ❌ **GET** `/JSSResource/mobiledevices/udid/{udid}/subset/{subset}`
  - `GetMobileDeviceByUUIDAndDataSubset` operation retrieves a mobile device by its UUID and a data subset.

- [] ❌ **GET** `/JSSResource/mobiledevices/serialnumber/{serialnumber}`
  - `GetMobileDeviceBySerialNumber` operation retrieves a mobile device by its serial number.

- [] ❌ **GET** `/JSSResource/mobiledevices/serialnumber/{serialnumber}/subset/{subset}`
  - `GetMobileDeviceBySerialNumberAndDataSubset` operation retrieves a mobile device by its Serial Number and a data subset.

- [x] ✅ **POST** `/JSSResource/mobiledevices/id/0`
  - `CreateMobileDevice` operation creates a new mobile device.

- [x] ✅ **PUT** `/JSSResource/mobiledevices/id/{id}`
  - `UpdateMobileDeviceByID` operation updates an existing mobile device by its ID.

- [x] ✅ **PUT** `/JSSResource/mobiledevices/name/{name}`
  - `UpdateMobileDeviceByName` operation updates an existing mobile device by its name.

- [] ❌ **PUT** `/JSSResource/mobiledevices/udid/{udid}`
  - `UpdateMobileDeviceByUDID` operation updates an existing mobile device by its UDID.

- [] ❌ **PUT** `/JSSResource/mobiledevices/serialnumber/{serialnumber}`
  - `UpdateMobileDeviceBySerialNumber` operation updates an existing mobile device by its Serial number.

- [] ❌ **PUT** `/JSSResource/mobiledevices/macaddress/{macaddress}`
  - `UpdateMobileDeviceByMACAddress` operation updates an existing mobile device by its MAC Address.

- [x] ✅ **DELETE** `/JSSResource/mobiledevices/id/{id}`
  - `DeleteMobileDeviceByID` operation deletes a mobile device by its ID.

- [x] ✅ **DELETE** `/JSSResource/mobiledevices/name/{name}`
  - `DeleteMobileDeviceByName` operation deletes a mobile device by its name.

- [] ❌ **DELETE** `/JSSResource/computers/udid/{udid}`
  - `DeleteComputerByUUID` operation deletes a computer by its UUID.

- [] ❌ **DELETE** `/JSSResource/mobiledevices/serialnumber/{serialnumber}`
  - `DeletemobiledevicesBySerialNumber` operation deletes a computer by its Serial Number.

- [] ❌ **DELETE** `/JSSResource/mobiledevices/macaddress/{macaddress}`
  - `DeletemobiledevicesByMacAddress` operation deletes a computer by its Mac Address.

## Summary

- Total Endpoints Covered: 10
  - `/JSSResource/mobiledevices`
  - `/JSSResource/mobiledevices/id/{id}`
  - `/JSSResource/mobiledevices/name/{name}`
  - `/JSSResource/mobiledevices/id/{id}/subset/{subset}`
  - `/JSSResource/mobiledevices/name/{name}/subset/{subset}`

- Total Operations Covered: 10
- Total Operations Not Covered: 11


### Jamf Pro Classic API - Patch Policies

This documentation outlines the operations available for managing Patch Policies within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/patchpolicies/id/{id}`
  - `GetPatchPoliciesByID` operation retrieves the details of a patch policy by its ID.

- [x] ✅ **GET** `/JSSResource/patchpolicies/id/{id}/subset/{subset}`
  - `GetPatchPolicyByIDAndDataSubset` operation fetches a specific subset of data for a patch policy by its ID.

- [x] ✅ **POST** `/JSSResource/patchpolicies/softwaretitleconfig/id/{softwaretitleconfigid}`
  - `CreatePatchPolicy` operation creates a new patch policy.

- [x] ✅ **PUT** `/JSSResource/patchpolicies/id/{id}`
  - `UpdatePatchPolicy` operation updates an existing patch policy by its ID.

- [x] ✅ **DELETE** `/JSSResource/patchpolicies/id/{id}`
  - `DeletePatchPolicyByID` operation deletes a patch policy by its ID.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/patchpolicies/id/{id}`
  - `/JSSResource/patchpolicies/id/{id}/subset/{subset}`
  - `/JSSResource/patchpolicies/softwaretitleconfig/id/{softwaretitleconfigid}`

- Total Operations Covered: 5

### Jamf Pro API - Computer Inventory

This documentation details the operations available for managing Computer Inventory within Jamf Pro using the API, which supports JSON data structures.

## Operations

- [x] ✅ **GET** `/api/v1/computers-inventory`
  - `GetComputersInventory` retrieves a paginated list of all computer inventory information. It supports sorting and section filters.

- [x] ✅ **GET** `/api/v1/computers-inventory/{id}`
  - `GetComputerInventoryByID` fetches a specific computer's inventory information by its ID.

- [x] ✅ **GET** `/api/v1/computers-inventory/filevault`
  - `GetComputersFileVaultInventory` retrieves all computer inventory FileVault information.

- [x] ✅ **GET** `/api/v1/computers-inventory/{id}/filevault`
  - `GetComputerFileVaultInventoryByID` returns FileVault details for a specific computer by its ID.

- [x] ✅ **GET** `/api/v1/computers-inventory/{id}/view-recovery-lock-password`
  - `GetComputerRecoveryLockPasswordByID` retrieves a computer's recovery lock password by the computer ID.

- [x] ✅ **PATCH** `/api/v1/computers-inventory/{id}`
  - `UpdateComputerInventoryByID` updates a specific computer's inventory information by its ID.

- [x] ✅ **DELETE** `/api/v1/computers-inventory/{id}`
  - `DeleteComputerInventoryByID` deletes a computer's inventory information by its ID.

- [x] ✅ **POST** `/api/v1/computers-inventory/{id}/attachments`
  - `UploadAttachmentAndAssignToComputerByID` uploads a file attachment and assigns it to a computer by the computer ID.

- [x] ✅ **DELETE** `/api/v1/computers-inventory/{computerID}/attachments/{attachmentID}`
  - `DeleteAttachmentByIDAndComputerID` deletes a computer's inventory attachment by the computer ID and attachment ID.

## Summary

- Total Endpoints Covered: 6
  - `/api/v1/computers-inventory`
  - `/api/v1/computers-inventory/{id}`
  - `/api/v1/computers-inventory/filevault`
  - `/api/v1/computers-inventory/{id}/filevault`
  - `/api/v1/computers-inventory/{id}/view-recovery-lock-password`
  - `/api/v1/computers-inventory/{id}/attachments`

- Total Operations Covered: 9
- Total Operations Covered: 2

### Jamf Pro Classic API - Removable Mac Addresses

This documentation outlines the operations available for managing Removable Mac Addresses within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/removablemacaddresses`
  - `GetRemovableMACAddresses` operation retrieves a list of all removable MAC addresses.

- [x] ✅ **GET** `/JSSResource/removablemacaddresses/id/{id}`
  - `GetRemovableMACAddressByID` operation retrieves the details of a removable MAC address by its ID.

- [x] ✅ **GET** `/JSSResource/removablemacaddresses/name/{name}`
  - `GetRemovableMACAddressByName` operation retrieves the details of a removable MAC address by its name.

- [x] ✅ **POST** `/JSSResource/removablemacaddresses/id/{id}`
  - `CreateRemovableMACAddress` operation creates a new removable MAC address.

- [x] ✅ **PUT** `/JSSResource/removablemacaddresses/id/{id}`
  - `UpdateRemovableMACAddressByID` operation updates an existing removable MAC address by its ID.

- [x] ✅ **PUT** `/JSSResource/removablemacaddresses/name/{name}`
  - `UpdateRemovableMACAddressByName` operation updates an existing removable MAC address by its name.

- [x] ✅ **DELETE** `/JSSResource/removablemacaddresses/id/{id}`
  - `DeleteRemovableMACAddressByID` operation deletes a removable MAC address by its ID.

- [x] ✅ **DELETE** `/JSSResource/removablemacaddresses/name/{name}`
  - `DeleteRemovableMACAddressByName` operation deletes a removable MAC address by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/removablemacaddresses`
  - `/JSSResource/removablemacaddresses/id/{id}`
  - `/JSSResource/removablemacaddresses/name/{name}`

- Total Operations Covered: 8

### Jamf Pro Classic API - Restricted Software

This documentation outlines the operations available for managing Restricted Software within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/restrictedsoftware`
  - `GetRestrictedSoftwares` operation retrieves a list of all restricted software entries.

- [x] ✅ **GET** `/JSSResource/restrictedsoftware/id/{id}`
  - `GetRestrictedSoftwareByID` operation retrieves the details of a specific restricted software entry by its ID.

- [x] ✅ **GET** `/JSSResource/restrictedsoftware/name/{name}`
  - `GetRestrictedSoftwareByName` operation retrieves the details of a specific restricted software entry by its name.

- [x] ✅ **POST** `/JSSResource/restrictedsoftware/id/{id}`
  - `CreateRestrictedSoftware` operation creates a new restricted software entry.

- [x] ✅ **PUT** `/JSSResource/restrictedsoftware/id/{id}`
  - `UpdateRestrictedSoftwareByID` operation updates an existing restricted software entry by its ID.

- [x] ✅ **PUT** `/JSSResource/restrictedsoftware/name/{name}`
  - `UpdateRestrictedSoftwareByName` operation updates an existing restricted software entry by its name.

- [x] ✅ **DELETE** `/JSSResource/restrictedsoftware/id/{id}`
  - `DeleteRestrictedSoftwareByID` operation deletes a restricted software entry by its ID.

- [x] ✅ **DELETE** `/JSSResource/restrictedsoftware/name/{name}`
  - `DeleteRestrictedSoftwareByName` operation deletes a restricted software entry by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/restrictedsoftware`
  - `/JSSResource/restrictedsoftware/id/{id}`
  - `/JSSResource/restrictedsoftware/name/{name}`

- Total Operations Covered: 8

### Jamf Pro Classic API - Software Update Servers

This documentation outlines the operations available for managing Software Update Servers within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/softwareupdateservers`
  - `GetSoftwareUpdateServers` operation retrieves a list of all software update servers.

- [x] ✅ **GET** `/JSSResource/softwareupdateservers/id/{id}`
  - `GetSoftwareUpdateServersByID` operation retrieves the details of a specific software update server by its ID.

- [x] ✅ **GET** `/JSSResource/softwareupdateservers/name/{name}`
  - `GetSoftwareUpdateServersByName` operation retrieves the details of a specific software update server by its name.

- [x] ✅ **POST** `/JSSResource/softwareupdateservers/id/0`
  - `CreateSoftwareUpdateServer` operation creates a new software update server.

- [x] ✅ **PUT** `/JSSResource/softwareupdateservers/id/{id}`
  - `UpdateSoftwareUpdateServerByID` operation updates an existing software update server by its ID.

- [x] ✅ **PUT** `/JSSResource/softwareupdateservers/name/{name}`
  - `UpdateSoftwareUpdateServerByName` operation updates an existing software update server by its name.

- [x] ✅ **DELETE** `/JSSResource/softwareupdateservers/id/{id}`
  - `DeleteSoftwareUpdateServerByID` operation deletes a software update server by its ID.

- [x] ✅ **DELETE** `/JSSResource/softwareupdateservers/name/{name}`
  - `DeleteSoftwareUpdateServerByName` operation deletes a software update server by its name.

## Summary

- Total Endpoints Covered: 3
  - `/JSSResource/softwareupdateservers`
  - `/JSSResource/softwareupdateservers/id/{id}`
  - `/JSSResource/softwareupdateservers/name/{name}`

- Total Operations Covered: 8

### Jamf Pro Classic API - VPP Accounts

This documentation outlines the operations available for managing VPP (Volume Purchase Program) Accounts within Jamf Pro using the Classic API, which supports XML data structures.

## Operations

- [x] ✅ **GET** `/JSSResource/vppaccounts`
  - `GetVPPAccounts` operation retrieves a list of all VPP accounts.

- [x] ✅ **GET** `/JSSResource/vppaccounts/id/{id}`
  - `GetVPPAccountByID` operation retrieves the details of a specific VPP account by its ID.

- [x] ✅ **POST** `/JSSResource/vppaccounts/id/0`
  - `CreateVPPAccount` operation creates a new VPP account.

- [x] ✅ **PUT** `/JSSResource/vppaccounts/id/{id}`
  - `UpdateVPPAccountByID` operation updates an existing VPP account by its ID.

- [x] ✅ **DELETE** `/JSSResource/vppaccounts/id/{id}`
  - `DeleteVPPAccountByID` operation deletes a VPP account by its ID.

## Summary

- Total Endpoints Covered: 2
  - `/JSSResource/vppaccounts`
  - `/JSSResource/vppaccounts/id/{id}`

- Total Operations Covered: 5

## Jamf Pro Classic API - Webhooks

This documentation outlines the operations available for managing Webhooks within Jamf Pro using the Classic API, which supports XML data structures.

### Operations

- [x] ✅ **GET** `/JSSResource/webhooks`
  - `GetWebhooks` operation retrieves a list of all webhooks.

- [x] ✅ **GET** `/JSSResource/webhooks/id/{id}`
  - `GetWebhookByID` operation retrieves the details of a specific webhook by its ID.

- [x] ✅ **GET** `/JSSResource/webhooks/name/{name}`
  - `GetWebhookByName` operation retrieves the details of a specific webhook by its name.

- [x] ✅ **POST** `/JSSResource/webhooks/id/0`
  - `CreateWebhook` operation creates a new webhook.

- [x] ✅ **PUT** `/JSSResource/webhooks/id/{id}`
  - `UpdateWebhookByID` operation updates an existing webhook by its ID.

- [x] ✅ **PUT** `/JSSResource/webhooks/name/{name}`
  - `UpdateWebhookByName` operation updates an existing webhook by its name.

- [x] ✅ **DELETE** `/JSSResource/webhooks/id/{id}`
  - `DeleteWebhookByID` operation deletes a webhook by its ID.

- [x] ✅ **DELETE** `/JSSResource/webhooks/name/{name}`
  - `DeleteWebhookByName` operation deletes a webhook by its name.

### Summary

- Total Endpoints Covered: 3
  - `/JSSResource/webhooks`
  - `/JSSResource/webhooks/id/{id}`
  - `/JSSResource/webhooks/name/{name}`

- Total Operations Covered: 8

## Jamf Pro Classic API - Computer Checkin

This documentation outlines the operations available for managing Computer Checkin settings within Jamf Pro using the Classic API, which supports XML data structures.

### Operations

- [x] ✅ **GET** `/JSSResource/computercheckin`
  - `GetComputerCheckinInformation` operation retrieves the current computer check-in settings.

- [x] ✅ **PUT** `/JSSResource/computercheckin`
  - `UpdateComputerCheckinInformation` operation updates the computer check-in settings.

### Summary

- Total Endpoints Covered: 1
  - `/JSSResource/computercheckin`

- Total Operations Covered: 2
  - Retrieving current computer check-in settings.
  - Updating computer check-in settings.

## Jamf Pro Classic API - GSX Connection

This documentation outlines the operations available for managing the GSX Connection settings within Jamf Pro using the Classic API, which supports XML data structures.

### Operations

- [x] ✅ **GET** `/JSSResource/gsxconnection`
  - `GetGSXConnectionInformation` operation retrieves the current GSX connection settings.

- [x] ✅ **PUT** `/JSSResource/gsxconnection`
  - `UpdateGSXConnectionInformation` operation updates the GSX connection settings.

### Summary

- Total Endpoints Covered: 1
  - `/JSSResource/gsxconnection`

- Total Operations Covered: 2
  - Retrieving current GSX connection settings.
  - Updating GSX connection settings.

## Jamf Pro Classic API - SMTP Server

This documentation outlines the operations available for managing SMTP Server settings within Jamf Pro using the Classic API, which supports XML data structures.

### Operations

- [x] ✅ **GET** `/JSSResource/smtpserver`
  - `GetSMTPServerInformation` operation retrieves the current SMTP server settings.

- [x] ✅ **PUT** `/JSSResource/smtpserver`
  - `UpdateSMTPServerInformation` operation updates the SMTP server settings.

### Summary

- Total Endpoints Covered: 1
  - `/JSSResource/smtpserver`

- Total Operations Covered: 2
  - Retrieving current SMTP server settings.
  - Updating SMTP server settings.

## Jamf Pro Classic API - VPP Assignments

This documentation outlines the operations available for managing VPP Assignments within Jamf Pro using the Classic API, which supports XML data structures.

### Operations

- [x] ✅ **GET** `/JSSResource/vppassignments`
  - `GetVPPAssignments` operation fetches a list of all VPP assignments.

- [x] ✅ **GET** `/JSSResource/vppassignments/id/{id}`
  - `GetVPPAssignmentByID` operation fetches a specific VPP assignment by its ID.

- [x] ✅ **POST** `/JSSResource/vppassignments/id/0`
  - `CreateVPPAssignment` operation creates a new VPP assignment.

- [x] ✅ **PUT** `/JSSResource/vppassignments/id/{id}`
  - `UpdateVPPAssignmentByID` operation updates an existing VPP assignment by its ID.

- [x] ✅ **DELETE** `/JSSResource/vppassignments/id/{id}`
  - `DeleteVPPAssignmentByID` operation deletes a VPP assignment by its ID.

### Summary

- Total Endpoints Covered: 2
  - `/JSSResource/vppassignments`
  - `/JSSResource/vppassignments/id/{id}`

- Total Operations Covered: 5
  - Fetching all VPP assignments.
  - Fetching a specific VPP assignment by ID.
  - Creating a new VPP assignment.
  - Updating an existing VPP assignment by ID.
  - Deleting a VPP assignment by ID.

## Progress Summary

- Total Operations: 470
- Total Covered Operations: 435
- Not Covered: 35
- Partially Covered: 0
- Deprecated: 


## Notes

- No preview api endpoints will be covered by this sdk. Only generally available endpoints will be covered.

