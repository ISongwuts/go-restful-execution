package infrastructure

import (
	"fmt"
	"io"
	"log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type (
	Router struct {
		gin.Engine
		Port string
		Path string
	}

	RouterImpl interface {
		NewRouter(port, path string) RouterImpl
		Post()
	}
)

func (r *Router) NewRouter(port, path string) RouterImpl {
	r.Use(cors.Default())
	return &Router{
		Port: port,
		Path: path,
	}
}

func (r *Router) Post() {
	r.POST(r.Path, func(ctx *gin.Context) {
		data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			log.Fatal("Err:", err)
		}
		fmt.Println(data)

	})
}