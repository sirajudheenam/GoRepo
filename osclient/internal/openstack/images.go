package openstack

import (
	"fmt"

	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
)

// ListImages lists all images from the OpenStack Glance service
func (c *OpenStackClient) ListImages() ([]images.Image, error) {
	imageClient, err := c.NewImageServiceClient()
	if err != nil {
		fmt.Printf("Error creating image service client: %v", err)
		return nil, fmt.Errorf("failed to create image service client: %w", err)
	}

	// Use our Wrapper struct to include logging
	svc := ServiceClient{Client: imageClient, Log: c.Log}

	imagesListOpts := images.ListOpts{
		Limit: 10,
	}
	allPages, err := images.List(svc.Client, imagesListOpts).AllPages()

	if err != nil {
		fmt.Printf("Error listing images: %v", err)
		svc.Log.Errorf("Error listing images: %v", err)
		return nil, fmt.Errorf("failed to list images: %w", err)
	}

	allImages, err := images.ExtractImages(allPages)
	if err != nil {
		fmt.Printf("Error extracting images: %v", err)
		svc.Log.Errorf("Error extracting images: %v", err)
		return nil, fmt.Errorf("failed to extract images: %w", err)
	}

	return allImages, nil
}
