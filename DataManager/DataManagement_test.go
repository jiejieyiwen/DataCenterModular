package DataManager

import (
	"DataCenterModular/DataDefine"
	"os"
	"strconv"
	"testing"
)

func TestDataManagement(t *testing.T) {
	nRedisWriteType := DataDefine.WRITE_TYPE_PROTO
	if len(os.Args) > 1 {
		for index, k := range os.Args {
			switch k {
			case "-w":
				{
					nTempRedisWriteType, err := strconv.Atoi(os.Args[index+1])
					if err != nil {
						return
					}
					if nTempRedisWriteType < DataDefine.WRITE_SIZE {
						nRedisWriteType = nTempRedisWriteType
					}
				}
			case "-t":
				{
					DataDefine.Token = os.Args[index+1]
				}
			case "-p":
				{
					DataDefine.SuperiorPlantID = os.Args[index+1]
				}
			case "-u":
				{
					DataDefine.HTTP_URL = os.Args[index+1]
				}
			}

		}
	}
	var temp DataManagement
	temp.Init(uint(nRedisWriteType))
	temp.Start()
}
