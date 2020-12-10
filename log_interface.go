package logger

type LogInterface interface {
	Init()
	SetLevel(level int)
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
	//以下三个为同步参数
	syncAdd()
	syncDone()
	SyncWait()


}
