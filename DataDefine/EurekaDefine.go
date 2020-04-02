package DataDefine

//-----------EurekaConfig---------------
const (
	EUREKA_TIME_UPDATE = 30
	EUREKA_TIME_RETRY  = 5
)

//----------------Eureka-------------------------
type Instance struct {
	InstanceID                    string      `xml:"instanceId"`
	HostName                      interface{} `xml:"hostName"`
	App                           interface{} `xml:"app"`
	IPAddr                        string      `xml:"ipAddr"`
	Status                        interface{} `xml:"status"`
	Overriddenstatus              interface{} `xml:"overriddenstatus"`
	Port                          string      `xml:"port"`
	SecurePort                    interface{} `xml:"securePort"`
	CountryID                     interface{} `xml:"countryId"`
	LeaseInfo                     interface{} `xml:"leaseInfo"`
	Metadata                      interface{} `xml:"metadata"`
	HomePageURL                   interface{} `xml:"homePageUrl"`
	StatusPageURL                 interface{} `xml:"statusPageUrl"`
	HealthCheckURL                interface{} `xml:"healthCheckUrl"`
	VipAddress                    interface{} `xml:"vipAddress"`
	SecureVipAddress              interface{} `xml:"secureVipAddress"`
	IsCoordinatingDiscoveryServer interface{} `xml:"isCoordinatingDiscoveryServer"`
	LastUpdatedTimestamp          interface{} `xml:"lastUpdatedTimestamp"`
	LastDirtyTimestamp            interface{} `xml:"lastDirtyTimestamp"`
	ActionType                    interface{} `xml:"actionType"`
}

type Application struct {
	Name     string   `xml:"name"`
	Instance Instance `xml:"instance"`
}

type Applications struct {
	VersionsDelta interface{}   `xml:"versions__delta"`
	AppsHashcode  interface{}   `xml:"apps__hashcode"`
	Application   []Application `xml:"application"`
}

//-----------------------------------------------
