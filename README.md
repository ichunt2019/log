# log
20201203最新版本log


## 部分参数说明

```json
log_chan_size 最高协程处理日志
log_path     日志路径 默认为logs
log_name     文件前缀 如果是 "save" 文件名将是 save_20201210
open_sync    如果值是 "1" 允许开启同步(可使用 wait方法阻塞直至日志完全存入文件,一般用于一次性处理脚本)
     
```

## 示例
``````
func TestLog(t *testing.T){
	for i:=0;i<=100;i++{
		go func() {
			logConfig := make(map[string]string)
			logConfig["log_path"] = "logs/test1"
			logConfig["log_chan_size"] = "10"
			log ,_:=InitLogger("file",logConfig)
			log.Init()

			for i:=0;i<100;i++{
				log.Warn("5555555555555555555555555")
				log.Info("5555555555555555555555555")
				log.Error("5555555555555555555555555")
				log.Fatal("5555555555555555555555555")
			}
		}()

		go func() {
			logConfig := make(map[string]string)
			logConfig["log_path"] = "logs/test2"
			logConfig["log_chan_size"] = "10"
			log1 ,_:=InitLogger("file",logConfig)
			log1.Init()

			for i:=0;i<100;i++{
				log1.Warn("666666666666666666666")
				log1.Info("66666666666666666666")
				log1.Error("66666666666666666666")
				log1.Fatal("666666666666666666666")
			}
		}()
	}
	time.Sleep(time.Second*10000)

}
``````

elk日志格式：
```
{"msg":"5555555555555555555555555","dateStr":"2020-12-04 15:27:34.602","levelStr":"INFO","fileName":"asm_amd64.s","method":"runtime.goexit","lineNo":1337}
```