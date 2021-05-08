package main

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	validator "gopkg.in/go-playground/validator.v8"
)

type User struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `binding:"required,email"`
}

func UserStructLevelValidation(v *validator.Validate, structLevel *validator.StructLevel) {
	user := structLevel.CurrentStruct.Interface().(User)

	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		structLevel.ReportError(
			reflect.ValueOf(user.FirstName), "FirstName", "fname", "fnameorlname",
		)

		structLevel.ReportError(
			reflect.ValueOf(user.LastName), "LastName", "lname", "fnameorlname",
		)
	}
}

func main() {
	route := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(UserStructLevelValidation, User{})
	}

	route.POST("/user", validateUser)
	route.Run(":8085")
}

func validateUser(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "User validation successful."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User validation failed!",
			"error":   err.Error(),
		})
	}
}
