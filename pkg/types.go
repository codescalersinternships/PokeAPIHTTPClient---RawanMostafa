package pokemon

type Resource struct {
	Count    int
	Next     string
	Previous string
	Result   []Result
}
type Result struct {
	Url  string
	Name string
}
