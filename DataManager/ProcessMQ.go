package DataManager

import (
	DataCenterDefine "DataCenterModular/DataDefine"
	"DataCenterModular/GRPCTest/Client"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"iPublic/AMQPModular"
	"iPublic/DataFactory/DataDefine"
	"strings"
	"time"
)

//sub thread 传租户名字
func (pThis *DataManagement) GoDealMQMessage(tenant string, consumer string, exChange string) error {
	/*
	  通过声明临时队列，接收数据刷新请求
	*/
	var err error
	var queue *AMQPModular.MqQueue // topic 发送时只绑定了id
	arrString := strings.Split(tenant, "t")
	RoutingKey := fmt.Sprintf("storage.%s", arrString[1])                    // 根据网关类型 租户 设置路由名称
	queueName := fmt.Sprintf("storage-gateway-queue-data-center-%s", tenant) // 根据网关类型 租户 设置队列名称
	fmt.Println("queueName:", queueName)
	// TODO 需要添加一个MQ拨号掉线重回, 掉线后，队列无法自动恢复，需要手动重新定义
	for {
		for {
			queue, err = pThis.m_pMQConn.QueueDeclare(queueName, true, false) // 声明队列, 设置为排他队列，链接断开后自动关闭删除
			if err != nil {
				pThis.m_logger.Errorf("topic-exchange: %v", err)
				time.Sleep(time.Second * 5)
			} else {
				errBind := pThis.m_pMQConn.BindExchangeQueue(queue.QueueObj, exChange, RoutingKey) //绑定交换机
				if errBind != nil {
					pThis.m_logger.Errorf("Bind Queue Error: %v", errBind)
					time.Sleep(time.Second * 5)
				} else {
					//Ques  Consumer Name
					err = pThis.m_pMQConn.AddConsumer(consumer, queue) //添加消费者
					if err != nil {
						pThis.m_logger.Errorf("AddConsumer Error: %v", err)
						return err
					} else {
						break
					}
				}
			}
		}

		//只能有一个消费者
		for _, delivery := range queue.Consumes {
			pThis.m_logger.Infof("MQ Consumer : %s", consumer)
			pThis.m_pMQConn.HandleMessage(delivery, pThis.HandleMessage)
		}
	}
	return nil
}

//sub thread 处理MQ消息
func (pThis *DataManagement) HandleMessage(data []byte) error {
	var msgBody DataCenterDefine.MessageBody
	err := json.Unmarshal(data, &msgBody)
	if nil == err {
		pThis.m_logger.Debugf("Received a message: %v", msgBody)
		pThis.m_MQMessageLock.Lock()
		defer pThis.m_MQMessageLock.Unlock()
		pThis.m_MQMessage = append(pThis.m_MQMessage, msgBody)
		return nil
	}
	return err
}

//main thread 处理MQ消息
func (pThis *DataManagement) HanddleMQMessage() {
	for {
		time.Sleep(time.Millisecond * 20)
		if 0 == len(pThis.m_MQMessage) {
			continue
		}
		pThis.m_MQMessageLock.Lock()
		tmpMQMessege := pThis.m_MQMessage
		pThis.m_MQMessage = nil
		pThis.m_MQMessageLock.Unlock()

		//处理MQ消息
		for _, value := range tmpMQMessege {

			var data []byte //转换为byte切片
			var err error
			//删除操作
			if value.OperationType == DataCenterDefine.OPERATION_DELETE || value.OperationType == DataCenterDefine.OPERATION_BATCH_DELETE {
				strData, ok := value.Data.(string)
				if ok {
					data = []byte(strData)
				}
				//增加更改操作
			} else {
				data, err = json.Marshal(value.Data)
				if nil != err {
					pThis.m_logger.Error(err.Error())
					continue
				}
			}
			switch value.OperationTarget {
			case DataCenterDefine.TARGET_SERVICE_INFO:
				pThis.HanddleServiceInfo(value)
			case DataCenterDefine.TARGET_IP_MANAGE:
				pThis.HanddleIPManage(value)
			case DataCenterDefine.TARGET_PORT_MAPPING:
				pThis.HanddlePortMapping(value)
			case DataCenterDefine.TARGET_NETTYPE:
				pThis.HanddleNetType(value)
			case DataCenterDefine.TARGET_CHANNEL_STORAGE_RELATIONSHIP:
				pThis.OperationChannelInfoByMQ(data, value.OperationType, nil)
			case DataCenterDefine.TARGET_CHANNEL_STORAGE_INFO:
				pThis.OperationChannelStorageInfoByMQ(data, value.OperationType, nil, nil)
			case DataCenterDefine.TARGET_STORAGE_MEDIUM_INFO:
				pThis.HanddleStorageMediumInfo(value)
			case DataCenterDefine.TARGET_STORAGE_POLICY_INFO:
				pThis.HanddleStoragePolicyInfo(value)
			case DataCenterDefine.TARGET_STORAGE_SCHEME_INFO:
				pThis.HanddleStorageSchemeInfo(value)
			case DataCenterDefine.TARGET_STORAGE_SCHEME_DETAIL_INFO:
				pThis.HanddleStorageSchemeDetailInfo(value)
			case DataCenterDefine.TARGET_STORAGE_CHANNEL_INFO:
				pThis.HanddleStorageChannelInfo(value)
			case DataCenterDefine.TARGET_DYNAMIC_STORAGE:
				pThis.HanddleDynamicStorage(value)
			case DataCenterDefine.TARGET_DEVICE_ID_STORAGE_INFO:
				pThis.HanddleDeviceIDStorageInfo(value)
			}
		}
	}
}

//main thread 处理serviceInfo  1
func (pThis *DataManagement) HanddleServiceInfo(msg DataCenterDefine.MessageBody) {
	switch msg.OperationType {
	case DataCenterDefine.OPERATION_ADD:
	case DataCenterDefine.OPERATION_UPDATE:
	case DataCenterDefine.OPERATION_DELETE:
	case DataCenterDefine.OPERATION_BATCH_ADD:
	case DataCenterDefine.OPERATION_BATCH_UPDATE:
	case DataCenterDefine.OPERATION_BATCH_DELETE:
	}
}

