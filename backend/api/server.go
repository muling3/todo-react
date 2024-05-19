package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/muling3/go-todos-api/db/sqlc"
)

type Server struct {
	queries *db.Queries
	router  *gin.Engine
}

func NewServer(q *db.Queries) *Server {
	server := &Server{queries: q}
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// todos
	todos := router.Group("/todos")

	todos.GET("/", server.GetToDoes)
	todos.GET("/:id", server.GetToDo)
	todos.POST("/", server.CreateTodo)
	todos.PUT("/:id", server.UpdateToDo)
	todos.DELETE("/:id", server.DeleteTodo)

	// users
	users := router.Group("/users")

	users.GET("/", server.GetUsers)
	users.GET("/:id", server.GetUser)
	users.POST("/", server.CreateUser)
	users.POST("/login", server.LoginUser)
	users.PUT("/:id", server.UpdateUser)
	users.DELETE("/:id", server.DeleteUser)

	server.router = router

	return server
}

func (s *Server) StartServer(adr string) error {
	return s.router.Run(adr)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		ctx.Next()
	}
}
