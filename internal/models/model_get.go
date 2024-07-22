package models

import "fmt"

type Filter struct {
	PassportSerie  string `schema:"passport_serie,omitempty"`
	PassportNumber string `schema:"passport_number,omitempty"`
	Surname        string `schema:"surname,omitempty"`
	Name           string `schema:"name,omitempty"`
	Patronymic     string `schema:"patronymic,omitempty"`
	Address        string `schema:"address,omitempty"`
	Page           int    `schema:"page"`
	PerPage        int    `schema:"per_page"`
}

func (f *Filter) String() string {
	return fmt.Sprintf("passport_serie = %s, passport_number = %s, name = %s, surname = %s, patronymic = %s, address = %s, page = %d, per_page = %d",
		f.PassportSerie, f.PassportNumber, f.Name, f.Surname, f.Patronymic, f.Address, f.Page, f.PerPage)
}

func NewFilter() *Filter {
	return &Filter{
		PerPage: 10,
	}
}
