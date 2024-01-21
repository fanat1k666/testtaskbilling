package handler

import (
	"TestTask/internal/entity"
	"TestTask/internal/repository"
	"fmt"
)

func (h *Handler) ShowTariffs(user entity.User) ([]repository.ShowUserTariffs, error) {
	line, err := h.us.ShowTariffs(user.UserId)
	if err != nil {
		return nil, fmt.Errorf("can't show tariffs: %w", err)
	}

	return line, nil
}
