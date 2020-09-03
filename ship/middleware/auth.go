package middleware

import (
	"github.com/gin-gonic/gin"
	"laya-go/ship"
	"strconv"
)

// implements the server.HandlerWrapper
func Auth(c *gin.Context) {
	token := c.GetHeader("Token")
	uid, err := ship.Redis.Get("user:token:" + token).Result()
	if err != nil {
		c.AbortWithStatusJSON(200, ship.TokenErr)
		return
    }
	ID, _ := strconv.ParseInt(uid, 10, 64)
	c.Set("uid", ID)
	c.Next()
}