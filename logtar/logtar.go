package logtar

import (
	"fmt"

	"modify/field"

	"github.com/imdario/mergo"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// FieldsIndexInfo
type FieldsIndexInfo struct {
	MsgType int `yaml:"MsgType"`
	MD5     int `yaml:"MD5"`
}

// WatchGroups
type WatchGroups struct {
	Suffix       string          `yaml:"Suffix"`
	SamplePath   string          `yaml:"SamplePath"`
	SampleCopy   bool            `yaml:"SampleCopy"`
	LogType      int             `yaml:"LogType"`
	CompressType int             `yaml:"CompressType"`
	FieldsCount  int             `yaml:"FieldsCount"`
	LogPath      string          `yaml:"LogPath"`
	FieldsIndex  FieldsIndexInfo `yaml:"FieldsIndex"`
}

// LOG
type LOG struct {
	WriteBufferSize    int           `yaml:"WriteBufferSize"`
	LogRecordCountType int           `yaml:"LogRecordCountType"`
	LogRecordCount     int           `yaml:"LogRecordCount"`
	LogWritePeriod     int           `yaml:"LogWritePeriod"`
	WatchGroups        []WatchGroups `yaml:"WatchGroups"`
}

// BASE
type BASE struct {
	DeviceID           string `yaml:"DeviceID"`
	DepSampleTime      int    `yaml:"DepSampleTime"`
	BackupLogSwitch    bool   `yaml:"BackupLogSwitch"`
	CMDPort            int    `yaml:"CMDPort"`
	Manufactor         string `yaml:"Manufactor"`
	CompressRoutineNum int    `yaml:"CompressRoutineNum"`
	BackupLogSave      string `yaml:"BackupLogSave"`
	DepSampleSwitch    bool   `yaml:"DepSampleSwitch"`
	BackupLogConf      string `yaml:"BackupLogConf"`
	LogLevel           int    `yaml:"LogLevel"`
	LogProcRoutineNum  int    `yaml:"LogProcRoutineNum"`
	WorkMode           int    `yaml:"WorkMode"`
	DepSampleCap       int    `yaml:"DepSampleCap"`
}

type LogtarCfg struct {
	BASE BASE `yaml:"BASE"`
	LOG  LOG  `yaml:"LOG"`
}

func logtarYaml2Conf(path string) (LogtarCfg, error) {
	viper.SetConfigFile(path)

	var config LogtarCfg
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

func MergeLogtarCfg(tpl, use string) (data []byte, err error) {
	tplConf, err := logtarYaml2Conf(tpl)
	if err != nil {
		fmt.Printf("transfer tpl yaml to conf failed: %v\n", err)
		return data, err
	}
	useConf, err := logtarYaml2Conf(use)
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

func ModifyLogtarCfg(key, value, module, filename string) (data []byte, err error) {
	cfg, err := logtarYaml2Conf(filename)
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
