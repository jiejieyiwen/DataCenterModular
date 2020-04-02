package DataManager

import (
	"DataCenterModular/DataDefine"
	"DataCenterModular/GRpcClient"
	"encoding/xml"
	DataStruct "iPublic/DataFactory/DataDefine"
	"io/ioutil"
	"net/http"
)

//sub thread 获取在线设备信息  会定时调用
func (pThis *DataManagement) GetEurekaServerList() error {
	resp, err := http.Get(pThis.m_strEurekaURL)
	if err != nil {
		pThis.m_logger.Error(err.Error())
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//反序列化
	err = xml.Unmarshal(body, &pThis.m_AllOnlineApps)

	if err != nil {
		pThis.m_logger.Error(err.Error())
		return err
	}
	pThis.m_logger.Info("ServerList Update ~")
	return nil
}

//sub thread 连接服务器 跟 GetEurekaServerList 属于同一个子线程
func (pThis *DataManagement) ConnectToAllServer() {
	pThis.m_mapAllGRpcClientsLock.Lock()
	defer pThis.m_mapAllGRpcClientsLock.Unlock()
	//gRpc
	TestAddr := "192.168.0.130:50051"
	if pClient := pThis.m_mapAllGRpcClients[TestAddr]; pClient == nil {
		pGRPCClient := &GRpcClient.GRpcClient{}
		//建立连接
		err := pGRPCClient.GRpcDial(TestAddr)
		if err != nil {
			pThis.m_logger.Errorf("GRpcClient Connect %s Failed", TestAddr)
		}
		//添加到自己表中
		pThis.m_mapAllGRpcClients[TestAddr] = pGRPCClient
	}
	//在自己表中
	pThis.m_logger.Infoln("Connect GRpcClients Complete")
}

//main thread 通知所有服务器
func (pThis *DataManagement) notifyToAllServer(operType int, nObjType int, data interface{}) {

	//通知所有
	pThis.m_mapAllGRpcClientsLock.Lock()
	defer pThis.m_mapAllGRpcClientsLock.Unlock()
	for strAddr, client := range pThis.m_mapAllGRpcClients {

		//通知服务器
		err := pThis.notifyToServer(client, operType, nObjType, data)
		//通知失败
		if err != nil {
			bConn := false
			//重连3次
			for i := 0; i < 3; i++ {
				err := client.GRpcDial(strAddr)
				if err == nil {
					bConn = true
					break
				}
			}
			//服务器下线
			if false == bConn {
				//表中删除了
				delete(pThis.m_mapAllGRpcClients, strAddr)
			} else {
				//通知服务器
				err := pThis.notifyToServer(client, operType, nObjType, data)
				//通知失败  删掉  服务器不稳定或者没实现grpc接口
				if err != nil {
					delete(pThis.m_mapAllGRpcClients, strAddr)
				}
			}
		}
	}
}

//main thread 通知单个服务器
func (pThis *DataManagement) notifyToServer(client *GRpcClient.GRpcClient, operType int, nObjType int, data interface{}) (err error) {
	switch nObjType {
	case DataDefine.TARGET_NETTYPE:
		{
			netType, ok := data.(DataStruct.NetType)
			if ok {
				_, err = client.OperationNetType(operType, netType)
			}

		}
	case DataDefine.TARGET_CHANNEL_STORAGE_RELATIONSHIP:
		{
			channelInfo, ok := data.(DataStruct.ChannelInfo)
			if ok {
				_, err = client.OperationChannelInfo(operType, channelInfo)
			}
		}
	case DataDefine.TARGET_CHANNEL_STORAGE_INFO:
		{
			channelStorageInfo, ok := data.(DataStruct.ChannelStorageInfo)
			if ok {
				_, err = client.OperationChannelStorageInfo(operType, channelStorageInfo)
			}
		}
	case DataDefine.TARGET_STORAGE_MEDIUM_INFO:
		{
			storageMediumInfo, ok := data.(DataStruct.MQStorageMediumData)
			if ok {
				_, err = client.OperationStorageMediumInfo(operType, storageMediumInfo)
			}
		}
	case DataDefine.TARGET_STORAGE_POLICY_INFO:
		{
			storagePolicyInfo, ok := data.(DataStruct.MQStoragePolicyData)
			if ok {
				_, err = client.OperationStoragePolicyInfo(operType, storagePolicyInfo)
			}

		}
	case DataDefine.TARGET_STORAGE_SCHEME_INFO:
		{
			storageSchemeData, ok := data.(DataStruct.StorageSchemeInfo)
			if ok {
				_, err = client.OperationStorageSchemeInfo(operType, storageSchemeData)
			}
		}
	case DataDefine.TARGET_STORAGE_SCHEME_DETAIL_INFO:
		{
			storageSchemeDetailInfo, ok := data.(DataStruct.MQStorageSchemeDetailData)
			if ok {
				_, err = client.OperationStorageSchemeDetailInfo(operType, storageSchemeDetailInfo)
			}

		}
	}
	return
}
