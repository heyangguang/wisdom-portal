package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// 日志对象
var Logger *zap.Logger


// 初始化Logger
func InitLogger(filename string, maxSize, maxAge, maxBackup int, level string) (err error) {
	// 日志切割第三方包
	ws := getLogWriter(filename, maxSize, maxAge, maxBackup)
	encoderWriteFileJson := getEncoderWriteFileJson()
	encoderConsole := getEncoderConsole()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(level))
	if err != nil {
		return
	}
	core := zapcore.NewTee(
		zapcore.NewCore(encoderWriteFileJson, ws, l),
		zapcore.NewCore(encoderConsole, zapcore.AddSync(os.Stdout), l),
	)
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return
}

// 日志切割
func getLogWriter(filename string, maxSize, maxAge, maxBackup int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		// 单个文件最大尺寸，默认单位M
		MaxSize:    maxSize,
		// 日志最大时间
		MaxAge:     maxAge,
		// 备份日志的数量
		MaxBackups: maxBackup,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 日志格式JSON写到文件里
func getEncoderWriteFileJson() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 人类识别时间
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	// 日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 秒级间隔
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	// 函数调用关系
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 返回JSON
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 日志格式打印控制台  可读性高
func getEncoderConsole() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 人类识别时间
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	// 日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// 秒级间隔
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	// 函数调用关系
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 返回JSON
	return zapcore.NewConsoleEncoder(encoderConfig)
}



// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
			zap.String("user-agent", c.Request.UserAgent()),
		)
	}
}

// GinRecovery recover可能出现的panic，使用zap记录
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				// If the connection is dead, we can't write a status to it.
				if brokenPipe {
					logger.Error(c.Request.URL.Path, zap.Any("error", err), zap.String("request", string(httpRequest)))
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error(
						"[Recovery from panic]", zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error(
						"[Recovery from panic]", zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func Debug(msg string, fields ...zap.Field)  {
	Logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field)  {
	Logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field)  {
	Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field)  {
	Logger.Error(msg, fields...)
}