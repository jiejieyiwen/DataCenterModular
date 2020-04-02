package main

import (
	"DataCenterModular/Config"
	"DataCenterModular/DataManager"
	"iPublic/EnvLoad"
	"iPublic/LoggerModular"
)

func init() {
	EnvLoad.GetCmdLineConfig()
}

func TestDataCenter() {
	logger := LoggerModular.GetLogger()

	if err := Config.ReadConfig(); err != nil {
		logger.Error(err)
		return
	}

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