//main thread 处理IPManage  2
func (pThis *DataManagement) HanddleIPManage(msg DataCenterDefine.MessageBody) {
	switch msg.OperationType {
	case DataCenterDefine.OPERATION_ADD:
	case DataCenterDefine.OPERATION_UPDATE:
	case DataCenterDefine.OPERATION_DELETE:
	case DataCenterDefine.OPERATION_BATCH_ADD:
	case DataCenterDefine.OPERATION_BATCH_UPDATE:
	case DataCenterDefine.OPERATION_BATCH_DELETE:
	}
}

//main thread 处理PortMapping  3
func (pThis *DataManagement) HanddlePortMapping(msg DataCenterDefine.MessageBody) {
	switch msg.OperationType {
	case DataCenterDefine.OPERATION_ADD:
	case DataCenterDefine.OPERATION_UPDATE:
	case DataCenterDefine.OPERATION_DELETE:
	case DataCenterDefine.OPERATION_BATCH_ADD:
	case DataCenterDefine.OPERATION_BATCH_UPDATE:
	case DataCenterDefine.OPERATION_BATCH_DELETE:
	}
}

//main thread 处理NetType  4 OK
func (pThis *DataManagement) HanddleNetType(msg DataCenterDefine.MessageBody) {
	switch msg.OperationType {
	case DataCenterDefine.OPERATION_ADD:
		fallthrough
	case DataCenterDefine.OPERATION_UPDATE:
		{
			arrData, err := json.Marshal(msg.Data)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}
			var netType DataDefine.NetType
			err = json.Unmarshal(arrData, &netType)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}
			//写入Redis
			data, _ := pThis.getWriteInRedisData(&netType)
			strPrefix := pThis.getWriteInRedisKey(DataCenterDefine.NAME_NET_TYPE)
			pThis.m_pRedisConn.Client.LPush(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_LIST, netType.NetTypeID)
			pThis.m_pRedisConn.Client.HSet(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_DATA, netType.NetTypeID, data)
			pThis.m_logger.Infof("Write NetType Success,NetTypeID [%s]", netType.NetTypeID)
			//通知客户端
			var client Client.GRpcClient
			err, res := client.GrpcSendNotifyToStorage(int32(msg.OperationTarget), int32(msg.OperationType), arrData)
			if err != nil {
				pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			}
			pThis.m_logger.Infof("Get Respond: [%v]", res)
			pThis.notifyToAllServer(msg.OperationType, DataCenterDefine.TARGET_NETTYPE, netType)
		}
	case DataCenterDefine.OPERATION_DELETE:
		{
			//{"operationTarget":4,
			//"operationType":3,
			//"data":"1071413585922359297"}
			arrData, err := json.Marshal(msg.Data)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}
			netTypeID, _ := msg.Data.(string)
			strPrefix := pThis.getWriteInRedisKey(DataCenterDefine.NAME_NET_TYPE)
			strList := strPrefix + DataCenterDefine.REDIS_DATA_SUFFIX_LIST
			strData := strPrefix + DataCenterDefine.REDIS_DATA_SUFFIX_DATA
			//ques
			pThis.m_pRedisConn.Client.LRem(strList, 0, netTypeID)
			//删除
			pIntCmd := pThis.m_pRedisConn.Client.HDel(strData, netTypeID)
			deleteNum, err := pIntCmd.Result()
			if err != nil || deleteNum == 0 {
				pThis.m_logger.Infof("Delete netType Failed ,netTypeID[%s]", netTypeID)
				return
			}
			var client Client.GRpcClient
			err, res := client.GrpcSendNotifyToStorage(int32(msg.OperationTarget), int32(msg.OperationType), arrData)
			if err != nil {
				pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			}
			pThis.m_logger.Infof("Get Respond: [%v]", res)
			pThis.m_logger.Infof("Delete netType Success ,netTypeID[%s]", netTypeID)
		}
	case DataCenterDefine.OPERATION_BATCH_ADD:
	case DataCenterDefine.OPERATION_BATCH_UPDATE:
	case DataCenterDefine.OPERATION_BATCH_DELETE:
	}
}

//------------------------------------------------------------------
type GetChannelInfo struct {
	ChannelInfo      DataDefine.ChannelInfo `json:"channelInfo"`
	Channelinfoid    string                 `json:"channelinfoid"`
	Channelstorageid string                 `json:"channelstorageid"`
	Id               string                 `json:"id"`
	Storagemediumid  string                 `json:"storagemediumid"`
	Storagepolicyid  string                 `json:"storagepolicyid"`
	Storageschemeid  string                 `json:"storageschemeid"`
}

var tmpChannelInfo GetChannelInfo

