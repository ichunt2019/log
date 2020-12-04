package logger

import (
	"testing"
	"time"
)




func TestLog(t *testing.T){



	for i:=0;i<=100;i++{
		go func() {
			logConfig := make(map[string]string)
			logConfig["log_path"] = "logs/test1"
			logConfig["log_chan_size"] = "10"
			logConfig["log_name"] = "xxxxx"
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
				log1.Fatal("666666666\"6666\"66666666")
			}
		}()
	}



	time.Sleep(time.Second*5)

}