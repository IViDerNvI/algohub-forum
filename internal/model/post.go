package model

type Post struct {
	ObjectMeta
	UserId  string `json:"userid" gorm:"column:userid"`
	Content string `json:"content" gorm:"column:content"`
	Title   string `json:"title" gorm:"column:title"`
}

type PostList struct {
	ListMeta
	Items []Post
}

func (p *Post) GetDatabaseName() string {
	return "post"
}
