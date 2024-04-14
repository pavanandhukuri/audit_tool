package adapters

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"security_audit_tool/domain"
	"security_audit_tool/domain/entities"
)

type OrgResponse struct {
	MembersCanCreatePublicRepositories bool `json:"members_can_create_public_repositories"`
}

type GithubVersionControlSystemAdapter struct {
	config domain.Configuration
}

func (adapter *GithubVersionControlSystemAdapter) GetInfo() (entities.VersionControlData, error) {

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "https://api.github.com/orgs/"+adapter.config.GithubOrgName, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	// Add headers to the request
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+adapter.config.GithubToken)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	// Send the request using the client
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return entities.VersionControlData{}, err

	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return entities.VersionControlData{}, err
	}

	// Unmarshal the JSON response into the struct
	var orgResponse OrgResponse
	err = json.Unmarshal(body, &orgResponse)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return entities.VersionControlData{}, err
	}

	return entities.VersionControlData{
		CanMembersCreatePublicRepositories: orgResponse.MembersCanCreatePublicRepositories,
	}, nil
}

func NewGithubVersionControlSystemAdapter(configuration domain.Configuration) *GithubVersionControlSystemAdapter {
	return &GithubVersionControlSystemAdapter{configuration}
}
