package taskService

import (
	"log"
	"net/http"
	dbconn "task-management/server/db"
	dl "task-management/server/taskservice/dl"

	"github.com/go-resty/resty/v2"
)

type taskService struct {
}

func NewTaskService() *taskService {
	return &taskService{}
}

func (cs *taskService) InitTaskServiceHandler() http.Handler {
	CustomerDB := "taskservice"
	client := resty.New()

	client.BaseURL = "http://localhost:8080"
	Db, err := dbconn.NewDB(CustomerDB)
	if err != nil {
		return nil
	}

	Dl := dl.NewTaskDL(Db)
	Bl := NewTaskBL(Dl)

	TaskEps, err := NewTaskServiceEndPoints(Bl)
	if err != nil {
		log.Printf("Failed to create taskservice endpoints %s", err.Error())
		return nil
	}

	HttpHandler := MakeHandler(TaskEps)
	return HttpHandler
}
