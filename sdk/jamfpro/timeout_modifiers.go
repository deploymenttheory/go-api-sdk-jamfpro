package jamfpro

import "time"

// Amends the HTTP timeout time
func (c *Client) ModifyHttpTimeout(newTimeout time.Duration) {
	c.HTTP.ModifyHttpTimeout(newTimeout)
}

// Resets HTTP timeout time back to 10 seconds
func (c *Client) ResetTimeout() {
	c.HTTP.ResetTimeout()
}

func (c *Client) HttpTimeout() time.Duration {
	return c.HTTP.HttpTimeout()
}
