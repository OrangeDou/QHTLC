/*
*@Author: Wenqiang
*@Date:   2023/6/17
*@Time:   14:43
 */
package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*zap.Logger
}

func (l *Logger) NewLogger() *Logger {

	coreList := make([]zapcore.Core, 0)

	//==============================输出终端==============================
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	coreList = append(coreList, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel))

	//==============================输出文件(INFO)==============================
	infoBlogLogger := &lumberjack.Logger{
		Filename:   "server/log/info.log",
		MaxSize:    50,
		MaxBackups: 30,
		MaxAge:     90, //日志最大保存天数
		LocalTime:  true,
		Compress:   false,
	}
	coreList = append(coreList, zapcore.NewCore(encoder, zapcore.AddSync(infoBlogLogger), zapcore.InfoLevel))

	//==============================输出文件(ERROR)==============================
	errBlogLogger := &lumberjack.Logger{
		Filename:   "server/log/err.log",
		MaxSize:    50,
		MaxBackups: 30,
		MaxAge:     90, //日志最大保存天数
		LocalTime:  true,
		Compress:   false,
	}
	coreList = append(coreList, zapcore.NewCore(encoder, zapcore.AddSync(errBlogLogger), zapcore.ErrorLevel))

	//==============================把所有的output,都汇聚成coreList，一次性初始化==============================
	core := zapcore.NewTee(coreList...)

	//==============================设置是否打印堆栈==============================
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))

	return &Logger{
		Logger: zapLogger,
	}

}