//main thread 处理ChannelStorageRelationship  5 OK
func (pThis *DataManagement) OperationChannelInfoByMQ(
	data []byte, opType int, ChannelList *[]DataDefine.ChannelInfo) (string, error) {
	//特殊处理

	switch opType {
	case DataCenterDefine.OPERATION_ADD:
		{
			err := json.Unmarshal(data, &tmpChannelInfo)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return "ERROR", err
			}

			channelStorage := DataDefine.ChannelStorageInfo{
				tmpChannelInfo.Channelstorageid,
				tmpChannelInfo.Storageschemeid,
				tmpChannelInfo.Storagemediumid,
				tmpChannelInfo.Channelinfoid,
				tmpChannelInfo.Storagepolicyid,
				0,
				0,
				tmpChannelInfo.Id,
			}

			//写入Redis
			data, _ := pThis.getWriteInRedisData(&channelStorage)
			strPrefix := pThis.getWriteInRedisKey(DataCenterDefine.NAME_CHANNEL_STORAGE_INFO)
			pThis.m_pRedisConn.Client.LPush(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_LIST, channelStorage.Id)
			if err := pThis.m_pRedisConn.HSet("DC_ChannelStorageInfo:Data", channelStorage.Id, data); err != nil {
				pThis.m_logger.Errorf("Write New ChannelStorageInfo:Data Failed, Error[%v]", err)
				return "Write New ChannelStorageInfo:Data Failed", err
			}
			pThis.m_logger.Infof("Write New ChannelStorageInfo:Data Success, Id[%v]", channelStorage.Id)

			////通知客户端
			var client Client.GRpcClient
			err, res := client.GrpcSendNotifyToStorage(int32(opType), int32(DataCenterDefine.OPERATION_ADD), data)
			if err != nil {
				pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			}
			pThis.m_logger.Infof("Get Respond: [%v]", res)
			//pThis.notifyToAllServer(opType, DataCenterDefine.TARGET_CHANNEL_STORAGE_INFO, channelStorage)
			//
			//data, _ = pThis.getWriteInRedisData(&tmpChannelInfo.ChannelInfo)
			//strPrefix = pThis.getWriteInRedisKey(DataCenterDefine.NAME_CHANNEL_INFO)
			//pThis.m_pRedisConn.Client.LPush(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_LIST, tmpChannelInfo.ChannelInfo.ChannelInfoID)
			//pThis.m_pRedisConn.Client.HSet(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_DATA, tmpChannelInfo.ChannelInfo.ChannelInfoID, data)
			//pThis.m_logger.Infof("Write ChannelInfo Success,ChannelInfoID [%s]", tmpChannelInfo.ChannelInfo.ChannelInfoID)
			////通知客户端
			//pThis.notifyToAllServer(opType, DataCenterDefine.TARGET_CHANNEL_STORAGE_RELATIONSHIP, tmpChannelInfo.ChannelInfo)

			return "", nil
		}
	case DataCenterDefine.OPERATION_UPDATE:
		{
			err := json.Unmarshal(data, &tmpChannelInfo)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return "ERROR", err
			}
			//pThis.m_logger.Infof("GetChannelInfo id [%s]", tmpChannelInfo.Id)

			//查
			//strPrefix := pThis.getWriteInRedisKey(DataCenterDefine.NAME_CHANNEL_STORAGE_INFO)
			//prx = strPrefix
			//pStringCmd := pThis.m_pRedisConn.Client.HGet(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_DATA, tmpChannelInfo.Id)
			//if pStringCmd.Err() != nil {
			//	pThis.m_logger.Errorf("No Find channelStorage ,Id[%s]", tmpChannelInfo.Id)
			//	return "ERROR", errors.New("No Find channelStorage ")
			//}
			////strData, _ := pStringCmd.Result()
			//err = json.Unmarshal([]byte(strData), &channelStorage)
			//if err != nil {
			//	pThis.m_logger.Errorf("Unmarshal err[%s] ,channelID[%s]", err.Error(), tmpChannelInfo.Channelinfoid)
			//	return "ERROR", errors.New("Unmarshal err")
			//}
			////更新
			//channelStorage.ChannelStorageInfoID = tmpChannelInfo.Channelstorageid
			//channelStorage.ChannelID = tmpChannelInfo.Channelinfoid
			//time.Sleep(time.Second * 1)
			//channelStorage.StorageSchemeID = deviceIDStorageData.StorageSchemeId
			//pThis.m_logger.Infof("StorageSchemeID[%s]", channelStorage.StorageSchemeID)
			//
			////写入Redis
			//data, _ := pThis.getWriteInRedisData(&channelStorage)
			////pThis.m_pRedisConn.Client.LPush(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_LIST, channelStorage.ChannelID)
			//pThis.m_pRedisConn.Client.HSet(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_DATA, channelStorage.Id, data)
			//pThis.m_logger.Infof("Update ChannelStorageInfo Success,Key :ChannelID[%s]", channelStorage.Id)
			//通知客户端
			var client Client.GRpcClient
			err, res := client.GrpcSendNotifyToStorage(int32(opType), int32(DataCenterDefine.OPERATION_UPDATE), data)
			if err != nil {
				pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			}
			pThis.m_logger.Infof("Get Respond: [%v]", res)
			//pThis.notifyToAllServer(opType, DataCenterDefine.TARGET_CHANNEL_STORAGE_RELATIONSHIP, channelStorage)
			//return "", nil
		}
	case DataCenterDefine.OPERATION_DELETE:
		{
			//var channelStorage DataDefine.ChannelStorageInfo
			//err := json.Unmarshal(data, &channelStorage)
			//if nil != err {
			//	pThis.m_logger.Error(err.Error())
			//	return "ERROR", err
			//}
			//
			//deleteNum1 := 0
			//for _, v := range TempChannelStorageInfo {
			//	if v.ChannelID == tmpChannelInfo.Channelinfoid {
			//		strPrefix := pThis.getWriteInRedisKey(DataCenterDefine.NAME_CHANNEL_STORAGE_INFO)
			//		//删除
			//		pIntCmd := pThis.m_pRedisConn.Client.HDel(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_DATA, channelStorage.Id)
			//		deleteNum, err := pIntCmd.Result()
			//		if err != nil || deleteNum == 0 {
			//			deleteNum1 = int(deleteNum)
			//			pThis.m_logger.Errorf("Delete ChannelStorage:Data Failed, Id[%s]", channelStorage.Id)
			//			return "ERROR", errors.New("No Find channelStorage")
			//		}
			//		pThis.m_logger.Infof("Delete ChannelStorage:Data Success, Id[%s]", channelStorage.Id)
			//		pThis.m_pRedisConn.Client.LRem(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_LIST, 0, channelStorage.Id)
			//		return "", nil
			//	}
			//}
			//if deleteNum1 == 0 {
			//	pThis.m_logger.Errorf("Delete ChannelStorage:Data Failed, DeviceId[%s]", channelStorage.ChannelID)
			//	return "ERROR", errors.New("No Find channelStorage")
			//}
			var client Client.GRpcClient
			err, res := client.GrpcSendNotifyToStorage(int32(opType), int32(DataCenterDefine.OPERATION_DELETE), data)
			if err != nil {
				pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			}
			pThis.m_logger.Infof("Get Respond: [%v]", res)
		}
	}
	return "", nil
}

