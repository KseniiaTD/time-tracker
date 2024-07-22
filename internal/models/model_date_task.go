package models

import (
	"fmt"
	"time"
)

type TaskDate struct {
	TaskId int       `json:"id"`
	Date   time.Time `json:"date"`
}

func (t *TaskDate) String() string {
	return fmt.Sprintf("task_id = %d, date = %s", t.TaskId, t.Date.String())
}
