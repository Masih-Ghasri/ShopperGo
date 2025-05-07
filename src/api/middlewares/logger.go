package middlewares

import (
	"bytes"
	"github.com/Masih-Ghasri/GolangBackend/config"
	"github.com/Masih-Ghasri/GolangBackend/pkg/logging"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func DefualtStructuredLogger(cfg *config.Config) gin.HandlerFunc {
	logger := logging.NewLogger(cfg)
	return StructuredLogger(logger)
}

func StructuredLogger(logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		start := time.Now()
		path := c.FullPath()
		raw := c.Request.URL.RawQuery

		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body.Close()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		c.Writer = blw
		c.Next()

		params := gin.LogFormatterParams{}
		params.TimeStamp = time.Now()
		params.Latency = params.TimeStamp.Sub(start)
		params.ClientIP = c.ClientIP()
		params.Method = c.Request.Method
		params.StatusCode = c.Writer.Status()
		params.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		params.BodySize = c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}
		params.Path = path

		keys := map[logging.ExtraKey]interface{}{}
		keys[logging.Path] = params.Path
		keys[logging.ClientIp] = params.ClientIP
		keys[logging.Method] = params.Method
		keys[logging.Latency] = params.Latency
		keys[logging.StatusCode] = params.StatusCode
		keys[logging.ErrorMessage] = params.ErrorMessage
		keys[logging.BodySize] = params.BodySize
		keys[logging.RequestBody] = string(bodyBytes)
		keys[logging.ResponseBody] = blw.body.String()

		logger.Info(logging.RequestResponse, logging.Api, "nil", keys)
	}
}
