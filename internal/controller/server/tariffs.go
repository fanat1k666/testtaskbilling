package server

import "net/http"

func (s *Server) HandleTariffs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.HandleShowTariffs(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
