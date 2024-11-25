package model

type Comment struct {
	Id          int
	UserId      int
	ReferPostId int
	Content     string
	LikeCnt     int
	BadCnt      int
}

type CommentLike struct {
	UserId    int
	CommentId int
}
