package server

import (
	"TestTask/internal/handler"
	"encoding/json"
	"log"
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
	//mux.HandleFunc("/segment", middleware.AddLogger(middleware.RequestLogger(s.HandleSegment), s.l))
	//mux.HandleFunc("/user_segment", middleware.AddLogger(middleware.RequestLogger(s.HandleSegmentToUser), s.l))
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
