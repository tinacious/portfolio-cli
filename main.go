package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/stoewer/go-strcase"
	lib "github.com/tinacious/portfolio-cli/lib"
)

func main() {
	jwtSecret := os.Getenv("TINACIOUS_DESIGN_JWT_SECRET")
	if jwtSecret == "" {
		fmt.Println("💥 environment variable TINACIOUS_DESIGN_JWT_SECRET not set")
		os.Exit(1)
	}

	// Final form values
	var clientName string
	var pickedTechnologies []string

	// Get Client name
	inputClientName := huh.NewGroup(
		huh.NewInput().Title("Client name").Value(&clientName),
	)

	// Get technologies
	technologyOptions, err := lib.GetTechnologyOptions()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputTechnologyOptions := huh.NewGroup(
		huh.NewMultiSelect[string]().
			Options(technologyOptions...).
			Title("Choose technologies").
			Value(&pickedTechnologies),
	)

	// Render the form with theme
	form := huh.NewForm(
		inputClientName,
		inputTechnologyOptions,
	)
	form.WithTheme(lib.TinaciousDesignTheme())
	form.Run()

	tokenData, err := lib.GetJwtTokenForTokenData(clientName, pickedTechnologies, jwtSecret)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("https://tinaciousdesign.com/curation/%s/%s\n", strcase.KebabCase(clientName), tokenData)
	// fmt.Printf("http://localhost:7777/curation/%s/%s\n", strcase.KebabCase(clientName), tokenData)
}
