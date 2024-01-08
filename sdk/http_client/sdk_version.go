// sdk_version.go
package http_client

import "fmt"

const (
	SDKVersion    = "1.0.0"
	UserAgentBase = "go-jamfpro-api-sdk"
)

func GetUserAgentHeader() string {
	return fmt.Sprintf("%s/%s", UserAgentBase, SDKVersion)
}
