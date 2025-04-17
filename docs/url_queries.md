# Querying the Jamf Pro API with RSQL

Jamf Pro's **Jamf Pro API (v1+)** supports querying resources using **RSQL filters** via HTTP requests. This guide explains how to construct API queries using URLs with RSQL in Go.


## 📘 Basic URL Format

Jamf Pro API endpoints follow this pattern:

```
https://<jamf-host>/api/<resource-path>?page=0&page-size=50&sort=fieldName:asc&filter=<rsql-expression>
```

- `page`: zero-based page index.
- `page-size`: number of results per page.
- `sort`: optional; `fieldName:asc` or `fieldName:desc`.
- `filter`: your RSQL query string (URL-encoded).

---

## 🔍 Example: Query Computers by Name

**Goal:** Find all computers whose name contains `test`.

**Endpoint:**

```
GET /api/v1/computers-inventory
```

**RSQL Filter:**

```
general.name==*test*
```

**Full URL (URL-encoded):**

```
https://your-jamf-url.com/api/v1/computers-inventory?filter=general.name%3D%3D%2Atest%2A
```



## 🧑‍💻 Creating an RSQL Query using `url.Values{}` in Go

```go
import (
	"net/http"
	"net/url"
)

func buildJamfQuery() (*http.Request, error) {
	baseURL := "https://your-jamf-url.com/api/v1/computers-inventory"

	params := url.Values{}
	params.Set("page", "0")
	params.Set("page-size", "50")
	params.Set("sort", "general.name:asc")
	params.Set("filter", `general.name==*test*`)

	fullURL := baseURL + "?" + params.Encode()

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer YOUR_API_TOKEN")
	return req, nil
}
```

- `url.Values{}.Encode()` handles correct URL encoding for you.
- The RSQL filter will be encoded automatically (e.g. `==` → `%3D%3D`, `*` → `%2A`).

---
