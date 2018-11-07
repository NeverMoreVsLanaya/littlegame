package main

import "github.com/gin-gonic/gin"

func initRouter()*gin.Engine  {
	r:=gin.New()

	r.LoadHTMLGlob()

	return r
}