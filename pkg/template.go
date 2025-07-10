package pkg

type Template struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	AuthorURL   string `json:"author_url"`
	CloneURL    string `json:"clone_url"`
	Description string `json:"description"`
	IsOfficial  bool   `json:"-"`
}
