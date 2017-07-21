package handler

import (
	"errors"
	"log"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/niwho/school_management_demo/model"
)

var HandlerManager Handler

type Handler struct {
	Db *gorm.DB
}

type ClassTotalScore struct {
	Total int64
}

var FiveDigitPatern *regexp.Regexp = regexp.MustCompile("^[0-9]{5,5}$")

// 注意此函数在程序启动是初始化，不可重复调用
func (handler *Handler) Init() {
	var err error
	handler.Db, err = gorm.Open("sqlite3", "student.db")
	if err != nil {
		log.Fatal("db connection faild!disater!")
	}
	handler.Db.AutoMigrate(&model.Student{}, &model.Class{})
}

func (handler *Handler) RegisterStudent(id string, classNumber uint8, score uint8) error {
	// 参数校验
	if !FiveDigitPatern.MatchString(id) || classNumber > 99 || score > 100 {
		return errors.New("params invalid")
	}
	//

	// 同步问题交给数据库， 高并发插入，增加异步缓存
	db := handler.Db.FirstOrCreate(&model.Student{}, model.Student{ID: id, ClassNumber: classNumber, Score: score})
	if len(db.GetErrors()) > 0 {
		return db.GetErrors()[0]
	}
	return nil
}

func (handler *Handler) RegisterStudentV2(st model.Student) error {
	// 参数校验
	if !FiveDigitPatern.MatchString(st.ID) || st.ClassNumber > 99 || st.Score > 100 {
		return errors.New("params invalid")
	}
	//没有外键约束，业务代码保证
	var count int
	db := handler.Db.Model(&model.Class{}).Where(model.Class{ClassNumber: st.ClassNumber}).Count(&count)
	if len(db.GetErrors()) > 0 {
		return db.GetErrors()[0]
	}
	if count == 0 {
		return errors.New("no such class")
	}

	// 同步问题交给数据库， 高并发插入，增加异步缓存
	db = handler.Db.Where(model.Student{ID: st.ID}).Assign(st).FirstOrCreate(&st)
	if len(db.GetErrors()) > 0 {
		return db.GetErrors()[0]
	}
	return nil
}

func (handler *Handler) RegisterClass(classNumber uint8, teacherName string) error {
	// 参数校验
	if classNumber > 99 || len(teacherName) > 20 || strings.TrimSpace(teacherName) == "" {
		return errors.New("params invalid")
	}
	// 同步问题交给数据库， 高并发插入，增加异步缓存
	db := handler.Db.FirstOrCreate(&model.Class{}, model.Class{ClassNumber: classNumber, TeacherName: teacherName})
	if len(db.GetErrors()) > 0 {
		return db.GetErrors()[0]
	}
	return nil
}

func (handler *Handler) RegisterClassV2(cl model.Class) error {
	// 参数校验
	if cl.ClassNumber > 99 || len(cl.TeacherName) > 20 || strings.TrimSpace(cl.TeacherName) == "" {
		return errors.New("params invalid")
	}
	// 同步问题交给数据库， 高并发插入，增加异步缓存
	db := handler.Db.Where(model.Class{ClassNumber: cl.ClassNumber}).Assign(cl).FirstOrCreate(&cl)

	if len(db.GetErrors()) > 0 {
		return db.GetErrors()[0]
	}
	return nil
}

func (handler *Handler) GetClassTotalScore(studentId string) (int64, error) {
	// 参数校验
	if !FiveDigitPatern.MatchString(studentId) {
		return 0, errors.New("params invalid")
	}

	// join query be more effective?
	// 获取到class
	// class := model.Class{}
	// db := handler.Db.Model(&model.Student{ID: studentId}).Related(&class)
	student := model.Student{}
	db := handler.Db.Where(&model.Student{ID: studentId}).First(&student)
	if len(db.GetErrors()) > 0 {
		return 0, db.GetErrors()[0]
	}
	cs := ClassTotalScore{}
	db = handler.Db.Table("students").Select("sum(score) as total").Group("class_number").Where("class_number=?", student.ClassNumber).First(&cs)
	if len(db.GetErrors()) > 0 {
		return 0, db.GetErrors()[0]
	}
	return cs.Total, nil
}

func (handler *Handler) GetTopTeacher() (string, error) {
	st := model.Student{}
	db := handler.Db.Order("score desc, id").First(&st)
	if len(db.GetErrors()) > 0 {
		return "", db.GetErrors()[0]
	}

	// join query be more effective?
	class := model.Class{}
	// db = handler.Db.Model(&st).Related(&class)
	db = handler.Db.Where(&model.Class{ClassNumber: st.ClassNumber}).First(&class)
	if len(db.GetErrors()) > 0 {
		return "", db.GetErrors()[0]
	}
	return class.TeacherName, nil
}
