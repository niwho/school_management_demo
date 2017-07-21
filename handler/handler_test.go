package handler

import (
	"log"
	"testing"
)

func Test(t *testing.T) {
	HandlerManager.Init()

	err := HandlerManager.RegisterClass(1, "zhanshan")
	log.Printf("registerClass err=%v", err)

	err = HandlerManager.RegisterClass(2, "lisi")
	log.Printf("registerClass err=%v", err)

	err = HandlerManager.RegisterStudent("00001", 1, 90)
	log.Printf("registerStudent err=%v", err)

	err = HandlerManager.RegisterStudent("00002", 1, 90)
	log.Printf("registerStudent err=%v", err)

	err = HandlerManager.RegisterStudent("00004", 1, 82)
	log.Printf("registerStudent err=%v", err)

	total, err1 := HandlerManager.GetClassTotalScore("00001")
	log.Printf("totol=%+v, err1=%+v", total, err1)

	total, err1 = HandlerManager.GetClassTotalScore("00005")
	log.Printf("no exist totol=%+v, err1=%+v", total, err1)

	top, err2 := HandlerManager.GetTopTeacher()
	log.Printf("toptt=%+v, err2=%+v", top, err2)
	if top != "zhanshan" {
		t.Fail()
	}
}
