package main

import (
	"github.com/ISongwuts/go-restful-execution/packages/infrastructure"
)

type (
)

func main() {
	router := infrastructure.NewRouter(":8000", "/api/submit")
	router.Post()
	router.Listen()
}