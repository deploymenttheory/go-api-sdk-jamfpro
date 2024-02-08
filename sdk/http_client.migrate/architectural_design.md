# Architectural Design Decision: Configuration and Initialization of Jamf Pro SDK

## Decision Topic: Defining a flexible, modular, and secure mechanism for the initialization and configuration of the SDK to interact with various environments of the Jamf Pro API

## Context

The Jamf Pro api SDK will be used in diverse contexts including Terraform providers, standalone scripts, compiled apps, and CI/CD pipelines. It needs to cater to multiple Jamf Pro environments like development, pre-production, and production. This decision outlines the configuration and initialization approach to ensure flexibility, modularity, and security.

## Decision

Configuration Structure:

Implement a Go struct (Config) to encapsulate configuration parameters required by the SDK.
go

```go
type Client struct {
	BaseURL                    string
	authMethod                 string 
	Token                      string
	oAuthCredentials           OAuthCredentials           
	bearerTokenAuthCredentials BearerTokenAuthCredentials
	Expiry                     time.Time
	httpClient                 *http.Client
	tokenLock                  sync.Mutex
	config                     Config
	logger                     Logger
	ConcurrencyMgr             *ConcurrencyManager
}
```

## SDK Initialization

Create an initialization function (NewClient). It accepts an optional Config object, which if not provided, defaults to retrieving configuration from environment variables.
go

```go
func NewClient(cfg *Config) (*JamfClient, error) {
    // Initialization logic here
}
```

### Configuration Loading Utility:

Implement utility functions that will allow loading the configuration from a file, streamlining the process for users who prefer configuration files over environment variables or direct parameters.

### Automatic Token Lifecycle Management:

During the SDK initialization, automate the retrieval of the Bearer Token using the provided authentication details.
Internally manage the token's lifecycle, which includes auto-refreshing when nearing expiry, thus abstracting token management from the SDK users and ensuring uninterrupted API interactions. Include a buffer period that if reached will trigger the token refresh logic the next time a request is performed. Token's should be locked

## Rationale

Flexibility: By supporting environment variables, direct parameters, and configuration files, the SDK can be seamlessly integrated into diverse contexts, from scripts to CI/CD pipelines.

Security: Handling token lifecycle internally reduces the risk of token mismanagement. Direct exposure of client_id and client_secret is minimized, as they can be stored in secure environment variables or encrypted configuration files.

User Experience: By abstracting complexities like token management and environment selection, the SDK offers a streamlined experience to developers, ensuring they focus on business logic rather than SDK intricacies.

## Implications

Maintenance: As Jamf Pro evolves, the SDK may need updates to accommodate any changes in the authentication process or endpoint structures.

Security Practices: Users must ensure they follow best practices in securely storing and managing client_id and client_secret, especially when using environment variables or configuration files

Alright, let's break down the error handling architectural design decision based on your requirements.

---

## Architectural Design Decision: Implementing a Comprehensive Logging System in Jamf Pro SDK

**Decision Topic**: The introduction and implementation of a comprehensive, multi-tiered logging system within the SDK to provide detailed and adjustable logging capabilities.

### Context:
The SDK is used by a diverse audience who have varying needs regarding the amount and type of information required from logging. There needs to be a balance between too much information, which can overwhelm users, and too little, which can hinder problem resolution.

### Decision:

1. **LogLevel Enumeration**:
   - Define a `LogLevel` type as an enumeration to represent various logging levels, allowing users to set the granularity of logs they receive.
   ```go
   type LogLevel int
   
   const (
       LogLevelNone LogLevel = iota
       LogLevelWarning
       LogLevelInfo
       LogLevelDebug
   )
   ```

 **Logging Interface**:
   - Introduce a `Logger` interface with methods corresponding to different logging levels (`Trace`, `Debug`, `Info`, `Warn`, `Error`, `Fatal`). This enables a consistent logging approach across different parts of the SDK.

```go
   type Logger interface {
       // ... method signatures ...
   }
```

3. **Default Logger**:
   - Implement a `defaultLogger` that utilizes Go’s standard logging library, encapsulating the logic for checking log levels before emitting logs. The default logger's level can be set at runtime.

4. **Flexible Log Level Setting**:
   - Provide a method to set the log level (`SetLevel`) dynamically, allowing users to adjust the verbosity of logs as needed without changing the code.

   ```go
   func (d *defaultLogger) SetLevel(level LogLevel) {
       d.logLevel = level
   }
   ```

