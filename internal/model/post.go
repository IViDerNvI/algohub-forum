package model

type Post struct {
	ObjectMeta
	Id         int
	UserId     int
	Content    string
	LikeCnt    int
	BadCnt     int
	Title      string
	PostTime   string
	CommentCnt int
}

type PostLike struct {
	ObjectMeta
	UserId int
	PostId int
}

type PostList struct {
	ListMeta
	Items []Post
}
