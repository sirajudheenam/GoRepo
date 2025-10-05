package main

import (
	"fmt"
	"log"

	"github.com/gophercloud/gophercloud"

	"github.com/gophercloud/gophercloud/openstack/keymanager/v1/secrets"
	"github.com/sirupsen/logrus"

	"github.com/sirajudheenam/osclient/internal/openstack"
)

// OpenStackClient holds the authenticated provider and allows creating service clients
type OpenStackClient struct {
	Provider *gophercloud.ProviderClient
}

func createBulkOfOpaqueSecrets(client *openstack.OpenStackClient, count int) ([]*secrets.Secret, error) {
	var createdSecrets []*secrets.Secret
	if count < 1 {
		return createdSecrets, fmt.Errorf("count must be at least 1")
	}
	for i := range count {
		secretOpts := &secrets.CreateOpts{
			Name:               fmt.Sprintf("my-secret-%d", i),
			Algorithm:          "aes",
			BitLength:          256,
			Mode:               "cbc",
			Payload:            fmt.Sprintf("my-secret-payload-%d", i),
			PayloadContentType: "text/plain", // ['text/plain', 'text/plain;charset=utf-8', 'text/plain; charset=utf-8', 'application/octet-stream', 'application/pkcs8']
		}

		createdSecret, err := client.CreateKeyManagerSecret(*secretOpts)
		if err != nil {
			logrus.Fatalf("Error creating barbican secret using osclient: %v", err)
			return nil, err
		}
		logrus.Printf("Created secret using osclient: | Name: %s | SecretRef: %s\n", createdSecret.Name, createdSecret.SecretRef)
		createdSecrets = append(createdSecrets, createdSecret)
	}
	return createdSecrets, nil
}

// Delete all Opaque secrets in Barbican (OpenStack Key Manager)

func deleteBulkOpaqueSecrets(client *openstack.OpenStackClient) ([]*secrets.Secret, error) {
	var deletedSecrets []*secrets.Secret

	secretType := secrets.OpaqueSecret

	// List all secrets
	allSecrets, err := client.ListAllKeyManagerSecretsByType(&secretType)
	if err != nil {
		return nil, fmt.Errorf("failed to list secrets: %w", err)
	}
	// Delete each secret
	for _, secret := range allSecrets {
		if secretID, err := openstack.GetSecretIDFromURL(secret.SecretRef); err != nil {
			logrus.Errorf("Error extracting secret ID from URL: %v", err)
			return nil, err
		} else {
			fmt.Printf("Deleting secret with ID: %s\n", secretID)
			if delRes := client.DeleteKeyManagerSecretByID(secretID); delRes.Err != nil {
				fmt.Printf("Error deleting barbican secret using osclient: %v\n", delRes.Err)
				return nil, delRes.Err
			}
			fmt.Printf("Deleted secret with ID: %s\n", secretID)
		}
	}
	return deletedSecrets, nil
}

func deleteAllSecretsFromKeyManager(client *openstack.OpenStackClient) error {
	// List all secrets
	allSecrets, err := client.ListAllKeyManagerSecrets()
	if err != nil {
		return fmt.Errorf("failed to list secrets: %w", err)
	}
	// Delete each secret
	for _, secret := range allSecrets {
		if secretID, err := openstack.GetSecretIDFromURL(secret.SecretRef); err != nil {
			fmt.Printf("Error extracting secret ID from URL: %v", err)
			return err
		} else {
			fmt.Printf("Deleting secret with ID: %s\n", secretID)
			if delRes := client.DeleteKeyManagerSecretByID(secretID); delRes.Err != nil {
				fmt.Printf("Error deleting barbican secret using osclient: %v", delRes.Err)
				return delRes.Err
			}
			fmt.Printf("Deleted secret with ID: %s\n", secretID)
		}
	}
	return nil
}

