package routes

import (
	"github.com/difmaj/password-validator/cmd/server/routes/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Routes is the struct that defines the routes.
type Routes struct {
	Engine  *gin.Engine
	Handler *handlers.PasswordHandler
}

// New returns a new Routes instance.
// It also sets up the Gin engine and the CORS middleware.
func New() *Routes {

	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	return &Routes{
		Engine:  engine,
		Handler: handlers.NewPasswordHandler(),
	}
}

// Register registers all routes.
func (r *Routes) Register(group string) {
	api := r.Engine.Group(group)
	api.POST("password/verify", r.Handler.Verify)
}
