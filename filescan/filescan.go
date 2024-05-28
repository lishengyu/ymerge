package filescan

import (
	"fmt"

	"modify/field"

	"github.com/imdario/mergo"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// Base
type Base struct {
	DeepInspect       bool    `yaml:"DeepInspect"`
	CMDPort           int     `yaml:"CMDPort"`
	CommandID         string  `yaml:"CommandID"`
	LogIDReorder      bool    `yaml:"LogIDReorder"`
	ForceUploadFile   bool    `yaml:"ForceUploadFile"`
	RecogRatio        float64 `yaml:"RecogRatio"`
	IgnoreNormalInfo  bool    `yaml:"IgnoreNormalInfo"`
	DsThreshold       int     `yaml:"DsThreshold"`
	DsLogFilter       bool    `yaml:"DsLogFilter"`
	Debug             bool    `yaml:"Debug"`
	DeviceNo          string  `yaml:"DeviceNo"`
	LogProcNum        int     `yaml:"LogProcNum"`
	RecordUnrecogLog  bool    `yaml:"RecordUnrecogLog"`
	PathReadPeriod    int     `yaml:"PathReadPeriod"`
	EnableOcr         int     `yaml:"EnableOcr"`
	TargetIP          string  `yaml:"TargetIP"`
	DeployMent        string  `yaml:"DeployMent"`
	Manufactor        string  `yaml:"Manufactor"`
	DecompressTempDir string  `yaml:"DecompressTempDir"`
	WorkMode          int     `yaml:"WorkMode"`
	MimeCheck         bool    `yaml:"MimeCheck"`
}

// WatchGroupsFieldsIndex
type WatchGroupsFieldsIndex struct {
	ApplicationProto int `yaml:"ApplicationProto,omitempty"`
	AssetsNum        int `yaml:"AssetsNum,omitempty"`
	AttachMent       int `yaml:"AttachMent,omitempty"`
	BusinessProto    int `yaml:"BusinessProto,omitempty"`
	CommandID        int `yaml:"CommandID,omitempty"`
	DataContent      int `yaml:"DataContent,omitempty"`
	DataInfoNum      int `yaml:"DataInfoNum,omitempty"`
	DataLevel        int `yaml:"DataLevel,omitempty"`
	DataNum          int `yaml:"DataNum,omitempty"`
	DataType         int `yaml:"DataType,omitempty"`
	DestIP           int `yaml:"DestIP,omitempty"`
	DestPort         int `yaml:"DestPort,omitempty"`
	Domain           int `yaml:"Domain,omitempty"`
	EventSubType     int `yaml:"EventSubType,omitempty"`
	EventTypeId      int `yaml:"EventTypeId,omitempty"`
	FileSize         int `yaml:"FileSize,omitempty"`
	FileType         int `yaml:"FileType,omitempty"`
	HouseID          int `yaml:"HouseID,omitempty"`
	IsMatchEvent     int `yaml:"IsMatchEvent,omitempty"`
	IsUploadFile     int `yaml:"IsUploadFile,omitempty"`
	LogID            int `yaml:"LogID,omitempty"`
	Proto            int `yaml:"Proto,omitempty"`
	ProtocolType     int `yaml:"ProtocolType,omitempty"`
	RuleDes          int `yaml:"RuleDes,omitempty"`
	RuleID           int `yaml:"RuleID,omitempty"`
	SrcIP            int `yaml:"SrcIP,omitempty"`
	TimeStamp        int `yaml:"TimeStamp,omitempty"`
	Url              int `yaml:"Url,omitempty"`
}

// WatchGroups
type WatchGroups struct {
	UploadPath   string                 `yaml:"UploadPath"`
	FieldsCount  int                    `yaml:"FieldsCount"`
	LogPath      string                 `yaml:"LogPath"`
	SamplePath   string                 `yaml:"SamplePath"`
	CheckSamples bool                   `yaml:"CheckSamples"`
	Suffix       string                 `yaml:"Suffix"`
	LogType      int                    `yaml:"LogType"`
	FieldsIndex  WatchGroupsFieldsIndex `yaml:"FieldsIndex"`
}

// Log
type Log struct {
	LogZipCount          int           `yaml:"LogZipCount"`
	DedupSwitch          bool          `yaml:"DedupSwitch"`
	DsHighFreqStatPeriod int           `yaml:"DsHighFreqStatPeriod"`
	AttachName           string        `yaml:"AttachName"`
	UnRecog06c0LogName   string        `yaml:"UnRecog06c0LogName"`
	FileScanCount        int           `yaml:"FileScanCount"`
	DsLogZipSwitch       bool          `yaml:"DsLogZipSwitch"`
	DedupPeriod          int           `yaml:"DedupPeriod"`
	IsDiffer             bool          `yaml:"IsDiffer"`
	EnableSampleSuffix   bool          `yaml:"EnableSampleSuffix"`
	WatchGroups          []WatchGroups `yaml:"WatchGroups"`
	RuleLogName          string        `yaml:"RuleLogName"`
	UnRecog06c1LogName   string        `yaml:"UnRecog06c1LogName"`
	MonitorLogPath       string        `yaml:"MonitorLogPath"`
	AttachZipSuffix      bool          `yaml:"AttachZipSuffix"`
	RuleLogPath          string        `yaml:"RuleLogPath"`
	HighFreqLogName      string        `yaml:"HighFreqLogName"`
}

// FsCfg
type FsCfg struct {
	Base Base `yaml:"Base"`
	Log  Log  `yaml:"Log"`
}

func fsYaml2Conf(path string) (FsCfg, error) {
	viper.SetConfigFile(path)

	var config FsCfg
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

func MergeFilescanCfg(tpl, use string) (data []byte, err error) {
	tplConf, err := fsYaml2Conf(tpl)
	if err != nil {
		fmt.Printf("transfer tpl yaml to conf failed: %v\n", err)
		return data, err
	}
	useConf, err := fsYaml2Conf(use)
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

func ModifyFilescanCfg(key, value, module, filename string) (data []byte, err error) {
	cfg, err := fsYaml2Conf(filename)
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
