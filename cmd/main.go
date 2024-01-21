package main

import (
	"TestTask/internal/controller/server"
	"TestTask/internal/handler"
	"TestTask/internal/postgres"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	db, err := postgres.NewPostgres("192.168.56.101", 5432, "postgres", "password", "testdb")
	if err != nil {
		panic(err)
	}

	h := handler.NewHandler(db)
	l := logrus.New()
	s := &http.Server{Addr: ":8000"}
	c := server.NewServer(s, h, l)

	err = c.Serve()
	if err != nil {
		panic(err)
	}
}
