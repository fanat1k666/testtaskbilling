package handler

import (
	"TestTask/internal/entity"
	"TestTask/internal/repository"
	"fmt"
)

func (h *Handler) ShowTariff(user entity.User) ([]repository.ShowUserTariff, error) {
	line, err := h.us.ShowTariff(user.UserId)
	if err != nil {
		return nil, fmt.Errorf("can't show user's tariff: %w", err)
	}
	return line, nil
}