//main thread 处理ChannelStorageInfo  6 OK
func (pThis *DataManagement) OperationChannelStorageInfoByMQ(
	data []byte, opType int, ChannelStorageList *[]DataDefine.ChannelStorageInfo, ChannelInfo *[]DataDefine.ChannelInfo) (channel string, err error) {
	switch opType {
	case DataCenterDefine.OPERATION_ADD:
		fallthrough
	case DataCenterDefine.OPERATION_UPDATE:
		{
			var channelStorageInfo DataDefine.ChannelStorageInfo
			err = json.Unmarshal(data, &channelStorageInfo)
			if nil != err {
				return "ERROR", err
			}

			//写入Redis
			data, _ := pThis.getWriteInRedisData(&channelStorageInfo)
			strPrefix := pThis.getWriteInRedisKey(DataCenterDefine.NAME_CHANNEL_STORAGE_INFO)
			pThis.m_pRedisConn.Client.LPush(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_LIST, channelStorageInfo.Id)
			pThis.m_pRedisConn.Client.HSet(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_DATA, channelStorageInfo.Id, data)
			pThis.m_logger.Infof("Write ChannelStorageInfo Success,ChannelStorageID [%s]", channelStorageInfo.Id)
			//通知客户端
			var client Client.GRpcClient
			err, res := client.GrpcSendNotifyToStorage(int32(opType), int32(DataCenterDefine.OPERATION_UPDATE), data)
			if err != nil {
				pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			}
			pThis.m_logger.Infof("Get Respond: [%v]", res)
			pThis.notifyToAllServer(opType, DataCenterDefine.TARGET_CHANNEL_STORAGE_INFO, channelStorageInfo)
			return "", nil
		}
		//暂时不考虑
	case DataCenterDefine.OPERATION_DELETE:
		{
			id := string(data)
			strPrefix := pThis.getWriteInRedisKey(DataCenterDefine.NAME_CHANNEL_STORAGE_INFO)
			//删除
			pIntCmd := pThis.m_pRedisConn.Client.HDel(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_DATA, id)
			deleteNum, err := pIntCmd.Result()
			if err != nil || deleteNum == 0 {
				pThis.m_logger.Infof("Delete channelStorage Failed ,id[%s]", id)
				return "ERROR", errors.New("No Find channelStorage ")
			}
			pThis.m_logger.Infof("Delete channelStorage Success ,id[%s]", id)
			pThis.m_pRedisConn.Client.LRem(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_LIST, 0, id)
			var client Client.GRpcClient
			err, res := client.GrpcSendNotifyToStorage(int32(opType), int32(DataCenterDefine.OPERATION_DELETE), data)
			if err != nil {
				pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			}
			pThis.m_logger.Infof("Get Respond: [%v]", res)
			return "", nil
		}
	case DataCenterDefine.OPERATION_BATCH_ADD:
	case DataCenterDefine.OPERATION_BATCH_UPDATE:
	case DataCenterDefine.OPERATION_BATCH_DELETE:
	}
	return "", nil
}

//main thread 处理StorageMediumInfo  7 OK
func (pThis *DataManagement) HanddleStorageMediumInfo(msg DataCenterDefine.MessageBody) {
	switch msg.OperationType {
	case DataCenterDefine.OPERATION_UPDATE:
		fallthrough
	case DataCenterDefine.OPERATION_ADD:
		{
			arrData, err := json.Marshal(msg.Data)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}
			var storageMediumData DataDefine.MQStorageMediumData
			err = json.Unmarshal(arrData, &storageMediumData)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}

			storageMediumInfo := DataDefine.StorageMediumInfo{
				storageMediumData.Id,
				storageMediumData.Type,
				storageMediumData.Url,
				storageMediumData.UserName,
				storageMediumData.Passwd,
				storageMediumData.Path,
				storageMediumData.Expires,
			}
			//写入Redis
			data, _ := pThis.getWriteInRedisData(&storageMediumInfo)
			strPrefix := pThis.getWriteInRedisKey(DataCenterDefine.NAME_STORAGE_MEDIUM_INFO)
			pThis.m_pRedisConn.Client.LPush(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_LIST, storageMediumInfo.StorageMediumInfoID)
			pThis.m_pRedisConn.Client.HSet(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_DATA, storageMediumInfo.StorageMediumInfoID, data)
			pThis.m_logger.Infof("Write StorageMediumData Success,StorageMediumInfoID [%s]", storageMediumInfo.StorageMediumInfoID)
			//通知服务器
			var client Client.GRpcClient
			err, res := client.GrpcSendNotifyToStorage(int32(msg.OperationTarget), int32(msg.OperationType), arrData)
			if err != nil {
				pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			}
			pThis.m_logger.Infof("Get Respond: [%v]", res)
			pThis.notifyToAllServer(msg.OperationType, DataCenterDefine.TARGET_STORAGE_MEDIUM_INFO, storageMediumData)
		}

	case DataCenterDefine.OPERATION_DELETE:
		{
			//{"operationTarget":7,
			//"operationType":3,
			//"data":"1071413585922359297"}
			arrData, err := json.Marshal(msg.Data)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}
			storageMediumInfoID, _ := msg.Data.(string)
			strPrefix := pThis.getWriteInRedisKey(DataCenterDefine.NAME_STORAGE_MEDIUM_INFO)
			//删除
			pIntCmd := pThis.m_pRedisConn.Client.HDel(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_DATA, storageMediumInfoID)
			deleteNum, err := pIntCmd.Result()
			if err != nil || deleteNum == 0 {
				pThis.m_logger.Infof("Delete channelStorage Failed ,storageMediumInfoID[%s]", storageMediumInfoID)
				return
			}
			pThis.m_logger.Infof("Delete channelStorage Success ,storageMediumInfoID[%s]", storageMediumInfoID)
			pThis.m_pRedisConn.Client.LRem(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_LIST, 0, storageMediumInfoID)
			var client Client.GRpcClient
			err, res := client.GrpcSendNotifyToStorage(int32(msg.OperationTarget), int32(msg.OperationType), arrData)
			if err != nil {
				pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			}
			pThis.m_logger.Infof("Get Respond: [%v]", res)
		}
	case DataCenterDefine.OPERATION_BATCH_ADD:
	case DataCenterDefine.OPERATION_BATCH_UPDATE:
	case DataCenterDefine.OPERATION_BATCH_DELETE:
	}
}

