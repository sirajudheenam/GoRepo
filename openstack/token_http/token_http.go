package token_http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Config represents the structure of the configuration file
type Config struct {
	AuthURL    string `json:"auth_url"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	ProjectID  string `json:"project_id"`
	DomainName string `json:"domain_name"`
}

// TokenResponse represents the structure of the token response from OpenStack
type TokenResponse struct {
	Token struct {
		ID string `json:"id"`
	} `json:"token"`
}

// Read from JSON file
// LoadConfig reads the configuration file and unmarshals it into a Config struct
// func LoadConfig(filename string) (*Config, error) {
// 	data, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read config file: %v", err)
// 	}

// 	var config Config
// 	err = json.Unmarshal(data, &config)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal config file: %v", err)
// 	}

// 	return &config, nil
// }

// LoadConfig reads the configuration from environment variables
// LoadConfig reads the configuration from environment variables
func LoadConfig() (*Config, error) {
	config := &Config{
		AuthURL:    os.Getenv("OS_AUTH_URL"),
		Username:   os.Getenv("OS_USERNAME"),
		Password:   os.Getenv("OS_PASSWORD"),
		ProjectID:  os.Getenv("OS_PROJECT_ID"),
		DomainName: os.Getenv("OS_DOMAIN_NAME"),
	}

	// Check if any required environment variables are missing and log individual errors
	missingVars := false

	if config.AuthURL == "" {
		log.Println("Error: OS_AUTH_URL environment variable is missing")
		missingVars = true
	}
	if config.Username == "" {
		log.Println("Error: OS_USERNAME environment variable is missing")
		missingVars = true
	}
	if config.Password == "" {
		log.Println("Error: OS_PASSWORD environment variable is missing")
		missingVars = true
	}
	if config.ProjectID == "" {
		log.Println("Error: OS_PROJECT_ID environment variable is missing")
		missingVars = true
	}
	if config.DomainName == "" {
		log.Println("Error: OS_DOMAIN_NAME environment variable is missing")
		missingVars = true
	}

	if missingVars {
		return nil, fmt.Errorf("one or more required environment variables are missing")
	}

	return config, nil
}

// GetToken retrieves a token from the OpenStack Identity service (Keystone)
func GetToken(config *Config) (string, error) {
	authData := map[string]interface{}{
		"auth": map[string]interface{}{
			"identity": map[string]interface{}{
				"methods": []string{"password"},
				"password": map[string]interface{}{
					"user": map[string]interface{}{
						"name":     config.Username,
						"password": config.Password,
						"domain": map[string]interface{}{
							"name": config.DomainName,
						},
					},
				},
			},
			"scope": map[string]interface{}{
				"project": map[string]interface{}{
					"id": config.ProjectID,
				},
			},
		},
	}

	authJSON, err := json.Marshal(authData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal auth data: %v", err)
	}

	req, err := http.NewRequest("POST", config.AuthURL+"/auth/tokens", bytes.NewBuffer(authJSON))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to get token: %s", body)
	}

	token := resp.Header.Get("X-Subject-Token")
	if token == "" {
		return "", fmt.Errorf("token not found in response headers")
	}

	return token, nil
}

func GetOSTokenHttp() (token string) {
	// Load the configuration
	config, err := LoadConfig()
	fmt.Println("config [main]: ", config)

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	token, err = GetToken(config)
	if err != nil {
		log.Fatalf("Error getting token: %v", err)
	}

	fmt.Printf("Successfully retrieved token from OpenStack via raw HTTP method: %s\n", token)
	return token
}
