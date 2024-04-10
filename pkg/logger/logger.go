package logger

import (
	"fmt"
	"github.com/xian1367/layout-go/pkg/app"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

// Logger 全局 Logger 对象
var Logger *zap.Logger

// InitLogger 日志初始化
func InitLogger(maxSize, maxBackup, maxAge int, level, filePath string) {
	// 获取日志写入介质
	writeSyncer := getLogWriter(maxSize, maxBackup, maxAge, filePath)

	// 设置日志等级
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		fmt.Println("日志初始化错误，日志级别设置有误。请修改 log.level 配置项")
	}

	// 初始化 core
	core := zapcore.NewCore(getEncoder(), writeSyncer, logLevel)

	// 初始化 Logger
	Logger = zap.New(core,
		zap.AddCaller(),                   // 调用文件和行号，内部使用 runtime.Caller
		zap.AddCallerSkip(1),              // 封装了一层，调用文件去除一层(runtime.Caller(1))
		zap.AddStacktrace(zap.ErrorLevel), // Error 时才会显示 stacktrace
	)

	// 将自定义的 logger 替换为全局的 logger
	// zap.L().Fatal() 调用时，就会使用我们自定的 Logger
	zap.ReplaceGlobals(Logger)
}

// getEncoder 设置日志存储格式
func getEncoder() zapcore.Encoder {
	// 日志格式规则
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller", // 代码调用，如 paginator/paginator.go:148
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      // 每行日志的结尾添加 "\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 日志级别名称大写，如 ERROR、INFO
		EncodeTime:     customTimeEncoder,              // 时间格式，我们自定义为 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller 短格式，如：types/converter.go:17，长格式为绝对路径
	}

	// 本地环境配置
	if app.IsLocal() {
		// 终端输出的关键词高亮
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		// 本地设置内置的 Console 解码器（支持 stacktrace 换行）
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 线上环境使用 JSON 编码器
	return zapcore.NewJSONEncoder(encoderConfig)
}

// customTimeEncoder 自定义友好的时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// getLogWriter 日志记录介质。使用了两种介质，os.Stdout 和文件
func getLogWriter(maxSize, maxBackup, maxAge int, filePath string) zapcore.WriteSyncer {
	// 滚动日志
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filePath + time.Now().Format("2006-01-02.log"),
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		LocalTime:  true, // 使用本地时间
		Compress:   true, // 是否压缩 disabled by default
	}
	// 配置输出介质
	if app.IsLocal() {
		// 本地开发终端打印和记录文件
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	} else {
		// 生产环境只记录文件
		return zapcore.AddSync(lumberJackLogger)
	}
}

// InfoIf 当 err != nil 时记录 info 等级的日志
func InfoIf(err error) {
	if err != nil {
		Logger.Info("Error Occurred:", zap.Error(err))
	}
}

// WarnIf 当 err != nil 时记录 warning 等级的日志
func WarnIf(err error) {
	if err != nil {
		Logger.Warn("Error Occurred:", zap.Error(err))
	}
}

// ErrorIf 当 err != nil 时记录 error 等级的日志
func ErrorIf(err error) {
	if err != nil {
		Logger.Error("Error Occurred:", zap.Error(err))
	}
}

// FatalIf 级别同 Error(), 写完 log 后调用 os.Exit(1) 退出程序
func FatalIf(err error) {
	if err != nil {
		Logger.Fatal("Error Occurred:", zap.Error(err))
	}
}

// DebugField 调试日志，详尽的程序日志
// 调用示例：
//
//	logger.Debug("Database", zap.String("sql", sql))
func DebugField(moduleName string, fields ...zap.Field) {
	if app.IsDebug() {
		Logger.Debug(moduleName, fields...)
	}
}

// InfoField 告知类日志
func InfoField(moduleName string, fields ...zap.Field) {
	Logger.Info(moduleName, fields...)
}

// WarnField 警告类
func WarnField(moduleName string, fields ...zap.Field) {
	Logger.Warn(moduleName, fields...)
}

// ErrorField 错误时记录，不应该中断程序，查看日志时重点关注
func ErrorField(moduleName string, fields ...zap.Field) {
	Logger.Error(moduleName, fields...)
}

// FatalField 级别同 Error(), 写完 log 后调用 os.Exit(1) 退出程序
func FatalField(moduleName string, fields ...zap.Field) {
	Logger.Fatal(moduleName, fields...)
}

// Debug 记录对象类型的 debug 日志，使用 json.Marshal 进行编码。调用示例：
//
//	logger.DebugJSON("读取登录用户", auth.CurrentUser())
func Debug(moduleName string, value interface{}) {
	Logger.Sugar().Debugf(moduleName, value)
}

func Info(moduleName string, value interface{}) {
	Logger.Sugar().Infof(moduleName, value)
}

func Warn(moduleName string, value interface{}) {
	Logger.Sugar().Warnf(moduleName, value)
}

func Error(moduleName string, value interface{}) {
	Logger.Sugar().Errorf(moduleName, value)
}

func Fatal(moduleName, name string, value interface{}) {
	Logger.Sugar().Fatalf(moduleName, value)
}

func ErrorName(moduleName string, name string, value interface{}) {
	Logger.Error(moduleName, zap.Any(name, value))
}
