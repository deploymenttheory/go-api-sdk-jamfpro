// jamfproapi_enrollment.go
// Jamf Pro Api - Enrollment
// api reference: https://developer.jamf.com/jamf-pro/reference/put_v4-enrollment
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

const uriEnrollmentV3 = "/api/v3/enrollment"
const uriEnrollmentV4 = "/api/v4/enrollment"

// List

// ResponseAccountDrivenUserEnrollmentAccessGroups represents the structure of the response for a list of access groups

type ResponseAccountDrivenUserEnrollmentAccessGroupsList struct {
	TotalCount int                                              `json:"totalCount"`
	Results    []ResourceAccountDrivenUserEnrollmentAccessGroup `json:"results"`
}

type ResponseAccountDrivenUserEnrollmentAccessGroupCreateAndUpdate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// Request

// RequestDeleteMultipleLanguages represents the request body for deleting multiple languages
type RequestDeleteMultipleLanguages struct {
	IDs []string `json:"ids"`
}

// Resource

type ResourceAccountDrivenUserEnrollmentAccessGroup struct {
	ID                                 string `json:"id"`
	GroupID                            string `json:"groupId"`
	LdapServerID                       string `json:"ldapServerId"`
	Name                               string `json:"name"`
	SiteID                             string `json:"siteId"`
	EnterpriseEnrollmentEnabled        bool   `json:"enterpriseEnrollmentEnabled"`
	PersonalEnrollmentEnabled          bool   `json:"personalEnrollmentEnabled"`
	AccountDrivenUserEnrollmentEnabled bool   `json:"accountDrivenUserEnrollmentEnabled"`
	RequireEula                        bool   `json:"requireEula"`
}

type ResourceEnrollment struct {
	InstallSingleProfile                         bool                           `json:"installSingleProfile"`
	SigningMdmProfileEnabled                     bool                           `json:"signingMdmProfileEnabled"`
	MdmSigningCertificate                        *ResourceEnrollmentCertificate `json:"mdmSigningCertificate"`
	RestrictReenrollment                         bool                           `json:"restrictReenrollment"`
	FlushLocationInformation                     bool                           `json:"flushLocationInformation"`
	FlushLocationHistoryInformation              bool                           `json:"flushLocationHistoryInformation"`
	FlushPolicyHistory                           bool                           `json:"flushPolicyHistory"`
	FlushExtensionAttributes                     bool                           `json:"flushExtensionAttributes"`
	FlushSoftwareUpdatePlans                     bool                           `json:"flushSoftwareUpdatePlans"`
	FlushMdmCommandsOnReenroll                   string                         `json:"flushMdmCommandsOnReenroll"`
	MacOsEnterpriseEnrollmentEnabled             bool                           `json:"macOsEnterpriseEnrollmentEnabled"`
	ManagementUsername                           string                         `json:"managementUsername"`
	CreateManagementAccount                      bool                           `json:"createManagementAccount"`
	HideManagementAccount                        bool                           `json:"hideManagementAccount"`
	AllowSshOnlyManagementAccount                bool                           `json:"allowSshOnlyManagementAccount"`
	EnsureSshRunning                             bool                           `json:"ensureSshRunning"`
	LaunchSelfService                            bool                           `json:"launchSelfService"`
	SignQuickAdd                                 bool                           `json:"signQuickAdd"`
	DeveloperCertificateIdentity                 *ResourceEnrollmentCertificate `json:"developerCertificateIdentity"`
	DeveloperCertificateIdentityDetails          ResourceCertificateDetails     `json:"developerCertificateIdentityDetails"`
	MdmSigningCertificateDetails                 ResourceCertificateDetails     `json:"mdmSigningCertificateDetails"`
	IosEnterpriseEnrollmentEnabled               bool                           `json:"iosEnterpriseEnrollmentEnabled"`
	IosPersonalEnrollmentEnabled                 bool                           `json:"iosPersonalEnrollmentEnabled"`
	PersonalDeviceEnrollmentType                 string                         `json:"personalDeviceEnrollmentType"`
	AccountDrivenUserEnrollmentEnabled           bool                           `json:"accountDrivenUserEnrollmentEnabled"`
	AccountDrivenDeviceIosEnrollmentEnabled      bool                           `json:"accountDrivenDeviceIosEnrollmentEnabled"`
	AccountDrivenDeviceMacosEnrollmentEnabled    bool                           `json:"accountDrivenDeviceMacosEnrollmentEnabled"`
	AccountDrivenUserVisionosEnrollmentEnabled   bool                           `json:"accountDrivenUserVisionosEnrollmentEnabled"`
	AccountDrivenDeviceVisionosEnrollmentEnabled bool                           `json:"accountDrivenDeviceVisionosEnrollmentEnabled"`
}

