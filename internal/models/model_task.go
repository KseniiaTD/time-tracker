package models

type Task struct {
	TaskId  int    `json:"task_id"`
	Name    string `json:"name"`
	Hours   int    `json:"hours"`
	Minutes int    `json:"minutes"`
}
