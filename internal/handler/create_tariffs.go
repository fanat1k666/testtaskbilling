package handler

import (
	"TestTask/internal/entity"
	"fmt"
)

func (h *Handler) CreateTariff(tariff entity.Tariffs) error {
	err := h.us.CreateTariff(tariff.Name, tariff.Price)
	if err != nil {
		return fmt.Errorf("can't create tariff: %w", err)
	}
	return nil
}
