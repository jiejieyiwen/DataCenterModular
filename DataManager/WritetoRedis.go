package DataManager

import (
	DataCenterDefine "DataCenterModular/DataDefine"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"github.com/golang/protobuf/proto"
	"iPublic/DataFactory/DataDefine"
)

//获取写入时的Key值
func (pThis *DataManagement) getWriteInRedisKey(strTypeName string) string {
	return DataCenterDefine.REDIS_DATA_PREFIX + strTypeName
}

//获取写入方式 pb 或 json
func (pThis *DataManagement) getWriteInRedisData(pb proto.Message) ([]byte, error) {
	switch pThis.m_nRedisWriteType {
	case DataCenterDefine.WRITE_TYPE_JASON:
		{
			return json.Marshal(pb)
		}
	case DataCenterDefine.WRITE_TYPE_PROTO:
		{
			return proto.Marshal(pb)
		}
	}
	return nil, errors.New("No This Write Type")
}

//数据格式化
func (pThis *DataManagement) writeInRedisFormat(pipeliner redis.Pipeliner, strStructName string, strKey string, data []byte, platID string) error {
	strPrefix := pThis.getWriteInRedisKey(strStructName)
	if "" != platID {
		strPrefix = platID + ":" + strPrefix
	}
	//写数据
	BoolCmd := pipeliner.HSet(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_DATA, strKey, data)
	if BoolCmd.Err() != nil {
		pThis.m_logger.Errorf("error :%v", BoolCmd.Err().Error())
		return BoolCmd.Err()
	}
	//写映射key值的list表
	IntCmd := pipeliner.LPush(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_LIST, strKey)
	if nil != IntCmd.Err() {
		pThis.m_logger.Errorf("error :%v", IntCmd.Err().Error())
		return IntCmd.Err()
	}
	return nil
}

