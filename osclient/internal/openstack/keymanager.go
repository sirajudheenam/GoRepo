package openstack

import (
	"fmt"
	"log"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/gophercloud/gophercloud/openstack/keymanager/v1/secrets"
)

type CreateOpts struct {
	Name        string
	Payload     string
	ContentType string
	SecretType  string
	Algorithm   string
	BitLength   int
	Mode        string
	Expiration  string
}

// CreateKeyManagerSecret creates a new secret in Barbican (OpenStack Key Manager)
func (c *OpenStackClient) CreateKeyManagerSecret(opts secrets.CreateOpts) (*secrets.Secret, error) {
	svcClient, err := c.NewKeyManagerClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create key manager service client: %w", err)
	}
	fmt.Printf("Creating secret with options: %+v\n", opts)

	cResult := secrets.Create(svcClient, opts)
	fmt.Printf("Create secret result: %+v\n", cResult)
	fmt.Printf("StatusCode: %+v\n", cResult.StatusCode)
	mySecret, err := cResult.Extract()
	return mySecret, err
}

// ListAllKeyManagerSecrets lists all secrets in Barbican (OpenStack Key Manager)
func (c *OpenStackClient) ListAllKeyManagerSecretsByType(secretType *secrets.SecretType) ([]secrets.Secret, error) {

	if secretType == nil {
		return []secrets.Secret{}, fmt.Errorf("secret type can't be empty")
	}

	svcClient, err := c.NewKeyManagerClient()
	if err != nil {
		fmt.Printf("failed to create key manager service client: %v", err)
		return []secrets.Secret{}, err
	}

	aclOnly := false
	listOpts := secrets.ListOpts{
		SecretType: *secretType,
		ACLOnly:    &aclOnly,
		Sort:       "created:desc",
		Limit:      3000,
	}

	log.Printf("%s: get barbican secret", svcClient.Context)
	logrus.Infof("%s: get barbican secret", svcClient.Context)
	allPages, err := secrets.List(svcClient, listOpts).AllPages()
	if err != nil {
		fmt.Printf("ERROR: %s: get barbican secret: %v", svcClient.Context, err)
		return nil, fmt.Errorf("unable to list secrets: %v", err)
	}

	return secrets.ExtractSecrets(allPages)
}

// ListAllKeyManagerSecrets lists all secrets in Barbican (OpenStack Key Manager)
func (c *OpenStackClient) ListAllKeyManagerSecrets() ([]secrets.Secret, error) {

	svcClient, err := c.NewKeyManagerClient()
	if err != nil {
		fmt.Printf("failed to create key manager service client: %w", err)
		return []secrets.Secret{}, err
	}

	aclOnly := false
	listOpts := secrets.ListOpts{
		// SecretType: "opaque",
		ACLOnly: &aclOnly,
		Sort:    "created:desc",
		Limit:   3000,
	}

	log.Printf("%s: get barbican secret", svcClient.Context)
	logrus.Infof("%s: get barbican secret", svcClient.Context)
	allPages, err := secrets.List(svcClient, listOpts).AllPages()
	if err != nil {
		fmt.Printf("ERROR: %s: get barbican secret: %v", svcClient.Context, err)
		return nil, fmt.Errorf("unable to list secrets: %v", err)
	}

	return secrets.ExtractSecrets(allPages)
}

// DeleteSecret deletes a secret in Barbican (OpenStack Key Manager) by its ID
func (c *OpenStackClient) DeleteKeyManagerSecretByID(secretID string) *secrets.DeleteResult {

	svcClient, err := c.NewKeyManagerClient()
	if err != nil {
		fmt.Printf("failed to create key manager service client: %+v", err)
		return nil
	}
	delResults := secrets.Delete(svcClient, secretID)
	return &delResults
}

// Retrieve Secret ID from URL

func GetSecretIDFromURL(secretRefURL string) (string, error) {
	// secretRefURL: https://keymanager-3.region.url:443/v1/secrets/<ID>
	parts := strings.Split(secretRefURL, "/")
	if len(parts) < 1 {
		return "", fmt.Errorf("invalid secret URL: %s", secretRefURL)
	}
	return parts[len(parts)-1], nil
}
