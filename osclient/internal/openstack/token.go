package openstack

import "fmt"

func (c *OpenStackClient) GetToken() (string, error) {
	logg := c.Log
	// Get the token from the provider
	token := c.Provider.Token()
	if token == "" {
		logg.Errorf("Failed to get Token")
		return "", fmt.Errorf("failed to get token: ")
	}
	return token, nil
}
