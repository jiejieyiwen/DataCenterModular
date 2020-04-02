package Client

import (
	DataDefine2 "DataCenterModular/DataDefine"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"iPublic/LoggerModular"
	"time"
)

type GRpcClient struct {
	m_pClientCon *grpc.ClientConn
}

func (pThis *GRpcClient) GRpcDial(StrURL string) error {
	logger := LoggerModular.GetLogger()
	//如果已经连接，则断开连接，重新连接
	pThis.Close()
	clientCon, err := grpc.Dial(StrURL, grpc.WithInsecure())
	if err != nil {
		logger.Errorf("Start Grpc Link Failed：[%v]", err)
		return err
	}
	logger.Infof("Start Grpc Link Success: [%v]", StrURL)
	pThis.m_pClientCon = clientCon
	return nil
}

func (pThis *GRpcClient) Close() {
	logger := LoggerModular.GetLogger()
	if nil != pThis.m_pClientCon {
		err := pThis.m_pClientCon.Close()
		if err != nil {
			logger.Errorf("Close Grpc Link Errpr：[%v]", err)
			return
		}
		pThis.m_pClientCon = nil
	}
}

func (pThis *GRpcClient) SendMsg() (*DC_Respond, error) {
	con := NewDC_NotificationClient(pThis.m_pClientCon)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := DC_Request{BGetNew: true}
	res, err := con.DC_Notify(ctx, &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (pThis *GRpcClient) GrpcSendNotify() (error, int32) {
	//var con GRpcClient
	//conf := EnvLoad.GetConf()
	//DataDefine2.TEMP_GRPC += strconv.Itoa(conf.HttpPort)
	if err := pThis.GRpcDial(DataDefine2.TEMP_GRPC); err != nil {
		return err, -1
	}
	res, err := pThis.SendMsg()
	if err != nil {
		return err, -2
	}
	return nil, res.StrRespond
}