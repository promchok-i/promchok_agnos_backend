package middleware

import (
	"bytes"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/promchok-i/promchok_agnos_backend/models"
	"gorm.io/gorm"
)

func LogRequestResponseMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody string
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			requestBody = string(bodyBytes)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		responseWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = responseWriter

		c.Next()

		logEntry := models.RequestLog{
			Method:       c.Request.Method,
			Path:         c.Request.URL.Path,
			StatusCode:   c.Writer.Status(),
			RequestBody:  requestBody,
			ResponseBody: responseWriter.body.String(),
		}

		db.Create(&logEntry)
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
