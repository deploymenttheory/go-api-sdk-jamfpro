// sdk_version.go
package http_client

import "fmt"

const (
	SDKVersion    = "0.0.54"
	UserAgentBase = "go-jamfpro-api-sdk"
)

func GetUserAgentHeader() string {
	return fmt.Sprintf("%s/%s", UserAgentBase, SDKVersion)
}
