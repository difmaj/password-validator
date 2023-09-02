package routes

import (
	"github.com/difmaj/password-validator-challenge/cmd/server/routes/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	Engine  *gin.Engine
	Handler *handlers.Handler
}

func New() *Routes {

	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	return &Routes{
		Engine:  engine,
		Handler: handlers.New(),
	}
}

func (r *Routes) Register(group string) {
	api := r.Engine.Group(group)
	api.POST("password/verify", r.Handler.Verify)
}
