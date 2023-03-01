package entity

type Article struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type ArticleParam struct {
	Author string `json:"author"`
	Query  string `json:"query"`
}

func (Article) TableName() string {
	return "article"
}
