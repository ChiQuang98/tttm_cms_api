package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"tttm_cms_api/services"
	"tttm_cms_api/lp-libs/redis"
	"tttm_cms_api/lp-libs/settings"
	"github.com/golang/glog"
)

const version string = "dev_1_1"

func init() {
	//glog
	//create logs folder
	os.Mkdir("./logs", 0777)
	flag.Lookup("stderrthreshold").Value.Set("[INFO|WARN|FATAL]")
	flag.Lookup("logtostderr").Value.Set("false")
	flag.Lookup("alsologtostderr").Value.Set("true")

	flag.Lookup("log_dir").Value.Set("./logs")
	glog.MaxSize = 1024 * 1024 * settings.GetGlogConfig().MaxSize
	flag.Lookup("v").Value.Set(fmt.Sprintf("%d", settings.GetGlogConfig().V))
	flag.Parse()

}

func main() {
	flag.Parse()
	defer func() {
		if err := recover(); err != nil {
			glog.Error("-------------RECOVER err: ", err)
		}
	}()
	glog.Infof("START LP-CMS-CLIENT %v, at: %v", version, time.Now())
	glog.Info("Init Redis database...")

	redis.Init1(settings.GetRedisInfo().Host1,settings.GetRedisInfo().Host2,settings.GetRedisInfo().Host3)
	if !redis.Ping() {
		glog.Info("redis ping error")
		os.Exit(0)
	} else {
		glog.Info("OK")
	}
	//init mqtt
	err := services.SetSuperUser()
	if err == nil {
		glog.Info("Connecting to MQTT broker... ")
		err := services.ConnectMqtt()
		if err != nil {
			glog.Info("Failed, exit")
			os.Exit(1)
		} else {
			glog.Info("OK")
		}
		glog.Info("Connecting to MQTT broker Auth... ")
		err = services.ConnectMQTTOpts()
		if err != nil {
			glog.Info("Failed, exit")
			os.Exit(1)
		} else {
			glog.Info("OK")
		}
	}

	ticker := time.NewTicker(60 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				services.CheckMqttStatusClients()
			}
		}
	}()

	stats()
}

func stats() {
	var s string
	for {
		fmt.Scanln(&s)
		switch s {
		case "q":
			services.DisconnectMqtt()
			glog.Info("exited")
			os.Exit(0)
		case "t":
			services.Test()
		}
	}
}
