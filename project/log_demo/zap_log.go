package log_demo

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// logger 和 sugaredLogger 的区别是：
// logger 是 zap.Logger 类型，提供了结构化日志的功能，但是使用起来不够方便。在每一微秒和每一次内存分配都很重要的上下文中，使用Logger。它甚至比SugaredLogger更快，内存分配次数也更少，但它只支持强类型的结构化日志记录。
// sugaredLogger 是 zap.SugaredLogger 类型，对 logger 进行了封装，提供了更方便的使用方式。在性能很好但不是很关键的上下文中，使用SugaredLogger。它比其他结构化日志记录包快4-10倍，并且支持结构化和printf风格的日志记录。
var (
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger
)

func InitLogger() {
	// logger, _ = zap.NewProduction()
	logger, _ = zap.NewDevelopment()
	sugaredLogger = logger.Sugar()
}

func ZapPrintLog() {
	InitLogger()
	// 在程序结束时，确保日志缓冲区中的所有日志都被写入
	defer logger.Sync()

	logger.Info("This is a log message", zap.String("key1", "value1"), zap.Float64s("key2", []float64{1.0, 2.0, 3.0}))
}

func ZapPrintLog1() {
	InitLogger()
	defer logger.Sync()

	sugaredLogger.Debugf("This is sugared logger debug message %s", "debug")
	sugaredLogger.Infow("This is a log message", "key1", "value1", "key2", []float64{1.0, 2.0, 3.0})
}

func InitCustomLogger() {
	// 将日志写到哪里去
	// writeSyncer := getLogWriter()
	writeSyncer := getLogWriter1()
	// 编码器，如何写入日志
	encoder := getEncoder()

	var core zapcore.Core
	mode := "prod" // dev or prod
	if mode == "dev" {
		// 开发模式下，日志输出到控制台，同时也输出到文件
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		// 生产模式，日志输出到文件
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	}

	// zap.AddCaller() 会在日志中加入调用函数的文件名和行号
	// zap.AddCallerSkip(1) 会跳过调用函数的文件名和行号
	// 当我们不是直接使用初始化好的logger实例记录日志，而是将其包装成一个函数等，此时日录日志的函数调用链会增加，想要获得准确的调用信息就需要通过AddCallerSkip函数来跳过
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugaredLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 大小写编码器
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 使用 JSON 格式输出日志
	// return zapcore.NewJSONEncoder(encoderConfig)
	// 使用 Console 格式输出日志
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.OpenFile("zap_local.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// 也可以利用 io.MultiWriter() 函数同时写到文件和控制台
	ws := io.MultiWriter(os.Stdout, file)
	return zapcore.AddSync(ws)
}

func getLogWriter1() zapcore.WriteSyncer {
	// 因为 zap 本身不支持切割归档日志文件，因此需要借助第三方库 lumberjack 来实现
	// lumberjack.Logger 实现了 io.WriteSyncer 接口，可以直接作为 zap 的 WriteSyncer 使用
	// lumberjack.Logger 的参数如下：
	// Filename 日志文件的位置
	// MaxSize 每个日志文件保存的最大尺寸 单位：MB
	// MaxBackups 保留旧文件的最大个数
	// MaxAge 保留旧文件的最大天数
	// Compress 是否压缩
	// LocalTime 是否使用本地时间
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "zap_local.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
		LocalTime:  true,
	}

	writeSyncer := zapcore.AddSync(lumberJackLogger)
	return writeSyncer
}

func ZapPrintLog2() {
	InitCustomLogger()
	defer logger.Sync()

	logger.Debug("This is a log message", zap.String("key1", "value1"), zap.Float64s("key2", []float64{1.0, 2.0, 3.0}))
}
