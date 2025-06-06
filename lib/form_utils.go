package lib

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmbracelet/huh"
)

// TechnologyOptions fetches technology options from the API and builds the options list
func GetTechnologyOptions() ([]huh.Option[string], error) {
	emptyTechnologyOptions := []huh.Option[string]{}
	res, err := http.Get("https://tinaciousdesign.com/api/technologies")
	if err != nil {
		return emptyTechnologyOptions, err
	}

	var technologies TinaciousDesignResponse[[]TechnologyItem]
	err = json.NewDecoder(res.Body).Decode(&technologies)
	if err != nil {
		return emptyTechnologyOptions, err
	}

	technologyOptions := make([]huh.Option[string], len(technologies.Data))
	for i, technology := range technologies.Data {
		technologyOptions[i] = huh.NewOption(technology.Title, technology.Slug)
	}

	return technologyOptions, nil
}

// AllPortfolioItems fetches all portfolio items from the API and builds the options list
func GetAllPortfolioItems() ([]huh.Option[string], error) {
	emptyPortfolioItemOptions := []huh.Option[string]{}
	res, err := http.Get("https://tinaciousdesign.com/api/portfolio")
	if err != nil {
		return emptyPortfolioItemOptions, err
	}

	var portfolioItems TinaciousDesignResponse[[]PortfolioItem]
	err = json.NewDecoder(res.Body).Decode(&portfolioItems)
	if err != nil {
		return emptyPortfolioItemOptions, err
	}

	portfolioItemOptions := make([]huh.Option[string], len(portfolioItems.Data))
	for i, item := range portfolioItems.Data {
		portfolioItemOptions[i] = huh.NewOption(fmt.Sprintf("%s <%s>", item.Title, item.Slug), item.Id)
	}

	return portfolioItemOptions, nil
}
