package api

import (
	"github.com/gin-gonic/gin"
)

type WebRouter struct{}

type ProxyReq struct {
	DestUrl string            `json:"dest_url" binding:"required"`
	Header  map[string]string `json:"header" binding:"required"`

	Body    map[string]interface{} `json:"body"`
	ReqType int                    `json:"req_type"` //0:get 1:post
	Index   int                    `json:"index"`
}

func (c WebRouter) DoProxy(ctx *gin.Context) {
	var req ProxyReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
}
