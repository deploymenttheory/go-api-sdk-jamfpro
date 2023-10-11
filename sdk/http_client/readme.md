## jamf_api_negotiation.go.

The file jamf_api_negotiation.go handles the content type negotiation and response decoding for different types of JAMF APIs.

### Constants and Types:

ContentType: A type alias for a string that represents the content type.
JSON and XML: Constants representing the "application/json" and "application/xml" content types, respectively.
Functions:

determineAPIType: Determines the type of API (Classic or JamfPro) based on the endpoint string. It logs a warning if it's unable to determine the API type and defaults to "Unknown".
DetermineContentType: Determines the content type (JSON or XML) based on the API type and the HTTP method. By default, it returns JSON.
UnmarshalResponse: Decodes the response body based on the JAMF API type and the HTTP method. It can handle both JSON and XML responses depending on the API type and method.
Now, I'll move on to the next file, http_request.go, to understand its contents and how it relates to the issue you mentioned.

## http_request.go
The file http_request.go handles the construction and execution of HTTP requests.

Comments:

constructRequest: This function determines the API type and content type based on the endpoint and method, marshals the request body to the appropriate format (JSON or XML) based on the content type, and sends the HTTP request.
doRequest: This function calls constructRequest to send the HTTP request, uses UnmarshalResponse to decode the response into the desired format based on the API type and method, and handles the decision for retries based on the response status code.
doRequestWithRetry: This function retries the request if a rate limit error is encountered and implements a backoff strategy to avoid overloading the API server.
Functions:

constructRequest: Constructs and sends an HTTP request. It first determines the API type and content type, then marshals the request body based on the content type, and finally constructs and sends the HTTP request.
The snippet provided seems to be truncated, and I haven't yet seen the implementations of doRequest and doRequestWithRetry.

Functions (Continued from constructRequest):
constructRequest (continued): After constructing the HTTP request, the function adds necessary headers, including the authorization token. If the token is invalid, it attempts to refresh it. It then sends the request and logs any errors.

DoRequest: This function sends an HTTP request and handles retries if necessary.

Starts by logging the initiation of the request if in DebugMode.
Checks the HTTP method to decide if it should use retries or send the request directly. Some methods (POST and PUT) are set to send directly without retries.
Calls constructRequest to send the actual request.
Uses UnmarshalResponse to decode the response based on the API type and HTTP method.
If the HTTP method is in the directMethods map or if the response status code indicates success (200-299 range), it returns the response. Otherwise, it seems to indicate that the request might need to be retried, but the logic is truncated.
From the provided portion, I see that the DoRequest method returns the response and any error that occurs during the request. This function seems to be the one called in the ssoFailover.go file, which you mentioned earlier.

## http_methods.go

The http_methods.go file provides convenient methods for making various HTTP requests (GET, POST, PUT, DELETE) using the Client type. Here's a brief overview:

Methods:
Get: Sends a GET request to the specified endpoint and unmarshals the response into the provided out variable. The caller is responsible for closing the response body.
Post: Sends a POST request to the specified endpoint with the provided body and unmarshals the response into the out variable.
Put: Sends a PUT request to the specified endpoint with the provided body and unmarshals the response into the out variable.
Delete: Sends a DELETE request to the specified endpoint and unmarshals the response into the out variable. (The implementation seems to be truncated in the provided snippet.)
Each of these methods logs the request details if the client is in DebugMode. They all use the DoRequest method from the http_request.go file to actually send the request and handle the response.

With this understanding, I'll proceed to the next file, http_logging.go, to further understand the client's behavior and functionalities.

## http_logging.go

The http_logging.go file is focused on logging functionalities. Here's a breakdown:

Logger Interface:

Logger is an interface that defines methods for different logging levels: Trace, Debug, Info, Warn, Error, and Fatal.
Default Logger:

defaultLogger is a struct that implements the Logger interface using Go's standard log package.
Each method (e.g., Trace, Debug, Info, etc.) logs messages with corresponding prefixes like "[TRACE]", "[DEBUG]", etc.
The Fatal method logs the message and then calls log.Fatalln, which will exit the program.
Function:

NewDefaultLogger: This function returns a new instance of the default logger.
This file provides a standardized way to handle logging throughout the client's codebase. The use of an interface (Logger) allows for flexibility in terms of using other logging libraries or custom logging mechanisms if needed.

## http_helpers.go

The http_helpers.go file contains utility functions to assist with HTTP operations. Here's a summary:

Functions:
ParseISO8601Date: Tries to parse a string date in ISO 8601 format using the RFC3339 format. It returns a time.Time object and an error.

EnsureHTTPScheme: Ensures that a given URL string is prefixed with either "http://" or "https://". If neither prefix is present, it defaults to adding "http://".

These helper functions are used throughout the client code to handle specific tasks related to HTTP operations and data manipulation.

## http_error_handling.go

The http_error_handling.go file seems to focus on handling API errors and converting them into structured errors when possible. Here's a summary:

Types:

APIError: Represents a structured API error response with fields for status code and an error message.
Methods & Functions:

handleAPIError: Handles error responses from the API and tries to convert them into a structured error. If the error response from the API is structured (contains an error message and code), it logs and returns a structured APIError. If the error response isn't structured, it tries to decode the error message directly from the response body. If this also fails, it defaults to using the HTTP status message as the error message.

Error: This method on the APIError type returns a string representation of the error, combining the status code and the error message.

translateStatusCode: Provides human-readable messages for various HTTP status codes. It uses a map to match status codes to their corresponding messages.

Given the focus on error handling in this file, it provides context on how the client processes and presents API errors to the user or caller.

## http_client.go

The http_client.go file provides core functionalities and structures related to the HTTP client's behavior.
Methods:

NewClient: Creates a new instance of the Client. If no logger is provided, it defaults to the defaultLogger. It applies any additional options provided to it.
SetAuthMethod: Sets the authentication method for the client. Supported methods are "bearer" and "oauth".
constructAPIEndpoint: Constructs the full URL for a given API endpoint path.
ObtainToken: Fetches and sets an authentication token using the provided username and password. The method constructs the authentication URL, sends a POST request with basic authentication, handles potential deprecation warnings, checks the response status, and then decodes the token from the response. If successful, it sets the token and expiry time on the client.
Error Handling:

In the ObtainToken method, if the response status code isn't http.StatusOK, it logs a warning and invokes c.handleAPIError(resp) to handle the error.

## http_client_auth.go

The http_client_auth.go file focuses on authentication functionalities, particularly around handling OAuth tokens. Here's a summary:

Types:

ClientAuthConfig: Represents the structure to read authentication details from a JSON configuration file. It includes fields for BaseURL, Username, Password, ClientID, and ClientSecret.
OAuthResponse: Represents the response structure when obtaining an OAuth access token. Fields include AccessToken, ExpiresIn, TokenType, RefreshToken, and Error.
OAuthCredentials: Contains the client ID and client secret required for OAuth authentication.
Functions & Methods:

LoadClientAuthConfig: Reads a JSON configuration file and decodes it into a ClientAuthConfig struct. It retrieves authentication details like BaseURL, Username, and Password for the client.
SetOAuthCredentials: Sets the OAuth credentials (Client ID and Client Secret) for the client instance. These credentials are used for obtaining and refreshing OAuth tokens for authentication.
ObtainOAuthToken: Fetches an OAuth access token using the provided OAuthCredentials (Client ID and Client Secret).
The file provides functionalities for loading authentication configurations, setting OAuth credentials, and obtaining OAuth tokens, which are crucial for interacting with protected endpoints.