package main

import (
	"fmt"
	"github.com/Abdur-Rohman/exam_project/api"
	"github.com/Abdur-Rohman/exam_project/config"
	"github.com/Abdur-Rohman/exam_project/pkg/logger"
	"github.com/Abdur-Rohman/exam_project/service"
	"github.com/Abdur-Rohman/exam_project/storage/postgres"
	"net/http"
)

func main() {
	l := logger.New()

	c := api.NewController(service.New(postgres.New(l)))
	e := api.Server(c)
	l.Info(fmt.Sprintf("running on port: %d and host: %s", config.Load().Port, config.Load().Host))
	http.ListenAndServe("localhost:8080", e)
}
