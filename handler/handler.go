package handler

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/niwho/school_management_demo/model"
)

type Handler struct {
	Db *gorm.DB
}

func (handler *Handler) Init() {
	var err error
	handler.Db, err = gorm.Open("sqlite3", "student.db")
	if err != nil {
		log.Fatal("db connection faild!disater!")
	}
	handler.Db.AutoMigrate(&model.Student{}, &model.Class{})
}

func (handler *Handler) RegisterStudent(id string, classNumber int8, score int8) error {
	// 参数校验
	//

	db := handler.Db.FirstOrCreate(&model.Student, model.Student{ID: id, ClassNumber: classNumber, Score: score})
	if len(db.GetErrors()) > 0 {
		return db.GetErrors()[0]
	}
	return nil
}

func (handler *Handler) RegisterClass(classNumber int8, teacherName string) error {
	return nil
}

func (handler *Handler) GetClassTotalScore(studentId string) error {
	return nil
}

func (handler *Handler) GetTopTeacher() error {
	return nil
}