//拆分数据
func (pThis *DataManagement) splitStruct(data interface{}, name string) (result interface{}) {
	switch name {
	case DataCenterDefine.NAME_CHANNEL_INFO:
		{
			info, ok := data.(*[]DataDefine.ChannelInfo)
			if !ok {
				return nil
			}
			if len(*info) == 0 {
				return info
			}
			if len(*info) < DataCenterDefine.MaxWrite {
				return *info
			}
			a := (*info)[:DataCenterDefine.MaxWrite]
			*info = (*info)[DataCenterDefine.MaxWrite:]
			return a
		}
	case DataCenterDefine.NAME_CHANNEL_STORAGE_INFO:
		{
			info, ok := data.(*[]DataDefine.ChannelStorageInfo)
			if !ok {
				return nil
			}
			if len(*info) == 0 {
				return info
			}
			if len(*info) < DataCenterDefine.MaxWrite {
				return *info
			}
			a := (*info)[:DataCenterDefine.MaxWrite]
			*info = (*info)[DataCenterDefine.MaxWrite:]
			return a
		}
	case DataCenterDefine.NAME_DEVICE_INFO:
		{
			info, ok := data.(*[]DataDefine.Deviceinfo)
			if !ok {
				return nil
			}
			if len(*info) == 0 {
				return info
			}
			if len(*info) < DataCenterDefine.MaxWrite {
				return *info
			}
			a := (*info)[:DataCenterDefine.MaxWrite]
			*info = (*info)[DataCenterDefine.MaxWrite:]
			return a
		}
	case DataCenterDefine.NAME_DEV_MANUFACTURER:
		{
			info, ok := data.(*[]DataDefine.DevManufacturer)
			if !ok {
				return nil
			}
			if len(*info) == 0 {
				return info
			}
			if len(*info) < DataCenterDefine.MaxWrite {
				return *info
			}
			a := (*info)[:DataCenterDefine.MaxWrite]
			*info = (*info)[DataCenterDefine.MaxWrite:]
			return a
		}
	case DataCenterDefine.NAME_OBTAIN_NATIONAL_STANDARD:
		{
			info, ok := data.([]*DataDefine.ObtainNationalStandardsStruct)
			if !ok {
				return nil
			}
			if len(info) == 0 {
				return info
			}
			if len(info) < DataCenterDefine.MaxWrite {
				return info
			}
			a := (info)[:DataCenterDefine.MaxWrite]
			info = (info)[DataCenterDefine.MaxWrite:]
			return a
		}
	case DataCenterDefine.NAME_SERVER_UPDATE:
		{
			info, ok := data.([]*DataDefine.ServerUpdate)
			if !ok {
				return nil
			}
			if len(info) == 0 {
				return info
			}
			if len(info) < DataCenterDefine.MaxWrite {
				return info
			}
			a := (info)[:DataCenterDefine.MaxWrite]
			info = (info)[DataCenterDefine.MaxWrite:]
			return a
		}
	case DataCenterDefine.NAME_UPPER_SVR:
		{
			info, ok := data.([]*DataDefine.UpperSvr)
			if !ok {
				return nil
			}
			if len(info) == 0 {
				return info
			}
			if len(info) < DataCenterDefine.MaxWrite {
				return info
			}
			a := (info)[:DataCenterDefine.MaxWrite]
			info = (info)[DataCenterDefine.MaxWrite:]
			return a
		}
	case DataCenterDefine.NAME_STORAGE_MEDIUM_INFO:
		{
			info, ok := data.(*[]DataDefine.StorageMediumInfo)
			if !ok {
				return nil
			}
			if len(*info) == 0 {
				return info
			}
			if len(*info) < DataCenterDefine.MaxWrite {
				return *info
			}
			a := (*info)[:DataCenterDefine.MaxWrite]
			*info = (*info)[DataCenterDefine.MaxWrite:]
			return a
		}
	case DataCenterDefine.NAME_STORAGE_POLICY_INFO:
		{
			info, ok := data.(*[]DataDefine.StoragePolicyInfo)
			if !ok {
				return nil
			}
			if len(*info) == 0 {
				return info
			}
			if len(*info) < DataCenterDefine.MaxWrite {
				return *info
			}
			a := (*info)[:DataCenterDefine.MaxWrite]
			*info = (*info)[DataCenterDefine.MaxWrite:]
			return a
		}
	case DataCenterDefine.NAME_STORAGE_SCHEME_INFO:
		{
			info, ok := data.(*[]DataDefine.StorageSchemeInfo)
			if !ok {
				return nil
			}
			if len(*info) == 0 {
				return info
			}
			if len(*info) < DataCenterDefine.MaxWrite {
				return *info
			}
			a := (*info)[:DataCenterDefine.MaxWrite]
			*info = (*info)[DataCenterDefine.MaxWrite:]
			return a
		}
	case DataCenterDefine.NAME_CHANNEL_DATA:
		{
			info, ok := data.([]*DataDefine.ChannelData)
			if !ok {
				return nil
			}
			if len(info) == 0 {
				return info
			}
			if len(info) < DataCenterDefine.MaxWrite {
				return info
			}
			a := (info)[:DataCenterDefine.MaxWrite]
			info = (info)[DataCenterDefine.MaxWrite:]
			return a
		}
	case DataCenterDefine.NAME_ORGANIZED_DATA:
		{
			info, ok := data.([]*DataDefine.OrganizedData)
			if !ok {
				return nil
			}
			if len(info) == 0 {
				return info
			}
			if len(info) < DataCenterDefine.MaxWrite {
				return info
			}
			a := (info)[:DataCenterDefine.MaxWrite]
			info = (info)[DataCenterDefine.MaxWrite:]
			return a
		}
	case DataCenterDefine.NAME_NATION_STANDARD_CHAN:
		{
			info, ok := data.([]*DataDefine.NationStandardChan)
			if !ok {
				return nil
			}
			if len(info) == 0 {
				return info
			}
			if len(info) < DataCenterDefine.MaxWrite {
				return info
			}
			a := (info)[:DataCenterDefine.MaxWrite]
			info = (info)[DataCenterDefine.MaxWrite:]
			return a
		}
	case DataCenterDefine.NAME_STORAGE_CHANNEL_INFO:
		{
			info, ok := data.(*[]*DataDefine.StorageData)
			if !ok {
				return nil
			}
			if len(*info) == 0 {
				return *info
			}
			if len(*info) < DataCenterDefine.MaxWrite {
				return *info
			}
			a := (*info)[:DataCenterDefine.MaxWrite]
			*info = (*info)[DataCenterDefine.MaxWrite:]
			return a
		}
	}
	return nil
}

