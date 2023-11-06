// sdk_version.go
package http_client

import "fmt"

const (
	SDKVersion    = "0.0.51"
	UserAgentBase = "go-jamfpro-api-sdk"
)

func GetUserAgent() string {
	return fmt.Sprintf("%s/%s", UserAgentBase, SDKVersion)
}