//main thread 处理StoragePolicyInfo  8 OK
func (pThis *DataManagement) HanddleStoragePolicyInfo(msg DataCenterDefine.MessageBody) {
	switch msg.OperationType {
	case DataCenterDefine.OPERATION_UPDATE:
		fallthrough
	case DataCenterDefine.OPERATION_ADD:
		{
			arrData, err := json.Marshal(msg.Data)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}
			var storagePolicyData DataDefine.MQStoragePolicyData
			err = json.Unmarshal(arrData, &storagePolicyData)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}

			var storagePolicyInfo = DataDefine.StoragePolicyInfo{
				storagePolicyData.Id,
				storagePolicyData.SliceType,
				int32(storagePolicyData.Slicetime),
			}
			//写入Redis
			data, _ := pThis.getWriteInRedisData(&storagePolicyInfo)
			strPrefix := pThis.getWriteInRedisKey(DataCenterDefine.NAME_STORAGE_POLICY_INFO)
			pThis.m_pRedisConn.Client.LPush(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_LIST, storagePolicyInfo.StoragePolicyID)
			pThis.m_pRedisConn.HSet(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_DATA, storagePolicyInfo.StoragePolicyID, data)
			pThis.m_logger.Infof("Write StoragePolicyData Success,StoragePolicyDataID [%s]", storagePolicyInfo.StoragePolicyID)

			//通知服务器
			var client Client.GRpcClient
			err, res := client.GrpcSendNotifyToStorage(int32(msg.OperationTarget), int32(msg.OperationType), arrData)
			if err != nil {
				pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			}
			pThis.m_logger.Infof("Get Respond: [%v]", res)
			pThis.notifyToAllServer(msg.OperationType, DataCenterDefine.TARGET_STORAGE_POLICY_INFO, storagePolicyInfo)
		}
	case DataCenterDefine.OPERATION_DELETE:
		{
			//{"operationTarget":8,
			//"operationType":3,
			//"data":"1071413585922359297"}
			arrData, err := json.Marshal(msg.Data)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}
			storagePolicyInfoID, _ := msg.Data.(string)
			strPrefix := pThis.getWriteInRedisKey(DataCenterDefine.NAME_STORAGE_POLICY_INFO)
			//删除
			pIntCmd := pThis.m_pRedisConn.Client.HDel(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_DATA, storagePolicyInfoID)
			deleteNum, err := pIntCmd.Result()
			if err != nil || deleteNum == 0 {
				pThis.m_logger.Infof("Delete channelStorage Failed ,storagePolicyInfoID[%s]", storagePolicyInfoID)
				return
			}
			pThis.m_logger.Infof("Delete channelStorage Success ,storagePolicyInfoID[%s]", storagePolicyInfoID)
			pThis.m_pRedisConn.Client.LRem(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_LIST, 0, storagePolicyInfoID)
			var client Client.GRpcClient
			err, res := client.GrpcSendNotifyToStorage(int32(msg.OperationTarget), int32(msg.OperationType), arrData)
			if err != nil {
				pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			}
			pThis.m_logger.Infof("Get Respond: [%v]", res)

		}
	case DataCenterDefine.OPERATION_BATCH_ADD:
	case DataCenterDefine.OPERATION_BATCH_UPDATE:
	case DataCenterDefine.OPERATION_BATCH_DELETE:
	}
}

//main thread 处理StorageSchemeInfo  9 OK
func (pThis *DataManagement) HanddleStorageSchemeInfo(msg DataCenterDefine.MessageBody) {
	switch msg.OperationType {
	case DataCenterDefine.OPERATION_UPDATE:
		fallthrough
	case DataCenterDefine.OPERATION_ADD:
		{
			arrData, err := json.Marshal(msg.Data)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}
			var storageSchemeData DataDefine.StorageSchemeInfo
			err = json.Unmarshal(arrData, &storageSchemeData)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}
			//写入Redis
			data, _ := pThis.getWriteInRedisData(&storageSchemeData)
			strPrefix := pThis.getWriteInRedisKey(DataCenterDefine.NAME_STORAGE_SCHEME_INFO)
			pThis.m_pRedisConn.Client.LPush(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_LIST, storageSchemeData.StorageSchemeInfoID)
			pThis.m_pRedisConn.HSet(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_DATA, storageSchemeData.StorageSchemeInfoID, data)
			pThis.m_logger.Infof("Write StorageSchemeInfo Success,StorageSchemeInfoID [%s]", storageSchemeData.StorageSchemeInfoID)
			//通知服务器
			var client Client.GRpcClient
			err, res := client.GrpcSendNotifyToStorage(int32(msg.OperationTarget), int32(msg.OperationType), arrData)
			if err != nil {
				pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			}
			pThis.m_logger.Infof("Get Respond: [%v]", res)
			pThis.notifyToAllServer(msg.OperationType, DataCenterDefine.TARGET_STORAGE_SCHEME_INFO, storageSchemeData)
		}

	case DataCenterDefine.OPERATION_DELETE:
		{
			arrData, err := json.Marshal(msg.Data)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}
			storageSchemeDataID, _ := msg.Data.(string)
			strPrefix := pThis.getWriteInRedisKey(DataCenterDefine.NAME_STORAGE_SCHEME_INFO)
			//删除
			pIntCmd := pThis.m_pRedisConn.Client.HDel(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_DATA, storageSchemeDataID)
			deleteNum, err := pIntCmd.Result()
			if err != nil || deleteNum == 0 {
				pThis.m_logger.Infof("Delete channelStorage Failed ,storageSchemeDataID[%s]", storageSchemeDataID)
				return
			}
			pThis.m_logger.Infof("Delete channelStorage Success ,storageSchemeDataID[%s]", storageSchemeDataID)
			pThis.m_pRedisConn.Client.LRem(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_LIST, 0, storageSchemeDataID)
			var client Client.GRpcClient
			err, res := client.GrpcSendNotifyToStorage(int32(msg.OperationTarget), int32(msg.OperationType), arrData)
			if err != nil {
				pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			}
			pThis.m_logger.Infof("Get Respond: [%v]", res)
		}
	case DataCenterDefine.OPERATION_BATCH_ADD:
	case DataCenterDefine.OPERATION_BATCH_UPDATE:
	case DataCenterDefine.OPERATION_BATCH_DELETE:
	}
}

