package controllers

import "github.com/gin-gonic/gin"

type ProxyController struct{}

func (proxy *ProxyController) HandleTraffic(c *gin.Context) {
	c.String(200, "TODO")
}