// Resource struct for enrollment language messaging
type ResourceEnrollmentLanguage struct {
	LanguageCode                     string `json:"languageCode"`
	Name                             string `json:"name"`
	Title                            string `json:"title"`
	LoginDescription                 string `json:"loginDescription"`
	Username                         string `json:"username"`
	Password                         string `json:"password"`
	LoginButton                      string `json:"loginButton"`
	DeviceClassDescription           string `json:"deviceClassDescription"`
	DeviceClassPersonal              string `json:"deviceClassPersonal"`
	DeviceClassPersonalDescription   string `json:"deviceClassPersonalDescription"`
	DeviceClassEnterprise            string `json:"deviceClassEnterprise"`
	DeviceClassEnterpriseDescription string `json:"deviceClassEnterpriseDescription"`
	DeviceClassButton                string `json:"deviceClassButton"`
	PersonalEula                     string `json:"personalEula"`
	EnterpriseEula                   string `json:"enterpriseEula"`
	EulaButton                       string `json:"eulaButton"`
	SiteDescription                  string `json:"siteDescription"`
	CertificateText                  string `json:"certificateText"`
	CertificateButton                string `json:"certificateButton"`
	CertificateProfileName           string `json:"certificateProfileName"`
	CertificateProfileDescription    string `json:"certificateProfileDescription"`
	PersonalText                     string `json:"personalText"`
	PersonalButton                   string `json:"personalButton"`
	PersonalProfileName              string `json:"personalProfileName"`
	PersonalProfileDescription       string `json:"personalProfileDescription"`
	UserEnrollmentText               string `json:"userEnrollmentText"`
	UserEnrollmentButton             string `json:"userEnrollmentButton"`
	UserEnrollmentProfileName        string `json:"userEnrollmentProfileName"`
	UserEnrollmentProfileDescription string `json:"userEnrollmentProfileDescription"`
	EnterpriseText                   string `json:"enterpriseText"`
	EnterpriseButton                 string `json:"enterpriseButton"`
	EnterpriseProfileName            string `json:"enterpriseProfileName"`
	EnterpriseProfileDescription     string `json:"enterpriseProfileDescription"`
	EnterprisePending                string `json:"enterprisePending"`
	QuickAddText                     string `json:"quickAddText"`
	QuickAddButton                   string `json:"quickAddButton"`
	QuickAddName                     string `json:"quickAddName"`
	QuickAddPending                  string `json:"quickAddPending"`
	CompleteMessage                  string `json:"completeMessage"`
	FailedMessage                    string `json:"failedMessage"`
	TryAgainButton                   string `json:"tryAgainButton"`
	CheckNowButton                   string `json:"checkNowButton"`
	CheckEnrollmentMessage           string `json:"checkEnrollmentMessage"`
	LogoutButton                     string `json:"logoutButton"`
}

type ResourceEnrollmentCertificate struct {
	Filename         string `json:"filename"`
	KeystorePassword string `json:"keystorePassword,omitempty"`
	IdentityKeystore string `json:"identityKeystore,omitempty"`
	Md5Sum           string `json:"md5Sum,omitempty"`
}

type ResourceCertificateDetails struct {
	Subject      string `json:"subject"`
	SerialNumber string `json:"serialNumber"`
}

// Resource struct for language codes
type ResourceLanguageCode struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

// CRUD

// GetAccountDrivenUserEnrollmentAccessGroups fetches all ADUE access groups
func (c *Client) GetAccountDrivenUserEnrollmentAccessGroups(sort_filter string) (*ResponseAccountDrivenUserEnrollmentAccessGroupsList, error) {
	endpoint := fmt.Sprintf("%s/access-groups", uriEnrollmentV3)
	resp, err := c.DoPaginatedGet(endpoint, standardPageSize, 0, sort_filter)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "ADUE Access Group List", err)
	}

	var OutStruct ResponseAccountDrivenUserEnrollmentAccessGroupsList
	OutStruct.TotalCount = resp.Size
	for _, value := range resp.Results {
		var newObj ResourceAccountDrivenUserEnrollmentAccessGroup
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "ADUE Access Group List", err)
		}
		OutStruct.Results = append(OutStruct.Results, newObj)
	}

	return &OutStruct, nil
}

