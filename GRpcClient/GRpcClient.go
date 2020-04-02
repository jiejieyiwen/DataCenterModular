package GRpcClient

import (
	DataCenterDefine "DataCenterModular/DataDefine"
	"context"
	"fmt"
	"iPublic/DataFactory/DataDefine"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

type GRpcClient struct {
	m_pClientCon *grpc.ClientConn
	m_sw         sync.WaitGroup
}

func (pThis *GRpcClient) GRpcDial(strIp string) error {

	//如果已经连接，则断开连接，重新连接
	pThis.Close()

	clientCon, err := grpc.Dial(strIp, grpc.WithInsecure())
	if err != nil {
		return err
	}
	pThis.m_pClientCon = clientCon
	return nil
}

func (pThis *GRpcClient) Close() {
	if nil != pThis.m_pClientCon {
		pThis.m_pClientCon.Close()
		pThis.m_pClientCon = nil
	}
}

//func (pThis *GRpcClient) Notify(strData string) (string, error) {
//	if nil == pThis.m_pClientCon {
//		return "", errors.New("Client Has No Connected")
//	}
//	pThis.m_sw.Add(1)
//	defer pThis.m_sw.Done()
//	client := Message.NewNotificationClient(pThis.m_pClientCon)
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//
//	req := Message.Request{strData}
//	res, err := client.Notify(ctx, &req)
//	if err != nil {
//		return "", err
//	}
//	return res.GetStrRespon(), nil
//}

//4
func (pThis *GRpcClient) OperationNetType(operType int, netType DataDefine.NetType) (string, error) {

	switch operType {
	case DataCenterDefine.OPERATION_UPDATE:
		fallthrough
	case DataCenterDefine.OPERATION_ADD:
		{
			client := DataDefine.NewNetTypefoUpdateClient(pThis.m_pClientCon)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			res, err := client.Add(ctx, &netType)
			if err != nil {
				fmt.Println("channelStorageInfo发送失败")
				return "FAILED", err
			}
			fmt.Println(res.GetRespon())
		}
	}
	return "SUCCESS", nil
}

//5
func (pThis *GRpcClient) OperationChannelInfo(operType int, channelInfo DataDefine.ChannelInfo) (string, error) {

	switch operType {
	case DataCenterDefine.OPERATION_UPDATE:
		fallthrough
	case DataCenterDefine.OPERATION_ADD:
		{
			client := DataDefine.NewChannelInfoUpdateClient(pThis.m_pClientCon)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			res, err := client.Add(ctx, &channelInfo)
			if err != nil {
				fmt.Println("channelStorageInfo发送失败")
				return "FAILED", err
			}
			fmt.Println(res.GetRespon())
		}
	}
	return "SUCCESS", nil
}

//6
func (pThis *GRpcClient) OperationChannelStorageInfo(operType int, channelStorageInfo DataDefine.ChannelStorageInfo) (string, error) {
	switch operType {
	case DataCenterDefine.OPERATION_UPDATE:
		fallthrough
	case DataCenterDefine.OPERATION_ADD:
		{
			client := DataDefine.NewChannelStorageInfoUpdateClient(pThis.m_pClientCon)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			res, err := client.Add(ctx, &channelStorageInfo)
			if err != nil {
				fmt.Println("channelStorageInfo发送失败")
				return "FAILED", err
			}
			fmt.Println(res.GetRespon())
		}
	}
	return "SUCCESS", nil
}

//7
func (pThis *GRpcClient) OperationStorageMediumInfo(operType int, storageMediumInfo DataDefine.MQStorageMediumData) (string, error) {
	switch operType {
	case DataCenterDefine.OPERATION_UPDATE:
		fallthrough
	case DataCenterDefine.OPERATION_ADD:
		{
			client := DataDefine.NewMQStorageMediumDataUpdateClient(pThis.m_pClientCon)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			res, err := client.Add(ctx, &storageMediumInfo)
			if err != nil {
				fmt.Println("发送失败")
				return "FAILED", err
			}
			fmt.Println(res.GetRespon())
		}
	}
	return "SUCCESS", nil
}

//8
func (pThis *GRpcClient) OperationStoragePolicyInfo(operType int, storagePolicyInfo DataDefine.MQStoragePolicyData) (string, error) {
	switch operType {
	case DataCenterDefine.OPERATION_UPDATE:
		fallthrough
	case DataCenterDefine.OPERATION_ADD:
		{
			client := DataDefine.NewMQStoragePolicyDataUpdateClient(pThis.m_pClientCon)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			res, err := client.Add(ctx, &storagePolicyInfo)
			if err != nil {
				fmt.Println("发送失败")
				return "FAILED", err
			}
			fmt.Println(res.GetRespon())
		}
	}
	return "SUCCESS", nil
}

//9
func (pThis *GRpcClient) OperationStorageSchemeInfo(operType int, storageSchemeData DataDefine.StorageSchemeInfo) (string, error) {
	switch operType {
	case DataCenterDefine.OPERATION_UPDATE:
		fallthrough
	case DataCenterDefine.OPERATION_ADD:
		{
			client := DataDefine.NewStorageSchemeInfoUpdateClient(pThis.m_pClientCon)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			res, err := client.Add(ctx, &storageSchemeData)
			if err != nil {
				fmt.Println("发送失败")
				return "FAILED", err
			}
			fmt.Println(res.GetRespon())
		}
	}
	return "SUCCESS", nil
}

//10
func (pThis *GRpcClient) OperationStorageSchemeDetailInfo(operType int, storageSchemeDetailInfo DataDefine.MQStorageSchemeDetailData) (string, error) {
	switch operType {
	case DataCenterDefine.OPERATION_UPDATE:
		fallthrough
	case DataCenterDefine.OPERATION_ADD:
		{
			client := DataDefine.NewMQStorageSchemeDetailDataUpdateClient(pThis.m_pClientCon)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			res, err := client.Add(ctx, &storageSchemeDetailInfo)
			if err != nil {
				fmt.Println("发送失败")
				return "FAILED", err
			}
			fmt.Println(res.GetRespon())
		}
	}
	return "SUCCESS", nil
}
