package models

import "fmt"

type UpdUser struct {
	UserId     int    `json:"user_id"`
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic,omitempty"`
	Address    string `json:"address"`
}

func (u *UpdUser) String() string {
	return fmt.Sprintf("user_id = %d, name = %s, surname = %s, patronymic = %s, address = %s",
		u.UserId, u.Name, u.Surname, u.Patronymic, u.Address)
}
