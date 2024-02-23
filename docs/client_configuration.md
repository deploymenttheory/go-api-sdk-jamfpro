# Configuring the `httpclient`

The `httpclient` package offers a customizable HTTP client designed for interfacing with specific APIs, featuring a variety of configuration options.

## 1. `Client` Structure

At the heart of the `httpclient` package is the `Client` structure, which represents the HTTP client designed to interact with a particular API.

### Key Attributes:

- **BaseURL**: The base URL for the API you're connecting to.
- **APIType**: Specifies the type of API ('Classic' or 'JamfPro').
- **authMethod**: Defines the authentication method; either "bearer" or "oauth".
- **Token**: Holds the authentication token.
- **oAuthCredentials**: Holds OAuth credentials when using OAuth for authentication.
- **Expiry**: Token expiry time.
- **httpClient**: Embedded standard HTTP client.
- **tokenLock**: Mutex for handling token-related concurrency.
- **config**: Client's configuration options.
- **logger**: Logger instance for logging activities.
- **ConcurrencyMgr**: Manages concurrency for requests.

## 2. `Config` Structure

The `Config` structure holds various configuration options for the HTTP Client.

### Attributes:

- **DebugMode**: Enables or disables debug mode.
- **MaxRetryAttempts**: Maximum number of retry attempts for failed requests.
- **CustomBackoff**: Custom backoff strategy function.
- **EnableDynamicRateLimiting**: Enables or disables dynamic rate limiting.
- **Logger**: Logger instance for the client.
- **MaxConcurrentRequests**: Maximum concurrent requests allowed.
- **TokenLifespan**: Lifespan of the token before it needs refreshing.
- **BufferPeriod**: Buffer time before the token expires.

## 3. Initializing a New Client

You can initialize a new `Client` instance using the `NewClient` function:

```go
client := httpclient.NewClient(baseURL, config, logger)
```

### Parameters:

- **baseURL**: The base URL for the API.
- **config**: A `Config` instance holding various client configuration options.
- **logger**: Logger instance for the client. If none provided, a default logger is used.

### Notes:

- Default values for `TokenLifespan` and `BufferPeriod` are 30 minutes and 5 minutes, respectively, if not explicitly set in the config.
- Additional client options can be provided during initialization.

## 4. Setting Authentication Method

The `SetAuthMethod` function allows you to specify the client's authentication method:

```go
client.SetAuthMethod("oauth")
```

Possible values for the method are:
- **"bearer"**: Bearer Token Authentication.
- **"oauth"**: OAuth based Authentication.

## 5. Customizing the Client

You can further customize the client using `ClientOption` functions during initialization.

---

Certainly! Let's enhance the documentation with an example.

---

## Example: Setting Up the `httpclient`

Below is an example demonstrating how to configure and initialize the `httpclient`.

### 1. Define Configuration:

First, you need to define the configuration for the client. Here's how you can create a `Config` instance:

```go
config := httpclient.Config{
    DebugMode:                 true,  // Enable debug mode
    MaxRetryAttempts:          3,     // Set maximum retry attempts
    EnableDynamicRateLimiting: true,  // Enable dynamic rate limiting
    Logger:                    nil,   // Use default logger
    MaxConcurrentRequests:     5,     // Set max concurrent requests
    TokenLifespan:             45 * time.Minute,   // Set token lifespan
    BufferPeriod:              10 * time.Minute,   // Set buffer period before token expiry
}
```

### 2. Initialize the Client:

Use the `NewClient` function to initialize a new client with the defined configuration:

```go
baseURL := "your-jamf-instance.jamfcloud.com"
client := httpclient.NewClient(baseURL, config, nil)
```

In this example, we're using the baseURL "https://exampleapi.com". We're also using the default logger by passing `nil`.

### 3. Set Authentication Method:

Now, set the desired authentication method for the client:

```go
client.SetAuthMethod("oauth")  // Use OAuth for authentication
```

You can switch between "bearer" and "oauth" based on your requirements.

### 4. Using the Client:

Now that the client is configured and initialized, you can use it to make requests to the API.

---

By following the example above, you can easily configure and set up the `httpclient` for your specific needs.