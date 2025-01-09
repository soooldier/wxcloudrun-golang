package main

import (
	"fmt"
	"io/ioutil"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}
	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/", func(c *gin.Context) {
		b, err := ioutil.ReadFile("./index.html")
		if err != nil {
			c.HTML(500, "Internal Error", nil)
			c.Abort()
		}
		c.HTML(200, string(b), nil)
	})
	r.GET("/MP_verify_32iWga2EVle6QTQm.txt", func(c *gin.Context) {
		c.HTML(200, string("32iWga2EVle6QTQm"), nil)
	})
	r.GET("/api/callback", service.LoginHandler)
	r.Run(":80")
}
