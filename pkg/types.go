package pokemon

type Resource struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}
type Result struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}