// GetAccountDrivenUserEnrollmentAccessGroupByID retrieves an Account Driven User Enrollment Access Group by its ID
func (c *Client) GetAccountDrivenUserEnrollmentAccessGroupByID(id string) (*ResourceAccountDrivenUserEnrollmentAccessGroup, error) {
	endpoint := fmt.Sprintf("%s/access-groups/%s", uriEnrollmentV3, id)

	var ADUEGroup ResourceAccountDrivenUserEnrollmentAccessGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ADUEGroup)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "ADUE Access Group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ADUEGroup, nil
}

// GetAccountDrivenUserEnrollmentAccessGroupByName retrieves an Account Driven User Enrollment Access Group by its name
func (c *Client) GetAccountDrivenUserEnrollmentAccessGroupByName(name string) (*ResourceAccountDrivenUserEnrollmentAccessGroup, error) {
	accessGroupsList, err := c.GetAccountDrivenUserEnrollmentAccessGroups("")
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "ADUE access group", err)
	}

	for _, group := range accessGroupsList.Results {
		if group.Name == name {
			return &group, nil
		}
	}

	return nil, fmt.Errorf(errMsgFailedGetByName, "ADUE access group", name, errMsgNoName)
}

// CreateAccountDrivenUserEnrollmentAccessGroup creates a new ADUE access group
func (c *Client) CreateAccountDrivenUserEnrollmentAccessGroup(script *ResourceAccountDrivenUserEnrollmentAccessGroup) (*ResponseAccountDrivenUserEnrollmentAccessGroupCreateAndUpdate, error) {
	endpoint := uriEnrollmentV3
	var out ResponseAccountDrivenUserEnrollmentAccessGroupCreateAndUpdate

	resp, err := c.HTTP.DoRequest("POST", endpoint, script, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "ADUE access group", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateAccountDrivenUserEnrollmentAccessGroupByID updates an ADUE access group by resource ID
func (c *Client) UpdateAccountDrivenUserEnrollmentAccessGroupByID(id string, groupUpdate *ResourceAccountDrivenUserEnrollmentAccessGroup) (*ResourceAccountDrivenUserEnrollmentAccessGroup, error) {
	endpoint := fmt.Sprintf("%s/access-groups/%s", uriEnrollmentV3, id)
	var out ResourceAccountDrivenUserEnrollmentAccessGroup

	resp, err := c.HTTP.DoRequest("PUT", endpoint, groupUpdate, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "ADUE Access Group", id, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateAccountDrivenUserEnrollmentAccessGroupByName updates an ADUE access group by resource name
func (c *Client) UpdateAccountDrivenUserEnrollmentAccessGroupByName(targetName string, groupUpdate *ResourceAccountDrivenUserEnrollmentAccessGroup) (*ResourceAccountDrivenUserEnrollmentAccessGroup, error) {
	target, err := c.GetAccountDrivenUserEnrollmentAccessGroupByName(targetName)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "ADUE access group", targetName, err)
	}

	resp, err := c.UpdateAccountDrivenUserEnrollmentAccessGroupByID(target.ID, groupUpdate)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "ADUE access group", targetName, err)
	}

	return resp, nil
}

// DeleteAccountDrivenUserEnrollmentAccessGroupByID deletes an ADUE access group with given id
func (c *Client) DeleteAccountDrivenUserEnrollmentAccessGroupByID(id string) error {
	endpoint := fmt.Sprintf("%s/access-groups/%s", uriEnrollmentV3, id)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)

	if err != nil || resp.StatusCode != 204 {
		return fmt.Errorf(errMsgFailedDeleteByID, "ADUE access group", id, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteAccountDrivenUserEnrollmentAccessGroupByName deletes an ADUE access group with given name, leverages GetAccountDrivenUserEnrollmentAccessGroupByName
func (c *Client) DeleteAccountDrivenUserEnrollmentAccessGroupByName(targetName string) error {
	target, err := c.GetAccountDrivenUserEnrollmentAccessGroupByName(targetName)
	if err != nil {
		return fmt.Errorf(errMsgFailedGetByName, "ADUE access group", targetName, err)
	}

	err = c.DeleteAccountDrivenUserEnrollmentAccessGroupByID(target.ID)

	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "ADUE access group", targetName, err)
	}

	return nil
}

