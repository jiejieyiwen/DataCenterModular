package DataManager

import (
	"errors"
	"iPublic/AMQPModular"
)

//sub thread 初始化MQ
func (pThis *DataManagement) initMQ() error {
	defer pThis.m_WaitGroup.Done()
	if nil != pThis.m_pMQConn {
		pThis.m_logger.Error("MQ not connected ")
		return errors.New("MQ not connected")
	}
	pThis.m_pMQConn = new(AMQPModular.RabbServer)
	//pThis.m_strMQURL = "amqp://guest:guest@192.168.0.56:30001/"
	//pThis.m_strMQURL = "amqp://user:1syhl9t@192.168.60.31:5672/"
	//conf := EnvLoad.GetConf()
	//	//pThis.m_strMQURL = conf.ServerConfig.RabbitURL
	err := AMQPModular.GetRabbitMQServ(pThis.m_strMQURL, pThis.m_pMQConn)
	if err != nil {
		pThis.m_logger.Errorf("initMQ  Failed ,errors : %v", err.Error())
		return err
	}
	return nil
}
