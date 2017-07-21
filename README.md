####部署####
`godep go build`

####test####
注册班级
`curl -H "Content-Type: application/json" -X POST -d '{"classNumber":2,"teacherName":"ssssss"}' http://localhost:8111/register-class`

注册学生
`curl -H "Content-Type: application/json" -X POST -d '{"classNumber":2,"id":"00006", "score": 97}' http://localhost:8111/register-student`

查询总分
`curl -L http://localhost:8111/get-class-total-score/00006/`

查询名师
`curl http://localhost:8111/get-top-teacher`
