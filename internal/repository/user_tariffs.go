package repository

type UserTariffs interface {
	ShowTariff(userId int) ([]ShowUserTariff, error)
	CreateTariff(name string, price int) error
}
type ShowUserTariff struct {
	UserId int
	Name   string
}
