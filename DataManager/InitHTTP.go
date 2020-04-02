package DataManager

import (
	"errors"
	"iPublic/DataFactory/HTTPModular"
)

//sub thread 初始化HTTP
func (pThis *DataManagement) InitHTTP() error {
	defer pThis.m_WaitGroup.Done()
	if nil != pThis.m_HTTPCon {
		pThis.m_logger.Error("HTTP Has Connected")
		return errors.New("HTTP Has Connected")
	}
	pThis.m_HTTPCon = &HTTPModular.HttpDataModular{}
	err := pThis.m_HTTPCon.Init(pThis.m_strDataURL)
	if err != nil {
		pThis.m_logger.Errorf("HTTP Connect Failed , error : %v \n", err.Error())
		return err
	}
	return nil
}