//main thread 处理StorageSchemeDetailInfo  10 OK
func (pThis *DataManagement) HanddleStorageSchemeDetailInfo(msg DataCenterDefine.MessageBody) {
	switch msg.OperationType {
	case DataCenterDefine.OPERATION_ADD:
		{
			arrData, err := json.Marshal(msg.Data)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}
			var storageSchemeDetailData DataDefine.MQStorageSchemeDetailData
			err = json.Unmarshal(arrData, &storageSchemeDetailData)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}
			var storageSchemeDetailInfo = DataDefine.StorageSchemeDetailInfo{
				storageSchemeDetailData.StorageSchemeDetailInfoID,
				storageSchemeDetailData.StorageSchemeID,
				storageSchemeDetailData.WeekNum,
				storageSchemeDetailData.StartDateTime,
				storageSchemeDetailData.EndDateTime,
			}

			strPrefix := pThis.getWriteInRedisKey(DataCenterDefine.NAME_STORAGE_SCHEME_INFO)
			pStringCmd := pThis.m_pRedisConn.Client.HGet(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_DATA, storageSchemeDetailInfo.StorageSchemeID)
			if pStringCmd.Err() != nil {
				pThis.m_logger.Errorf("No Find StorageSchemeID[%s]", storageSchemeDetailInfo.StorageSchemeID)
				return
			}
			strData, _ := pStringCmd.Result()

			var storageSchemeInfo DataDefine.StorageSchemeInfo

			err = json.Unmarshal([]byte(strData), &storageSchemeInfo)
			if nil != err {
				pThis.m_logger.Errorf("No Find StorageSchemeID[%s]", storageSchemeDetailInfo.StorageSchemeID)
				return
			}

			if nil == storageSchemeInfo.List {
				storageSchemeInfo.List = []DataDefine.StorageSchemeDetailInfo{storageSchemeDetailInfo}
			} else {
				storageSchemeInfo.List = append(storageSchemeInfo.List, storageSchemeDetailInfo)
			}
			data, _ := pThis.getWriteInRedisData(&storageSchemeInfo)
			strPrefix = pThis.getWriteInRedisKey(DataCenterDefine.NAME_STORAGE_SCHEME_INFO)
			pThis.m_pRedisConn.HSet(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_DATA, storageSchemeInfo.StorageSchemeInfoID, data)
			pThis.m_logger.Infof("Write StorageSchemeInfo Success,StorageSchemeInfoID [%s]", storageSchemeInfo.StorageSchemeInfoID)
			//通知服务器
			var client Client.GRpcClient
			err, res := client.GrpcSendNotifyToStorage(int32(msg.OperationTarget), int32(msg.OperationType), arrData)
			if err != nil {
				pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			}
			pThis.m_logger.Infof("Get Respond: [%v]", res)
			pThis.notifyToAllServer(msg.OperationType, DataCenterDefine.TARGET_STORAGE_SCHEME_INFO, storageSchemeInfo)
		}
	case DataCenterDefine.OPERATION_UPDATE:
		{
			arrData, err := json.Marshal(msg.Data)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}
			var storageSchemeDetailData DataDefine.MQStorageSchemeDetailData
			err = json.Unmarshal(arrData, &storageSchemeDetailData)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}
			var storageSchemeDetailInfo = DataDefine.StorageSchemeDetailInfo{
				storageSchemeDetailData.StorageSchemeDetailInfoID,
				storageSchemeDetailData.StorageSchemeID,
				storageSchemeDetailData.WeekNum,
				storageSchemeDetailData.StartDateTime,
				storageSchemeDetailData.EndDateTime,
			}

			strPrefix := pThis.getWriteInRedisKey(DataCenterDefine.NAME_STORAGE_SCHEME_INFO)
			pStringCmd := pThis.m_pRedisConn.Client.HGet(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_DATA, storageSchemeDetailInfo.StorageSchemeID)
			if pStringCmd.Err() != nil {
				pThis.m_logger.Errorf("No Find StorageSchemeID[%s]", storageSchemeDetailInfo.StorageSchemeID)
				return
			}
			strData, _ := pStringCmd.Result()

			var storageSchemeInfo DataDefine.StorageSchemeInfo

			err = json.Unmarshal([]byte(strData), &storageSchemeInfo)
			if nil != err {
				pThis.m_logger.Errorf("No Find StorageSchemeID[%s]", storageSchemeDetailInfo.StorageSchemeID)
				return
			}
			bFind := false
			for i := 0; i < len(storageSchemeInfo.List); i++ {
				if storageSchemeInfo.List[i].StorageSchemeDetailInfoID == storageSchemeDetailInfo.StorageSchemeDetailInfoID {
					storageSchemeInfo.List[i] = storageSchemeDetailInfo
					bFind = true
					break
				}
			}
			if false == bFind {
				pThis.m_logger.Errorf("No Find StorageSchemeDetailInfoID[%s] in StorageSchemeID[%s]", storageSchemeDetailInfo.StorageSchemeDetailInfoID, storageSchemeDetailInfo.StorageSchemeID)
				return
			}
			data, _ := pThis.getWriteInRedisData(&storageSchemeInfo)
			strPrefix = pThis.getWriteInRedisKey(DataCenterDefine.NAME_STORAGE_SCHEME_INFO)
			pThis.m_pRedisConn.HSet(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_DATA, storageSchemeInfo.StorageSchemeInfoID, data)
			pThis.m_logger.Infof("Write StorageSchemeInfo Success,StorageSchemeInfoID [%s]", storageSchemeInfo.StorageSchemeInfoID)
			//通知服务器
			var client Client.GRpcClient
			err, res := client.GrpcSendNotifyToStorage(int32(msg.OperationTarget), int32(msg.OperationType), arrData)
			if err != nil {
				pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			}
			pThis.m_logger.Infof("Get Respond: [%v]", res)
			pThis.notifyToAllServer(msg.OperationType, DataCenterDefine.TARGET_STORAGE_SCHEME_INFO, storageSchemeInfo)
		}
	case DataCenterDefine.OPERATION_DELETE:
		{
			arrData, err := json.Marshal(msg.Data)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}
			storageSchemeDetailDataID, _ := msg.Data.(string)
			fmt.Println(storageSchemeDetailDataID)
			var client Client.GRpcClient
			err, res := client.GrpcSendNotifyToStorage(int32(msg.OperationTarget), int32(msg.OperationType), arrData)
			if err != nil {
				pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			}
			pThis.m_logger.Infof("Get Respond: [%v]", res)
		}
	case DataCenterDefine.OPERATION_BATCH_ADD:
		{
			batchMap, ok := msg.Data.([]interface{})
			if ok {
				for _, mapSource := range batchMap {
					arrData, err := json.Marshal(mapSource)
					if nil == err {
						var storageSchemeDetailData DataDefine.MQStorageSchemeDetailData
						err = json.Unmarshal(arrData, &storageSchemeDetailData)
						if nil == err {
							fmt.Println(storageSchemeDetailData)
						}
					}
				}
			}
		}
	case DataCenterDefine.OPERATION_BATCH_UPDATE:
		{
			batchMap, ok := msg.Data.([]interface{})
			if ok {
				for _, mapSource := range batchMap {
					arrData, err := json.Marshal(mapSource)
					if nil == err {
						var storageSchemeDetailData DataDefine.MQStorageSchemeDetailData
						err = json.Unmarshal(arrData, &storageSchemeDetailData)
						if nil == err {
							fmt.Println(storageSchemeDetailData)
						}
					}
				}
			}
		}
	case DataCenterDefine.OPERATION_BATCH_DELETE:
		{
			storageSchemeDetailInfoIDs, ok := msg.Data.([]interface{})
			if ok {
				for _, storageSchemeDetailInfoIDSource := range storageSchemeDetailInfoIDs {
					storageSchemeDetailInfoID, ok := storageSchemeDetailInfoIDSource.(string)
					if ok {
						fmt.Println(storageSchemeDetailInfoID)
					}
				}
			}
		}
	}
}

