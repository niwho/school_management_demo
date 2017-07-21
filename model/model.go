package model

type Class struct {
	ClassNumber uint8  `gorm:"primary_key" json:"classNumber"`
	TeacherName string `gorm:"size:"20" json:"teacherName"`
}

type Student struct {
	ID    string `gorm:"size:5;primary_key" json:"id"`
	Score uint8  `gorm:"index:score" json:"score"`
	// Class       Class  `gorm:"ForeignKey:ClassNumber`
	ClassNumber uint8 `json:"classNumber"`
}
