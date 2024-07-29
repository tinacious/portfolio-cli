package lib

import (
	"encoding/json"
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