5. **Level-Based Logging Logic**:
   - Ensure that each logging method in the `defaultLogger` checks the current `logLevel` to decide whether to output the log message, effectively filtering logs based on the set level.

### Rationale:

- **Customizability**: Users can customize the verbosity of the logs by setting the appropriate `LogLevel`, making the SDK adaptable for different environments and use cases.

- **Clarity and Relevance**: By filtering logs according to the set level, users receive only the most relevant information, reducing noise and focusing on the appropriate details for their needs.

- **Simplicity and Familiarity**: Using Go’s standard logging library for the default implementation keeps the SDK simple and familiar to Go developers.

- **Flexibility and Maintenance**: The logging interface allows for different logging implementations to be integrated in the future, ensuring the SDK can evolve without breaking existing functionality.

### Implications:

- **Performance**: The logging system is designed to minimize performance impacts by checking log levels before constructing log messages or performing output operations.

- **Maintenance**: Developers must ensure that log messages at different levels are meaningful and appropriate, necessitating thoughtful logging throughout the SDK's development.

- **Consistency**: A consistent logging interface across the SDK ensures that different components and external contributors adhere to the same logging standards and practices.

---

This document outlines the architectural decision-making for implementing a comprehensive logging system, detailing the context, decisions, rationale, and implications, and aligns with the current implementation of the logging system in the SDK.

## Architectural Design Decision: Dynamic Rate Limiting and Retrying Mechanism in Jamf Pro SDK

**Decision Topic**: Implementing a responsive mechanism that automatically adapts rate limiting and retry strategies based on real-time API response behaviors and headers.

### Context:
Given that APIs can change their rate-limiting behaviors dynamically due to various reasons such as server loads, maintenance, and other external factors, it is crucial for the SDK to adjust its request patterns accordingly to ensure optimal performance and minimize failed requests.

### Decision:

1. **Dynamic Inspection of Rate Limit Headers**:
   - After every API call, parse the response headers to extract information regarding the rate limits. Specifically, focus on headers like `X-RateLimit-Remaining` (indicating how many requests are left in the current window) and `X-RateLimit-Reset` (indicating when the rate limit window resets).
   - Based on the parsed information, adjust the rate at which the SDK sends requests. For instance, if the `X-RateLimit-Remaining` indicates only a few requests are left and the reset time is far off, the SDK should slow down its request rate.

2. **Intelligent Retrying**:
   - If a request fails due to rate limiting (typically indicated by a `429 Too Many Requests` response), the SDK should wait for the time indicated by `X-RateLimit-Reset` before retrying.
   - For other types of errors, use an exponential backoff strategy with a cap to ensure the SDK doesn't retry indefinitely.

3. **User Configurability**:
   - Provide configuration options allowing users to set maximum retry attempts, define custom backoff strategies, or even disable the dynamic rate limiting if needed.

    ```go
    type Config struct {
        // ... other fields ...
        MaxRetryAttempts int
        CustomBackoff    func(attempt int) time.Duration
    }
    ```

### Rationale:

- **Adaptability**: By inspecting the API's response headers in real-time and adjusting the request rate dynamically, the SDK can adapt to varying server behaviors, ensuring optimal performance and minimizing disruptions.
  
- **User Experience**: The built-in intelligence of dynamically adjusting request rates and retry strategies abstracts these complexities from the users, offering a more seamless experience.

- **Flexibility**: Offering configuration options ensures that advanced users can fine-tune the SDK's behavior to best fit their unique scenarios.

### Implications:

- **Complexity**: The dynamic nature of this mechanism might introduce additional complexities in terms of maintenance and debugging.
  
- **Latency**: In cases where the API is frequently rate-limiting or there are consistent transient errors, operations might experience added latency due to the wait times.

---

### Architectural Design Decision: Timeouts and Deadlines in the SDK

**Decision**: The SDK will provide a mechanism for users to configure custom timeout values. However, a default timeout will be enforced to ensure that requests don't hang indefinitely.

**Justification**:
- **Usability**: Providing a default timeout value ensures that, out-of-the-box, users won't face issues with requests hanging indefinitely due to unforeseen network or server-side issues.
  
- **Flexibility**: By allowing timeout values to be configurable, the SDK caters to advanced users who may have specific timeout requirements depending on their use case or environment.
  
- **Robustness**: By explicitly handling timeouts, the SDK becomes more resilient to potential disruptions and can give meaningful error messages to the user.

#### Implementation Details:

1. **Default Timeout**: A reasonable default timeout can be set. Let's say `10 seconds` as an example.

2. **Configurable Timeout**: Users should be able to easily configure this value based on their needs.

