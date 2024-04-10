package controller

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/xian1367/layout-go/pkg/gin/response"
	"github.com/xian1367/layout-go/pkg/jwt"
)

type TestController struct {
	BaseAPIController
}

func (ctrl *TestController) Index(c *gin.Context) {
	spew.Dump()

	response.Data(c, gin.H{
		"result": jwt.NewJWT().IssueToken("1"),
	})
}
