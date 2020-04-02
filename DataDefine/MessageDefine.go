package DataDefine

//------------Message structName ---------------
const (
	NAME_SERVER_NAT_INFO            = "ServerNatInfo"
	NAME_CHANNEL_INFO               = "ChannelInfo"
	NAME_CHANNEL_STORAGE_INFO       = "ChannelStorageInfo"
	NAME_DEVICE_INFO                = "DeviceInfo"
	NAME_DEV_MANUFACTURER           = "DevManufacturer"
	NAME_NATION_STANDARD_CHAN       = "GBChannel"
	NAME_OBTAIN_NATIONAL_STANDARD   = "GetGB"
	NAME_SERVER_UPDATE              = "ServerUpdate"
	NAME_UPPER_SVR                  = "SuperiorPlatform"
	NAME_CHANNEL_DATA               = "SuperiorSharingChannelInfo"
	NAME_ORGANIZED_DATA             = "SuperiorSharingOrganizedInfo"
	NAME_STORAGE_MEDIUM_INFO        = "StorageMediumInfo"
	NAME_STORAGE_POLICY_INFO        = "StoragePolicyInfo"
	NAME_STORAGE_SCHEME_INFO        = "StorageSchemeInfo"
	NAME_NET_TYPE                   = "NetType"
	NAME_STORAGE_MEDIUM_DATA        = "StorageMediumData"
	NAME_STORAGE_POLICY_DATA        = "StoragePolicyData"
	NAME_STORAGE_SCHEME_DATA        = "StorageSchemeData"
	NAME_STORAGE_SCHEME_DETAIL_DATA = "StorageSchemeDetailData"
	NAME_STORAGE_CHANNEL_INFO       = "StorageChannelInfo"
)

//------------Message operationTarget ---------------
const (
	NONE = iota
	TARGET_SERVICE_INFO
	TARGET_IP_MANAGE
	TARGET_PORT_MAPPING
	TARGET_NETTYPE
	TARGET_CHANNEL_STORAGE_RELATIONSHIP
	TARGET_CHANNEL_STORAGE_INFO
	TARGET_STORAGE_MEDIUM_INFO
	TARGET_STORAGE_POLICY_INFO
	TARGET_STORAGE_SCHEME_INFO
	TARGET_STORAGE_SCHEME_DETAIL_INFO
	TARGET_STORAGE_CHANNEL_INFO
	TARGET_DYNAMIC_STORAGE
	TARGET_DEVICE_ID_STORAGE_INFO = 14
)

//------------Message Operation -----------------
const (
	NONE_INFO = iota
	OPERATION_ADD
	OPERATION_UPDATE
	OPERATION_DELETE
	OPERATION_BATCH_ADD
	OPERATION_BATCH_UPDATE
	OPERATION_BATCH_DELETE
)

//----------------Message Body-------------------------
type MessageBody struct {
	OperationTarget int         `json:"operationTarget"`
	OperationType   int         `json:"operationType"`
	Data            interface{} `json:"data"`
}
