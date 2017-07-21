package views

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/niwho/school_management_demo/handler"
	"github.com/niwho/school_management_demo/model"
	// "github.com/niwho/school_management_demo/utils"
)

func RegisterStudent(c *gin.Context) {
	var st model.Student
	c.BindJSON(&st)
	log.Printf("%+v", st)
	status := 0
	message := "ok"
	if err := handler.HandlerManager.RegisterStudentV2(st); err != nil {
		status = -1
		message = err.Error()
	}
	c.JSON(200, gin.H{
		"status":  status,
		"message": message,
	})
}

func RegisterClass(c *gin.Context) {
	var cl model.Class
	c.BindJSON(&cl)
	log.Printf("%+v", cl)
	status := 0
	message := "ok"
	if err := handler.HandlerManager.RegisterClassV2(cl); err != nil {
		status = -1
		message = err.Error()
	}
	c.JSON(200, gin.H{
		"status":  status,
		"message": message,
	})
}

func GetClassTotalScore(c *gin.Context) {
	mt := handler.FiveDigitPatern.FindStringSubmatch(c.Param("studentId"))
	rt := gin.H{
		"error": "student-not-found",
	}
	if len(mt) > 0 {
		score, err := handler.HandlerManager.GetClassTotalScore(mt[0])
		if err == nil {
			rt = gin.H{
				"total": score,
			}
		}
	}
	c.JSON(200, rt)
}

func GetTopTeacher(c *gin.Context) {
	name, err := handler.HandlerManager.GetTopTeacher()
	if err != nil {
		name = err.Error()
	}
	c.JSON(200, gin.H{
		"teacher": name,
	})
}
