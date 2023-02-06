package biz

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"bytes"
)

func Handler() gin.HandlerFunc {
	return func(c *gin.Context){
		data,_ := ioutil.ReadAll(c.Request.Body)

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	}
}