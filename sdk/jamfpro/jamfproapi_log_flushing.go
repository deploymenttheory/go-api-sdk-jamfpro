// jamfproapi_log_flushing.go
// Jamf Pro Api - Log Flushing
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-log-flushing
// Jamf Pro API requires the structs to support an JSON data structure.
package jamfpro

import "fmt"

const uriLogFlushing = "/api/v1/log-flushing"

// Resource/Response structs for log flushing settings
type ResponseLogFlushingSettings struct {
	RetentionPolicies []LogRetentionPolicy `json:"retentionPolicies"`
	HourOfDay         int                  `json:"hourOfDay"`
}

type LogRetentionPolicy struct {
	DisplayName         string `json:"displayName"`
	Qualifier           string `json:"qualifier"`
	RetentionPeriod     int    `json:"retentionPeriod"`
	RetentionPeriodUnit string `json:"retentionPeriodUnit"`
}

// Request struct for creating log flushing task
type ResourceLogFlushingTask struct {
	Qualifier           string `json:"qualifier"`
	RetentionPeriod     int    `json:"retentionPeriod"`
	RetentionPeriodUnit string `json:"retentionPeriodUnit"`
}

// Response struct for created log flushing task
type ResponseLogFlushingTaskCreated struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// Response struct for log flushing tasks
type LogFlushingTask struct {
	ID                  string `json:"id"`
	Qualifier           string `json:"qualifier"`
	RetentionPeriod     int    `json:"retentionPeriod"`
	RetentionPeriodUnit string `json:"retentionPeriodUnit"`
	State               string `json:"state"`
}

// GetLogFlushingTasks retrieves the list of log flushing tasks from Jamf Pro
func (c *Client) GetLogFlushingTasks() ([]LogFlushingTask, error) {
	endpoint := fmt.Sprintf("%s/task", uriLogFlushing)

	var out []LogFlushingTask
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "log flushing tasks", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return out, nil
}

// GetLogFlushingTaskByID retrieves a specific log flushing task by its ID from Jamf Pro
func (c *Client) GetLogFlushingTaskByID(id string) (*LogFlushingTask, error) {
	endpoint := fmt.Sprintf("%s/task/%s", uriLogFlushing, id)

	var task LogFlushingTask
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &task)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, fmt.Sprintf("log flushing task with ID %s", id), err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &task, nil
}

// GetLogFlushingSettings retrieves the current log flushing settings from Jamf Pro
func (c *Client) GetLogFlushingSettings() (*ResponseLogFlushingSettings, error) {
	endpoint := uriLogFlushing

	var out ResponseLogFlushingSettings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "log flushing settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// QueueLogFlushingTaskcreates a new log flushing task in Jamf Pro
func (c *Client) QueueLogFlushingTask(task *ResourceLogFlushingTask) (*ResponseLogFlushingTaskCreated, error) {
	endpoint := fmt.Sprintf("%s/task", uriLogFlushing)

	var response ResponseLogFlushingTaskCreated
	resp, err := c.HTTP.DoRequest("POST", endpoint, task, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "log flushing task", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeleteLogFlushingTaskByID deletes a specific log flushing task by its ID from Jamf Pro
func (c *Client) DeleteLogFlushingTaskByID(id string) error {
	endpoint := fmt.Sprintf("%s/task/%s", uriLogFlushing, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDelete, fmt.Sprintf("log flushing task with ID %s", id), err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
