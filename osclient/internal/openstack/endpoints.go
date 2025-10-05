package openstack

import (
	"fmt"
	"strings"

	"github.com/gophercloud/gophercloud/openstack/identity/v3/endpoints"
)

type CronusEndpoints struct {
	Cronus string
	Nebula string
}

func (c *OpenStackClient) GetCronusEndpoints() (*CronusEndpoints, error) {

	// create identity service client
	svcClient, err := c.NewEndpointServiceClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create identity service client: %w", err)
	}
	// list all the openstack endpoints
	allPages, err := endpoints.List(svcClient, endpoints.ListOpts{}).AllPages()
	if err != nil {
		return nil, fmt.Errorf("failed to list endpoints: %w", err)
	}

	allEndpoints, err := endpoints.ExtractEndpoints(allPages)
	if err != nil {
		return nil, fmt.Errorf("failed to extract endpoints: %w", err)
	}

	ce := &CronusEndpoints{}

	for _, ep := range allEndpoints {
		// fmt.Printf("ID: %s | Name: %s | URL: %s | ServiceID: %s | Interface: %s\n", ep.ID, ep.Name, ep.URL, ep.ServiceID, ep.Availability)
		if strings.Contains(ep.URL, "https://cronus.") && strings.Contains(fmt.Sprintf("%v", ep.Availability), "public") {
			ce.Cronus = ep.URL
		}
		if strings.Contains(ep.URL, "https://nebula.") && strings.Contains(fmt.Sprintf("%v", ep.Availability), "public") {
			ce.Nebula = ep.URL
		}
	}
	return ce, nil
}

func (c *OpenStackClient) GetAllEndpoints() (*[]endpoints.Endpoint, error) {
	// list all the openstack endpoints
	svcClient, err := c.NewEndpointServiceClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create identity service client: %w", err)
	}

	allPages, err := endpoints.List(svcClient, endpoints.ListOpts{}).AllPages()
	if err != nil {
		return nil, fmt.Errorf("failed to list endpoints: %w", err)
	}

	allEndpoints, err := endpoints.ExtractEndpoints(allPages)
	if err != nil {
		return nil, fmt.Errorf("failed to extract endpoints: %w", err)
	}
	return &allEndpoints, nil
}
