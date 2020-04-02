package DataManager

import (
	"DataCenterModular/Config"
	DataCenterDefine "DataCenterModular/DataDefine"
	"DataCenterModular/GRpcClient"
	"github.com/sirupsen/logrus"
	"iPublic/LoggerModular"
	"iPublic/iAuthorize"
	"time"
)

//main thread 初始化
func (pThis *DataManagement) Init() error {
	//初始化数据
	pThis.initData()
	//获取服务器列表
	//err := pThis.GetServerURL()
	//if err != nil {
	//	pThis.m_logger.Info("DataManagement init Fail")
	//	return err
	//}
	////获取Http
	//err = pThis.getDataURL()
	//if err != nil {
	//	pThis.m_logger.Info("DataManagement init Fail")
	//	return err
	//}

	err := pThis.GetURL()
	if err != nil {
		pThis.m_logger.Info("DataManagement init Fail")
		return err
	}

	//定时更新在线设备信息  30s
	go func() {
		for {
			tmpErr := pThis.GetEurekaServerList()
			if nil != tmpErr {
				time.Sleep(time.Second * DataCenterDefine.EUREKA_TIME_RETRY)
				continue
			}
			pThis.ConnectToAllServer()
			//休息30s
			time.Sleep(time.Second * DataCenterDefine.EUREKA_TIME_UPDATE)
		}
	}()

	pThis.m_WaitGroup.Add(3)
	//连接MQ
	go pThis.initMQ()
	//连接Redis
	go pThis.initRedis()
	//连接HTTP
	go pThis.InitHTTP()
	pThis.m_WaitGroup.Wait()

	if /*nil == pThis.m_pMQConn || */ nil == pThis.m_pRedisConn || nil == pThis.m_HTTPCon {
		pThis.m_logger.Info("DataManagement init Fail")
		return err
	}

	pThis.m_logger.Info("DataManagement init OK")
	return nil
}

//main thread 初始化数据
func (pThis *DataManagement) initData() {
	pThis.m_mapAllGRpcClients = make(map[string]*GRpcClient.GRpcClient)
	pThis.m_MQMessage = nil
	pThis.m_mapProductsInfo = make(map[string]interface{})
	pThis.m_logger = LoggerModular.GetLogger().WithFields(logrus.Fields{})
	pThis.m_HTTPCon = nil
	pThis.m_pRedisConn = nil
	pThis.m_pMQConn = nil
	pThis.m_nRedisWriteType = DataCenterDefine.WRITE_TYPE_JASON
	pThis.m_pDataInfo = NewDataInfo()
	pThis.m_logger.Info("InitDataManagement Data Success")
	//conf := EnvLoad.GetConf()
	//var err error
	//DataCenterDefine.Token, err = iAuthorize.AesAuthorize(conf.Url.AuthURL, conf.AuthConfig.DataUserName, conf.AuthConfig.DataPassWord, conf.AuthConfig.ClientID, conf.AuthConfig.ClientSecret)
	//if err != nil {
	//	pThis.m_logger.Errorf("Get User Token Error")
	//} else {
	//	pThis.m_logger.Info("Get User Token Success~!")
	//}
	var err error
	DataCenterDefine.Token, err = iAuthorize.AesAuthorize(Config.GetConfig().DataUrlConfig.AuthURL, Config.GetConfig().DataUrlConfig.DataUserName, Config.GetConfig().DataUrlConfig.DataPassWord, Config.GetConfig().DataUrlConfig.ClientID, Config.GetConfig().DataUrlConfig.ClientSecret)
	if err != nil {
		pThis.m_logger.Errorf("Get User Token Error")
		pThis.m_logger.Info("Init DC Failed")
		return
	} else {
		pThis.m_logger.Info("Get User Token Success~!")
	}
}
