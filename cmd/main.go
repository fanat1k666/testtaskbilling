package main

import (
	"TestTask/internal/controller/server"
	"TestTask/internal/handler"
	"net/http"
)

func main() {
	db, err := postgres.NewPostgres("pgdb", 5432, "postgres", "postgres", "avito")
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
