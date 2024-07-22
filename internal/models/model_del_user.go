package models

import "fmt"

type DelUser struct {
	UserId int `json:"user_id"`
}

func (d *DelUser) String() string {
	return fmt.Sprintf("user_id = %d", d.UserId)
}
