package tokens

import (
	"fmt"
	"log"
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

func GetOSToken() (tokenID string) {
	// Read OpenStack authentication details from environment variables
	authURL := os.Getenv("OS_AUTH_URL")
	username := os.Getenv("OS_USERNAME")
	password := os.Getenv("OS_PASSWORD")
	domainName := os.Getenv("OS_DOMAIN_NAME")
	projectID := os.Getenv("OS_PROJECT_ID")

	// Check for missing environment variables
	if authURL == "" {
		log.Fatal("Error: OS_AUTH_URL environment variable is missing")
	}
	if username == "" {
		log.Fatal("Error: OS_USERNAME environment variable is missing")
	}
	if password == "" {
		log.Fatal("Error: OS_PASSWORD environment variable is missing")
	}
	if domainName == "" {
		log.Fatal("Error: OS_DOMAIN_NAME environment variable is missing")
	}
	if projectID == "" {
		log.Fatal("Error: OS_PROJECT_ID environment variable is missing")
	}

	// Define the authentication options
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: authURL,
		Username:         username,
		Password:         password,
		DomainName:       domainName,
		Scope:            &gophercloud.AuthScope{ProjectID: projectID},
	}

	// Authenticate and obtain a provider client
	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		log.Fatalf("Error authenticating: %v", err)
	}

	// Extract the token ID from the provider client
	tokenID = provider.TokenID

	fmt.Printf("[gophercloud][openstack][provider][TokenID]: %s\n", tokenID)

	return tokenID
}
