package cors

import (
	"github.com/cloudwego/hertz/pkg/app"
	hcors "github.com/hertz-contrib/cors"
	"time"
)

func InitCors() app.HandlerFunc {
	config := hcors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	handlerFunc := hcors.New(config)
	return handlerFunc
}
