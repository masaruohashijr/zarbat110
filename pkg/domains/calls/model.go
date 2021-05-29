package calls

type Call struct {
	From string `json:"from"`
	To   string `json:"to"`
	Url  string `json:"url"`
}
