# Extracting API Calls Using Developer Tools in Edge/Chrome

## Quick Steps

1. Open Developer Tools (F12 or Ctrl+Shift+J on Windows/Linux, Command+Option+J on macOS)
2. Select the Network tab
3. Perform the action that triggers the API call
4. Find and click the API request in the Network Log
5. Check the Headers tab to get the HTTP method
6. Go to the Preview tab to see the formatted request/response data

## Step-by-Step Instructions

### 1. Open Developer Tools

- Log into Jamf Pro
[001](/docs/screen_shots/debug_api_calls/001.png)
- Right-click on the webpage and select "Inspect"
- Or press F12 (or Ctrl+Shift+J on Windows/Linux, Command+Option+J on macOS)

[002](/docs/screen_shots/debug_api_calls/002.png)

### 2. Access the Network Tab

- Click on the "Network" tab in DevTools
[003](/docs/screen_shots/debug_api_calls/003.png)
- Make sure recording is enabled (it should be by default)
- Filter the logs with XHR and Fetch
[004](/docs/screen_shots/debug_api_calls/004.png)
- You may want to clear the current log by clicking the 🗑️ icon

### 3. Trigger the API Call

- Perform the action on the website that will make the API call
- This could be creating a resource, deleting a resource, updating etc
[005](/docs/screen_shots/debug_api_calls/004.png)

### 4. Locate the API Request

- Look for requests in the Network Log that appear to be API calls
- These often have jamf pro resource name in log
[006](/docs/screen_shots/debug_api_calls/004.png)
- XHR and Fetch request types usually indicate API calls

### 5. Extract API Information

- Click on the request to open the details panel
- Go to the **Headers** tab to find:
  - Request URL (the endpoint)
  - Request Method (GET, POST, PUT, DELETE, etc.)
  - Request headers (including authorization)
[007](/docs/screen_shots/debug_api_calls/007.png)

- Go to the **Payload** tab to see:
  - For POST/PUT requests: The formatted request data
  - For all requests: The formatted response data
[008](/docs/screen_shots/debug_api_calls/008.png)

- Go to the **Preview** tab to see:
  - For POST/PUT requests: The formatted request data
  - For all requests: The formatted response data
[009](/docs/screen_shots/debug_api_calls/009.png)

- Go to the **Response** tab to see:
  - For POST/PUT requests: The formatted response data for create and update
  - For all other requests: The formatted response data e.g for GET
[010](/docs/screen_shots/debug_api_calls/010.png)
  
### 6. Additional Tips

- **Filter requests**: Type in the filter box to show only certain requests
  - Type "XHR" or "Fetch" to show only API calls
  - Type a domain or keyword to filter by URL

- **See request payload**: For POST requests, check the "Request Payload" section under Headers or use the Payload tab

- **Copy as cURL/fetch**: Right-click the request and select "Copy" → "Copy as cURL" or "Copy as fetch" to get code you can reuse

- **Simulate slow connections**: Use the Throttling dropdown to test how API calls perform under different network conditions

- **Preserve log**: Check "Preserve log" to keep requests visible when navigating between pages
