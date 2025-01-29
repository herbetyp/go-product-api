package configs

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	zapLog "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zapLog.Logger

	LOG_LEVEL = os.Getenv("LOG_LEVEL")
)

func init() {
	log = zapLog.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		}),
		zapcore.AddSync(os.Stdout),
		zapLog.NewAtomicLevelAt(getLevelLogs()),
	))

	defer log.Sync()
}
func InitDefaultLogs(c *gin.Context) *zapLog.Logger {
	return log.With(
		zapLog.String("request_id", c.GetHeader("X-Request-Id")),
		zapLog.String("ip", c.ClientIP()),
		zapLog.String("method", c.Request.Method),
		zapLog.String("path", c.Request.URL.Path),
		zapLog.String("user_agent", c.GetHeader("User-Agent")),
	)
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}

func Info(message string, tags ...zapLog.Field) {
	log.Info(message, tags...)
	log.Sync()
}

func Error(message string, err error, tags ...zapLog.Field) {
	tags = append(tags, zapLog.NamedError("reason", err))
	log.Error(message, tags...)
	log.Sync()
}

func Debug(message string, tags ...zapLog.Field) {
	log.Debug(message, tags...)
	log.Sync()
}
