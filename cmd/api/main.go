// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/cors"
	"time"
)

func main() {

	h := server.New(
		server.WithHostPorts(":8888"),
		server.WithMaxRequestBodySize(1024*1024*1024*5), //最大请求体大小
	)
	hlog.SetLevel(hlog.LevelInfo)

	h.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3001"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"Origin", "content-type", "authorization"},

		ExposeHeaders: []string{"Content-Length", "content-type"}, // Request headers allowed in the upload_file

		AllowCredentials: true, //// 是否允许在跨域的情况下传递Cookie
		MaxAge:           12 * time.Hour,
	}))

	register(h)
	h.Spin()
}
