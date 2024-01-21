package server

import (
	"TestTask/internal/entity"
	"encoding/json"
	"io"
	"net/http"
)

type ShowTariffsIn struct {
	UserId int `json:"user_id"`
}

type ShowTariffsOut struct {
	UserId  int      `json:"userId"`
	Tariffs []string `json:"tariffs"`
	Err     *string  `json:"err,omitempty"`
}

func (s *Server) HandleShowTariffs(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()
	if err != nil {
		errorStr := err.Error()
		ans := ShowTariffsOut{Err: &errorStr}
		s.SendAnswer(w, ans, http.StatusBadGateway)
		return
	}

	in := ShowTariffsIn{}
	err = json.Unmarshal(body, &in)
	if err != nil {
		errorStr := err.Error()
		ans := ShowTariffsOut{Err: &errorStr}
		s.SendAnswer(w, ans, http.StatusBadRequest)
		return
	}

	eUs := entity.User{
		UserId: in.UserId,
	}

	rows, err := s.h.ShowTariffs(eUs)
	if err != nil {
		errorStr := err.Error()
		ans := ShowTariffsOut{Err: &errorStr}
		s.SendAnswer(w, ans, http.StatusInternalServerError)
		return
	}
	var end []string
	var u int
	if len(rows) != 0 {
		u = rows[0].UserId
		for i := range rows {
			end = append(end, rows[i].Name)
		}
	}
	s.SendAnswer(w, ShowTariffsOut{UserId: u, Tariffs: end}, http.StatusOK)
}
