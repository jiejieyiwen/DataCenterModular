package main

import (
	DataCenterDefine "DataCenterModular/DataDefine"
	"DataCenterModular/DataManager"

	//"DataCenterModular/DataManager"
	"os"
	"strconv"
)

func TestDataCenter2() {
	//nRedisWriteType := DataCenterDefine.WRITE_TYPE_JASON
	if len(os.Args) > 1 {
		for index, k := range os.Args {
			switch k {
			case "-w":
				{
					nTempRedisWriteType, err := strconv.Atoi(os.Args[index+1])
					if err != nil {
						return
					}
					if nTempRedisWriteType < DataCenterDefine.WRITE_SIZE {
						//nRedisWriteType = nTempRedisWriteType
					}
				}
			case "-t": //token
				{
					DataCenterDefine.Token = os.Args[index+1]
				}
			case "-u": //配置中心
				{
					DataCenterDefine.HTTP_URL = os.Args[index+1]
				}
			case "-rs": //redis存储数据库
				{
					DataCenterDefine.REDIS_STORAGE = os.Args[index+1]
				}
			}

		}
	}
	temp := DataManager.GetDataManagement()
	if err := temp.Init(); nil == err {
		temp.Start()
	}
}

func main1() {
	TestDataCenter2()
}
