package config

import (
	"github.com/joho/godotenv"
	"os"
	"security_audit_tool/domain"
	"security_audit_tool/logger"
)

func Load() (*domain.Configuration, error) {
	err := godotenv.Load(".env")
	if err != nil {
		logger.LogWarn("Unable to load .env file. Please make sure necessary environment variables are set.")
		return nil, err
	}

	return &domain.Configuration{
		GithubToken:   os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"),
		GithubOrgName: os.Getenv("GITHUB_ORG_NAME"),
	}, nil
}
