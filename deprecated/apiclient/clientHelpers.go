// clientHelpers.go
// For utility/helper functions to support from the main package
package apiClient

import (
	"fmt"
	"net/http"
)

// Print request headers for troubleshooting
func PrintRequestHeaders(req *http.Request) {
	fmt.Println("Request Headers:")
	for name, values := range req.Header {
		// Each value is a slice of strings since headers can be repeated.
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}
}
