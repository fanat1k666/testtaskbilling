package handler

type Handler struct {
	us repository.UserSegment
}

func NewHandler(us repository.UserSegment) *Handler {
	return &Handler{us: us}
}
