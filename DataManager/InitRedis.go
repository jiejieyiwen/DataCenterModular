package DataManager

import (
	"errors"
	"iPublic/RedisModular"
)

//sub thread 初始化Redis
func (pThis *DataManagement) initRedis() error {
	defer pThis.m_WaitGroup.Done()
	if nil != pThis.m_pRedisConn {
		pThis.m_logger.Error("Redis has connected")
		return errors.New("Redis has connected")
	}
	pThis.m_pRedisConn = RedisModular.GetRedisPool()

	//pThis.M_strRedisURL = "redis://:B9OxgC3HYg@192.168.0.56:30003/8"
	//pThis.M_strRedisURL = "redis://:S0o9l@7&PO@49.234.88.77:8888/8"
	//pThis.m_strRedisURL = "redis://:inphase123.@192.168.2.64:23680/" + DataCenterDefine.REDIS_STORAGE

	//pThis.M_strRedisURL = "redis://:inphase123.@192.168.2.64:23680/0"

	pThis.M_strRedisURL = "redis://:R7OxmC3HYg@10.0.1.209:6379,10.0.1.210:6379,10.0.1.211:6379,10.0.1.212:6379,10.0.1.228:6379,10.0.1.229:6379/0?PoolSize=5"

	//内部实现了断开重连
	err := pThis.m_pRedisConn.DaliWithURL(pThis.M_strRedisURL)
	if err != nil {
		pThis.m_logger.Errorf("initRedis  Failed ,URL:%v,errors : %v", pThis.M_strRedisURL, err.Error())
		return err
	}
	pThis.m_logger.Infof("initRedis Success, URL: [%v]", pThis.M_strRedisURL)
	return nil
}
