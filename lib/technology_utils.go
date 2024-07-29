package lib

type TinaciousDesignResponse[T any] struct {
	Data T `json:"data"`
}

type TechnologyItem struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}
