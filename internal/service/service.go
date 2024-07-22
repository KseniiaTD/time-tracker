package service

import (
	"context"

	"github.com/KseniiaTD/time-tracker/internal/database"
	"github.com/KseniiaTD/time-tracker/internal/logger"
	"github.com/KseniiaTD/time-tracker/internal/models"
)

type Service interface {
	GetUserList(filter models.Filter) (users []models.User, err error)
	GetTaskListByUser(userId int) (tasks []models.Task, err error)
	SetBeginDateTask(taskInfo models.TaskDate) (startedTaskId int)
	SetEndDateTask(taskInfo models.TaskDate) (finishedTaskId int)
	DeleteUser(userId models.DelUser) (delUserId int)
	UpdateUser(user models.UpdUser) (udpUserId int)
	CreateUser(user models.User) (userId int, err error)
}

type srv struct {
	db  database.Database
	ctx context.Context
}

func New(db database.Database, ctx context.Context) Service {
	return &srv{db: db, ctx: ctx}
}

func (s *srv) GetUserList(filter models.Filter) ([]models.User, error) {
	logger.Logger().Debug("GetUserList called")
	logger.Logger().Debug(filter)
	if filter.Page == 0 {
		filter.Page = 1
	}

	rows, err := s.db.DB().QueryContext(s.ctx,
		`
	with data as (
	select u.user_id,
		   u.name,
		   u.surname,
		   u.patronymic,
		   u.passport_serie,
		   u.passport_number,
		   u.address,
		   row_number() over (order by u.surname, u.name, u.patronymic, u.passport_serie, u.passport_number) rn
	from public.users u
	where (u.name ilike '%'||$1||'%' or $1 is null)
	      and (u.surname ilike '%'||$2||'%' or $2 is null)
		  and (u.patronymic ilike '%'||$3||'%' or $3 is null)
		  and (u.passport_serie ilike '%'||$4||'%' or $4 is null)
		  and (u.passport_number ilike '%'||$5||'%' or $5 is null)
		  and (u.address ilike '%'||$6||'%' or $6 is null)
		  and is_deleted = false
)
	select u.user_id,
		   u.name,
		   u.surname,
		   u.patronymic,
		   u.passport_serie,
		   u.passport_number,
		   u.address
	from data u
	offset ($7 - 1)*$8 LIMIT $8
	`,
		filter.Name,
		filter.Surname,
		filter.Patronymic,
		filter.PassportSerie,
		filter.PassportNumber,
		filter.Address,
		filter.Page,
		filter.PerPage,
	)
	if err != nil {
		logger.Logger().Debug(err)
		return nil, err
	}
	defer rows.Close()
	users := make([]models.User, 0)
	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.UserId, &user.Name, &user.Surname, &user.Patronymic, &user.PassportSerie, &user.PassportNumber, &user.Address)
		if err != nil {
			logger.Logger().Debug(err)
			return nil, err
		}
		logger.Logger().Debug(user)
		users = append(users, user)
	}

	logger.Logger().Debug(users)
	return users, nil
}

func (s *srv) GetTaskListByUser(userId int) ([]models.Task, error) {
	logger.Logger().Debug("GetTaskListByUser called")
	logger.Logger().Debugf("user_id = %d", userId)

	rows, err := s.db.DB().QueryContext(s.ctx,
		`
	with basis as (
		select t.task_id,
			t.task_name,
			DATE_PART('day', COALESCE(date_end, now()) - COALESCE(date_begin, now())) * 24 + 
				DATE_PART('hour', COALESCE(date_end, now()) - COALESCE(date_begin, now())) hours,
			DATE_PART('minute', COALESCE(date_end, now()) - COALESCE(date_begin, now())) minutes
		from public.users u
		inner join public.tasks t on u.user_id = t.user_id
		where u.user_id = $1 and u.is_deleted = false)
	select t.task_id,
	   t.task_name,
	   t.hours,
	   t.minutes
	from basis t
	order by hours desc, minutes desc
	`,
		userId)
	if err != nil {
		logger.Logger().Debug(err)
		return nil, err
	}
	defer rows.Close()
	tasks := make([]models.Task, 0)
	for rows.Next() {
		task := models.Task{}
		err = rows.Scan(&task.TaskId, &task.Name, &task.Hours, &task.Minutes)
		if err != nil {
			logger.Logger().Debug(err)
			return nil, err
		}
		logger.Logger().Debug(task)
		tasks = append(tasks, task)
	}

	logger.Logger().Debug(tasks)
	return tasks, nil
}

func (s *srv) SetBeginDateTask(taskInfo models.TaskDate) int {
	logger.Logger().Debug("SetBeginDateTask called")
	logger.Logger().Debug(taskInfo)

	row := s.db.DB().QueryRowContext(s.ctx,
		`
	update public.tasks 
	set date_begin = $1
	where task_id = $2 and date_begin is null 
	returning task_id`,
		taskInfo.Date,
		taskInfo.TaskId)

	var startedTaskId int
	row.Scan(&startedTaskId)

	logger.Logger().Debugf("Task_id = %d", startedTaskId)
	return startedTaskId
}

func (s *srv) SetEndDateTask(taskInfo models.TaskDate) int {
	logger.Logger().Debug("SetEndDateTask called")
	logger.Logger().Debug(taskInfo)

	row := s.db.DB().QueryRowContext(s.ctx,
		`
	update public.tasks 
	set date_end = $1
	where task_id = $2 and date_end is null
	returning task_id`,
		taskInfo.Date,
		taskInfo.TaskId)

	var finishedTaskId int
	row.Scan(&finishedTaskId)

	logger.Logger().Debugf("Task_id = %d", finishedTaskId)
	return finishedTaskId

}

func (s *srv) DeleteUser(userId models.DelUser) int {
	logger.Logger().Debug("DeleteUser called")
	logger.Logger().Debug(userId)

	row := s.db.DB().QueryRowContext(s.ctx,
		`
	update public.users 
	set is_deleted = true
	where user_id = $1
		and is_deleted = false
	returning user_id`,
		userId.UserId)

	var delUserId int
	row.Scan(&delUserId)

	logger.Logger().Debugf("User_id = %d", delUserId)
	return delUserId
}

func (s *srv) UpdateUser(user models.UpdUser) int {
	logger.Logger().Debug("UpdateUser called")
	logger.Logger().Debug(user)

	row := s.db.DB().QueryRowContext(s.ctx,
		`
		update public.users 
		set name = $1,
			surname = $2,
			patronymic = $3,
			address = $4
		where user_id = $5
		  and is_deleted = false
		returning user_id`,
		user.Name,
		user.Surname,
		user.Patronymic,
		user.Address,
		user.UserId,
	)

	var updUserId int
	row.Scan(&updUserId)

	logger.Logger().Debugf("User_id = %d", updUserId)
	return updUserId
}

func (s *srv) CreateUser(user models.User) (int, error) {
	logger.Logger().Debug("CreateUser called")
	logger.Logger().Debug(user)

	row := s.db.DB().QueryRowContext(s.ctx,
		`
		insert into public.users(name, surname, patronymic, passport_serie, passport_number, address)
		values ($1, $2, $3, $4, $5, $6)
		returning user_id`,
		user.Name,
		user.Surname,
		user.Patronymic,
		user.PassportSerie,
		user.PassportNumber,
		user.Address)

	var userId int
	err := row.Scan(&userId)

	logger.Logger().Debug(err)
	logger.Logger().Debugf("userId = %d", userId)
	return userId, err
}