//main thread 处理StorageChannelInfo  11
func (pThis *DataManagement) HanddleStorageChannelInfo(msg DataCenterDefine.MessageBody) {
	switch msg.OperationType {
	case DataCenterDefine.OPERATION_ADD:
	case DataCenterDefine.OPERATION_UPDATE:
	case DataCenterDefine.OPERATION_DELETE:
	case DataCenterDefine.OPERATION_BATCH_ADD:
	case DataCenterDefine.OPERATION_BATCH_UPDATE:
	case DataCenterDefine.OPERATION_BATCH_DELETE:
	}
}

//main thread 处理DynamicStorage  12
func (pThis *DataManagement) HanddleDynamicStorage(msg DataCenterDefine.MessageBody) {
	switch msg.OperationType {
	case DataCenterDefine.OPERATION_ADD:
	case DataCenterDefine.OPERATION_UPDATE:
	case DataCenterDefine.OPERATION_DELETE:
	case DataCenterDefine.OPERATION_BATCH_ADD:
	case DataCenterDefine.OPERATION_BATCH_UPDATE:
	case DataCenterDefine.OPERATION_BATCH_DELETE:
	}
}

//动存 14
func (pThis *DataManagement) HanddleDeviceIDStorageInfo(msg DataCenterDefine.MessageBody) {
	pThis.m_logger.Infof("Get MQ OperationTarget: [%v], OperationType: [%v]", msg.OperationTarget, msg.OperationType)
	switch msg.OperationType {
	case DataCenterDefine.OPERATION_ADD:
		{
			arrData, err := json.Marshal(msg.Data)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}
			var deviceIDStorageData DataDefine.StorageData
			err = json.Unmarshal(arrData, &deviceIDStorageData)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}

			//写入Redis
			data, _ := pThis.getWriteInRedisData(&deviceIDStorageData)
			pThis.m_pRedisConn.Client.LPush("DC_StorageChannelInfo:List", deviceIDStorageData.DeviceId)
			if err := pThis.m_pRedisConn.HSet("DC_StorageChannelInfo:Data", deviceIDStorageData.DeviceId, data); err != nil {
				pThis.m_logger.Errorf("Write New DC_StorageChannelInfo:Data To Redis Failed: [%v]", err)
				return
			}
			pThis.m_logger.Infof("Write New DC_StorageChannelInfo:Data To Redis Success, DeviceId[%v]", deviceIDStorageData.DeviceId)

			//通知服务器
			var client Client.GRpcClient
			//err, res := client.GrpcSendNotify()
			//if err != nil {
			//	pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			//}
			//pThis.m_logger.Infof("Get Respond: [%v]", res)

			err, res := client.GrpcSendNotifyToStorage(int32(msg.OperationTarget), int32(msg.OperationType), arrData)
			if err != nil {
				pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			}
			pThis.m_logger.Infof("Get Respond: [%v]", res)
			pThis.notifyToAllServer(msg.OperationType, DataCenterDefine.TARGET_STORAGE_SCHEME_DETAIL_INFO, deviceIDStorageData)
		}
	case DataCenterDefine.OPERATION_UPDATE:
		{
			arrData, err := json.Marshal(msg.Data)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}
			var deviceIDStorageData DataDefine.StorageData
			err = json.Unmarshal(arrData, &deviceIDStorageData)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}

			data2, _ := pThis.getWriteInRedisData(&deviceIDStorageData)
			pThis.m_pRedisConn.Client.LPush("DC_StorageChannelInfo:List", deviceIDStorageData.DeviceId)
			if err := pThis.m_pRedisConn.HSet("DC_StorageChannelInfo:Data", deviceIDStorageData.DeviceId, data2); err != nil {
				pThis.m_logger.Errorf("Update DC_StorageChannelInfo:Data To Redis Failed: [%v]", err)
				return
			}
			pThis.m_logger.Infof("Update DC_StorageChannelInfo:Data To Redis Success, DeviceId[%v]", deviceIDStorageData.DeviceId)

			//找到相应的ChannelStorageInfo
			data, err := pThis.m_pRedisConn.Client.HGet("DC_ChannelStorageInfo:Data", tmpChannelInfo.Id).Result()
			if err != nil {
				pThis.m_logger.Errorf("Get ChannelStorageInfo From Redis Error: [%v]", err)
				return
			}
			var cs DataDefine.ChannelStorageInfo
			err = json.Unmarshal([]byte(data), &cs)
			if err != nil {
				pThis.m_logger.Errorf("Unmarshal ChannelStorageInfo Error: [%v]", err)
				return
			}

			//更新
			cs.Id = tmpChannelInfo.Id
			cs.ChannelID = tmpChannelInfo.Channelinfoid
			cs.ChannelStorageInfoID = tmpChannelInfo.Channelstorageid
			cs.StorageSchemeID = deviceIDStorageData.StorageSchemeId
			pThis.m_logger.Infof("StorageSchemeID[%s]", cs.StorageSchemeID)

			//写入Redis
			data1, _ := json.Marshal(cs)
			pThis.m_pRedisConn.Client.HSet("DC_ChannelStorageInfo:Data", cs.Id, data1)
			pThis.m_logger.Infof("Update ChannelStorageInfo Success, Key :Id[%s]", cs.Id)

			//通知服务器
			var client Client.GRpcClient
			//err, res := client.GrpcSendNotify()
			//if err != nil {
			//	pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			//}
			//pThis.m_logger.Infof("Get Respond: [%v]", res)

			err, res := client.GrpcSendNotifyToStorage(int32(msg.OperationTarget), int32(msg.OperationType), arrData)
			if err != nil {
				pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			}
			pThis.m_logger.Infof("Get Respond: [%v]", res)
			pThis.notifyToAllServer(msg.OperationType, DataCenterDefine.TARGET_STORAGE_SCHEME_DETAIL_INFO, deviceIDStorageData)
		}
	case DataCenterDefine.OPERATION_DELETE:
		{
			arrData, err := json.Marshal(msg.Data)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}
			var deviceIDStorageData DataDefine.StorageData
			err = json.Unmarshal(arrData, &deviceIDStorageData)
			if nil != err {
				pThis.m_logger.Error(err.Error())
				return
			}

			//删除ChannelStorage
			deleteNum1 := 0
			pThis.m_logger.Infof("TempChannelStorageInfo len [%v]", len(TempChannelStorageInfo))
			for _, v := range TempChannelStorageInfo {
				if v.ChannelID == deviceIDStorageData.DeviceId {
					strPrefix := pThis.getWriteInRedisKey(DataCenterDefine.NAME_CHANNEL_STORAGE_INFO)
					//删除
					pIntCmd := pThis.m_pRedisConn.Client.HDel(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_DATA, v.Id)
					deleteNum, err := pIntCmd.Result()
					if err != nil || deleteNum == 0 {
						pThis.m_logger.Errorf("Delete ChannelStorage:Data Failed, Id[%s]", v.Id)
						break
					}
					pThis.m_logger.Infof("Delete ChannelStorage:Data Success, Id[%s]", v.Id)
					deleteNum1++
					pThis.m_pRedisConn.Client.LRem(strPrefix+DataCenterDefine.REDIS_DATA_SUFFIX_LIST, 0, v.Id)
					break
				}
			}

			if deleteNum1 == 0 {
				pThis.m_logger.Errorf("Delete ChannelStorage:Data Failed, DeviceId[%s]", deviceIDStorageData.DeviceId)
			}

			//删除
			result := pThis.m_pRedisConn.Client.HDel("DC_StorageChannelInfo:Data", deviceIDStorageData.DeviceId)
			deleteNum, err := result.Result()
			if err != nil || deleteNum == 0 {
				pThis.m_logger.Infof("Delete From DC_StorageChannelInfo:Data Failed, DeviceId: [%v]", deviceIDStorageData.DeviceId)
				return
			}
			pThis.m_logger.Infof("Delete From DC_StorageChannelInfo:Data Success, DeviceId: [%v]", deviceIDStorageData.DeviceId)
			pThis.m_pRedisConn.Client.LRem("DC_StorageChannelInfo:Data", 0, deviceIDStorageData.DeviceId)

			//通知更新
			var client Client.GRpcClient
			//err, res := client.GrpcSendNotify()
			//if err != nil {
			//	pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			//}
			//pThis.m_logger.Infof("Get Respond: [%v]", res)

			err, res := client.GrpcSendNotifyToStorage(int32(msg.OperationTarget), int32(msg.OperationType), arrData)
			if err != nil {
				pThis.m_logger.Infof("Send Notify To SearchServer Error: [%v]", err)
			}
			pThis.m_logger.Infof("Get Respond: [%v]", res)

			pThis.notifyToAllServer(msg.OperationType, DataCenterDefine.TARGET_STORAGE_SCHEME_DETAIL_INFO, deviceIDStorageData)
		}
	case DataCenterDefine.OPERATION_BATCH_ADD:
	case DataCenterDefine.OPERATION_BATCH_UPDATE:
	case DataCenterDefine.OPERATION_BATCH_DELETE:
	}
}
