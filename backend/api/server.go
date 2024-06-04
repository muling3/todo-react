package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/muling3/go-todos-api/db/sqlc"
	"github.com/o1egl/paseto"
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

	// authentication middleware
	// router.Use(AuthTokenMiddleware())

	router.Use(func(c *gin.Context) {
		log.Printf("Incoming request: %s %s", c.Request.Method, c.Request.URL.Path)
		log.Printf("Request headers: %v", c.Request.Header)
		c.Next()
	})

	// users
	users := router.Group("/users")
	{
		users.GET("/", AuthTokenMiddleware(), server.GetUsers)
		users.GET("/:id", AuthTokenMiddleware(), server.GetUser)
		users.POST("/", server.CreateUser)
		users.POST("/login", server.LoginUser)
		users.PUT("/:id", AuthTokenMiddleware(), server.UpdateUser)
		users.DELETE("/:id", AuthTokenMiddleware(), server.DeleteUser)
	}

	// todos
	todos := router.Group("/todos")
	{
		todos.Use(AuthTokenMiddleware())
		todos.GET("/", server.GetToDoes)
		todos.GET("/:id", server.GetToDo)
		todos.POST("/", server.CreateTodo)
		todos.PUT("/:id", server.UpdateToDo)
		todos.DELETE("/:id", server.DeleteTodo)
	}

	server.router = router

	return server
}

func (s *Server) StartServer(adr string) error {
	return s.router.Run(adr)
}

func AuthTokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get the token from header
		var AuthHeader = ctx.GetHeader("Authorization")
		log.Println("AUTH HEADER " + AuthHeader)

		if strings.EqualFold(AuthHeader, "") {
			// throw unauthorized
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Not Authorized"})
			return
		}

		token := strings.Split(AuthHeader, "Bearer ")
		log.Println("TOKEN IS " + token[1])

		// // Decrypt data
		symmetricKey := []byte("YELLOW SUBMARINE, BLACK WIZARDRY") // Must be 32 bytes
		var newJsonToken paseto.JSONToken
		var newFooter string
		err := paseto.NewV2().Decrypt(token[1], symmetricKey, &newJsonToken, &newFooter)

		if err != nil {
			// throw unauthorised exception
			log.Println("Error verifyng token " + err.Error())
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// add the user id in the request
		ctx.Set("Username", newJsonToken.Subject)

		ctx.Next()
	}
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
