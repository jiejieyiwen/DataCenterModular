package DataManager

import (
	DataCenterDefine "DataCenterModular/DataDefine"
	data "iPublic/DataFactory/DataDefine"
	"strconv"
	"sync"
)

func (pThis *DataManagement) GetSearchSever() (string, error) {
	if pThis.m_pRedisConn == nil {
		pThis.m_logger.Error("m_pRedisConn is nil")
	}
	if pThis.m_pRedisConn.Client == nil {
		pThis.m_logger.Error("m_pRedisConn.Client is nil")
	}
	res, err := pThis.m_pRedisConn.Client.Get("SearchServer").Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

//channelinfo
func (pThis *DataManagement) goGetAndWriteChannelInfoData(token string, platId string) {
	write := true
	defer pThis.m_WaitGroup.Done()
	pipeliner := pThis.m_pRedisConn.Client.Pipeline()
	defer pipeliner.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	pThis.m_HTTPCon.GetChannelInfo(token, &wg, &pThis.m_pDataInfo.m_sliceChannelInfo)
	wg.Wait()
	len := len(pThis.m_pDataInfo.m_sliceChannelInfo)/DataCenterDefine.MaxWrite + 1
	for i := 1; i <= len; i++ {
		err := pThis.WriteToRedis(&pThis.m_pDataInfo.m_sliceChannelInfo, DataCenterDefine.NAME_CHANNEL_INFO, pipeliner)
		if err != nil {
			write = false
			pThis.m_logger.Error("Write Current ChannelInfo Data to Redis Error: %v", err)
		}
	}
	if write {
		pThis.m_logger.Info("All ChannelInfo Data Has Writen to Redis Success~~!!")
	}
}

//ServerNatInfo
func (pThis *DataManagement) goGetAndWriteServerNatInfoData(token string, platId string) {
	defer pThis.m_WaitGroup.Done()
	pipeliner := pThis.m_pRedisConn.Client.Pipeline()
	defer pipeliner.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	pThis.m_HTTPCon.GetNetInfo(token, &wg, &pThis.m_pDataInfo.m_mapServerNatInfo)
	wg.Wait()
	for _, value := range pThis.m_pDataInfo.m_mapServerNatInfo {
		data, _ := pThis.getWriteInRedisData(&value)
		pThis.writeInRedisFormat(pipeliner, DataCenterDefine.NAME_SERVER_NAT_INFO, value.ServerIP, data, "")
	}
	_, err := pipeliner.Exec()
	if err != nil {
		pThis.m_logger.Errorf("Write ServerNetinfo to Redis Error: %v", err)
	} else {
		pThis.m_logger.Infof("Write ServerNetinfo to Redis Success")
	}
}

var TempChannelStorageInfo []data.ChannelStorageInfo

//ChannelStorageInfo
func (pThis *DataManagement) goGetAndWriteSChannelStorageInfoData(token string, platId string) {

	defer pThis.m_WaitGroup.Done()
	pipeliner := pThis.m_pRedisConn.Client.Pipeline()
	defer pipeliner.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	pThis.m_HTTPCon.GetChannelStorageInfo(token, &wg, &pThis.m_pDataInfo.m_sliceChannelStorageInfo)
	wg.Wait()

	//复制一份临时的
	TempChannelStorageInfo = make([]data.ChannelStorageInfo, len(pThis.m_pDataInfo.m_sliceChannelStorageInfo))
	copy(TempChannelStorageInfo, pThis.m_pDataInfo.m_sliceChannelStorageInfo)
	//for _, value := range pThis.m_pDataInfo.m_sliceChannelStorageInfo {
	//	data, _ := pThis.getWriteInRedisData(&value)
	//	pThis.writeInRedisFormat(pipeliner, DataCenterDefine.NAME_CHANNEL_STORAGE_INFO, value.ChannelStorageInfoID, data, "")
	//}
	//
	//_, err := pipeliner.Exec()
	//if err != nil {
	//	pThis.m_logger.Errorf("WriteRedis failed: [%v]", err)
	//	return
	//}
	//pThis.m_logger.Info("WriteRedis Success")

	write := true
	len := len(pThis.m_pDataInfo.m_sliceChannelStorageInfo)/DataCenterDefine.MaxWrite + 1
	for i := 1; i <= len; i++ {
		err := pThis.WriteToRedis(&pThis.m_pDataInfo.m_sliceChannelStorageInfo, DataCenterDefine.NAME_CHANNEL_STORAGE_INFO, pipeliner)
		if err != nil {
			write = false
			pThis.m_logger.Error("Write Current ChannelStorageInfo Data to Redis Error: %v", err)
		}
	}
	if write {
		pThis.m_logger.Info("All ChannelStorageInfo Data Has Writen to Redis Success~~!!")
	}
}

//Deviceinfo
func (pThis *DataManagement) goGetAndWriteDeviceinfoData(token string, platId string) {
	write := true
	defer pThis.m_WaitGroup.Done()
	pipeliner := pThis.m_pRedisConn.Client.Pipeline()
	defer pipeliner.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	pThis.m_HTTPCon.GetDeviceInfo(token, &wg, &pThis.m_pDataInfo.m_sliceDeviceinfo)
	wg.Wait()
	len := len(pThis.m_pDataInfo.m_sliceDeviceinfo)/DataCenterDefine.MaxWrite + 1
	for i := 1; i <= len; i++ {
		err := pThis.WriteToRedis(&pThis.m_pDataInfo.m_sliceDeviceinfo, DataCenterDefine.NAME_DEVICE_INFO, pipeliner)
		if err != nil {
			write = false
			pThis.m_logger.Error("Write Current Deviceinfo Data to Redis Error: %v", err)
		}
	}
	if write {
		pThis.m_logger.Info("All Deviceinfo Data Has Writen to Redis Success~~!!")
	}
}

//DevManufacturer
func (pThis *DataManagement) goGetAndWriteDevManufacturerData(token string, platId string) {
	write := true
	defer pThis.m_WaitGroup.Done()
	pipeliner := pThis.m_pRedisConn.Client.Pipeline()
	defer pipeliner.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	pThis.m_HTTPCon.GetManufacturerInfo(token, &wg, &pThis.m_pDataInfo.m_sliceDevManufacturer)
	wg.Wait()
	len := len(pThis.m_pDataInfo.m_sliceDevManufacturer)/DataCenterDefine.MaxWrite + 1
	for i := 1; i <= len; i++ {
		err := pThis.WriteToRedis(&pThis.m_pDataInfo.m_sliceDevManufacturer, DataCenterDefine.NAME_DEV_MANUFACTURER, pipeliner)
		if err != nil {
			write = false
			pThis.m_logger.Error("Write Current DevManufacturer Data to Redis Error: %v", err)
		}
	}
	if write {
		pThis.m_logger.Info("All DevManufacturer Data Has Writen to Redis Success~~!!")
	}
}

//GBInfo
func (pThis *DataManagement) goGetAndWriteObtainNationalStandards(token string, platId string) {
	write := true
	defer pThis.m_WaitGroup.Done()
	pipeliner := pThis.m_pRedisConn.Client.Pipeline()
	defer pipeliner.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	pThis.m_HTTPCon.GetGBInfo(token, &wg, &pThis.m_pDataInfo.m_sliceObtainNationalStandards)
	wg.Wait()
	len := len(pThis.m_pDataInfo.m_sliceObtainNationalStandards)/DataCenterDefine.MaxWrite + 1
	for i := 1; i <= len; i++ {
		err := pThis.WriteToRedis(pThis.m_pDataInfo.m_sliceObtainNationalStandards, DataCenterDefine.NAME_OBTAIN_NATIONAL_STANDARD, pipeliner)
		if err != nil {
			write = false
			pThis.m_logger.Error("Write Current GBInfo Data to Redis Error: %v", err)
		}
	}
	if write {
		pThis.m_logger.Info("All GBInfo Data Has Writen to Redis Success~~!!")
	}
}

//ServerUpdate
func (pThis *DataManagement) goGetAndWriteServerUpdate(token string, platId string) {
	write := true
	defer pThis.m_WaitGroup.Done()
	pipeliner := pThis.m_pRedisConn.Client.Pipeline()
	defer pipeliner.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	pThis.m_HTTPCon.GetServerUpdateInfoByTime(token, &wg, 0, &pThis.m_pDataInfo.m_sliceServerUpdate)
	wg.Wait()
	len := len(pThis.m_pDataInfo.m_sliceServerUpdate)/DataCenterDefine.MaxWrite + 1
	for i := 1; i <= len; i++ {
		err := pThis.WriteToRedis(pThis.m_pDataInfo.m_sliceServerUpdate, DataCenterDefine.NAME_SERVER_UPDATE, pipeliner)
		if err != nil {
			write = false
			pThis.m_logger.Error("Write Current ServerUpdate Data to Redis Error: %v", err)
		}
	}
	if write {
		pThis.m_logger.Info("All ServerUpdate Data Has Writen to Redis Success~~!!")
	}
}

//UpperSvr
func (pThis *DataManagement) goGetAndWriteUpperSvr(token string, platId string) {
	write := true
	defer pThis.m_WaitGroup.Done()
	pipeliner := pThis.m_pRedisConn.Client.Pipeline()
	defer pipeliner.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	pThis.m_HTTPCon.GetSuperiorPlatformInfo(token, &wg, &pThis.m_pDataInfo.m_sliceUpperSvr)
	wg.Wait()
	len := len(pThis.m_pDataInfo.m_sliceUpperSvr)/DataCenterDefine.MaxWrite + 1
	for i := 1; i <= len; i++ {
		err := pThis.WriteToRedis(pThis.m_pDataInfo.m_sliceUpperSvr, DataCenterDefine.NAME_UPPER_SVR, pipeliner)
		if err != nil {
			write = false
			pThis.m_logger.Error("Write Current UpperSvr Data to Redis Error: %v", err)
		}
	}
	if write {
		pThis.m_logger.Info("All UpperSvr Data Has Writen to Redis Success~~!!")
	}
}

//StorageMediumInfo
func (pThis *DataManagement) goGetAndWriteStorageMediumInfo(token string, platId string) {
	write := true
	defer pThis.m_WaitGroup.Done()
	pipeliner := pThis.m_pRedisConn.Client.Pipeline()
	defer pipeliner.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	pThis.m_HTTPCon.GetStorageMediumInfo(token, &wg, &pThis.m_pDataInfo.m_sliceStorageMediumInfo)
	len := len(pThis.m_pDataInfo.m_sliceStorageMediumInfo)/DataCenterDefine.MaxWrite + 1
	for i := 1; i <= len; i++ {
		err := pThis.WriteToRedis(&pThis.m_pDataInfo.m_sliceStorageMediumInfo, DataCenterDefine.NAME_STORAGE_MEDIUM_INFO, pipeliner)
		if err != nil {
			write = false
			pThis.m_logger.Error("Write Current StorageMediumInfo Data to Redis Error: %v", err)
		}
	}
	if write {
		pThis.m_logger.Info("All StorageMediumInfo Data Has Writen to Redis Success~~!!")
	}
}

//StoragePolicyInfo
func (pThis *DataManagement) goGetAndWriteStoragePolicyInfo(token string, platId string) {
	write := true
	defer pThis.m_WaitGroup.Done()
	pipeliner := pThis.m_pRedisConn.Client.Pipeline()
	defer pipeliner.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	pThis.m_HTTPCon.GetStoragePolicyInfo(token, &wg, &pThis.m_pDataInfo.m_sliceStoragePolicyInfo)
	len := len(pThis.m_pDataInfo.m_sliceStoragePolicyInfo)/DataCenterDefine.MaxWrite + 1
	for i := 1; i <= len; i++ {
		err := pThis.WriteToRedis(&pThis.m_pDataInfo.m_sliceStoragePolicyInfo, DataCenterDefine.NAME_STORAGE_POLICY_INFO, pipeliner)
		if err != nil {
			write = false
			pThis.m_logger.Error("Write Current StoragePolicyInfo Data to Redis Error: %v", err)
		}
	}
	if write {
		pThis.m_logger.Info("All StoragePolicyInfo Data Has Writen to Redis Success~~!!")
	}
}

//StorageSchemeInfo
func (pThis *DataManagement) goGetAndWriteStorageSchemeInfo(token string, platId string) {
	write := true
	defer pThis.m_WaitGroup.Done()
	pipeliner := pThis.m_pRedisConn.Client.Pipeline()
	defer pipeliner.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	pThis.m_HTTPCon.GetStorageSchemeInfo(token, &wg, &pThis.m_pDataInfo.m_sliceStorageSchemeInfo)
	len := len(pThis.m_pDataInfo.m_sliceStorageSchemeInfo)/DataCenterDefine.MaxWrite + 1
	for i := 1; i <= len; i++ {
		err := pThis.WriteToRedis(&pThis.m_pDataInfo.m_sliceStorageSchemeInfo, DataCenterDefine.NAME_STORAGE_SCHEME_INFO, pipeliner)
		if err != nil {
			write = false
			pThis.m_logger.Error("Write Current StorageSchemeInfo Data to Redis Error: %v", err)
		}
	}
	if write {
		pThis.m_logger.Info("All StorageSchemeInfo Data Has Writen to Redis Success~~!!")
	}
}

//StorageChannelInfo 动态云存
func (pThis *DataManagement) goGetAndWriteStorageChannelInfo(token string, platId string) {
	write := true
	defer pThis.m_WaitGroup.Done()
	pipeliner := pThis.m_pRedisConn.Client.Pipeline()
	defer pipeliner.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	pThis.m_HTTPCon.GetAllStorageChannelInfo(token, &wg, 0, 0, &pThis.m_pDataInfo.m_sliceStorageChannelInfo)
	wg.Wait()
	len := len(pThis.m_pDataInfo.m_sliceStorageChannelInfo)/DataCenterDefine.MaxWrite + 1
	for i := 1; i <= len; i++ {
		err := pThis.WriteToRedis(&pThis.m_pDataInfo.m_sliceStorageChannelInfo, DataCenterDefine.NAME_STORAGE_CHANNEL_INFO, pipeliner)
		if err != nil {
			write = false
			pThis.m_logger.Error("Write Current StorageChannelInfo Data to Redis Error: %v", err)
		}
	}
	if write {
		pThis.m_logger.Info("All StorageChannelInfo Data Has Writen to Redis Success~~!!")
	}
}

//获取与上级相关的数据
func (pThis *DataManagement) getSuperiorPlatInfo(token string, platId string) {
	pipeliner := pThis.m_pRedisConn.Client.Pipeline()
	defer pipeliner.Close()

	pThis.m_WaitGroup.Add(3)
	//ChannelData
	go pThis.m_HTTPCon.GetSuperiorSharingChannelDataInfo(token, &pThis.m_WaitGroup, platId, &pThis.m_pDataInfo.m_sliceChannelData)
	//OrganizedData
	go pThis.m_HTTPCon.GetSuperiorSharingOrgnizedDataInfo(token, &pThis.m_WaitGroup, platId, &pThis.m_pDataInfo.m_sliceOrganizedData)
	//NationStandardChan
	go pThis.m_HTTPCon.GetGBChannelInfo(token, &pThis.m_WaitGroup, platId, &pThis.m_pDataInfo.m_sliceNationStandardChan)

	pThis.m_WaitGroup.Wait()

	pThis.m_WaitGroup.Add(3)
	//ChannelData
	go func() {
		defer pThis.m_WaitGroup.Done()
		for _, value := range pThis.m_pDataInfo.m_sliceChannelData {
			//strKey := REDIS_DATA_PREFIX + "ChannelData_" + strconv.Itoa(int(value.CorpID))
			data, _ := pThis.getWriteInRedisData(value)
			pThis.writeInRedisFormat(pipeliner, DataCenterDefine.NAME_CHANNEL_DATA, value.ChanAor, data, platId)
		}
	}()
	//OrganizedData
	go func() {
		defer pThis.m_WaitGroup.Done()
		for _, value := range pThis.m_pDataInfo.m_sliceOrganizedData {
			//strKey := REDIS_DATA_PREFIX + "OrganizedData_" + value.OrgCode
			data, _ := pThis.getWriteInRedisData(value)
			pThis.writeInRedisFormat(pipeliner, DataCenterDefine.NAME_ORGANIZED_DATA, value.OrgCode, data, platId)
		}
	}()
	//NationStandardChan
	go func() {
		defer pThis.m_WaitGroup.Done()
		for _, value := range pThis.m_pDataInfo.m_sliceNationStandardChan {
			data, _ := pThis.getWriteInRedisData(value)
			pThis.writeInRedisFormat(pipeliner, DataCenterDefine.NAME_NATION_STANDARD_CHAN, strconv.Itoa(int(value.CorpID)), data, platId)
		}
	}()

	pThis.m_WaitGroup.Wait()

	_, err := pipeliner.Exec()
	if err != nil {
		pThis.m_logger.Error("WriteRedis failed")
		return
	}
}
