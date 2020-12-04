package logger

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

type LogData struct {
	Message      string `json:"msg"`
	TimeStr      string `json:"dateStr"`
	LevelStr     string `json:"levelStr"`
	Filename     string `json:"fileName"`
	FuncName     string `json:"method"`
	LineNo       int `json:"lineNo"`
	WarnAndFatal bool `json:"warnAndFatal"`
	Req string `json:"request"`
}

//util.go 10

func GetLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(4)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}
	return
}

/*
1. 当业务调用打日志的方法时，我们把日志相关的数据写入到chan（队列）
2. 然后我们有一个后台的线程不断的从chan里面获取这些日志，最终写入到文件。
*/
func writeLog(level int, format string, args ...interface{}) *LogData {
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")
	levelStr := getLevelText(level)

	fileName, funcName, lineNo := GetLineInfo()
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(format, args...)

	logData := &LogData{
		Message:      msg,
		TimeStr:      nowStr,
		LevelStr:     levelStr,
		Filename:     fileName,
		FuncName:     funcName,
		LineNo:       lineNo,
		WarnAndFatal: false,
	}

	if level == LogLevelError || level == LogLevelWarn || level == LogLevelFatal {
		logData.WarnAndFatal = true
	}

	return logData
	//fmt.Fprintf(file, "%s %s (%s:%s:%d) %s\n", nowStr, levelStr, fileName, funcName, lineNo, msg)
}
