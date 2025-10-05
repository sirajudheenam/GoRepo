package openstack

import (
	"context"
	"fmt"
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/sirupsen/logrus"
)

type OpenStackClient struct {
	Provider *gophercloud.ProviderClient
	Log      *logrus.Logger
}

type ServiceClient struct {
	Client *gophercloud.ServiceClient
	Log    *logrus.Logger
}

func getRegion() string {
	if os.Getenv("OS_REGION_NAME") == "" {
		fmt.Println("Warning: OS_REGION_NAME is not set. Using default region behavior.")
		return ""
	}
	return os.Getenv("OS_REGION_NAME")
}

// NewOpenStackClient authenticates using environment variables
func NewOpenStackClient() (*OpenStackClient, error) {
	authOptions, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, fmt.Errorf("auth from env failed: %w", err)
	}

	provider, err := openstack.AuthenticatedClient(authOptions)
	if err != nil {
		return nil, fmt.Errorf("auth failed: %w", err)
	}

	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	provider.Context = context.Background()

	region := getRegion()
	if region == "" {
		log.Warn("OS_REGION_NAME not set; using default region behavior")
	}

	return &OpenStackClient{
		Provider: provider,
		Log:      log,
	}, nil
}

// Image (Glance) client
func (c *OpenStackClient) NewImageServiceClient() (*gophercloud.ServiceClient, error) {
	return openstack.NewImageServiceV2(c.Provider, gophercloud.EndpointOpts{
		Region: getRegion(),
	})
}

// Identity (Keystone) client
func (c *OpenStackClient) NewEndpointServiceClient() (*gophercloud.ServiceClient, error) {
	return openstack.NewIdentityV3(c.Provider, gophercloud.EndpointOpts{
		Region: getRegion(),
	})
}

// Key Manager (Barbican) client
func (c *OpenStackClient) NewKeyManagerClient() (*gophercloud.ServiceClient, error) {
	return openstack.NewKeyManagerV1(c.Provider, gophercloud.EndpointOpts{
		Region: getRegion(),
	})
}

// Object Storage (Swift) client
func (c *OpenStackClient) NewObjectStorageSwiftClient() (*gophercloud.ServiceClient, error) {
	return openstack.NewObjectStorageV1(c.Provider, gophercloud.EndpointOpts{
		Region: getRegion(),
	})
}
