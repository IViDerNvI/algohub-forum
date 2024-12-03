package model

type Problem struct {
	ObjectMeta
	Title          string `json:"title,omitempty" gorm:"column:title;unique;index"`
	SubmissionTime int    `json:"submissionTime,omitempty" gorm:"column:submissionTime"`
	PassTime       int    `json:"passTime,omitempty" gorm:"column:passTime"`
	Difficulty     int    `json:"difficulty" gorm:"column:difficulty"`
	Tags           string `json:"tags,omitempty" gorm:"column:tags"`
}

type ProblemList struct {
	ListMeta
	Items []Problem
}
