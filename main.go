package main

import (
	"github.com/gin-gonic/gin"
	"github.com/niwho/school_management_demo/handler"
	"github.com/niwho/school_management_demo/views"
)

func main() {
	//init
	handler.HandlerManager.Init()
	r := gin.Default()
	r.POST("/register-student", views.RegisterStudent)
	r.POST("/register-class", views.RegisterClass)
	r.GET("/get-class-total-score/:studentId", views.GetClassTotalScore)
	r.GET("/get-top-teacher", views.GetTopTeacher)
	r.Run(":8111") // listen and serve on 0.0.0.0:8080
}
