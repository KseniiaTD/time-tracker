package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/KseniiaTD/time-tracker/internal/logger"
	"github.com/KseniiaTD/time-tracker/internal/models"
	"github.com/KseniiaTD/time-tracker/internal/service"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type handler struct {
	srv service.Service
}

// @Summary      List users
// @Description  get users by filter
// @Produce      json
// @Param        name             query   string  false  "user name"
// @Param        surname          query   string  false  "user surname"
// @Param        patronymic       query   string  false  "user patronymic"
// @Param        address          query   string  false  "user address"
// @Param        passport_serie   query   string  false  "user passport serie" example(1111)
// @Param        passport_number  query   string  false  "user passport number" example(222222)
// @Success      200  {array}   models.User
// @Failure      400
// @Failure      500
// @Router       /users [get]
func (h *handler) userListHandler(w http.ResponseWriter, r *http.Request) {
	logger.Logger().Debug("userListHandler called")
	if err := r.ParseForm(); err != nil {
		logger.Logger().Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	filter := models.NewFilter()
	err := schema.NewDecoder().Decode(filter, r.Form)
	if err != nil {
		logger.Logger().Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	users, err := h.srv.GetUserList(*filter)
	if err != nil {
		logger.Logger().Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(users)
	if err != nil {
		logger.Logger().Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Logger().Debugf("Status = OK, data = %s", data)
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

// @Summary      List tasks
// @Description  get tasks by user_id
// @Produce      json
// @Param 		 user_id path int true "user id"
// @Success      200  {array}   models.Task
// @Failure      400
// @Failure      500
// @Router       /users/{user_id}/tasks [get]
func (h *handler) taskListHandler(w http.ResponseWriter, r *http.Request) {
	logger.Logger().Debug("taskListHandler called")
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		logger.Logger().Debug(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	logger.Logger().Debugf("userId = %d", userId)
	if userId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tasks, err := h.srv.GetTaskListByUser(userId)
	if err != nil {
		logger.Logger().Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(tasks)

	if err != nil {
		logger.Logger().Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logger.Logger().Debugf("Status = OK, data = %s", data)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// @Summary      Task start
// @Description  start time counting
// @Accept       json
// @Param        task body models.TaskDate true "date format: 2006-01-02T15:04:05Z"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /task/begin [put]
func (h *handler) taskBeginDateHandler(w http.ResponseWriter, r *http.Request) {
	logger.Logger().Debug("taskBeginDateHandler called")

	var task models.TaskDate
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		logger.Logger().Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Logger().Debug(task)

	taskId := h.srv.SetBeginDateTask(task)
	if taskId == 0 {
		logger.Logger().Debug("Task has already started")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.Logger().Debug("Status = OK")
	w.WriteHeader(http.StatusOK)
}

// @Summary      Task finish
// @Description  stop time counting
// @Accept       json
// @Param        task body models.TaskDate true "date format: 2006-01-02T15:04:05Z"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /task/end [put]
func (h *handler) taskEndDateHandler(w http.ResponseWriter, r *http.Request) {
	logger.Logger().Debug("taskEndDateHandler called")

	var task models.TaskDate
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		logger.Logger().Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Logger().Debug(task)

	taskId := h.srv.SetEndDateTask(task)
	if taskId == 0 {
		logger.Logger().Debug("Task has already finished")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.Logger().Debug("Status = OK")
	w.WriteHeader(http.StatusOK)
}

// @Summary      User delete
// @Description  delete user by id
// @Accept       json
// @Param        user body models.DelUser true "json"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /user [delete]
func (h *handler) delUserHandler(w http.ResponseWriter, r *http.Request) {
	logger.Logger().Debug("delUserHandler called")

	var userId models.DelUser
	err := json.NewDecoder(r.Body).Decode(&userId)
	if err != nil {
		logger.Logger().Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Logger().Debug(userId)

	delUserId := h.srv.DeleteUser(userId)
	if delUserId == 0 {
		logger.Logger().Debug("User not found")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.Logger().Debug("Status = OK")
	w.WriteHeader(http.StatusOK)
}

// @Summary      User update
// @Description  update user by id
// @Accept       json
// @Param        user body models.UpdUser true "json"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /user [put]
func (h *handler) updUserHandler(w http.ResponseWriter, r *http.Request) {
	logger.Logger().Debug("updUserHandler called")

	var user models.UpdUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Logger().Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Logger().Debug(user)

	if user.UserId == 0 || len(user.Name) == 0 || len(user.Address) == 0 || len(user.Surname) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userId := h.srv.UpdateUser(user)
	if userId == 0 {
		logger.Logger().Debug("User not found")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.Logger().Debug("Status = OK")
	w.WriteHeader(http.StatusOK)
}

// @Summary      User create
// @Description  create user by passport info
// @Accept       json
// @Produce      json
// @Param        user body models.UserPassport true "json"
// @Success      201 {object} models.User
// @Failure      400
// @Failure      500
// @Router       /user [post]
func (h *handler) newUserHandler(w http.ResponseWriter, r *http.Request) {
	logger.Logger().Debug("newUserHandler called")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Logger().Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Logger().Debug(user)

	if len(user.PassportSerie) == 0 ||
		len(user.PassportNumber) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userList, err := h.srv.GetUserList(models.Filter{
		PassportSerie:  user.PassportSerie,
		PassportNumber: user.PassportNumber,
		Page:           1,
		PerPage:        1})
	if err != nil {
		logger.Logger().Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(userList) != 0 {
		logger.Logger().Debug("User has already existed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	port := os.Getenv("EXTERNAL_SERVICE_PORT")
	logger.Logger().Debug("Get data from swagger")
	resp, err := http.Get("http://localhost:" + port + "/info")
	if err != nil {
		logger.Logger().Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		logger.Logger().Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Logger().Debug(user)

	userId, err := h.srv.CreateUser(user)
	if err != nil {
		logger.Logger().Debug(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logger.Logger().Debugf("Status = OK, data = %d", userId)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{"user_id": %d}`, userId)))
}

func New(srv service.Service) *mux.Router {
	router := mux.NewRouter()

	h := handler{srv: srv}
	router.HandleFunc("/users", h.userListHandler).Methods(http.MethodGet)
	router.HandleFunc("/users/{user_id}/tasks", h.taskListHandler).Methods(http.MethodGet)
	router.HandleFunc("/task/begin", h.taskBeginDateHandler).Methods(http.MethodPut)
	router.HandleFunc("/task/end", h.taskEndDateHandler).Methods(http.MethodPut)
	router.HandleFunc("/user", h.delUserHandler).Methods(http.MethodDelete)
	router.HandleFunc("/user", h.updUserHandler).Methods(http.MethodPut)
	router.HandleFunc("/user", h.newUserHandler).Methods(http.MethodPost)

	return router
}
