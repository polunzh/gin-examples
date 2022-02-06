package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func hello(c *gin.Context) {
	c.String(http.StatusOK, strings.Join(c.HandlerNames(), "\n"))
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", hello)

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok == true {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	authorize := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":       "bar",
		"zhenqiang": "zhang",
	}))

	authorize.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		var json struct {
			Value string `json:"value" binding:"required"`
		}

		fmt.Println(user)

		bindResult := c.Bind(&json)
		if bindResult == nil {
			db[user] = json.Value
			fmt.Println(json)
			fmt.Println(db["foo"])
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"error": bindResult.Error()})
	})

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
