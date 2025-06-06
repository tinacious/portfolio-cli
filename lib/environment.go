package lib

import (
	"fmt"
	"os"
)

type EnvironmentConfiguration struct {
	JwtSecret string
	BaseUrl   string
}

func GetEnvironmentConfiguration() EnvironmentConfiguration {
	jwtSecret := os.Getenv("TINACIOUS_DESIGN_JWT_SECRET")
	if jwtSecret == "" {
		fmt.Println("💥 environment variable TINACIOUS_DESIGN_JWT_SECRET not set")
		os.Exit(1)
	}

	baseUrl := os.Getenv("TINACIOUS_DESIGN_BASE_URL")
	if baseUrl == "" {
		baseUrl = "https://tinaciousdesign.com"
	}

	return EnvironmentConfiguration{
		JwtSecret: jwtSecret,
		BaseUrl:   baseUrl,
	}
}
