package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct{}

func ErrMsg(err error, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
}
