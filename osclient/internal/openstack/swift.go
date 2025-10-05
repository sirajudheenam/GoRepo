package openstack

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gophercloud/gophercloud/openstack/objectstorage/v1/accounts"
	"github.com/gophercloud/gophercloud/openstack/objectstorage/v1/containers"
	"github.com/gophercloud/gophercloud/openstack/objectstorage/v1/objects"
	"github.com/gophercloud/gophercloud/pagination"

	"strings"
)

func (c *OpenStackClient) CreateNewSwiftContainer(containerName string) error {

	if strings.TrimSpace(containerName) == "" {
		return fmt.Errorf("container name cannot be empty")
	}

	createOpts := containers.CreateOpts{
		ContainerRead:     ".r:*,.rlistings", // e.g., ".r:*,.rlistings" for public read access
		ContainerWrite:    ".w:*,.rlistings", // e.g., ".w:*" for public write access
		ContainerSyncTo:   "",
		ContainerSyncKey:  "",
		ContentType:       "", // MIME type for the container metadata request (usually "application/json" or empty
		DetectContentType: true,
		IfNoneMatch:       "*",        // only create if it doesn’t already exist
		VersionsLocation:  "Versions", // Only one of x-versions-location or x-history-location may be specified
		HistoryLocation:   "",         // Only one of x-versions-location or x-history-location may be specified
		TempURLKey:        "mySecretTempURLKey",
		TempURLKey2:       "mySecretTempURLKey2",
		StoragePolicy:     "standard", // or "high" for high availability or "cold" storage
		VersionsEnabled:   true,
	}

	svcClient, err := c.NewObjectStorageSwiftClient()
	if err != nil {
		return fmt.Errorf("failed to create object storage (Swift) service client: %w", err)
	}
	r := containers.Create(svcClient, containerName, createOpts)
	if r.Err != nil {
		return fmt.Errorf("failed to create container: %w", r.Err)
	}

	c.Log.Infof("Container %s created successfully", containerName)

	return nil
}
func (c *OpenStackClient) ListAllSwiftContainers() ([]containers.Container, error) {
	var allContainers []containers.Container

	svcClient, err := c.NewObjectStorageSwiftClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create object storage (Swift) service client: %w", err)
	}

	listOpts := containers.ListOpts{
		Full:  true, // Retrieve full details
		Limit: 1000,
	}

	pager := containers.List(svcClient, listOpts)

	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		containerList, err := containers.ExtractInfo(page)
		if err != nil {
			return false, fmt.Errorf("failed to extract containers: %w", err)
		}
		allContainers = append(allContainers, containerList...)
		return true, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list containers: %w", err)
	}

	return allContainers, nil
}

// UploadFolder recursively uploads a folder's contents to a Swift container.
func (c *OpenStackClient) UploadFolder(containerName string, localPath string) error {
	// Clean up any trailing slashes
	localPath = filepath.Clean(localPath)

	return filepath.Walk(localPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories — Swift has no real directories
		if info.IsDir() {
			return nil
		}

		// Compute relative path to preserve folder structure
		relPath, err := filepath.Rel(localPath, path)
		if err != nil {
			return err
		}

		// Use forward slashes for Swift (even on Windows)
		relPath = strings.ReplaceAll(relPath, string(filepath.Separator), "/")

		// Open file
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		// Upload object
		createOpts := objects.CreateOpts{
			Content:     f,
			ContentType: "application/octet-stream",
			IfNoneMatch: "*",
		}

		log.Printf("Uploading: %s → %s/%s", path, containerName, relPath)

		svcClient, err := c.NewObjectStorageSwiftClient()
		if err != nil {
			return fmt.Errorf("failed to create object storage (Swift) service client: %w", err)
		}
		result := objects.Create(svcClient, containerName, relPath, createOpts)
		if result.Err != nil {
			return fmt.Errorf("failed to upload %s: %w", path, result.Err)
		}

		return nil
	})
}

// DeleteFolder recursively deletes all objects within a folder prefix.
func (c *OpenStackClient) DeleteFolder(containerName, folderPrefix string) error {

	svcClient, err := c.NewObjectStorageSwiftClient()
	if err != nil {
		return fmt.Errorf("failed to create object storage (Swift) service client: %w", err)
	}
	// Normalize prefix: ensure it ends with a slash
	if folderPrefix != "" && folderPrefix[len(folderPrefix)-1] != '/' {
		folderPrefix += "/"
	}

	listOpts := objects.ListOpts{
		Full:   false,
		Prefix: folderPrefix,
	}

	pager := objects.List(svcClient, containerName, listOpts)

	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		objectNames, err := objects.ExtractNames(page)
		if err != nil {
			return false, fmt.Errorf("failed to extract objects: %w", err)
		}

		for _, objName := range objectNames {
			log.Printf("Deleting: %s/%s", containerName, objName)
			result := objects.Delete(svcClient, containerName, objName, nil)
			if result.Err != nil {
				log.Printf("⚠️ failed to delete %s: %v", objName, result.Err)
			}
		}

		return true, nil
	})
	if err != nil {
		return fmt.Errorf("failed to list or delete objects: %w", err)
	}

	log.Printf("✅ Folder '%s' and its contents deleted successfully.", folderPrefix)
	return nil
}

func (c *OpenStackClient) GetAccountMetaData() (map[string]string, error) {
	var rAccountsResult accounts.GetResult
	svcClient, err := c.NewObjectStorageSwiftClient()
	if err != nil {
		return nil, err
	}
	rAccountsResult = accounts.Get(svcClient, accounts.GetOpts{
		Newest: true,
	})
	metaData, err := rAccountsResult.ExtractMetadata()
	return metaData, err
}

func (c *OpenStackClient) UpdateMetaData(metaData map[string]string) accounts.UpdateResult {
	svcClient, _ := c.NewObjectStorageSwiftClient()
	updateOpts := accounts.UpdateOpts{
		Metadata: metaData,
	}
	updateResult := accounts.Update(svcClient, updateOpts)
	return updateResult
}
