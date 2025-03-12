package jamfpro

import "time"

// Modifies the HTTP client timeout time for longer requests (uploads)
func (c *Client) ModifyHttpTimeout(newTimeout time.Duration) {
	c.HTTP.ModifyHttpTimeout(newTimeout)
}

// Resets the timeout time back to 10s
func (c *Client) ResetHttpTimeout() {
	c.HTTP.ResetTimeout()
}
