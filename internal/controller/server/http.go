package server

import (
	"TestTask/internal/controller/middleware"
	"TestTask/internal/handler"
	"TestTask/pkg/log"
	"encoding/json"
	"net/http"
)

type Server struct {
	s *http.Server
	h *handler.Handler
	l log.Logger
}

func NewServer(s *http.Server, h *handler.Handler, l log.Logger) *Server {
	return &Server{
		s: s,
		h: h,
		l: l,
	}
}

func (s *Server) Serve() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/user_tariffs", middleware.AddLogger(middleware.RequestLogger(s.HandleShowTariffs), s.l))
	s.s.Handler = mux

	return s.s.ListenAndServe()
}

func (s *Server) SendAnswer(w http.ResponseWriter, answer interface{}, status int) {
	data, err := json.Marshal(answer)
	if err != nil {
		s.l.Errorf("can't write response: %v", err)
		return
	}
	w.WriteHeader(status)
	_, err = w.Write(data)
	if err != nil {
		s.l.Errorf("can't write response: %v", err)
		return
	}
}
