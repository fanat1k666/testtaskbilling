package repository

type UserTariffs interface {
	ShowTariffs(userId int) ([]ShowUserTariffs, error)
}
type ShowUserTariffs struct {
	UserId int
	Name   string
}
