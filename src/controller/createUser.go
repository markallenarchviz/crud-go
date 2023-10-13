package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/markallenarchviz/crud-go/src/configuration/rest_err"
	request "github.com/markallenarchviz/crud-go/src/controller/model/resquest"
)

func CreateUser(c *gin.Context) {
	var userResquest request.UserRequest

	if err := c.ShouldBindJSON(&userResquest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorret fields, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userResquest)
}
