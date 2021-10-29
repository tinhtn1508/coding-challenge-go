package v2

import "github.com/gin-gonic/gin"

type IController interface {
	List(c *gin.Context)
	Get(c *gin.Context)
	Post(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
}
