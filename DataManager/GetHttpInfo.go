package DataManager

import (
	DataCenterDefine "DataCenterModular/DataDefine"
	"iPublic/DataFactory/DataDefine"
)

type DataInfo struct {
	m_mapServerNatInfo             map[string]DataDefine.ServerNatInfo         //ServerNatInfo
	m_sliceChannelInfo             []DataDefine.ChannelInfo                    //channelinfo
	m_sliceChannelStorageInfo      []DataDefine.ChannelStorageInfo             //ChannelStorageInfo
	m_sliceDeviceinfo              []DataDefine.Deviceinfo                     //Deviceinfo
	m_sliceDevManufacturer         []DataDefine.DevManufacturer                //DevManufacturer
	m_sliceNationStandardChan      []*DataDefine.NationStandardChan            //NationStandardChan
	m_sliceObtainNationalStandards []*DataDefine.ObtainNationalStandardsStruct //ObtainNationalStandards
	m_sliceServerUpdate            []*DataDefine.ServerUpdate                  //ServerUpdate
	m_sliceUpperSvr                []*DataDefine.UpperSvr                      //UpperSvr
	m_sliceChannelData             []*DataDefine.ChannelData                   //ChannelData
	m_sliceOrganizedData           []*DataDefine.OrganizedData                 //OrganizedData
	m_sliceStorageMediumInfo       []DataDefine.StorageMediumInfo              //StorageMediumInfo
	m_sliceStoragePolicyInfo       []DataDefine.StoragePolicyInfo              //StoragePolicyInfo
	m_sliceStorageSchemeInfo       []DataDefine.StorageSchemeInfo              //StorageSchemeInfo
	m_sliceStorageChannelInfo      []*DataDefine.StorageData                   //StorageChannelInfo
}

func NewDataInfo() *DataInfo {
	return &DataInfo{
		m_mapServerNatInfo:             make(map[string]DataDefine.ServerNatInfo),               //ServerNatInfo
		m_sliceChannelInfo:             make([]DataDefine.ChannelInfo, 0, 1),                    //channelinfo
		m_sliceChannelStorageInfo:      make([]DataDefine.ChannelStorageInfo, 0, 1),             //ChannelStorageInfo
		m_sliceDeviceinfo:              make([]DataDefine.Deviceinfo, 0, 1),                     //Deviceinfo
		m_sliceDevManufacturer:         make([]DataDefine.DevManufacturer, 0, 1),                //DevManufacturer
		m_sliceNationStandardChan:      make([]*DataDefine.NationStandardChan, 0, 1),            //NationStandardChan
		m_sliceObtainNationalStandards: make([]*DataDefine.ObtainNationalStandardsStruct, 0, 1), //ObtainNationalStandards
		m_sliceServerUpdate:            make([]*DataDefine.ServerUpdate, 0, 1),                  //ServerUpdate
		m_sliceUpperSvr:                make([]*DataDefine.UpperSvr, 0, 1),                      //UpperSvr
		m_sliceChannelData:             make([]*DataDefine.ChannelData, 0, 1),                   //ChannelData
		m_sliceOrganizedData:           make([]*DataDefine.OrganizedData, 0, 1),                 //OrganizedData
		m_sliceStorageMediumInfo:       make([]DataDefine.StorageMediumInfo, 0, 1),              //StorageMediumInfo
		m_sliceStoragePolicyInfo:       make([]DataDefine.StoragePolicyInfo, 0, 1),              //StoragePolicyInfo
		m_sliceStorageSchemeInfo:       make([]DataDefine.StorageSchemeInfo, 0, 1),              //StorageSchemeInfo
		m_sliceStorageChannelInfo:      make([]*DataDefine.StorageData, 0, 1),                   //StorageChannelInfo
	}
}

//main thread 从HTTP获取数据
func (pThis *DataManagement) GetHTTPInfo() {
	//获取数据
	pThis.getAllDataMessage(DataCenterDefine.Token, DataCenterDefine.SuperiorPlantID)
}

//获取所有数据
func (pThis *DataManagement) getAllDataMessage(token string, platId string) (err error) {
	pThis.m_WaitGroup.Add(12)
	go pThis.goGetAndWriteChannelInfoData(token, platId)
	go pThis.goGetAndWriteServerNatInfoData(token, platId)
	go pThis.goGetAndWriteSChannelStorageInfoData(token, platId)
	go pThis.goGetAndWriteDeviceinfoData(token, platId)
	go pThis.goGetAndWriteDevManufacturerData(token, platId)
	go pThis.goGetAndWriteObtainNationalStandards(token, platId)
	go pThis.goGetAndWriteServerUpdate(token, platId)
	go pThis.goGetAndWriteUpperSvr(token, platId)
	go pThis.goGetAndWriteStorageMediumInfo(token, platId)
	go pThis.goGetAndWriteStoragePolicyInfo(token, platId)
	go pThis.goGetAndWriteStorageSchemeInfo(token, platId)
	go pThis.goGetAndWriteStorageChannelInfo(token, platId)
	pThis.m_WaitGroup.Wait()

	//上级
	for _, upperSvr := range pThis.m_pDataInfo.m_sliceUpperSvr {
		platId = upperSvr.PublicID
		pThis.getSuperiorPlatInfo(token, platId)
	}
	pThis.m_logger.Info("read all")
	return nil
}
