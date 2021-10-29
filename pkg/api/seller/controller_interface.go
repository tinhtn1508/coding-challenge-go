package seller

import "github.com/gin-gonic/gin"

type IController interface {
	List(c *gin.Context)
	GetTop10(c *gin.Context)
}
