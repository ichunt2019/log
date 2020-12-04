package logger

import (
	"testing"
	"time"
)




func TestLog(t *testing.T){

	logConfig := make(map[string]string)
	logConfig["log_path"] = "logs/test1"
	logConfig["log_chan_size"] = "1000"
	logConfig["log_name"] = "xxxxx"
	log ,_:=InitLogger("file",logConfig)
	log.Init()


	logConfig = make(map[string]string)
	logConfig["log_path"] = "logs/test2"
	logConfig["log_chan_size"] = "1000"
	log1 ,_:=InitLogger("file",logConfig)
	log1.Init()


		go func() {
			for i:=0;i<10000;i++{
				log.Warn("5555555555555555555555555")
				log.Info("5555555555555555555555555")
				log.Error("5555555555555555555555555")
				log.Fatal("5555555555555555555555555")
				time.Sleep(time.Second*1)
			}
		}()

		go func() {

			for i:=0;i<10000;i++{
				log1.Warn("666666666666666666666")
				log1.Info("66666666666666666666")
				log1.Error("66666666666666666666")
				log1.Fatal("66666666666666")
				time.Sleep(time.Second*1)
			}
		}()




	time.Sleep(time.Second*3600*2)

}