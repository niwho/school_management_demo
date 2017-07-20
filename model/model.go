package model

type Student struct{
	ID string `gorm:"size:5;primary_key"`
	Score int8 `gorm:"index:score"`
    Class Class `gorm:"ForeignKey:ClassNumber`
    ClassNumber string
}

type Class struct {
	ClassNumber int8 `gorm:"primary_key"`
	TeacherName string `gorm:"size:"20"`
}

