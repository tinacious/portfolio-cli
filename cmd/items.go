package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/stoewer/go-strcase"
	lib "github.com/tinacious/portfolio-cli/lib"
)

var itemsCmd = &cobra.Command{
	Use:   "items",
	Short: "Create a curated portfolio by choosing the portfolio items",
	Long:  `Choose portfolio items individually. Only items chosen will be included in the curated portfolio.`,
	Run: func(cmd *cobra.Command, args []string) {
		jwtSecret := lib.GetEnvironmentConfiguration().JwtSecret

		// Final form values
		var clientName string
		var pickedItemIds []string

		// Get Client name
		inputClientName := huh.NewGroup(
			huh.NewInput().Title("Client name").Value(&clientName),
		)

		// Get portfolio items
		portfolioItemOptions, err := lib.GetAllPortfolioItems()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		inputPortfolioItemsOptions := huh.NewGroup(
			huh.NewMultiSelect[string]().
				Options(portfolioItemOptions...).
				Title("Choose portfolio items").
				Value(&pickedItemIds),
		)

		// Render the form with theme
		form := huh.NewForm(
			inputClientName,
			inputPortfolioItemsOptions,
		)
		form.WithTheme(lib.TinaciousDesignTheme())
		form.Run()

		tokenData, err := lib.GetJwtTokenForPortfolioItemsData(clientName, pickedItemIds, jwtSecret)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		baseUrl := lib.GetEnvironmentConfiguration().BaseUrl

		fmt.Printf("%s/curation/%s/%s\n", baseUrl, strcase.KebabCase(clientName), tokenData)
	},
}

func init() {
	rootCmd.AddCommand(itemsCmd)
}
