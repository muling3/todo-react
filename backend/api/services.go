package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/muling3/go-todos-api/db/sqlc"
)

type todoRequest struct {
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

	args := db.CreateTodoParams{
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
		ID: int32(idReq.Id),
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
