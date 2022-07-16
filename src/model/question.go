package model

type Question struct {
	Id           string `json:"id"`
	QuestionName string `json:"questionName"`
	Answer       string `json:"answer"`
	ExamId       string `json:"exam_id"`
}