Here's a potential implementation using Go:

```go
package apiClient

import (
    "net/http"
    "time"
)

const DefaultTimeout = 10 * time.Second

type Client struct {
    httpClient *http.Client
    // ... other fields
}

type ClientOption func(*Client)

func WithTimeout(timeout time.Duration) ClientOption {
    return func(c *Client) {
        c.httpClient.Timeout = timeout
    }
}

func New(options ...ClientOption) *Client {
    client := &Client{
        httpClient: &http.Client{
            Timeout: DefaultTimeout,
        },
    }

    for _, opt := range options {
        opt(client)
    }

    return client
}
```

#### Usage:

For users who want to use the default timeout:

```go
client := apiClient.New()  // Uses the 10 second default timeout
```

For users who want to set a custom timeout:

```go
client := apiClient.New(apiClient.WithTimeout(30 * time.Second))  // Custom 30 second timeout
```

With this approach, the SDK provides sensible defaults while still offering configurability for varied requirements.
---

### Architectural Design Decision: Bearer Token Management in the SDK

**Decision**: The SDK will handle the expiration and renewal of Bearer Tokens automatically without requiring intervention from the user.

**Justification**:

- **Seamlessness**: Automatic token management ensures uninterrupted SDK operations, thus providing a smoother user experience.
  
- **Reliability**: By internally managing token renewals, the SDK reduces the potential for manual errors and enhances the reliability of any tool or script using it.
  
- **Encapsulation**: Users of the SDK should focus on their core requirements, and not be bogged down with the intricacies of token management. The SDK will abstract these details.

#### Implementation Details:

1. **Token Storage**: The SDK will store the Bearer Token and its expiration date internally.
   
2. **Automatic Token Renewal**: Before any API call, the SDK will check the token's expiration date. If the token is close to expiry (or expired), the SDK will proactively renew it using the `/v1/auth/keep-alive` endpoint or by obtaining a new one via `/v1/auth/token`.
   