//写入Redis
func (pThis *DataManagement) WriteToRedis(data interface{}, name string, pipeliner redis.Pipeliner) error {
	result := pThis.splitStruct(data, name)
	if result == nil {
		return errors.New("Split Struct Error!~~")
	}
	switch name {
	case DataCenterDefine.NAME_CHANNEL_INFO:
		{
			info, ok := result.([]DataDefine.ChannelInfo)
			if !ok {
				break
			}
			for _, value := range info {
				data, _ := pThis.getWriteInRedisData(&value)
				pThis.writeInRedisFormat(pipeliner, DataCenterDefine.NAME_CHANNEL_INFO, value.ChannelInfoID, data, "")
			}
		}
	case DataCenterDefine.NAME_CHANNEL_STORAGE_INFO:
		{
			info, ok := result.([]DataDefine.ChannelStorageInfo)
			if !ok {
				break
			}
			for _, value := range info {
				data, _ := pThis.getWriteInRedisData(&value)
				pThis.writeInRedisFormat(pipeliner, DataCenterDefine.NAME_CHANNEL_STORAGE_INFO, value.Id, data, "")
			}
		}
	case DataCenterDefine.NAME_DEVICE_INFO:
		{
			info, ok := result.([]DataDefine.Deviceinfo)
			if !ok {
				break
			}
			for _, value := range info {
				data, _ := pThis.getWriteInRedisData(&value)
				pThis.writeInRedisFormat(pipeliner, DataCenterDefine.NAME_DEVICE_INFO, value.DeviceInfoID, data, "")
			}
		}
	case DataCenterDefine.NAME_DEV_MANUFACTURER:
		{
			info, ok := result.([]DataDefine.DevManufacturer)
			if !ok {
				break
			}
			for _, value := range info {
				data, _ := pThis.getWriteInRedisData(&value)
				pThis.writeInRedisFormat(pipeliner, DataCenterDefine.NAME_DEV_MANUFACTURER, value.DmID, data, "")
			}
		}
	case DataCenterDefine.NAME_OBTAIN_NATIONAL_STANDARD:
		{
			info, ok := result.([]*DataDefine.ObtainNationalStandardsStruct)
			if !ok {
				break
			}
			for _, value := range info {
				data, _ := pThis.getWriteInRedisData(value)
				pThis.writeInRedisFormat(pipeliner, DataCenterDefine.NAME_OBTAIN_NATIONAL_STANDARD, value.PublicID, data, "")
			}
		}
	case DataCenterDefine.NAME_SERVER_UPDATE:
		{
			info, ok := result.([]*DataDefine.ServerUpdate)
			if !ok {
				break
			}
			for _, value := range info {
				data, _ := pThis.getWriteInRedisData(value)
				pThis.writeInRedisFormat(pipeliner, DataCenterDefine.NAME_SERVER_UPDATE, value.PublicID, data, "")
			}
		}
	case DataCenterDefine.NAME_UPPER_SVR:
		{
			info, ok := result.([]*DataDefine.UpperSvr)
			if !ok {
				break
			}
			for _, value := range info {
				data, _ := pThis.getWriteInRedisData(value)
				pThis.writeInRedisFormat(pipeliner, DataCenterDefine.NAME_UPPER_SVR, value.PublicID, data, "")
			}
		}
	case DataCenterDefine.NAME_STORAGE_MEDIUM_INFO:
		{
			info, ok := result.([]DataDefine.StorageMediumInfo)
			if !ok {
				break
			}
			for _, value := range info {
				data, _ := pThis.getWriteInRedisData(&value)
				pThis.writeInRedisFormat(pipeliner, DataCenterDefine.NAME_STORAGE_MEDIUM_INFO, value.StorageMediumInfoID, data, "")
			}
		}
	case DataCenterDefine.NAME_STORAGE_POLICY_INFO:
		{
			info, ok := result.([]DataDefine.StoragePolicyInfo)
			if !ok {
				break
			}
			for _, value := range info {
				data, _ := pThis.getWriteInRedisData(&value)
				pThis.writeInRedisFormat(pipeliner, DataCenterDefine.NAME_STORAGE_POLICY_INFO, value.StoragePolicyID, data, "")
			}
		}
	case DataCenterDefine.NAME_STORAGE_SCHEME_INFO:
		{
			info, ok := result.([]DataDefine.StorageSchemeInfo)
			if !ok {
				break
			}
			for _, value := range info {
				data, _ := pThis.getWriteInRedisData(&value)
				pThis.writeInRedisFormat(pipeliner, DataCenterDefine.NAME_STORAGE_SCHEME_INFO, value.StorageSchemeInfoID, data, "")
			}
		}
	case DataCenterDefine.NAME_STORAGE_CHANNEL_INFO:
		{
			info, ok := result.([]*DataDefine.StorageData)
			if !ok {
				break
			}
			for _, value := range info {
				data, _ := pThis.getWriteInRedisData(value)
				pThis.writeInRedisFormat(pipeliner, DataCenterDefine.NAME_STORAGE_CHANNEL_INFO, value.DeviceId, data, "")
			}
		}
	}
	_, err := pipeliner.Exec()
	if err != nil {
		return err
	}
	return nil
}
