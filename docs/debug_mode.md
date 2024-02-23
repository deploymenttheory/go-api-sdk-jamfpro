# Using Debug Mode with the SDK

The SDK provides a `DebugMode` feature which, when enabled, outputs detailed logs about the HTTP requests, responses, and internal operations. This can be particularly useful during development, troubleshooting, or when you want to understand the internals of the SDK's operations.

## Enabling Debug Mode

To enable `DebugMode`, you'll need to set the `DebugMode` property of the SDK configuration to `true`.

### Example

```go
// Configuration for the httpclient
config := httpclient.Config{
    DebugMode: true, // This enables the Debug Mode
    Logger:    httpclient.NewDefaultLogger(),
}

// Create a new jamfpro client instance using the loaded BaseURL
client := jamfpro.NewClient(baseURL, config)
```

## Understanding Debug Logs

When `DebugMode` is enabled, the SDK will print detailed logs, which include:

- Initialization logs
- HTTP request methods, endpoints, headers, and body content (if applicable)
- HTTP response status, headers, and body content
- Token validation and renewal operations
- Any errors or issues encountered during the SDK's operations

## Precautions

While `DebugMode` is a powerful tool, there are some precautions to keep in mind:

1. **Sensitive Information**: Logs might contain sensitive information, especially if you're working with authentication tokens, secrets, or any private data. Always be cautious about where and how you store or display these logs.
2. **Performance**: Generating and printing detailed logs might have a minor impact on performance. It's recommended to disable `DebugMode` in production environments.
3. **Log Volume**: The SDK can produce a significant volume of logs, especially if it's used for many operations in a short time. Ensure that your logging system can handle this volume if you plan to keep `DebugMode` enabled for extended periods.

## Disabling Debug Mode

If you wish to disable the `DebugMode`, simply set the `DebugMode` property in the configuration to `false`:

\```go
config := httpclient.Config{
    DebugMode: false,
}
\```

---

**Note**: Always use the Debug Mode responsibly, and remember to disable it in production environments to ensure the security and performance of your applications.
