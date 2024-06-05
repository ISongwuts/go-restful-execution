package infrastructure

import (
	"net/http"

	"github.com/ISongwuts/go-restful-execution/packages/domain"
	"github.com/ISongwuts/go-restful-execution/packages/usecase"
	"github.com/gin-gonic/gin"
)

type (
	Router struct {
		*gin.Engine
		Port string
		Path string
	}

	RouterImpl interface {
		Post()
		Listen()
	}
)

func NewRouter(port, path string) RouterImpl {
	return &Router{
		Engine: gin.Default(),
		Port:   port,
		Path:   path,
	}
}

func (r *Router) Post() {
	r.POST(r.Path, func(ctx *gin.Context) {
		var data domain.Submit

		if err := ctx.BindJSON(&data); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		submit := usecase.NewSubmit(data.Language, data.Code, data.SubmitAt, data.Status)
		output, status, err := submit.TestCase()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"output": output,
			"status": status,
		})

	})
}

func (r *Router) Listen() {
	r.Run(r.Port)
}