func main() {

	client, err := openstack.NewOpenStackClient()

	if err != nil {
		fmt.Printf("Error creating OpenStack client: %s", err)
	}

	// Assign the provider to OpenStackClient
	openstackClient := &openstack.OpenStackClient{Provider: client.Provider}

	// Get Token
	if token, err := openstackClient.GetToken(); err != nil {
		log.Fatalf("Error getting token: %v", err)
	} else {
		fmt.Printf("Token: %s\n", token)
	}

	// List all endpoints
	endpointList, err := openstackClient.GetAllEndpoints()
	if err != nil {
		fmt.Printf("Error listing endpoints: %v", err)
	}
	fmt.Printf("Found %d endpoints using osclient\n", len(*endpointList))

	for _, ep := range *endpointList {
		fmt.Printf("ID: %s | Name: %s | URL: %s | ServiceID: %s | Interface: %s\n", ep.ID, ep.Name, ep.URL, ep.ServiceID, ep.Availability)
	}

	// Get Cronus and Nebula endpoints
	if cronusEndpoints, err := openstackClient.GetCronusEndpoints(); err != nil {
		fmt.Printf("Error listing endpoints: %v", err)
	} else {
		fmt.Printf("Cronus Endpoint: %s\n", cronusEndpoints.Cronus)
		fmt.Printf("Nebula Endpoint: %s\n", cronusEndpoints.Nebula)
	}

	// List all images
	allImages, err := openstackClient.ListImages()
	if err != nil {
		fmt.Printf("Error listing images using osclient: %v", err)
	}
	fmt.Printf("Found %d images using osclient\n", len(allImages))
	for _, img := range allImages {
		fmt.Printf("ID: %s | Name: %s | Status: %s\n", img.ID, img.Name, img.Status)
	}

	// Key Manager (Barbican) operations

	// Bulk Creation of Opaque secrets in Barbican (OpenStack Key Manager)
	opaqueSecrets, err := createBulkOfOpaqueSecrets(openstackClient, 5)
	if err != nil {
		fmt.Printf("Error creating opaque secrets: %v", err)
	} else {
		fmt.Printf("Created %d opaque secrets\n", len(opaqueSecrets))
	}

	// List all secrets
	allSecrets, err := openstackClient.ListAllKeyManagerSecrets()
	if err != nil {
		fmt.Printf("Error listing all secrets: %v", err)
	} else {
		fmt.Printf("Listed %d secrets\n", len(allSecrets))
		for _, secret := range allSecrets {
			fmt.Printf("Creator ID: %s | Name: %s | Created: %s | Updated: %s | SecretRef: %s\n", secret.CreatorID, secret.Name, secret.Created, secret.Updated, secret.SecretRef)
		}
	}

	// Bulk Deletion of Opaque secrets in Barbican (OpenStack Key Manager)
	deletedResults, err := deleteBulkOpaqueSecrets(openstackClient)
	if err != nil {
		fmt.Printf("Error deleting opaque secrets: %v", err)
	} else {
		fmt.Printf("Deleted %d opaque secrets\n", len(deletedResults))
	}

	// Delete all secrets in Barbican (OpenStack Key Manager)
	if err := deleteAllSecretsFromKeyManager(openstackClient); err != nil {
		fmt.Printf("Error deleting all secrets: %v", err)
	} else {
		fmt.Printf("Deleted all secrets successfully\n")
	}

	// Container operations in Swift (OpenStack Object Storage) TODO: ERROR
	// containerName := "my-container"
	// if err := openstackClient.CreateNewSwiftContainer(containerName); err != nil {
	// 	fmt.Printf("Error creating Swift container: %v", err)
	// } else {
	// 	fmt.Printf("Swift container '%s' created successfully\n", containerName)
	// }

	// List all Swift containers - WORKS
	containers, err := openstackClient.ListAllSwiftContainers()
	if err != nil {
		fmt.Printf("Error listing Swift containers: %v", err)
	} else {
		fmt.Printf("Found %d Swift containers:\n", len(containers))
		for _, container := range containers {
			fmt.Printf("Name: %s | Count: %d | Bytes: %d\n", container.Name, container.Count, container.Bytes)
		}
	}

	// Upload a folder to Swift container - WORKS
	localFolderPath := "internal"
	containerName := "internal-container"
	if err := openstackClient.UploadFolder(containerName, localFolderPath); err != nil {
		fmt.Printf("Error uploading folder to Swift container: %v", err)
	} else {
		fmt.Printf("Folder '%s' uploaded to Swift container '%s' successfully\n", localFolderPath, containerName)
	}

	// Delete a folder from Swift container - WORKS
	folderPrefix := "openstack"
	if err := openstackClient.DeleteFolder(containerName, folderPrefix); err != nil {
		fmt.Printf("Error deleting folder from Swift container: %v", err)
	} else {
		fmt.Printf("Folder '%s' deleted from Swift container '%s' successfully\n", folderPrefix, containerName)
	}

	// Get account metadata
	if metadata, err := openstackClient.GetAccountMetaData(); err != nil {
		fmt.Printf("Error getting account metadata: %v", err)
	} else {
		fmt.Printf("Account Metadata:\n")
		for key, value := range metadata {
			fmt.Printf("%+v: %+v\n", key, value)
		}
	}

	// Update account metadata
	newMetadata := map[string]string{
		"X-Account-Meta-Test": "TestValue",
	}
	if updateResult := openstackClient.UpdateMetaData(newMetadata); err != nil {
		fmt.Printf("Error updating account metadata: %v", err)
	} else {
		fmt.Printf("Account metadata updated successfully. New Metadata:\n")
		updateResult.ExtractInto(&newMetadata)
		for key, value := range newMetadata {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
	// Sample Output:
	// Account Metadata:
	// Quota-Bytes: 4398046511104
	// Account metadata updated successfully. New Metadata:
	// Date: Sun, 05 Oct 2025 16:47:12 GMT
	// X-Openstack-Request-Id: txead8e3bfd5a24718abbba-0068e2a110
	// X-Trans-Id: <RANDOM-UUID>
	// X-Account-Meta-Test: TestValue
	// Content-Length: 0
	// Content-Type: text/html; charset=UTF-8

}
