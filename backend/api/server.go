package api

import (
	"log"

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

	// router.Use(CORSMiddleware())
	router.Use(cors.New(CORSConfig()))

	router.Use(func(c *gin.Context) {
		log.Printf("Incoming request: %s %s", c.Request.Method, c.Request.URL.Path)
		log.Printf("Request headers: %v", c.Request.Header)
		c.Next()
	})

	// todos
	todos := router.Group("/todos")
	{
		todos.GET("/", server.GetToDoes)
		todos.GET("/:id", server.GetToDo)
		todos.POST("/", server.CreateTodo)
		todos.PUT("/:id", server.UpdateToDo)
		todos.DELETE("/:id", server.DeleteTodo)
	}

	// users
	users := router.Group("/users")
	{
		users.GET("/", server.GetUsers)
		users.GET("/:id", server.GetUser)
		users.POST("/", server.CreateUser)
		users.POST("/login", server.LoginUser)
		users.PUT("/:id", server.UpdateUser)
		users.DELETE("/:id", server.DeleteUser)
	}

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
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Access-Control-Allow-Origin")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		ctx.Next()
	}
}

func CORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE", "OPTIONS")
	return corsConfig
}
