// jamfproapi_patch_policies.go
// Jamf Pro Api - Patch Policies On Dashboard
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-policies
// Jamf Pro Api requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

const uriPatchPoliciesJamfProAPI = "/api/v2/patch-policies"

// List

// Struct for paginated response for patch policies
type ResponsePatchPoliciesList struct {
	Size    int                   `json:"totalCount"`
	Results []ResourcePatchPolicy `json:"results"`
}

// Response

// Response struct for creating a patch policy
type ResponsePatchPolicyCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// Resource

// Resource struct representing a Patch Policy object from Pro API
type ResourcePatchPolicy struct {
	ID                           string `json:"id"`
	PolicyName                   string `json:"policyName"`
	PolicyEnabled                bool   `json:"policyEnabled"`
	PolicyTargetVersion          string `json:"policyTargetVersion"`
	PolicyDeploymentMethod       string `json:"policyDeploymentMethod"`
	SoftwareTitle                string `json:"softwareTitle"`
	SoftwareTitleConfigurationId string `json:"softwareTitleConfigurationId"`
	Pending                      int    `json:"pending"`
	Completed                    int    `json:"completed"`
	Deferred                     int    `json:"deferred"`
	Failed                       int    `json:"failed"`
}

// Gets full list of patch policies & handles pagination
func (c *Client) GetPatchPolicies(sortFilter string) (*ResponsePatchPoliciesList, error) {
	resp, err := c.DoPaginatedGet(
		uriPatchPoliciesJamfProAPI+"/policy-details",
		standardPageSize,
		startingPageNumber,
		sortFilter,
	)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "patch policies", err)
	}

	var out ResponsePatchPoliciesList
	out.Size = resp.Size

	for _, value := range resp.Results {
		var newObj ResourcePatchPolicy
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "patch policy", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}
