package DataManager

/*
@Time : 2019-11-1
@Author : 白云飞/邓亦闻
@Software : Golang
@Explain :
*/

import (
	"DataCenterModular/DataDefine"
	"DataCenterModular/GRpcClient"
	"github.com/sirupsen/logrus"
	"iPublic/AMQPModular"
	"iPublic/DataFactory/HTTPModular"
	"iPublic/EnvLoad"
	"iPublic/RedisModular"
	"sync"
)

type DataManagement struct {
	m_logger      *logrus.Entry
	m_WaitGroup   sync.WaitGroup
	m_pRedisConn  *RedisModular.RedisConn //Redis连接
	M_strRedisURL string                  //Redis连接地址

	m_pMQConn  *AMQPModular.RabbServer //MQ连接
	m_strMQURL string                  //MQ连接地址

	m_strEurekaURL  string                  //Eureka连接地址
	m_AllOnlineApps DataDefine.Applications //在线设备  同一个子线程访问

	m_HTTPCon         *HTTPModular.HttpDataModular //Http链接
	m_strDataURL      string                       //HttpData链接
	m_strAuthURL      string                       //Http授权链接
	m_mapProductsInfo map[string]interface{}       //临时存产品信息

	m_mapAllGRpcClients     map[string]*GRpcClient.GRpcClient //Addr Client长连接客户端列表
	m_mapAllGRpcClientsLock sync.Mutex

	m_MQMessage     []DataDefine.MessageBody //临时表 MQ  存数据的
	m_MQMessageLock sync.Mutex               //临时表锁

	m_nRedisWriteType uint      //写入Redis类型
	m_pDataInfo       *DataInfo //写入数据类型
}

var g_DataManagement DataManagement

func GetDataManagement() *DataManagement {
	return &g_DataManagement
}

//main thread 程序启动
func (pThis *DataManagement) Start() {
	DataDefine.TEMP_GRPC, _ = pThis.GetSearchSever()
	//连接MQ通道
	conf := EnvLoad.GetConf()
	if pThis.m_pMQConn != nil {
		//go pThis.GoDealMQMessage("t2000", "test1", DataDefine.MQ_TOPIC_EXCHANGE)
		go pThis.GoDealMQMessage(conf.Tenant, "test", DataDefine.MQ_TOPIC_EXCHANGE)
	}
	//从HTTP获取信息
	pThis.GetHTTPInfo()
	//处理MQ消息
	pThis.HanddleMQMessage()
}
