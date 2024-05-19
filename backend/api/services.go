package api

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/muling3/go-todos-api/db/sqlc"
)

type todoRequest struct {
	UserID   int32  `json:"user_id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Body     string `json:"body" binding:"required"`
	Due      int    `json:"due" binding:"required"`
	Priority string `json:"priority" binding:"required,oneof=LOW HIGH MEDIUM"`
}

// creating a single todo
func (s *Server) CreateTodo(ctx *gin.Context) {
	var request todoRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dueDate time.Time
	if request.Due <= 0 {
		dueDate = time.Now().Add(time.Hour * 12)
	} else {
		dueDate = time.Now().AddDate(0, 0, request.Due)
	}

	// confirm the user exists
	user, err := s.queries.GetUser(ctx, request.UserID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("USER FOUND " + string(user.ID))

	args := db.CreateTodoParams{
		UserID: sql.NullInt32{
			Int32: user.ID,
		},
		Title:    request.Title,
		Body:     request.Body,
		Priority: request.Priority,
		DueDate: sql.NullTime{
			Time:  dueDate, //time.Now().AddDate(0, 0, request.Due),
			Valid: true,
		},
	}

	if err := s.queries.CreateTodo(ctx, args); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, "Success")
}

// getting a single todo
type idRequest struct {
	Id int `uri:"id" binding:"required,min=1"`
}

func (s *Server) GetToDo(ctx *gin.Context) {
	var getRequest idRequest
	if err := ctx.Copy().ShouldBindUri(&getRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := s.queries.GetTodo(ctx, int32(getRequest.Id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

// getting  all  todoes
func (s *Server) GetToDoes(ctx *gin.Context) {
	todoes, err := s.queries.ListTodos(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todoes)
}

type todoUpdateRequest struct {
	Body     string `json:"body" binding:"required"`
	Priority string `json:"priority" binding:"required,oneof=LOW HIGH MEDIUM"`
}

// updating a todo
func (s *Server) UpdateToDo(ctx *gin.Context) {
	var idReq idRequest
	var request todoUpdateRequest

	if err := ctx.Copy().ShouldBindUri(&idReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	args := db.UpdateTodoParams{
		Body:     request.Body,
		Priority: request.Priority,
		ID:       int32(idReq.Id),
	}

	err := s.queries.UpdateTodo(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}

// delete todo
func (s *Server) DeleteTodo(ctx *gin.Context) {
	var getRequest idRequest
	if err := ctx.Copy().ShouldBindUri(&getRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.queries.DeleteTodo(ctx, int32(getRequest.Id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}

type userRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// creating a single user
func (s *Server) CreateUser(ctx *gin.Context) {
	var request userRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	args := db.CreateUserParams{
		Username: request.Username,
		Password: request.Password,
	}

	if err := s.queries.CreateUser(ctx, args); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, "Success")
}

// getting a single user
func (s *Server) GetUser(ctx *gin.Context) {
	var getRequest idRequest
	if err := ctx.Copy().ShouldBindUri(&getRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.queries.GetUser(ctx, int32(getRequest.Id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// getting  all  users
func (s *Server) GetUsers(ctx *gin.Context) {
	users, err := s.queries.ListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

type userUpdateRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// login user
func (s *Server) LoginUser(ctx *gin.Context) {
	var idReq idRequest
	var request userUpdateRequest

	if err := ctx.Copy().ShouldBindUri(&idReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	args := db.LoginUserParams{
		Username: request.Username,
		Password: request.Password,
	}

	user, err := s.queries.LoginUser(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("USER LOGGED IN " + user.Username)

	ctx.JSON(http.StatusOK, "Login Succesful")
}

// updating a user
func (s *Server) UpdateUser(ctx *gin.Context) {
	var idReq idRequest
	var request userUpdateRequest

	if err := ctx.Copy().ShouldBindUri(&idReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	args := db.UpdateUserParams{
		Username: request.Username,
		Password: request.Password,
		ID:       int32(idReq.Id),
	}

	err := s.queries.UpdateUser(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}

// delete user
func (s *Server) DeleteUser(ctx *gin.Context) {
	var getRequest idRequest
	if err := ctx.Copy().ShouldBindUri(&getRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.queries.DeleteUser(ctx, int32(getRequest.Id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}
