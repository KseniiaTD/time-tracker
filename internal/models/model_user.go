package models

import "fmt"

type UserPassport struct {
	PassportSerie  string `json:"passport_serie"`
	PassportNumber string `json:"passport_number"`
}

type User struct {
	UserPassport
	UserId     int    `json:"user_id"`
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic,omitempty"`
	Address    string `json:"address"`
}

func (u *User) String() string {
	return fmt.Sprintf("user_id = %d, passport_serie = %s, passport_number = %s, surname = %s, name = %s, patronymic = %s, address = %s",
		u.UserId, u.PassportSerie, u.PassportNumber, u.Name, u.Surname, u.Patronymic, u.Address)
}
