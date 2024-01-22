package handler

import (
	"TestTask/internal/entity"
	"fmt"
)

func (h *Handler) UpdateTariff(tariff entity.Tariffs) error {
	err := h.us.UpdateTariff(tariff.Name, tariff.Price)
	if err != nil {
		return fmt.Errorf("can't update tariff: %w", err)
	}
	return nil
}
