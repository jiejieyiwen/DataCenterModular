package DataManager

import (
	"Config"
	DataCenterDefine "DataCenterModular/DataDefine"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//main thread 获取ServerURL中的source

func (pThis *DataManagement) GetURL() error {
	pThis.m_strMQURL = Config.GetConfig().PublicConfig.AMQPURL
	pThis.m_strAuthURL = Config.GetConfig().DataUrlConfig.AuthURL
	pThis.m_strDataURL = Config.GetConfig().DataUrlConfig.DataURL
	pThis.M_strRedisURL = Config.GetConfig().PublicConfig.RedisURL
	pThis.m_strEurekaURL = Config.GetConfig().PublicConfig.EurekaURL
	pThis.m_strEurekaURL += "/apps"
	if pThis.m_strMQURL == "" || pThis.m_strAuthURL == "" || pThis.m_strDataURL == "" || pThis.M_strRedisURL == "" {
		pThis.m_logger.Errorf("Get All URL Failed, mq:[%v], auth:[%v], data:[%v], redis:[%v]",
			pThis.m_strMQURL, pThis.m_strAuthURL, pThis.m_strDataURL, pThis.M_strRedisURL)
		return errors.New("Get All URL Failed")
	}
	pThis.m_logger.Infof("Get All URL Success, mq:[%v], auth:[%v], data:[%v], redis:[%v]",
		pThis.m_strMQURL, pThis.m_strAuthURL, pThis.m_strDataURL, pThis.M_strRedisURL)
	return nil
}

func (pThis *DataManagement) GetServerURL() error {
	httpServerURL := DataCenterDefine.HTTP_URL + "/ServConfig"
	resp, err := http.Get(httpServerURL)
	if err != nil {
		pThis.m_logger.Errorf("getServerURL Failed error : %v\n", err.Error())
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var serverConfig DataCenterDefine.ServConfig
	err = json.Unmarshal(body, &serverConfig)
	if err != nil {
		pThis.m_logger.Errorf("getServerURL Unmarshal Failed,error :%v\n ", err.Error())
		return err
	}
	propertySources, ok := serverConfig.PropertySources[0].(map[string]interface{})
	if ok {
		serverURL, ok := propertySources["source"].(map[string]interface{})
		if ok {
			pThis.M_strRedisURL = serverURL["RedisURL"].(string)
			pThis.m_strMQURL = serverURL["AMQPURL"].(string)
			pThis.m_strEurekaURL = serverURL["EurekaURL"].(string)
			pThis.m_strEurekaURL += "/apps"
			pThis.m_logger.Info("getServerURL OK ~")
			return nil
		}
	}
	pThis.m_logger.Errorln("GetServerUrl Find Failed")
	return errors.New("getServerURL Fail")
}

//main thread 获取DataURL中的source
func (pThis *DataManagement) getDataURL() error {
	httpDataURL := DataCenterDefine.HTTP_URL + "/DataURL"
	resp, err := http.Get(httpDataURL)
	if err != nil {
		pThis.m_logger.Errorf("GetDataURL Connect Failed,error:%v\n", err.Error())
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var serverConfig DataCenterDefine.ServConfig
	err = json.Unmarshal(body, &serverConfig)
	if err != nil {
		pThis.m_logger.Errorf("GetDataURL Unmarshal Failed,error:%v\n", err.Error())
		return err
	}
	propertySources, ok := serverConfig.PropertySources[0].(map[string]interface{})
	if ok {
		serverURL, ok := propertySources["source"].(map[string]interface{})
		if ok {
			pThis.m_strDataURL = serverURL["DataURL"].(string)
			pThis.m_strAuthURL = serverURL["AuthURL"].(string)
			pThis.m_logger.Info("GetDataURL OK ~")
			return nil
		}
	}
	pThis.m_logger.Errorln("GetDataURL Find Failed")
	return errors.New("GetDataURL Fail")
}

func (pThis *DataManagement) GetDataURL1() error {
	httpDataURL := DataCenterDefine.HTTP_URL + "/DataURL"
	resp, err := http.Get(httpDataURL)
	if err != nil {
		pThis.m_logger.Errorf("GetDataURL Connect Failed,error:%v\n", err.Error())
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var serverConfig DataCenterDefine.ServConfig
	err = json.Unmarshal(body, &serverConfig)
	if err != nil {
		pThis.m_logger.Errorf("GetDataURL Unmarshal Failed,error:%v\n", err.Error())
		return err
	}
	propertySources, ok := serverConfig.PropertySources[0].(map[string]interface{})
	if ok {
		serverURL, ok := propertySources["source"].(map[string]interface{})
		if ok {
			pThis.m_strDataURL = serverURL["DataURL"].(string)
			pThis.m_strAuthURL = serverURL["AuthURL"].(string)
			pThis.m_logger.Info("GetDataURL OK ~")
			return nil
		}
	}
	pThis.m_logger.Errorln("GetDataURL Find Failed")
	return errors.New("GetDataURL Fail")
}
