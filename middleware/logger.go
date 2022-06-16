package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {

	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)

	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:03"})

	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		latentTime:= endTime.Sub(startTime)

		reqMethod := c.Request.Method

		reqUri := c.Request.RequestURI

		statusCode := c.Writer.Status()

		clientIP := c.ClientIP()

		logger.Infof("| %3d | %13v | %15s | %s | %s |", statusCode, latentTime, clientIP, reqMethod, reqUri)
	}
}