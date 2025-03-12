package jamfpro

import "time"

func (c *Client) ModifyHttpTimeout(newTimeout time.Duration) {
	c.HTTP.ModifyHttpTimeout(newTimeout)
}
