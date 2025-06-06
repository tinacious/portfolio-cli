package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/stoewer/go-strcase"
	lib "github.com/tinacious/portfolio-cli/lib"
)

var techCmd = &cobra.Command{
	Use:   "tech",
	Short: "Create a curated portfolio by technologies",
	Long:  `Choose portfolio items based on the technologies. All portfolio items tagged with this technology will be included in the curated portfolio.`,
	Run: func(cmd *cobra.Command, args []string) {
		jwtSecret := lib.GetEnvironmentConfiguration().JwtSecret

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

		tokenData, err := lib.GetJwtTokenForTechnologiesData(clientName, pickedTechnologies, jwtSecret)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		baseUrl := lib.GetEnvironmentConfiguration().BaseUrl

		fmt.Printf("%s/curation/%s/%s\n", baseUrl, strcase.KebabCase(clientName), tokenData)
	},
}

func init() {
	rootCmd.AddCommand(techCmd)
}
