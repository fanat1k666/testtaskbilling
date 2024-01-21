package handler

import "TestTask/internal/repository"

type Handler struct {
	us repository.UserTariffs
}

func NewHandler(us repository.UserTariffs) *Handler {
	return &Handler{us: us}
}
