package github

import (
	"security_audit_tool/domain"
	"security_audit_tool/domain/entities"
	"security_audit_tool/util"
)

type OrgResponse struct {
	CanMembersCreatePublicRepositories bool `json:"members_can_create_public_repositories"`
}

type RepoResponse struct {
	Name      string `json:"full_name"`
	IsPrivate bool   `json:"private"`
}

func (repo RepoResponse) toRepository() entities.Repository {
	return entities.Repository{
		Name:      repo.Name,
		IsPrivate: repo.IsPrivate,
	}
}

func (adapter *VersionControlSystemAdapter) convertRepos(repos []RepoResponse) []entities.Repository {
	var repositories []entities.Repository
	for _, repo := range repos {
		repositories = append(repositories, repo.toRepository())
	}
	return repositories

}

type VersionControlSystemAdapter struct {
	config domain.Configuration
}

func NewGithubVersionControlSystemAdapter(configuration domain.Configuration) *VersionControlSystemAdapter {
	return &VersionControlSystemAdapter{configuration}
}

func (adapter *VersionControlSystemAdapter) GetData() (entities.VersionControlSystem, error) {

	// Create a new HTTP request
	orgResponse, err := adapter.getOrgData()
	if err != nil {
		return entities.VersionControlSystem{}, err
	}
	repositories, err := adapter.getRepositories()
	if err != nil {
		return entities.VersionControlSystem{}, err
	}

	return entities.VersionControlSystem{
		CanMembersCreatePublicRepositories: orgResponse.CanMembersCreatePublicRepositories,
		Repositories:                       adapter.convertRepos(repositories),
	}, nil
}

func (adapter *VersionControlSystemAdapter) getOrgData() (OrgResponse, error) {
	url := githubAPIURLPrefix + adapter.config.GithubOrgName
	orgResponse, err := util.GetAsType[OrgResponse](url, adapter.getHeaders())
	if err != nil {
		return OrgResponse{}, err
	}
	return orgResponse, nil
}

func (adapter *VersionControlSystemAdapter) getRepositories() ([]RepoResponse, error) {
	url := githubAPIURLPrefix + adapter.config.GithubOrgName + "/repos"
	repoResponses, err := util.GetAsType[[]RepoResponse](url, adapter.getHeaders())
	if err != nil {
		return []RepoResponse{}, err
	}
	return repoResponses, nil
}

func (adapter *VersionControlSystemAdapter) getHeaders() map[string]string {
	return map[string]string{
		"Accept":               githubAPIAcceptHeader,
		"Authorization":        "Bearer " + adapter.config.GithubToken,
		"X-GitHub-Api-Version": githubAPIVersion,
	}
}
