package DataDefine

//-----------Redis operation ---------------
var (
	REDIS_SAVE_TIME        = 0                   //Redis保存时间
	REDIS_DATA_PREFIX      = "DC_"               //Redis数据前缀
	REDIS_IMITATE_GB       = "51040200002150000" //Redis模拟分文件
	REDIS_STORAGE          = "3"
	REDIS_DATA_SUFFIX_LIST = ":List"
	REDIS_DATA_SUFFIX_DATA = ":Data"
)

//-----------Redis Writen Type ---------------
const (
	WRITE_NONE = iota
	WRITE_TYPE_JASON
	WRITE_TYPE_PROTO
	WRITE_SIZE
)
