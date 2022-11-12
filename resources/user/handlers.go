package user

import (
	"github.com/gabrielgaspar447/go-blog-api/constants"
	"github.com/gabrielgaspar447/go-blog-api/models"
	"github.com/gabrielgaspar447/go-blog-api/utils"
	"github.com/gin-gonic/gin"
)

func signupHandler(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(constants.BadRequest, gin.H{"error": utils.GetErrorResponse(err)})
		return
	}

	err := signupService(&input)

	if err != nil {
		switch errMsg := err.Error(); errMsg {
		case constants.UserAlreadyExists:
			c.JSON(constants.Conflict, gin.H{"error": constants.UserAlreadyExists})
			return
		default:
			c.JSON(constants.InternalServerError, gin.H{"error": constants.SomethingWentWrong})
			return
		}
	}

	c.JSON(constants.Created, gin.H{"data": input})
}