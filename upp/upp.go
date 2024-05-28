package upp

import (
	"fmt"

	"modify/field"

	"github.com/imdario/mergo"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// YGSH
type YGSH struct {
	UserAddURI      string `yaml:"UserAddURI"`
	BlockUserReport string `yaml:"BlockUserReport"`
	BlockUserSuffix string `yaml:"BlockUserSuffix"`
	Switch          bool   `yaml:"Switch"`
	UserAllURI      string `yaml:"UserAllURI"`
}

// Yaml2Go
type Yaml2Go struct {
	SYSTEM  SYSTEM  `yaml:"SYSTEM"`
	MALWARE MALWARE `yaml:"MALWARE"`
	MTX     MTX     `yaml:"MTX"`
	JMR     JMR     `yaml:"JMR"`
	HTTP    HTTP    `yaml:"HTTP"`
	LOG     LOG     `yaml:"LOG"`
	RULE    RULE    `yaml:"RULE"`
	POLICY  POLICY  `yaml:"POLICY"`
	YGSH    YGSH    `yaml:"YGSH"`
}

// SYSTEM
type SYSTEM struct {
	CMDPort  int    `yaml:"CMDPort"`
	DeviceIP string `yaml:"DeviceIP"`
	Debug    bool   `yaml:"Debug"`
}

// MTX
type MTX struct {
	MtxAllUrl         string `yaml:"MtxAllUrl"`
	SampleScanUrl     string `yaml:"SampleScanUrl"`
	MtxResultPath     string `yaml:"MtxResultPath"`
	Switch            bool   `yaml:"Switch"`
	MtxAddUrl         string `yaml:"MtxAddUrl"`
	SampleResultPath  string `yaml:"SampleResultPath"`
	AESKEY            string `yaml:"AESKEY"`
	SampleTimeout     int    `yaml:"SampleTimeout"`
	MtxMsCallback     string `yaml:"MtxMsCallback"`
	MtxCallbackUrl    string `yaml:"MtxCallbackUrl"`
	SampleCallbackUrl string `yaml:"SampleCallbackUrl"`
	SampleLoadPath    string `yaml:"SampleLoadPath"`
	MtxTimeout        int    `yaml:"MtxTimeout"`
}

// JMR
type JMR struct {
	CommandUrl          string `yaml:"CommandUrl"`
	MatchItemNumPerJSON int    `yaml:"MatchItemNumPerJSON"`
	MtxPath             string `yaml:"MtxPath"`
	DownloadRoutineNum  int    `yaml:"DownloadRoutineNum"`
	Switch              bool   `yaml:"Switch"`
	AuthUser            string `yaml:"AuthUser"`
	AuthPasswd          string `yaml:"AuthPasswd"`
	DeviceStatusUrl     string `yaml:"DeviceStatusUrl"`
	AuthUrl             string `yaml:"AuthUrl"`
	DbUpdateUrl         string `yaml:"DbUpdateUrl"`
	SnortPath           string `yaml:"SnortPath"`
	DownloadPath        string `yaml:"DownloadPath"`
	DeviceInfoUrl       string `yaml:"DeviceInfoUrl"`
}

// RULE
type RULE struct {
	SnortPath          string `yaml:"SnortPath"`
	SnortImportUrl     string `yaml:"SnortImportUrl"`
	Command            string `yaml:"Command"`
	SnortDeleteUrl     string `yaml:"SnortDeleteUrl"`
	TokenSwitch        bool   `yaml:"TokenSwitch"`
	TokenUrl           string `yaml:"TokenUrl"`
	AppId              string `yaml:"AppId"`
	AppKey             string `yaml:"AppKey"`
	SnortSwitch        bool   `yaml:"SnortSwitch"`
	SnortFileNameDef   bool   `yaml:"SnortFileNameDef"`
	SnortDeleteBulkUrl string `yaml:"SnortDeleteBulkUrl"`
	SnortSyncUrl       string `yaml:"SnortSyncUrl"`
	SnortFileName      string `yaml:"SnortFileName"`
	SnortSyncFile      string `yaml:"SnortSyncFile"`
	CommandWait        int    `yaml:"CommandWait"`
}

// POLICY
type POLICY struct {
	FtpUrlFile    string `yaml:"FtpUrlFile"`
	TaskPeriod    string `yaml:"TaskPeriod"`
	IsSftp        bool   `yaml:"IsSftp"`
	Password      string `yaml:"Password"`
	FtpServerIP   string `yaml:"FtpServerIP"`
	FtpServerPort int    `yaml:"FtpServerPort"`
	Switch        bool   `yaml:"Switch"`
	User          string `yaml:"User"`
	FtpDomainFile string `yaml:"FtpDomainFile"`
}

// MALWARE
type MALWARE struct {
	CmccDefaultRedirURL   string  `yaml:"CmccDefaultRedirURL"`
	CmccComDevURI         string  `yaml:"CmccComDevURI"`
	GLDevIP               string  `yaml:"GLDevIP"`
	DevVersion            string  `yaml:"DevVersion"`
	HostStripperSwitch    bool    `yaml:"HostStripperSwitch"`
	DeleteFilePeriod      int     `yaml:"DeleteFilePeriod"`
	YaraPath              string  `yaml:"YaraPath"`
	PolicyMetrics         string  `yaml:"Policy_Metrics"`
	PolicyBackPath        string  `yaml:"PolicyBackPath"`
	ReportMetrics         string  `yaml:"Report_Metrics"`
	DevType               string  `yaml:"DevType"`
	Operator              int     `yaml:"Operator"`
	DownloadPath          string  `yaml:"DownloadPath"`
	UseDefaultPeriod      bool    `yaml:"UseDefaultPeriod"`
	DbFileProcRoutineNum  int     `yaml:"DbFileProcRoutineNum"`
	DownloadRoutineNum    int     `yaml:"DownloadRoutineNum"`
	DevNum                string  `yaml:"DevNum"`
	KeepAlivePort         int     `yaml:"KeepAlivePort"`
	Version               float64 `yaml:"Version"`
	MtxPath               string  `yaml:"MtxPath"`
	MtxCapSwitch          bool    `yaml:"MtxCapSwitch"`
	MtxCapFeature         string  `yaml:"MtxCapFeature"`
	TzRuleMode            int     `yaml:"TzRuleMode"`
	PolicyPath            string  `yaml:"PolicyPath"`
	ReportStatusPeriod    int     `yaml:"ReportStatusPeriod"`
	SnortCapFeature       string  `yaml:"SnortCapFeature"`
	TzRuleAllMap          int     `yaml:"TzRuleAllMap"`
	CmccMalwarePktURI     string  `yaml:"CmccMalwarePktURI"`
	UpdateURI             string  `yaml:"UpdateURI"`
	SaveDownLoadFileNum   int     `yaml:"SaveDownLoadFileNum"`
	TzRuleAllFile         string  `yaml:"TzRuleAllFile"`
	UpgradeReportURI      string  `yaml:"UpgradeReportURI"`
	SnortPath             string  `yaml:"SnortPath"`
	TzExtPath             string  `yaml:"TzExtPath"`
	DevVendor             string  `yaml:"DevVendor"`
	ProvinceNum           string  `yaml:"ProvinceNum"`
	AliveCheckURI         string  `yaml:"AliveCheckURI"`
	StatusReportURI       string  `yaml:"StatusReportURI"`
	LocalFeatureDir       string  `yaml:"LocalFeatureDir"`
	SavePolicyTime        int     `yaml:"SavePolicyTime"`
	KeepAlivePeriod       int     `yaml:"KeepAlivePeriod"`
	SnortCapSwitch        bool    `yaml:"SnortCapSwitch"`
	MalwareSwitch         bool    `yaml:"MalwareSwitch"`
	CmccMalwareFileURI    string  `yaml:"CmccMalwareFileURI"`
	DevTypeInt            int     `yaml:"DevTypeInt"`
	UseRelaPath           bool    `yaml:"UseRelaPath"`
	SpecVersion           int     `yaml:"SpecVersion"`
	Disable86Suffix       bool    `yaml:"Disable86Suffix"`
	EnableUrlBase64Decode bool    `yaml:"EnableUrlBase64Decode"`
}

// HTTP
type HTTP struct {
	LocalListenIP       string   `yaml:"LocalListenIP"`
	JmrRemoteServerIP   string   `yaml:"JmrRemoteServerIP"`
	ServerCert          string   `yaml:"ServerCert"`
	SlaveServerInfo     []string `yaml:"SlaveServerInfo"`
	DumpRequestSwtch    bool     `yaml:"DumpRequestSwtch"`
	JmrRemoteServerPort int      `yaml:"JmrRemoteServerPort"`
	JmrLocalListenPort  int      `yaml:"JmrLocalListenPort"`
	JmrLocalListenIP    string   `yaml:"JmrLocalListenIP"`
	EnableSSL           bool     `yaml:"EnableSSL"`
	ServerKey           string   `yaml:"ServerKey"`
	WorkMode            int      `yaml:"WorkMode"`
	LocalListenPort     int      `yaml:"LocalListenPort"`
	HttpUserName        string   `yaml:"HttpUserName"`
	HttpPassWord        string   `yaml:"HttpPassWord"`
	RemoteServerPort    int      `yaml:"RemoteServerPort"`
	RemoteServerIP      string   `yaml:"RemoteServerIP"`
	CACert              string   `yaml:"CACert"`
	Token               string   `yaml:"Token"`
	AuthAlgo            int      `yaml:"AuthAlgo"`
	AuthPeriod          int      `yaml:"AuthPeriod"`
}

// LOG
type LOG struct {
	LogDir          string `yaml:"LogDir"`
	LogUploadJSON   string `yaml:"LogUploadJSON"`
	LogWriteSwitch  int    `yaml:"LogWriteSwitch"`
	LogWriteRoutine int    `yaml:"LogWriteRoutine"`
	LogRecordCount  int    `yaml:"LogRecordCount"`
	LogFileSize     int    `yaml:"LogFileSize"`
	LogWritePeriod  int    `yaml:"LogWritePeriod"`
}

// stConfig
type CfgUpp struct {
	SYSTEM  SYSTEM  `yaml:"SYSTEM"`
	HTTP    HTTP    `yaml:"HTTP"`
	MALWARE MALWARE `yaml:"MALWARE"`
	LOG     LOG     `yaml:"LOG"`
	RULE    RULE    `yaml:"RULE"`
	POLICY  POLICY  `yaml:"POLICY"`
	YGSH    YGSH    `yaml:"YGSH"`
	MTX     MTX     `yaml:"MTX"`
	JMR     JMR     `yaml:"JMR"`
}

func uppYaml2Conf(path string) (CfgUpp, error) {
	viper.SetConfigFile(path)

	var config CfgUpp
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return config, err
	}

	return config, nil
}

func MergeUppCfg(tpl, use string) (data []byte, err error) {
	tplConf, err := uppYaml2Conf(tpl)
	if err != nil {
		fmt.Printf("transfer tpl yaml to conf failed: %v\n", err)
		return data, err
	}
	useConf, err := uppYaml2Conf(use)
	if err != nil {
		fmt.Printf("transfer use yaml to conf failed: %v\n", err)
		return data, err
	}

	if err := mergo.Merge(&useConf, tplConf); err != nil {
		fmt.Printf("merge tpl to use yaml failed: %v\n", err)
		return data, err
	}

	return yaml.Marshal(&useConf)
}

func ModifyUppCfg(key, value, module, filename string) (data []byte, err error) {
	cfg, err := uppYaml2Conf(filename)
	if err != nil {
		fmt.Printf("yaml to struct failed: %v\n", err)
		return data, err
	}
	err = field.SetField(&cfg, key, value)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return data, err
	} else {
		//fmt.Printf("change [%s] key [%s] value to [%s] succ!\n", module, key, value)
	}

	data, err = yaml.Marshal(&cfg)
	if err != nil {
		fmt.Printf("marshal yaml file failed: %v\n", err)
		return data, err
	}

	return data, err
}
