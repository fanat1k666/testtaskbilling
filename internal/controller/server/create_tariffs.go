package server

import (
	"TestTask/internal/entity"
	"encoding/json"
	"io"
	"net/http"
)

type CreateTariffIn struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type CreateTariffOut struct {
	Err *string `json:"err,omitempty"`
}

func (s *Server) HandleCreateTariff(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()
	if err != nil {
		errorStr := err.Error()
		ans := CreateTariffOut{Err: &errorStr}
		s.SendAnswer(w, ans, http.StatusBadGateway)
		return
	}

	in := CreateTariffIn{}
	err = json.Unmarshal(body, &in)
	if err != nil {
		errorStr := err.Error()
		ans := CreateTariffOut{Err: &errorStr}
		s.SendAnswer(w, ans, http.StatusBadRequest)
		return
	}

	eSeg := entity.Tariffs{
		Name:  in.Name,
		Price: in.Price,
	}

	err = s.h.CreateTariff(eSeg)
	if err != nil {
		errorStr := err.Error()
		ans := CreateTariffOut{Err: &errorStr}
		s.SendAnswer(w, ans, http.StatusInternalServerError)
		return
	}

	s.SendAnswer(w, CreateTariffOut{}, http.StatusOK)
}
