package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

type BodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w BodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Response() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := &BodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = w
		c.Next()
	}
}