3. **Transparent to User**: The token renewal process will be transparent to the user. They will simply receive the results of their intended API call without any indication of the token being renewed (unless they're in debug mode, where such internal operations might be logged).

Here's a rough outline in Go:

```go
package apiClient

import (
    "sync"
    "time"
)

type Client struct {
    token         string
    tokenExpiry   time.Time
    httpClient    *http.Client
    tokenLock     sync.Mutex
    // ... other fields
}

// This function is called before every API call to ensure the token is valid.
func (c *Client) ensureValidToken() error {
    c.tokenLock.Lock()
    defer c.tokenLock.Unlock()
    
    // If token is close to expiry or already expired, refresh it.
    // The "5 minutes" buffer is just an example; it can be adjusted as needed.
    if time.Until(c.tokenExpiry) < 5*time.Minute {
        err := c.refreshToken()
        if err != nil {
            return err
        }
    }
    return nil
}

// refreshToken reaches out to Jamf Pro API to get a new token.
func (c *Client) refreshToken() error {
    // Logic to send a POST request to /v1/auth/keep-alive or /v1/auth/token.
    // Update c.token and c.tokenExpiry based on the response.
    // ...
    return nil
}
```
---


# Architectural Decision Record (ADR): Content Negotiation

Jamf Pro offers two key APIs: the Classic API and the Jamf Pro API. Each API has its own set of nuances in terms of base URL, authentication mechanisms, privileges, HTTP methods, data schema, and response codes. In order to interact with both these APIs seamlessly, we require an HTTP client that can modularly address these nuances.

#### **Decision**:

1. **Modular Design**: The Go-based HTTP client will be modular, with separate modules dedicated to handling specifics of the Classic API and the Jamf Pro API.

2. **Unified Interface**: Despite the modularity, there will be a unified interface to interact with both APIs to provide a seamless experience for users of the client.

3. **Authentication**:
   - **Classic API**: Implement both Client Credentials and Bearer Token authentication methods.
   - **Jamf Pro API**: Implement Client Credentials and Bearer Token authentication with provisions to refresh tokens upon expiry.

4. **Data Formats**:
   - **Classic API**: Support both XML and JSON formats for GET requests, and exclusively XML for POST and PUT requests.
   - **Jamf Pro API**: Primarily interact using JSON, with exceptions handled for specific workflows such as file uploads.

5. **Error Handling**: Implement robust error handling to interpret and handle the various HTTP status codes returned by the APIs, and provide descriptive error messages to the user.

6. **Extensibility**: The design will ensure easy extensibility to cater to any future changes or additional features in the Jamf Pro APIs.

#### **Consequences**:

1. The modular design will make it easier to maintain and update the client for individual API changes without affecting other parts of the system.
2. Users will benefit from a unified interface, simplifying the integration process with the Jamf Pro system.
3. Robust error handling will ensure that the client gracefully handles failures, providing clear feedback to the users.
4. The extensible nature of the design will future-proof the client against potential updates in the Jamf Pro system.

---

# Handling of Concurrent Requests in the Go-Based HTTP Client for Jamf Pro APIs

## Context:
The SDK, which is intended to be used by a Terraform provider, needs to handle potential concurrent operations gracefully, ensuring data integrity and avoiding potential issues like race conditions. While Terraform itself manages much of this concurrency, the HTTP client should be designed to safely and effectively handle concurrent requests.

## Decision:
Rate Limiting: Introduce a rate limiter to control the rate of requests sent to the Jamf Pro APIs. This ensures that we don't overwhelm the API with too many requests at once, respecting any API rate limits, and prevents potential throttling.

## Concurrency Control:

Mutexes: Use mutexes (from Go's sync package) to lock critical sections of the code, especially if you have shared state or resources that shouldn't be accessed concurrently. This can prevent race conditions.

Error Handling: Ensure that the client can handle API errors related to concurrency gracefully. For instance, if two concurrent operations result in a conflict (e.g., trying to create a resource that already exists), the client should be able to recognize this and respond appropriately.

Connection Pooling: Utilize connection pooling for the HTTP client. This ensures that multiple concurrent requests can reuse existing connections, rather than opening a new connection for every request, which is less efficient.

State Management: If the client maintains any state, ensure it's designed to be thread-safe. This often involves a combination of mutexes and careful design to ensure that concurrent operations don't produce unpredictable results.

Idempotency: Whenever possible, design operations to be idempotent. This means that even if an operation is executed multiple times (e.g., due to retries), the result remains consistent. This is especially crucial for Terraform providers.

## Consequences:
Proper handling of concurrency will ensure data consistency and prevent potential race conditions, making the Terraform provider robust and reliable.

Introducing rate limiting and connection pooling will optimize performance without compromising on the integrity of operations.

Designing for idempotency will provide more reliable outcomes, especially in the face of intermittent network or service issues.

This decision can then guide the development process, ensuring that concurrency is handled in an effective and safe manner within the HTTP client used by the Terraform provider.

# Intelligent Retrying & Rate Handling in the Go-Based HTTP Client for APIs

## Context:

The HTTP client is being designed with the primary goal of interacting with Jamf Pro APIs, but with the flexibility to accommodate other APIs in the future. Given the unique nature of Jamf Pro's rate limiting recommendations and the absence of built-in rate limit headers, there's a need for an intelligent mechanism to manage request rates and handle retries.

### Decision:

Exponential Backoff with Jitter:
Implement an exponential backoff strategy for retries. This means that for each consecutive retry, the wait time before the next retry will double. Adding "jitter" (a random variation) to the backoff will prevent many clients from starting their retries simultaneously, mitigating the risk of overwhelming the server.

### Response Time Monitoring:

Monitor the average response time of successful requests. If the observed response time is significantly higher than the average, it may indicate server stress. The client should dynamically adjust its behavior, increasing delay between subsequent requests or pausing for a period.

### Generic Rate Limit Header Handling:

For APIs that provide built-in rate limiting (through headers), the client should be able to parse these headers and adjust request rates accordingly.

### Maximum Retries:

Introduce a maximum retry count to ensure that the client doesn't end up in an infinite retry loop in cases of persistent failure.

### Configurability:

Make the client's behavior configurable, allowing users to set values for parameters like base retry delay, maximum retries, and even provide their own backoff strategy function if desired.

### Concurrency Management:

Introduce a mechanism to ensure that the number of concurrent requests does not exceed the recommended limit (e.g., 5 for Jamf Pro). This can be achieved using a semaphore or a similar construct in Go.

### Error Classification:

Implement logic to classify errors. Only transient errors, which indicate temporary server-side issues, should trigger retries. Client-side errors, which indicate issues like bad requests, should not be retried.

## Consequences:

By implementing an intelligent retry and rate-handling mechanism, the HTTP client will be robust, able to gracefully handle server-side issues, and respectful of server resources. This ensures optimal performance, minimizes the risk of overwhelming the server, and provides a consistent user experience.

The flexibility and configurability of the client make it versatile enough to be adapted for other APIs in the future, ensuring longevity and reducing the need for significant rewrites.