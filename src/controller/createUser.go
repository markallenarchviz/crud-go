package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markallenarchviz/crud-go/src/configuration/rest_err"
	request "github.com/markallenarchviz/crud-go/src/controller/model/resquest"
)

func CreateUser(c *gin.Context) {
	var userResquest request.UserRequest

	if err := c.ShouldBindJSON(&userResquest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorret fields"))

		c.JSON(restErr.Code, restErr)
		return
	}
	c.JSON(http.StatusCreated, userResquest)
	fmt.Println(userResquest)
}
