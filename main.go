package main

import (
	"Config"
	"DataCenterModular/DataManager"
	"iPublic/EnvLoad"
	"iPublic/LoggerModular"
)

func init() {
	EnvLoad.GetCmdLineConfig()
}

func TestDataCenter() {
	logger := LoggerModular.GetLogger()

	conf := Config.GetConfig()
	if err := Config.ReadConfig(); err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("conf is [%v]", conf)

	temp := DataManager.GetDataManagement()
	if err := temp.Init(); nil != err {
		logger.Errorf("Init DC Error: [%v]", err)
		return
	}
	temp.Start()
}

func main() {
	TestDataCenter()
}
