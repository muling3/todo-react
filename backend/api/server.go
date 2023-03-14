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

	router.GET("/", server.GetToDoes)
	router.GET("/:id", server.GetToDo)
	router.POST("/", server.CreateTodo)
	router.PUT("/:id", server.UpdateToDo)
	router.DELETE("/:id", server.DeleteTodo)

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
