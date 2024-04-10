package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xian1367/layout-go/config"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func OpenTelemetry() gin.HandlerFunc {
	return otelgin.Middleware(config.Get().Http[0].Name)
}
