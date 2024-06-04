package api

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/muling3/go-todos-api/db/sqlc"
	"github.com/o1egl/paseto"
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

	// get set username
	username, exists := ctx.Get("Username")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unathorized"})
		return
	}

	// get the user
	user, err := s.queries.GetUserByUsername(ctx, username.(string))
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

	// get set username
	username, exists := ctx.Get("Username")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unathorized"})
		return
	}

	// get the user
	user, err := s.queries.GetUserByUsername(ctx, username.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	args := db.GetUuserTodoParams{
		UserID: sql.NullInt32{
			Int32: int32(user.ID),
		},
		ID: int32(getRequest.Id),
	}

	todo, err := s.queries.GetUuserTodo(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

// getting  all  todoes
func (s *Server) GetToDoes(ctx *gin.Context) {
	// get set username
	username, exists := ctx.Get("Username")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unathorized"})
		return
	}

	// get the user
	user, err := s.queries.GetUserByUsername(ctx, username.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	todoes, err := s.queries.ListUserTodos(ctx, sql.NullInt32{Valid: true, Int32: int32(user.ID)})
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

	// get set username
	username, exists := ctx.Get("Username")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unathorized"})
		return
	}

	// get the user
	user, err := s.queries.GetUserByUsername(ctx, username.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	args := db.UpdateUserTodoParams{
		Body:     request.Body,
		Priority: request.Priority,
		ID:       int32(idReq.Id),
		UserID:   sql.NullInt32{Int32: int32(user.ID)},
	}

	err = s.queries.UpdateUserTodo(ctx, args)
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

	// get set username
	username, exists := ctx.Get("Username")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unathorized"})
		return
	}

	// get the user
	user, err := s.queries.GetUserByUsername(ctx, username.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	args := db.DeleteUserTodoParams{
		UserID: sql.NullInt32{Int32: int32(user.ID)},
	}

	err = s.queries.DeleteUserTodo(ctx, args)
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
	var request userUpdateRequest

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

	// create the token
	symmetricKey := []byte("YELLOW SUBMARINE, BLACK WIZARDRY") // Must be 32 bytes
	now := time.Now()
	exp := now.Add(24 * time.Hour)
	nbt := now

	jsonToken := paseto.JSONToken{
		Audience:   "todo-frontend",
		Issuer:     "todo.api",
		Jti:        fmt.Sprintf("%s, %v ", args.Username, rand.Intn(20)),
		Subject:    args.Username,
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  nbt,
	}
	// Add custom claim    to the token
	jsonToken.Set("data", "this is a signed message")
	footer := "some footer"

	// Encrypt data
	token, err := paseto.NewV2().Encrypt(symmetricKey, jsonToken, footer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("\n\n USER LOGGED IN %v \n\n", user.Username)

	ctx.JSON(http.StatusOK, gin.H{"token": token, "username": args.Username})
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

	// get set username
	username, exists := ctx.Get("Username")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unathorized"})
		return
	}

	// get the user
	_, err := s.queries.GetUserByUsername(ctx, username.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	args := db.UpdateUserParams{
		Username: request.Username,
		Password: request.Password,
		ID:       int32(idReq.Id),
	}

	err = s.queries.UpdateUser(ctx, args)
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

	// get set username
	username, exists := ctx.Get("Username")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unathorized"})
		return
	}

	// get the user
	_, err := s.queries.GetUserByUsername(ctx, username.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = s.queries.DeleteUser(ctx, int32(getRequest.Id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}
