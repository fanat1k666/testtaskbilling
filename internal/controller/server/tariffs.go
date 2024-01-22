package server

import "net/http"

func (s *Server) HandleTariffs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.HandleCreateTariff(w, r)
	case http.MethodGet:
		s.HandleShowTariff(w, r)
	case http.MethodPut:
		s.HandleUpdateTariff(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