// GetEnrollment retrieves the current enrollment configuration
func (c *Client) GetEnrollment() (*ResourceEnrollment, error) {
	endpoint := uriEnrollmentV4
	var enrollment ResourceEnrollment

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &enrollment)
	if err != nil {
		return nil, fmt.Errorf("failed to get enrollment configuration: %v", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &enrollment, nil
}

// UpdateEnrollment updates the enrollment configuration
func (c *Client) UpdateEnrollment(enrollmentUpdate *ResourceEnrollment) (*ResourceEnrollment, error) {
	endpoint := uriEnrollmentV4
	var out ResourceEnrollment

	resp, err := c.HTTP.DoRequest("PUT", endpoint, enrollmentUpdate, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to update enrollment configuration: %v", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetEnrollmentMessageByLanguageID retrieves the enrollment language messaging for a specific language ID
func (c *Client) GetEnrollmentMessageByLanguageID(languageId string) (*ResourceEnrollmentLanguage, error) {
	endpoint := fmt.Sprintf("%s/languages/%s", uriEnrollmentV3, languageId)
	var languageMsg ResourceEnrollmentLanguage

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &languageMsg)
	if err != nil {
		return nil, fmt.Errorf("failed to get enrollment language messaging for language ID '%s': %v", languageId, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &languageMsg, nil
}

// UpdateEnrollmentMessageByLanguageID updates the enrollment messaging for a specific language ID
func (c *Client) UpdateEnrollmentMessageByLanguageID(languageId string, enrollmentMessage *ResourceEnrollmentLanguage) (*ResourceEnrollmentLanguage, error) {
	endpoint := fmt.Sprintf("%s/languages/%s", uriEnrollmentV3, languageId)
	var updatedMessage ResourceEnrollmentLanguage

	resp, err := c.HTTP.DoRequest("PUT", endpoint, enrollmentMessage, &updatedMessage)
	if err != nil {
		return nil, fmt.Errorf("failed to update enrollment language messaging for language ID '%s': %v", languageId, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &updatedMessage, nil
}

// DeleteEnrollmentMessageByLanguageID deletes the enrollment messaging for a specific language ID
func (c *Client) DeleteEnrollmentMessageByLanguageID(languageId string) error {
	endpoint := fmt.Sprintf("%s/languages/%s", uriEnrollmentV3, languageId)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "enrollment language message", languageId, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	var respErr ResponseError
	if err := json.NewDecoder(resp.Body).Decode(&respErr); err == nil {
		if respErr.HTTPStatus != 0 {
			return fmt.Errorf("deletion failed for enrollment language '%s' with status %d: %v", languageId, respErr.HTTPStatus, respErr.Errors)
		}
	}

	return nil
}

// DeleteMultipleEnrollmentMessagesByLanguageIDs deletes multiple enrollment language messages by their IDs
func (c *Client) DeleteMultipleEnrollmentMessagesByLanguageIDs(languageIds []string) error {
	endpoint := fmt.Sprintf("%s/languages/delete-multiple", uriEnrollmentV3)

	// Create request body
	request := RequestDeleteMultipleLanguages{
		IDs: languageIds,
	}

	resp, err := c.HTTP.DoRequest("POST", endpoint, request, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDelete, "multiple enrollment language messages", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	var respErr ResponseError
	if err := json.NewDecoder(resp.Body).Decode(&respErr); err == nil {
		if respErr.HTTPStatus != 0 {
			return fmt.Errorf("deletion failed for multiple enrollment languages with status %d: %v", respErr.HTTPStatus, respErr.Errors)
		}
	}

	return nil
}

// GetEnrollmentLanguageCodes retrieves the list of available languages and their ISO 639-1 codes
func (c *Client) GetEnrollmentLanguageCodes() ([]ResourceLanguageCode, error) {
	endpoint := fmt.Sprintf("%s/language-codes", uriEnrollmentV3)
	var languageCodes []ResourceLanguageCode

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &languageCodes)
	if err != nil {
		return nil, fmt.Errorf("failed to get enrollment language codes: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return languageCodes, nil
}
